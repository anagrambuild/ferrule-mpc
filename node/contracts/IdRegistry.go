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

// IIdRegistryBulkRegisterData is an auto generated low-level Go binding around an user-defined struct.
type IIdRegistryBulkRegisterData struct {
	Fid      *big.Int
	Custody  common.Address
	Recovery common.Address
}

// IIdRegistryBulkRegisterDefaultRecoveryData is an auto generated low-level Go binding around an user-defined struct.
type IIdRegistryBulkRegisterDefaultRecoveryData struct {
	Fid     *big.Int
	Custody common.Address
}

// IdRegistryMetaData contains all meta data concerning the IdRegistry contract.
var IdRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_migrator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyMigrated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GatewayFrozen\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HasId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HasNoId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGuardian\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyMigrator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissionRevoked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"Add\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"}],\"name\":\"AdminReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"}],\"name\":\"ChangeRecoveryAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"idGateway\",\"type\":\"address\"}],\"name\":\"FreezeIdGateway\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"migratedAt\",\"type\":\"uint256\"}],\"name\":\"Migrated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Recover\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"}],\"name\":\"Register\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"Remove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCounter\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCounter\",\"type\":\"uint256\"}],\"name\":\"SetIdCounter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldIdGateway\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newIdGateway\",\"type\":\"address\"}],\"name\":\"SetIdGateway\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldMigrator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newMigrator\",\"type\":\"address\"}],\"name\":\"SetMigrator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHANGE_RECOVERY_ADDRESS_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_AND_CHANGE_RECOVERY_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"addGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint24\",\"name\":\"fid\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"custody\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"}],\"internalType\":\"structIIdRegistry.BulkRegisterData[]\",\"name\":\"ids\",\"type\":\"tuple[]\"}],\"name\":\"bulkRegisterIds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint24\",\"name\":\"fid\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"custody\",\"type\":\"address\"}],\"internalType\":\"structIIdRegistry.BulkRegisterDefaultRecoveryData[]\",\"name\":\"ids\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"}],\"name\":\"bulkRegisterIdsWithDefaultRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24[]\",\"name\":\"ids\",\"type\":\"uint24[]\"}],\"name\":\"bulkResetIds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"}],\"name\":\"changeRecoveryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"changeRecoveryAddressFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"}],\"name\":\"custodyOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"custody\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparatorV4\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"freezeIdGateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gatewayFrozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gracePeriod\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"guardians\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isGuardian\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"structHash\",\"type\":\"bytes32\"}],\"name\":\"hashTypedDataV4\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"idCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"idGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"idOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMigrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migratedAt\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"recover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"recoveryDeadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recoverySig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"toDeadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"toSig\",\"type\":\"bytes\"}],\"name\":\"recoverFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"}],\"name\":\"recoveryOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"removeGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_counter\",\"type\":\"uint256\"}],\"name\":\"setIdCounter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_idGateway\",\"type\":\"address\"}],\"name\":\"setIdGateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_migrator\",\"type\":\"address\"}],\"name\":\"setMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"transferAndChangeRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromDeadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"toDeadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"toSig\",\"type\":\"bytes\"}],\"name\":\"transferAndChangeRecoveryFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromDeadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"toDeadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"toSig\",\"type\":\"bytes\"}],\"name\":\"transferFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"custodyAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"verifyFidSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IdRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use IdRegistryMetaData.ABI instead.
var IdRegistryABI = IdRegistryMetaData.ABI

// IdRegistry is an auto generated Go binding around an Ethereum contract.
type IdRegistry struct {
	IdRegistryCaller     // Read-only binding to the contract
	IdRegistryTransactor // Write-only binding to the contract
	IdRegistryFilterer   // Log filterer for contract events
}

// IdRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdRegistrySession struct {
	Contract     *IdRegistry       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdRegistryCallerSession struct {
	Contract *IdRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IdRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdRegistryTransactorSession struct {
	Contract     *IdRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IdRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdRegistryRaw struct {
	Contract *IdRegistry // Generic contract binding to access the raw methods on
}

// IdRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdRegistryCallerRaw struct {
	Contract *IdRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IdRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdRegistryTransactorRaw struct {
	Contract *IdRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdRegistry creates a new instance of IdRegistry, bound to a specific deployed contract.
func NewIdRegistry(address common.Address, backend bind.ContractBackend) (*IdRegistry, error) {
	contract, err := bindIdRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdRegistry{IdRegistryCaller: IdRegistryCaller{contract: contract}, IdRegistryTransactor: IdRegistryTransactor{contract: contract}, IdRegistryFilterer: IdRegistryFilterer{contract: contract}}, nil
}

// NewIdRegistryCaller creates a new read-only instance of IdRegistry, bound to a specific deployed contract.
func NewIdRegistryCaller(address common.Address, caller bind.ContractCaller) (*IdRegistryCaller, error) {
	contract, err := bindIdRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdRegistryCaller{contract: contract}, nil
}

// NewIdRegistryTransactor creates a new write-only instance of IdRegistry, bound to a specific deployed contract.
func NewIdRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IdRegistryTransactor, error) {
	contract, err := bindIdRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdRegistryTransactor{contract: contract}, nil
}

// NewIdRegistryFilterer creates a new log filterer instance of IdRegistry, bound to a specific deployed contract.
func NewIdRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IdRegistryFilterer, error) {
	contract, err := bindIdRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdRegistryFilterer{contract: contract}, nil
}

// bindIdRegistry binds a generic wrapper to an already deployed contract.
func bindIdRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IdRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdRegistry *IdRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdRegistry.Contract.IdRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdRegistry *IdRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdRegistry.Contract.IdRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdRegistry *IdRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdRegistry.Contract.IdRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdRegistry *IdRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdRegistry *IdRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdRegistry *IdRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdRegistry.Contract.contract.Transact(opts, method, params...)
}

// CHANGERECOVERYADDRESSTYPEHASH is a free data retrieval call binding the contract method 0xd5bac7f3.
//
// Solidity: function CHANGE_RECOVERY_ADDRESS_TYPEHASH() view returns(bytes32)
func (_IdRegistry *IdRegistryCaller) CHANGERECOVERYADDRESSTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "CHANGE_RECOVERY_ADDRESS_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CHANGERECOVERYADDRESSTYPEHASH is a free data retrieval call binding the contract method 0xd5bac7f3.
//
// Solidity: function CHANGE_RECOVERY_ADDRESS_TYPEHASH() view returns(bytes32)
func (_IdRegistry *IdRegistrySession) CHANGERECOVERYADDRESSTYPEHASH() ([32]byte, error) {
	return _IdRegistry.Contract.CHANGERECOVERYADDRESSTYPEHASH(&_IdRegistry.CallOpts)
}

// CHANGERECOVERYADDRESSTYPEHASH is a free data retrieval call binding the contract method 0xd5bac7f3.
//
// Solidity: function CHANGE_RECOVERY_ADDRESS_TYPEHASH() view returns(bytes32)
func (_IdRegistry *IdRegistryCallerSession) CHANGERECOVERYADDRESSTYPEHASH() ([32]byte, error) {
	return _IdRegistry.Contract.CHANGERECOVERYADDRESSTYPEHASH(&_IdRegistry.CallOpts)
}

// TRANSFERANDCHANGERECOVERYTYPEHASH is a free data retrieval call binding the contract method 0xea2bbb83.
//
// Solidity: function TRANSFER_AND_CHANGE_RECOVERY_TYPEHASH() view returns(bytes32)
func (_IdRegistry *IdRegistryCaller) TRANSFERANDCHANGERECOVERYTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "TRANSFER_AND_CHANGE_RECOVERY_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TRANSFERANDCHANGERECOVERYTYPEHASH is a free data retrieval call binding the contract method 0xea2bbb83.
//
// Solidity: function TRANSFER_AND_CHANGE_RECOVERY_TYPEHASH() view returns(bytes32)
func (_IdRegistry *IdRegistrySession) TRANSFERANDCHANGERECOVERYTYPEHASH() ([32]byte, error) {
	return _IdRegistry.Contract.TRANSFERANDCHANGERECOVERYTYPEHASH(&_IdRegistry.CallOpts)
}

// TRANSFERANDCHANGERECOVERYTYPEHASH is a free data retrieval call binding the contract method 0xea2bbb83.
//
// Solidity: function TRANSFER_AND_CHANGE_RECOVERY_TYPEHASH() view returns(bytes32)
func (_IdRegistry *IdRegistryCallerSession) TRANSFERANDCHANGERECOVERYTYPEHASH() ([32]byte, error) {
	return _IdRegistry.Contract.TRANSFERANDCHANGERECOVERYTYPEHASH(&_IdRegistry.CallOpts)
}

// TRANSFERTYPEHASH is a free data retrieval call binding the contract method 0x00bf26f4.
//
// Solidity: function TRANSFER_TYPEHASH() view returns(bytes32)
func (_IdRegistry *IdRegistryCaller) TRANSFERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "TRANSFER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TRANSFERTYPEHASH is a free data retrieval call binding the contract method 0x00bf26f4.
//
// Solidity: function TRANSFER_TYPEHASH() view returns(bytes32)
func (_IdRegistry *IdRegistrySession) TRANSFERTYPEHASH() ([32]byte, error) {
	return _IdRegistry.Contract.TRANSFERTYPEHASH(&_IdRegistry.CallOpts)
}

