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

// IKeyRegistryBulkAddData is an auto generated low-level Go binding around an user-defined struct.
type IKeyRegistryBulkAddData struct {
	Fid  *big.Int
	Keys []IKeyRegistryBulkAddKey
}

// IKeyRegistryBulkAddKey is an auto generated low-level Go binding around an user-defined struct.
type IKeyRegistryBulkAddKey struct {
	Key      []byte
	Metadata []byte
}

// IKeyRegistryBulkResetData is an auto generated low-level Go binding around an user-defined struct.
type IKeyRegistryBulkResetData struct {
	Fid  *big.Int
	Keys [][]byte
}

// IKeyRegistryKeyData is an auto generated low-level Go binding around an user-defined struct.
type IKeyRegistryKeyData struct {
	State   uint8
	KeyType uint32
}

// KeyRegistryMetaData contains all meta data concerning the KeyRegistry contract.
var KeyRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_idRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_migrator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxKeysPerFid\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyMigrated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedsMaximum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GatewayFrozen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxKeys\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMetadata\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMetadataType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGuardian\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyMigrator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissionRevoked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"keyType\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"metadataType\",\"type\":\"uint8\"}],\"name\":\"ValidatorNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"keyType\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"keyBytes\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"metadataType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"Add\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"Add\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"keyBytes\",\"type\":\"bytes\"}],\"name\":\"AdminReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"keyGateway\",\"type\":\"address\"}],\"name\":\"FreezeKeyGateway\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"migratedAt\",\"type\":\"uint256\"}],\"name\":\"Migrated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"keyBytes\",\"type\":\"bytes\"}],\"name\":\"Remove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"Remove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldIdRegistry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newIdRegistry\",\"type\":\"address\"}],\"name\":\"SetIdRegistry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldKeyGateway\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newKeyGateway\",\"type\":\"address\"}],\"name\":\"SetKeyGateway\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMax\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMax\",\"type\":\"uint256\"}],\"name\":\"SetMaxKeysPerFid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldMigrator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newMigrator\",\"type\":\"address\"}],\"name\":\"SetMigrator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"keyType\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"metadataType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldValidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValidator\",\"type\":\"address\"}],\"name\":\"SetValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"REMOVE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fidOwner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"keyType\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"metadataType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"addGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structIKeyRegistry.BulkAddKey[]\",\"name\":\"keys\",\"type\":\"tuple[]\"}],\"internalType\":\"structIKeyRegistry.BulkAddData[]\",\"name\":\"items\",\"type\":\"tuple[]\"}],\"name\":\"bulkAddKeysForMigration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"keys\",\"type\":\"bytes[]\"}],\"internalType\":\"structIKeyRegistry.BulkResetData[]\",\"name\":\"items\",\"type\":\"tuple[]\"}],\"name\":\"bulkResetKeysForMigration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparatorV4\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"freezeKeyGateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gatewayFrozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gracePeriod\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"guardians\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isGuardian\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"structHash\",\"type\":\"bytes32\"}],\"name\":\"hashTypedDataV4\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"idRegistry\",\"outputs\":[{\"internalType\":\"contractIdRegistryLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMigrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"internalType\":\"enumIKeyRegistry.KeyState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"keyAt\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"keyDataOf\",\"outputs\":[{\"components\":[{\"internalType\":\"enumIKeyRegistry.KeyState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"keyType\",\"type\":\"uint32\"}],\"internalType\":\"structIKeyRegistry.KeyData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keyGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"keys\",\"outputs\":[{\"internalType\":\"enumIKeyRegistry.KeyState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"keyType\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"internalType\":\"enumIKeyRegistry.KeyState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"keysOf\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"internalType\":\"enumIKeyRegistry.KeyState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"}],\"name\":\"keysOf\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"page\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"nextIdx\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxKeysPerFid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migratedAt\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fidOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"removeFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"removeGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_idRegistry\",\"type\":\"address\"}],\"name\":\"setIdRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_keyGateway\",\"type\":\"address\"}],\"name\":\"setKeyGateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxKeysPerFid\",\"type\":\"uint256\"}],\"name\":\"setMaxKeysPerFid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_migrator\",\"type\":\"address\"}],\"name\":\"setMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"keyType\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"metadataType\",\"type\":\"uint8\"},{\"internalType\":\"contractIMetadataValidator\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"internalType\":\"enumIKeyRegistry.KeyState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"totalKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"keyType\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"metadataType\",\"type\":\"uint8\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"contractIMetadataValidator\",\"name\":\"validator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// KeyRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use KeyRegistryMetaData.ABI instead.
var KeyRegistryABI = KeyRegistryMetaData.ABI

// KeyRegistry is an auto generated Go binding around an Ethereum contract.
type KeyRegistry struct {
	KeyRegistryCaller     // Read-only binding to the contract
	KeyRegistryTransactor // Write-only binding to the contract
	KeyRegistryFilterer   // Log filterer for contract events
}

// KeyRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeyRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeyRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KeyRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeyRegistrySession struct {
	Contract     *KeyRegistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KeyRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeyRegistryCallerSession struct {
	Contract *KeyRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// KeyRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeyRegistryTransactorSession struct {
	Contract     *KeyRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// KeyRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeyRegistryRaw struct {
	Contract *KeyRegistry // Generic contract binding to access the raw methods on
}

// KeyRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeyRegistryCallerRaw struct {
	Contract *KeyRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// KeyRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeyRegistryTransactorRaw struct {
	Contract *KeyRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeyRegistry creates a new instance of KeyRegistry, bound to a specific deployed contract.
func NewKeyRegistry(address common.Address, backend bind.ContractBackend) (*KeyRegistry, error) {
	contract, err := bindKeyRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KeyRegistry{KeyRegistryCaller: KeyRegistryCaller{contract: contract}, KeyRegistryTransactor: KeyRegistryTransactor{contract: contract}, KeyRegistryFilterer: KeyRegistryFilterer{contract: contract}}, nil
}

// NewKeyRegistryCaller creates a new read-only instance of KeyRegistry, bound to a specific deployed contract.
func NewKeyRegistryCaller(address common.Address, caller bind.ContractCaller) (*KeyRegistryCaller, error) {
	contract, err := bindKeyRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeyRegistryCaller{contract: contract}, nil
}

// NewKeyRegistryTransactor creates a new write-only instance of KeyRegistry, bound to a specific deployed contract.
func NewKeyRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*KeyRegistryTransactor, error) {
	contract, err := bindKeyRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeyRegistryTransactor{contract: contract}, nil
}

// NewKeyRegistryFilterer creates a new log filterer instance of KeyRegistry, bound to a specific deployed contract.
func NewKeyRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*KeyRegistryFilterer, error) {
	contract, err := bindKeyRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeyRegistryFilterer{contract: contract}, nil
}

// bindKeyRegistry binds a generic wrapper to an already deployed contract.
func bindKeyRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KeyRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyRegistry *KeyRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyRegistry.Contract.KeyRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyRegistry *KeyRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyRegistry.Contract.KeyRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyRegistry *KeyRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyRegistry.Contract.KeyRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyRegistry *KeyRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyRegistry *KeyRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyRegistry *KeyRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyRegistry.Contract.contract.Transact(opts, method, params...)
}

// REMOVETYPEHASH is a free data retrieval call binding the contract method 0xb5775561.
//
// Solidity: function REMOVE_TYPEHASH() view returns(bytes32)
func (_KeyRegistry *KeyRegistryCaller) REMOVETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "REMOVE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REMOVETYPEHASH is a free data retrieval call binding the contract method 0xb5775561.
//
// Solidity: function REMOVE_TYPEHASH() view returns(bytes32)
func (_KeyRegistry *KeyRegistrySession) REMOVETYPEHASH() ([32]byte, error) {
	return _KeyRegistry.Contract.REMOVETYPEHASH(&_KeyRegistry.CallOpts)
}

