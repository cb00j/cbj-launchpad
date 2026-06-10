// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// CBJSaleMetaData contains all meta data concerning the CBJSale contract.
var CBJSaleMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_allocationStaking\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"allocationStaking\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationStaking\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositTokens\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"extendRegistrationTime\",\"inputs\":[{\"name\":\"timeToExtend\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"factory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISalesFactory\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNumberOfRegisteredUsers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getParticipation\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSaleInfo\",\"inputs\":[],\"outputs\":[{\"name\":\"totalTokensSold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalETHRaised\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountOfTokensToSell\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenPriceInETH\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVestingInfo\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isParticipated\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRegistered\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxVestingTimeShift\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numberOfParticipants\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"participate\",\"inputs\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"portionVestingPrecision\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"postponeSale\",\"inputs\":[{\"name\":\"timeToShift\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerForSale\",\"inputs\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registration\",\"inputs\":[],\"outputs\":[{\"name\":\"registrationTimeStarts\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"registrationTimeEnds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"numberOfRegistrants\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sale\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"isCreated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"earningWithdrawn\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"leftoverWithdrawn\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokensDeposited\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"saleOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenPriceInETH\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountOfTokensToSell\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalTokensSold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalETHRaised\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxParticipation\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"saleStart\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"saleEnd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensUnlockTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setCap\",\"inputs\":[{\"name\":\"cap\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRegistrationTime\",\"inputs\":[{\"name\":\"_registrationTimeStarts\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_registrationTimeEnds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSaleParams\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_saleOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenPriceInETH\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_amountOfTokensToSell\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_saleEnd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_tokensUnlockTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_maxParticipation\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_portionVestingPrecision\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSaleStart\",\"inputs\":[{\"name\":\"_saleStart\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setToken\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setVestingParams\",\"inputs\":[{\"name\":\"_unlockTimes\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_percentPerPortion\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_maxVestingTimeShift\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shiftVestingPortionsUnlockTime\",\"inputs\":[{\"name\":\"timeToShift\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTokenPriceInETH\",\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userToParticipation\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amountBought\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountETHPaid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeParticipated\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vestingPercentPerPortion\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vestingPortionsUnlockTime\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawEarnings\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawEarningsAndLeftover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawLeftover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawMultiplePortions\",\"inputs\":[{\"name\":\"portionIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawTokens\",\"inputs\":[{\"name\":\"portionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"MaxParticipationSet\",\"inputs\":[{\"name\":\"newMaxParticipation\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegistrationTimeSet\",\"inputs\":[{\"name\":\"registrationTimeStarts\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"registrationTimeEnds\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SaleCreated\",\"inputs\":[{\"name\":\"saleOwner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenPriceInETH\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountOfTokensToSell\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"saleEnd\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StartTimeSet\",\"inputs\":[{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenPriceSet\",\"inputs\":[{\"name\":\"newPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensSold\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensWithdrawn\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserRegistered\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pid\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// CBJSaleABI is the input ABI used to generate the binding from.
// Deprecated: Use CBJSaleMetaData.ABI instead.
var CBJSaleABI = CBJSaleMetaData.ABI

// CBJSale is an auto generated Go binding around an Ethereum contract.
type CBJSale struct {
	CBJSaleCaller     // Read-only binding to the contract
	CBJSaleTransactor // Write-only binding to the contract
	CBJSaleFilterer   // Log filterer for contract events
}

// CBJSaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type CBJSaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CBJSaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CBJSaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CBJSaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CBJSaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CBJSaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CBJSaleSession struct {
	Contract     *CBJSale          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CBJSaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CBJSaleCallerSession struct {
	Contract *CBJSaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CBJSaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CBJSaleTransactorSession struct {
	Contract     *CBJSaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CBJSaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type CBJSaleRaw struct {
	Contract *CBJSale // Generic contract binding to access the raw methods on
}

// CBJSaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CBJSaleCallerRaw struct {
	Contract *CBJSaleCaller // Generic read-only contract binding to access the raw methods on
}

// CBJSaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CBJSaleTransactorRaw struct {
	Contract *CBJSaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCBJSale creates a new instance of CBJSale, bound to a specific deployed contract.
func NewCBJSale(address common.Address, backend bind.ContractBackend) (*CBJSale, error) {
	contract, err := bindCBJSale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CBJSale{CBJSaleCaller: CBJSaleCaller{contract: contract}, CBJSaleTransactor: CBJSaleTransactor{contract: contract}, CBJSaleFilterer: CBJSaleFilterer{contract: contract}}, nil
}

// NewCBJSaleCaller creates a new read-only instance of CBJSale, bound to a specific deployed contract.
func NewCBJSaleCaller(address common.Address, caller bind.ContractCaller) (*CBJSaleCaller, error) {
	contract, err := bindCBJSale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CBJSaleCaller{contract: contract}, nil
}

// NewCBJSaleTransactor creates a new write-only instance of CBJSale, bound to a specific deployed contract.
func NewCBJSaleTransactor(address common.Address, transactor bind.ContractTransactor) (*CBJSaleTransactor, error) {
	contract, err := bindCBJSale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CBJSaleTransactor{contract: contract}, nil
}

// NewCBJSaleFilterer creates a new log filterer instance of CBJSale, bound to a specific deployed contract.
func NewCBJSaleFilterer(address common.Address, filterer bind.ContractFilterer) (*CBJSaleFilterer, error) {
	contract, err := bindCBJSale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CBJSaleFilterer{contract: contract}, nil
}

// bindCBJSale binds a generic wrapper to an already deployed contract.
func bindCBJSale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CBJSaleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CBJSale *CBJSaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CBJSale.Contract.CBJSaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CBJSale *CBJSaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CBJSale.Contract.CBJSaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CBJSale *CBJSaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CBJSale.Contract.CBJSaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CBJSale *CBJSaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CBJSale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CBJSale *CBJSaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CBJSale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CBJSale *CBJSaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CBJSale.Contract.contract.Transact(opts, method, params...)
}