// TRANSFERTYPEHASH is a free data retrieval call binding the contract method 0x00bf26f4.
//
// Solidity: function TRANSFER_TYPEHASH() view returns(bytes32)
func (_IdRegistry *IdRegistryCallerSession) TRANSFERTYPEHASH() ([32]byte, error) {
	return _IdRegistry.Contract.TRANSFERTYPEHASH(&_IdRegistry.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_IdRegistry *IdRegistryCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_IdRegistry *IdRegistrySession) VERSION() (string, error) {
	return _IdRegistry.Contract.VERSION(&_IdRegistry.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_IdRegistry *IdRegistryCallerSession) VERSION() (string, error) {
	return _IdRegistry.Contract.VERSION(&_IdRegistry.CallOpts)
}

// CustodyOf is a free data retrieval call binding the contract method 0x65269e47.
//
// Solidity: function custodyOf(uint256 fid) view returns(address custody)
func (_IdRegistry *IdRegistryCaller) CustodyOf(opts *bind.CallOpts, fid *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "custodyOf", fid)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CustodyOf is a free data retrieval call binding the contract method 0x65269e47.
//
// Solidity: function custodyOf(uint256 fid) view returns(address custody)
func (_IdRegistry *IdRegistrySession) CustodyOf(fid *big.Int) (common.Address, error) {
	return _IdRegistry.Contract.CustodyOf(&_IdRegistry.CallOpts, fid)
}

// CustodyOf is a free data retrieval call binding the contract method 0x65269e47.
//
// Solidity: function custodyOf(uint256 fid) view returns(address custody)
func (_IdRegistry *IdRegistryCallerSession) CustodyOf(fid *big.Int) (common.Address, error) {
	return _IdRegistry.Contract.CustodyOf(&_IdRegistry.CallOpts, fid)
}

// DomainSeparatorV4 is a free data retrieval call binding the contract method 0x78e890ba.
//
// Solidity: function domainSeparatorV4() view returns(bytes32)
func (_IdRegistry *IdRegistryCaller) DomainSeparatorV4(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "domainSeparatorV4")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparatorV4 is a free data retrieval call binding the contract method 0x78e890ba.
//
// Solidity: function domainSeparatorV4() view returns(bytes32)
func (_IdRegistry *IdRegistrySession) DomainSeparatorV4() ([32]byte, error) {
	return _IdRegistry.Contract.DomainSeparatorV4(&_IdRegistry.CallOpts)
}

// DomainSeparatorV4 is a free data retrieval call binding the contract method 0x78e890ba.
//
// Solidity: function domainSeparatorV4() view returns(bytes32)
func (_IdRegistry *IdRegistryCallerSession) DomainSeparatorV4() ([32]byte, error) {
	return _IdRegistry.Contract.DomainSeparatorV4(&_IdRegistry.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_IdRegistry *IdRegistryCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "eip712Domain")

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
func (_IdRegistry *IdRegistrySession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _IdRegistry.Contract.Eip712Domain(&_IdRegistry.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_IdRegistry *IdRegistryCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _IdRegistry.Contract.Eip712Domain(&_IdRegistry.CallOpts)
}

// GatewayFrozen is a free data retrieval call binding the contract method 0x95e7549f.
//
// Solidity: function gatewayFrozen() view returns(bool)
func (_IdRegistry *IdRegistryCaller) GatewayFrozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "gatewayFrozen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GatewayFrozen is a free data retrieval call binding the contract method 0x95e7549f.
//
// Solidity: function gatewayFrozen() view returns(bool)
func (_IdRegistry *IdRegistrySession) GatewayFrozen() (bool, error) {
	return _IdRegistry.Contract.GatewayFrozen(&_IdRegistry.CallOpts)
}

// GatewayFrozen is a free data retrieval call binding the contract method 0x95e7549f.
//
// Solidity: function gatewayFrozen() view returns(bool)
func (_IdRegistry *IdRegistryCallerSession) GatewayFrozen() (bool, error) {
	return _IdRegistry.Contract.GatewayFrozen(&_IdRegistry.CallOpts)
}

// GracePeriod is a free data retrieval call binding the contract method 0xa06db7dc.
//
// Solidity: function gracePeriod() view returns(uint24)
func (_IdRegistry *IdRegistryCaller) GracePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "gracePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GracePeriod is a free data retrieval call binding the contract method 0xa06db7dc.
//
// Solidity: function gracePeriod() view returns(uint24)
func (_IdRegistry *IdRegistrySession) GracePeriod() (*big.Int, error) {
	return _IdRegistry.Contract.GracePeriod(&_IdRegistry.CallOpts)
}

// GracePeriod is a free data retrieval call binding the contract method 0xa06db7dc.
//
// Solidity: function gracePeriod() view returns(uint24)
func (_IdRegistry *IdRegistryCallerSession) GracePeriod() (*big.Int, error) {
	return _IdRegistry.Contract.GracePeriod(&_IdRegistry.CallOpts)
}

// Guardians is a free data retrieval call binding the contract method 0x0633b14a.
//
// Solidity: function guardians(address guardian) view returns(bool isGuardian)
func (_IdRegistry *IdRegistryCaller) Guardians(opts *bind.CallOpts, guardian common.Address) (bool, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "guardians", guardian)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Guardians is a free data retrieval call binding the contract method 0x0633b14a.
//
// Solidity: function guardians(address guardian) view returns(bool isGuardian)
func (_IdRegistry *IdRegistrySession) Guardians(guardian common.Address) (bool, error) {
	return _IdRegistry.Contract.Guardians(&_IdRegistry.CallOpts, guardian)
}

// Guardians is a free data retrieval call binding the contract method 0x0633b14a.
//
// Solidity: function guardians(address guardian) view returns(bool isGuardian)
func (_IdRegistry *IdRegistryCallerSession) Guardians(guardian common.Address) (bool, error) {
	return _IdRegistry.Contract.Guardians(&_IdRegistry.CallOpts, guardian)
}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x4980f288.
//
// Solidity: function hashTypedDataV4(bytes32 structHash) view returns(bytes32)
func (_IdRegistry *IdRegistryCaller) HashTypedDataV4(opts *bind.CallOpts, structHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "hashTypedDataV4", structHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x4980f288.
//
// Solidity: function hashTypedDataV4(bytes32 structHash) view returns(bytes32)
func (_IdRegistry *IdRegistrySession) HashTypedDataV4(structHash [32]byte) ([32]byte, error) {
	return _IdRegistry.Contract.HashTypedDataV4(&_IdRegistry.CallOpts, structHash)
}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x4980f288.
//
// Solidity: function hashTypedDataV4(bytes32 structHash) view returns(bytes32)
func (_IdRegistry *IdRegistryCallerSession) HashTypedDataV4(structHash [32]byte) ([32]byte, error) {
	return _IdRegistry.Contract.HashTypedDataV4(&_IdRegistry.CallOpts, structHash)
}

// IdCounter is a free data retrieval call binding the contract method 0xeb08ab28.
//
// Solidity: function idCounter() view returns(uint256)
func (_IdRegistry *IdRegistryCaller) IdCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "idCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IdCounter is a free data retrieval call binding the contract method 0xeb08ab28.
//
// Solidity: function idCounter() view returns(uint256)
func (_IdRegistry *IdRegistrySession) IdCounter() (*big.Int, error) {
	return _IdRegistry.Contract.IdCounter(&_IdRegistry.CallOpts)
}

// IdCounter is a free data retrieval call binding the contract method 0xeb08ab28.
//
// Solidity: function idCounter() view returns(uint256)
func (_IdRegistry *IdRegistryCallerSession) IdCounter() (*big.Int, error) {
	return _IdRegistry.Contract.IdCounter(&_IdRegistry.CallOpts)
}

// IdGateway is a free data retrieval call binding the contract method 0x4b57a600.
//
// Solidity: function idGateway() view returns(address)
func (_IdRegistry *IdRegistryCaller) IdGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "idGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdGateway is a free data retrieval call binding the contract method 0x4b57a600.
//
// Solidity: function idGateway() view returns(address)
func (_IdRegistry *IdRegistrySession) IdGateway() (common.Address, error) {
	return _IdRegistry.Contract.IdGateway(&_IdRegistry.CallOpts)
}

// IdGateway is a free data retrieval call binding the contract method 0x4b57a600.
//
// Solidity: function idGateway() view returns(address)
func (_IdRegistry *IdRegistryCallerSession) IdGateway() (common.Address, error) {
	return _IdRegistry.Contract.IdGateway(&_IdRegistry.CallOpts)
}

// IdOf is a free data retrieval call binding the contract method 0xd94fe832.
//
// Solidity: function idOf(address owner) view returns(uint256 fid)
func (_IdRegistry *IdRegistryCaller) IdOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "idOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IdOf is a free data retrieval call binding the contract method 0xd94fe832.
//
// Solidity: function idOf(address owner) view returns(uint256 fid)
func (_IdRegistry *IdRegistrySession) IdOf(owner common.Address) (*big.Int, error) {
	return _IdRegistry.Contract.IdOf(&_IdRegistry.CallOpts, owner)
}

// IdOf is a free data retrieval call binding the contract method 0xd94fe832.
//
// Solidity: function idOf(address owner) view returns(uint256 fid)
func (_IdRegistry *IdRegistryCallerSession) IdOf(owner common.Address) (*big.Int, error) {
	return _IdRegistry.Contract.IdOf(&_IdRegistry.CallOpts, owner)
}

// IsMigrated is a free data retrieval call binding the contract method 0xb06faf62.
//
// Solidity: function isMigrated() view returns(bool)
func (_IdRegistry *IdRegistryCaller) IsMigrated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "isMigrated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMigrated is a free data retrieval call binding the contract method 0xb06faf62.
//
// Solidity: function isMigrated() view returns(bool)
func (_IdRegistry *IdRegistrySession) IsMigrated() (bool, error) {
	return _IdRegistry.Contract.IsMigrated(&_IdRegistry.CallOpts)
}

// IsMigrated is a free data retrieval call binding the contract method 0xb06faf62.
//
// Solidity: function isMigrated() view returns(bool)
func (_IdRegistry *IdRegistryCallerSession) IsMigrated() (bool, error) {
	return _IdRegistry.Contract.IsMigrated(&_IdRegistry.CallOpts)
}

// MigratedAt is a free data retrieval call binding the contract method 0x8b21e484.
//
// Solidity: function migratedAt() view returns(uint40)
func (_IdRegistry *IdRegistryCaller) MigratedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "migratedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MigratedAt is a free data retrieval call binding the contract method 0x8b21e484.
//
// Solidity: function migratedAt() view returns(uint40)
func (_IdRegistry *IdRegistrySession) MigratedAt() (*big.Int, error) {
	return _IdRegistry.Contract.MigratedAt(&_IdRegistry.CallOpts)
}

// MigratedAt is a free data retrieval call binding the contract method 0x8b21e484.
//
// Solidity: function migratedAt() view returns(uint40)
func (_IdRegistry *IdRegistryCallerSession) MigratedAt() (*big.Int, error) {
	return _IdRegistry.Contract.MigratedAt(&_IdRegistry.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_IdRegistry *IdRegistryCaller) Migrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "migrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_IdRegistry *IdRegistrySession) Migrator() (common.Address, error) {
	return _IdRegistry.Contract.Migrator(&_IdRegistry.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_IdRegistry *IdRegistryCallerSession) Migrator() (common.Address, error) {
	return _IdRegistry.Contract.Migrator(&_IdRegistry.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IdRegistry *IdRegistryCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IdRegistry *IdRegistrySession) Name() (string, error) {
	return _IdRegistry.Contract.Name(&_IdRegistry.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IdRegistry *IdRegistryCallerSession) Name() (string, error) {
	return _IdRegistry.Contract.Name(&_IdRegistry.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IdRegistry *IdRegistryCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IdRegistry *IdRegistrySession) Nonces(owner common.Address) (*big.Int, error) {
	return _IdRegistry.Contract.Nonces(&_IdRegistry.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IdRegistry *IdRegistryCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IdRegistry.Contract.Nonces(&_IdRegistry.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdRegistry *IdRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdRegistry *IdRegistrySession) Owner() (common.Address, error) {
	return _IdRegistry.Contract.Owner(&_IdRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdRegistry *IdRegistryCallerSession) Owner() (common.Address, error) {
	return _IdRegistry.Contract.Owner(&_IdRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IdRegistry *IdRegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IdRegistry *IdRegistrySession) Paused() (bool, error) {
	return _IdRegistry.Contract.Paused(&_IdRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IdRegistry *IdRegistryCallerSession) Paused() (bool, error) {
	return _IdRegistry.Contract.Paused(&_IdRegistry.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_IdRegistry *IdRegistryCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_IdRegistry *IdRegistrySession) PendingOwner() (common.Address, error) {
	return _IdRegistry.Contract.PendingOwner(&_IdRegistry.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_IdRegistry *IdRegistryCallerSession) PendingOwner() (common.Address, error) {
	return _IdRegistry.Contract.PendingOwner(&_IdRegistry.CallOpts)
}

// RecoveryOf is a free data retrieval call binding the contract method 0xfa1a1b25.
//
// Solidity: function recoveryOf(uint256 fid) view returns(address recovery)
func (_IdRegistry *IdRegistryCaller) RecoveryOf(opts *bind.CallOpts, fid *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "recoveryOf", fid)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoveryOf is a free data retrieval call binding the contract method 0xfa1a1b25.
//
// Solidity: function recoveryOf(uint256 fid) view returns(address recovery)
func (_IdRegistry *IdRegistrySession) RecoveryOf(fid *big.Int) (common.Address, error) {
	return _IdRegistry.Contract.RecoveryOf(&_IdRegistry.CallOpts, fid)
}

// RecoveryOf is a free data retrieval call binding the contract method 0xfa1a1b25.
//
// Solidity: function recoveryOf(uint256 fid) view returns(address recovery)
func (_IdRegistry *IdRegistryCallerSession) RecoveryOf(fid *big.Int) (common.Address, error) {
	return _IdRegistry.Contract.RecoveryOf(&_IdRegistry.CallOpts, fid)
}

// VerifyFidSignature is a free data retrieval call binding the contract method 0x32faac70.
//
// Solidity: function verifyFidSignature(address custodyAddress, uint256 fid, bytes32 digest, bytes sig) view returns(bool isValid)
func (_IdRegistry *IdRegistryCaller) VerifyFidSignature(opts *bind.CallOpts, custodyAddress common.Address, fid *big.Int, digest [32]byte, sig []byte) (bool, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "verifyFidSignature", custodyAddress, fid, digest, sig)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyFidSignature is a free data retrieval call binding the contract method 0x32faac70.
//
// Solidity: function verifyFidSignature(address custodyAddress, uint256 fid, bytes32 digest, bytes sig) view returns(bool isValid)
func (_IdRegistry *IdRegistrySession) VerifyFidSignature(custodyAddress common.Address, fid *big.Int, digest [32]byte, sig []byte) (bool, error) {
	return _IdRegistry.Contract.VerifyFidSignature(&_IdRegistry.CallOpts, custodyAddress, fid, digest, sig)
}

// VerifyFidSignature is a free data retrieval call binding the contract method 0x32faac70.
//
// Solidity: function verifyFidSignature(address custodyAddress, uint256 fid, bytes32 digest, bytes sig) view returns(bool isValid)
func (_IdRegistry *IdRegistryCallerSession) VerifyFidSignature(custodyAddress common.Address, fid *big.Int, digest [32]byte, sig []byte) (bool, error) {
	return _IdRegistry.Contract.VerifyFidSignature(&_IdRegistry.CallOpts, custodyAddress, fid, digest, sig)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_IdRegistry *IdRegistryTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_IdRegistry *IdRegistrySession) AcceptOwnership() (*types.Transaction, error) {
	return _IdRegistry.Contract.AcceptOwnership(&_IdRegistry.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_IdRegistry *IdRegistryTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _IdRegistry.Contract.AcceptOwnership(&_IdRegistry.TransactOpts)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_IdRegistry *IdRegistryTransactor) AddGuardian(opts *bind.TransactOpts, guardian common.Address) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "addGuardian", guardian)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_IdRegistry *IdRegistrySession) AddGuardian(guardian common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.AddGuardian(&_IdRegistry.TransactOpts, guardian)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_IdRegistry *IdRegistryTransactorSession) AddGuardian(guardian common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.AddGuardian(&_IdRegistry.TransactOpts, guardian)
}

// BulkRegisterIds is a paid mutator transaction binding the contract method 0x55c5b358.
//
// Solidity: function bulkRegisterIds((uint24,address,address)[] ids) returns()
func (_IdRegistry *IdRegistryTransactor) BulkRegisterIds(opts *bind.TransactOpts, ids []IIdRegistryBulkRegisterData) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "bulkRegisterIds", ids)
}

// BulkRegisterIds is a paid mutator transaction binding the contract method 0x55c5b358.
//
// Solidity: function bulkRegisterIds((uint24,address,address)[] ids) returns()
func (_IdRegistry *IdRegistrySession) BulkRegisterIds(ids []IIdRegistryBulkRegisterData) (*types.Transaction, error) {
	return _IdRegistry.Contract.BulkRegisterIds(&_IdRegistry.TransactOpts, ids)
}

// BulkRegisterIds is a paid mutator transaction binding the contract method 0x55c5b358.
//
// Solidity: function bulkRegisterIds((uint24,address,address)[] ids) returns()
func (_IdRegistry *IdRegistryTransactorSession) BulkRegisterIds(ids []IIdRegistryBulkRegisterData) (*types.Transaction, error) {
	return _IdRegistry.Contract.BulkRegisterIds(&_IdRegistry.TransactOpts, ids)
}

// BulkRegisterIdsWithDefaultRecovery is a paid mutator transaction binding the contract method 0x8d8043e2.
//
// Solidity: function bulkRegisterIdsWithDefaultRecovery((uint24,address)[] ids, address recovery) returns()
func (_IdRegistry *IdRegistryTransactor) BulkRegisterIdsWithDefaultRecovery(opts *bind.TransactOpts, ids []IIdRegistryBulkRegisterDefaultRecoveryData, recovery common.Address) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "bulkRegisterIdsWithDefaultRecovery", ids, recovery)
}

// BulkRegisterIdsWithDefaultRecovery is a paid mutator transaction binding the contract method 0x8d8043e2.
//
// Solidity: function bulkRegisterIdsWithDefaultRecovery((uint24,address)[] ids, address recovery) returns()
func (_IdRegistry *IdRegistrySession) BulkRegisterIdsWithDefaultRecovery(ids []IIdRegistryBulkRegisterDefaultRecoveryData, recovery common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.BulkRegisterIdsWithDefaultRecovery(&_IdRegistry.TransactOpts, ids, recovery)
}

// BulkRegisterIdsWithDefaultRecovery is a paid mutator transaction binding the contract method 0x8d8043e2.
//
// Solidity: function bulkRegisterIdsWithDefaultRecovery((uint24,address)[] ids, address recovery) returns()
func (_IdRegistry *IdRegistryTransactorSession) BulkRegisterIdsWithDefaultRecovery(ids []IIdRegistryBulkRegisterDefaultRecoveryData, recovery common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.BulkRegisterIdsWithDefaultRecovery(&_IdRegistry.TransactOpts, ids, recovery)
}

// BulkResetIds is a paid mutator transaction binding the contract method 0xff126441.
//
// Solidity: function bulkResetIds(uint24[] ids) returns()
func (_IdRegistry *IdRegistryTransactor) BulkResetIds(opts *bind.TransactOpts, ids []*big.Int) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "bulkResetIds", ids)
}

// BulkResetIds is a paid mutator transaction binding the contract method 0xff126441.
//
// Solidity: function bulkResetIds(uint24[] ids) returns()
func (_IdRegistry *IdRegistrySession) BulkResetIds(ids []*big.Int) (*types.Transaction, error) {
	return _IdRegistry.Contract.BulkResetIds(&_IdRegistry.TransactOpts, ids)
}

// BulkResetIds is a paid mutator transaction binding the contract method 0xff126441.
//
// Solidity: function bulkResetIds(uint24[] ids) returns()
func (_IdRegistry *IdRegistryTransactorSession) BulkResetIds(ids []*big.Int) (*types.Transaction, error) {
	return _IdRegistry.Contract.BulkResetIds(&_IdRegistry.TransactOpts, ids)
}

// ChangeRecoveryAddress is a paid mutator transaction binding the contract method 0xf1f0b224.
//
// Solidity: function changeRecoveryAddress(address recovery) returns()
func (_IdRegistry *IdRegistryTransactor) ChangeRecoveryAddress(opts *bind.TransactOpts, recovery common.Address) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "changeRecoveryAddress", recovery)
}

// ChangeRecoveryAddress is a paid mutator transaction binding the contract method 0xf1f0b224.
//
// Solidity: function changeRecoveryAddress(address recovery) returns()
func (_IdRegistry *IdRegistrySession) ChangeRecoveryAddress(recovery common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.ChangeRecoveryAddress(&_IdRegistry.TransactOpts, recovery)
}

// ChangeRecoveryAddress is a paid mutator transaction binding the contract method 0xf1f0b224.
//
// Solidity: function changeRecoveryAddress(address recovery) returns()
func (_IdRegistry *IdRegistryTransactorSession) ChangeRecoveryAddress(recovery common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.ChangeRecoveryAddress(&_IdRegistry.TransactOpts, recovery)
}

// ChangeRecoveryAddressFor is a paid mutator transaction binding the contract method 0x9cbef8dc.
//
// Solidity: function changeRecoveryAddressFor(address owner, address recovery, uint256 deadline, bytes sig) returns()
func (_IdRegistry *IdRegistryTransactor) ChangeRecoveryAddressFor(opts *bind.TransactOpts, owner common.Address, recovery common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "changeRecoveryAddressFor", owner, recovery, deadline, sig)
}

// ChangeRecoveryAddressFor is a paid mutator transaction binding the contract method 0x9cbef8dc.
//
// Solidity: function changeRecoveryAddressFor(address owner, address recovery, uint256 deadline, bytes sig) returns()
func (_IdRegistry *IdRegistrySession) ChangeRecoveryAddressFor(owner common.Address, recovery common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _IdRegistry.Contract.ChangeRecoveryAddressFor(&_IdRegistry.TransactOpts, owner, recovery, deadline, sig)
}

// ChangeRecoveryAddressFor is a paid mutator transaction binding the contract method 0x9cbef8dc.
//
// Solidity: function changeRecoveryAddressFor(address owner, address recovery, uint256 deadline, bytes sig) returns()
func (_IdRegistry *IdRegistryTransactorSession) ChangeRecoveryAddressFor(owner common.Address, recovery common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _IdRegistry.Contract.ChangeRecoveryAddressFor(&_IdRegistry.TransactOpts, owner, recovery, deadline, sig)
}

// FreezeIdGateway is a paid mutator transaction binding the contract method 0xddd76649.
//
// Solidity: function freezeIdGateway() returns()
func (_IdRegistry *IdRegistryTransactor) FreezeIdGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "freezeIdGateway")
}

// FreezeIdGateway is a paid mutator transaction binding the contract method 0xddd76649.
//
// Solidity: function freezeIdGateway() returns()
func (_IdRegistry *IdRegistrySession) FreezeIdGateway() (*types.Transaction, error) {
	return _IdRegistry.Contract.FreezeIdGateway(&_IdRegistry.TransactOpts)
}

// FreezeIdGateway is a paid mutator transaction binding the contract method 0xddd76649.
//
// Solidity: function freezeIdGateway() returns()
func (_IdRegistry *IdRegistryTransactorSession) FreezeIdGateway() (*types.Transaction, error) {
	return _IdRegistry.Contract.FreezeIdGateway(&_IdRegistry.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_IdRegistry *IdRegistryTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_IdRegistry *IdRegistrySession) Migrate() (*types.Transaction, error) {
	return _IdRegistry.Contract.Migrate(&_IdRegistry.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_IdRegistry *IdRegistryTransactorSession) Migrate() (*types.Transaction, error) {
	return _IdRegistry.Contract.Migrate(&_IdRegistry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IdRegistry *IdRegistryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IdRegistry *IdRegistrySession) Pause() (*types.Transaction, error) {
	return _IdRegistry.Contract.Pause(&_IdRegistry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IdRegistry *IdRegistryTransactorSession) Pause() (*types.Transaction, error) {
	return _IdRegistry.Contract.Pause(&_IdRegistry.TransactOpts)
}

// Recover is a paid mutator transaction binding the contract method 0x2a42ede3.
//
// Solidity: function recover(address from, address to, uint256 deadline, bytes sig) returns()
func (_IdRegistry *IdRegistryTransactor) Recover(opts *bind.TransactOpts, from common.Address, to common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "recover", from, to, deadline, sig)
}

// Recover is a paid mutator transaction binding the contract method 0x2a42ede3.
//
// Solidity: function recover(address from, address to, uint256 deadline, bytes sig) returns()
func (_IdRegistry *IdRegistrySession) Recover(from common.Address, to common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _IdRegistry.Contract.Recover(&_IdRegistry.TransactOpts, from, to, deadline, sig)
}

// Recover is a paid mutator transaction binding the contract method 0x2a42ede3.
//
// Solidity: function recover(address from, address to, uint256 deadline, bytes sig) returns()
func (_IdRegistry *IdRegistryTransactorSession) Recover(from common.Address, to common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _IdRegistry.Contract.Recover(&_IdRegistry.TransactOpts, from, to, deadline, sig)
}

// RecoverFor is a paid mutator transaction binding the contract method 0xba656434.
//
// Solidity: function recoverFor(address from, address to, uint256 recoveryDeadline, bytes recoverySig, uint256 toDeadline, bytes toSig) returns()
func (_IdRegistry *IdRegistryTransactor) RecoverFor(opts *bind.TransactOpts, from common.Address, to common.Address, recoveryDeadline *big.Int, recoverySig []byte, toDeadline *big.Int, toSig []byte) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "recoverFor", from, to, recoveryDeadline, recoverySig, toDeadline, toSig)
}

// RecoverFor is a paid mutator transaction binding the contract method 0xba656434.
//
// Solidity: function recoverFor(address from, address to, uint256 recoveryDeadline, bytes recoverySig, uint256 toDeadline, bytes toSig) returns()
func (_IdRegistry *IdRegistrySession) RecoverFor(from common.Address, to common.Address, recoveryDeadline *big.Int, recoverySig []byte, toDeadline *big.Int, toSig []byte) (*types.Transaction, error) {
	return _IdRegistry.Contract.RecoverFor(&_IdRegistry.TransactOpts, from, to, recoveryDeadline, recoverySig, toDeadline, toSig)
}

// RecoverFor is a paid mutator transaction binding the contract method 0xba656434.
//
// Solidity: function recoverFor(address from, address to, uint256 recoveryDeadline, bytes recoverySig, uint256 toDeadline, bytes toSig) returns()
func (_IdRegistry *IdRegistryTransactorSession) RecoverFor(from common.Address, to common.Address, recoveryDeadline *big.Int, recoverySig []byte, toDeadline *big.Int, toSig []byte) (*types.Transaction, error) {
	return _IdRegistry.Contract.RecoverFor(&_IdRegistry.TransactOpts, from, to, recoveryDeadline, recoverySig, toDeadline, toSig)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address to, address recovery) returns(uint256 fid)
func (_IdRegistry *IdRegistryTransactor) Register(opts *bind.TransactOpts, to common.Address, recovery common.Address) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "register", to, recovery)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address to, address recovery) returns(uint256 fid)
func (_IdRegistry *IdRegistrySession) Register(to common.Address, recovery common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.Register(&_IdRegistry.TransactOpts, to, recovery)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address to, address recovery) returns(uint256 fid)
func (_IdRegistry *IdRegistryTransactorSession) Register(to common.Address, recovery common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.Register(&_IdRegistry.TransactOpts, to, recovery)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_IdRegistry *IdRegistryTransactor) RemoveGuardian(opts *bind.TransactOpts, guardian common.Address) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "removeGuardian", guardian)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_IdRegistry *IdRegistrySession) RemoveGuardian(guardian common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.RemoveGuardian(&_IdRegistry.TransactOpts, guardian)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_IdRegistry *IdRegistryTransactorSession) RemoveGuardian(guardian common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.RemoveGuardian(&_IdRegistry.TransactOpts, guardian)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IdRegistry *IdRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IdRegistry *IdRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _IdRegistry.Contract.RenounceOwnership(&_IdRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IdRegistry *IdRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _IdRegistry.Contract.RenounceOwnership(&_IdRegistry.TransactOpts)
}

// SetIdCounter is a paid mutator transaction binding the contract method 0xa5ed6a6a.
//
// Solidity: function setIdCounter(uint256 _counter) returns()
func (_IdRegistry *IdRegistryTransactor) SetIdCounter(opts *bind.TransactOpts, _counter *big.Int) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "setIdCounter", _counter)
}

