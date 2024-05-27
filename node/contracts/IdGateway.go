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

// IdGatewayMetaData contains all meta data concerning the IdGateway contract.
var IdGatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_idRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_storageRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGuardian\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"Add\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"Remove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldStorageRegistry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newStorageRegistry\",\"type\":\"address\"}],\"name\":\"SetStorageRegistry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"REGISTER_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"addGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparatorV4\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"guardians\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isGuardian\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"structHash\",\"type\":\"bytes32\"}],\"name\":\"hashTypedDataV4\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"idRegistry\",\"outputs\":[{\"internalType\":\"contractIIdRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"extraStorage\",\"type\":\"uint256\"}],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"extraStorage\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overpayment\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"registerFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"extraStorage\",\"type\":\"uint256\"}],\"name\":\"registerFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overpayment\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"removeGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_storageRegistry\",\"type\":\"address\"}],\"name\":\"setStorageRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storageRegistry\",\"outputs\":[{\"internalType\":\"contractIStorageRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// IdGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use IdGatewayMetaData.ABI instead.
var IdGatewayABI = IdGatewayMetaData.ABI

// IdGateway is an auto generated Go binding around an Ethereum contract.
type IdGateway struct {
	IdGatewayCaller     // Read-only binding to the contract
	IdGatewayTransactor // Write-only binding to the contract
	IdGatewayFilterer   // Log filterer for contract events
}

// IdGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdGatewaySession struct {
	Contract     *IdGateway        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdGatewayCallerSession struct {
	Contract *IdGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IdGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdGatewayTransactorSession struct {
	Contract     *IdGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IdGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdGatewayRaw struct {
	Contract *IdGateway // Generic contract binding to access the raw methods on
}

// IdGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdGatewayCallerRaw struct {
	Contract *IdGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// IdGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdGatewayTransactorRaw struct {
	Contract *IdGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdGateway creates a new instance of IdGateway, bound to a specific deployed contract.
func NewIdGateway(address common.Address, backend bind.ContractBackend) (*IdGateway, error) {
	contract, err := bindIdGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdGateway{IdGatewayCaller: IdGatewayCaller{contract: contract}, IdGatewayTransactor: IdGatewayTransactor{contract: contract}, IdGatewayFilterer: IdGatewayFilterer{contract: contract}}, nil
}

// NewIdGatewayCaller creates a new read-only instance of IdGateway, bound to a specific deployed contract.
func NewIdGatewayCaller(address common.Address, caller bind.ContractCaller) (*IdGatewayCaller, error) {
	contract, err := bindIdGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdGatewayCaller{contract: contract}, nil
}

// NewIdGatewayTransactor creates a new write-only instance of IdGateway, bound to a specific deployed contract.
func NewIdGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*IdGatewayTransactor, error) {
	contract, err := bindIdGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdGatewayTransactor{contract: contract}, nil
}

// NewIdGatewayFilterer creates a new log filterer instance of IdGateway, bound to a specific deployed contract.
func NewIdGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*IdGatewayFilterer, error) {
	contract, err := bindIdGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdGatewayFilterer{contract: contract}, nil
}

// bindIdGateway binds a generic wrapper to an already deployed contract.
func bindIdGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IdGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdGateway *IdGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdGateway.Contract.IdGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdGateway *IdGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdGateway.Contract.IdGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdGateway *IdGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdGateway.Contract.IdGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdGateway *IdGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdGateway *IdGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdGateway *IdGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdGateway.Contract.contract.Transact(opts, method, params...)
}

