// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package answer

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// QuestionAnswerContractAnswerRecord is an auto generated low-level Go binding around an user-defined struct.
type QuestionAnswerContractAnswerRecord struct {
	User       common.Address
	Answer     string
	QuestionId *big.Int
	Timestamp  *big.Int
	Day        *big.Int
}

// AnswerMetaData contains all meta data concerning the Answer contract.
var AnswerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"questionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"name\":\"AlreadySubmittedToday\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyAnswer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidQuestionId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"answer\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"questionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"name\":\"AnswerSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"questionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"name\":\"DailyLimitReached\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"SBTMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"answerRecords\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"answer\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"questionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dailySubmissions\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"recordId\",\"type\":\"uint256\"}],\"name\":\"getAnswerRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"answer\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"questionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"internalType\":\"structQuestionAnswerContract.AnswerRecord\",\"name\":\"record\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentDayOffset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"questionId\",\"type\":\"uint256\"}],\"name\":\"getQuestionAnswerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalAnswerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserAnswerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasMinted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"hasSBT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"questionId\",\"type\":\"uint256\"}],\"name\":\"hasSubmittedOnDay\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"questionId\",\"type\":\"uint256\"}],\"name\":\"hasSubmittedToday\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"questionAnswerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"answer\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"submit_answer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userAnswerHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AnswerABI is the input ABI used to generate the binding from.
// Deprecated: Use AnswerMetaData.ABI instead.
var AnswerABI = AnswerMetaData.ABI

// Answer is an auto generated Go binding around an Ethereum contract.
type Answer struct {
	AnswerCaller     // Read-only binding to the contract
	AnswerTransactor // Write-only binding to the contract
	AnswerFilterer   // Log filterer for contract events
}

// AnswerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AnswerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnswerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AnswerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnswerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AnswerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnswerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AnswerSession struct {
	Contract     *Answer           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AnswerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AnswerCallerSession struct {
	Contract *AnswerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AnswerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AnswerTransactorSession struct {
	Contract     *AnswerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AnswerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AnswerRaw struct {
	Contract *Answer // Generic contract binding to access the raw methods on
}

// AnswerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AnswerCallerRaw struct {
	Contract *AnswerCaller // Generic read-only contract binding to access the raw methods on
}

// AnswerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AnswerTransactorRaw struct {
	Contract *AnswerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAnswer creates a new instance of Answer, bound to a specific deployed contract.
func NewAnswer(address common.Address, backend bind.ContractBackend) (*Answer, error) {
	contract, err := bindAnswer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Answer{AnswerCaller: AnswerCaller{contract: contract}, AnswerTransactor: AnswerTransactor{contract: contract}, AnswerFilterer: AnswerFilterer{contract: contract}}, nil
}

// NewAnswerCaller creates a new read-only instance of Answer, bound to a specific deployed contract.
func NewAnswerCaller(address common.Address, caller bind.ContractCaller) (*AnswerCaller, error) {
	contract, err := bindAnswer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AnswerCaller{contract: contract}, nil
}

// NewAnswerTransactor creates a new write-only instance of Answer, bound to a specific deployed contract.
func NewAnswerTransactor(address common.Address, transactor bind.ContractTransactor) (*AnswerTransactor, error) {
	contract, err := bindAnswer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AnswerTransactor{contract: contract}, nil
}

// NewAnswerFilterer creates a new log filterer instance of Answer, bound to a specific deployed contract.
func NewAnswerFilterer(address common.Address, filterer bind.ContractFilterer) (*AnswerFilterer, error) {
	contract, err := bindAnswer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AnswerFilterer{contract: contract}, nil
}

// bindAnswer binds a generic wrapper to an already deployed contract.
func bindAnswer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AnswerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Answer *AnswerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Answer.Contract.AnswerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Answer *AnswerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Answer.Contract.AnswerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Answer *AnswerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Answer.Contract.AnswerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Answer *AnswerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Answer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Answer *AnswerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Answer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Answer *AnswerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Answer.Contract.contract.Transact(opts, method, params...)
}