// SetIdCounter is a paid mutator transaction binding the contract method 0xa5ed6a6a.
//
// Solidity: function setIdCounter(uint256 _counter) returns()
func (_IdRegistry *IdRegistrySession) SetIdCounter(_counter *big.Int) (*types.Transaction, error) {
	return _IdRegistry.Contract.SetIdCounter(&_IdRegistry.TransactOpts, _counter)
}

// SetIdCounter is a paid mutator transaction binding the contract method 0xa5ed6a6a.
//
// Solidity: function setIdCounter(uint256 _counter) returns()
func (_IdRegistry *IdRegistryTransactorSession) SetIdCounter(_counter *big.Int) (*types.Transaction, error) {
	return _IdRegistry.Contract.SetIdCounter(&_IdRegistry.TransactOpts, _counter)
}

// SetIdGateway is a paid mutator transaction binding the contract method 0x033e2cb3.
//
// Solidity: function setIdGateway(address _idGateway) returns()
func (_IdRegistry *IdRegistryTransactor) SetIdGateway(opts *bind.TransactOpts, _idGateway common.Address) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "setIdGateway", _idGateway)
}

// SetIdGateway is a paid mutator transaction binding the contract method 0x033e2cb3.
//
// Solidity: function setIdGateway(address _idGateway) returns()
func (_IdRegistry *IdRegistrySession) SetIdGateway(_idGateway common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.SetIdGateway(&_IdRegistry.TransactOpts, _idGateway)
}

