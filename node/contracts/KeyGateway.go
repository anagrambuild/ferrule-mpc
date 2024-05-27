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

// KeyGatewayMetaData contains all meta data concerning the KeyGateway contract.
var KeyGatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_keyRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGuardian\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"Add\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"Remove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADD_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"keyType\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"metadataType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fidOwner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"keyType\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"metadataType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"addFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"addGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparatorV4\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"guardians\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isGuardian\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"structHash\",\"type\":\"bytes32\"}],\"name\":\"hashTypedDataV4\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keyRegistry\",\"outputs\":[{\"internalType\":\"contractIKeyRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"removeGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// KeyGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use KeyGatewayMetaData.ABI instead.
var KeyGatewayABI = KeyGatewayMetaData.ABI

// KeyGateway is an auto generated Go binding around an Ethereum contract.
type KeyGateway struct {
	KeyGatewayCaller     // Read-only binding to the contract
	KeyGatewayTransactor // Write-only binding to the contract
	KeyGatewayFilterer   // Log filterer for contract events
}

// KeyGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeyGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeyGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KeyGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeyGatewaySession struct {
	Contract     *KeyGateway       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KeyGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeyGatewayCallerSession struct {
	Contract *KeyGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// KeyGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeyGatewayTransactorSession struct {
	Contract     *KeyGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// KeyGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeyGatewayRaw struct {
	Contract *KeyGateway // Generic contract binding to access the raw methods on
}

// KeyGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeyGatewayCallerRaw struct {
	Contract *KeyGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// KeyGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeyGatewayTransactorRaw struct {
	Contract *KeyGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeyGateway creates a new instance of KeyGateway, bound to a specific deployed contract.
func NewKeyGateway(address common.Address, backend bind.ContractBackend) (*KeyGateway, error) {
	contract, err := bindKeyGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KeyGateway{KeyGatewayCaller: KeyGatewayCaller{contract: contract}, KeyGatewayTransactor: KeyGatewayTransactor{contract: contract}, KeyGatewayFilterer: KeyGatewayFilterer{contract: contract}}, nil
}

// NewKeyGatewayCaller creates a new read-only instance of KeyGateway, bound to a specific deployed contract.
func NewKeyGatewayCaller(address common.Address, caller bind.ContractCaller) (*KeyGatewayCaller, error) {
	contract, err := bindKeyGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeyGatewayCaller{contract: contract}, nil
}

// NewKeyGatewayTransactor creates a new write-only instance of KeyGateway, bound to a specific deployed contract.
func NewKeyGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*KeyGatewayTransactor, error) {
	contract, err := bindKeyGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeyGatewayTransactor{contract: contract}, nil
}

// NewKeyGatewayFilterer creates a new log filterer instance of KeyGateway, bound to a specific deployed contract.
func NewKeyGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*KeyGatewayFilterer, error) {
	contract, err := bindKeyGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeyGatewayFilterer{contract: contract}, nil
}

// bindKeyGateway binds a generic wrapper to an already deployed contract.
func bindKeyGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KeyGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyGateway *KeyGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyGateway.Contract.KeyGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyGateway *KeyGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyGateway.Contract.KeyGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyGateway *KeyGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyGateway.Contract.KeyGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyGateway *KeyGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyGateway *KeyGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyGateway *KeyGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyGateway.Contract.contract.Transact(opts, method, params...)
}