// REMOVETYPEHASH is a free data retrieval call binding the contract method 0xb5775561.
//
// Solidity: function REMOVE_TYPEHASH() view returns(bytes32)
func (_KeyRegistry *KeyRegistryCallerSession) REMOVETYPEHASH() ([32]byte, error) {
	return _KeyRegistry.Contract.REMOVETYPEHASH(&_KeyRegistry.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_KeyRegistry *KeyRegistryCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_KeyRegistry *KeyRegistrySession) VERSION() (string, error) {
	return _KeyRegistry.Contract.VERSION(&_KeyRegistry.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_KeyRegistry *KeyRegistryCallerSession) VERSION() (string, error) {
	return _KeyRegistry.Contract.VERSION(&_KeyRegistry.CallOpts)
}

// DomainSeparatorV4 is a free data retrieval call binding the contract method 0x78e890ba.
//
// Solidity: function domainSeparatorV4() view returns(bytes32)
func (_KeyRegistry *KeyRegistryCaller) DomainSeparatorV4(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "domainSeparatorV4")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparatorV4 is a free data retrieval call binding the contract method 0x78e890ba.
//
// Solidity: function domainSeparatorV4() view returns(bytes32)
func (_KeyRegistry *KeyRegistrySession) DomainSeparatorV4() ([32]byte, error) {
	return _KeyRegistry.Contract.DomainSeparatorV4(&_KeyRegistry.CallOpts)
}

// DomainSeparatorV4 is a free data retrieval call binding the contract method 0x78e890ba.
//
// Solidity: function domainSeparatorV4() view returns(bytes32)
func (_KeyRegistry *KeyRegistryCallerSession) DomainSeparatorV4() ([32]byte, error) {
	return _KeyRegistry.Contract.DomainSeparatorV4(&_KeyRegistry.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_KeyRegistry *KeyRegistryCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "eip712Domain")

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
func (_KeyRegistry *KeyRegistrySession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _KeyRegistry.Contract.Eip712Domain(&_KeyRegistry.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_KeyRegistry *KeyRegistryCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _KeyRegistry.Contract.Eip712Domain(&_KeyRegistry.CallOpts)
}

// GatewayFrozen is a free data retrieval call binding the contract method 0x95e7549f.
//
// Solidity: function gatewayFrozen() view returns(bool)
func (_KeyRegistry *KeyRegistryCaller) GatewayFrozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "gatewayFrozen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GatewayFrozen is a free data retrieval call binding the contract method 0x95e7549f.
//
// Solidity: function gatewayFrozen() view returns(bool)
func (_KeyRegistry *KeyRegistrySession) GatewayFrozen() (bool, error) {
	return _KeyRegistry.Contract.GatewayFrozen(&_KeyRegistry.CallOpts)
}

// GatewayFrozen is a free data retrieval call binding the contract method 0x95e7549f.
//
// Solidity: function gatewayFrozen() view returns(bool)
func (_KeyRegistry *KeyRegistryCallerSession) GatewayFrozen() (bool, error) {
	return _KeyRegistry.Contract.GatewayFrozen(&_KeyRegistry.CallOpts)
}

// GracePeriod is a free data retrieval call binding the contract method 0xa06db7dc.
//
// Solidity: function gracePeriod() view returns(uint24)
func (_KeyRegistry *KeyRegistryCaller) GracePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "gracePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GracePeriod is a free data retrieval call binding the contract method 0xa06db7dc.
//
// Solidity: function gracePeriod() view returns(uint24)
func (_KeyRegistry *KeyRegistrySession) GracePeriod() (*big.Int, error) {
	return _KeyRegistry.Contract.GracePeriod(&_KeyRegistry.CallOpts)
}

// GracePeriod is a free data retrieval call binding the contract method 0xa06db7dc.
//
// Solidity: function gracePeriod() view returns(uint24)
func (_KeyRegistry *KeyRegistryCallerSession) GracePeriod() (*big.Int, error) {
	return _KeyRegistry.Contract.GracePeriod(&_KeyRegistry.CallOpts)
}

// Guardians is a free data retrieval call binding the contract method 0x0633b14a.
//
// Solidity: function guardians(address guardian) view returns(bool isGuardian)
func (_KeyRegistry *KeyRegistryCaller) Guardians(opts *bind.CallOpts, guardian common.Address) (bool, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "guardians", guardian)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Guardians is a free data retrieval call binding the contract method 0x0633b14a.
//
// Solidity: function guardians(address guardian) view returns(bool isGuardian)
func (_KeyRegistry *KeyRegistrySession) Guardians(guardian common.Address) (bool, error) {
	return _KeyRegistry.Contract.Guardians(&_KeyRegistry.CallOpts, guardian)
}

// Guardians is a free data retrieval call binding the contract method 0x0633b14a.
//
// Solidity: function guardians(address guardian) view returns(bool isGuardian)
func (_KeyRegistry *KeyRegistryCallerSession) Guardians(guardian common.Address) (bool, error) {
	return _KeyRegistry.Contract.Guardians(&_KeyRegistry.CallOpts, guardian)
}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x4980f288.
//
// Solidity: function hashTypedDataV4(bytes32 structHash) view returns(bytes32)
func (_KeyRegistry *KeyRegistryCaller) HashTypedDataV4(opts *bind.CallOpts, structHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "hashTypedDataV4", structHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x4980f288.
//
// Solidity: function hashTypedDataV4(bytes32 structHash) view returns(bytes32)
func (_KeyRegistry *KeyRegistrySession) HashTypedDataV4(structHash [32]byte) ([32]byte, error) {
	return _KeyRegistry.Contract.HashTypedDataV4(&_KeyRegistry.CallOpts, structHash)
}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x4980f288.
//
// Solidity: function hashTypedDataV4(bytes32 structHash) view returns(bytes32)
func (_KeyRegistry *KeyRegistryCallerSession) HashTypedDataV4(structHash [32]byte) ([32]byte, error) {
	return _KeyRegistry.Contract.HashTypedDataV4(&_KeyRegistry.CallOpts, structHash)
}

// IdRegistry is a free data retrieval call binding the contract method 0x0aa13b8c.
//
// Solidity: function idRegistry() view returns(address)
func (_KeyRegistry *KeyRegistryCaller) IdRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "idRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdRegistry is a free data retrieval call binding the contract method 0x0aa13b8c.
//
// Solidity: function idRegistry() view returns(address)
func (_KeyRegistry *KeyRegistrySession) IdRegistry() (common.Address, error) {
	return _KeyRegistry.Contract.IdRegistry(&_KeyRegistry.CallOpts)
}

// IdRegistry is a free data retrieval call binding the contract method 0x0aa13b8c.
//
// Solidity: function idRegistry() view returns(address)
func (_KeyRegistry *KeyRegistryCallerSession) IdRegistry() (common.Address, error) {
	return _KeyRegistry.Contract.IdRegistry(&_KeyRegistry.CallOpts)
}

// IsMigrated is a free data retrieval call binding the contract method 0xb06faf62.
//
// Solidity: function isMigrated() view returns(bool)
func (_KeyRegistry *KeyRegistryCaller) IsMigrated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "isMigrated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMigrated is a free data retrieval call binding the contract method 0xb06faf62.
//
// Solidity: function isMigrated() view returns(bool)
func (_KeyRegistry *KeyRegistrySession) IsMigrated() (bool, error) {
	return _KeyRegistry.Contract.IsMigrated(&_KeyRegistry.CallOpts)
}

// IsMigrated is a free data retrieval call binding the contract method 0xb06faf62.
//
// Solidity: function isMigrated() view returns(bool)
func (_KeyRegistry *KeyRegistryCallerSession) IsMigrated() (bool, error) {
	return _KeyRegistry.Contract.IsMigrated(&_KeyRegistry.CallOpts)
}

// KeyAt is a free data retrieval call binding the contract method 0x0ea9442c.
//
// Solidity: function keyAt(uint256 fid, uint8 state, uint256 index) view returns(bytes)
func (_KeyRegistry *KeyRegistryCaller) KeyAt(opts *bind.CallOpts, fid *big.Int, state uint8, index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "keyAt", fid, state, index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// KeyAt is a free data retrieval call binding the contract method 0x0ea9442c.
//
// Solidity: function keyAt(uint256 fid, uint8 state, uint256 index) view returns(bytes)
func (_KeyRegistry *KeyRegistrySession) KeyAt(fid *big.Int, state uint8, index *big.Int) ([]byte, error) {
	return _KeyRegistry.Contract.KeyAt(&_KeyRegistry.CallOpts, fid, state, index)
}

// KeyAt is a free data retrieval call binding the contract method 0x0ea9442c.
//
// Solidity: function keyAt(uint256 fid, uint8 state, uint256 index) view returns(bytes)
func (_KeyRegistry *KeyRegistryCallerSession) KeyAt(fid *big.Int, state uint8, index *big.Int) ([]byte, error) {
	return _KeyRegistry.Contract.KeyAt(&_KeyRegistry.CallOpts, fid, state, index)
}

// KeyDataOf is a free data retrieval call binding the contract method 0xac34cc5a.
//
// Solidity: function keyDataOf(uint256 fid, bytes key) view returns((uint8,uint32))
func (_KeyRegistry *KeyRegistryCaller) KeyDataOf(opts *bind.CallOpts, fid *big.Int, key []byte) (IKeyRegistryKeyData, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "keyDataOf", fid, key)

	if err != nil {
		return *new(IKeyRegistryKeyData), err
	}

	out0 := *abi.ConvertType(out[0], new(IKeyRegistryKeyData)).(*IKeyRegistryKeyData)

	return out0, err

}

// KeyDataOf is a free data retrieval call binding the contract method 0xac34cc5a.
//
// Solidity: function keyDataOf(uint256 fid, bytes key) view returns((uint8,uint32))
func (_KeyRegistry *KeyRegistrySession) KeyDataOf(fid *big.Int, key []byte) (IKeyRegistryKeyData, error) {
	return _KeyRegistry.Contract.KeyDataOf(&_KeyRegistry.CallOpts, fid, key)
}

// KeyDataOf is a free data retrieval call binding the contract method 0xac34cc5a.
//
// Solidity: function keyDataOf(uint256 fid, bytes key) view returns((uint8,uint32))
func (_KeyRegistry *KeyRegistryCallerSession) KeyDataOf(fid *big.Int, key []byte) (IKeyRegistryKeyData, error) {
	return _KeyRegistry.Contract.KeyDataOf(&_KeyRegistry.CallOpts, fid, key)
}

// KeyGateway is a free data retrieval call binding the contract method 0x80334737.
//
// Solidity: function keyGateway() view returns(address)
func (_KeyRegistry *KeyRegistryCaller) KeyGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "keyGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeyGateway is a free data retrieval call binding the contract method 0x80334737.
//
// Solidity: function keyGateway() view returns(address)
func (_KeyRegistry *KeyRegistrySession) KeyGateway() (common.Address, error) {
	return _KeyRegistry.Contract.KeyGateway(&_KeyRegistry.CallOpts)
}

// KeyGateway is a free data retrieval call binding the contract method 0x80334737.
//
// Solidity: function keyGateway() view returns(address)
func (_KeyRegistry *KeyRegistryCallerSession) KeyGateway() (common.Address, error) {
	return _KeyRegistry.Contract.KeyGateway(&_KeyRegistry.CallOpts)
}

// Keys is a free data retrieval call binding the contract method 0x0e24e34c.
//
// Solidity: function keys(uint256 fid, bytes key) view returns(uint8 state, uint32 keyType)
func (_KeyRegistry *KeyRegistryCaller) Keys(opts *bind.CallOpts, fid *big.Int, key []byte) (struct {
	State   uint8
	KeyType uint32
}, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "keys", fid, key)

	outstruct := new(struct {
		State   uint8
		KeyType uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.State = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.KeyType = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// Keys is a free data retrieval call binding the contract method 0x0e24e34c.
//
// Solidity: function keys(uint256 fid, bytes key) view returns(uint8 state, uint32 keyType)
func (_KeyRegistry *KeyRegistrySession) Keys(fid *big.Int, key []byte) (struct {
	State   uint8
	KeyType uint32
}, error) {
	return _KeyRegistry.Contract.Keys(&_KeyRegistry.CallOpts, fid, key)
}

// Keys is a free data retrieval call binding the contract method 0x0e24e34c.
//
// Solidity: function keys(uint256 fid, bytes key) view returns(uint8 state, uint32 keyType)
func (_KeyRegistry *KeyRegistryCallerSession) Keys(fid *big.Int, key []byte) (struct {
	State   uint8
	KeyType uint32
}, error) {
	return _KeyRegistry.Contract.Keys(&_KeyRegistry.CallOpts, fid, key)
}

// KeysOf is a free data retrieval call binding the contract method 0x1f64222f.
//
// Solidity: function keysOf(uint256 fid, uint8 state) view returns(bytes[])
func (_KeyRegistry *KeyRegistryCaller) KeysOf(opts *bind.CallOpts, fid *big.Int, state uint8) ([][]byte, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "keysOf", fid, state)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// KeysOf is a free data retrieval call binding the contract method 0x1f64222f.
//
// Solidity: function keysOf(uint256 fid, uint8 state) view returns(bytes[])
func (_KeyRegistry *KeyRegistrySession) KeysOf(fid *big.Int, state uint8) ([][]byte, error) {
	return _KeyRegistry.Contract.KeysOf(&_KeyRegistry.CallOpts, fid, state)
}

// KeysOf is a free data retrieval call binding the contract method 0x1f64222f.
//
// Solidity: function keysOf(uint256 fid, uint8 state) view returns(bytes[])
func (_KeyRegistry *KeyRegistryCallerSession) KeysOf(fid *big.Int, state uint8) ([][]byte, error) {
	return _KeyRegistry.Contract.KeysOf(&_KeyRegistry.CallOpts, fid, state)
}

// KeysOf0 is a free data retrieval call binding the contract method 0xf27995e3.
//
// Solidity: function keysOf(uint256 fid, uint8 state, uint256 startIdx, uint256 batchSize) view returns(bytes[] page, uint256 nextIdx)
func (_KeyRegistry *KeyRegistryCaller) KeysOf0(opts *bind.CallOpts, fid *big.Int, state uint8, startIdx *big.Int, batchSize *big.Int) (struct {
	Page    [][]byte
	NextIdx *big.Int
}, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "keysOf0", fid, state, startIdx, batchSize)

	outstruct := new(struct {
		Page    [][]byte
		NextIdx *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Page = *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)
	outstruct.NextIdx = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// KeysOf0 is a free data retrieval call binding the contract method 0xf27995e3.
//
// Solidity: function keysOf(uint256 fid, uint8 state, uint256 startIdx, uint256 batchSize) view returns(bytes[] page, uint256 nextIdx)
func (_KeyRegistry *KeyRegistrySession) KeysOf0(fid *big.Int, state uint8, startIdx *big.Int, batchSize *big.Int) (struct {
	Page    [][]byte
	NextIdx *big.Int
}, error) {
	return _KeyRegistry.Contract.KeysOf0(&_KeyRegistry.CallOpts, fid, state, startIdx, batchSize)
}

// KeysOf0 is a free data retrieval call binding the contract method 0xf27995e3.
//
// Solidity: function keysOf(uint256 fid, uint8 state, uint256 startIdx, uint256 batchSize) view returns(bytes[] page, uint256 nextIdx)
func (_KeyRegistry *KeyRegistryCallerSession) KeysOf0(fid *big.Int, state uint8, startIdx *big.Int, batchSize *big.Int) (struct {
	Page    [][]byte
	NextIdx *big.Int
}, error) {
	return _KeyRegistry.Contract.KeysOf0(&_KeyRegistry.CallOpts, fid, state, startIdx, batchSize)
}

// MaxKeysPerFid is a free data retrieval call binding the contract method 0xe33acf38.
//
// Solidity: function maxKeysPerFid() view returns(uint256)
func (_KeyRegistry *KeyRegistryCaller) MaxKeysPerFid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "maxKeysPerFid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxKeysPerFid is a free data retrieval call binding the contract method 0xe33acf38.
//
// Solidity: function maxKeysPerFid() view returns(uint256)
func (_KeyRegistry *KeyRegistrySession) MaxKeysPerFid() (*big.Int, error) {
	return _KeyRegistry.Contract.MaxKeysPerFid(&_KeyRegistry.CallOpts)
}

// MaxKeysPerFid is a free data retrieval call binding the contract method 0xe33acf38.
//
// Solidity: function maxKeysPerFid() view returns(uint256)
func (_KeyRegistry *KeyRegistryCallerSession) MaxKeysPerFid() (*big.Int, error) {
	return _KeyRegistry.Contract.MaxKeysPerFid(&_KeyRegistry.CallOpts)
}

// MigratedAt is a free data retrieval call binding the contract method 0x8b21e484.
//
// Solidity: function migratedAt() view returns(uint40)
func (_KeyRegistry *KeyRegistryCaller) MigratedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "migratedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MigratedAt is a free data retrieval call binding the contract method 0x8b21e484.
//
// Solidity: function migratedAt() view returns(uint40)
func (_KeyRegistry *KeyRegistrySession) MigratedAt() (*big.Int, error) {
	return _KeyRegistry.Contract.MigratedAt(&_KeyRegistry.CallOpts)
}

// MigratedAt is a free data retrieval call binding the contract method 0x8b21e484.
//
// Solidity: function migratedAt() view returns(uint40)
func (_KeyRegistry *KeyRegistryCallerSession) MigratedAt() (*big.Int, error) {
	return _KeyRegistry.Contract.MigratedAt(&_KeyRegistry.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_KeyRegistry *KeyRegistryCaller) Migrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "migrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_KeyRegistry *KeyRegistrySession) Migrator() (common.Address, error) {
	return _KeyRegistry.Contract.Migrator(&_KeyRegistry.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_KeyRegistry *KeyRegistryCallerSession) Migrator() (common.Address, error) {
	return _KeyRegistry.Contract.Migrator(&_KeyRegistry.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_KeyRegistry *KeyRegistryCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_KeyRegistry *KeyRegistrySession) Nonces(owner common.Address) (*big.Int, error) {
	return _KeyRegistry.Contract.Nonces(&_KeyRegistry.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_KeyRegistry *KeyRegistryCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _KeyRegistry.Contract.Nonces(&_KeyRegistry.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KeyRegistry *KeyRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KeyRegistry *KeyRegistrySession) Owner() (common.Address, error) {
	return _KeyRegistry.Contract.Owner(&_KeyRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KeyRegistry *KeyRegistryCallerSession) Owner() (common.Address, error) {
	return _KeyRegistry.Contract.Owner(&_KeyRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KeyRegistry *KeyRegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KeyRegistry *KeyRegistrySession) Paused() (bool, error) {
	return _KeyRegistry.Contract.Paused(&_KeyRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KeyRegistry *KeyRegistryCallerSession) Paused() (bool, error) {
	return _KeyRegistry.Contract.Paused(&_KeyRegistry.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_KeyRegistry *KeyRegistryCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_KeyRegistry *KeyRegistrySession) PendingOwner() (common.Address, error) {
	return _KeyRegistry.Contract.PendingOwner(&_KeyRegistry.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_KeyRegistry *KeyRegistryCallerSession) PendingOwner() (common.Address, error) {
	return _KeyRegistry.Contract.PendingOwner(&_KeyRegistry.CallOpts)
}

// TotalKeys is a free data retrieval call binding the contract method 0x6840b75e.
//
// Solidity: function totalKeys(uint256 fid, uint8 state) view returns(uint256)
func (_KeyRegistry *KeyRegistryCaller) TotalKeys(opts *bind.CallOpts, fid *big.Int, state uint8) (*big.Int, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "totalKeys", fid, state)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalKeys is a free data retrieval call binding the contract method 0x6840b75e.
//
// Solidity: function totalKeys(uint256 fid, uint8 state) view returns(uint256)
func (_KeyRegistry *KeyRegistrySession) TotalKeys(fid *big.Int, state uint8) (*big.Int, error) {
	return _KeyRegistry.Contract.TotalKeys(&_KeyRegistry.CallOpts, fid, state)
}

// TotalKeys is a free data retrieval call binding the contract method 0x6840b75e.
//
// Solidity: function totalKeys(uint256 fid, uint8 state) view returns(uint256)
func (_KeyRegistry *KeyRegistryCallerSession) TotalKeys(fid *big.Int, state uint8) (*big.Int, error) {
	return _KeyRegistry.Contract.TotalKeys(&_KeyRegistry.CallOpts, fid, state)
}

// Validators is a free data retrieval call binding the contract method 0xd8810395.
//
// Solidity: function validators(uint32 keyType, uint8 metadataType) view returns(address validator)
func (_KeyRegistry *KeyRegistryCaller) Validators(opts *bind.CallOpts, keyType uint32, metadataType uint8) (common.Address, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "validators", keyType, metadataType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Validators is a free data retrieval call binding the contract method 0xd8810395.
//
// Solidity: function validators(uint32 keyType, uint8 metadataType) view returns(address validator)
func (_KeyRegistry *KeyRegistrySession) Validators(keyType uint32, metadataType uint8) (common.Address, error) {
	return _KeyRegistry.Contract.Validators(&_KeyRegistry.CallOpts, keyType, metadataType)
}

// Validators is a free data retrieval call binding the contract method 0xd8810395.
//
// Solidity: function validators(uint32 keyType, uint8 metadataType) view returns(address validator)
func (_KeyRegistry *KeyRegistryCallerSession) Validators(keyType uint32, metadataType uint8) (common.Address, error) {
	return _KeyRegistry.Contract.Validators(&_KeyRegistry.CallOpts, keyType, metadataType)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_KeyRegistry *KeyRegistryTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyRegistry.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_KeyRegistry *KeyRegistrySession) AcceptOwnership() (*types.Transaction, error) {
	return _KeyRegistry.Contract.AcceptOwnership(&_KeyRegistry.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_KeyRegistry *KeyRegistryTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _KeyRegistry.Contract.AcceptOwnership(&_KeyRegistry.TransactOpts)
}

// Add is a paid mutator transaction binding the contract method 0x207e3d69.
//
// Solidity: function add(address fidOwner, uint32 keyType, bytes key, uint8 metadataType, bytes metadata) returns()
func (_KeyRegistry *KeyRegistryTransactor) Add(opts *bind.TransactOpts, fidOwner common.Address, keyType uint32, key []byte, metadataType uint8, metadata []byte) (*types.Transaction, error) {
	return _KeyRegistry.contract.Transact(opts, "add", fidOwner, keyType, key, metadataType, metadata)
}

// Add is a paid mutator transaction binding the contract method 0x207e3d69.
//
// Solidity: function add(address fidOwner, uint32 keyType, bytes key, uint8 metadataType, bytes metadata) returns()
func (_KeyRegistry *KeyRegistrySession) Add(fidOwner common.Address, keyType uint32, key []byte, metadataType uint8, metadata []byte) (*types.Transaction, error) {
	return _KeyRegistry.Contract.Add(&_KeyRegistry.TransactOpts, fidOwner, keyType, key, metadataType, metadata)
}

// Add is a paid mutator transaction binding the contract method 0x207e3d69.
//
// Solidity: function add(address fidOwner, uint32 keyType, bytes key, uint8 metadataType, bytes metadata) returns()
func (_KeyRegistry *KeyRegistryTransactorSession) Add(fidOwner common.Address, keyType uint32, key []byte, metadataType uint8, metadata []byte) (*types.Transaction, error) {
	return _KeyRegistry.Contract.Add(&_KeyRegistry.TransactOpts, fidOwner, keyType, key, metadataType, metadata)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_KeyRegistry *KeyRegistryTransactor) AddGuardian(opts *bind.TransactOpts, guardian common.Address) (*types.Transaction, error) {
	return _KeyRegistry.contract.Transact(opts, "addGuardian", guardian)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_KeyRegistry *KeyRegistrySession) AddGuardian(guardian common.Address) (*types.Transaction, error) {
	return _KeyRegistry.Contract.AddGuardian(&_KeyRegistry.TransactOpts, guardian)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_KeyRegistry *KeyRegistryTransactorSession) AddGuardian(guardian common.Address) (*types.Transaction, error) {
	return _KeyRegistry.Contract.AddGuardian(&_KeyRegistry.TransactOpts, guardian)
}

// BulkAddKeysForMigration is a paid mutator transaction binding the contract method 0x708e9c70.
//
// Solidity: function bulkAddKeysForMigration((uint256,(bytes,bytes)[])[] items) returns()
func (_KeyRegistry *KeyRegistryTransactor) BulkAddKeysForMigration(opts *bind.TransactOpts, items []IKeyRegistryBulkAddData) (*types.Transaction, error) {
	return _KeyRegistry.contract.Transact(opts, "bulkAddKeysForMigration", items)
}

// BulkAddKeysForMigration is a paid mutator transaction binding the contract method 0x708e9c70.
//
// Solidity: function bulkAddKeysForMigration((uint256,(bytes,bytes)[])[] items) returns()
func (_KeyRegistry *KeyRegistrySession) BulkAddKeysForMigration(items []IKeyRegistryBulkAddData) (*types.Transaction, error) {
	return _KeyRegistry.Contract.BulkAddKeysForMigration(&_KeyRegistry.TransactOpts, items)
}

// BulkAddKeysForMigration is a paid mutator transaction binding the contract method 0x708e9c70.
//
// Solidity: function bulkAddKeysForMigration((uint256,(bytes,bytes)[])[] items) returns()
func (_KeyRegistry *KeyRegistryTransactorSession) BulkAddKeysForMigration(items []IKeyRegistryBulkAddData) (*types.Transaction, error) {
	return _KeyRegistry.Contract.BulkAddKeysForMigration(&_KeyRegistry.TransactOpts, items)
}

// BulkResetKeysForMigration is a paid mutator transaction binding the contract method 0x46b3f429.
//
// Solidity: function bulkResetKeysForMigration((uint256,bytes[])[] items) returns()
func (_KeyRegistry *KeyRegistryTransactor) BulkResetKeysForMigration(opts *bind.TransactOpts, items []IKeyRegistryBulkResetData) (*types.Transaction, error) {
	return _KeyRegistry.contract.Transact(opts, "bulkResetKeysForMigration", items)
}

// BulkResetKeysForMigration is a paid mutator transaction binding the contract method 0x46b3f429.
//
// Solidity: function bulkResetKeysForMigration((uint256,bytes[])[] items) returns()
func (_KeyRegistry *KeyRegistrySession) BulkResetKeysForMigration(items []IKeyRegistryBulkResetData) (*types.Transaction, error) {
	return _KeyRegistry.Contract.BulkResetKeysForMigration(&_KeyRegistry.TransactOpts, items)
}

// BulkResetKeysForMigration is a paid mutator transaction binding the contract method 0x46b3f429.
//
// Solidity: function bulkResetKeysForMigration((uint256,bytes[])[] items) returns()
func (_KeyRegistry *KeyRegistryTransactorSession) BulkResetKeysForMigration(items []IKeyRegistryBulkResetData) (*types.Transaction, error) {
	return _KeyRegistry.Contract.BulkResetKeysForMigration(&_KeyRegistry.TransactOpts, items)
}

// FreezeKeyGateway is a paid mutator transaction binding the contract method 0x47cf8d4d.
//
// Solidity: function freezeKeyGateway() returns()
func (_KeyRegistry *KeyRegistryTransactor) FreezeKeyGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyRegistry.contract.Transact(opts, "freezeKeyGateway")
}

// FreezeKeyGateway is a paid mutator transaction binding the contract method 0x47cf8d4d.
//
// Solidity: function freezeKeyGateway() returns()
func (_KeyRegistry *KeyRegistrySession) FreezeKeyGateway() (*types.Transaction, error) {
	return _KeyRegistry.Contract.FreezeKeyGateway(&_KeyRegistry.TransactOpts)
}

// FreezeKeyGateway is a paid mutator transaction binding the contract method 0x47cf8d4d.
//
// Solidity: function freezeKeyGateway() returns()
func (_KeyRegistry *KeyRegistryTransactorSession) FreezeKeyGateway() (*types.Transaction, error) {
	return _KeyRegistry.Contract.FreezeKeyGateway(&_KeyRegistry.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_KeyRegistry *KeyRegistryTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyRegistry.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_KeyRegistry *KeyRegistrySession) Migrate() (*types.Transaction, error) {
	return _KeyRegistry.Contract.Migrate(&_KeyRegistry.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_KeyRegistry *KeyRegistryTransactorSession) Migrate() (*types.Transaction, error) {
	return _KeyRegistry.Contract.Migrate(&_KeyRegistry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KeyRegistry *KeyRegistryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyRegistry.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KeyRegistry *KeyRegistrySession) Pause() (*types.Transaction, error) {
	return _KeyRegistry.Contract.Pause(&_KeyRegistry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KeyRegistry *KeyRegistryTransactorSession) Pause() (*types.Transaction, error) {
	return _KeyRegistry.Contract.Pause(&_KeyRegistry.TransactOpts)
}

// Remove is a paid mutator transaction binding the contract method 0x58edef4c.
//
// Solidity: function remove(bytes key) returns()
func (_KeyRegistry *KeyRegistryTransactor) Remove(opts *bind.TransactOpts, key []byte) (*types.Transaction, error) {
	return _KeyRegistry.contract.Transact(opts, "remove", key)
}

// Remove is a paid mutator transaction binding the contract method 0x58edef4c.
//
// Solidity: function remove(bytes key) returns()
func (_KeyRegistry *KeyRegistrySession) Remove(key []byte) (*types.Transaction, error) {
	return _KeyRegistry.Contract.Remove(&_KeyRegistry.TransactOpts, key)
}

// Remove is a paid mutator transaction binding the contract method 0x58edef4c.
//
// Solidity: function remove(bytes key) returns()
func (_KeyRegistry *KeyRegistryTransactorSession) Remove(key []byte) (*types.Transaction, error) {
	return _KeyRegistry.Contract.Remove(&_KeyRegistry.TransactOpts, key)
}

// RemoveFor is a paid mutator transaction binding the contract method 0x787bd966.
//
// Solidity: function removeFor(address fidOwner, bytes key, uint256 deadline, bytes sig) returns()
func (_KeyRegistry *KeyRegistryTransactor) RemoveFor(opts *bind.TransactOpts, fidOwner common.Address, key []byte, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _KeyRegistry.contract.Transact(opts, "removeFor", fidOwner, key, deadline, sig)
}

// RemoveFor is a paid mutator transaction binding the contract method 0x787bd966.
//
// Solidity: function removeFor(address fidOwner, bytes key, uint256 deadline, bytes sig) returns()
func (_KeyRegistry *KeyRegistrySession) RemoveFor(fidOwner common.Address, key []byte, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _KeyRegistry.Contract.RemoveFor(&_KeyRegistry.TransactOpts, fidOwner, key, deadline, sig)
}

// RemoveFor is a paid mutator transaction binding the contract method 0x787bd966.
//
// Solidity: function removeFor(address fidOwner, bytes key, uint256 deadline, bytes sig) returns()
func (_KeyRegistry *KeyRegistryTransactorSession) RemoveFor(fidOwner common.Address, key []byte, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _KeyRegistry.Contract.RemoveFor(&_KeyRegistry.TransactOpts, fidOwner, key, deadline, sig)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_KeyRegistry *KeyRegistryTransactor) RemoveGuardian(opts *bind.TransactOpts, guardian common.Address) (*types.Transaction, error) {
	return _KeyRegistry.contract.Transact(opts, "removeGuardian", guardian)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_KeyRegistry *KeyRegistrySession) RemoveGuardian(guardian common.Address) (*types.Transaction, error) {
	return _KeyRegistry.Contract.RemoveGuardian(&_KeyRegistry.TransactOpts, guardian)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_KeyRegistry *KeyRegistryTransactorSession) RemoveGuardian(guardian common.Address) (*types.Transaction, error) {
	return _KeyRegistry.Contract.RemoveGuardian(&_KeyRegistry.TransactOpts, guardian)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KeyRegistry *KeyRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KeyRegistry *KeyRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _KeyRegistry.Contract.RenounceOwnership(&_KeyRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KeyRegistry *KeyRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _KeyRegistry.Contract.RenounceOwnership(&_KeyRegistry.TransactOpts)
}

// SetIdRegistry is a paid mutator transaction binding the contract method 0x81749468.
//
// Solidity: function setIdRegistry(address _idRegistry) returns()
func (_KeyRegistry *KeyRegistryTransactor) SetIdRegistry(opts *bind.TransactOpts, _idRegistry common.Address) (*types.Transaction, error) {
	return _KeyRegistry.contract.Transact(opts, "setIdRegistry", _idRegistry)
}

// SetIdRegistry is a paid mutator transaction binding the contract method 0x81749468.
//
// Solidity: function setIdRegistry(address _idRegistry) returns()
func (_KeyRegistry *KeyRegistrySession) SetIdRegistry(_idRegistry common.Address) (*types.Transaction, error) {
	return _KeyRegistry.Contract.SetIdRegistry(&_KeyRegistry.TransactOpts, _idRegistry)
}

// SetIdRegistry is a paid mutator transaction binding the contract method 0x81749468.
//
// Solidity: function setIdRegistry(address _idRegistry) returns()
func (_KeyRegistry *KeyRegistryTransactorSession) SetIdRegistry(_idRegistry common.Address) (*types.Transaction, error) {
	return _KeyRegistry.Contract.SetIdRegistry(&_KeyRegistry.TransactOpts, _idRegistry)
}

// SetKeyGateway is a paid mutator transaction binding the contract method 0xb221dac4.
//
// Solidity: function setKeyGateway(address _keyGateway) returns()
func (_KeyRegistry *KeyRegistryTransactor) SetKeyGateway(opts *bind.TransactOpts, _keyGateway common.Address) (*types.Transaction, error) {
	return _KeyRegistry.contract.Transact(opts, "setKeyGateway", _keyGateway)
}

// SetKeyGateway is a paid mutator transaction binding the contract method 0xb221dac4.
//
// Solidity: function setKeyGateway(address _keyGateway) returns()
func (_KeyRegistry *KeyRegistrySession) SetKeyGateway(_keyGateway common.Address) (*types.Transaction, error) {
	return _KeyRegistry.Contract.SetKeyGateway(&_KeyRegistry.TransactOpts, _keyGateway)
}

// SetKeyGateway is a paid mutator transaction binding the contract method 0xb221dac4.
//
// Solidity: function setKeyGateway(address _keyGateway) returns()
func (_KeyRegistry *KeyRegistryTransactorSession) SetKeyGateway(_keyGateway common.Address) (*types.Transaction, error) {
	return _KeyRegistry.Contract.SetKeyGateway(&_KeyRegistry.TransactOpts, _keyGateway)
}

// SetMaxKeysPerFid is a paid mutator transaction binding the contract method 0xd4c04809.
//
// Solidity: function setMaxKeysPerFid(uint256 _maxKeysPerFid) returns()
func (_KeyRegistry *KeyRegistryTransactor) SetMaxKeysPerFid(opts *bind.TransactOpts, _maxKeysPerFid *big.Int) (*types.Transaction, error) {
	return _KeyRegistry.contract.Transact(opts, "setMaxKeysPerFid", _maxKeysPerFid)
}

// SetMaxKeysPerFid is a paid mutator transaction binding the contract method 0xd4c04809.
//
// Solidity: function setMaxKeysPerFid(uint256 _maxKeysPerFid) returns()
func (_KeyRegistry *KeyRegistrySession) SetMaxKeysPerFid(_maxKeysPerFid *big.Int) (*types.Transaction, error) {
	return _KeyRegistry.Contract.SetMaxKeysPerFid(&_KeyRegistry.TransactOpts, _maxKeysPerFid)
}

// SetMaxKeysPerFid is a paid mutator transaction binding the contract method 0xd4c04809.
//
// Solidity: function setMaxKeysPerFid(uint256 _maxKeysPerFid) returns()
func (_KeyRegistry *KeyRegistryTransactorSession) SetMaxKeysPerFid(_maxKeysPerFid *big.Int) (*types.Transaction, error) {
	return _KeyRegistry.Contract.SetMaxKeysPerFid(&_KeyRegistry.TransactOpts, _maxKeysPerFid)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_KeyRegistry *KeyRegistryTransactor) SetMigrator(opts *bind.TransactOpts, _migrator common.Address) (*types.Transaction, error) {
	return _KeyRegistry.contract.Transact(opts, "setMigrator", _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_KeyRegistry *KeyRegistrySession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _KeyRegistry.Contract.SetMigrator(&_KeyRegistry.TransactOpts, _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_KeyRegistry *KeyRegistryTransactorSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _KeyRegistry.Contract.SetMigrator(&_KeyRegistry.TransactOpts, _migrator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x368ab610.
//
// Solidity: function setValidator(uint32 keyType, uint8 metadataType, address validator) returns()
func (_KeyRegistry *KeyRegistryTransactor) SetValidator(opts *bind.TransactOpts, keyType uint32, metadataType uint8, validator common.Address) (*types.Transaction, error) {
	return _KeyRegistry.contract.Transact(opts, "setValidator", keyType, metadataType, validator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x368ab610.
//
// Solidity: function setValidator(uint32 keyType, uint8 metadataType, address validator) returns()
func (_KeyRegistry *KeyRegistrySession) SetValidator(keyType uint32, metadataType uint8, validator common.Address) (*types.Transaction, error) {
	return _KeyRegistry.Contract.SetValidator(&_KeyRegistry.TransactOpts, keyType, metadataType, validator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x368ab610.
//
// Solidity: function setValidator(uint32 keyType, uint8 metadataType, address validator) returns()
func (_KeyRegistry *KeyRegistryTransactorSession) SetValidator(keyType uint32, metadataType uint8, validator common.Address) (*types.Transaction, error) {
	return _KeyRegistry.Contract.SetValidator(&_KeyRegistry.TransactOpts, keyType, metadataType, validator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KeyRegistry *KeyRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KeyRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KeyRegistry *KeyRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KeyRegistry.Contract.TransferOwnership(&_KeyRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KeyRegistry *KeyRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KeyRegistry.Contract.TransferOwnership(&_KeyRegistry.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KeyRegistry *KeyRegistryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyRegistry.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KeyRegistry *KeyRegistrySession) Unpause() (*types.Transaction, error) {
	return _KeyRegistry.Contract.Unpause(&_KeyRegistry.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KeyRegistry *KeyRegistryTransactorSession) Unpause() (*types.Transaction, error) {
	return _KeyRegistry.Contract.Unpause(&_KeyRegistry.TransactOpts)
}

// UseNonce is a paid mutator transaction binding the contract method 0x69615a4c.
//
// Solidity: function useNonce() returns(uint256)
func (_KeyRegistry *KeyRegistryTransactor) UseNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyRegistry.contract.Transact(opts, "useNonce")
}

// UseNonce is a paid mutator transaction binding the contract method 0x69615a4c.
//
// Solidity: function useNonce() returns(uint256)
func (_KeyRegistry *KeyRegistrySession) UseNonce() (*types.Transaction, error) {
	return _KeyRegistry.Contract.UseNonce(&_KeyRegistry.TransactOpts)
}

// UseNonce is a paid mutator transaction binding the contract method 0x69615a4c.
//
// Solidity: function useNonce() returns(uint256)
func (_KeyRegistry *KeyRegistryTransactorSession) UseNonce() (*types.Transaction, error) {
	return _KeyRegistry.Contract.UseNonce(&_KeyRegistry.TransactOpts)
}

// KeyRegistryAddIterator is returned from FilterAdd and is used to iterate over the raw logs and unpacked data for Add events raised by the KeyRegistry contract.
type KeyRegistryAddIterator struct {
	Event *KeyRegistryAdd // Event containing the contract specifics and raw log

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
func (it *KeyRegistryAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyRegistryAdd)
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
		it.Event = new(KeyRegistryAdd)
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
func (it *KeyRegistryAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyRegistryAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyRegistryAdd represents a Add event raised by the KeyRegistry contract.
type KeyRegistryAdd struct {
	Fid          *big.Int
	KeyType      uint32
	Key          common.Hash
	KeyBytes     []byte
	MetadataType uint8
	Metadata     []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAdd is a free log retrieval operation binding the contract event 0x7d285df41058466977811345cd453c0c52e8d841ffaabc74fc050f277ad4de02.
//
// Solidity: event Add(uint256 indexed fid, uint32 indexed keyType, bytes indexed key, bytes keyBytes, uint8 metadataType, bytes metadata)
func (_KeyRegistry *KeyRegistryFilterer) FilterAdd(opts *bind.FilterOpts, fid []*big.Int, keyType []uint32, key [][]byte) (*KeyRegistryAddIterator, error) {

	var fidRule []interface{}
	for _, fidItem := range fid {
		fidRule = append(fidRule, fidItem)
	}
	var keyTypeRule []interface{}
	for _, keyTypeItem := range keyType {
		keyTypeRule = append(keyTypeRule, keyTypeItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _KeyRegistry.contract.FilterLogs(opts, "Add", fidRule, keyTypeRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &KeyRegistryAddIterator{contract: _KeyRegistry.contract, event: "Add", logs: logs, sub: sub}, nil
}

// WatchAdd is a free log subscription operation binding the contract event 0x7d285df41058466977811345cd453c0c52e8d841ffaabc74fc050f277ad4de02.
//
// Solidity: event Add(uint256 indexed fid, uint32 indexed keyType, bytes indexed key, bytes keyBytes, uint8 metadataType, bytes metadata)
func (_KeyRegistry *KeyRegistryFilterer) WatchAdd(opts *bind.WatchOpts, sink chan<- *KeyRegistryAdd, fid []*big.Int, keyType []uint32, key [][]byte) (event.Subscription, error) {

	var fidRule []interface{}
	for _, fidItem := range fid {
		fidRule = append(fidRule, fidItem)
	}
	var keyTypeRule []interface{}
	for _, keyTypeItem := range keyType {
		keyTypeRule = append(keyTypeRule, keyTypeItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _KeyRegistry.contract.WatchLogs(opts, "Add", fidRule, keyTypeRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyRegistryAdd)
				if err := _KeyRegistry.contract.UnpackLog(event, "Add", log); err != nil {
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

// ParseAdd is a log parse operation binding the contract event 0x7d285df41058466977811345cd453c0c52e8d841ffaabc74fc050f277ad4de02.
//
// Solidity: event Add(uint256 indexed fid, uint32 indexed keyType, bytes indexed key, bytes keyBytes, uint8 metadataType, bytes metadata)
func (_KeyRegistry *KeyRegistryFilterer) ParseAdd(log types.Log) (*KeyRegistryAdd, error) {
	event := new(KeyRegistryAdd)
	if err := _KeyRegistry.contract.UnpackLog(event, "Add", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyRegistryAdd0Iterator is returned from FilterAdd0 and is used to iterate over the raw logs and unpacked data for Add0 events raised by the KeyRegistry contract.
type KeyRegistryAdd0Iterator struct {
	Event *KeyRegistryAdd0 // Event containing the contract specifics and raw log

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
func (it *KeyRegistryAdd0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyRegistryAdd0)
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
		it.Event = new(KeyRegistryAdd0)
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
func (it *KeyRegistryAdd0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyRegistryAdd0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyRegistryAdd0 represents a Add0 event raised by the KeyRegistry contract.
type KeyRegistryAdd0 struct {
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdd0 is a free log retrieval operation binding the contract event 0x87dc5eecd6d6bdeae407c426da6bfba5b7190befc554ed5d4d62dd5cf939fbae.
//
// Solidity: event Add(address indexed guardian)
func (_KeyRegistry *KeyRegistryFilterer) FilterAdd0(opts *bind.FilterOpts, guardian []common.Address) (*KeyRegistryAdd0Iterator, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _KeyRegistry.contract.FilterLogs(opts, "Add0", guardianRule)
	if err != nil {
		return nil, err
	}
	return &KeyRegistryAdd0Iterator{contract: _KeyRegistry.contract, event: "Add0", logs: logs, sub: sub}, nil
}

// WatchAdd0 is a free log subscription operation binding the contract event 0x87dc5eecd6d6bdeae407c426da6bfba5b7190befc554ed5d4d62dd5cf939fbae.
//
// Solidity: event Add(address indexed guardian)
func (_KeyRegistry *KeyRegistryFilterer) WatchAdd0(opts *bind.WatchOpts, sink chan<- *KeyRegistryAdd0, guardian []common.Address) (event.Subscription, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _KeyRegistry.contract.WatchLogs(opts, "Add0", guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyRegistryAdd0)
				if err := _KeyRegistry.contract.UnpackLog(event, "Add0", log); err != nil {
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

// ParseAdd0 is a log parse operation binding the contract event 0x87dc5eecd6d6bdeae407c426da6bfba5b7190befc554ed5d4d62dd5cf939fbae.
//
// Solidity: event Add(address indexed guardian)
func (_KeyRegistry *KeyRegistryFilterer) ParseAdd0(log types.Log) (*KeyRegistryAdd0, error) {
	event := new(KeyRegistryAdd0)
	if err := _KeyRegistry.contract.UnpackLog(event, "Add0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyRegistryAdminResetIterator is returned from FilterAdminReset and is used to iterate over the raw logs and unpacked data for AdminReset events raised by the KeyRegistry contract.
type KeyRegistryAdminResetIterator struct {
	Event *KeyRegistryAdminReset // Event containing the contract specifics and raw log

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
func (it *KeyRegistryAdminResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyRegistryAdminReset)
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
		it.Event = new(KeyRegistryAdminReset)
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
func (it *KeyRegistryAdminResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyRegistryAdminResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyRegistryAdminReset represents a AdminReset event raised by the KeyRegistry contract.
type KeyRegistryAdminReset struct {
	Fid      *big.Int
	Key      common.Hash
	KeyBytes []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminReset is a free log retrieval operation binding the contract event 0x1ecc1009ebad5d2fb61239462f4f9f6f152662defe1845fc87f07d96bd1c60b4.
//
// Solidity: event AdminReset(uint256 indexed fid, bytes indexed key, bytes keyBytes)
func (_KeyRegistry *KeyRegistryFilterer) FilterAdminReset(opts *bind.FilterOpts, fid []*big.Int, key [][]byte) (*KeyRegistryAdminResetIterator, error) {

	var fidRule []interface{}
	for _, fidItem := range fid {
		fidRule = append(fidRule, fidItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _KeyRegistry.contract.FilterLogs(opts, "AdminReset", fidRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &KeyRegistryAdminResetIterator{contract: _KeyRegistry.contract, event: "AdminReset", logs: logs, sub: sub}, nil
}

// WatchAdminReset is a free log subscription operation binding the contract event 0x1ecc1009ebad5d2fb61239462f4f9f6f152662defe1845fc87f07d96bd1c60b4.
//
// Solidity: event AdminReset(uint256 indexed fid, bytes indexed key, bytes keyBytes)
func (_KeyRegistry *KeyRegistryFilterer) WatchAdminReset(opts *bind.WatchOpts, sink chan<- *KeyRegistryAdminReset, fid []*big.Int, key [][]byte) (event.Subscription, error) {

	var fidRule []interface{}
	for _, fidItem := range fid {
		fidRule = append(fidRule, fidItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _KeyRegistry.contract.WatchLogs(opts, "AdminReset", fidRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyRegistryAdminReset)
				if err := _KeyRegistry.contract.UnpackLog(event, "AdminReset", log); err != nil {
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

// ParseAdminReset is a log parse operation binding the contract event 0x1ecc1009ebad5d2fb61239462f4f9f6f152662defe1845fc87f07d96bd1c60b4.
//
// Solidity: event AdminReset(uint256 indexed fid, bytes indexed key, bytes keyBytes)
func (_KeyRegistry *KeyRegistryFilterer) ParseAdminReset(log types.Log) (*KeyRegistryAdminReset, error) {
	event := new(KeyRegistryAdminReset)
	if err := _KeyRegistry.contract.UnpackLog(event, "AdminReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyRegistryEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the KeyRegistry contract.
type KeyRegistryEIP712DomainChangedIterator struct {
	Event *KeyRegistryEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *KeyRegistryEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyRegistryEIP712DomainChanged)
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
		it.Event = new(KeyRegistryEIP712DomainChanged)
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
func (it *KeyRegistryEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyRegistryEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyRegistryEIP712DomainChanged represents a EIP712DomainChanged event raised by the KeyRegistry contract.
type KeyRegistryEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_KeyRegistry *KeyRegistryFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*KeyRegistryEIP712DomainChangedIterator, error) {

	logs, sub, err := _KeyRegistry.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &KeyRegistryEIP712DomainChangedIterator{contract: _KeyRegistry.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_KeyRegistry *KeyRegistryFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *KeyRegistryEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _KeyRegistry.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyRegistryEIP712DomainChanged)
				if err := _KeyRegistry.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_KeyRegistry *KeyRegistryFilterer) ParseEIP712DomainChanged(log types.Log) (*KeyRegistryEIP712DomainChanged, error) {
	event := new(KeyRegistryEIP712DomainChanged)
	if err := _KeyRegistry.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyRegistryFreezeKeyGatewayIterator is returned from FilterFreezeKeyGateway and is used to iterate over the raw logs and unpacked data for FreezeKeyGateway events raised by the KeyRegistry contract.
type KeyRegistryFreezeKeyGatewayIterator struct {
	Event *KeyRegistryFreezeKeyGateway // Event containing the contract specifics and raw log

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
func (it *KeyRegistryFreezeKeyGatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyRegistryFreezeKeyGateway)
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
		it.Event = new(KeyRegistryFreezeKeyGateway)
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
func (it *KeyRegistryFreezeKeyGatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyRegistryFreezeKeyGatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyRegistryFreezeKeyGateway represents a FreezeKeyGateway event raised by the KeyRegistry contract.
type KeyRegistryFreezeKeyGateway struct {
	KeyGateway common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFreezeKeyGateway is a free log retrieval operation binding the contract event 0xcb685c7ba5a65fe9e6be9b7decbb5dc8ebba92bcae3cb09fc2a5a075b1eb5772.
//
// Solidity: event FreezeKeyGateway(address keyGateway)
func (_KeyRegistry *KeyRegistryFilterer) FilterFreezeKeyGateway(opts *bind.FilterOpts) (*KeyRegistryFreezeKeyGatewayIterator, error) {

	logs, sub, err := _KeyRegistry.contract.FilterLogs(opts, "FreezeKeyGateway")
	if err != nil {
		return nil, err
	}
	return &KeyRegistryFreezeKeyGatewayIterator{contract: _KeyRegistry.contract, event: "FreezeKeyGateway", logs: logs, sub: sub}, nil
}

// WatchFreezeKeyGateway is a free log subscription operation binding the contract event 0xcb685c7ba5a65fe9e6be9b7decbb5dc8ebba92bcae3cb09fc2a5a075b1eb5772.
//
// Solidity: event FreezeKeyGateway(address keyGateway)
func (_KeyRegistry *KeyRegistryFilterer) WatchFreezeKeyGateway(opts *bind.WatchOpts, sink chan<- *KeyRegistryFreezeKeyGateway) (event.Subscription, error) {

	logs, sub, err := _KeyRegistry.contract.WatchLogs(opts, "FreezeKeyGateway")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyRegistryFreezeKeyGateway)
				if err := _KeyRegistry.contract.UnpackLog(event, "FreezeKeyGateway", log); err != nil {
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

// ParseFreezeKeyGateway is a log parse operation binding the contract event 0xcb685c7ba5a65fe9e6be9b7decbb5dc8ebba92bcae3cb09fc2a5a075b1eb5772.
//
// Solidity: event FreezeKeyGateway(address keyGateway)
func (_KeyRegistry *KeyRegistryFilterer) ParseFreezeKeyGateway(log types.Log) (*KeyRegistryFreezeKeyGateway, error) {
	event := new(KeyRegistryFreezeKeyGateway)
	if err := _KeyRegistry.contract.UnpackLog(event, "FreezeKeyGateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyRegistryMigratedIterator is returned from FilterMigrated and is used to iterate over the raw logs and unpacked data for Migrated events raised by the KeyRegistry contract.
type KeyRegistryMigratedIterator struct {
	Event *KeyRegistryMigrated // Event containing the contract specifics and raw log

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
func (it *KeyRegistryMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyRegistryMigrated)
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
		it.Event = new(KeyRegistryMigrated)
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
func (it *KeyRegistryMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyRegistryMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyRegistryMigrated represents a Migrated event raised by the KeyRegistry contract.
type KeyRegistryMigrated struct {
	MigratedAt *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMigrated is a free log retrieval operation binding the contract event 0xe4a25c0c2cbe89d6ad8b64c61a7dbdd20d1f781f6023f1ab94ebb7fe0aef6ab8.
//
// Solidity: event Migrated(uint256 indexed migratedAt)
func (_KeyRegistry *KeyRegistryFilterer) FilterMigrated(opts *bind.FilterOpts, migratedAt []*big.Int) (*KeyRegistryMigratedIterator, error) {

	var migratedAtRule []interface{}
	for _, migratedAtItem := range migratedAt {
		migratedAtRule = append(migratedAtRule, migratedAtItem)
	}

	logs, sub, err := _KeyRegistry.contract.FilterLogs(opts, "Migrated", migratedAtRule)
	if err != nil {
		return nil, err
	}
	return &KeyRegistryMigratedIterator{contract: _KeyRegistry.contract, event: "Migrated", logs: logs, sub: sub}, nil
}

// WatchMigrated is a free log subscription operation binding the contract event 0xe4a25c0c2cbe89d6ad8b64c61a7dbdd20d1f781f6023f1ab94ebb7fe0aef6ab8.
//
// Solidity: event Migrated(uint256 indexed migratedAt)
func (_KeyRegistry *KeyRegistryFilterer) WatchMigrated(opts *bind.WatchOpts, sink chan<- *KeyRegistryMigrated, migratedAt []*big.Int) (event.Subscription, error) {

	var migratedAtRule []interface{}
	for _, migratedAtItem := range migratedAt {
		migratedAtRule = append(migratedAtRule, migratedAtItem)
	}

	logs, sub, err := _KeyRegistry.contract.WatchLogs(opts, "Migrated", migratedAtRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyRegistryMigrated)
				if err := _KeyRegistry.contract.UnpackLog(event, "Migrated", log); err != nil {
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
func (_KeyRegistry *KeyRegistryFilterer) ParseMigrated(log types.Log) (*KeyRegistryMigrated, error) {
	event := new(KeyRegistryMigrated)
	if err := _KeyRegistry.contract.UnpackLog(event, "Migrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyRegistryOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the KeyRegistry contract.
type KeyRegistryOwnershipTransferStartedIterator struct {
	Event *KeyRegistryOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *KeyRegistryOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyRegistryOwnershipTransferStarted)
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
		it.Event = new(KeyRegistryOwnershipTransferStarted)
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
func (it *KeyRegistryOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyRegistryOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyRegistryOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the KeyRegistry contract.
type KeyRegistryOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_KeyRegistry *KeyRegistryFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KeyRegistryOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KeyRegistry.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KeyRegistryOwnershipTransferStartedIterator{contract: _KeyRegistry.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_KeyRegistry *KeyRegistryFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *KeyRegistryOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KeyRegistry.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyRegistryOwnershipTransferStarted)
				if err := _KeyRegistry.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_KeyRegistry *KeyRegistryFilterer) ParseOwnershipTransferStarted(log types.Log) (*KeyRegistryOwnershipTransferStarted, error) {
	event := new(KeyRegistryOwnershipTransferStarted)
	if err := _KeyRegistry.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the KeyRegistry contract.
type KeyRegistryOwnershipTransferredIterator struct {
	Event *KeyRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *KeyRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyRegistryOwnershipTransferred)
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
		it.Event = new(KeyRegistryOwnershipTransferred)
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
func (it *KeyRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the KeyRegistry contract.
type KeyRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KeyRegistry *KeyRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KeyRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KeyRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KeyRegistryOwnershipTransferredIterator{contract: _KeyRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KeyRegistry *KeyRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KeyRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KeyRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyRegistryOwnershipTransferred)
				if err := _KeyRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_KeyRegistry *KeyRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*KeyRegistryOwnershipTransferred, error) {
	event := new(KeyRegistryOwnershipTransferred)
	if err := _KeyRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyRegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the KeyRegistry contract.
type KeyRegistryPausedIterator struct {
	Event *KeyRegistryPaused // Event containing the contract specifics and raw log

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
func (it *KeyRegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyRegistryPaused)
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
		it.Event = new(KeyRegistryPaused)
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
func (it *KeyRegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyRegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyRegistryPaused represents a Paused event raised by the KeyRegistry contract.
type KeyRegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_KeyRegistry *KeyRegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*KeyRegistryPausedIterator, error) {

	logs, sub, err := _KeyRegistry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &KeyRegistryPausedIterator{contract: _KeyRegistry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_KeyRegistry *KeyRegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *KeyRegistryPaused) (event.Subscription, error) {

	logs, sub, err := _KeyRegistry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyRegistryPaused)
				if err := _KeyRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_KeyRegistry *KeyRegistryFilterer) ParsePaused(log types.Log) (*KeyRegistryPaused, error) {
	event := new(KeyRegistryPaused)
	if err := _KeyRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyRegistryRemoveIterator is returned from FilterRemove and is used to iterate over the raw logs and unpacked data for Remove events raised by the KeyRegistry contract.
type KeyRegistryRemoveIterator struct {
	Event *KeyRegistryRemove // Event containing the contract specifics and raw log

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
func (it *KeyRegistryRemoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyRegistryRemove)
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
		it.Event = new(KeyRegistryRemove)
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
func (it *KeyRegistryRemoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyRegistryRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyRegistryRemove represents a Remove event raised by the KeyRegistry contract.
type KeyRegistryRemove struct {
	Fid      *big.Int
	Key      common.Hash
	KeyBytes []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemove is a free log retrieval operation binding the contract event 0x09e77066e0155f46785be12f6938a6b2e4be4381e59058129ce15f355cb96958.
//
// Solidity: event Remove(uint256 indexed fid, bytes indexed key, bytes keyBytes)
func (_KeyRegistry *KeyRegistryFilterer) FilterRemove(opts *bind.FilterOpts, fid []*big.Int, key [][]byte) (*KeyRegistryRemoveIterator, error) {

	var fidRule []interface{}
	for _, fidItem := range fid {
		fidRule = append(fidRule, fidItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _KeyRegistry.contract.FilterLogs(opts, "Remove", fidRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &KeyRegistryRemoveIterator{contract: _KeyRegistry.contract, event: "Remove", logs: logs, sub: sub}, nil
}

// WatchRemove is a free log subscription operation binding the contract event 0x09e77066e0155f46785be12f6938a6b2e4be4381e59058129ce15f355cb96958.
//
// Solidity: event Remove(uint256 indexed fid, bytes indexed key, bytes keyBytes)
func (_KeyRegistry *KeyRegistryFilterer) WatchRemove(opts *bind.WatchOpts, sink chan<- *KeyRegistryRemove, fid []*big.Int, key [][]byte) (event.Subscription, error) {

	var fidRule []interface{}
	for _, fidItem := range fid {
		fidRule = append(fidRule, fidItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _KeyRegistry.contract.WatchLogs(opts, "Remove", fidRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyRegistryRemove)
				if err := _KeyRegistry.contract.UnpackLog(event, "Remove", log); err != nil {
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

// ParseRemove is a log parse operation binding the contract event 0x09e77066e0155f46785be12f6938a6b2e4be4381e59058129ce15f355cb96958.
//
// Solidity: event Remove(uint256 indexed fid, bytes indexed key, bytes keyBytes)
func (_KeyRegistry *KeyRegistryFilterer) ParseRemove(log types.Log) (*KeyRegistryRemove, error) {
	event := new(KeyRegistryRemove)
	if err := _KeyRegistry.contract.UnpackLog(event, "Remove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyRegistryRemove0Iterator is returned from FilterRemove0 and is used to iterate over the raw logs and unpacked data for Remove0 events raised by the KeyRegistry contract.
type KeyRegistryRemove0Iterator struct {
	Event *KeyRegistryRemove0 // Event containing the contract specifics and raw log

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
func (it *KeyRegistryRemove0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyRegistryRemove0)
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
		it.Event = new(KeyRegistryRemove0)
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
func (it *KeyRegistryRemove0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyRegistryRemove0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyRegistryRemove0 represents a Remove0 event raised by the KeyRegistry contract.
type KeyRegistryRemove0 struct {
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemove0 is a free log retrieval operation binding the contract event 0xbe7c7ac3248df4581c206a84aab3cb4e7d521b5398b42b681757f78a5a7d411e.
//
// Solidity: event Remove(address indexed guardian)
func (_KeyRegistry *KeyRegistryFilterer) FilterRemove0(opts *bind.FilterOpts, guardian []common.Address) (*KeyRegistryRemove0Iterator, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _KeyRegistry.contract.FilterLogs(opts, "Remove0", guardianRule)
	if err != nil {
		return nil, err
	}
	return &KeyRegistryRemove0Iterator{contract: _KeyRegistry.contract, event: "Remove0", logs: logs, sub: sub}, nil
}

// WatchRemove0 is a free log subscription operation binding the contract event 0xbe7c7ac3248df4581c206a84aab3cb4e7d521b5398b42b681757f78a5a7d411e.
//
// Solidity: event Remove(address indexed guardian)
func (_KeyRegistry *KeyRegistryFilterer) WatchRemove0(opts *bind.WatchOpts, sink chan<- *KeyRegistryRemove0, guardian []common.Address) (event.Subscription, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _KeyRegistry.contract.WatchLogs(opts, "Remove0", guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyRegistryRemove0)
				if err := _KeyRegistry.contract.UnpackLog(event, "Remove0", log); err != nil {
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

// ParseRemove0 is a log parse operation binding the contract event 0xbe7c7ac3248df4581c206a84aab3cb4e7d521b5398b42b681757f78a5a7d411e.
//
// Solidity: event Remove(address indexed guardian)
func (_KeyRegistry *KeyRegistryFilterer) ParseRemove0(log types.Log) (*KeyRegistryRemove0, error) {
	event := new(KeyRegistryRemove0)
	if err := _KeyRegistry.contract.UnpackLog(event, "Remove0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyRegistrySetIdRegistryIterator is returned from FilterSetIdRegistry and is used to iterate over the raw logs and unpacked data for SetIdRegistry events raised by the KeyRegistry contract.
type KeyRegistrySetIdRegistryIterator struct {
	Event *KeyRegistrySetIdRegistry // Event containing the contract specifics and raw log

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
func (it *KeyRegistrySetIdRegistryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyRegistrySetIdRegistry)
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
		it.Event = new(KeyRegistrySetIdRegistry)
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
func (it *KeyRegistrySetIdRegistryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyRegistrySetIdRegistryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyRegistrySetIdRegistry represents a SetIdRegistry event raised by the KeyRegistry contract.
type KeyRegistrySetIdRegistry struct {
	OldIdRegistry common.Address
	NewIdRegistry common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetIdRegistry is a free log retrieval operation binding the contract event 0x940dcf34ec2e245e837ee4599997e577ce274d7b87c73238e2878ac7ea1af2f1.
//
// Solidity: event SetIdRegistry(address oldIdRegistry, address newIdRegistry)
func (_KeyRegistry *KeyRegistryFilterer) FilterSetIdRegistry(opts *bind.FilterOpts) (*KeyRegistrySetIdRegistryIterator, error) {

	logs, sub, err := _KeyRegistry.contract.FilterLogs(opts, "SetIdRegistry")
	if err != nil {
		return nil, err
	}
	return &KeyRegistrySetIdRegistryIterator{contract: _KeyRegistry.contract, event: "SetIdRegistry", logs: logs, sub: sub}, nil
}

// WatchSetIdRegistry is a free log subscription operation binding the contract event 0x940dcf34ec2e245e837ee4599997e577ce274d7b87c73238e2878ac7ea1af2f1.
//
// Solidity: event SetIdRegistry(address oldIdRegistry, address newIdRegistry)
func (_KeyRegistry *KeyRegistryFilterer) WatchSetIdRegistry(opts *bind.WatchOpts, sink chan<- *KeyRegistrySetIdRegistry) (event.Subscription, error) {

	logs, sub, err := _KeyRegistry.contract.WatchLogs(opts, "SetIdRegistry")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyRegistrySetIdRegistry)
				if err := _KeyRegistry.contract.UnpackLog(event, "SetIdRegistry", log); err != nil {
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

// ParseSetIdRegistry is a log parse operation binding the contract event 0x940dcf34ec2e245e837ee4599997e577ce274d7b87c73238e2878ac7ea1af2f1.
//
// Solidity: event SetIdRegistry(address oldIdRegistry, address newIdRegistry)
func (_KeyRegistry *KeyRegistryFilterer) ParseSetIdRegistry(log types.Log) (*KeyRegistrySetIdRegistry, error) {
	event := new(KeyRegistrySetIdRegistry)
	if err := _KeyRegistry.contract.UnpackLog(event, "SetIdRegistry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyRegistrySetKeyGatewayIterator is returned from FilterSetKeyGateway and is used to iterate over the raw logs and unpacked data for SetKeyGateway events raised by the KeyRegistry contract.
type KeyRegistrySetKeyGatewayIterator struct {
	Event *KeyRegistrySetKeyGateway // Event containing the contract specifics and raw log

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
func (it *KeyRegistrySetKeyGatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyRegistrySetKeyGateway)
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
		it.Event = new(KeyRegistrySetKeyGateway)
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
func (it *KeyRegistrySetKeyGatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyRegistrySetKeyGatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyRegistrySetKeyGateway represents a SetKeyGateway event raised by the KeyRegistry contract.
type KeyRegistrySetKeyGateway struct {
	OldKeyGateway common.Address
	NewKeyGateway common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetKeyGateway is a free log retrieval operation binding the contract event 0x56785750704201befc0a27dae1e5d37835a8ad6e35affc87136ed24d1ac694ac.
//
// Solidity: event SetKeyGateway(address oldKeyGateway, address newKeyGateway)
func (_KeyRegistry *KeyRegistryFilterer) FilterSetKeyGateway(opts *bind.FilterOpts) (*KeyRegistrySetKeyGatewayIterator, error) {

	logs, sub, err := _KeyRegistry.contract.FilterLogs(opts, "SetKeyGateway")
	if err != nil {
		return nil, err
	}
	return &KeyRegistrySetKeyGatewayIterator{contract: _KeyRegistry.contract, event: "SetKeyGateway", logs: logs, sub: sub}, nil
}

// WatchSetKeyGateway is a free log subscription operation binding the contract event 0x56785750704201befc0a27dae1e5d37835a8ad6e35affc87136ed24d1ac694ac.
//
// Solidity: event SetKeyGateway(address oldKeyGateway, address newKeyGateway)
func (_KeyRegistry *KeyRegistryFilterer) WatchSetKeyGateway(opts *bind.WatchOpts, sink chan<- *KeyRegistrySetKeyGateway) (event.Subscription, error) {

	logs, sub, err := _KeyRegistry.contract.WatchLogs(opts, "SetKeyGateway")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyRegistrySetKeyGateway)
				if err := _KeyRegistry.contract.UnpackLog(event, "SetKeyGateway", log); err != nil {
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

// ParseSetKeyGateway is a log parse operation binding the contract event 0x56785750704201befc0a27dae1e5d37835a8ad6e35affc87136ed24d1ac694ac.
//
// Solidity: event SetKeyGateway(address oldKeyGateway, address newKeyGateway)
func (_KeyRegistry *KeyRegistryFilterer) ParseSetKeyGateway(log types.Log) (*KeyRegistrySetKeyGateway, error) {
	event := new(KeyRegistrySetKeyGateway)
	if err := _KeyRegistry.contract.UnpackLog(event, "SetKeyGateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyRegistrySetMaxKeysPerFidIterator is returned from FilterSetMaxKeysPerFid and is used to iterate over the raw logs and unpacked data for SetMaxKeysPerFid events raised by the KeyRegistry contract.
type KeyRegistrySetMaxKeysPerFidIterator struct {
	Event *KeyRegistrySetMaxKeysPerFid // Event containing the contract specifics and raw log

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
func (it *KeyRegistrySetMaxKeysPerFidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyRegistrySetMaxKeysPerFid)
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
		it.Event = new(KeyRegistrySetMaxKeysPerFid)
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
func (it *KeyRegistrySetMaxKeysPerFidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyRegistrySetMaxKeysPerFidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyRegistrySetMaxKeysPerFid represents a SetMaxKeysPerFid event raised by the KeyRegistry contract.
type KeyRegistrySetMaxKeysPerFid struct {
	OldMax *big.Int
	NewMax *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetMaxKeysPerFid is a free log retrieval operation binding the contract event 0x6c336d0e5e74bb26a5e2d4646801801837b0fdbaddd9131923fd42d740449731.
//
// Solidity: event SetMaxKeysPerFid(uint256 oldMax, uint256 newMax)
func (_KeyRegistry *KeyRegistryFilterer) FilterSetMaxKeysPerFid(opts *bind.FilterOpts) (*KeyRegistrySetMaxKeysPerFidIterator, error) {

	logs, sub, err := _KeyRegistry.contract.FilterLogs(opts, "SetMaxKeysPerFid")
	if err != nil {
		return nil, err
	}
	return &KeyRegistrySetMaxKeysPerFidIterator{contract: _KeyRegistry.contract, event: "SetMaxKeysPerFid", logs: logs, sub: sub}, nil
}

// WatchSetMaxKeysPerFid is a free log subscription operation binding the contract event 0x6c336d0e5e74bb26a5e2d4646801801837b0fdbaddd9131923fd42d740449731.
//
// Solidity: event SetMaxKeysPerFid(uint256 oldMax, uint256 newMax)
func (_KeyRegistry *KeyRegistryFilterer) WatchSetMaxKeysPerFid(opts *bind.WatchOpts, sink chan<- *KeyRegistrySetMaxKeysPerFid) (event.Subscription, error) {

	logs, sub, err := _KeyRegistry.contract.WatchLogs(opts, "SetMaxKeysPerFid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyRegistrySetMaxKeysPerFid)
				if err := _KeyRegistry.contract.UnpackLog(event, "SetMaxKeysPerFid", log); err != nil {
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

// ParseSetMaxKeysPerFid is a log parse operation binding the contract event 0x6c336d0e5e74bb26a5e2d4646801801837b0fdbaddd9131923fd42d740449731.
//
// Solidity: event SetMaxKeysPerFid(uint256 oldMax, uint256 newMax)
func (_KeyRegistry *KeyRegistryFilterer) ParseSetMaxKeysPerFid(log types.Log) (*KeyRegistrySetMaxKeysPerFid, error) {
	event := new(KeyRegistrySetMaxKeysPerFid)
	if err := _KeyRegistry.contract.UnpackLog(event, "SetMaxKeysPerFid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyRegistrySetMigratorIterator is returned from FilterSetMigrator and is used to iterate over the raw logs and unpacked data for SetMigrator events raised by the KeyRegistry contract.
type KeyRegistrySetMigratorIterator struct {
	Event *KeyRegistrySetMigrator // Event containing the contract specifics and raw log

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
func (it *KeyRegistrySetMigratorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyRegistrySetMigrator)
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
		it.Event = new(KeyRegistrySetMigrator)
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
func (it *KeyRegistrySetMigratorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyRegistrySetMigratorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyRegistrySetMigrator represents a SetMigrator event raised by the KeyRegistry contract.
type KeyRegistrySetMigrator struct {
	OldMigrator common.Address
	NewMigrator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetMigrator is a free log retrieval operation binding the contract event 0xd8ad954fe808212ab9ed7139873e40807dff7995fe36e3d6cdeb8fa00fcebf10.
//
// Solidity: event SetMigrator(address oldMigrator, address newMigrator)
func (_KeyRegistry *KeyRegistryFilterer) FilterSetMigrator(opts *bind.FilterOpts) (*KeyRegistrySetMigratorIterator, error) {

	logs, sub, err := _KeyRegistry.contract.FilterLogs(opts, "SetMigrator")
	if err != nil {
		return nil, err
	}
	return &KeyRegistrySetMigratorIterator{contract: _KeyRegistry.contract, event: "SetMigrator", logs: logs, sub: sub}, nil
}

// WatchSetMigrator is a free log subscription operation binding the contract event 0xd8ad954fe808212ab9ed7139873e40807dff7995fe36e3d6cdeb8fa00fcebf10.
//
// Solidity: event SetMigrator(address oldMigrator, address newMigrator)
func (_KeyRegistry *KeyRegistryFilterer) WatchSetMigrator(opts *bind.WatchOpts, sink chan<- *KeyRegistrySetMigrator) (event.Subscription, error) {

	logs, sub, err := _KeyRegistry.contract.WatchLogs(opts, "SetMigrator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyRegistrySetMigrator)
				if err := _KeyRegistry.contract.UnpackLog(event, "SetMigrator", log); err != nil {
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
func (_KeyRegistry *KeyRegistryFilterer) ParseSetMigrator(log types.Log) (*KeyRegistrySetMigrator, error) {
	event := new(KeyRegistrySetMigrator)
	if err := _KeyRegistry.contract.UnpackLog(event, "SetMigrator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyRegistrySetValidatorIterator is returned from FilterSetValidator and is used to iterate over the raw logs and unpacked data for SetValidator events raised by the KeyRegistry contract.
type KeyRegistrySetValidatorIterator struct {
	Event *KeyRegistrySetValidator // Event containing the contract specifics and raw log

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
func (it *KeyRegistrySetValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyRegistrySetValidator)
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
		it.Event = new(KeyRegistrySetValidator)
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
func (it *KeyRegistrySetValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyRegistrySetValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyRegistrySetValidator represents a SetValidator event raised by the KeyRegistry contract.
type KeyRegistrySetValidator struct {
	KeyType      uint32
	MetadataType uint8
	OldValidator common.Address
	NewValidator common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetValidator is a free log retrieval operation binding the contract event 0xd964242236f6208120d76a25cd886db49c82403f50d88dfd1bc865ee60ad462d.
//
// Solidity: event SetValidator(uint32 keyType, uint8 metadataType, address oldValidator, address newValidator)
func (_KeyRegistry *KeyRegistryFilterer) FilterSetValidator(opts *bind.FilterOpts) (*KeyRegistrySetValidatorIterator, error) {

	logs, sub, err := _KeyRegistry.contract.FilterLogs(opts, "SetValidator")
	if err != nil {
		return nil, err
	}
	return &KeyRegistrySetValidatorIterator{contract: _KeyRegistry.contract, event: "SetValidator", logs: logs, sub: sub}, nil
}

// WatchSetValidator is a free log subscription operation binding the contract event 0xd964242236f6208120d76a25cd886db49c82403f50d88dfd1bc865ee60ad462d.
//
// Solidity: event SetValidator(uint32 keyType, uint8 metadataType, address oldValidator, address newValidator)
func (_KeyRegistry *KeyRegistryFilterer) WatchSetValidator(opts *bind.WatchOpts, sink chan<- *KeyRegistrySetValidator) (event.Subscription, error) {

	logs, sub, err := _KeyRegistry.contract.WatchLogs(opts, "SetValidator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyRegistrySetValidator)
				if err := _KeyRegistry.contract.UnpackLog(event, "SetValidator", log); err != nil {
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

// ParseSetValidator is a log parse operation binding the contract event 0xd964242236f6208120d76a25cd886db49c82403f50d88dfd1bc865ee60ad462d.
//
// Solidity: event SetValidator(uint32 keyType, uint8 metadataType, address oldValidator, address newValidator)
func (_KeyRegistry *KeyRegistryFilterer) ParseSetValidator(log types.Log) (*KeyRegistrySetValidator, error) {
	event := new(KeyRegistrySetValidator)
	if err := _KeyRegistry.contract.UnpackLog(event, "SetValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyRegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the KeyRegistry contract.
type KeyRegistryUnpausedIterator struct {
	Event *KeyRegistryUnpaused // Event containing the contract specifics and raw log

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
func (it *KeyRegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyRegistryUnpaused)
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
		it.Event = new(KeyRegistryUnpaused)
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
func (it *KeyRegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyRegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyRegistryUnpaused represents a Unpaused event raised by the KeyRegistry contract.
type KeyRegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_KeyRegistry *KeyRegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*KeyRegistryUnpausedIterator, error) {

	logs, sub, err := _KeyRegistry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &KeyRegistryUnpausedIterator{contract: _KeyRegistry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_KeyRegistry *KeyRegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *KeyRegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _KeyRegistry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyRegistryUnpaused)
				if err := _KeyRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_KeyRegistry *KeyRegistryFilterer) ParseUnpaused(log types.Log) (*KeyRegistryUnpaused, error) {
	event := new(KeyRegistryUnpaused)
	if err := _KeyRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