// SetIdGateway is a paid mutator transaction binding the contract method 0x033e2cb3.
//
// Solidity: function setIdGateway(address _idGateway) returns()
func (_IdRegistry *IdRegistryTransactorSession) SetIdGateway(_idGateway common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.SetIdGateway(&_IdRegistry.TransactOpts, _idGateway)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_IdRegistry *IdRegistryTransactor) SetMigrator(opts *bind.TransactOpts, _migrator common.Address) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "setMigrator", _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_IdRegistry *IdRegistrySession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.SetMigrator(&_IdRegistry.TransactOpts, _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_IdRegistry *IdRegistryTransactorSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.SetMigrator(&_IdRegistry.TransactOpts, _migrator)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(address to, uint256 deadline, bytes sig) returns()
func (_IdRegistry *IdRegistryTransactor) Transfer(opts *bind.TransactOpts, to common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "transfer", to, deadline, sig)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(address to, uint256 deadline, bytes sig) returns()
func (_IdRegistry *IdRegistrySession) Transfer(to common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _IdRegistry.Contract.Transfer(&_IdRegistry.TransactOpts, to, deadline, sig)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(address to, uint256 deadline, bytes sig) returns()
func (_IdRegistry *IdRegistryTransactorSession) Transfer(to common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _IdRegistry.Contract.Transfer(&_IdRegistry.TransactOpts, to, deadline, sig)
}

// TransferAndChangeRecovery is a paid mutator transaction binding the contract method 0x3ab8465d.
//
// Solidity: function transferAndChangeRecovery(address to, address recovery, uint256 deadline, bytes sig) returns()
func (_IdRegistry *IdRegistryTransactor) TransferAndChangeRecovery(opts *bind.TransactOpts, to common.Address, recovery common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "transferAndChangeRecovery", to, recovery, deadline, sig)
}

// TransferAndChangeRecovery is a paid mutator transaction binding the contract method 0x3ab8465d.
//
// Solidity: function transferAndChangeRecovery(address to, address recovery, uint256 deadline, bytes sig) returns()
func (_IdRegistry *IdRegistrySession) TransferAndChangeRecovery(to common.Address, recovery common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _IdRegistry.Contract.TransferAndChangeRecovery(&_IdRegistry.TransactOpts, to, recovery, deadline, sig)
}

// TransferAndChangeRecovery is a paid mutator transaction binding the contract method 0x3ab8465d.
//
// Solidity: function transferAndChangeRecovery(address to, address recovery, uint256 deadline, bytes sig) returns()
func (_IdRegistry *IdRegistryTransactorSession) TransferAndChangeRecovery(to common.Address, recovery common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _IdRegistry.Contract.TransferAndChangeRecovery(&_IdRegistry.TransactOpts, to, recovery, deadline, sig)
}

// TransferAndChangeRecoveryFor is a paid mutator transaction binding the contract method 0x4c5cbb34.
//
// Solidity: function transferAndChangeRecoveryFor(address from, address to, address recovery, uint256 fromDeadline, bytes fromSig, uint256 toDeadline, bytes toSig) returns()
func (_IdRegistry *IdRegistryTransactor) TransferAndChangeRecoveryFor(opts *bind.TransactOpts, from common.Address, to common.Address, recovery common.Address, fromDeadline *big.Int, fromSig []byte, toDeadline *big.Int, toSig []byte) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "transferAndChangeRecoveryFor", from, to, recovery, fromDeadline, fromSig, toDeadline, toSig)
}

