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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"questionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"name\":\"AlreadySubmittedToday\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyAnswer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidQuestionId\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"answer\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"questionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"name\":\"AnswerSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"questionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"name\":\"DailyLimitReached\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"answerRecords\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"answer\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"questionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dailySubmissions\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"recordId\",\"type\":\"uint256\"}],\"name\":\"getAnswerRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"answer\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"questionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"internalType\":\"structQuestionAnswerContract.AnswerRecord\",\"name\":\"record\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getAnswerRecords\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"answer\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"questionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"internalType\":\"structQuestionAnswerContract.AnswerRecord[]\",\"name\":\"records\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"questionId\",\"type\":\"uint256\"}],\"name\":\"getQuestionAnswerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"questionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getQuestionAnswers\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"answer\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"questionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"internalType\":\"structQuestionAnswerContract.AnswerRecord[]\",\"name\":\"records\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalAnswerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserAnswerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getUserAnswerHistory\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"answer\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"questionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"internalType\":\"structQuestionAnswerContract.AnswerRecord[]\",\"name\":\"records\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"questionId\",\"type\":\"uint256\"}],\"name\":\"hasSubmittedOnDay\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"questionId\",\"type\":\"uint256\"}],\"name\":\"hasSubmittedToday\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"questionAnswerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"answer\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"submit_answer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userAnswerHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// GetAnswerRecords is a free data retrieval call binding the contract method 0xcf4b44ef.
//
// Solidity: function getAnswerRecords(uint256 offset, uint256 limit) view returns((address,string,uint256,uint256,uint256)[] records)
func (_Answer *AnswerCaller) GetAnswerRecords(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]QuestionAnswerContractAnswerRecord, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "getAnswerRecords", offset, limit)

	if err != nil {
		return *new([]QuestionAnswerContractAnswerRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]QuestionAnswerContractAnswerRecord)).(*[]QuestionAnswerContractAnswerRecord)

	return out0, err

}

// GetAnswerRecords is a free data retrieval call binding the contract method 0xcf4b44ef.
//
// Solidity: function getAnswerRecords(uint256 offset, uint256 limit) view returns((address,string,uint256,uint256,uint256)[] records)
func (_Answer *AnswerSession) GetAnswerRecords(offset *big.Int, limit *big.Int) ([]QuestionAnswerContractAnswerRecord, error) {
	return _Answer.Contract.GetAnswerRecords(&_Answer.CallOpts, offset, limit)
}

// GetAnswerRecords is a free data retrieval call binding the contract method 0xcf4b44ef.
//
// Solidity: function getAnswerRecords(uint256 offset, uint256 limit) view returns((address,string,uint256,uint256,uint256)[] records)
func (_Answer *AnswerCallerSession) GetAnswerRecords(offset *big.Int, limit *big.Int) ([]QuestionAnswerContractAnswerRecord, error) {
	return _Answer.Contract.GetAnswerRecords(&_Answer.CallOpts, offset, limit)
}

// GetCurrentDay is a free data retrieval call binding the contract method 0x3e6968b6.
//
// Solidity: function getCurrentDay() view returns(uint256)
func (_Answer *AnswerCaller) GetCurrentDay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "getCurrentDay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentDay is a free data retrieval call binding the contract method 0x3e6968b6.
//
// Solidity: function getCurrentDay() view returns(uint256)
func (_Answer *AnswerSession) GetCurrentDay() (*big.Int, error) {
	return _Answer.Contract.GetCurrentDay(&_Answer.CallOpts)
}

