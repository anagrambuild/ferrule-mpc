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

// BundlerMetaData contains all meta data concerning the Bundler contract.
var BundlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_keyRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGuardian\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"Add\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"Remove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADD_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"keyType\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"metadataType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fidOwner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"keyType\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"metadataType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"addFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"addGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparatorV4\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"guardians\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isGuardian\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"structHash\",\"type\":\"bytes32\"}],\"name\":\"hashTypedDataV4\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keyRegistry\",\"outputs\":[{\"internalType\":\"contractIKeyRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"removeGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BundlerABI is the input ABI used to generate the binding from.
// Deprecated: Use BundlerMetaData.ABI instead.
var BundlerABI = BundlerMetaData.ABI

// Bundler is an auto generated Go binding around an Ethereum contract.
type Bundler struct {
	BundlerCaller     // Read-only binding to the contract
	BundlerTransactor // Write-only binding to the contract
	BundlerFilterer   // Log filterer for contract events
}

// BundlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BundlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BundlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BundlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BundlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BundlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BundlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BundlerSession struct {
	Contract     *Bundler          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BundlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BundlerCallerSession struct {
	Contract *BundlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BundlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BundlerTransactorSession struct {
	Contract     *BundlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BundlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BundlerRaw struct {
	Contract *Bundler // Generic contract binding to access the raw methods on
}

// BundlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BundlerCallerRaw struct {
	Contract *BundlerCaller // Generic read-only contract binding to access the raw methods on
}

// BundlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BundlerTransactorRaw struct {
	Contract *BundlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBundler creates a new instance of Bundler, bound to a specific deployed contract.
func NewBundler(address common.Address, backend bind.ContractBackend) (*Bundler, error) {
	contract, err := bindBundler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bundler{BundlerCaller: BundlerCaller{contract: contract}, BundlerTransactor: BundlerTransactor{contract: contract}, BundlerFilterer: BundlerFilterer{contract: contract}}, nil
}

// NewBundlerCaller creates a new read-only instance of Bundler, bound to a specific deployed contract.
func NewBundlerCaller(address common.Address, caller bind.ContractCaller) (*BundlerCaller, error) {
	contract, err := bindBundler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BundlerCaller{contract: contract}, nil
}

// NewBundlerTransactor creates a new write-only instance of Bundler, bound to a specific deployed contract.
func NewBundlerTransactor(address common.Address, transactor bind.ContractTransactor) (*BundlerTransactor, error) {
	contract, err := bindBundler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BundlerTransactor{contract: contract}, nil
}

// NewBundlerFilterer creates a new log filterer instance of Bundler, bound to a specific deployed contract.
func NewBundlerFilterer(address common.Address, filterer bind.ContractFilterer) (*BundlerFilterer, error) {
	contract, err := bindBundler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BundlerFilterer{contract: contract}, nil
}

// bindBundler binds a generic wrapper to an already deployed contract.
func bindBundler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BundlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bundler *BundlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bundler.Contract.BundlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bundler *BundlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bundler.Contract.BundlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bundler *BundlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bundler.Contract.BundlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bundler *BundlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bundler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bundler *BundlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bundler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bundler *BundlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bundler.Contract.contract.Transact(opts, method, params...)
}

