package jobs

import (
	answer "Ephyra-genesis-api/biz/contracts"
	"Ephyra-genesis-api/biz/dal"
	"Ephyra-genesis-api/biz/dal/redis"
	"Ephyra-genesis-api/conf"
	"context"
	"math"
	"time"

	"github.com/bytedance/gopkg/util/gopool"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cast"
)

const (
	confirmedBlockNum = 6
	defaultBlockRange = 80
)

type AnswerEventListenerJob struct {
	startBlock            uint64
	answerContract        *answer.Answer
	answerContractAddress string
	ethClient             *ethclient.Client
	workers               gopool.Pool
}

func NewAnswerEventListenerJob() *AnswerEventListenerJob {

	c := conf.GetConf().Contract[0]
	cAddress := common.HexToAddress(c.Address)

	ethClient, err := ethclient.Dial(c.RpcUrl)
	if err != nil {
		panic(err)
	}

	TokenBridge, err := answer.NewAnswer(cAddress, ethClient)
	if err != nil {
		panic(err)
	}

	return &AnswerEventListenerJob{
		startBlock:            0,
		answerContract:        TokenBridge,
		answerContractAddress: c.Address,
		ethClient:             ethClient,
		workers:               gopool.NewPool("AnswerEventListenerJob-workers", 100, gopool.NewConfig()),
	}
}

func (c *AnswerEventListenerJob) Do(ctx context.Context) {
	hlog.CtxInfof(ctx, "[AnswerEventListenerJob] start")

	var err error
	lastListenOnBlockNumberKey := redis.ChainEventListenKey(c.answerContractAddress)
	if c.startBlock == 0 {
		// 从 redis 获取上一次访问的Block
		ret, err := redis.RedisClient.Get(ctx, lastListenOnBlockNumberKey).Result()
		if err == nil {
			c.startBlock = cast.ToUint64(ret)
		}
	}
	endBlockNum, err := c.ethClient.BlockNumber(ctx)
	if err != nil || endBlockNum > math.MaxUint32 || endBlockNum == 0 {
		// 获取区块失败
		hlog.CtxErrorf(ctx, "[AnswerEventListenerJob] contract address: %v, get block number failed, err: %v, latest block number: %v", c.answerContractAddress, err, endBlockNum)
		return
	}
	hlog.CtxInfof(ctx, "[AnswerEventListenerJob] get end block: %v", endBlockNum)

	// 还是等于0
	if c.startBlock == 0 {
		c.startBlock = endBlockNum - defaultBlockRange
	}

	endBlockNum -= 2

	// 开始区块和结束区块相同时先退出等待下一轮
	if c.startBlock >= endBlockNum {
		return
	}

	if endBlockNum-c.startBlock > defaultBlockRange {
		start := c.startBlock
		for start <= endBlockNum {
			end := start + defaultBlockRange - 1
			if end > endBlockNum {
				end = endBlockNum
			}
			hlog.CtxInfof(ctx, "[AnswerEventListenerJob] contract address: %v, start listen on block: %v, end block: %v", c.answerContractAddress, start, end)

			opts := bind.FilterOpts{
				Start:   start,
				End:     &end,
				Context: ctx,
			}

			err = c.processEvent(ctx, &opts)
			if err != nil {
				return
			}

			start = end + 1
			c.startBlock = start

			hlog.CtxInfof(ctx, "[AnswerEventListenerJob] contract address: %v, update last listen on block number: %d", c.answerContractAddress, c.startBlock)
			_, _ = redis.RedisClient.Set(ctx, lastListenOnBlockNumberKey, c.startBlock, 0).Result()
			time.Sleep(100 * time.Millisecond)
		}
	} else {
		endBlockOffset := endBlockNum - defaultBlockRange
		hlog.CtxInfof(ctx, "[AnswerEventListenerJob] contract address: %v, start listen on block: %v, end block: %v", c.answerContractAddress, c.startBlock, endBlockNum)
		opts := bind.FilterOpts{
			Start:   c.startBlock,
			End:     &endBlockNum,
			Context: ctx,
		}

		err = c.processEvent(ctx, &opts)
		if err != nil {
			return
		}

		c.startBlock = max(c.startBlock, endBlockOffset) + 1
		hlog.CtxInfof(ctx, "[AnswerEventListenerJob] contract address: %v, update last listen on block number: %d", c.answerContractAddress, c.startBlock)
		_, _ = redis.RedisClient.Set(ctx, lastListenOnBlockNumberKey, c.startBlock, 0).Result()
	}

	hlog.CtxInfof(ctx, "[AnswerEventListenerJob] end")
}

func (c *AnswerEventListenerJob) processEvent(ctx context.Context, opts *bind.FilterOpts) error {
	eventIterator, err := c.answerContract.FilterAnswerSubmitted(opts, nil)
	if err != nil {
		hlog.CtxErrorf(ctx, "[AnswerEventListenerJob] address: %v, FilterUserClaimEvent error: %v", c.answerContractAddress, err)
		return err
	}

	for {
		ret := eventIterator.Next()
		if !ret {
			break
		}

		event_ := eventIterator.Event
		if event_.Raw.BlockNumber > c.startBlock {
			c.startBlock = event_.Raw.BlockNumber
		}

		// 处理链上答案提交事件
		hlog.CtxInfof(ctx, "[AnswerEventListenerJob] Processing answer event: user=%s, questionId=%d, day=%d, txHash=%s",
			event_.User.Hex(), event_.QuestionId.Int64(), event_.Day.Int64(), event_.Raw.TxHash.Hex())

		err = dal.ProcessAnswerOnChainEvent(
			ctx,
			event_.User.Hex(),
			event_.Answer,
			event_.QuestionId.Int64(),
			event_.Timestamp.Int64(),
			event_.Day.Int64(),
			event_.Raw.TxHash.Hex())

		if err != nil {
			hlog.CtxErrorf(ctx, "[AnswerEventListenerJob] ProcessAnswerOnChainEvent err: %v", err)
			return err
		}
	}

	return nil
}

func (c *AnswerEventListenerJob) Interval() time.Duration {
	return 1 * time.Second
}

func (c *AnswerEventListenerJob) DisableSerializable() bool {
	return false
}