// AllocationStaking is a free data retrieval call binding the contract method 0x135fb504.
//
// Solidity: function allocationStaking() view returns(address)
func (_CBJSale *CBJSaleCaller) AllocationStaking(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CBJSale.contract.Call(opts, &out, "allocationStaking")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationStaking is a free data retrieval call binding the contract method 0x135fb504.
//
// Solidity: function allocationStaking() view returns(address)
func (_CBJSale *CBJSaleSession) AllocationStaking() (common.Address, error) {
	return _CBJSale.Contract.AllocationStaking(&_CBJSale.CallOpts)
}

// AllocationStaking is a free data retrieval call binding the contract method 0x135fb504.
//
// Solidity: function allocationStaking() view returns(address)
func (_CBJSale *CBJSaleCallerSession) AllocationStaking() (common.Address, error) {
	return _CBJSale.Contract.AllocationStaking(&_CBJSale.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CBJSale *CBJSaleCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CBJSale.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CBJSale *CBJSaleSession) Factory() (common.Address, error) {
	return _CBJSale.Contract.Factory(&_CBJSale.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CBJSale *CBJSaleCallerSession) Factory() (common.Address, error) {
	return _CBJSale.Contract.Factory(&_CBJSale.CallOpts)
}

// GetNumberOfRegisteredUsers is a free data retrieval call binding the contract method 0xab7589b5.
//
// Solidity: function getNumberOfRegisteredUsers() view returns(uint256)
func (_CBJSale *CBJSaleCaller) GetNumberOfRegisteredUsers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CBJSale.contract.Call(opts, &out, "getNumberOfRegisteredUsers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfRegisteredUsers is a free data retrieval call binding the contract method 0xab7589b5.
//
// Solidity: function getNumberOfRegisteredUsers() view returns(uint256)
func (_CBJSale *CBJSaleSession) GetNumberOfRegisteredUsers() (*big.Int, error) {
	return _CBJSale.Contract.GetNumberOfRegisteredUsers(&_CBJSale.CallOpts)
}

// GetNumberOfRegisteredUsers is a free data retrieval call binding the contract method 0xab7589b5.
//
// Solidity: function getNumberOfRegisteredUsers() view returns(uint256)
func (_CBJSale *CBJSaleCallerSession) GetNumberOfRegisteredUsers() (*big.Int, error) {
	return _CBJSale.Contract.GetNumberOfRegisteredUsers(&_CBJSale.CallOpts)
}

// GetParticipation is a free data retrieval call binding the contract method 0xcad925ef.
//
// Solidity: function getParticipation(address _user) view returns(uint256, uint256, uint256, bool[])
func (_CBJSale *CBJSaleCaller) GetParticipation(opts *bind.CallOpts, _user common.Address) (*big.Int, *big.Int, *big.Int, []bool, error) {
	var out []interface{}
	err := _CBJSale.contract.Call(opts, &out, "getParticipation", _user)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new([]bool)).(*[]bool)

	return out0, out1, out2, out3, err

}

// GetParticipation is a free data retrieval call binding the contract method 0xcad925ef.
//
// Solidity: function getParticipation(address _user) view returns(uint256, uint256, uint256, bool[])
func (_CBJSale *CBJSaleSession) GetParticipation(_user common.Address) (*big.Int, *big.Int, *big.Int, []bool, error) {
	return _CBJSale.Contract.GetParticipation(&_CBJSale.CallOpts, _user)
}

// GetParticipation is a free data retrieval call binding the contract method 0xcad925ef.
//
// Solidity: function getParticipation(address _user) view returns(uint256, uint256, uint256, bool[])
func (_CBJSale *CBJSaleCallerSession) GetParticipation(_user common.Address) (*big.Int, *big.Int, *big.Int, []bool, error) {
	return _CBJSale.Contract.GetParticipation(&_CBJSale.CallOpts, _user)
}

// GetSaleInfo is a free data retrieval call binding the contract method 0xdb83694c.
//
// Solidity: function getSaleInfo() view returns(uint256 totalTokensSold, uint256 totalETHRaised, uint256 amountOfTokensToSell, uint256 tokenPriceInETH)
func (_CBJSale *CBJSaleCaller) GetSaleInfo(opts *bind.CallOpts) (struct {
	TotalTokensSold      *big.Int
	TotalETHRaised       *big.Int
	AmountOfTokensToSell *big.Int
	TokenPriceInETH      *big.Int
}, error) {
	var out []interface{}
	err := _CBJSale.contract.Call(opts, &out, "getSaleInfo")

	outstruct := new(struct {
		TotalTokensSold      *big.Int
		TotalETHRaised       *big.Int
		AmountOfTokensToSell *big.Int
		TokenPriceInETH      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalTokensSold = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalETHRaised = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AmountOfTokensToSell = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TokenPriceInETH = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSaleInfo is a free data retrieval call binding the contract method 0xdb83694c.
//
// Solidity: function getSaleInfo() view returns(uint256 totalTokensSold, uint256 totalETHRaised, uint256 amountOfTokensToSell, uint256 tokenPriceInETH)
func (_CBJSale *CBJSaleSession) GetSaleInfo() (struct {
	TotalTokensSold      *big.Int
	TotalETHRaised       *big.Int
	AmountOfTokensToSell *big.Int
	TokenPriceInETH      *big.Int
}, error) {
	return _CBJSale.Contract.GetSaleInfo(&_CBJSale.CallOpts)
}

// GetSaleInfo is a free data retrieval call binding the contract method 0xdb83694c.
//
// Solidity: function getSaleInfo() view returns(uint256 totalTokensSold, uint256 totalETHRaised, uint256 amountOfTokensToSell, uint256 tokenPriceInETH)
func (_CBJSale *CBJSaleCallerSession) GetSaleInfo() (struct {
	TotalTokensSold      *big.Int
	TotalETHRaised       *big.Int
	AmountOfTokensToSell *big.Int
	TokenPriceInETH      *big.Int
}, error) {
	return _CBJSale.Contract.GetSaleInfo(&_CBJSale.CallOpts)
}

// GetVestingInfo is a free data retrieval call binding the contract method 0xdc25a300.
//
// Solidity: function getVestingInfo() view returns(uint256[], uint256[])
func (_CBJSale *CBJSaleCaller) GetVestingInfo(opts *bind.CallOpts) ([]*big.Int, []*big.Int, error) {
	var out []interface{}
	err := _CBJSale.contract.Call(opts, &out, "getVestingInfo")

	if err != nil {
		return *new([]*big.Int), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetVestingInfo is a free data retrieval call binding the contract method 0xdc25a300.
//
// Solidity: function getVestingInfo() view returns(uint256[], uint256[])
func (_CBJSale *CBJSaleSession) GetVestingInfo() ([]*big.Int, []*big.Int, error) {
	return _CBJSale.Contract.GetVestingInfo(&_CBJSale.CallOpts)
}

// GetVestingInfo is a free data retrieval call binding the contract method 0xdc25a300.
//
// Solidity: function getVestingInfo() view returns(uint256[], uint256[])
func (_CBJSale *CBJSaleCallerSession) GetVestingInfo() ([]*big.Int, []*big.Int, error) {
	return _CBJSale.Contract.GetVestingInfo(&_CBJSale.CallOpts)
}

// IsParticipated is a free data retrieval call binding the contract method 0xcf5b8d4b.
//
// Solidity: function isParticipated(address ) view returns(bool)
func (_CBJSale *CBJSaleCaller) IsParticipated(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CBJSale.contract.Call(opts, &out, "isParticipated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsParticipated is a free data retrieval call binding the contract method 0xcf5b8d4b.
//
// Solidity: function isParticipated(address ) view returns(bool)
func (_CBJSale *CBJSaleSession) IsParticipated(arg0 common.Address) (bool, error) {
	return _CBJSale.Contract.IsParticipated(&_CBJSale.CallOpts, arg0)
}

// IsParticipated is a free data retrieval call binding the contract method 0xcf5b8d4b.
//
// Solidity: function isParticipated(address ) view returns(bool)
func (_CBJSale *CBJSaleCallerSession) IsParticipated(arg0 common.Address) (bool, error) {
	return _CBJSale.Contract.IsParticipated(&_CBJSale.CallOpts, arg0)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_CBJSale *CBJSaleCaller) IsRegistered(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CBJSale.contract.Call(opts, &out, "isRegistered", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_CBJSale *CBJSaleSession) IsRegistered(arg0 common.Address) (bool, error) {
	return _CBJSale.Contract.IsRegistered(&_CBJSale.CallOpts, arg0)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_CBJSale *CBJSaleCallerSession) IsRegistered(arg0 common.Address) (bool, error) {
	return _CBJSale.Contract.IsRegistered(&_CBJSale.CallOpts, arg0)
}

// MaxVestingTimeShift is a free data retrieval call binding the contract method 0xccc171f5.
//
// Solidity: function maxVestingTimeShift() view returns(uint256)
func (_CBJSale *CBJSaleCaller) MaxVestingTimeShift(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CBJSale.contract.Call(opts, &out, "maxVestingTimeShift")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxVestingTimeShift is a free data retrieval call binding the contract method 0xccc171f5.
//
// Solidity: function maxVestingTimeShift() view returns(uint256)
func (_CBJSale *CBJSaleSession) MaxVestingTimeShift() (*big.Int, error) {
	return _CBJSale.Contract.MaxVestingTimeShift(&_CBJSale.CallOpts)
}

// MaxVestingTimeShift is a free data retrieval call binding the contract method 0xccc171f5.
//
// Solidity: function maxVestingTimeShift() view returns(uint256)
func (_CBJSale *CBJSaleCallerSession) MaxVestingTimeShift() (*big.Int, error) {
	return _CBJSale.Contract.MaxVestingTimeShift(&_CBJSale.CallOpts)
}

// NumberOfParticipants is a free data retrieval call binding the contract method 0x7417040e.
//
// Solidity: function numberOfParticipants() view returns(uint256)
func (_CBJSale *CBJSaleCaller) NumberOfParticipants(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CBJSale.contract.Call(opts, &out, "numberOfParticipants")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfParticipants is a free data retrieval call binding the contract method 0x7417040e.
//
// Solidity: function numberOfParticipants() view returns(uint256)
func (_CBJSale *CBJSaleSession) NumberOfParticipants() (*big.Int, error) {
	return _CBJSale.Contract.NumberOfParticipants(&_CBJSale.CallOpts)
}

// NumberOfParticipants is a free data retrieval call binding the contract method 0x7417040e.
//
// Solidity: function numberOfParticipants() view returns(uint256)
func (_CBJSale *CBJSaleCallerSession) NumberOfParticipants() (*big.Int, error) {
	return _CBJSale.Contract.NumberOfParticipants(&_CBJSale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CBJSale *CBJSaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CBJSale.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CBJSale *CBJSaleSession) Owner() (common.Address, error) {
	return _CBJSale.Contract.Owner(&_CBJSale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CBJSale *CBJSaleCallerSession) Owner() (common.Address, error) {
	return _CBJSale.Contract.Owner(&_CBJSale.CallOpts)
}

// PortionVestingPrecision is a free data retrieval call binding the contract method 0x2a7c35de.
//
// Solidity: function portionVestingPrecision() view returns(uint256)
func (_CBJSale *CBJSaleCaller) PortionVestingPrecision(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CBJSale.contract.Call(opts, &out, "portionVestingPrecision")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PortionVestingPrecision is a free data retrieval call binding the contract method 0x2a7c35de.
//
// Solidity: function portionVestingPrecision() view returns(uint256)
func (_CBJSale *CBJSaleSession) PortionVestingPrecision() (*big.Int, error) {
	return _CBJSale.Contract.PortionVestingPrecision(&_CBJSale.CallOpts)
}

// PortionVestingPrecision is a free data retrieval call binding the contract method 0x2a7c35de.
//
// Solidity: function portionVestingPrecision() view returns(uint256)
func (_CBJSale *CBJSaleCallerSession) PortionVestingPrecision() (*big.Int, error) {
	return _CBJSale.Contract.PortionVestingPrecision(&_CBJSale.CallOpts)
}

// Registration is a free data retrieval call binding the contract method 0x443bd1d0.
//
// Solidity: function registration() view returns(uint256 registrationTimeStarts, uint256 registrationTimeEnds, uint256 numberOfRegistrants)
func (_CBJSale *CBJSaleCaller) Registration(opts *bind.CallOpts) (struct {
	RegistrationTimeStarts *big.Int
	RegistrationTimeEnds   *big.Int
	NumberOfRegistrants    *big.Int
}, error) {
	var out []interface{}
	err := _CBJSale.contract.Call(opts, &out, "registration")

	outstruct := new(struct {
		RegistrationTimeStarts *big.Int
		RegistrationTimeEnds   *big.Int
		NumberOfRegistrants    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RegistrationTimeStarts = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RegistrationTimeEnds = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NumberOfRegistrants = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Registration is a free data retrieval call binding the contract method 0x443bd1d0.
//
// Solidity: function registration() view returns(uint256 registrationTimeStarts, uint256 registrationTimeEnds, uint256 numberOfRegistrants)
func (_CBJSale *CBJSaleSession) Registration() (struct {
	RegistrationTimeStarts *big.Int
	RegistrationTimeEnds   *big.Int
	NumberOfRegistrants    *big.Int
}, error) {
	return _CBJSale.Contract.Registration(&_CBJSale.CallOpts)
}

// Registration is a free data retrieval call binding the contract method 0x443bd1d0.
//
// Solidity: function registration() view returns(uint256 registrationTimeStarts, uint256 registrationTimeEnds, uint256 numberOfRegistrants)
func (_CBJSale *CBJSaleCallerSession) Registration() (struct {
	RegistrationTimeStarts *big.Int
	RegistrationTimeEnds   *big.Int
	NumberOfRegistrants    *big.Int
}, error) {
	return _CBJSale.Contract.Registration(&_CBJSale.CallOpts)
}

// Sale is a free data retrieval call binding the contract method 0x6ad1fe02.
//
// Solidity: function sale() view returns(address token, bool isCreated, bool earningWithdrawn, bool leftoverWithdrawn, bool tokensDeposited, address saleOwner, uint256 tokenPriceInETH, uint256 amountOfTokensToSell, uint256 totalTokensSold, uint256 totalETHRaised, uint256 maxParticipation, uint256 saleStart, uint256 saleEnd, uint256 tokensUnlockTime)
func (_CBJSale *CBJSaleCaller) Sale(opts *bind.CallOpts) (struct {
	Token                common.Address
	IsCreated            bool
	EarningWithdrawn     bool
	LeftoverWithdrawn    bool
	TokensDeposited      bool
	SaleOwner            common.Address
	TokenPriceInETH      *big.Int
	AmountOfTokensToSell *big.Int
	TotalTokensSold      *big.Int
	TotalETHRaised       *big.Int
	MaxParticipation     *big.Int
	SaleStart            *big.Int
	SaleEnd              *big.Int
	TokensUnlockTime     *big.Int
}, error) {
	var out []interface{}
	err := _CBJSale.contract.Call(opts, &out, "sale")

	outstruct := new(struct {
		Token                common.Address
		IsCreated            bool
		EarningWithdrawn     bool
		LeftoverWithdrawn    bool
		TokensDeposited      bool
		SaleOwner            common.Address
		TokenPriceInETH      *big.Int
		AmountOfTokensToSell *big.Int
		TotalTokensSold      *big.Int
		TotalETHRaised       *big.Int
		MaxParticipation     *big.Int
		SaleStart            *big.Int
		SaleEnd              *big.Int
		TokensUnlockTime     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.IsCreated = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.EarningWithdrawn = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.LeftoverWithdrawn = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.TokensDeposited = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.SaleOwner = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.TokenPriceInETH = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.AmountOfTokensToSell = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.TotalTokensSold = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.TotalETHRaised = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.MaxParticipation = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.SaleStart = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.SaleEnd = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)
	outstruct.TokensUnlockTime = *abi.ConvertType(out[13], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Sale is a free data retrieval call binding the contract method 0x6ad1fe02.
//
// Solidity: function sale() view returns(address token, bool isCreated, bool earningWithdrawn, bool leftoverWithdrawn, bool tokensDeposited, address saleOwner, uint256 tokenPriceInETH, uint256 amountOfTokensToSell, uint256 totalTokensSold, uint256 totalETHRaised, uint256 maxParticipation, uint256 saleStart, uint256 saleEnd, uint256 tokensUnlockTime)
func (_CBJSale *CBJSaleSession) Sale() (struct {
	Token                common.Address
	IsCreated            bool
	EarningWithdrawn     bool
	LeftoverWithdrawn    bool
	TokensDeposited      bool
	SaleOwner            common.Address
	TokenPriceInETH      *big.Int
	AmountOfTokensToSell *big.Int
	TotalTokensSold      *big.Int
	TotalETHRaised       *big.Int
	MaxParticipation     *big.Int
	SaleStart            *big.Int
	SaleEnd              *big.Int
	TokensUnlockTime     *big.Int
}, error) {
	return _CBJSale.Contract.Sale(&_CBJSale.CallOpts)
}

// Sale is a free data retrieval call binding the contract method 0x6ad1fe02.
//
// Solidity: function sale() view returns(address token, bool isCreated, bool earningWithdrawn, bool leftoverWithdrawn, bool tokensDeposited, address saleOwner, uint256 tokenPriceInETH, uint256 amountOfTokensToSell, uint256 totalTokensSold, uint256 totalETHRaised, uint256 maxParticipation, uint256 saleStart, uint256 saleEnd, uint256 tokensUnlockTime)
func (_CBJSale *CBJSaleCallerSession) Sale() (struct {
	Token                common.Address
	IsCreated            bool
	EarningWithdrawn     bool
	LeftoverWithdrawn    bool
	TokensDeposited      bool
	SaleOwner            common.Address
	TokenPriceInETH      *big.Int
	AmountOfTokensToSell *big.Int
	TotalTokensSold      *big.Int
	TotalETHRaised       *big.Int
	MaxParticipation     *big.Int
	SaleStart            *big.Int
	SaleEnd              *big.Int
	TokensUnlockTime     *big.Int
}, error) {
	return _CBJSale.Contract.Sale(&_CBJSale.CallOpts)
}

// UserToParticipation is a free data retrieval call binding the contract method 0x5e7464f6.
//
// Solidity: function userToParticipation(address ) view returns(uint256 amountBought, uint256 amountETHPaid, uint256 timeParticipated)
func (_CBJSale *CBJSaleCaller) UserToParticipation(opts *bind.CallOpts, arg0 common.Address) (struct {
	AmountBought     *big.Int
	AmountETHPaid    *big.Int
	TimeParticipated *big.Int
}, error) {
	var out []interface{}
	err := _CBJSale.contract.Call(opts, &out, "userToParticipation", arg0)

	outstruct := new(struct {
		AmountBought     *big.Int
		AmountETHPaid    *big.Int
		TimeParticipated *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountBought = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountETHPaid = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TimeParticipated = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserToParticipation is a free data retrieval call binding the contract method 0x5e7464f6.
//
// Solidity: function userToParticipation(address ) view returns(uint256 amountBought, uint256 amountETHPaid, uint256 timeParticipated)
func (_CBJSale *CBJSaleSession) UserToParticipation(arg0 common.Address) (struct {
	AmountBought     *big.Int
	AmountETHPaid    *big.Int
	TimeParticipated *big.Int
}, error) {
	return _CBJSale.Contract.UserToParticipation(&_CBJSale.CallOpts, arg0)
}

// UserToParticipation is a free data retrieval call binding the contract method 0x5e7464f6.
//
// Solidity: function userToParticipation(address ) view returns(uint256 amountBought, uint256 amountETHPaid, uint256 timeParticipated)
func (_CBJSale *CBJSaleCallerSession) UserToParticipation(arg0 common.Address) (struct {
	AmountBought     *big.Int
	AmountETHPaid    *big.Int
	TimeParticipated *big.Int
}, error) {
	return _CBJSale.Contract.UserToParticipation(&_CBJSale.CallOpts, arg0)
}

// VestingPercentPerPortion is a free data retrieval call binding the contract method 0x927f6aee.
//
// Solidity: function vestingPercentPerPortion(uint256 ) view returns(uint256)
func (_CBJSale *CBJSaleCaller) VestingPercentPerPortion(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CBJSale.contract.Call(opts, &out, "vestingPercentPerPortion", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VestingPercentPerPortion is a free data retrieval call binding the contract method 0x927f6aee.
//
// Solidity: function vestingPercentPerPortion(uint256 ) view returns(uint256)
func (_CBJSale *CBJSaleSession) VestingPercentPerPortion(arg0 *big.Int) (*big.Int, error) {
	return _CBJSale.Contract.VestingPercentPerPortion(&_CBJSale.CallOpts, arg0)
}

// VestingPercentPerPortion is a free data retrieval call binding the contract method 0x927f6aee.
//
// Solidity: function vestingPercentPerPortion(uint256 ) view returns(uint256)
func (_CBJSale *CBJSaleCallerSession) VestingPercentPerPortion(arg0 *big.Int) (*big.Int, error) {
	return _CBJSale.Contract.VestingPercentPerPortion(&_CBJSale.CallOpts, arg0)
}

// VestingPortionsUnlockTime is a free data retrieval call binding the contract method 0xf1ef7ff2.
//
// Solidity: function vestingPortionsUnlockTime(uint256 ) view returns(uint256)
func (_CBJSale *CBJSaleCaller) VestingPortionsUnlockTime(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CBJSale.contract.Call(opts, &out, "vestingPortionsUnlockTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VestingPortionsUnlockTime is a free data retrieval call binding the contract method 0xf1ef7ff2.
//
// Solidity: function vestingPortionsUnlockTime(uint256 ) view returns(uint256)
func (_CBJSale *CBJSaleSession) VestingPortionsUnlockTime(arg0 *big.Int) (*big.Int, error) {
	return _CBJSale.Contract.VestingPortionsUnlockTime(&_CBJSale.CallOpts, arg0)
}

// VestingPortionsUnlockTime is a free data retrieval call binding the contract method 0xf1ef7ff2.
//
// Solidity: function vestingPortionsUnlockTime(uint256 ) view returns(uint256)
func (_CBJSale *CBJSaleCallerSession) VestingPortionsUnlockTime(arg0 *big.Int) (*big.Int, error) {
	return _CBJSale.Contract.VestingPortionsUnlockTime(&_CBJSale.CallOpts, arg0)
}

// DepositTokens is a paid mutator transaction binding the contract method 0xdd49756e.
//
// Solidity: function depositTokens(uint256 amount) returns()
func (_CBJSale *CBJSaleTransactor) DepositTokens(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CBJSale.contract.Transact(opts, "depositTokens", amount)
}

// DepositTokens is a paid mutator transaction binding the contract method 0xdd49756e.
//
// Solidity: function depositTokens(uint256 amount) returns()
func (_CBJSale *CBJSaleSession) DepositTokens(amount *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.DepositTokens(&_CBJSale.TransactOpts, amount)
}

// DepositTokens is a paid mutator transaction binding the contract method 0xdd49756e.
//
// Solidity: function depositTokens(uint256 amount) returns()
func (_CBJSale *CBJSaleTransactorSession) DepositTokens(amount *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.DepositTokens(&_CBJSale.TransactOpts, amount)
}

// ExtendRegistrationTime is a paid mutator transaction binding the contract method 0x1d321241.
//
// Solidity: function extendRegistrationTime(uint256 timeToExtend) returns()
func (_CBJSale *CBJSaleTransactor) ExtendRegistrationTime(opts *bind.TransactOpts, timeToExtend *big.Int) (*types.Transaction, error) {
	return _CBJSale.contract.Transact(opts, "extendRegistrationTime", timeToExtend)
}

// ExtendRegistrationTime is a paid mutator transaction binding the contract method 0x1d321241.
//
// Solidity: function extendRegistrationTime(uint256 timeToExtend) returns()
func (_CBJSale *CBJSaleSession) ExtendRegistrationTime(timeToExtend *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.ExtendRegistrationTime(&_CBJSale.TransactOpts, timeToExtend)
}

// ExtendRegistrationTime is a paid mutator transaction binding the contract method 0x1d321241.
//
// Solidity: function extendRegistrationTime(uint256 timeToExtend) returns()
func (_CBJSale *CBJSaleTransactorSession) ExtendRegistrationTime(timeToExtend *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.ExtendRegistrationTime(&_CBJSale.TransactOpts, timeToExtend)
}

// Participate is a paid mutator transaction binding the contract method 0x931d81c9.
//
// Solidity: function participate(bytes signature, uint256 amount) payable returns()
func (_CBJSale *CBJSaleTransactor) Participate(opts *bind.TransactOpts, signature []byte, amount *big.Int) (*types.Transaction, error) {
	return _CBJSale.contract.Transact(opts, "participate", signature, amount)
}

// Participate is a paid mutator transaction binding the contract method 0x931d81c9.
//
// Solidity: function participate(bytes signature, uint256 amount) payable returns()
func (_CBJSale *CBJSaleSession) Participate(signature []byte, amount *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.Participate(&_CBJSale.TransactOpts, signature, amount)
}

// Participate is a paid mutator transaction binding the contract method 0x931d81c9.
//
// Solidity: function participate(bytes signature, uint256 amount) payable returns()
func (_CBJSale *CBJSaleTransactorSession) Participate(signature []byte, amount *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.Participate(&_CBJSale.TransactOpts, signature, amount)
}

// PostponeSale is a paid mutator transaction binding the contract method 0x1f11cb1e.
//
// Solidity: function postponeSale(uint256 timeToShift) returns()
func (_CBJSale *CBJSaleTransactor) PostponeSale(opts *bind.TransactOpts, timeToShift *big.Int) (*types.Transaction, error) {
	return _CBJSale.contract.Transact(opts, "postponeSale", timeToShift)
}

// PostponeSale is a paid mutator transaction binding the contract method 0x1f11cb1e.
//
// Solidity: function postponeSale(uint256 timeToShift) returns()
func (_CBJSale *CBJSaleSession) PostponeSale(timeToShift *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.PostponeSale(&_CBJSale.TransactOpts, timeToShift)
}

// PostponeSale is a paid mutator transaction binding the contract method 0x1f11cb1e.
//
// Solidity: function postponeSale(uint256 timeToShift) returns()
func (_CBJSale *CBJSaleTransactorSession) PostponeSale(timeToShift *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.PostponeSale(&_CBJSale.TransactOpts, timeToShift)
}

// RegisterForSale is a paid mutator transaction binding the contract method 0xe9d8479e.
//
// Solidity: function registerForSale(bytes signature, uint256 pid) returns()
func (_CBJSale *CBJSaleTransactor) RegisterForSale(opts *bind.TransactOpts, signature []byte, pid *big.Int) (*types.Transaction, error) {
	return _CBJSale.contract.Transact(opts, "registerForSale", signature, pid)
}

// RegisterForSale is a paid mutator transaction binding the contract method 0xe9d8479e.
//
// Solidity: function registerForSale(bytes signature, uint256 pid) returns()
func (_CBJSale *CBJSaleSession) RegisterForSale(signature []byte, pid *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.RegisterForSale(&_CBJSale.TransactOpts, signature, pid)
}

// RegisterForSale is a paid mutator transaction binding the contract method 0xe9d8479e.
//
// Solidity: function registerForSale(bytes signature, uint256 pid) returns()
func (_CBJSale *CBJSaleTransactorSession) RegisterForSale(signature []byte, pid *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.RegisterForSale(&_CBJSale.TransactOpts, signature, pid)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CBJSale *CBJSaleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CBJSale.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CBJSale *CBJSaleSession) RenounceOwnership() (*types.Transaction, error) {
	return _CBJSale.Contract.RenounceOwnership(&_CBJSale.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CBJSale *CBJSaleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CBJSale.Contract.RenounceOwnership(&_CBJSale.TransactOpts)
}

// SetCap is a paid mutator transaction binding the contract method 0x47786d37.
//
// Solidity: function setCap(uint256 cap) returns()
func (_CBJSale *CBJSaleTransactor) SetCap(opts *bind.TransactOpts, cap *big.Int) (*types.Transaction, error) {
	return _CBJSale.contract.Transact(opts, "setCap", cap)
}

// SetCap is a paid mutator transaction binding the contract method 0x47786d37.
//
// Solidity: function setCap(uint256 cap) returns()
func (_CBJSale *CBJSaleSession) SetCap(cap *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.SetCap(&_CBJSale.TransactOpts, cap)
}

// SetCap is a paid mutator transaction binding the contract method 0x47786d37.
//
// Solidity: function setCap(uint256 cap) returns()
func (_CBJSale *CBJSaleTransactorSession) SetCap(cap *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.SetCap(&_CBJSale.TransactOpts, cap)
}

// SetRegistrationTime is a paid mutator transaction binding the contract method 0xe099cf64.
//
// Solidity: function setRegistrationTime(uint256 _registrationTimeStarts, uint256 _registrationTimeEnds) returns()
func (_CBJSale *CBJSaleTransactor) SetRegistrationTime(opts *bind.TransactOpts, _registrationTimeStarts *big.Int, _registrationTimeEnds *big.Int) (*types.Transaction, error) {
	return _CBJSale.contract.Transact(opts, "setRegistrationTime", _registrationTimeStarts, _registrationTimeEnds)
}

// SetRegistrationTime is a paid mutator transaction binding the contract method 0xe099cf64.
//
// Solidity: function setRegistrationTime(uint256 _registrationTimeStarts, uint256 _registrationTimeEnds) returns()
func (_CBJSale *CBJSaleSession) SetRegistrationTime(_registrationTimeStarts *big.Int, _registrationTimeEnds *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.SetRegistrationTime(&_CBJSale.TransactOpts, _registrationTimeStarts, _registrationTimeEnds)
}

// SetRegistrationTime is a paid mutator transaction binding the contract method 0xe099cf64.
//
// Solidity: function setRegistrationTime(uint256 _registrationTimeStarts, uint256 _registrationTimeEnds) returns()
func (_CBJSale *CBJSaleTransactorSession) SetRegistrationTime(_registrationTimeStarts *big.Int, _registrationTimeEnds *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.SetRegistrationTime(&_CBJSale.TransactOpts, _registrationTimeStarts, _registrationTimeEnds)
}

// SetSaleParams is a paid mutator transaction binding the contract method 0xc4fbe091.
//
// Solidity: function setSaleParams(address _token, address _saleOwner, uint256 _tokenPriceInETH, uint256 _amountOfTokensToSell, uint256 _saleEnd, uint256 _tokensUnlockTime, uint256 _maxParticipation, uint256 _portionVestingPrecision) returns()
func (_CBJSale *CBJSaleTransactor) SetSaleParams(opts *bind.TransactOpts, _token common.Address, _saleOwner common.Address, _tokenPriceInETH *big.Int, _amountOfTokensToSell *big.Int, _saleEnd *big.Int, _tokensUnlockTime *big.Int, _maxParticipation *big.Int, _portionVestingPrecision *big.Int) (*types.Transaction, error) {
	return _CBJSale.contract.Transact(opts, "setSaleParams", _token, _saleOwner, _tokenPriceInETH, _amountOfTokensToSell, _saleEnd, _tokensUnlockTime, _maxParticipation, _portionVestingPrecision)
}

// SetSaleParams is a paid mutator transaction binding the contract method 0xc4fbe091.
//
// Solidity: function setSaleParams(address _token, address _saleOwner, uint256 _tokenPriceInETH, uint256 _amountOfTokensToSell, uint256 _saleEnd, uint256 _tokensUnlockTime, uint256 _maxParticipation, uint256 _portionVestingPrecision) returns()
func (_CBJSale *CBJSaleSession) SetSaleParams(_token common.Address, _saleOwner common.Address, _tokenPriceInETH *big.Int, _amountOfTokensToSell *big.Int, _saleEnd *big.Int, _tokensUnlockTime *big.Int, _maxParticipation *big.Int, _portionVestingPrecision *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.SetSaleParams(&_CBJSale.TransactOpts, _token, _saleOwner, _tokenPriceInETH, _amountOfTokensToSell, _saleEnd, _tokensUnlockTime, _maxParticipation, _portionVestingPrecision)
}

// SetSaleParams is a paid mutator transaction binding the contract method 0xc4fbe091.
//
// Solidity: function setSaleParams(address _token, address _saleOwner, uint256 _tokenPriceInETH, uint256 _amountOfTokensToSell, uint256 _saleEnd, uint256 _tokensUnlockTime, uint256 _maxParticipation, uint256 _portionVestingPrecision) returns()
func (_CBJSale *CBJSaleTransactorSession) SetSaleParams(_token common.Address, _saleOwner common.Address, _tokenPriceInETH *big.Int, _amountOfTokensToSell *big.Int, _saleEnd *big.Int, _tokensUnlockTime *big.Int, _maxParticipation *big.Int, _portionVestingPrecision *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.SetSaleParams(&_CBJSale.TransactOpts, _token, _saleOwner, _tokenPriceInETH, _amountOfTokensToSell, _saleEnd, _tokensUnlockTime, _maxParticipation, _portionVestingPrecision)
}

// SetSaleStart is a paid mutator transaction binding the contract method 0x2f181f54.
//
// Solidity: function setSaleStart(uint256 _saleStart) returns()
func (_CBJSale *CBJSaleTransactor) SetSaleStart(opts *bind.TransactOpts, _saleStart *big.Int) (*types.Transaction, error) {
	return _CBJSale.contract.Transact(opts, "setSaleStart", _saleStart)
}

// SetSaleStart is a paid mutator transaction binding the contract method 0x2f181f54.
//
// Solidity: function setSaleStart(uint256 _saleStart) returns()
func (_CBJSale *CBJSaleSession) SetSaleStart(_saleStart *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.SetSaleStart(&_CBJSale.TransactOpts, _saleStart)
}

// SetSaleStart is a paid mutator transaction binding the contract method 0x2f181f54.
//
// Solidity: function setSaleStart(uint256 _saleStart) returns()
func (_CBJSale *CBJSaleTransactorSession) SetSaleStart(_saleStart *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.SetSaleStart(&_CBJSale.TransactOpts, _saleStart)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _token) returns()
func (_CBJSale *CBJSaleTransactor) SetToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _CBJSale.contract.Transact(opts, "setToken", _token)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _token) returns()
func (_CBJSale *CBJSaleSession) SetToken(_token common.Address) (*types.Transaction, error) {
	return _CBJSale.Contract.SetToken(&_CBJSale.TransactOpts, _token)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _token) returns()
func (_CBJSale *CBJSaleTransactorSession) SetToken(_token common.Address) (*types.Transaction, error) {
	return _CBJSale.Contract.SetToken(&_CBJSale.TransactOpts, _token)
}

// SetVestingParams is a paid mutator transaction binding the contract method 0xd937d456.
//
// Solidity: function setVestingParams(uint256[] _unlockTimes, uint256[] _percentPerPortion, uint256 _maxVestingTimeShift) returns()
func (_CBJSale *CBJSaleTransactor) SetVestingParams(opts *bind.TransactOpts, _unlockTimes []*big.Int, _percentPerPortion []*big.Int, _maxVestingTimeShift *big.Int) (*types.Transaction, error) {
	return _CBJSale.contract.Transact(opts, "setVestingParams", _unlockTimes, _percentPerPortion, _maxVestingTimeShift)
}

// SetVestingParams is a paid mutator transaction binding the contract method 0xd937d456.
//
// Solidity: function setVestingParams(uint256[] _unlockTimes, uint256[] _percentPerPortion, uint256 _maxVestingTimeShift) returns()
func (_CBJSale *CBJSaleSession) SetVestingParams(_unlockTimes []*big.Int, _percentPerPortion []*big.Int, _maxVestingTimeShift *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.SetVestingParams(&_CBJSale.TransactOpts, _unlockTimes, _percentPerPortion, _maxVestingTimeShift)
}

// SetVestingParams is a paid mutator transaction binding the contract method 0xd937d456.
//
// Solidity: function setVestingParams(uint256[] _unlockTimes, uint256[] _percentPerPortion, uint256 _maxVestingTimeShift) returns()
func (_CBJSale *CBJSaleTransactorSession) SetVestingParams(_unlockTimes []*big.Int, _percentPerPortion []*big.Int, _maxVestingTimeShift *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.SetVestingParams(&_CBJSale.TransactOpts, _unlockTimes, _percentPerPortion, _maxVestingTimeShift)
}

// ShiftVestingPortionsUnlockTime is a paid mutator transaction binding the contract method 0xef693532.
//
// Solidity: function shiftVestingPortionsUnlockTime(uint256 timeToShift) returns()
func (_CBJSale *CBJSaleTransactor) ShiftVestingPortionsUnlockTime(opts *bind.TransactOpts, timeToShift *big.Int) (*types.Transaction, error) {
	return _CBJSale.contract.Transact(opts, "shiftVestingPortionsUnlockTime", timeToShift)
}

// ShiftVestingPortionsUnlockTime is a paid mutator transaction binding the contract method 0xef693532.
//
// Solidity: function shiftVestingPortionsUnlockTime(uint256 timeToShift) returns()
func (_CBJSale *CBJSaleSession) ShiftVestingPortionsUnlockTime(timeToShift *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.ShiftVestingPortionsUnlockTime(&_CBJSale.TransactOpts, timeToShift)
}

// ShiftVestingPortionsUnlockTime is a paid mutator transaction binding the contract method 0xef693532.
//
// Solidity: function shiftVestingPortionsUnlockTime(uint256 timeToShift) returns()
func (_CBJSale *CBJSaleTransactorSession) ShiftVestingPortionsUnlockTime(timeToShift *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.ShiftVestingPortionsUnlockTime(&_CBJSale.TransactOpts, timeToShift)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CBJSale *CBJSaleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CBJSale.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CBJSale *CBJSaleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CBJSale.Contract.TransferOwnership(&_CBJSale.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CBJSale *CBJSaleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CBJSale.Contract.TransferOwnership(&_CBJSale.TransactOpts, newOwner)
}

// UpdateTokenPriceInETH is a paid mutator transaction binding the contract method 0x7843990a.
//
// Solidity: function updateTokenPriceInETH(uint256 price) returns()
func (_CBJSale *CBJSaleTransactor) UpdateTokenPriceInETH(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _CBJSale.contract.Transact(opts, "updateTokenPriceInETH", price)
}

// UpdateTokenPriceInETH is a paid mutator transaction binding the contract method 0x7843990a.
//
// Solidity: function updateTokenPriceInETH(uint256 price) returns()
func (_CBJSale *CBJSaleSession) UpdateTokenPriceInETH(price *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.UpdateTokenPriceInETH(&_CBJSale.TransactOpts, price)
}

// UpdateTokenPriceInETH is a paid mutator transaction binding the contract method 0x7843990a.
//
// Solidity: function updateTokenPriceInETH(uint256 price) returns()
func (_CBJSale *CBJSaleTransactorSession) UpdateTokenPriceInETH(price *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.UpdateTokenPriceInETH(&_CBJSale.TransactOpts, price)
}

// WithdrawEarnings is a paid mutator transaction binding the contract method 0xb73c6ce9.
//
// Solidity: function withdrawEarnings() returns()
func (_CBJSale *CBJSaleTransactor) WithdrawEarnings(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CBJSale.contract.Transact(opts, "withdrawEarnings")
}

// WithdrawEarnings is a paid mutator transaction binding the contract method 0xb73c6ce9.
//
// Solidity: function withdrawEarnings() returns()
func (_CBJSale *CBJSaleSession) WithdrawEarnings() (*types.Transaction, error) {
	return _CBJSale.Contract.WithdrawEarnings(&_CBJSale.TransactOpts)
}

// WithdrawEarnings is a paid mutator transaction binding the contract method 0xb73c6ce9.
//
// Solidity: function withdrawEarnings() returns()
func (_CBJSale *CBJSaleTransactorSession) WithdrawEarnings() (*types.Transaction, error) {
	return _CBJSale.Contract.WithdrawEarnings(&_CBJSale.TransactOpts)
}

// WithdrawEarningsAndLeftover is a paid mutator transaction binding the contract method 0xda4d4fbf.
//
// Solidity: function withdrawEarningsAndLeftover() returns()
func (_CBJSale *CBJSaleTransactor) WithdrawEarningsAndLeftover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CBJSale.contract.Transact(opts, "withdrawEarningsAndLeftover")
}

// WithdrawEarningsAndLeftover is a paid mutator transaction binding the contract method 0xda4d4fbf.
//
// Solidity: function withdrawEarningsAndLeftover() returns()
func (_CBJSale *CBJSaleSession) WithdrawEarningsAndLeftover() (*types.Transaction, error) {
	return _CBJSale.Contract.WithdrawEarningsAndLeftover(&_CBJSale.TransactOpts)
}

// WithdrawEarningsAndLeftover is a paid mutator transaction binding the contract method 0xda4d4fbf.
//
// Solidity: function withdrawEarningsAndLeftover() returns()
func (_CBJSale *CBJSaleTransactorSession) WithdrawEarningsAndLeftover() (*types.Transaction, error) {
	return _CBJSale.Contract.WithdrawEarningsAndLeftover(&_CBJSale.TransactOpts)
}

// WithdrawLeftover is a paid mutator transaction binding the contract method 0xa525d237.
//
// Solidity: function withdrawLeftover() returns()
func (_CBJSale *CBJSaleTransactor) WithdrawLeftover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CBJSale.contract.Transact(opts, "withdrawLeftover")
}

// WithdrawLeftover is a paid mutator transaction binding the contract method 0xa525d237.
//
// Solidity: function withdrawLeftover() returns()
func (_CBJSale *CBJSaleSession) WithdrawLeftover() (*types.Transaction, error) {
	return _CBJSale.Contract.WithdrawLeftover(&_CBJSale.TransactOpts)
}

// WithdrawLeftover is a paid mutator transaction binding the contract method 0xa525d237.
//
// Solidity: function withdrawLeftover() returns()
func (_CBJSale *CBJSaleTransactorSession) WithdrawLeftover() (*types.Transaction, error) {
	return _CBJSale.Contract.WithdrawLeftover(&_CBJSale.TransactOpts)
}

// WithdrawMultiplePortions is a paid mutator transaction binding the contract method 0x718af7e6.
//
// Solidity: function withdrawMultiplePortions(uint256[] portionIds) returns()
func (_CBJSale *CBJSaleTransactor) WithdrawMultiplePortions(opts *bind.TransactOpts, portionIds []*big.Int) (*types.Transaction, error) {
	return _CBJSale.contract.Transact(opts, "withdrawMultiplePortions", portionIds)
}

// WithdrawMultiplePortions is a paid mutator transaction binding the contract method 0x718af7e6.
//
// Solidity: function withdrawMultiplePortions(uint256[] portionIds) returns()
func (_CBJSale *CBJSaleSession) WithdrawMultiplePortions(portionIds []*big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.WithdrawMultiplePortions(&_CBJSale.TransactOpts, portionIds)
}

// WithdrawMultiplePortions is a paid mutator transaction binding the contract method 0x718af7e6.
//
// Solidity: function withdrawMultiplePortions(uint256[] portionIds) returns()
func (_CBJSale *CBJSaleTransactorSession) WithdrawMultiplePortions(portionIds []*big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.WithdrawMultiplePortions(&_CBJSale.TransactOpts, portionIds)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x315a095d.
//
// Solidity: function withdrawTokens(uint256 portionId) returns()
func (_CBJSale *CBJSaleTransactor) WithdrawTokens(opts *bind.TransactOpts, portionId *big.Int) (*types.Transaction, error) {
	return _CBJSale.contract.Transact(opts, "withdrawTokens", portionId)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x315a095d.
//
// Solidity: function withdrawTokens(uint256 portionId) returns()
func (_CBJSale *CBJSaleSession) WithdrawTokens(portionId *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.WithdrawTokens(&_CBJSale.TransactOpts, portionId)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x315a095d.
//
// Solidity: function withdrawTokens(uint256 portionId) returns()
func (_CBJSale *CBJSaleTransactorSession) WithdrawTokens(portionId *big.Int) (*types.Transaction, error) {
	return _CBJSale.Contract.WithdrawTokens(&_CBJSale.TransactOpts, portionId)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CBJSale *CBJSaleTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CBJSale.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CBJSale *CBJSaleSession) Receive() (*types.Transaction, error) {
	return _CBJSale.Contract.Receive(&_CBJSale.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CBJSale *CBJSaleTransactorSession) Receive() (*types.Transaction, error) {
	return _CBJSale.Contract.Receive(&_CBJSale.TransactOpts)
}

// CBJSaleMaxParticipationSetIterator is returned from FilterMaxParticipationSet and is used to iterate over the raw logs and unpacked data for MaxParticipationSet events raised by the CBJSale contract.
type CBJSaleMaxParticipationSetIterator struct {
	Event *CBJSaleMaxParticipationSet // Event containing the contract specifics and raw log

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
func (it *CBJSaleMaxParticipationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJSaleMaxParticipationSet)
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
		it.Event = new(CBJSaleMaxParticipationSet)
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
func (it *CBJSaleMaxParticipationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJSaleMaxParticipationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJSaleMaxParticipationSet represents a MaxParticipationSet event raised by the CBJSale contract.
type CBJSaleMaxParticipationSet struct {
	NewMaxParticipation *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMaxParticipationSet is a free log retrieval operation binding the contract event 0x37cfb0744bdb923d0300a0c37fac13cfec5fc2ee313cb9b217d284122980bada.
//
// Solidity: event MaxParticipationSet(uint256 newMaxParticipation)
func (_CBJSale *CBJSaleFilterer) FilterMaxParticipationSet(opts *bind.FilterOpts) (*CBJSaleMaxParticipationSetIterator, error) {

	logs, sub, err := _CBJSale.contract.FilterLogs(opts, "MaxParticipationSet")
	if err != nil {
		return nil, err
	}
	return &CBJSaleMaxParticipationSetIterator{contract: _CBJSale.contract, event: "MaxParticipationSet", logs: logs, sub: sub}, nil
}

// WatchMaxParticipationSet is a free log subscription operation binding the contract event 0x37cfb0744bdb923d0300a0c37fac13cfec5fc2ee313cb9b217d284122980bada.
//
// Solidity: event MaxParticipationSet(uint256 newMaxParticipation)
func (_CBJSale *CBJSaleFilterer) WatchMaxParticipationSet(opts *bind.WatchOpts, sink chan<- *CBJSaleMaxParticipationSet) (event.Subscription, error) {

	logs, sub, err := _CBJSale.contract.WatchLogs(opts, "MaxParticipationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJSaleMaxParticipationSet)
				if err := _CBJSale.contract.UnpackLog(event, "MaxParticipationSet", log); err != nil {
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

// ParseMaxParticipationSet is a log parse operation binding the contract event 0x37cfb0744bdb923d0300a0c37fac13cfec5fc2ee313cb9b217d284122980bada.
//
// Solidity: event MaxParticipationSet(uint256 newMaxParticipation)
func (_CBJSale *CBJSaleFilterer) ParseMaxParticipationSet(log types.Log) (*CBJSaleMaxParticipationSet, error) {
	event := new(CBJSaleMaxParticipationSet)
	if err := _CBJSale.contract.UnpackLog(event, "MaxParticipationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJSaleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CBJSale contract.
type CBJSaleOwnershipTransferredIterator struct {
	Event *CBJSaleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CBJSaleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJSaleOwnershipTransferred)
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
		it.Event = new(CBJSaleOwnershipTransferred)
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
func (it *CBJSaleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJSaleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJSaleOwnershipTransferred represents a OwnershipTransferred event raised by the CBJSale contract.
type CBJSaleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CBJSale *CBJSaleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CBJSaleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CBJSale.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CBJSaleOwnershipTransferredIterator{contract: _CBJSale.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CBJSale *CBJSaleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CBJSaleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CBJSale.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJSaleOwnershipTransferred)
				if err := _CBJSale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CBJSale *CBJSaleFilterer) ParseOwnershipTransferred(log types.Log) (*CBJSaleOwnershipTransferred, error) {
	event := new(CBJSaleOwnershipTransferred)
	if err := _CBJSale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJSaleRegistrationTimeSetIterator is returned from FilterRegistrationTimeSet and is used to iterate over the raw logs and unpacked data for RegistrationTimeSet events raised by the CBJSale contract.
type CBJSaleRegistrationTimeSetIterator struct {
	Event *CBJSaleRegistrationTimeSet // Event containing the contract specifics and raw log

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
func (it *CBJSaleRegistrationTimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJSaleRegistrationTimeSet)
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
		it.Event = new(CBJSaleRegistrationTimeSet)
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
func (it *CBJSaleRegistrationTimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJSaleRegistrationTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJSaleRegistrationTimeSet represents a RegistrationTimeSet event raised by the CBJSale contract.
type CBJSaleRegistrationTimeSet struct {
	RegistrationTimeStarts *big.Int
	RegistrationTimeEnds   *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRegistrationTimeSet is a free log retrieval operation binding the contract event 0xaf6e91c17885891414abbe7fa95074976b3e429980b0d8790034468e69650dd9.
//
// Solidity: event RegistrationTimeSet(uint256 registrationTimeStarts, uint256 registrationTimeEnds)
func (_CBJSale *CBJSaleFilterer) FilterRegistrationTimeSet(opts *bind.FilterOpts) (*CBJSaleRegistrationTimeSetIterator, error) {

	logs, sub, err := _CBJSale.contract.FilterLogs(opts, "RegistrationTimeSet")
	if err != nil {
		return nil, err
	}
	return &CBJSaleRegistrationTimeSetIterator{contract: _CBJSale.contract, event: "RegistrationTimeSet", logs: logs, sub: sub}, nil
}

// WatchRegistrationTimeSet is a free log subscription operation binding the contract event 0xaf6e91c17885891414abbe7fa95074976b3e429980b0d8790034468e69650dd9.
//
// Solidity: event RegistrationTimeSet(uint256 registrationTimeStarts, uint256 registrationTimeEnds)
func (_CBJSale *CBJSaleFilterer) WatchRegistrationTimeSet(opts *bind.WatchOpts, sink chan<- *CBJSaleRegistrationTimeSet) (event.Subscription, error) {

	logs, sub, err := _CBJSale.contract.WatchLogs(opts, "RegistrationTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJSaleRegistrationTimeSet)
				if err := _CBJSale.contract.UnpackLog(event, "RegistrationTimeSet", log); err != nil {
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

// ParseRegistrationTimeSet is a log parse operation binding the contract event 0xaf6e91c17885891414abbe7fa95074976b3e429980b0d8790034468e69650dd9.
//
// Solidity: event RegistrationTimeSet(uint256 registrationTimeStarts, uint256 registrationTimeEnds)
func (_CBJSale *CBJSaleFilterer) ParseRegistrationTimeSet(log types.Log) (*CBJSaleRegistrationTimeSet, error) {
	event := new(CBJSaleRegistrationTimeSet)
	if err := _CBJSale.contract.UnpackLog(event, "RegistrationTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJSaleSaleCreatedIterator is returned from FilterSaleCreated and is used to iterate over the raw logs and unpacked data for SaleCreated events raised by the CBJSale contract.
type CBJSaleSaleCreatedIterator struct {
	Event *CBJSaleSaleCreated // Event containing the contract specifics and raw log

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
func (it *CBJSaleSaleCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJSaleSaleCreated)
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
		it.Event = new(CBJSaleSaleCreated)
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
func (it *CBJSaleSaleCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJSaleSaleCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJSaleSaleCreated represents a SaleCreated event raised by the CBJSale contract.
type CBJSaleSaleCreated struct {
	SaleOwner            common.Address
	TokenPriceInETH      *big.Int
	AmountOfTokensToSell *big.Int
	SaleEnd              *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSaleCreated is a free log retrieval operation binding the contract event 0x2cd2dfcdeb2b58c4b80527e9df5e12da537fa4f6c958a4fb623a83ab74eeab63.
//
// Solidity: event SaleCreated(address saleOwner, uint256 tokenPriceInETH, uint256 amountOfTokensToSell, uint256 saleEnd)
func (_CBJSale *CBJSaleFilterer) FilterSaleCreated(opts *bind.FilterOpts) (*CBJSaleSaleCreatedIterator, error) {

	logs, sub, err := _CBJSale.contract.FilterLogs(opts, "SaleCreated")
	if err != nil {
		return nil, err
	}
	return &CBJSaleSaleCreatedIterator{contract: _CBJSale.contract, event: "SaleCreated", logs: logs, sub: sub}, nil
}

// WatchSaleCreated is a free log subscription operation binding the contract event 0x2cd2dfcdeb2b58c4b80527e9df5e12da537fa4f6c958a4fb623a83ab74eeab63.
//
// Solidity: event SaleCreated(address saleOwner, uint256 tokenPriceInETH, uint256 amountOfTokensToSell, uint256 saleEnd)
func (_CBJSale *CBJSaleFilterer) WatchSaleCreated(opts *bind.WatchOpts, sink chan<- *CBJSaleSaleCreated) (event.Subscription, error) {

	logs, sub, err := _CBJSale.contract.WatchLogs(opts, "SaleCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJSaleSaleCreated)
				if err := _CBJSale.contract.UnpackLog(event, "SaleCreated", log); err != nil {
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

// ParseSaleCreated is a log parse operation binding the contract event 0x2cd2dfcdeb2b58c4b80527e9df5e12da537fa4f6c958a4fb623a83ab74eeab63.
//
// Solidity: event SaleCreated(address saleOwner, uint256 tokenPriceInETH, uint256 amountOfTokensToSell, uint256 saleEnd)
func (_CBJSale *CBJSaleFilterer) ParseSaleCreated(log types.Log) (*CBJSaleSaleCreated, error) {
	event := new(CBJSaleSaleCreated)
	if err := _CBJSale.contract.UnpackLog(event, "SaleCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJSaleStartTimeSetIterator is returned from FilterStartTimeSet and is used to iterate over the raw logs and unpacked data for StartTimeSet events raised by the CBJSale contract.
type CBJSaleStartTimeSetIterator struct {
	Event *CBJSaleStartTimeSet // Event containing the contract specifics and raw log

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
func (it *CBJSaleStartTimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJSaleStartTimeSet)
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
		it.Event = new(CBJSaleStartTimeSet)
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
func (it *CBJSaleStartTimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJSaleStartTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJSaleStartTimeSet represents a StartTimeSet event raised by the CBJSale contract.
type CBJSaleStartTimeSet struct {
	StartTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStartTimeSet is a free log retrieval operation binding the contract event 0xaad53c4362ef2fe5a5390cc046e71fd8423a0a8dceebc156ac9bbcd15997eec2.
//
// Solidity: event StartTimeSet(uint256 startTime)
func (_CBJSale *CBJSaleFilterer) FilterStartTimeSet(opts *bind.FilterOpts) (*CBJSaleStartTimeSetIterator, error) {

	logs, sub, err := _CBJSale.contract.FilterLogs(opts, "StartTimeSet")
	if err != nil {
		return nil, err
	}
	return &CBJSaleStartTimeSetIterator{contract: _CBJSale.contract, event: "StartTimeSet", logs: logs, sub: sub}, nil
}

// WatchStartTimeSet is a free log subscription operation binding the contract event 0xaad53c4362ef2fe5a5390cc046e71fd8423a0a8dceebc156ac9bbcd15997eec2.
//
// Solidity: event StartTimeSet(uint256 startTime)
func (_CBJSale *CBJSaleFilterer) WatchStartTimeSet(opts *bind.WatchOpts, sink chan<- *CBJSaleStartTimeSet) (event.Subscription, error) {

	logs, sub, err := _CBJSale.contract.WatchLogs(opts, "StartTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJSaleStartTimeSet)
				if err := _CBJSale.contract.UnpackLog(event, "StartTimeSet", log); err != nil {
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

// ParseStartTimeSet is a log parse operation binding the contract event 0xaad53c4362ef2fe5a5390cc046e71fd8423a0a8dceebc156ac9bbcd15997eec2.
//
// Solidity: event StartTimeSet(uint256 startTime)
func (_CBJSale *CBJSaleFilterer) ParseStartTimeSet(log types.Log) (*CBJSaleStartTimeSet, error) {
	event := new(CBJSaleStartTimeSet)
	if err := _CBJSale.contract.UnpackLog(event, "StartTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJSaleTokenPriceSetIterator is returned from FilterTokenPriceSet and is used to iterate over the raw logs and unpacked data for TokenPriceSet events raised by the CBJSale contract.
type CBJSaleTokenPriceSetIterator struct {
	Event *CBJSaleTokenPriceSet // Event containing the contract specifics and raw log

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
func (it *CBJSaleTokenPriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJSaleTokenPriceSet)
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
		it.Event = new(CBJSaleTokenPriceSet)
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
func (it *CBJSaleTokenPriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJSaleTokenPriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJSaleTokenPriceSet represents a TokenPriceSet event raised by the CBJSale contract.
type CBJSaleTokenPriceSet struct {
	NewPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenPriceSet is a free log retrieval operation binding the contract event 0x4b59d61d9ffdc3db926d0ce7e06ebabb6bd1bf9dcdae262667e48be368227216.
//
// Solidity: event TokenPriceSet(uint256 newPrice)
func (_CBJSale *CBJSaleFilterer) FilterTokenPriceSet(opts *bind.FilterOpts) (*CBJSaleTokenPriceSetIterator, error) {

	logs, sub, err := _CBJSale.contract.FilterLogs(opts, "TokenPriceSet")
	if err != nil {
		return nil, err
	}
	return &CBJSaleTokenPriceSetIterator{contract: _CBJSale.contract, event: "TokenPriceSet", logs: logs, sub: sub}, nil
}

// WatchTokenPriceSet is a free log subscription operation binding the contract event 0x4b59d61d9ffdc3db926d0ce7e06ebabb6bd1bf9dcdae262667e48be368227216.
//
// Solidity: event TokenPriceSet(uint256 newPrice)
func (_CBJSale *CBJSaleFilterer) WatchTokenPriceSet(opts *bind.WatchOpts, sink chan<- *CBJSaleTokenPriceSet) (event.Subscription, error) {

	logs, sub, err := _CBJSale.contract.WatchLogs(opts, "TokenPriceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJSaleTokenPriceSet)
				if err := _CBJSale.contract.UnpackLog(event, "TokenPriceSet", log); err != nil {
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

// ParseTokenPriceSet is a log parse operation binding the contract event 0x4b59d61d9ffdc3db926d0ce7e06ebabb6bd1bf9dcdae262667e48be368227216.
//
// Solidity: event TokenPriceSet(uint256 newPrice)
func (_CBJSale *CBJSaleFilterer) ParseTokenPriceSet(log types.Log) (*CBJSaleTokenPriceSet, error) {
	event := new(CBJSaleTokenPriceSet)
	if err := _CBJSale.contract.UnpackLog(event, "TokenPriceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJSaleTokensSoldIterator is returned from FilterTokensSold and is used to iterate over the raw logs and unpacked data for TokensSold events raised by the CBJSale contract.
type CBJSaleTokensSoldIterator struct {
	Event *CBJSaleTokensSold // Event containing the contract specifics and raw log

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
func (it *CBJSaleTokensSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJSaleTokensSold)
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
		it.Event = new(CBJSaleTokensSold)
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
func (it *CBJSaleTokensSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJSaleTokensSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJSaleTokensSold represents a TokensSold event raised by the CBJSale contract.
type CBJSaleTokensSold struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensSold is a free log retrieval operation binding the contract event 0x57d61f3ccd4ccd25ec5d234d6049553a586fac134c85c98d0b0d9d5724f4e43e.
//
// Solidity: event TokensSold(address indexed user, uint256 amount)
func (_CBJSale *CBJSaleFilterer) FilterTokensSold(opts *bind.FilterOpts, user []common.Address) (*CBJSaleTokensSoldIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CBJSale.contract.FilterLogs(opts, "TokensSold", userRule)
	if err != nil {
		return nil, err
	}
	return &CBJSaleTokensSoldIterator{contract: _CBJSale.contract, event: "TokensSold", logs: logs, sub: sub}, nil
}

// WatchTokensSold is a free log subscription operation binding the contract event 0x57d61f3ccd4ccd25ec5d234d6049553a586fac134c85c98d0b0d9d5724f4e43e.
//
// Solidity: event TokensSold(address indexed user, uint256 amount)
func (_CBJSale *CBJSaleFilterer) WatchTokensSold(opts *bind.WatchOpts, sink chan<- *CBJSaleTokensSold, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CBJSale.contract.WatchLogs(opts, "TokensSold", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJSaleTokensSold)
				if err := _CBJSale.contract.UnpackLog(event, "TokensSold", log); err != nil {
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

// ParseTokensSold is a log parse operation binding the contract event 0x57d61f3ccd4ccd25ec5d234d6049553a586fac134c85c98d0b0d9d5724f4e43e.
//
// Solidity: event TokensSold(address indexed user, uint256 amount)
func (_CBJSale *CBJSaleFilterer) ParseTokensSold(log types.Log) (*CBJSaleTokensSold, error) {
	event := new(CBJSaleTokensSold)
	if err := _CBJSale.contract.UnpackLog(event, "TokensSold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJSaleTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the CBJSale contract.
type CBJSaleTokensWithdrawnIterator struct {
	Event *CBJSaleTokensWithdrawn // Event containing the contract specifics and raw log

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
func (it *CBJSaleTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJSaleTokensWithdrawn)
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
		it.Event = new(CBJSaleTokensWithdrawn)
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
func (it *CBJSaleTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJSaleTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJSaleTokensWithdrawn represents a TokensWithdrawn event raised by the CBJSale contract.
type CBJSaleTokensWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed user, uint256 amount)
func (_CBJSale *CBJSaleFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts, user []common.Address) (*CBJSaleTokensWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CBJSale.contract.FilterLogs(opts, "TokensWithdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &CBJSaleTokensWithdrawnIterator{contract: _CBJSale.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed user, uint256 amount)
func (_CBJSale *CBJSaleFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *CBJSaleTokensWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CBJSale.contract.WatchLogs(opts, "TokensWithdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJSaleTokensWithdrawn)
				if err := _CBJSale.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
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

// ParseTokensWithdrawn is a log parse operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed user, uint256 amount)
func (_CBJSale *CBJSaleFilterer) ParseTokensWithdrawn(log types.Log) (*CBJSaleTokensWithdrawn, error) {
	event := new(CBJSaleTokensWithdrawn)
	if err := _CBJSale.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJSaleUserRegisteredIterator is returned from FilterUserRegistered and is used to iterate over the raw logs and unpacked data for UserRegistered events raised by the CBJSale contract.
type CBJSaleUserRegisteredIterator struct {
	Event *CBJSaleUserRegistered // Event containing the contract specifics and raw log

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
func (it *CBJSaleUserRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJSaleUserRegistered)
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
		it.Event = new(CBJSaleUserRegistered)
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
func (it *CBJSaleUserRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJSaleUserRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJSaleUserRegistered represents a UserRegistered event raised by the CBJSale contract.
type CBJSaleUserRegistered struct {
	User common.Address
	Pid  *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUserRegistered is a free log retrieval operation binding the contract event 0xe29d35093005f4d575e1003753426b57a7f64378ba73332eef9c6ccc2b8decd6.
//
// Solidity: event UserRegistered(address indexed user, uint256 indexed pid)
func (_CBJSale *CBJSaleFilterer) FilterUserRegistered(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*CBJSaleUserRegisteredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _CBJSale.contract.FilterLogs(opts, "UserRegistered", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &CBJSaleUserRegisteredIterator{contract: _CBJSale.contract, event: "UserRegistered", logs: logs, sub: sub}, nil
}

// WatchUserRegistered is a free log subscription operation binding the contract event 0xe29d35093005f4d575e1003753426b57a7f64378ba73332eef9c6ccc2b8decd6.
//
// Solidity: event UserRegistered(address indexed user, uint256 indexed pid)
func (_CBJSale *CBJSaleFilterer) WatchUserRegistered(opts *bind.WatchOpts, sink chan<- *CBJSaleUserRegistered, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _CBJSale.contract.WatchLogs(opts, "UserRegistered", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJSaleUserRegistered)
				if err := _CBJSale.contract.UnpackLog(event, "UserRegistered", log); err != nil {
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

// ParseUserRegistered is a log parse operation binding the contract event 0xe29d35093005f4d575e1003753426b57a7f64378ba73332eef9c6ccc2b8decd6.
//
// Solidity: event UserRegistered(address indexed user, uint256 indexed pid)
func (_CBJSale *CBJSaleFilterer) ParseUserRegistered(log types.Log) (*CBJSaleUserRegistered, error) {
	event := new(CBJSaleUserRegistered)
	if err := _CBJSale.contract.UnpackLog(event, "UserRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