// TransferAndChangeRecoveryFor is a paid mutator transaction binding the contract method 0x4c5cbb34.
//
// Solidity: function transferAndChangeRecoveryFor(address from, address to, address recovery, uint256 fromDeadline, bytes fromSig, uint256 toDeadline, bytes toSig) returns()
func (_IdRegistry *IdRegistrySession) TransferAndChangeRecoveryFor(from common.Address, to common.Address, recovery common.Address, fromDeadline *big.Int, fromSig []byte, toDeadline *big.Int, toSig []byte) (*types.Transaction, error) {
	return _IdRegistry.Contract.TransferAndChangeRecoveryFor(&_IdRegistry.TransactOpts, from, to, recovery, fromDeadline, fromSig, toDeadline, toSig)
}

// TransferAndChangeRecoveryFor is a paid mutator transaction binding the contract method 0x4c5cbb34.
//
// Solidity: function transferAndChangeRecoveryFor(address from, address to, address recovery, uint256 fromDeadline, bytes fromSig, uint256 toDeadline, bytes toSig) returns()
func (_IdRegistry *IdRegistryTransactorSession) TransferAndChangeRecoveryFor(from common.Address, to common.Address, recovery common.Address, fromDeadline *big.Int, fromSig []byte, toDeadline *big.Int, toSig []byte) (*types.Transaction, error) {
	return _IdRegistry.Contract.TransferAndChangeRecoveryFor(&_IdRegistry.TransactOpts, from, to, recovery, fromDeadline, fromSig, toDeadline, toSig)
}

// TransferFor is a paid mutator transaction binding the contract method 0x16f72842.
//
// Solidity: function transferFor(address from, address to, uint256 fromDeadline, bytes fromSig, uint256 toDeadline, bytes toSig) returns()
func (_IdRegistry *IdRegistryTransactor) TransferFor(opts *bind.TransactOpts, from common.Address, to common.Address, fromDeadline *big.Int, fromSig []byte, toDeadline *big.Int, toSig []byte) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "transferFor", from, to, fromDeadline, fromSig, toDeadline, toSig)
}

// TransferFor is a paid mutator transaction binding the contract method 0x16f72842.
//
// Solidity: function transferFor(address from, address to, uint256 fromDeadline, bytes fromSig, uint256 toDeadline, bytes toSig) returns()
func (_IdRegistry *IdRegistrySession) TransferFor(from common.Address, to common.Address, fromDeadline *big.Int, fromSig []byte, toDeadline *big.Int, toSig []byte) (*types.Transaction, error) {
	return _IdRegistry.Contract.TransferFor(&_IdRegistry.TransactOpts, from, to, fromDeadline, fromSig, toDeadline, toSig)
}