// ADDTYPEHASH is a free data retrieval call binding the contract method 0xab583c1b.
//
// Solidity: function ADD_TYPEHASH() view returns(bytes32)
func (_Bundler *BundlerCaller) ADDTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bundler.contract.Call(opts, &out, "ADD_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADDTYPEHASH is a free data retrieval call binding the contract method 0xab583c1b.
//
// Solidity: function ADD_TYPEHASH() view returns(bytes32)
func (_Bundler *BundlerSession) ADDTYPEHASH() ([32]byte, error) {
	return _Bundler.Contract.ADDTYPEHASH(&_Bundler.CallOpts)
}

// ADDTYPEHASH is a free data retrieval call binding the contract method 0xab583c1b.
//
// Solidity: function ADD_TYPEHASH() view returns(bytes32)
func (_Bundler *BundlerCallerSession) ADDTYPEHASH() ([32]byte, error) {
	return _Bundler.Contract.ADDTYPEHASH(&_Bundler.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Bundler *BundlerCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bundler.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Bundler *BundlerSession) VERSION() (string, error) {
	return _Bundler.Contract.VERSION(&_Bundler.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Bundler *BundlerCallerSession) VERSION() (string, error) {
	return _Bundler.Contract.VERSION(&_Bundler.CallOpts)
}

// DomainSeparatorV4 is a free data retrieval call binding the contract method 0x78e890ba.
//
// Solidity: function domainSeparatorV4() view returns(bytes32)
func (_Bundler *BundlerCaller) DomainSeparatorV4(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bundler.contract.Call(opts, &out, "domainSeparatorV4")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparatorV4 is a free data retrieval call binding the contract method 0x78e890ba.
//
// Solidity: function domainSeparatorV4() view returns(bytes32)
func (_Bundler *BundlerSession) DomainSeparatorV4() ([32]byte, error) {
	return _Bundler.Contract.DomainSeparatorV4(&_Bundler.CallOpts)
}

// DomainSeparatorV4 is a free data retrieval call binding the contract method 0x78e890ba.
//
// Solidity: function domainSeparatorV4() view returns(bytes32)
func (_Bundler *BundlerCallerSession) DomainSeparatorV4() ([32]byte, error) {
	return _Bundler.Contract.DomainSeparatorV4(&_Bundler.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Bundler *BundlerCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Bundler.contract.Call(opts, &out, "eip712Domain")

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
func (_Bundler *BundlerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Bundler.Contract.Eip712Domain(&_Bundler.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Bundler *BundlerCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Bundler.Contract.Eip712Domain(&_Bundler.CallOpts)
}

// Guardians is a free data retrieval call binding the contract method 0x0633b14a.
//
// Solidity: function guardians(address guardian) view returns(bool isGuardian)
func (_Bundler *BundlerCaller) Guardians(opts *bind.CallOpts, guardian common.Address) (bool, error) {
	var out []interface{}
	err := _Bundler.contract.Call(opts, &out, "guardians", guardian)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Guardians is a free data retrieval call binding the contract method 0x0633b14a.
//
// Solidity: function guardians(address guardian) view returns(bool isGuardian)
func (_Bundler *BundlerSession) Guardians(guardian common.Address) (bool, error) {
	return _Bundler.Contract.Guardians(&_Bundler.CallOpts, guardian)
}

// Guardians is a free data retrieval call binding the contract method 0x0633b14a.
//
// Solidity: function guardians(address guardian) view returns(bool isGuardian)
func (_Bundler *BundlerCallerSession) Guardians(guardian common.Address) (bool, error) {
	return _Bundler.Contract.Guardians(&_Bundler.CallOpts, guardian)
}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x4980f288.
//
// Solidity: function hashTypedDataV4(bytes32 structHash) view returns(bytes32)
func (_Bundler *BundlerCaller) HashTypedDataV4(opts *bind.CallOpts, structHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Bundler.contract.Call(opts, &out, "hashTypedDataV4", structHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x4980f288.
//
// Solidity: function hashTypedDataV4(bytes32 structHash) view returns(bytes32)
func (_Bundler *BundlerSession) HashTypedDataV4(structHash [32]byte) ([32]byte, error) {
	return _Bundler.Contract.HashTypedDataV4(&_Bundler.CallOpts, structHash)
}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x4980f288.
//
// Solidity: function hashTypedDataV4(bytes32 structHash) view returns(bytes32)
func (_Bundler *BundlerCallerSession) HashTypedDataV4(structHash [32]byte) ([32]byte, error) {
	return _Bundler.Contract.HashTypedDataV4(&_Bundler.CallOpts, structHash)
}

// KeyRegistry is a free data retrieval call binding the contract method 0x086b5198.
//
// Solidity: function keyRegistry() view returns(address)
func (_Bundler *BundlerCaller) KeyRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bundler.contract.Call(opts, &out, "keyRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeyRegistry is a free data retrieval call binding the contract method 0x086b5198.
//
// Solidity: function keyRegistry() view returns(address)
func (_Bundler *BundlerSession) KeyRegistry() (common.Address, error) {
	return _Bundler.Contract.KeyRegistry(&_Bundler.CallOpts)
}

// KeyRegistry is a free data retrieval call binding the contract method 0x086b5198.
//
// Solidity: function keyRegistry() view returns(address)
func (_Bundler *BundlerCallerSession) KeyRegistry() (common.Address, error) {
	return _Bundler.Contract.KeyRegistry(&_Bundler.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Bundler *BundlerCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bundler.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Bundler *BundlerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Bundler.Contract.Nonces(&_Bundler.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Bundler *BundlerCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Bundler.Contract.Nonces(&_Bundler.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bundler *BundlerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bundler.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bundler *BundlerSession) Owner() (common.Address, error) {
	return _Bundler.Contract.Owner(&_Bundler.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bundler *BundlerCallerSession) Owner() (common.Address, error) {
	return _Bundler.Contract.Owner(&_Bundler.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bundler *BundlerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bundler.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bundler *BundlerSession) Paused() (bool, error) {
	return _Bundler.Contract.Paused(&_Bundler.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bundler *BundlerCallerSession) Paused() (bool, error) {
	return _Bundler.Contract.Paused(&_Bundler.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Bundler *BundlerCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bundler.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Bundler *BundlerSession) PendingOwner() (common.Address, error) {
	return _Bundler.Contract.PendingOwner(&_Bundler.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Bundler *BundlerCallerSession) PendingOwner() (common.Address, error) {
	return _Bundler.Contract.PendingOwner(&_Bundler.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Bundler *BundlerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bundler.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Bundler *BundlerSession) AcceptOwnership() (*types.Transaction, error) {
	return _Bundler.Contract.AcceptOwnership(&_Bundler.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Bundler *BundlerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Bundler.Contract.AcceptOwnership(&_Bundler.TransactOpts)
}

// Add is a paid mutator transaction binding the contract method 0x22b1a414.
//
// Solidity: function add(uint32 keyType, bytes key, uint8 metadataType, bytes metadata) returns()
func (_Bundler *BundlerTransactor) Add(opts *bind.TransactOpts, keyType uint32, key []byte, metadataType uint8, metadata []byte) (*types.Transaction, error) {
	return _Bundler.contract.Transact(opts, "add", keyType, key, metadataType, metadata)
}

// Add is a paid mutator transaction binding the contract method 0x22b1a414.
//
// Solidity: function add(uint32 keyType, bytes key, uint8 metadataType, bytes metadata) returns()
func (_Bundler *BundlerSession) Add(keyType uint32, key []byte, metadataType uint8, metadata []byte) (*types.Transaction, error) {
	return _Bundler.Contract.Add(&_Bundler.TransactOpts, keyType, key, metadataType, metadata)
}

// Add is a paid mutator transaction binding the contract method 0x22b1a414.
//
// Solidity: function add(uint32 keyType, bytes key, uint8 metadataType, bytes metadata) returns()
func (_Bundler *BundlerTransactorSession) Add(keyType uint32, key []byte, metadataType uint8, metadata []byte) (*types.Transaction, error) {
	return _Bundler.Contract.Add(&_Bundler.TransactOpts, keyType, key, metadataType, metadata)
}

// AddFor is a paid mutator transaction binding the contract method 0xa005d3d2.
//
// Solidity: function addFor(address fidOwner, uint32 keyType, bytes key, uint8 metadataType, bytes metadata, uint256 deadline, bytes sig) returns()
func (_Bundler *BundlerTransactor) AddFor(opts *bind.TransactOpts, fidOwner common.Address, keyType uint32, key []byte, metadataType uint8, metadata []byte, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _Bundler.contract.Transact(opts, "addFor", fidOwner, keyType, key, metadataType, metadata, deadline, sig)
}

// AddFor is a paid mutator transaction binding the contract method 0xa005d3d2.
//
// Solidity: function addFor(address fidOwner, uint32 keyType, bytes key, uint8 metadataType, bytes metadata, uint256 deadline, bytes sig) returns()
func (_Bundler *BundlerSession) AddFor(fidOwner common.Address, keyType uint32, key []byte, metadataType uint8, metadata []byte, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _Bundler.Contract.AddFor(&_Bundler.TransactOpts, fidOwner, keyType, key, metadataType, metadata, deadline, sig)
}

// AddFor is a paid mutator transaction binding the contract method 0xa005d3d2.
//
// Solidity: function addFor(address fidOwner, uint32 keyType, bytes key, uint8 metadataType, bytes metadata, uint256 deadline, bytes sig) returns()
func (_Bundler *BundlerTransactorSession) AddFor(fidOwner common.Address, keyType uint32, key []byte, metadataType uint8, metadata []byte, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _Bundler.Contract.AddFor(&_Bundler.TransactOpts, fidOwner, keyType, key, metadataType, metadata, deadline, sig)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_Bundler *BundlerTransactor) AddGuardian(opts *bind.TransactOpts, guardian common.Address) (*types.Transaction, error) {
	return _Bundler.contract.Transact(opts, "addGuardian", guardian)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_Bundler *BundlerSession) AddGuardian(guardian common.Address) (*types.Transaction, error) {
	return _Bundler.Contract.AddGuardian(&_Bundler.TransactOpts, guardian)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_Bundler *BundlerTransactorSession) AddGuardian(guardian common.Address) (*types.Transaction, error) {
	return _Bundler.Contract.AddGuardian(&_Bundler.TransactOpts, guardian)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bundler *BundlerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bundler.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bundler *BundlerSession) Pause() (*types.Transaction, error) {
	return _Bundler.Contract.Pause(&_Bundler.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bundler *BundlerTransactorSession) Pause() (*types.Transaction, error) {
	return _Bundler.Contract.Pause(&_Bundler.TransactOpts)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_Bundler *BundlerTransactor) RemoveGuardian(opts *bind.TransactOpts, guardian common.Address) (*types.Transaction, error) {
	return _Bundler.contract.Transact(opts, "removeGuardian", guardian)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_Bundler *BundlerSession) RemoveGuardian(guardian common.Address) (*types.Transaction, error) {
	return _Bundler.Contract.RemoveGuardian(&_Bundler.TransactOpts, guardian)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_Bundler *BundlerTransactorSession) RemoveGuardian(guardian common.Address) (*types.Transaction, error) {
	return _Bundler.Contract.RemoveGuardian(&_Bundler.TransactOpts, guardian)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bundler *BundlerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bundler.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bundler *BundlerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bundler.Contract.RenounceOwnership(&_Bundler.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bundler *BundlerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bundler.Contract.RenounceOwnership(&_Bundler.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bundler *BundlerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bundler.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bundler *BundlerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bundler.Contract.TransferOwnership(&_Bundler.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bundler *BundlerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bundler.Contract.TransferOwnership(&_Bundler.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bundler *BundlerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bundler.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bundler *BundlerSession) Unpause() (*types.Transaction, error) {
	return _Bundler.Contract.Unpause(&_Bundler.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bundler *BundlerTransactorSession) Unpause() (*types.Transaction, error) {
	return _Bundler.Contract.Unpause(&_Bundler.TransactOpts)
}

// UseNonce is a paid mutator transaction binding the contract method 0x69615a4c.
//
// Solidity: function useNonce() returns(uint256)
func (_Bundler *BundlerTransactor) UseNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bundler.contract.Transact(opts, "useNonce")
}

// UseNonce is a paid mutator transaction binding the contract method 0x69615a4c.
//
// Solidity: function useNonce() returns(uint256)
func (_Bundler *BundlerSession) UseNonce() (*types.Transaction, error) {
	return _Bundler.Contract.UseNonce(&_Bundler.TransactOpts)
}

// UseNonce is a paid mutator transaction binding the contract method 0x69615a4c.
//
// Solidity: function useNonce() returns(uint256)
func (_Bundler *BundlerTransactorSession) UseNonce() (*types.Transaction, error) {
	return _Bundler.Contract.UseNonce(&_Bundler.TransactOpts)
}

// BundlerAddIterator is returned from FilterAdd and is used to iterate over the raw logs and unpacked data for Add events raised by the Bundler contract.
type BundlerAddIterator struct {
	Event *BundlerAdd // Event containing the contract specifics and raw log

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
func (it *BundlerAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BundlerAdd)
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
		it.Event = new(BundlerAdd)
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
func (it *BundlerAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BundlerAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BundlerAdd represents a Add event raised by the Bundler contract.
type BundlerAdd struct {
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdd is a free log retrieval operation binding the contract event 0x87dc5eecd6d6bdeae407c426da6bfba5b7190befc554ed5d4d62dd5cf939fbae.
//
// Solidity: event Add(address indexed guardian)
func (_Bundler *BundlerFilterer) FilterAdd(opts *bind.FilterOpts, guardian []common.Address) (*BundlerAddIterator, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _Bundler.contract.FilterLogs(opts, "Add", guardianRule)
	if err != nil {
		return nil, err
	}
	return &BundlerAddIterator{contract: _Bundler.contract, event: "Add", logs: logs, sub: sub}, nil
}

// WatchAdd is a free log subscription operation binding the contract event 0x87dc5eecd6d6bdeae407c426da6bfba5b7190befc554ed5d4d62dd5cf939fbae.
//
// Solidity: event Add(address indexed guardian)
func (_Bundler *BundlerFilterer) WatchAdd(opts *bind.WatchOpts, sink chan<- *BundlerAdd, guardian []common.Address) (event.Subscription, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _Bundler.contract.WatchLogs(opts, "Add", guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BundlerAdd)
				if err := _Bundler.contract.UnpackLog(event, "Add", log); err != nil {
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
func (_Bundler *BundlerFilterer) ParseAdd(log types.Log) (*BundlerAdd, error) {
	event := new(BundlerAdd)
	if err := _Bundler.contract.UnpackLog(event, "Add", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BundlerEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Bundler contract.
type BundlerEIP712DomainChangedIterator struct {
	Event *BundlerEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *BundlerEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BundlerEIP712DomainChanged)
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
		it.Event = new(BundlerEIP712DomainChanged)
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
func (it *BundlerEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BundlerEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BundlerEIP712DomainChanged represents a EIP712DomainChanged event raised by the Bundler contract.
type BundlerEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Bundler *BundlerFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*BundlerEIP712DomainChangedIterator, error) {

	logs, sub, err := _Bundler.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &BundlerEIP712DomainChangedIterator{contract: _Bundler.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Bundler *BundlerFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *BundlerEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Bundler.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BundlerEIP712DomainChanged)
				if err := _Bundler.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_Bundler *BundlerFilterer) ParseEIP712DomainChanged(log types.Log) (*BundlerEIP712DomainChanged, error) {
	event := new(BundlerEIP712DomainChanged)
	if err := _Bundler.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BundlerOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Bundler contract.
type BundlerOwnershipTransferStartedIterator struct {
	Event *BundlerOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *BundlerOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BundlerOwnershipTransferStarted)
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
		it.Event = new(BundlerOwnershipTransferStarted)
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
func (it *BundlerOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BundlerOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BundlerOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Bundler contract.
type BundlerOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Bundler *BundlerFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BundlerOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bundler.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BundlerOwnershipTransferStartedIterator{contract: _Bundler.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Bundler *BundlerFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *BundlerOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bundler.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BundlerOwnershipTransferStarted)
				if err := _Bundler.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_Bundler *BundlerFilterer) ParseOwnershipTransferStarted(log types.Log) (*BundlerOwnershipTransferStarted, error) {
	event := new(BundlerOwnershipTransferStarted)
	if err := _Bundler.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BundlerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bundler contract.
type BundlerOwnershipTransferredIterator struct {
	Event *BundlerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BundlerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BundlerOwnershipTransferred)
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
		it.Event = new(BundlerOwnershipTransferred)
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
func (it *BundlerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BundlerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BundlerOwnershipTransferred represents a OwnershipTransferred event raised by the Bundler contract.
type BundlerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bundler *BundlerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BundlerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bundler.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BundlerOwnershipTransferredIterator{contract: _Bundler.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bundler *BundlerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BundlerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bundler.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BundlerOwnershipTransferred)
				if err := _Bundler.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bundler *BundlerFilterer) ParseOwnershipTransferred(log types.Log) (*BundlerOwnershipTransferred, error) {
	event := new(BundlerOwnershipTransferred)
	if err := _Bundler.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BundlerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Bundler contract.
type BundlerPausedIterator struct {
	Event *BundlerPaused // Event containing the contract specifics and raw log

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
func (it *BundlerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BundlerPaused)
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
		it.Event = new(BundlerPaused)
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
func (it *BundlerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BundlerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BundlerPaused represents a Paused event raised by the Bundler contract.
type BundlerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bundler *BundlerFilterer) FilterPaused(opts *bind.FilterOpts) (*BundlerPausedIterator, error) {

	logs, sub, err := _Bundler.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BundlerPausedIterator{contract: _Bundler.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bundler *BundlerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BundlerPaused) (event.Subscription, error) {

	logs, sub, err := _Bundler.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BundlerPaused)
				if err := _Bundler.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Bundler *BundlerFilterer) ParsePaused(log types.Log) (*BundlerPaused, error) {
	event := new(BundlerPaused)
	if err := _Bundler.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BundlerRemoveIterator is returned from FilterRemove and is used to iterate over the raw logs and unpacked data for Remove events raised by the Bundler contract.
type BundlerRemoveIterator struct {
	Event *BundlerRemove // Event containing the contract specifics and raw log

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
func (it *BundlerRemoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BundlerRemove)
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
		it.Event = new(BundlerRemove)
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
func (it *BundlerRemoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BundlerRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BundlerRemove represents a Remove event raised by the Bundler contract.
type BundlerRemove struct {
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemove is a free log retrieval operation binding the contract event 0xbe7c7ac3248df4581c206a84aab3cb4e7d521b5398b42b681757f78a5a7d411e.
//
// Solidity: event Remove(address indexed guardian)
func (_Bundler *BundlerFilterer) FilterRemove(opts *bind.FilterOpts, guardian []common.Address) (*BundlerRemoveIterator, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _Bundler.contract.FilterLogs(opts, "Remove", guardianRule)
	if err != nil {
		return nil, err
	}
	return &BundlerRemoveIterator{contract: _Bundler.contract, event: "Remove", logs: logs, sub: sub}, nil
}

// WatchRemove is a free log subscription operation binding the contract event 0xbe7c7ac3248df4581c206a84aab3cb4e7d521b5398b42b681757f78a5a7d411e.
//
// Solidity: event Remove(address indexed guardian)
func (_Bundler *BundlerFilterer) WatchRemove(opts *bind.WatchOpts, sink chan<- *BundlerRemove, guardian []common.Address) (event.Subscription, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _Bundler.contract.WatchLogs(opts, "Remove", guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BundlerRemove)
				if err := _Bundler.contract.UnpackLog(event, "Remove", log); err != nil {
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
func (_Bundler *BundlerFilterer) ParseRemove(log types.Log) (*BundlerRemove, error) {
	event := new(BundlerRemove)
	if err := _Bundler.contract.UnpackLog(event, "Remove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BundlerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Bundler contract.
type BundlerUnpausedIterator struct {
	Event *BundlerUnpaused // Event containing the contract specifics and raw log

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
func (it *BundlerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BundlerUnpaused)
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
		it.Event = new(BundlerUnpaused)
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
func (it *BundlerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BundlerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BundlerUnpaused represents a Unpaused event raised by the Bundler contract.
type BundlerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bundler *BundlerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BundlerUnpausedIterator, error) {

	logs, sub, err := _Bundler.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BundlerUnpausedIterator{contract: _Bundler.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bundler *BundlerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BundlerUnpaused) (event.Subscription, error) {

	logs, sub, err := _Bundler.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BundlerUnpaused)
				if err := _Bundler.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Bundler *BundlerFilterer) ParseUnpaused(log types.Log) (*BundlerUnpaused, error) {
	event := new(BundlerUnpaused)
	if err := _Bundler.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
