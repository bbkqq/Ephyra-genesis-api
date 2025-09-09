package web3

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"strconv"
	"strings"
)

func AddressFormat(addr string) string {
	return strings.ToLower(addr)
}

func CompareEthAddress(addr1 string, addr2 string) bool {
	return AddressFormat(addr1) == AddressFormat(addr2)
}

func GetTransactionReceiptStatus(ctx context.Context, client *ethclient.Client, txHashStr string) int {
	// 返回 0 表示合约执行失败
	// 返回 1 表示合约执行成功
	// 返回 -1 表示交易还没找到记录

	// 交易哈希
	txHash := common.HexToHash(txHashStr)

	// 获取交易回执
	receipt, err := client.TransactionReceipt(ctx, txHash)
	if err != nil {
		hlog.CtxErrorf(ctx, "Failed to get transaction receipt: %v", err)
		return -1
	}

	// 检查交易状态
	if receipt.Status == 1 {
		return 1
	} else {
		return 0
	}
}

func VerifyEthSignature(ctx context.Context, message, signature, expectedAddress string) (bool, error) {

	// Hash the unsigned message using EIP-191
	hashedMessage := []byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(message)) + message)
	hash := crypto.Keccak256Hash(hashedMessage)

	// Get the bytes of the signed message
	decodedMessage, err := hexutil.Decode(signature)
	if err != nil {
		hlog.CtxErrorf(ctx, "Failed to decode signed message: %v", err)
		return false, err
	}

	// Handles cases where EIP-115 is not implemented (most wallets don't implement it)
	if len(decodedMessage) != 65 {
		return false, errors.New("invalid signature length")
	}

	if decodedMessage[64] == 27 || decodedMessage[64] == 28 {
		decodedMessage[64] -= 27
	}

	// Recover a public key from the signed message
	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), decodedMessage)
	if err != nil {
		hlog.CtxErrorf(ctx, "Failed to recover public key: %v", err)
		return false, err
	}

	if sigPublicKeyECDSA == nil {
		return false, errors.New("invalid public key")
	}

	recoveredAddr := crypto.PubkeyToAddress(*sigPublicKeyECDSA).String()
	return CompareEthAddress(recoveredAddr, expectedAddress), nil
}
