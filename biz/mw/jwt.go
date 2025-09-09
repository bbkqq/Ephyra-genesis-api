package jwt

import (
	"Ephyra-genesis-api/biz/bizerror"
	"Ephyra-genesis-api/biz/dal/redis"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	"net/http"
	"time"

	"Ephyra-genesis-api/biz/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/cast"
)

type KeyVersion string

const (
	BscVersion KeyVersion = "v1"
)

var (
	keyVersions = map[string][]byte{
		string(BscVersion): []byte("182hhgu-i1doi12biubc"),
	}
	JwtMiddleware              = &jWTMiddleware{SigningAlgorithm: "HS256", KeyVersion: keyVersions}
	errInvalidSigningAlgorithm = errors.New("invalid signing algorithm")
	errInvalidHeaderVersion    = errors.New("invalid header version")
	errEmptyHeader             = errors.New("header is empty")
	authHeaderKey              = "Authorization"
	tokenNotEqualErr           = errors.New("token not equal")
)

const jwtExpiredTime = time.Duration(time.Hour * 6)

type jWTMiddleware struct {
	SigningAlgorithm string
	// Secret key used for signing. Required.
	Key        []byte
	KeyVersion map[string][]byte
	// ParseOptions allow to modify jwt's parser methods
	ParseOptions []jwt.ParserOption
}

type MapClaims map[string]interface{}

func NewJWTPayload(uid int64, address string, registerTime int64, chain string) MapClaims {
	payload := MapClaims{}
	payload["user_id"] = uid
	payload["address"] = address
	payload["register_time"] = registerTime
	payload["chain"] = chain
	return payload
}

func (mw *jWTMiddleware) MiddlewareFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		err := mw.JwtVerify(ctx, c)
		if err != nil {
			mw.unauthorized(ctx, c)
			return
		}
		c.Next(ctx)
	}
}

func (mw *jWTMiddleware) JwtVerify(ctx context.Context, c *app.RequestContext) error {
	authToken, err := mw.getKeyFromHeader(ctx, c, authHeaderKey)
	// 先从header中获取
	if err != nil {
		// header 中不存在从 url query 获取
		authToken, err = mw.getKeyFromURLQuery(ctx, c, authHeaderKey)
		if err != nil {
			hlog.CtxWarnf(ctx, "[JWTMiddleware] get jwt headerkey error: %v", err)
			return err
		}
	}

	jwtToken, err := jwt.Parse(authToken, func(t *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(mw.SigningAlgorithm) != t.Method {
			return nil, errInvalidSigningAlgorithm
		}
		// save token string if valid
		// c.Set("JWT_TOKEN", authToken)
		v, ok := t.Header["version"]
		// version not in header
		if !ok {
			return nil, errInvalidHeaderVersion
		}
		vr := cast.ToString(v)
		key, ok := mw.KeyVersion[vr]
		if !ok {
			return nil, errInvalidHeaderVersion
		}
		return key, nil
	}, mw.ParseOptions...)

	if err != nil {
		hlog.CtxErrorf(ctx, "[JWTMiddleware] parse jwt data error: %v", err)
		return err
	}

	payload := MapClaims{}
	for key, value := range jwtToken.Claims.(jwt.MapClaims) {
		payload[key] = value
	}
	// check address
	payloadAddress := cast.ToString(payload["address"])
	payloadUserId := cast.ToInt64(payload["user_id"])

	// check uid & address in redis
	remoteTokenMD5, _, err := mw.getRemoteJwtToken(ctx, payloadUserId, payloadAddress)
	if err != nil {
		hlog.CtxWarnf(ctx, "[JWTMiddleware] get login token error: %v", err)
		return err
	}

	if remoteTokenMD5 != utils.MD5(authToken) {
		hlog.CtxWarnf(ctx, "[JWTMiddleware] jwt token not equal with remote")
		return tokenNotEqualErr
	}

	//if !inRedis {
	//	mw.SetJwtToken(ctx, payloadUserId, payloadAddress, remoteTokenMD5)
	//}

	// 对 redis 中的 jwt token 续期
	mw.renewalJwtToken(ctx, payloadUserId, payloadAddress)

	c.Set("address", payloadAddress)
	c.Set("user_id", payloadUserId)
	c.Set("user_id_str", cast.ToString(payloadUserId))
	c.Set("chain", cast.ToString(payload["chain"]))
	c.Set("referrer_id", payload["referrer_id"])

	return nil
}

func (mw *jWTMiddleware) unauthorized(ctx context.Context, c *app.RequestContext) {
	c.Abort()
	utils.SendErrResponse(ctx, c, http.StatusOK, bizerror.UnauthorizedError)
}

func (mw *jWTMiddleware) getKeyFromHeader(_ context.Context, c *app.RequestContext, key string) (string, error) {
	authHeader := c.Request.Header.Get(key)

	if authHeader == "" {
		return "", errEmptyHeader
	}
	return authHeader, nil
}

func (mw *jWTMiddleware) getKeyFromURLQuery(_ context.Context, c *app.RequestContext, key string) (string, error) {
	authHeader := c.Query(key)

	if authHeader == "" {
		return "", errEmptyHeader
	}
	return authHeader, nil
}

func (mw *jWTMiddleware) CreateTokenWithVersion(version KeyVersion, payload MapClaims) (string, error) {
	// Create the token
	token := jwt.New(jwt.GetSigningMethod(mw.SigningAlgorithm))
	claims := token.Claims.(jwt.MapClaims)

	for key, value := range payload {
		claims[key] = value
	}

	// set key version to header
	token.Header["version"] = string(version)
	key, ok := mw.KeyVersion[string(version)]
	if !ok {
		return "", errors.New("unknown key version")
	}

	claims["login_timestamp"] = time.Now().Unix()
	claims["seed"] = utils.RandomString(10)
	tokenString, err := token.SignedString(key)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (mw *jWTMiddleware) getRemoteJwtToken(ctx context.Context, userId int64, address string) (string, bool, error) {
	key := redis.UserLoginTokenKey(userId, address)
	remoteToken, err := redis.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return "", false, err
	}
	if len(remoteToken) == 0 {
		return "", false, errors.New("remoteToken empty error")
	}
	return remoteToken, true, nil
}

func (mw *jWTMiddleware) SetJwtToken(ctx context.Context, userId int64, address string, token string) (string, error) {
	key := redis.UserLoginTokenKey(userId, address)
	return redis.RedisClient.Set(ctx, key, token, jwtExpiredTime).Result()
}

func (mw *jWTMiddleware) renewalJwtToken(ctx context.Context, userId int64, address string) {
	key := redis.UserLoginTokenKey(userId, address)
	redis.RedisClient.Expire(ctx, key, jwtExpiredTime)
}
