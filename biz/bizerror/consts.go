package bizerror

type ErrorLevel int

const (
	Toast ErrorLevel = iota
	Block
	Custom
	Ignore
)

type ErrorCode int

const (
	UnKnownErrorCode               ErrorCode = -1
	DBErrorCode                    ErrorCode = -8001
	InvalidRequestErrorCode        ErrorCode = -8002
	InternalErrorCode              ErrorCode = -8003
	InvalidTaskIDErrorCode         ErrorCode = -8004
	DBRowRowsAffectedZeroErrorCode ErrorCode = -8005
	TypeAssertionErrorCode         ErrorCode = -8006
	InvalidParamsErrorCode         ErrorCode = -8007
	ThirdPartyErrorCode            ErrorCode = -8008
	UnauthorizedErrorCode          ErrorCode = -8401
)

var ErrorCode2String = map[ErrorCode]string{
	UnKnownErrorCode:               "-1",
	DBErrorCode:                    "-8001",
	InvalidRequestErrorCode:        "-8002",
	InternalErrorCode:              "-8003",
	InvalidTaskIDErrorCode:         "-8004",
	DBRowRowsAffectedZeroErrorCode: "-8005",
	TypeAssertionErrorCode:         "-8006",
	InvalidParamsErrorCode:         "-8007",
	UnauthorizedErrorCode:          "-8401",
}

var DBError = New(Toast, DBErrorCode, "Internal database error")
var InvalidSignatureError = New(Toast, UnauthorizedErrorCode, "Signature verification failed")
var InvalidRequestError = New(Toast, InvalidRequestErrorCode, "Request parameters are invalid")
var UnauthorizedError = New(Ignore, UnauthorizedErrorCode, "User request authentication failed")
var InternalError = New(Block, InternalErrorCode, "Internal error")
var TypeAssertionError = New(Block, TypeAssertionErrorCode, "Type assertion error")
var InvalidParamsError = New(Block, InvalidParamsErrorCode, "illegal parameter")