// ADDTYPEHASH is a free data retrieval call binding the contract method 0xab583c1b.
//
// Solidity: function ADD_TYPEHASH() view returns(bytes32)
func (_KeyGateway *KeyGatewayCaller) ADDTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KeyGateway.contract.Call(opts, &out, "ADD_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADDTYPEHASH is a free data retrieval call binding the contract method 0xab583c1b.
//
// Solidity: function ADD_TYPEHASH() view returns(bytes32)
func (_KeyGateway *KeyGatewaySession) ADDTYPEHASH() ([32]byte, error) {
	return _KeyGateway.Contract.ADDTYPEHASH(&_KeyGateway.CallOpts)
}

// ADDTYPEHASH is a free data retrieval call binding the contract method 0xab583c1b.
//
// Solidity: function ADD_TYPEHASH() view returns(bytes32)
func (_KeyGateway *KeyGatewayCallerSession) ADDTYPEHASH() ([32]byte, error) {
	return _KeyGateway.Contract.ADDTYPEHASH(&_KeyGateway.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_KeyGateway *KeyGatewayCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KeyGateway.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_KeyGateway *KeyGatewaySession) VERSION() (string, error) {
	return _KeyGateway.Contract.VERSION(&_KeyGateway.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_KeyGateway *KeyGatewayCallerSession) VERSION() (string, error) {
	return _KeyGateway.Contract.VERSION(&_KeyGateway.CallOpts)
}

// DomainSeparatorV4 is a free data retrieval call binding the contract method 0x78e890ba.
//
// Solidity: function domainSeparatorV4() view returns(bytes32)
func (_KeyGateway *KeyGatewayCaller) DomainSeparatorV4(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KeyGateway.contract.Call(opts, &out, "domainSeparatorV4")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparatorV4 is a free data retrieval call binding the contract method 0x78e890ba.
//
// Solidity: function domainSeparatorV4() view returns(bytes32)
func (_KeyGateway *KeyGatewaySession) DomainSeparatorV4() ([32]byte, error) {
	return _KeyGateway.Contract.DomainSeparatorV4(&_KeyGateway.CallOpts)
}

// DomainSeparatorV4 is a free data retrieval call binding the contract method 0x78e890ba.
//
// Solidity: function domainSeparatorV4() view returns(bytes32)
func (_KeyGateway *KeyGatewayCallerSession) DomainSeparatorV4() ([32]byte, error) {
	return _KeyGateway.Contract.DomainSeparatorV4(&_KeyGateway.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_KeyGateway *KeyGatewayCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _KeyGateway.contract.Call(opts, &out, "eip712Domain")

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
func (_KeyGateway *KeyGatewaySession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _KeyGateway.Contract.Eip712Domain(&_KeyGateway.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_KeyGateway *KeyGatewayCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _KeyGateway.Contract.Eip712Domain(&_KeyGateway.CallOpts)
}

// Guardians is a free data retrieval call binding the contract method 0x0633b14a.
//
// Solidity: function guardians(address guardian) view returns(bool isGuardian)
func (_KeyGateway *KeyGatewayCaller) Guardians(opts *bind.CallOpts, guardian common.Address) (bool, error) {
	var out []interface{}
	err := _KeyGateway.contract.Call(opts, &out, "guardians", guardian)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Guardians is a free data retrieval call binding the contract method 0x0633b14a.
//
// Solidity: function guardians(address guardian) view returns(bool isGuardian)
func (_KeyGateway *KeyGatewaySession) Guardians(guardian common.Address) (bool, error) {
	return _KeyGateway.Contract.Guardians(&_KeyGateway.CallOpts, guardian)
}

// Guardians is a free data retrieval call binding the contract method 0x0633b14a.
//
// Solidity: function guardians(address guardian) view returns(bool isGuardian)
func (_KeyGateway *KeyGatewayCallerSession) Guardians(guardian common.Address) (bool, error) {
	return _KeyGateway.Contract.Guardians(&_KeyGateway.CallOpts, guardian)
}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x4980f288.
//
// Solidity: function hashTypedDataV4(bytes32 structHash) view returns(bytes32)
func (_KeyGateway *KeyGatewayCaller) HashTypedDataV4(opts *bind.CallOpts, structHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _KeyGateway.contract.Call(opts, &out, "hashTypedDataV4", structHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x4980f288.
//
// Solidity: function hashTypedDataV4(bytes32 structHash) view returns(bytes32)
func (_KeyGateway *KeyGatewaySession) HashTypedDataV4(structHash [32]byte) ([32]byte, error) {
	return _KeyGateway.Contract.HashTypedDataV4(&_KeyGateway.CallOpts, structHash)
}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x4980f288.
//
// Solidity: function hashTypedDataV4(bytes32 structHash) view returns(bytes32)
func (_KeyGateway *KeyGatewayCallerSession) HashTypedDataV4(structHash [32]byte) ([32]byte, error) {
	return _KeyGateway.Contract.HashTypedDataV4(&_KeyGateway.CallOpts, structHash)
}

// KeyRegistry is a free data retrieval call binding the contract method 0x086b5198.
//
// Solidity: function keyRegistry() view returns(address)
func (_KeyGateway *KeyGatewayCaller) KeyRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeyGateway.contract.Call(opts, &out, "keyRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeyRegistry is a free data retrieval call binding the contract method 0x086b5198.
//
// Solidity: function keyRegistry() view returns(address)
func (_KeyGateway *KeyGatewaySession) KeyRegistry() (common.Address, error) {
	return _KeyGateway.Contract.KeyRegistry(&_KeyGateway.CallOpts)
}

// KeyRegistry is a free data retrieval call binding the contract method 0x086b5198.
//
// Solidity: function keyRegistry() view returns(address)
func (_KeyGateway *KeyGatewayCallerSession) KeyRegistry() (common.Address, error) {
	return _KeyGateway.Contract.KeyRegistry(&_KeyGateway.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_KeyGateway *KeyGatewayCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KeyGateway.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_KeyGateway *KeyGatewaySession) Nonces(owner common.Address) (*big.Int, error) {
	return _KeyGateway.Contract.Nonces(&_KeyGateway.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_KeyGateway *KeyGatewayCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _KeyGateway.Contract.Nonces(&_KeyGateway.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KeyGateway *KeyGatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeyGateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KeyGateway *KeyGatewaySession) Owner() (common.Address, error) {
	return _KeyGateway.Contract.Owner(&_KeyGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KeyGateway *KeyGatewayCallerSession) Owner() (common.Address, error) {
	return _KeyGateway.Contract.Owner(&_KeyGateway.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KeyGateway *KeyGatewayCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KeyGateway.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KeyGateway *KeyGatewaySession) Paused() (bool, error) {
	return _KeyGateway.Contract.Paused(&_KeyGateway.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KeyGateway *KeyGatewayCallerSession) Paused() (bool, error) {
	return _KeyGateway.Contract.Paused(&_KeyGateway.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_KeyGateway *KeyGatewayCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeyGateway.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_KeyGateway *KeyGatewaySession) PendingOwner() (common.Address, error) {
	return _KeyGateway.Contract.PendingOwner(&_KeyGateway.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_KeyGateway *KeyGatewayCallerSession) PendingOwner() (common.Address, error) {
	return _KeyGateway.Contract.PendingOwner(&_KeyGateway.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_KeyGateway *KeyGatewayTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyGateway.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_KeyGateway *KeyGatewaySession) AcceptOwnership() (*types.Transaction, error) {
	return _KeyGateway.Contract.AcceptOwnership(&_KeyGateway.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_KeyGateway *KeyGatewayTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _KeyGateway.Contract.AcceptOwnership(&_KeyGateway.TransactOpts)
}

// Add is a paid mutator transaction binding the contract method 0x22b1a414.
//
// Solidity: function add(uint32 keyType, bytes key, uint8 metadataType, bytes metadata) returns()
func (_KeyGateway *KeyGatewayTransactor) Add(opts *bind.TransactOpts, keyType uint32, key []byte, metadataType uint8, metadata []byte) (*types.Transaction, error) {
	return _KeyGateway.contract.Transact(opts, "add", keyType, key, metadataType, metadata)
}

// Add is a paid mutator transaction binding the contract method 0x22b1a414.
//
// Solidity: function add(uint32 keyType, bytes key, uint8 metadataType, bytes metadata) returns()
func (_KeyGateway *KeyGatewaySession) Add(keyType uint32, key []byte, metadataType uint8, metadata []byte) (*types.Transaction, error) {
	return _KeyGateway.Contract.Add(&_KeyGateway.TransactOpts, keyType, key, metadataType, metadata)
}

// Add is a paid mutator transaction binding the contract method 0x22b1a414.
//
// Solidity: function add(uint32 keyType, bytes key, uint8 metadataType, bytes metadata) returns()
func (_KeyGateway *KeyGatewayTransactorSession) Add(keyType uint32, key []byte, metadataType uint8, metadata []byte) (*types.Transaction, error) {
	return _KeyGateway.Contract.Add(&_KeyGateway.TransactOpts, keyType, key, metadataType, metadata)
}

// AddFor is a paid mutator transaction binding the contract method 0xa005d3d2.
//
// Solidity: function addFor(address fidOwner, uint32 keyType, bytes key, uint8 metadataType, bytes metadata, uint256 deadline, bytes sig) returns()
func (_KeyGateway *KeyGatewayTransactor) AddFor(opts *bind.TransactOpts, fidOwner common.Address, keyType uint32, key []byte, metadataType uint8, metadata []byte, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _KeyGateway.contract.Transact(opts, "addFor", fidOwner, keyType, key, metadataType, metadata, deadline, sig)
}

// AddFor is a paid mutator transaction binding the contract method 0xa005d3d2.
//
// Solidity: function addFor(address fidOwner, uint32 keyType, bytes key, uint8 metadataType, bytes metadata, uint256 deadline, bytes sig) returns()
func (_KeyGateway *KeyGatewaySession) AddFor(fidOwner common.Address, keyType uint32, key []byte, metadataType uint8, metadata []byte, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _KeyGateway.Contract.AddFor(&_KeyGateway.TransactOpts, fidOwner, keyType, key, metadataType, metadata, deadline, sig)
}

// AddFor is a paid mutator transaction binding the contract method 0xa005d3d2.
//
// Solidity: function addFor(address fidOwner, uint32 keyType, bytes key, uint8 metadataType, bytes metadata, uint256 deadline, bytes sig) returns()
func (_KeyGateway *KeyGatewayTransactorSession) AddFor(fidOwner common.Address, keyType uint32, key []byte, metadataType uint8, metadata []byte, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _KeyGateway.Contract.AddFor(&_KeyGateway.TransactOpts, fidOwner, keyType, key, metadataType, metadata, deadline, sig)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_KeyGateway *KeyGatewayTransactor) AddGuardian(opts *bind.TransactOpts, guardian common.Address) (*types.Transaction, error) {
	return _KeyGateway.contract.Transact(opts, "addGuardian", guardian)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_KeyGateway *KeyGatewaySession) AddGuardian(guardian common.Address) (*types.Transaction, error) {
	return _KeyGateway.Contract.AddGuardian(&_KeyGateway.TransactOpts, guardian)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_KeyGateway *KeyGatewayTransactorSession) AddGuardian(guardian common.Address) (*types.Transaction, error) {
	return _KeyGateway.Contract.AddGuardian(&_KeyGateway.TransactOpts, guardian)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KeyGateway *KeyGatewayTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyGateway.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KeyGateway *KeyGatewaySession) Pause() (*types.Transaction, error) {
	return _KeyGateway.Contract.Pause(&_KeyGateway.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KeyGateway *KeyGatewayTransactorSession) Pause() (*types.Transaction, error) {
	return _KeyGateway.Contract.Pause(&_KeyGateway.TransactOpts)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_KeyGateway *KeyGatewayTransactor) RemoveGuardian(opts *bind.TransactOpts, guardian common.Address) (*types.Transaction, error) {
	return _KeyGateway.contract.Transact(opts, "removeGuardian", guardian)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_KeyGateway *KeyGatewaySession) RemoveGuardian(guardian common.Address) (*types.Transaction, error) {
	return _KeyGateway.Contract.RemoveGuardian(&_KeyGateway.TransactOpts, guardian)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_KeyGateway *KeyGatewayTransactorSession) RemoveGuardian(guardian common.Address) (*types.Transaction, error) {
	return _KeyGateway.Contract.RemoveGuardian(&_KeyGateway.TransactOpts, guardian)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KeyGateway *KeyGatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyGateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KeyGateway *KeyGatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _KeyGateway.Contract.RenounceOwnership(&_KeyGateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KeyGateway *KeyGatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _KeyGateway.Contract.RenounceOwnership(&_KeyGateway.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KeyGateway *KeyGatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KeyGateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KeyGateway *KeyGatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KeyGateway.Contract.TransferOwnership(&_KeyGateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KeyGateway *KeyGatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KeyGateway.Contract.TransferOwnership(&_KeyGateway.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KeyGateway *KeyGatewayTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyGateway.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KeyGateway *KeyGatewaySession) Unpause() (*types.Transaction, error) {
	return _KeyGateway.Contract.Unpause(&_KeyGateway.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KeyGateway *KeyGatewayTransactorSession) Unpause() (*types.Transaction, error) {
	return _KeyGateway.Contract.Unpause(&_KeyGateway.TransactOpts)
}

// UseNonce is a paid mutator transaction binding the contract method 0x69615a4c.
//
// Solidity: function useNonce() returns(uint256)
func (_KeyGateway *KeyGatewayTransactor) UseNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyGateway.contract.Transact(opts, "useNonce")
}

// UseNonce is a paid mutator transaction binding the contract method 0x69615a4c.
//
// Solidity: function useNonce() returns(uint256)
func (_KeyGateway *KeyGatewaySession) UseNonce() (*types.Transaction, error) {
	return _KeyGateway.Contract.UseNonce(&_KeyGateway.TransactOpts)
}

// UseNonce is a paid mutator transaction binding the contract method 0x69615a4c.
//
// Solidity: function useNonce() returns(uint256)
func (_KeyGateway *KeyGatewayTransactorSession) UseNonce() (*types.Transaction, error) {
	return _KeyGateway.Contract.UseNonce(&_KeyGateway.TransactOpts)
}

// KeyGatewayAddIterator is returned from FilterAdd and is used to iterate over the raw logs and unpacked data for Add events raised by the KeyGateway contract.
type KeyGatewayAddIterator struct {
	Event *KeyGatewayAdd // Event containing the contract specifics and raw log

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
func (it *KeyGatewayAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyGatewayAdd)
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
		it.Event = new(KeyGatewayAdd)
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
func (it *KeyGatewayAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyGatewayAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyGatewayAdd represents a Add event raised by the KeyGateway contract.
type KeyGatewayAdd struct {
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdd is a free log retrieval operation binding the contract event 0x87dc5eecd6d6bdeae407c426da6bfba5b7190befc554ed5d4d62dd5cf939fbae.
//
// Solidity: event Add(address indexed guardian)
func (_KeyGateway *KeyGatewayFilterer) FilterAdd(opts *bind.FilterOpts, guardian []common.Address) (*KeyGatewayAddIterator, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _KeyGateway.contract.FilterLogs(opts, "Add", guardianRule)
	if err != nil {
		return nil, err
	}
	return &KeyGatewayAddIterator{contract: _KeyGateway.contract, event: "Add", logs: logs, sub: sub}, nil
}

// WatchAdd is a free log subscription operation binding the contract event 0x87dc5eecd6d6bdeae407c426da6bfba5b7190befc554ed5d4d62dd5cf939fbae.
//
// Solidity: event Add(address indexed guardian)
func (_KeyGateway *KeyGatewayFilterer) WatchAdd(opts *bind.WatchOpts, sink chan<- *KeyGatewayAdd, guardian []common.Address) (event.Subscription, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _KeyGateway.contract.WatchLogs(opts, "Add", guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyGatewayAdd)
				if err := _KeyGateway.contract.UnpackLog(event, "Add", log); err != nil {
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
func (_KeyGateway *KeyGatewayFilterer) ParseAdd(log types.Log) (*KeyGatewayAdd, error) {
	event := new(KeyGatewayAdd)
	if err := _KeyGateway.contract.UnpackLog(event, "Add", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyGatewayEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the KeyGateway contract.
type KeyGatewayEIP712DomainChangedIterator struct {
	Event *KeyGatewayEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *KeyGatewayEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyGatewayEIP712DomainChanged)
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
		it.Event = new(KeyGatewayEIP712DomainChanged)
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
func (it *KeyGatewayEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyGatewayEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyGatewayEIP712DomainChanged represents a EIP712DomainChanged event raised by the KeyGateway contract.
type KeyGatewayEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_KeyGateway *KeyGatewayFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*KeyGatewayEIP712DomainChangedIterator, error) {

	logs, sub, err := _KeyGateway.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &KeyGatewayEIP712DomainChangedIterator{contract: _KeyGateway.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_KeyGateway *KeyGatewayFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *KeyGatewayEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _KeyGateway.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyGatewayEIP712DomainChanged)
				if err := _KeyGateway.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_KeyGateway *KeyGatewayFilterer) ParseEIP712DomainChanged(log types.Log) (*KeyGatewayEIP712DomainChanged, error) {
	event := new(KeyGatewayEIP712DomainChanged)
	if err := _KeyGateway.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyGatewayOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the KeyGateway contract.
type KeyGatewayOwnershipTransferStartedIterator struct {
	Event *KeyGatewayOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *KeyGatewayOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyGatewayOwnershipTransferStarted)
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
		it.Event = new(KeyGatewayOwnershipTransferStarted)
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
func (it *KeyGatewayOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyGatewayOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyGatewayOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the KeyGateway contract.
type KeyGatewayOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_KeyGateway *KeyGatewayFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KeyGatewayOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KeyGateway.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KeyGatewayOwnershipTransferStartedIterator{contract: _KeyGateway.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_KeyGateway *KeyGatewayFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *KeyGatewayOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KeyGateway.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyGatewayOwnershipTransferStarted)
				if err := _KeyGateway.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_KeyGateway *KeyGatewayFilterer) ParseOwnershipTransferStarted(log types.Log) (*KeyGatewayOwnershipTransferStarted, error) {
	event := new(KeyGatewayOwnershipTransferStarted)
	if err := _KeyGateway.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyGatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the KeyGateway contract.
type KeyGatewayOwnershipTransferredIterator struct {
	Event *KeyGatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *KeyGatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyGatewayOwnershipTransferred)
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
		it.Event = new(KeyGatewayOwnershipTransferred)
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
func (it *KeyGatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyGatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyGatewayOwnershipTransferred represents a OwnershipTransferred event raised by the KeyGateway contract.
type KeyGatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KeyGateway *KeyGatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KeyGatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KeyGateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KeyGatewayOwnershipTransferredIterator{contract: _KeyGateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KeyGateway *KeyGatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KeyGatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KeyGateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyGatewayOwnershipTransferred)
				if err := _KeyGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_KeyGateway *KeyGatewayFilterer) ParseOwnershipTransferred(log types.Log) (*KeyGatewayOwnershipTransferred, error) {
	event := new(KeyGatewayOwnershipTransferred)
	if err := _KeyGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyGatewayPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the KeyGateway contract.
type KeyGatewayPausedIterator struct {
	Event *KeyGatewayPaused // Event containing the contract specifics and raw log

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
func (it *KeyGatewayPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyGatewayPaused)
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
		it.Event = new(KeyGatewayPaused)
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
func (it *KeyGatewayPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyGatewayPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyGatewayPaused represents a Paused event raised by the KeyGateway contract.
type KeyGatewayPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_KeyGateway *KeyGatewayFilterer) FilterPaused(opts *bind.FilterOpts) (*KeyGatewayPausedIterator, error) {

	logs, sub, err := _KeyGateway.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &KeyGatewayPausedIterator{contract: _KeyGateway.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_KeyGateway *KeyGatewayFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *KeyGatewayPaused) (event.Subscription, error) {

	logs, sub, err := _KeyGateway.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyGatewayPaused)
				if err := _KeyGateway.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_KeyGateway *KeyGatewayFilterer) ParsePaused(log types.Log) (*KeyGatewayPaused, error) {
	event := new(KeyGatewayPaused)
	if err := _KeyGateway.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyGatewayRemoveIterator is returned from FilterRemove and is used to iterate over the raw logs and unpacked data for Remove events raised by the KeyGateway contract.
type KeyGatewayRemoveIterator struct {
	Event *KeyGatewayRemove // Event containing the contract specifics and raw log

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
func (it *KeyGatewayRemoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyGatewayRemove)
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
		it.Event = new(KeyGatewayRemove)
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
func (it *KeyGatewayRemoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyGatewayRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyGatewayRemove represents a Remove event raised by the KeyGateway contract.
type KeyGatewayRemove struct {
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemove is a free log retrieval operation binding the contract event 0xbe7c7ac3248df4581c206a84aab3cb4e7d521b5398b42b681757f78a5a7d411e.
//
// Solidity: event Remove(address indexed guardian)
func (_KeyGateway *KeyGatewayFilterer) FilterRemove(opts *bind.FilterOpts, guardian []common.Address) (*KeyGatewayRemoveIterator, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _KeyGateway.contract.FilterLogs(opts, "Remove", guardianRule)
	if err != nil {
		return nil, err
	}
	return &KeyGatewayRemoveIterator{contract: _KeyGateway.contract, event: "Remove", logs: logs, sub: sub}, nil
}

// WatchRemove is a free log subscription operation binding the contract event 0xbe7c7ac3248df4581c206a84aab3cb4e7d521b5398b42b681757f78a5a7d411e.
//
// Solidity: event Remove(address indexed guardian)
func (_KeyGateway *KeyGatewayFilterer) WatchRemove(opts *bind.WatchOpts, sink chan<- *KeyGatewayRemove, guardian []common.Address) (event.Subscription, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _KeyGateway.contract.WatchLogs(opts, "Remove", guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyGatewayRemove)
				if err := _KeyGateway.contract.UnpackLog(event, "Remove", log); err != nil {
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
func (_KeyGateway *KeyGatewayFilterer) ParseRemove(log types.Log) (*KeyGatewayRemove, error) {
	event := new(KeyGatewayRemove)
	if err := _KeyGateway.contract.UnpackLog(event, "Remove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyGatewayUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the KeyGateway contract.
type KeyGatewayUnpausedIterator struct {
	Event *KeyGatewayUnpaused // Event containing the contract specifics and raw log

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
func (it *KeyGatewayUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyGatewayUnpaused)
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
		it.Event = new(KeyGatewayUnpaused)
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
func (it *KeyGatewayUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyGatewayUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyGatewayUnpaused represents a Unpaused event raised by the KeyGateway contract.
type KeyGatewayUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_KeyGateway *KeyGatewayFilterer) FilterUnpaused(opts *bind.FilterOpts) (*KeyGatewayUnpausedIterator, error) {

	logs, sub, err := _KeyGateway.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &KeyGatewayUnpausedIterator{contract: _KeyGateway.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_KeyGateway *KeyGatewayFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *KeyGatewayUnpaused) (event.Subscription, error) {

	logs, sub, err := _KeyGateway.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyGatewayUnpaused)
				if err := _KeyGateway.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_KeyGateway *KeyGatewayFilterer) ParseUnpaused(log types.Log) (*KeyGatewayUnpaused, error) {
	event := new(KeyGatewayUnpaused)
	if err := _KeyGateway.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