// TransferFor is a paid mutator transaction binding the contract method 0x16f72842.
//
// Solidity: function transferFor(address from, address to, uint256 fromDeadline, bytes fromSig, uint256 toDeadline, bytes toSig) returns()
func (_IdRegistry *IdRegistryTransactorSession) TransferFor(from common.Address, to common.Address, fromDeadline *big.Int, fromSig []byte, toDeadline *big.Int, toSig []byte) (*types.Transaction, error) {
	return _IdRegistry.Contract.TransferFor(&_IdRegistry.TransactOpts, from, to, fromDeadline, fromSig, toDeadline, toSig)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IdRegistry *IdRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IdRegistry *IdRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.TransferOwnership(&_IdRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IdRegistry *IdRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.TransferOwnership(&_IdRegistry.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IdRegistry *IdRegistryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IdRegistry *IdRegistrySession) Unpause() (*types.Transaction, error) {
	return _IdRegistry.Contract.Unpause(&_IdRegistry.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IdRegistry *IdRegistryTransactorSession) Unpause() (*types.Transaction, error) {
	return _IdRegistry.Contract.Unpause(&_IdRegistry.TransactOpts)
}

// UseNonce is a paid mutator transaction binding the contract method 0x69615a4c.
//
// Solidity: function useNonce() returns(uint256)
func (_IdRegistry *IdRegistryTransactor) UseNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "useNonce")
}

// UseNonce is a paid mutator transaction binding the contract method 0x69615a4c.
//
// Solidity: function useNonce() returns(uint256)
func (_IdRegistry *IdRegistrySession) UseNonce() (*types.Transaction, error) {
	return _IdRegistry.Contract.UseNonce(&_IdRegistry.TransactOpts)
}

// UseNonce is a paid mutator transaction binding the contract method 0x69615a4c.
//
// Solidity: function useNonce() returns(uint256)
func (_IdRegistry *IdRegistryTransactorSession) UseNonce() (*types.Transaction, error) {
	return _IdRegistry.Contract.UseNonce(&_IdRegistry.TransactOpts)
}

// IdRegistryAddIterator is returned from FilterAdd and is used to iterate over the raw logs and unpacked data for Add events raised by the IdRegistry contract.
type IdRegistryAddIterator struct {
	Event *IdRegistryAdd // Event containing the contract specifics and raw log

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
func (it *IdRegistryAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistryAdd)
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
		it.Event = new(IdRegistryAdd)
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
func (it *IdRegistryAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistryAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistryAdd represents a Add event raised by the IdRegistry contract.
type IdRegistryAdd struct {
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdd is a free log retrieval operation binding the contract event 0x87dc5eecd6d6bdeae407c426da6bfba5b7190befc554ed5d4d62dd5cf939fbae.
//
// Solidity: event Add(address indexed guardian)
func (_IdRegistry *IdRegistryFilterer) FilterAdd(opts *bind.FilterOpts, guardian []common.Address) (*IdRegistryAddIterator, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "Add", guardianRule)
	if err != nil {
		return nil, err
	}
	return &IdRegistryAddIterator{contract: _IdRegistry.contract, event: "Add", logs: logs, sub: sub}, nil
}

// WatchAdd is a free log subscription operation binding the contract event 0x87dc5eecd6d6bdeae407c426da6bfba5b7190befc554ed5d4d62dd5cf939fbae.
//
// Solidity: event Add(address indexed guardian)
func (_IdRegistry *IdRegistryFilterer) WatchAdd(opts *bind.WatchOpts, sink chan<- *IdRegistryAdd, guardian []common.Address) (event.Subscription, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "Add", guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistryAdd)
				if err := _IdRegistry.contract.UnpackLog(event, "Add", log); err != nil {
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
func (_IdRegistry *IdRegistryFilterer) ParseAdd(log types.Log) (*IdRegistryAdd, error) {
	event := new(IdRegistryAdd)
	if err := _IdRegistry.contract.UnpackLog(event, "Add", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdRegistryAdminResetIterator is returned from FilterAdminReset and is used to iterate over the raw logs and unpacked data for AdminReset events raised by the IdRegistry contract.
type IdRegistryAdminResetIterator struct {
	Event *IdRegistryAdminReset // Event containing the contract specifics and raw log

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
func (it *IdRegistryAdminResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistryAdminReset)
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
		it.Event = new(IdRegistryAdminReset)
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
func (it *IdRegistryAdminResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistryAdminResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistryAdminReset represents a AdminReset event raised by the IdRegistry contract.
type IdRegistryAdminReset struct {
	Fid *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAdminReset is a free log retrieval operation binding the contract event 0x8b4b4c6da5b89da518fb865149e01ad2863b48861a8b952e11645f663959fa70.
//
// Solidity: event AdminReset(uint256 indexed fid)
func (_IdRegistry *IdRegistryFilterer) FilterAdminReset(opts *bind.FilterOpts, fid []*big.Int) (*IdRegistryAdminResetIterator, error) {

	var fidRule []interface{}
	for _, fidItem := range fid {
		fidRule = append(fidRule, fidItem)
	}

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "AdminReset", fidRule)
	if err != nil {
		return nil, err
	}
	return &IdRegistryAdminResetIterator{contract: _IdRegistry.contract, event: "AdminReset", logs: logs, sub: sub}, nil
}

// WatchAdminReset is a free log subscription operation binding the contract event 0x8b4b4c6da5b89da518fb865149e01ad2863b48861a8b952e11645f663959fa70.
//
// Solidity: event AdminReset(uint256 indexed fid)
func (_IdRegistry *IdRegistryFilterer) WatchAdminReset(opts *bind.WatchOpts, sink chan<- *IdRegistryAdminReset, fid []*big.Int) (event.Subscription, error) {

	var fidRule []interface{}
	for _, fidItem := range fid {
		fidRule = append(fidRule, fidItem)
	}

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "AdminReset", fidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistryAdminReset)
				if err := _IdRegistry.contract.UnpackLog(event, "AdminReset", log); err != nil {
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

// ParseAdminReset is a log parse operation binding the contract event 0x8b4b4c6da5b89da518fb865149e01ad2863b48861a8b952e11645f663959fa70.
//
// Solidity: event AdminReset(uint256 indexed fid)
func (_IdRegistry *IdRegistryFilterer) ParseAdminReset(log types.Log) (*IdRegistryAdminReset, error) {
	event := new(IdRegistryAdminReset)
	if err := _IdRegistry.contract.UnpackLog(event, "AdminReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdRegistryChangeRecoveryAddressIterator is returned from FilterChangeRecoveryAddress and is used to iterate over the raw logs and unpacked data for ChangeRecoveryAddress events raised by the IdRegistry contract.
type IdRegistryChangeRecoveryAddressIterator struct {
	Event *IdRegistryChangeRecoveryAddress // Event containing the contract specifics and raw log

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
func (it *IdRegistryChangeRecoveryAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistryChangeRecoveryAddress)
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
		it.Event = new(IdRegistryChangeRecoveryAddress)
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
func (it *IdRegistryChangeRecoveryAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistryChangeRecoveryAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistryChangeRecoveryAddress represents a ChangeRecoveryAddress event raised by the IdRegistry contract.
type IdRegistryChangeRecoveryAddress struct {
	Id       *big.Int
	Recovery common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChangeRecoveryAddress is a free log retrieval operation binding the contract event 0x8e700b803af43e14651431cd73c9fe7d11b131ad797576a70b893ce5766f65c3.
//
// Solidity: event ChangeRecoveryAddress(uint256 indexed id, address indexed recovery)
func (_IdRegistry *IdRegistryFilterer) FilterChangeRecoveryAddress(opts *bind.FilterOpts, id []*big.Int, recovery []common.Address) (*IdRegistryChangeRecoveryAddressIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var recoveryRule []interface{}
	for _, recoveryItem := range recovery {
		recoveryRule = append(recoveryRule, recoveryItem)
	}

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "ChangeRecoveryAddress", idRule, recoveryRule)
	if err != nil {
		return nil, err
	}
	return &IdRegistryChangeRecoveryAddressIterator{contract: _IdRegistry.contract, event: "ChangeRecoveryAddress", logs: logs, sub: sub}, nil
}

// WatchChangeRecoveryAddress is a free log subscription operation binding the contract event 0x8e700b803af43e14651431cd73c9fe7d11b131ad797576a70b893ce5766f65c3.
//
// Solidity: event ChangeRecoveryAddress(uint256 indexed id, address indexed recovery)
func (_IdRegistry *IdRegistryFilterer) WatchChangeRecoveryAddress(opts *bind.WatchOpts, sink chan<- *IdRegistryChangeRecoveryAddress, id []*big.Int, recovery []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var recoveryRule []interface{}
	for _, recoveryItem := range recovery {
		recoveryRule = append(recoveryRule, recoveryItem)
	}

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "ChangeRecoveryAddress", idRule, recoveryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistryChangeRecoveryAddress)
				if err := _IdRegistry.contract.UnpackLog(event, "ChangeRecoveryAddress", log); err != nil {
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

// ParseChangeRecoveryAddress is a log parse operation binding the contract event 0x8e700b803af43e14651431cd73c9fe7d11b131ad797576a70b893ce5766f65c3.
//
// Solidity: event ChangeRecoveryAddress(uint256 indexed id, address indexed recovery)
func (_IdRegistry *IdRegistryFilterer) ParseChangeRecoveryAddress(log types.Log) (*IdRegistryChangeRecoveryAddress, error) {
	event := new(IdRegistryChangeRecoveryAddress)
	if err := _IdRegistry.contract.UnpackLog(event, "ChangeRecoveryAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdRegistryEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the IdRegistry contract.
type IdRegistryEIP712DomainChangedIterator struct {
	Event *IdRegistryEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *IdRegistryEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistryEIP712DomainChanged)
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
		it.Event = new(IdRegistryEIP712DomainChanged)
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
func (it *IdRegistryEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistryEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistryEIP712DomainChanged represents a EIP712DomainChanged event raised by the IdRegistry contract.
type IdRegistryEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_IdRegistry *IdRegistryFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*IdRegistryEIP712DomainChangedIterator, error) {

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &IdRegistryEIP712DomainChangedIterator{contract: _IdRegistry.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_IdRegistry *IdRegistryFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *IdRegistryEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistryEIP712DomainChanged)
				if err := _IdRegistry.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_IdRegistry *IdRegistryFilterer) ParseEIP712DomainChanged(log types.Log) (*IdRegistryEIP712DomainChanged, error) {
	event := new(IdRegistryEIP712DomainChanged)
	if err := _IdRegistry.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdRegistryFreezeIdGatewayIterator is returned from FilterFreezeIdGateway and is used to iterate over the raw logs and unpacked data for FreezeIdGateway events raised by the IdRegistry contract.
type IdRegistryFreezeIdGatewayIterator struct {
	Event *IdRegistryFreezeIdGateway // Event containing the contract specifics and raw log

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
func (it *IdRegistryFreezeIdGatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistryFreezeIdGateway)
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
		it.Event = new(IdRegistryFreezeIdGateway)
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
func (it *IdRegistryFreezeIdGatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistryFreezeIdGatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistryFreezeIdGateway represents a FreezeIdGateway event raised by the IdRegistry contract.
type IdRegistryFreezeIdGateway struct {
	IdGateway common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreezeIdGateway is a free log retrieval operation binding the contract event 0x1f54688ee839cb2e57222a4f7482fd67a532a36666748891a7634428b2e8a153.
//
// Solidity: event FreezeIdGateway(address idGateway)
func (_IdRegistry *IdRegistryFilterer) FilterFreezeIdGateway(opts *bind.FilterOpts) (*IdRegistryFreezeIdGatewayIterator, error) {

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "FreezeIdGateway")
	if err != nil {
		return nil, err
	}
	return &IdRegistryFreezeIdGatewayIterator{contract: _IdRegistry.contract, event: "FreezeIdGateway", logs: logs, sub: sub}, nil
}

// WatchFreezeIdGateway is a free log subscription operation binding the contract event 0x1f54688ee839cb2e57222a4f7482fd67a532a36666748891a7634428b2e8a153.
//
// Solidity: event FreezeIdGateway(address idGateway)
func (_IdRegistry *IdRegistryFilterer) WatchFreezeIdGateway(opts *bind.WatchOpts, sink chan<- *IdRegistryFreezeIdGateway) (event.Subscription, error) {

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "FreezeIdGateway")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistryFreezeIdGateway)
				if err := _IdRegistry.contract.UnpackLog(event, "FreezeIdGateway", log); err != nil {
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

// ParseFreezeIdGateway is a log parse operation binding the contract event 0x1f54688ee839cb2e57222a4f7482fd67a532a36666748891a7634428b2e8a153.
//
// Solidity: event FreezeIdGateway(address idGateway)
func (_IdRegistry *IdRegistryFilterer) ParseFreezeIdGateway(log types.Log) (*IdRegistryFreezeIdGateway, error) {
	event := new(IdRegistryFreezeIdGateway)
	if err := _IdRegistry.contract.UnpackLog(event, "FreezeIdGateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdRegistryMigratedIterator is returned from FilterMigrated and is used to iterate over the raw logs and unpacked data for Migrated events raised by the IdRegistry contract.
type IdRegistryMigratedIterator struct {
	Event *IdRegistryMigrated // Event containing the contract specifics and raw log

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
func (it *IdRegistryMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistryMigrated)
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
		it.Event = new(IdRegistryMigrated)
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
func (it *IdRegistryMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistryMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistryMigrated represents a Migrated event raised by the IdRegistry contract.
type IdRegistryMigrated struct {
	MigratedAt *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMigrated is a free log retrieval operation binding the contract event 0xe4a25c0c2cbe89d6ad8b64c61a7dbdd20d1f781f6023f1ab94ebb7fe0aef6ab8.
//
// Solidity: event Migrated(uint256 indexed migratedAt)
func (_IdRegistry *IdRegistryFilterer) FilterMigrated(opts *bind.FilterOpts, migratedAt []*big.Int) (*IdRegistryMigratedIterator, error) {

	var migratedAtRule []interface{}
	for _, migratedAtItem := range migratedAt {
		migratedAtRule = append(migratedAtRule, migratedAtItem)
	}

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "Migrated", migratedAtRule)
	if err != nil {
		return nil, err
	}
	return &IdRegistryMigratedIterator{contract: _IdRegistry.contract, event: "Migrated", logs: logs, sub: sub}, nil
}

// WatchMigrated is a free log subscription operation binding the contract event 0xe4a25c0c2cbe89d6ad8b64c61a7dbdd20d1f781f6023f1ab94ebb7fe0aef6ab8.
//
// Solidity: event Migrated(uint256 indexed migratedAt)
func (_IdRegistry *IdRegistryFilterer) WatchMigrated(opts *bind.WatchOpts, sink chan<- *IdRegistryMigrated, migratedAt []*big.Int) (event.Subscription, error) {

	var migratedAtRule []interface{}
	for _, migratedAtItem := range migratedAt {
		migratedAtRule = append(migratedAtRule, migratedAtItem)
	}

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "Migrated", migratedAtRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistryMigrated)
				if err := _IdRegistry.contract.UnpackLog(event, "Migrated", log); err != nil {
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

// ParseMigrated is a log parse operation binding the contract event 0xe4a25c0c2cbe89d6ad8b64c61a7dbdd20d1f781f6023f1ab94ebb7fe0aef6ab8.
//
// Solidity: event Migrated(uint256 indexed migratedAt)
func (_IdRegistry *IdRegistryFilterer) ParseMigrated(log types.Log) (*IdRegistryMigrated, error) {
	event := new(IdRegistryMigrated)
	if err := _IdRegistry.contract.UnpackLog(event, "Migrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdRegistryOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the IdRegistry contract.
type IdRegistryOwnershipTransferStartedIterator struct {
	Event *IdRegistryOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *IdRegistryOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistryOwnershipTransferStarted)
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
		it.Event = new(IdRegistryOwnershipTransferStarted)
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
func (it *IdRegistryOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistryOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistryOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the IdRegistry contract.
type IdRegistryOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_IdRegistry *IdRegistryFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IdRegistryOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IdRegistryOwnershipTransferStartedIterator{contract: _IdRegistry.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_IdRegistry *IdRegistryFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *IdRegistryOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistryOwnershipTransferStarted)
				if err := _IdRegistry.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_IdRegistry *IdRegistryFilterer) ParseOwnershipTransferStarted(log types.Log) (*IdRegistryOwnershipTransferStarted, error) {
	event := new(IdRegistryOwnershipTransferStarted)
	if err := _IdRegistry.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IdRegistry contract.
type IdRegistryOwnershipTransferredIterator struct {
	Event *IdRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IdRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistryOwnershipTransferred)
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
		it.Event = new(IdRegistryOwnershipTransferred)
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
func (it *IdRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the IdRegistry contract.
type IdRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IdRegistry *IdRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IdRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IdRegistryOwnershipTransferredIterator{contract: _IdRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IdRegistry *IdRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IdRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistryOwnershipTransferred)
				if err := _IdRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_IdRegistry *IdRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*IdRegistryOwnershipTransferred, error) {
	event := new(IdRegistryOwnershipTransferred)
	if err := _IdRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdRegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the IdRegistry contract.
type IdRegistryPausedIterator struct {
	Event *IdRegistryPaused // Event containing the contract specifics and raw log

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
func (it *IdRegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistryPaused)
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
		it.Event = new(IdRegistryPaused)
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
func (it *IdRegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistryPaused represents a Paused event raised by the IdRegistry contract.
type IdRegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IdRegistry *IdRegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*IdRegistryPausedIterator, error) {

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &IdRegistryPausedIterator{contract: _IdRegistry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IdRegistry *IdRegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *IdRegistryPaused) (event.Subscription, error) {

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistryPaused)
				if err := _IdRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_IdRegistry *IdRegistryFilterer) ParsePaused(log types.Log) (*IdRegistryPaused, error) {
	event := new(IdRegistryPaused)
	if err := _IdRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdRegistryRecoverIterator is returned from FilterRecover and is used to iterate over the raw logs and unpacked data for Recover events raised by the IdRegistry contract.
type IdRegistryRecoverIterator struct {
	Event *IdRegistryRecover // Event containing the contract specifics and raw log

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
func (it *IdRegistryRecoverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistryRecover)
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
		it.Event = new(IdRegistryRecover)
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
func (it *IdRegistryRecoverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistryRecoverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistryRecover represents a Recover event raised by the IdRegistry contract.
type IdRegistryRecover struct {
	From common.Address
	To   common.Address
	Id   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRecover is a free log retrieval operation binding the contract event 0xf6891c84a6c6af32a6d052172a8acc4c631b1d5057ffa2bc1da268b6938ea2da.
//
// Solidity: event Recover(address indexed from, address indexed to, uint256 indexed id)
func (_IdRegistry *IdRegistryFilterer) FilterRecover(opts *bind.FilterOpts, from []common.Address, to []common.Address, id []*big.Int) (*IdRegistryRecoverIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "Recover", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return &IdRegistryRecoverIterator{contract: _IdRegistry.contract, event: "Recover", logs: logs, sub: sub}, nil
}

// WatchRecover is a free log subscription operation binding the contract event 0xf6891c84a6c6af32a6d052172a8acc4c631b1d5057ffa2bc1da268b6938ea2da.
//
// Solidity: event Recover(address indexed from, address indexed to, uint256 indexed id)
func (_IdRegistry *IdRegistryFilterer) WatchRecover(opts *bind.WatchOpts, sink chan<- *IdRegistryRecover, from []common.Address, to []common.Address, id []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "Recover", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistryRecover)
				if err := _IdRegistry.contract.UnpackLog(event, "Recover", log); err != nil {
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

// ParseRecover is a log parse operation binding the contract event 0xf6891c84a6c6af32a6d052172a8acc4c631b1d5057ffa2bc1da268b6938ea2da.
//
// Solidity: event Recover(address indexed from, address indexed to, uint256 indexed id)
func (_IdRegistry *IdRegistryFilterer) ParseRecover(log types.Log) (*IdRegistryRecover, error) {
	event := new(IdRegistryRecover)
	if err := _IdRegistry.contract.UnpackLog(event, "Recover", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdRegistryRegisterIterator is returned from FilterRegister and is used to iterate over the raw logs and unpacked data for Register events raised by the IdRegistry contract.
type IdRegistryRegisterIterator struct {
	Event *IdRegistryRegister // Event containing the contract specifics and raw log

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
func (it *IdRegistryRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistryRegister)
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
		it.Event = new(IdRegistryRegister)
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
func (it *IdRegistryRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistryRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistryRegister represents a Register event raised by the IdRegistry contract.
type IdRegistryRegister struct {
	To       common.Address
	Id       *big.Int
	Recovery common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRegister is a free log retrieval operation binding the contract event 0xf2e19a901b0748d8b08e428d0468896a039ac751ec4fec49b44b7b9c28097e45.
//
// Solidity: event Register(address indexed to, uint256 indexed id, address recovery)
func (_IdRegistry *IdRegistryFilterer) FilterRegister(opts *bind.FilterOpts, to []common.Address, id []*big.Int) (*IdRegistryRegisterIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "Register", toRule, idRule)
	if err != nil {
		return nil, err
	}
	return &IdRegistryRegisterIterator{contract: _IdRegistry.contract, event: "Register", logs: logs, sub: sub}, nil
}

// WatchRegister is a free log subscription operation binding the contract event 0xf2e19a901b0748d8b08e428d0468896a039ac751ec4fec49b44b7b9c28097e45.
//
// Solidity: event Register(address indexed to, uint256 indexed id, address recovery)
func (_IdRegistry *IdRegistryFilterer) WatchRegister(opts *bind.WatchOpts, sink chan<- *IdRegistryRegister, to []common.Address, id []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "Register", toRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistryRegister)
				if err := _IdRegistry.contract.UnpackLog(event, "Register", log); err != nil {
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

// ParseRegister is a log parse operation binding the contract event 0xf2e19a901b0748d8b08e428d0468896a039ac751ec4fec49b44b7b9c28097e45.
//
// Solidity: event Register(address indexed to, uint256 indexed id, address recovery)
func (_IdRegistry *IdRegistryFilterer) ParseRegister(log types.Log) (*IdRegistryRegister, error) {
	event := new(IdRegistryRegister)
	if err := _IdRegistry.contract.UnpackLog(event, "Register", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdRegistryRemoveIterator is returned from FilterRemove and is used to iterate over the raw logs and unpacked data for Remove events raised by the IdRegistry contract.
type IdRegistryRemoveIterator struct {
	Event *IdRegistryRemove // Event containing the contract specifics and raw log

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
func (it *IdRegistryRemoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistryRemove)
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
		it.Event = new(IdRegistryRemove)
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
func (it *IdRegistryRemoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistryRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistryRemove represents a Remove event raised by the IdRegistry contract.
type IdRegistryRemove struct {
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemove is a free log retrieval operation binding the contract event 0xbe7c7ac3248df4581c206a84aab3cb4e7d521b5398b42b681757f78a5a7d411e.
//
// Solidity: event Remove(address indexed guardian)
func (_IdRegistry *IdRegistryFilterer) FilterRemove(opts *bind.FilterOpts, guardian []common.Address) (*IdRegistryRemoveIterator, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "Remove", guardianRule)
	if err != nil {
		return nil, err
	}
	return &IdRegistryRemoveIterator{contract: _IdRegistry.contract, event: "Remove", logs: logs, sub: sub}, nil
}

// WatchRemove is a free log subscription operation binding the contract event 0xbe7c7ac3248df4581c206a84aab3cb4e7d521b5398b42b681757f78a5a7d411e.
//
// Solidity: event Remove(address indexed guardian)
func (_IdRegistry *IdRegistryFilterer) WatchRemove(opts *bind.WatchOpts, sink chan<- *IdRegistryRemove, guardian []common.Address) (event.Subscription, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "Remove", guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistryRemove)
				if err := _IdRegistry.contract.UnpackLog(event, "Remove", log); err != nil {
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
func (_IdRegistry *IdRegistryFilterer) ParseRemove(log types.Log) (*IdRegistryRemove, error) {
	event := new(IdRegistryRemove)
	if err := _IdRegistry.contract.UnpackLog(event, "Remove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdRegistrySetIdCounterIterator is returned from FilterSetIdCounter and is used to iterate over the raw logs and unpacked data for SetIdCounter events raised by the IdRegistry contract.
type IdRegistrySetIdCounterIterator struct {
	Event *IdRegistrySetIdCounter // Event containing the contract specifics and raw log

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
func (it *IdRegistrySetIdCounterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistrySetIdCounter)
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
		it.Event = new(IdRegistrySetIdCounter)
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
func (it *IdRegistrySetIdCounterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistrySetIdCounterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistrySetIdCounter represents a SetIdCounter event raised by the IdRegistry contract.
type IdRegistrySetIdCounter struct {
	OldCounter *big.Int
	NewCounter *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetIdCounter is a free log retrieval operation binding the contract event 0x562044dce594b5c0ac495e6cf3717dbef4dcc96bf978ff452457bfccd68a4eed.
//
// Solidity: event SetIdCounter(uint256 oldCounter, uint256 newCounter)
func (_IdRegistry *IdRegistryFilterer) FilterSetIdCounter(opts *bind.FilterOpts) (*IdRegistrySetIdCounterIterator, error) {

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "SetIdCounter")
	if err != nil {
		return nil, err
	}
	return &IdRegistrySetIdCounterIterator{contract: _IdRegistry.contract, event: "SetIdCounter", logs: logs, sub: sub}, nil
}

// WatchSetIdCounter is a free log subscription operation binding the contract event 0x562044dce594b5c0ac495e6cf3717dbef4dcc96bf978ff452457bfccd68a4eed.
//
// Solidity: event SetIdCounter(uint256 oldCounter, uint256 newCounter)
func (_IdRegistry *IdRegistryFilterer) WatchSetIdCounter(opts *bind.WatchOpts, sink chan<- *IdRegistrySetIdCounter) (event.Subscription, error) {

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "SetIdCounter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistrySetIdCounter)
				if err := _IdRegistry.contract.UnpackLog(event, "SetIdCounter", log); err != nil {
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

// ParseSetIdCounter is a log parse operation binding the contract event 0x562044dce594b5c0ac495e6cf3717dbef4dcc96bf978ff452457bfccd68a4eed.
//
// Solidity: event SetIdCounter(uint256 oldCounter, uint256 newCounter)
func (_IdRegistry *IdRegistryFilterer) ParseSetIdCounter(log types.Log) (*IdRegistrySetIdCounter, error) {
	event := new(IdRegistrySetIdCounter)
	if err := _IdRegistry.contract.UnpackLog(event, "SetIdCounter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdRegistrySetIdGatewayIterator is returned from FilterSetIdGateway and is used to iterate over the raw logs and unpacked data for SetIdGateway events raised by the IdRegistry contract.
type IdRegistrySetIdGatewayIterator struct {
	Event *IdRegistrySetIdGateway // Event containing the contract specifics and raw log

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
func (it *IdRegistrySetIdGatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistrySetIdGateway)
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
		it.Event = new(IdRegistrySetIdGateway)
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
func (it *IdRegistrySetIdGatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistrySetIdGatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistrySetIdGateway represents a SetIdGateway event raised by the IdRegistry contract.
type IdRegistrySetIdGateway struct {
	OldIdGateway common.Address
	NewIdGateway common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetIdGateway is a free log retrieval operation binding the contract event 0x306b123921c19a8629c68977f4dfea9ef9d5a6dedfafcd0d4a70ac6c9b763ac2.
//
// Solidity: event SetIdGateway(address oldIdGateway, address newIdGateway)
func (_IdRegistry *IdRegistryFilterer) FilterSetIdGateway(opts *bind.FilterOpts) (*IdRegistrySetIdGatewayIterator, error) {

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "SetIdGateway")
	if err != nil {
		return nil, err
	}
	return &IdRegistrySetIdGatewayIterator{contract: _IdRegistry.contract, event: "SetIdGateway", logs: logs, sub: sub}, nil
}

// WatchSetIdGateway is a free log subscription operation binding the contract event 0x306b123921c19a8629c68977f4dfea9ef9d5a6dedfafcd0d4a70ac6c9b763ac2.
//
// Solidity: event SetIdGateway(address oldIdGateway, address newIdGateway)
func (_IdRegistry *IdRegistryFilterer) WatchSetIdGateway(opts *bind.WatchOpts, sink chan<- *IdRegistrySetIdGateway) (event.Subscription, error) {

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "SetIdGateway")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistrySetIdGateway)
				if err := _IdRegistry.contract.UnpackLog(event, "SetIdGateway", log); err != nil {
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

// ParseSetIdGateway is a log parse operation binding the contract event 0x306b123921c19a8629c68977f4dfea9ef9d5a6dedfafcd0d4a70ac6c9b763ac2.
//
// Solidity: event SetIdGateway(address oldIdGateway, address newIdGateway)
func (_IdRegistry *IdRegistryFilterer) ParseSetIdGateway(log types.Log) (*IdRegistrySetIdGateway, error) {
	event := new(IdRegistrySetIdGateway)
	if err := _IdRegistry.contract.UnpackLog(event, "SetIdGateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdRegistrySetMigratorIterator is returned from FilterSetMigrator and is used to iterate over the raw logs and unpacked data for SetMigrator events raised by the IdRegistry contract.
type IdRegistrySetMigratorIterator struct {
	Event *IdRegistrySetMigrator // Event containing the contract specifics and raw log

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
func (it *IdRegistrySetMigratorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistrySetMigrator)
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
		it.Event = new(IdRegistrySetMigrator)
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
func (it *IdRegistrySetMigratorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistrySetMigratorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistrySetMigrator represents a SetMigrator event raised by the IdRegistry contract.
type IdRegistrySetMigrator struct {
	OldMigrator common.Address
	NewMigrator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetMigrator is a free log retrieval operation binding the contract event 0xd8ad954fe808212ab9ed7139873e40807dff7995fe36e3d6cdeb8fa00fcebf10.
//
// Solidity: event SetMigrator(address oldMigrator, address newMigrator)
func (_IdRegistry *IdRegistryFilterer) FilterSetMigrator(opts *bind.FilterOpts) (*IdRegistrySetMigratorIterator, error) {

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "SetMigrator")
	if err != nil {
		return nil, err
	}
	return &IdRegistrySetMigratorIterator{contract: _IdRegistry.contract, event: "SetMigrator", logs: logs, sub: sub}, nil
}

// WatchSetMigrator is a free log subscription operation binding the contract event 0xd8ad954fe808212ab9ed7139873e40807dff7995fe36e3d6cdeb8fa00fcebf10.
//
// Solidity: event SetMigrator(address oldMigrator, address newMigrator)
func (_IdRegistry *IdRegistryFilterer) WatchSetMigrator(opts *bind.WatchOpts, sink chan<- *IdRegistrySetMigrator) (event.Subscription, error) {

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "SetMigrator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistrySetMigrator)
				if err := _IdRegistry.contract.UnpackLog(event, "SetMigrator", log); err != nil {
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

// ParseSetMigrator is a log parse operation binding the contract event 0xd8ad954fe808212ab9ed7139873e40807dff7995fe36e3d6cdeb8fa00fcebf10.
//
// Solidity: event SetMigrator(address oldMigrator, address newMigrator)
func (_IdRegistry *IdRegistryFilterer) ParseSetMigrator(log types.Log) (*IdRegistrySetMigrator, error) {
	event := new(IdRegistrySetMigrator)
	if err := _IdRegistry.contract.UnpackLog(event, "SetMigrator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdRegistryTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IdRegistry contract.
type IdRegistryTransferIterator struct {
	Event *IdRegistryTransfer // Event containing the contract specifics and raw log

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
func (it *IdRegistryTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistryTransfer)
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
		it.Event = new(IdRegistryTransfer)
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
func (it *IdRegistryTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistryTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistryTransfer represents a Transfer event raised by the IdRegistry contract.
type IdRegistryTransfer struct {
	From common.Address
	To   common.Address
	Id   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_IdRegistry *IdRegistryFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, id []*big.Int) (*IdRegistryTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return &IdRegistryTransferIterator{contract: _IdRegistry.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_IdRegistry *IdRegistryFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IdRegistryTransfer, from []common.Address, to []common.Address, id []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistryTransfer)
				if err := _IdRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_IdRegistry *IdRegistryFilterer) ParseTransfer(log types.Log) (*IdRegistryTransfer, error) {
	event := new(IdRegistryTransfer)
	if err := _IdRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdRegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the IdRegistry contract.
type IdRegistryUnpausedIterator struct {
	Event *IdRegistryUnpaused // Event containing the contract specifics and raw log

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
func (it *IdRegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistryUnpaused)
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
		it.Event = new(IdRegistryUnpaused)
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
func (it *IdRegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistryUnpaused represents a Unpaused event raised by the IdRegistry contract.
type IdRegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IdRegistry *IdRegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*IdRegistryUnpausedIterator, error) {

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &IdRegistryUnpausedIterator{contract: _IdRegistry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IdRegistry *IdRegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *IdRegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistryUnpaused)
				if err := _IdRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_IdRegistry *IdRegistryFilterer) ParseUnpaused(log types.Log) (*IdRegistryUnpaused, error) {
	event := new(IdRegistryUnpaused)
	if err := _IdRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