// GetCurrentDay is a free data retrieval call binding the contract method 0x3e6968b6.
//
// Solidity: function getCurrentDay() view returns(uint256)
func (_Answer *AnswerCallerSession) GetCurrentDay() (*big.Int, error) {
	return _Answer.Contract.GetCurrentDay(&_Answer.CallOpts)
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

// GetQuestionAnswers is a free data retrieval call binding the contract method 0x66c37ca9.
//
// Solidity: function getQuestionAnswers(uint256 questionId, uint256 offset, uint256 limit) view returns((address,string,uint256,uint256,uint256)[] records, uint256 total)
func (_Answer *AnswerCaller) GetQuestionAnswers(opts *bind.CallOpts, questionId *big.Int, offset *big.Int, limit *big.Int) (struct {
	Records []QuestionAnswerContractAnswerRecord
	Total   *big.Int
}, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "getQuestionAnswers", questionId, offset, limit)

	outstruct := new(struct {
		Records []QuestionAnswerContractAnswerRecord
		Total   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Records = *abi.ConvertType(out[0], new([]QuestionAnswerContractAnswerRecord)).(*[]QuestionAnswerContractAnswerRecord)
	outstruct.Total = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetQuestionAnswers is a free data retrieval call binding the contract method 0x66c37ca9.
//
// Solidity: function getQuestionAnswers(uint256 questionId, uint256 offset, uint256 limit) view returns((address,string,uint256,uint256,uint256)[] records, uint256 total)
func (_Answer *AnswerSession) GetQuestionAnswers(questionId *big.Int, offset *big.Int, limit *big.Int) (struct {
	Records []QuestionAnswerContractAnswerRecord
	Total   *big.Int
}, error) {
	return _Answer.Contract.GetQuestionAnswers(&_Answer.CallOpts, questionId, offset, limit)
}

// GetQuestionAnswers is a free data retrieval call binding the contract method 0x66c37ca9.
//
// Solidity: function getQuestionAnswers(uint256 questionId, uint256 offset, uint256 limit) view returns((address,string,uint256,uint256,uint256)[] records, uint256 total)
func (_Answer *AnswerCallerSession) GetQuestionAnswers(questionId *big.Int, offset *big.Int, limit *big.Int) (struct {
	Records []QuestionAnswerContractAnswerRecord
	Total   *big.Int
}, error) {
	return _Answer.Contract.GetQuestionAnswers(&_Answer.CallOpts, questionId, offset, limit)
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

// GetUserAnswerHistory is a free data retrieval call binding the contract method 0x8d8e8efb.
//
// Solidity: function getUserAnswerHistory(address user, uint256 offset, uint256 limit) view returns((address,string,uint256,uint256,uint256)[] records)
func (_Answer *AnswerCaller) GetUserAnswerHistory(opts *bind.CallOpts, user common.Address, offset *big.Int, limit *big.Int) ([]QuestionAnswerContractAnswerRecord, error) {
	var out []interface{}
	err := _Answer.contract.Call(opts, &out, "getUserAnswerHistory", user, offset, limit)

	if err != nil {
		return *new([]QuestionAnswerContractAnswerRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]QuestionAnswerContractAnswerRecord)).(*[]QuestionAnswerContractAnswerRecord)

	return out0, err

}

// GetUserAnswerHistory is a free data retrieval call binding the contract method 0x8d8e8efb.
//
// Solidity: function getUserAnswerHistory(address user, uint256 offset, uint256 limit) view returns((address,string,uint256,uint256,uint256)[] records)
func (_Answer *AnswerSession) GetUserAnswerHistory(user common.Address, offset *big.Int, limit *big.Int) ([]QuestionAnswerContractAnswerRecord, error) {
	return _Answer.Contract.GetUserAnswerHistory(&_Answer.CallOpts, user, offset, limit)
}

// GetUserAnswerHistory is a free data retrieval call binding the contract method 0x8d8e8efb.
//
// Solidity: function getUserAnswerHistory(address user, uint256 offset, uint256 limit) view returns((address,string,uint256,uint256,uint256)[] records)
func (_Answer *AnswerCallerSession) GetUserAnswerHistory(user common.Address, offset *big.Int, limit *big.Int) ([]QuestionAnswerContractAnswerRecord, error) {
	return _Answer.Contract.GetUserAnswerHistory(&_Answer.CallOpts, user, offset, limit)
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