// AnswerRecords is a free data retrieval call binding the contract method 0x7c3884f1.
//
// Solidity: function answerRecords(uint256 ) view returns(address user, string answer, uint256 questionId, uint256 timestamp, uint256 day)
func (_Answer *AnswerCaller) AnswerRecords(opts *bind.CallOpts, arg0 *big.Int) (struct {
	User       common.Address
	Answer     string
	QuestionId *big.Int
	Timestamp  *big.Int
	Day        *big.Int
}, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "answerRecords", arg0)

	outstruct := new(struct {
		User       common.Address
		Answer     string
		QuestionId *big.Int
		Timestamp  *big.Int
		Day        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Answer = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.QuestionId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Day = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AnswerRecords is a free data retrieval call binding the contract method 0x7c3884f1.
//
// Solidity: function answerRecords(uint256 ) view returns(address user, string answer, uint256 questionId, uint256 timestamp, uint256 day)
func (_Answer *AnswerSession) AnswerRecords(arg0 *big.Int) (struct {
	User       common.Address
	Answer     string
	QuestionId *big.Int
	Timestamp  *big.Int
	Day        *big.Int
}, error) {
	return _Answer.Contract.AnswerRecords(&_Answer.CallOpts, arg0)
}

// AnswerRecords is a free data retrieval call binding the contract method 0x7c3884f1.
//
// Solidity: function answerRecords(uint256 ) view returns(address user, string answer, uint256 questionId, uint256 timestamp, uint256 day)
func (_Answer *AnswerCallerSession) AnswerRecords(arg0 *big.Int) (struct {
	User       common.Address
	Answer     string
	QuestionId *big.Int
	Timestamp  *big.Int
	Day        *big.Int
}, error) {
	return _Answer.Contract.AnswerRecords(&_Answer.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Answer *AnswerCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Answer *AnswerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Answer.Contract.BalanceOf(&_Answer.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Answer *AnswerCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Answer.Contract.BalanceOf(&_Answer.CallOpts, owner)
}

// DailySubmissions is a free data retrieval call binding the contract method 0x92182edb.
//
// Solidity: function dailySubmissions(address , uint256 , uint256 ) view returns(bool)
func (_Answer *AnswerCaller) DailySubmissions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (bool, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "dailySubmissions", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DailySubmissions is a free data retrieval call binding the contract method 0x92182edb.
//
// Solidity: function dailySubmissions(address , uint256 , uint256 ) view returns(bool)
func (_Answer *AnswerSession) DailySubmissions(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (bool, error) {
	return _Answer.Contract.DailySubmissions(&_Answer.CallOpts, arg0, arg1, arg2)
}

// DailySubmissions is a free data retrieval call binding the contract method 0x92182edb.
//
// Solidity: function dailySubmissions(address , uint256 , uint256 ) view returns(bool)
func (_Answer *AnswerCallerSession) DailySubmissions(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (bool, error) {
	return _Answer.Contract.DailySubmissions(&_Answer.CallOpts, arg0, arg1, arg2)
}

// GetAnswerRecord is a free data retrieval call binding the contract method 0x978827d0.
//
// Solidity: function getAnswerRecord(uint256 recordId) view returns((address,string,uint256,uint256,uint256) record)
func (_Answer *AnswerCaller) GetAnswerRecord(opts *bind.CallOpts, recordId *big.Int) (QuestionAnswerContractAnswerRecord, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "getAnswerRecord", recordId)

	if err != nil {
		return *new(QuestionAnswerContractAnswerRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(QuestionAnswerContractAnswerRecord)).(*QuestionAnswerContractAnswerRecord)

	return out0, err

}

// GetAnswerRecord is a free data retrieval call binding the contract method 0x978827d0.
//
// Solidity: function getAnswerRecord(uint256 recordId) view returns((address,string,uint256,uint256,uint256) record)
func (_Answer *AnswerSession) GetAnswerRecord(recordId *big.Int) (QuestionAnswerContractAnswerRecord, error) {
	return _Answer.Contract.GetAnswerRecord(&_Answer.CallOpts, recordId)
}

// GetAnswerRecord is a free data retrieval call binding the contract method 0x978827d0.
//
// Solidity: function getAnswerRecord(uint256 recordId) view returns((address,string,uint256,uint256,uint256) record)
func (_Answer *AnswerCallerSession) GetAnswerRecord(recordId *big.Int) (QuestionAnswerContractAnswerRecord, error) {
	return _Answer.Contract.GetAnswerRecord(&_Answer.CallOpts, recordId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Answer *AnswerCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Answer *AnswerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Answer.Contract.GetApproved(&_Answer.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Answer *AnswerCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Answer.Contract.GetApproved(&_Answer.CallOpts, tokenId)
}

// GetCurrentDayOffset is a free data retrieval call binding the contract method 0x0c04934c.
//
// Solidity: function getCurrentDayOffset() view returns(uint256)
func (_Answer *AnswerCaller) GetCurrentDayOffset(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "getCurrentDayOffset")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentDayOffset is a free data retrieval call binding the contract method 0x0c04934c.
//
// Solidity: function getCurrentDayOffset() view returns(uint256)
func (_Answer *AnswerSession) GetCurrentDayOffset() (*big.Int, error) {
	return _Answer.Contract.GetCurrentDayOffset(&_Answer.CallOpts)
}

// GetCurrentDayOffset is a free data retrieval call binding the contract method 0x0c04934c.
//
// Solidity: function getCurrentDayOffset() view returns(uint256)
func (_Answer *AnswerCallerSession) GetCurrentDayOffset() (*big.Int, error) {
	return _Answer.Contract.GetCurrentDayOffset(&_Answer.CallOpts)
}

// GetQuestionAnswerCount is a free data retrieval call binding the contract method 0x336e24aa.
//
// Solidity: function getQuestionAnswerCount(uint256 questionId) view returns(uint256)
func (_Answer *AnswerCaller) GetQuestionAnswerCount(opts *bind.CallOpts, questionId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "getQuestionAnswerCount", questionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuestionAnswerCount is a free data retrieval call binding the contract method 0x336e24aa.
//
// Solidity: function getQuestionAnswerCount(uint256 questionId) view returns(uint256)
func (_Answer *AnswerSession) GetQuestionAnswerCount(questionId *big.Int) (*big.Int, error) {
	return _Answer.Contract.GetQuestionAnswerCount(&_Answer.CallOpts, questionId)
}

// GetQuestionAnswerCount is a free data retrieval call binding the contract method 0x336e24aa.
//
// Solidity: function getQuestionAnswerCount(uint256 questionId) view returns(uint256)
func (_Answer *AnswerCallerSession) GetQuestionAnswerCount(questionId *big.Int) (*big.Int, error) {
	return _Answer.Contract.GetQuestionAnswerCount(&_Answer.CallOpts, questionId)
}

// GetTotalAnswerCount is a free data retrieval call binding the contract method 0x0a91db23.
//
// Solidity: function getTotalAnswerCount() view returns(uint256)
func (_Answer *AnswerCaller) GetTotalAnswerCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "getTotalAnswerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalAnswerCount is a free data retrieval call binding the contract method 0x0a91db23.
//
// Solidity: function getTotalAnswerCount() view returns(uint256)
func (_Answer *AnswerSession) GetTotalAnswerCount() (*big.Int, error) {
	return _Answer.Contract.GetTotalAnswerCount(&_Answer.CallOpts)
}

// GetTotalAnswerCount is a free data retrieval call binding the contract method 0x0a91db23.
//
// Solidity: function getTotalAnswerCount() view returns(uint256)
func (_Answer *AnswerCallerSession) GetTotalAnswerCount() (*big.Int, error) {
	return _Answer.Contract.GetTotalAnswerCount(&_Answer.CallOpts)
}

// GetUserAnswerCount is a free data retrieval call binding the contract method 0x2659ac8f.
//
// Solidity: function getUserAnswerCount(address user) view returns(uint256)
func (_Answer *AnswerCaller) GetUserAnswerCount(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "getUserAnswerCount", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserAnswerCount is a free data retrieval call binding the contract method 0x2659ac8f.
//
// Solidity: function getUserAnswerCount(address user) view returns(uint256)
func (_Answer *AnswerSession) GetUserAnswerCount(user common.Address) (*big.Int, error) {
	return _Answer.Contract.GetUserAnswerCount(&_Answer.CallOpts, user)
}

// GetUserAnswerCount is a free data retrieval call binding the contract method 0x2659ac8f.
//
// Solidity: function getUserAnswerCount(address user) view returns(uint256)
func (_Answer *AnswerCallerSession) GetUserAnswerCount(user common.Address) (*big.Int, error) {
	return _Answer.Contract.GetUserAnswerCount(&_Answer.CallOpts, user)
}

// GetUserTokenId is a free data retrieval call binding the contract method 0xfefd4e4f.
//
// Solidity: function getUserTokenId(address user) view returns(uint256)
func (_Answer *AnswerCaller) GetUserTokenId(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "getUserTokenId", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserTokenId is a free data retrieval call binding the contract method 0xfefd4e4f.
//
// Solidity: function getUserTokenId(address user) view returns(uint256)
func (_Answer *AnswerSession) GetUserTokenId(user common.Address) (*big.Int, error) {
	return _Answer.Contract.GetUserTokenId(&_Answer.CallOpts, user)
}

// GetUserTokenId is a free data retrieval call binding the contract method 0xfefd4e4f.
//
// Solidity: function getUserTokenId(address user) view returns(uint256)
func (_Answer *AnswerCallerSession) GetUserTokenId(user common.Address) (*big.Int, error) {
	return _Answer.Contract.GetUserTokenId(&_Answer.CallOpts, user)
}

// HasMinted is a free data retrieval call binding the contract method 0x38e21cce.
//
// Solidity: function hasMinted(address ) view returns(bool)
func (_Answer *AnswerCaller) HasMinted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "hasMinted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasMinted is a free data retrieval call binding the contract method 0x38e21cce.
//
// Solidity: function hasMinted(address ) view returns(bool)
func (_Answer *AnswerSession) HasMinted(arg0 common.Address) (bool, error) {
	return _Answer.Contract.HasMinted(&_Answer.CallOpts, arg0)
}

// HasMinted is a free data retrieval call binding the contract method 0x38e21cce.
//
// Solidity: function hasMinted(address ) view returns(bool)
func (_Answer *AnswerCallerSession) HasMinted(arg0 common.Address) (bool, error) {
	return _Answer.Contract.HasMinted(&_Answer.CallOpts, arg0)
}

// HasSBT is a free data retrieval call binding the contract method 0x56a9c454.
//
// Solidity: function hasSBT(address user) view returns(bool)
func (_Answer *AnswerCaller) HasSBT(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "hasSBT", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasSBT is a free data retrieval call binding the contract method 0x56a9c454.
//
// Solidity: function hasSBT(address user) view returns(bool)
func (_Answer *AnswerSession) HasSBT(user common.Address) (bool, error) {
	return _Answer.Contract.HasSBT(&_Answer.CallOpts, user)
}

// HasSBT is a free data retrieval call binding the contract method 0x56a9c454.
//
// Solidity: function hasSBT(address user) view returns(bool)
func (_Answer *AnswerCallerSession) HasSBT(user common.Address) (bool, error) {
	return _Answer.Contract.HasSBT(&_Answer.CallOpts, user)
}

// HasSubmittedOnDay is a free data retrieval call binding the contract method 0x890c3140.
//
// Solidity: function hasSubmittedOnDay(address user, uint256 day, uint256 questionId) view returns(bool)
func (_Answer *AnswerCaller) HasSubmittedOnDay(opts *bind.CallOpts, user common.Address, day *big.Int, questionId *big.Int) (bool, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "hasSubmittedOnDay", user, day, questionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasSubmittedOnDay is a free data retrieval call binding the contract method 0x890c3140.
//
// Solidity: function hasSubmittedOnDay(address user, uint256 day, uint256 questionId) view returns(bool)
func (_Answer *AnswerSession) HasSubmittedOnDay(user common.Address, day *big.Int, questionId *big.Int) (bool, error) {
	return _Answer.Contract.HasSubmittedOnDay(&_Answer.CallOpts, user, day, questionId)
}

// HasSubmittedOnDay is a free data retrieval call binding the contract method 0x890c3140.
//
// Solidity: function hasSubmittedOnDay(address user, uint256 day, uint256 questionId) view returns(bool)
func (_Answer *AnswerCallerSession) HasSubmittedOnDay(user common.Address, day *big.Int, questionId *big.Int) (bool, error) {
	return _Answer.Contract.HasSubmittedOnDay(&_Answer.CallOpts, user, day, questionId)
}

// HasSubmittedToday is a free data retrieval call binding the contract method 0xe1b23689.
//
// Solidity: function hasSubmittedToday(address user, uint256 questionId) view returns(bool)
func (_Answer *AnswerCaller) HasSubmittedToday(opts *bind.CallOpts, user common.Address, questionId *big.Int) (bool, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "hasSubmittedToday", user, questionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasSubmittedToday is a free data retrieval call binding the contract method 0xe1b23689.
//
// Solidity: function hasSubmittedToday(address user, uint256 questionId) view returns(bool)
func (_Answer *AnswerSession) HasSubmittedToday(user common.Address, questionId *big.Int) (bool, error) {
	return _Answer.Contract.HasSubmittedToday(&_Answer.CallOpts, user, questionId)
}

// HasSubmittedToday is a free data retrieval call binding the contract method 0xe1b23689.
//
// Solidity: function hasSubmittedToday(address user, uint256 questionId) view returns(bool)
func (_Answer *AnswerCallerSession) HasSubmittedToday(user common.Address, questionId *big.Int) (bool, error) {
	return _Answer.Contract.HasSubmittedToday(&_Answer.CallOpts, user, questionId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Answer *AnswerCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Answer *AnswerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Answer.Contract.IsApprovedForAll(&_Answer.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Answer *AnswerCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Answer.Contract.IsApprovedForAll(&_Answer.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Answer *AnswerCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Answer *AnswerSession) Name() (string, error) {
	return _Answer.Contract.Name(&_Answer.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Answer *AnswerCallerSession) Name() (string, error) {
	return _Answer.Contract.Name(&_Answer.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Answer *AnswerCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Answer *AnswerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Answer.Contract.OwnerOf(&_Answer.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Answer *AnswerCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Answer.Contract.OwnerOf(&_Answer.CallOpts, tokenId)
}

// QuestionAnswerCount is a free data retrieval call binding the contract method 0x2e44f456.
//
// Solidity: function questionAnswerCount(uint256 ) view returns(uint256)
func (_Answer *AnswerCaller) QuestionAnswerCount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "questionAnswerCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuestionAnswerCount is a free data retrieval call binding the contract method 0x2e44f456.
//
// Solidity: function questionAnswerCount(uint256 ) view returns(uint256)
func (_Answer *AnswerSession) QuestionAnswerCount(arg0 *big.Int) (*big.Int, error) {
	return _Answer.Contract.QuestionAnswerCount(&_Answer.CallOpts, arg0)
}

// QuestionAnswerCount is a free data retrieval call binding the contract method 0x2e44f456.
//
// Solidity: function questionAnswerCount(uint256 ) view returns(uint256)
func (_Answer *AnswerCallerSession) QuestionAnswerCount(arg0 *big.Int) (*big.Int, error) {
	return _Answer.Contract.QuestionAnswerCount(&_Answer.CallOpts, arg0)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Answer *AnswerCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Answer *AnswerSession) StartTime() (*big.Int, error) {
	return _Answer.Contract.StartTime(&_Answer.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Answer *AnswerCallerSession) StartTime() (*big.Int, error) {
	return _Answer.Contract.StartTime(&_Answer.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Answer *AnswerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Answer *AnswerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Answer.Contract.SupportsInterface(&_Answer.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Answer *AnswerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Answer.Contract.SupportsInterface(&_Answer.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Answer *AnswerCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Answer *AnswerSession) Symbol() (string, error) {
	return _Answer.Contract.Symbol(&_Answer.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Answer *AnswerCallerSession) Symbol() (string, error) {
	return _Answer.Contract.Symbol(&_Answer.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 ) pure returns(string)
func (_Answer *AnswerCaller) TokenURI(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "tokenURI", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 ) pure returns(string)
func (_Answer *AnswerSession) TokenURI(arg0 *big.Int) (string, error) {
	return _Answer.Contract.TokenURI(&_Answer.CallOpts, arg0)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 ) pure returns(string)
func (_Answer *AnswerCallerSession) TokenURI(arg0 *big.Int) (string, error) {
	return _Answer.Contract.TokenURI(&_Answer.CallOpts, arg0)
}

// UserAnswerHistory is a free data retrieval call binding the contract method 0x5f466156.
//
// Solidity: function userAnswerHistory(address , uint256 ) view returns(uint256)
func (_Answer *AnswerCaller) UserAnswerHistory(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "userAnswerHistory", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserAnswerHistory is a free data retrieval call binding the contract method 0x5f466156.
//
// Solidity: function userAnswerHistory(address , uint256 ) view returns(uint256)
func (_Answer *AnswerSession) UserAnswerHistory(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Answer.Contract.UserAnswerHistory(&_Answer.CallOpts, arg0, arg1)
}

// UserAnswerHistory is a free data retrieval call binding the contract method 0x5f466156.
//
// Solidity: function userAnswerHistory(address , uint256 ) view returns(uint256)
func (_Answer *AnswerCallerSession) UserAnswerHistory(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Answer.Contract.UserAnswerHistory(&_Answer.CallOpts, arg0, arg1)
}

// UserTokenId is a free data retrieval call binding the contract method 0x715ca542.
//
// Solidity: function userTokenId(address ) view returns(uint256)
func (_Answer *AnswerCaller) UserTokenId(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "userTokenId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserTokenId is a free data retrieval call binding the contract method 0x715ca542.
//
// Solidity: function userTokenId(address ) view returns(uint256)
func (_Answer *AnswerSession) UserTokenId(arg0 common.Address) (*big.Int, error) {
	return _Answer.Contract.UserTokenId(&_Answer.CallOpts, arg0)
}

// UserTokenId is a free data retrieval call binding the contract method 0x715ca542.
//
// Solidity: function userTokenId(address ) view returns(uint256)
func (_Answer *AnswerCallerSession) UserTokenId(arg0 common.Address) (*big.Int, error) {
	return _Answer.Contract.UserTokenId(&_Answer.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Answer *AnswerTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Answer.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Answer *AnswerSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Answer.Contract.Approve(&_Answer.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Answer *AnswerTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Answer.Contract.Approve(&_Answer.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Answer *AnswerTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Answer.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Answer *AnswerSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Answer.Contract.SafeTransferFrom(&_Answer.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Answer *AnswerTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Answer.Contract.SafeTransferFrom(&_Answer.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Answer *AnswerTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Answer.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Answer *AnswerSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Answer.Contract.SafeTransferFrom0(&_Answer.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Answer *AnswerTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Answer.Contract.SafeTransferFrom0(&_Answer.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Answer *AnswerTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Answer.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Answer *AnswerSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Answer.Contract.SetApprovalForAll(&_Answer.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Answer *AnswerTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Answer.Contract.SetApprovalForAll(&_Answer.TransactOpts, operator, approved)
}

// SubmitAnswer is a paid mutator transaction binding the contract method 0xf9ae8797.
//
// Solidity: function submit_answer(string answer, uint256 id) returns()
func (_Answer *AnswerTransactor) SubmitAnswer(opts *bind.TransactOpts, answer string, id *big.Int) (*types.Transaction, error) {
	return _Answer.contract.Transact(opts, "submit_answer", answer, id)
}

// SubmitAnswer is a paid mutator transaction binding the contract method 0xf9ae8797.
//
// Solidity: function submit_answer(string answer, uint256 id) returns()
func (_Answer *AnswerSession) SubmitAnswer(answer string, id *big.Int) (*types.Transaction, error) {
	return _Answer.Contract.SubmitAnswer(&_Answer.TransactOpts, answer, id)
}

// SubmitAnswer is a paid mutator transaction binding the contract method 0xf9ae8797.
//
// Solidity: function submit_answer(string answer, uint256 id) returns()
func (_Answer *AnswerTransactorSession) SubmitAnswer(answer string, id *big.Int) (*types.Transaction, error) {
	return _Answer.Contract.SubmitAnswer(&_Answer.TransactOpts, answer, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Answer *AnswerTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Answer.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Answer *AnswerSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Answer.Contract.TransferFrom(&_Answer.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Answer *AnswerTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Answer.Contract.TransferFrom(&_Answer.TransactOpts, from, to, tokenId)
}

// AnswerAnswerSubmittedIterator is returned from FilterAnswerSubmitted and is used to iterate over the raw logs and unpacked data for AnswerSubmitted events raised by the Answer contract.
type AnswerAnswerSubmittedIterator struct {
	Event *AnswerAnswerSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AnswerAnswerSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnswerAnswerSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AnswerAnswerSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AnswerAnswerSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnswerAnswerSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnswerAnswerSubmitted represents a AnswerSubmitted event raised by the Answer contract.
type AnswerAnswerSubmitted struct {
	User       common.Address
	Answer     string
	QuestionId *big.Int
	Timestamp  *big.Int
	Day        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAnswerSubmitted is a free log retrieval operation binding the contract event 0xd482948029b8d23badc9abe4680340bb5a48488f3c093001b3a8fcc845ebe9e9.
//
// Solidity: event AnswerSubmitted(address indexed user, string answer, uint256 questionId, uint256 timestamp, uint256 day)
func (_Answer *AnswerFilterer) FilterAnswerSubmitted(opts *bind.FilterOpts, user []common.Address) (*AnswerAnswerSubmittedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Answer.contract.FilterLogs(opts, "AnswerSubmitted", userRule)
	if err != nil {
		return nil, err
	}
	return &AnswerAnswerSubmittedIterator{contract: _Answer.contract, event: "AnswerSubmitted", logs: logs, sub: sub}, nil
}

// WatchAnswerSubmitted is a free log subscription operation binding the contract event 0xd482948029b8d23badc9abe4680340bb5a48488f3c093001b3a8fcc845ebe9e9.
//
// Solidity: event AnswerSubmitted(address indexed user, string answer, uint256 questionId, uint256 timestamp, uint256 day)
func (_Answer *AnswerFilterer) WatchAnswerSubmitted(opts *bind.WatchOpts, sink chan<- *AnswerAnswerSubmitted, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Answer.contract.WatchLogs(opts, "AnswerSubmitted", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnswerAnswerSubmitted)
				if err := _Answer.contract.UnpackLog(event, "AnswerSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAnswerSubmitted is a log parse operation binding the contract event 0xd482948029b8d23badc9abe4680340bb5a48488f3c093001b3a8fcc845ebe9e9.
//
// Solidity: event AnswerSubmitted(address indexed user, string answer, uint256 questionId, uint256 timestamp, uint256 day)
func (_Answer *AnswerFilterer) ParseAnswerSubmitted(log types.Log) (*AnswerAnswerSubmitted, error) {
	event := new(AnswerAnswerSubmitted)
	if err := _Answer.contract.UnpackLog(event, "AnswerSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnswerApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Answer contract.
type AnswerApprovalIterator struct {
	Event *AnswerApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AnswerApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnswerApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AnswerApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AnswerApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnswerApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnswerApproval represents a Approval event raised by the Answer contract.
type AnswerApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Answer *AnswerFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*AnswerApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Answer.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AnswerApprovalIterator{contract: _Answer.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Answer *AnswerFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AnswerApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Answer.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnswerApproval)
				if err := _Answer.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Answer *AnswerFilterer) ParseApproval(log types.Log) (*AnswerApproval, error) {
	event := new(AnswerApproval)
	if err := _Answer.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnswerApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Answer contract.
type AnswerApprovalForAllIterator struct {
	Event *AnswerApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AnswerApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnswerApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AnswerApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AnswerApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnswerApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnswerApprovalForAll represents a ApprovalForAll event raised by the Answer contract.
type AnswerApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Answer *AnswerFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*AnswerApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Answer.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &AnswerApprovalForAllIterator{contract: _Answer.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Answer *AnswerFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *AnswerApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Answer.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnswerApprovalForAll)
				if err := _Answer.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Answer *AnswerFilterer) ParseApprovalForAll(log types.Log) (*AnswerApprovalForAll, error) {
	event := new(AnswerApprovalForAll)
	if err := _Answer.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnswerDailyLimitReachedIterator is returned from FilterDailyLimitReached and is used to iterate over the raw logs and unpacked data for DailyLimitReached events raised by the Answer contract.
type AnswerDailyLimitReachedIterator struct {
	Event *AnswerDailyLimitReached // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AnswerDailyLimitReachedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnswerDailyLimitReached)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AnswerDailyLimitReached)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AnswerDailyLimitReachedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnswerDailyLimitReachedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnswerDailyLimitReached represents a DailyLimitReached event raised by the Answer contract.
type AnswerDailyLimitReached struct {
	User       common.Address
	QuestionId *big.Int
	Day        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDailyLimitReached is a free log retrieval operation binding the contract event 0xd669607932b4b86ab223c8921367d517f03e29d19f1bd518aa432a213ab257d7.
//
// Solidity: event DailyLimitReached(address indexed user, uint256 questionId, uint256 day)
func (_Answer *AnswerFilterer) FilterDailyLimitReached(opts *bind.FilterOpts, user []common.Address) (*AnswerDailyLimitReachedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Answer.contract.FilterLogs(opts, "DailyLimitReached", userRule)
	if err != nil {
		return nil, err
	}
	return &AnswerDailyLimitReachedIterator{contract: _Answer.contract, event: "DailyLimitReached", logs: logs, sub: sub}, nil
}

// WatchDailyLimitReached is a free log subscription operation binding the contract event 0xd669607932b4b86ab223c8921367d517f03e29d19f1bd518aa432a213ab257d7.
//
// Solidity: event DailyLimitReached(address indexed user, uint256 questionId, uint256 day)
func (_Answer *AnswerFilterer) WatchDailyLimitReached(opts *bind.WatchOpts, sink chan<- *AnswerDailyLimitReached, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Answer.contract.WatchLogs(opts, "DailyLimitReached", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnswerDailyLimitReached)
				if err := _Answer.contract.UnpackLog(event, "DailyLimitReached", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDailyLimitReached is a log parse operation binding the contract event 0xd669607932b4b86ab223c8921367d517f03e29d19f1bd518aa432a213ab257d7.
//
// Solidity: event DailyLimitReached(address indexed user, uint256 questionId, uint256 day)
func (_Answer *AnswerFilterer) ParseDailyLimitReached(log types.Log) (*AnswerDailyLimitReached, error) {
	event := new(AnswerDailyLimitReached)
	if err := _Answer.contract.UnpackLog(event, "DailyLimitReached", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnswerSBTMintedIterator is returned from FilterSBTMinted and is used to iterate over the raw logs and unpacked data for SBTMinted events raised by the Answer contract.
type AnswerSBTMintedIterator struct {
	Event *AnswerSBTMinted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AnswerSBTMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnswerSBTMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AnswerSBTMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AnswerSBTMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnswerSBTMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnswerSBTMinted represents a SBTMinted event raised by the Answer contract.
type AnswerSBTMinted struct {
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSBTMinted is a free log retrieval operation binding the contract event 0x69180e5630fe9bd5b50dab3216e4151093508a9dfe7eb5fa1166632c7bc65936.
//
// Solidity: event SBTMinted(address indexed to, uint256 tokenId)
func (_Answer *AnswerFilterer) FilterSBTMinted(opts *bind.FilterOpts, to []common.Address) (*AnswerSBTMintedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Answer.contract.FilterLogs(opts, "SBTMinted", toRule)
	if err != nil {
		return nil, err
	}
	return &AnswerSBTMintedIterator{contract: _Answer.contract, event: "SBTMinted", logs: logs, sub: sub}, nil
}

// WatchSBTMinted is a free log subscription operation binding the contract event 0x69180e5630fe9bd5b50dab3216e4151093508a9dfe7eb5fa1166632c7bc65936.
//
// Solidity: event SBTMinted(address indexed to, uint256 tokenId)
func (_Answer *AnswerFilterer) WatchSBTMinted(opts *bind.WatchOpts, sink chan<- *AnswerSBTMinted, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Answer.contract.WatchLogs(opts, "SBTMinted", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnswerSBTMinted)
				if err := _Answer.contract.UnpackLog(event, "SBTMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSBTMinted is a log parse operation binding the contract event 0x69180e5630fe9bd5b50dab3216e4151093508a9dfe7eb5fa1166632c7bc65936.
//
// Solidity: event SBTMinted(address indexed to, uint256 tokenId)
func (_Answer *AnswerFilterer) ParseSBTMinted(log types.Log) (*AnswerSBTMinted, error) {
	event := new(AnswerSBTMinted)
	if err := _Answer.contract.UnpackLog(event, "SBTMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnswerTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Answer contract.
type AnswerTransferIterator struct {
	Event *AnswerTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AnswerTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnswerTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AnswerTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AnswerTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnswerTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnswerTransfer represents a Transfer event raised by the Answer contract.
type AnswerTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Answer *AnswerFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*AnswerTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Answer.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AnswerTransferIterator{contract: _Answer.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Answer *AnswerFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AnswerTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Answer.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnswerTransfer)
				if err := _Answer.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Answer *AnswerFilterer) ParseTransfer(log types.Log) (*AnswerTransfer, error) {
	event := new(AnswerTransfer)
	if err := _Answer.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