// REGISTERTYPEHASH is a free data retrieval call binding the contract method 0x6a5306a3.
//
// Solidity: function REGISTER_TYPEHASH() view returns(bytes32)
func (_IdGateway *IdGatewayCaller) REGISTERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IdGateway.contract.Call(opts, &out, "REGISTER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REGISTERTYPEHASH is a free data retrieval call binding the contract method 0x6a5306a3.
//
// Solidity: function REGISTER_TYPEHASH() view returns(bytes32)
func (_IdGateway *IdGatewaySession) REGISTERTYPEHASH() ([32]byte, error) {
	return _IdGateway.Contract.REGISTERTYPEHASH(&_IdGateway.CallOpts)
}

// REGISTERTYPEHASH is a free data retrieval call binding the contract method 0x6a5306a3.
//
// Solidity: function REGISTER_TYPEHASH() view returns(bytes32)
func (_IdGateway *IdGatewayCallerSession) REGISTERTYPEHASH() ([32]byte, error) {
	return _IdGateway.Contract.REGISTERTYPEHASH(&_IdGateway.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_IdGateway *IdGatewayCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IdGateway.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_IdGateway *IdGatewaySession) VERSION() (string, error) {
	return _IdGateway.Contract.VERSION(&_IdGateway.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_IdGateway *IdGatewayCallerSession) VERSION() (string, error) {
	return _IdGateway.Contract.VERSION(&_IdGateway.CallOpts)
}

// DomainSeparatorV4 is a free data retrieval call binding the contract method 0x78e890ba.
//
// Solidity: function domainSeparatorV4() view returns(bytes32)
func (_IdGateway *IdGatewayCaller) DomainSeparatorV4(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IdGateway.contract.Call(opts, &out, "domainSeparatorV4")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparatorV4 is a free data retrieval call binding the contract method 0x78e890ba.
//
// Solidity: function domainSeparatorV4() view returns(bytes32)
func (_IdGateway *IdGatewaySession) DomainSeparatorV4() ([32]byte, error) {
	return _IdGateway.Contract.DomainSeparatorV4(&_IdGateway.CallOpts)
}

// DomainSeparatorV4 is a free data retrieval call binding the contract method 0x78e890ba.
//
// Solidity: function domainSeparatorV4() view returns(bytes32)
func (_IdGateway *IdGatewayCallerSession) DomainSeparatorV4() ([32]byte, error) {
	return _IdGateway.Contract.DomainSeparatorV4(&_IdGateway.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_IdGateway *IdGatewayCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _IdGateway.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_IdGateway *IdGatewaySession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _IdGateway.Contract.Eip712Domain(&_IdGateway.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_IdGateway *IdGatewayCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _IdGateway.Contract.Eip712Domain(&_IdGateway.CallOpts)
}

// Guardians is a free data retrieval call binding the contract method 0x0633b14a.
//
// Solidity: function guardians(address guardian) view returns(bool isGuardian)
func (_IdGateway *IdGatewayCaller) Guardians(opts *bind.CallOpts, guardian common.Address) (bool, error) {
	var out []interface{}
	err := _IdGateway.contract.Call(opts, &out, "guardians", guardian)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Guardians is a free data retrieval call binding the contract method 0x0633b14a.
//
// Solidity: function guardians(address guardian) view returns(bool isGuardian)
func (_IdGateway *IdGatewaySession) Guardians(guardian common.Address) (bool, error) {
	return _IdGateway.Contract.Guardians(&_IdGateway.CallOpts, guardian)
}

// Guardians is a free data retrieval call binding the contract method 0x0633b14a.
//
// Solidity: function guardians(address guardian) view returns(bool isGuardian)
func (_IdGateway *IdGatewayCallerSession) Guardians(guardian common.Address) (bool, error) {
	return _IdGateway.Contract.Guardians(&_IdGateway.CallOpts, guardian)
}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x4980f288.
//
// Solidity: function hashTypedDataV4(bytes32 structHash) view returns(bytes32)
func (_IdGateway *IdGatewayCaller) HashTypedDataV4(opts *bind.CallOpts, structHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IdGateway.contract.Call(opts, &out, "hashTypedDataV4", structHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x4980f288.
//
// Solidity: function hashTypedDataV4(bytes32 structHash) view returns(bytes32)
func (_IdGateway *IdGatewaySession) HashTypedDataV4(structHash [32]byte) ([32]byte, error) {
	return _IdGateway.Contract.HashTypedDataV4(&_IdGateway.CallOpts, structHash)
}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x4980f288.
//
// Solidity: function hashTypedDataV4(bytes32 structHash) view returns(bytes32)
func (_IdGateway *IdGatewayCallerSession) HashTypedDataV4(structHash [32]byte) ([32]byte, error) {
	return _IdGateway.Contract.HashTypedDataV4(&_IdGateway.CallOpts, structHash)
}

// IdRegistry is a free data retrieval call binding the contract method 0x0aa13b8c.
//
// Solidity: function idRegistry() view returns(address)
func (_IdGateway *IdGatewayCaller) IdRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IdGateway.contract.Call(opts, &out, "idRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdRegistry is a free data retrieval call binding the contract method 0x0aa13b8c.
//
// Solidity: function idRegistry() view returns(address)
func (_IdGateway *IdGatewaySession) IdRegistry() (common.Address, error) {
	return _IdGateway.Contract.IdRegistry(&_IdGateway.CallOpts)
}

// IdRegistry is a free data retrieval call binding the contract method 0x0aa13b8c.
//
// Solidity: function idRegistry() view returns(address)
func (_IdGateway *IdGatewayCallerSession) IdRegistry() (common.Address, error) {
	return _IdGateway.Contract.IdRegistry(&_IdGateway.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IdGateway *IdGatewayCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IdGateway.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IdGateway *IdGatewaySession) Nonces(owner common.Address) (*big.Int, error) {
	return _IdGateway.Contract.Nonces(&_IdGateway.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IdGateway *IdGatewayCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IdGateway.Contract.Nonces(&_IdGateway.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdGateway *IdGatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IdGateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdGateway *IdGatewaySession) Owner() (common.Address, error) {
	return _IdGateway.Contract.Owner(&_IdGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdGateway *IdGatewayCallerSession) Owner() (common.Address, error) {
	return _IdGateway.Contract.Owner(&_IdGateway.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IdGateway *IdGatewayCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IdGateway.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IdGateway *IdGatewaySession) Paused() (bool, error) {
	return _IdGateway.Contract.Paused(&_IdGateway.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IdGateway *IdGatewayCallerSession) Paused() (bool, error) {
	return _IdGateway.Contract.Paused(&_IdGateway.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_IdGateway *IdGatewayCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IdGateway.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_IdGateway *IdGatewaySession) PendingOwner() (common.Address, error) {
	return _IdGateway.Contract.PendingOwner(&_IdGateway.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_IdGateway *IdGatewayCallerSession) PendingOwner() (common.Address, error) {
	return _IdGateway.Contract.PendingOwner(&_IdGateway.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0x26a49e37.
//
// Solidity: function price(uint256 extraStorage) view returns(uint256)
func (_IdGateway *IdGatewayCaller) Price(opts *bind.CallOpts, extraStorage *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IdGateway.contract.Call(opts, &out, "price", extraStorage)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0x26a49e37.
//
// Solidity: function price(uint256 extraStorage) view returns(uint256)
func (_IdGateway *IdGatewaySession) Price(extraStorage *big.Int) (*big.Int, error) {
	return _IdGateway.Contract.Price(&_IdGateway.CallOpts, extraStorage)
}

// Price is a free data retrieval call binding the contract method 0x26a49e37.
//
// Solidity: function price(uint256 extraStorage) view returns(uint256)
func (_IdGateway *IdGatewayCallerSession) Price(extraStorage *big.Int) (*big.Int, error) {
	return _IdGateway.Contract.Price(&_IdGateway.CallOpts, extraStorage)
}

// Price0 is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_IdGateway *IdGatewayCaller) Price0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdGateway.contract.Call(opts, &out, "price0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price0 is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_IdGateway *IdGatewaySession) Price0() (*big.Int, error) {
	return _IdGateway.Contract.Price0(&_IdGateway.CallOpts)
}

// Price0 is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_IdGateway *IdGatewayCallerSession) Price0() (*big.Int, error) {
	return _IdGateway.Contract.Price0(&_IdGateway.CallOpts)
}

// StorageRegistry is a free data retrieval call binding the contract method 0x4ec77b45.
//
// Solidity: function storageRegistry() view returns(address)
func (_IdGateway *IdGatewayCaller) StorageRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IdGateway.contract.Call(opts, &out, "storageRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StorageRegistry is a free data retrieval call binding the contract method 0x4ec77b45.
//
// Solidity: function storageRegistry() view returns(address)
func (_IdGateway *IdGatewaySession) StorageRegistry() (common.Address, error) {
	return _IdGateway.Contract.StorageRegistry(&_IdGateway.CallOpts)
}

// StorageRegistry is a free data retrieval call binding the contract method 0x4ec77b45.
//
// Solidity: function storageRegistry() view returns(address)
func (_IdGateway *IdGatewayCallerSession) StorageRegistry() (common.Address, error) {
	return _IdGateway.Contract.StorageRegistry(&_IdGateway.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_IdGateway *IdGatewayTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdGateway.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_IdGateway *IdGatewaySession) AcceptOwnership() (*types.Transaction, error) {
	return _IdGateway.Contract.AcceptOwnership(&_IdGateway.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_IdGateway *IdGatewayTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _IdGateway.Contract.AcceptOwnership(&_IdGateway.TransactOpts)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_IdGateway *IdGatewayTransactor) AddGuardian(opts *bind.TransactOpts, guardian common.Address) (*types.Transaction, error) {
	return _IdGateway.contract.Transact(opts, "addGuardian", guardian)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_IdGateway *IdGatewaySession) AddGuardian(guardian common.Address) (*types.Transaction, error) {
	return _IdGateway.Contract.AddGuardian(&_IdGateway.TransactOpts, guardian)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_IdGateway *IdGatewayTransactorSession) AddGuardian(guardian common.Address) (*types.Transaction, error) {
	return _IdGateway.Contract.AddGuardian(&_IdGateway.TransactOpts, guardian)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IdGateway *IdGatewayTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdGateway.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IdGateway *IdGatewaySession) Pause() (*types.Transaction, error) {
	return _IdGateway.Contract.Pause(&_IdGateway.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IdGateway *IdGatewayTransactorSession) Pause() (*types.Transaction, error) {
	return _IdGateway.Contract.Pause(&_IdGateway.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address recovery) payable returns(uint256, uint256)
func (_IdGateway *IdGatewayTransactor) Register(opts *bind.TransactOpts, recovery common.Address) (*types.Transaction, error) {
	return _IdGateway.contract.Transact(opts, "register", recovery)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address recovery) payable returns(uint256, uint256)
func (_IdGateway *IdGatewaySession) Register(recovery common.Address) (*types.Transaction, error) {
	return _IdGateway.Contract.Register(&_IdGateway.TransactOpts, recovery)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address recovery) payable returns(uint256, uint256)
func (_IdGateway *IdGatewayTransactorSession) Register(recovery common.Address) (*types.Transaction, error) {
	return _IdGateway.Contract.Register(&_IdGateway.TransactOpts, recovery)
}

// Register0 is a paid mutator transaction binding the contract method 0x6d705ebb.
//
// Solidity: function register(address recovery, uint256 extraStorage) payable returns(uint256 fid, uint256 overpayment)
func (_IdGateway *IdGatewayTransactor) Register0(opts *bind.TransactOpts, recovery common.Address, extraStorage *big.Int) (*types.Transaction, error) {
	return _IdGateway.contract.Transact(opts, "register0", recovery, extraStorage)
}

// Register0 is a paid mutator transaction binding the contract method 0x6d705ebb.
//
// Solidity: function register(address recovery, uint256 extraStorage) payable returns(uint256 fid, uint256 overpayment)
func (_IdGateway *IdGatewaySession) Register0(recovery common.Address, extraStorage *big.Int) (*types.Transaction, error) {
	return _IdGateway.Contract.Register0(&_IdGateway.TransactOpts, recovery, extraStorage)
}

// Register0 is a paid mutator transaction binding the contract method 0x6d705ebb.
//
// Solidity: function register(address recovery, uint256 extraStorage) payable returns(uint256 fid, uint256 overpayment)
func (_IdGateway *IdGatewayTransactorSession) Register0(recovery common.Address, extraStorage *big.Int) (*types.Transaction, error) {
	return _IdGateway.Contract.Register0(&_IdGateway.TransactOpts, recovery, extraStorage)
}

// RegisterFor is a paid mutator transaction binding the contract method 0x3efa0b02.
//
// Solidity: function registerFor(address to, address recovery, uint256 deadline, bytes sig) payable returns(uint256, uint256)
func (_IdGateway *IdGatewayTransactor) RegisterFor(opts *bind.TransactOpts, to common.Address, recovery common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _IdGateway.contract.Transact(opts, "registerFor", to, recovery, deadline, sig)
}

// RegisterFor is a paid mutator transaction binding the contract method 0x3efa0b02.
//
// Solidity: function registerFor(address to, address recovery, uint256 deadline, bytes sig) payable returns(uint256, uint256)
func (_IdGateway *IdGatewaySession) RegisterFor(to common.Address, recovery common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _IdGateway.Contract.RegisterFor(&_IdGateway.TransactOpts, to, recovery, deadline, sig)
}

// RegisterFor is a paid mutator transaction binding the contract method 0x3efa0b02.
//
// Solidity: function registerFor(address to, address recovery, uint256 deadline, bytes sig) payable returns(uint256, uint256)
func (_IdGateway *IdGatewayTransactorSession) RegisterFor(to common.Address, recovery common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _IdGateway.Contract.RegisterFor(&_IdGateway.TransactOpts, to, recovery, deadline, sig)
}

// RegisterFor0 is a paid mutator transaction binding the contract method 0xa0c7529c.
//
// Solidity: function registerFor(address to, address recovery, uint256 deadline, bytes sig, uint256 extraStorage) payable returns(uint256 fid, uint256 overpayment)
func (_IdGateway *IdGatewayTransactor) RegisterFor0(opts *bind.TransactOpts, to common.Address, recovery common.Address, deadline *big.Int, sig []byte, extraStorage *big.Int) (*types.Transaction, error) {
	return _IdGateway.contract.Transact(opts, "registerFor0", to, recovery, deadline, sig, extraStorage)
}

// RegisterFor0 is a paid mutator transaction binding the contract method 0xa0c7529c.
//
// Solidity: function registerFor(address to, address recovery, uint256 deadline, bytes sig, uint256 extraStorage) payable returns(uint256 fid, uint256 overpayment)
func (_IdGateway *IdGatewaySession) RegisterFor0(to common.Address, recovery common.Address, deadline *big.Int, sig []byte, extraStorage *big.Int) (*types.Transaction, error) {
	return _IdGateway.Contract.RegisterFor0(&_IdGateway.TransactOpts, to, recovery, deadline, sig, extraStorage)
}

// RegisterFor0 is a paid mutator transaction binding the contract method 0xa0c7529c.
//
// Solidity: function registerFor(address to, address recovery, uint256 deadline, bytes sig, uint256 extraStorage) payable returns(uint256 fid, uint256 overpayment)
func (_IdGateway *IdGatewayTransactorSession) RegisterFor0(to common.Address, recovery common.Address, deadline *big.Int, sig []byte, extraStorage *big.Int) (*types.Transaction, error) {
	return _IdGateway.Contract.RegisterFor0(&_IdGateway.TransactOpts, to, recovery, deadline, sig, extraStorage)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_IdGateway *IdGatewayTransactor) RemoveGuardian(opts *bind.TransactOpts, guardian common.Address) (*types.Transaction, error) {
	return _IdGateway.contract.Transact(opts, "removeGuardian", guardian)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_IdGateway *IdGatewaySession) RemoveGuardian(guardian common.Address) (*types.Transaction, error) {
	return _IdGateway.Contract.RemoveGuardian(&_IdGateway.TransactOpts, guardian)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_IdGateway *IdGatewayTransactorSession) RemoveGuardian(guardian common.Address) (*types.Transaction, error) {
	return _IdGateway.Contract.RemoveGuardian(&_IdGateway.TransactOpts, guardian)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IdGateway *IdGatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdGateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IdGateway *IdGatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _IdGateway.Contract.RenounceOwnership(&_IdGateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IdGateway *IdGatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _IdGateway.Contract.RenounceOwnership(&_IdGateway.TransactOpts)
}

// SetStorageRegistry is a paid mutator transaction binding the contract method 0xfa0b7dc8.
//
// Solidity: function setStorageRegistry(address _storageRegistry) returns()
func (_IdGateway *IdGatewayTransactor) SetStorageRegistry(opts *bind.TransactOpts, _storageRegistry common.Address) (*types.Transaction, error) {
	return _IdGateway.contract.Transact(opts, "setStorageRegistry", _storageRegistry)
}

// SetStorageRegistry is a paid mutator transaction binding the contract method 0xfa0b7dc8.
//
// Solidity: function setStorageRegistry(address _storageRegistry) returns()
func (_IdGateway *IdGatewaySession) SetStorageRegistry(_storageRegistry common.Address) (*types.Transaction, error) {
	return _IdGateway.Contract.SetStorageRegistry(&_IdGateway.TransactOpts, _storageRegistry)
}

// SetStorageRegistry is a paid mutator transaction binding the contract method 0xfa0b7dc8.
//
// Solidity: function setStorageRegistry(address _storageRegistry) returns()
func (_IdGateway *IdGatewayTransactorSession) SetStorageRegistry(_storageRegistry common.Address) (*types.Transaction, error) {
	return _IdGateway.Contract.SetStorageRegistry(&_IdGateway.TransactOpts, _storageRegistry)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IdGateway *IdGatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IdGateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IdGateway *IdGatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IdGateway.Contract.TransferOwnership(&_IdGateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IdGateway *IdGatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IdGateway.Contract.TransferOwnership(&_IdGateway.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IdGateway *IdGatewayTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdGateway.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IdGateway *IdGatewaySession) Unpause() (*types.Transaction, error) {
	return _IdGateway.Contract.Unpause(&_IdGateway.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IdGateway *IdGatewayTransactorSession) Unpause() (*types.Transaction, error) {
	return _IdGateway.Contract.Unpause(&_IdGateway.TransactOpts)
}

// UseNonce is a paid mutator transaction binding the contract method 0x69615a4c.
//
// Solidity: function useNonce() returns(uint256)
func (_IdGateway *IdGatewayTransactor) UseNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdGateway.contract.Transact(opts, "useNonce")
}

// UseNonce is a paid mutator transaction binding the contract method 0x69615a4c.
//
// Solidity: function useNonce() returns(uint256)
func (_IdGateway *IdGatewaySession) UseNonce() (*types.Transaction, error) {
	return _IdGateway.Contract.UseNonce(&_IdGateway.TransactOpts)
}

// UseNonce is a paid mutator transaction binding the contract method 0x69615a4c.
//
// Solidity: function useNonce() returns(uint256)
func (_IdGateway *IdGatewayTransactorSession) UseNonce() (*types.Transaction, error) {
	return _IdGateway.Contract.UseNonce(&_IdGateway.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IdGateway *IdGatewayTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdGateway.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IdGateway *IdGatewaySession) Receive() (*types.Transaction, error) {
	return _IdGateway.Contract.Receive(&_IdGateway.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IdGateway *IdGatewayTransactorSession) Receive() (*types.Transaction, error) {
	return _IdGateway.Contract.Receive(&_IdGateway.TransactOpts)
}

// IdGatewayAddIterator is returned from FilterAdd and is used to iterate over the raw logs and unpacked data for Add events raised by the IdGateway contract.
type IdGatewayAddIterator struct {
	Event *IdGatewayAdd // Event containing the contract specifics and raw log

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
func (it *IdGatewayAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdGatewayAdd)
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
		it.Event = new(IdGatewayAdd)
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
func (it *IdGatewayAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdGatewayAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdGatewayAdd represents a Add event raised by the IdGateway contract.
type IdGatewayAdd struct {
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdd is a free log retrieval operation binding the contract event 0x87dc5eecd6d6bdeae407c426da6bfba5b7190befc554ed5d4d62dd5cf939fbae.
//
// Solidity: event Add(address indexed guardian)
func (_IdGateway *IdGatewayFilterer) FilterAdd(opts *bind.FilterOpts, guardian []common.Address) (*IdGatewayAddIterator, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _IdGateway.contract.FilterLogs(opts, "Add", guardianRule)
	if err != nil {
		return nil, err
	}
	return &IdGatewayAddIterator{contract: _IdGateway.contract, event: "Add", logs: logs, sub: sub}, nil
}

// WatchAdd is a free log subscription operation binding the contract event 0x87dc5eecd6d6bdeae407c426da6bfba5b7190befc554ed5d4d62dd5cf939fbae.
//
// Solidity: event Add(address indexed guardian)
func (_IdGateway *IdGatewayFilterer) WatchAdd(opts *bind.WatchOpts, sink chan<- *IdGatewayAdd, guardian []common.Address) (event.Subscription, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _IdGateway.contract.WatchLogs(opts, "Add", guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdGatewayAdd)
				if err := _IdGateway.contract.UnpackLog(event, "Add", log); err != nil {
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

// ParseAdd is a log parse operation binding the contract event 0x87dc5eecd6d6bdeae407c426da6bfba5b7190befc554ed5d4d62dd5cf939fbae.
//
// Solidity: event Add(address indexed guardian)
func (_IdGateway *IdGatewayFilterer) ParseAdd(log types.Log) (*IdGatewayAdd, error) {
	event := new(IdGatewayAdd)
	if err := _IdGateway.contract.UnpackLog(event, "Add", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdGatewayEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the IdGateway contract.
type IdGatewayEIP712DomainChangedIterator struct {
	Event *IdGatewayEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *IdGatewayEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdGatewayEIP712DomainChanged)
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
		it.Event = new(IdGatewayEIP712DomainChanged)
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
func (it *IdGatewayEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdGatewayEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdGatewayEIP712DomainChanged represents a EIP712DomainChanged event raised by the IdGateway contract.
type IdGatewayEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_IdGateway *IdGatewayFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*IdGatewayEIP712DomainChangedIterator, error) {

	logs, sub, err := _IdGateway.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &IdGatewayEIP712DomainChangedIterator{contract: _IdGateway.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_IdGateway *IdGatewayFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *IdGatewayEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _IdGateway.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdGatewayEIP712DomainChanged)
				if err := _IdGateway.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_IdGateway *IdGatewayFilterer) ParseEIP712DomainChanged(log types.Log) (*IdGatewayEIP712DomainChanged, error) {
	event := new(IdGatewayEIP712DomainChanged)
	if err := _IdGateway.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdGatewayOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the IdGateway contract.
type IdGatewayOwnershipTransferStartedIterator struct {
	Event *IdGatewayOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *IdGatewayOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdGatewayOwnershipTransferStarted)
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
		it.Event = new(IdGatewayOwnershipTransferStarted)
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
func (it *IdGatewayOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdGatewayOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdGatewayOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the IdGateway contract.
type IdGatewayOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_IdGateway *IdGatewayFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IdGatewayOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdGateway.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IdGatewayOwnershipTransferStartedIterator{contract: _IdGateway.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_IdGateway *IdGatewayFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *IdGatewayOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdGateway.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdGatewayOwnershipTransferStarted)
				if err := _IdGateway.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_IdGateway *IdGatewayFilterer) ParseOwnershipTransferStarted(log types.Log) (*IdGatewayOwnershipTransferStarted, error) {
	event := new(IdGatewayOwnershipTransferStarted)
	if err := _IdGateway.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdGatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IdGateway contract.
type IdGatewayOwnershipTransferredIterator struct {
	Event *IdGatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IdGatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdGatewayOwnershipTransferred)
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
		it.Event = new(IdGatewayOwnershipTransferred)
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
func (it *IdGatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdGatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdGatewayOwnershipTransferred represents a OwnershipTransferred event raised by the IdGateway contract.
type IdGatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IdGateway *IdGatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IdGatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdGateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IdGatewayOwnershipTransferredIterator{contract: _IdGateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IdGateway *IdGatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IdGatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdGateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdGatewayOwnershipTransferred)
				if err := _IdGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_IdGateway *IdGatewayFilterer) ParseOwnershipTransferred(log types.Log) (*IdGatewayOwnershipTransferred, error) {
	event := new(IdGatewayOwnershipTransferred)
	if err := _IdGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdGatewayPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the IdGateway contract.
type IdGatewayPausedIterator struct {
	Event *IdGatewayPaused // Event containing the contract specifics and raw log

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
func (it *IdGatewayPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdGatewayPaused)
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
		it.Event = new(IdGatewayPaused)
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
func (it *IdGatewayPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdGatewayPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdGatewayPaused represents a Paused event raised by the IdGateway contract.
type IdGatewayPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IdGateway *IdGatewayFilterer) FilterPaused(opts *bind.FilterOpts) (*IdGatewayPausedIterator, error) {

	logs, sub, err := _IdGateway.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &IdGatewayPausedIterator{contract: _IdGateway.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IdGateway *IdGatewayFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *IdGatewayPaused) (event.Subscription, error) {

	logs, sub, err := _IdGateway.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdGatewayPaused)
				if err := _IdGateway.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IdGateway *IdGatewayFilterer) ParsePaused(log types.Log) (*IdGatewayPaused, error) {
	event := new(IdGatewayPaused)
	if err := _IdGateway.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdGatewayRemoveIterator is returned from FilterRemove and is used to iterate over the raw logs and unpacked data for Remove events raised by the IdGateway contract.
type IdGatewayRemoveIterator struct {
	Event *IdGatewayRemove // Event containing the contract specifics and raw log

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
func (it *IdGatewayRemoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdGatewayRemove)
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
		it.Event = new(IdGatewayRemove)
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
func (it *IdGatewayRemoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdGatewayRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdGatewayRemove represents a Remove event raised by the IdGateway contract.
type IdGatewayRemove struct {
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemove is a free log retrieval operation binding the contract event 0xbe7c7ac3248df4581c206a84aab3cb4e7d521b5398b42b681757f78a5a7d411e.
//
// Solidity: event Remove(address indexed guardian)
func (_IdGateway *IdGatewayFilterer) FilterRemove(opts *bind.FilterOpts, guardian []common.Address) (*IdGatewayRemoveIterator, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _IdGateway.contract.FilterLogs(opts, "Remove", guardianRule)
	if err != nil {
		return nil, err
	}
	return &IdGatewayRemoveIterator{contract: _IdGateway.contract, event: "Remove", logs: logs, sub: sub}, nil
}

// WatchRemove is a free log subscription operation binding the contract event 0xbe7c7ac3248df4581c206a84aab3cb4e7d521b5398b42b681757f78a5a7d411e.
//
// Solidity: event Remove(address indexed guardian)
func (_IdGateway *IdGatewayFilterer) WatchRemove(opts *bind.WatchOpts, sink chan<- *IdGatewayRemove, guardian []common.Address) (event.Subscription, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _IdGateway.contract.WatchLogs(opts, "Remove", guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdGatewayRemove)
				if err := _IdGateway.contract.UnpackLog(event, "Remove", log); err != nil {
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

// ParseRemove is a log parse operation binding the contract event 0xbe7c7ac3248df4581c206a84aab3cb4e7d521b5398b42b681757f78a5a7d411e.
//
// Solidity: event Remove(address indexed guardian)
func (_IdGateway *IdGatewayFilterer) ParseRemove(log types.Log) (*IdGatewayRemove, error) {
	event := new(IdGatewayRemove)
	if err := _IdGateway.contract.UnpackLog(event, "Remove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdGatewaySetStorageRegistryIterator is returned from FilterSetStorageRegistry and is used to iterate over the raw logs and unpacked data for SetStorageRegistry events raised by the IdGateway contract.
type IdGatewaySetStorageRegistryIterator struct {
	Event *IdGatewaySetStorageRegistry // Event containing the contract specifics and raw log

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
func (it *IdGatewaySetStorageRegistryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdGatewaySetStorageRegistry)
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
		it.Event = new(IdGatewaySetStorageRegistry)
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
func (it *IdGatewaySetStorageRegistryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdGatewaySetStorageRegistryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdGatewaySetStorageRegistry represents a SetStorageRegistry event raised by the IdGateway contract.
type IdGatewaySetStorageRegistry struct {
	OldStorageRegistry common.Address
	NewStorageRegistry common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSetStorageRegistry is a free log retrieval operation binding the contract event 0x02fe9e072322317117700d5b8e41348ab279dac34a90a8302038a8e1ce324628.
//
// Solidity: event SetStorageRegistry(address oldStorageRegistry, address newStorageRegistry)
func (_IdGateway *IdGatewayFilterer) FilterSetStorageRegistry(opts *bind.FilterOpts) (*IdGatewaySetStorageRegistryIterator, error) {

	logs, sub, err := _IdGateway.contract.FilterLogs(opts, "SetStorageRegistry")
	if err != nil {
		return nil, err
	}
	return &IdGatewaySetStorageRegistryIterator{contract: _IdGateway.contract, event: "SetStorageRegistry", logs: logs, sub: sub}, nil
}

// WatchSetStorageRegistry is a free log subscription operation binding the contract event 0x02fe9e072322317117700d5b8e41348ab279dac34a90a8302038a8e1ce324628.
//
// Solidity: event SetStorageRegistry(address oldStorageRegistry, address newStorageRegistry)
func (_IdGateway *IdGatewayFilterer) WatchSetStorageRegistry(opts *bind.WatchOpts, sink chan<- *IdGatewaySetStorageRegistry) (event.Subscription, error) {

	logs, sub, err := _IdGateway.contract.WatchLogs(opts, "SetStorageRegistry")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdGatewaySetStorageRegistry)
				if err := _IdGateway.contract.UnpackLog(event, "SetStorageRegistry", log); err != nil {
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

// ParseSetStorageRegistry is a log parse operation binding the contract event 0x02fe9e072322317117700d5b8e41348ab279dac34a90a8302038a8e1ce324628.
//
// Solidity: event SetStorageRegistry(address oldStorageRegistry, address newStorageRegistry)
func (_IdGateway *IdGatewayFilterer) ParseSetStorageRegistry(log types.Log) (*IdGatewaySetStorageRegistry, error) {
	event := new(IdGatewaySetStorageRegistry)
	if err := _IdGateway.contract.UnpackLog(event, "SetStorageRegistry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdGatewayUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the IdGateway contract.
type IdGatewayUnpausedIterator struct {
	Event *IdGatewayUnpaused // Event containing the contract specifics and raw log

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
func (it *IdGatewayUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdGatewayUnpaused)
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
		it.Event = new(IdGatewayUnpaused)
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
func (it *IdGatewayUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdGatewayUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdGatewayUnpaused represents a Unpaused event raised by the IdGateway contract.
type IdGatewayUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IdGateway *IdGatewayFilterer) FilterUnpaused(opts *bind.FilterOpts) (*IdGatewayUnpausedIterator, error) {

	logs, sub, err := _IdGateway.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &IdGatewayUnpausedIterator{contract: _IdGateway.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IdGateway *IdGatewayFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *IdGatewayUnpaused) (event.Subscription, error) {

	logs, sub, err := _IdGateway.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdGatewayUnpaused)
				if err := _IdGateway.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IdGateway *IdGatewayFilterer) ParseUnpaused(log types.Log) (*IdGatewayUnpaused, error) {
	event := new(IdGatewayUnpaused)
	if err := _IdGateway.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
