// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package generated

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MicroPaySystemABI is the input ABI used to generate the binding from.
const MicroPaySystemABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ta\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"packet\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tonken\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"micrNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimNonce\",\"type\":\"uint256\"}],\"name\":\"PoolClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"PoolDestroy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gtn\",\"type\":\"uint256\"}],\"name\":\"PoolRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"packets\",\"type\":\"uint256\"}],\"name\":\"UserRecharge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"packets\",\"type\":\"uint256\"}],\"name\":\"UserRefund\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"AllMinersOfPool\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"AllMyPools\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenNo\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"poolAddr\",\"type\":\"address\"}],\"name\":\"BuyPacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"mainAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"email\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"ChangeBasicInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"subAddr\",\"type\":\"bytes32\"}],\"name\":\"ChangePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"credit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"micNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cn\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"Claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CurMsgHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddr\",\"type\":\"address\"}],\"name\":\"DestroyPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenNo\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"subAddr\",\"type\":\"bytes32\"},{\"internalType\":\"bytes2\",\"name\":\"zone\",\"type\":\"bytes2\"}],\"name\":\"JoinPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"MinerData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"poolAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"subAddr\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"GTN\",\"type\":\"uint256\"},{\"internalType\":\"bytes2\",\"name\":\"zone\",\"type\":\"bytes2\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"MinerNoOfPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"PartOfMiners\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PoolData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mainAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payerAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"GTN\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"shortName\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"email\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Pools\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PoolsIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RecoveredAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenNo\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"poolAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"email\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"RegPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subAddr\",\"type\":\"bytes32\"}],\"name\":\"RetireFromPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"micNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cn\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"TestClaimSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"TokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"UserData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remindPacket\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimedMicNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pGTN\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mGTN\",\"type\":\"uint256\"}],\"name\":\"changeSetting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"minersUnderPool\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolsUsedByUser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setting\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"MBytesPerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MinPoolGuarantee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MinMinerGuarantee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"usersOfPool\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usersUnderPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vm\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"versionOfPoolData\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"versionOfMinerData\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"versionOfUserData\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"versionOfSysSetting\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MicroPaySystem is an auto generated Go binding around an Ethereum contract.
type MicroPaySystem struct {
	MicroPaySystemCaller     // Read-only binding to the contract
	MicroPaySystemTransactor // Write-only binding to the contract
	MicroPaySystemFilterer   // Log filterer for contract events
}

// MicroPaySystemCaller is an auto generated read-only Go binding around an Ethereum contract.
type MicroPaySystemCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MicroPaySystemTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MicroPaySystemTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MicroPaySystemFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MicroPaySystemFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MicroPaySystemSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MicroPaySystemSession struct {
	Contract     *MicroPaySystem   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MicroPaySystemCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MicroPaySystemCallerSession struct {
	Contract *MicroPaySystemCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MicroPaySystemTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MicroPaySystemTransactorSession struct {
	Contract     *MicroPaySystemTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MicroPaySystemRaw is an auto generated low-level Go binding around an Ethereum contract.
type MicroPaySystemRaw struct {
	Contract *MicroPaySystem // Generic contract binding to access the raw methods on
}

// MicroPaySystemCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MicroPaySystemCallerRaw struct {
	Contract *MicroPaySystemCaller // Generic read-only contract binding to access the raw methods on
}

// MicroPaySystemTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MicroPaySystemTransactorRaw struct {
	Contract *MicroPaySystemTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMicroPaySystem creates a new instance of MicroPaySystem, bound to a specific deployed contract.
func NewMicroPaySystem(address common.Address, backend bind.ContractBackend) (*MicroPaySystem, error) {
	contract, err := bindMicroPaySystem(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MicroPaySystem{MicroPaySystemCaller: MicroPaySystemCaller{contract: contract}, MicroPaySystemTransactor: MicroPaySystemTransactor{contract: contract}, MicroPaySystemFilterer: MicroPaySystemFilterer{contract: contract}}, nil
}

// NewMicroPaySystemCaller creates a new read-only instance of MicroPaySystem, bound to a specific deployed contract.
func NewMicroPaySystemCaller(address common.Address, caller bind.ContractCaller) (*MicroPaySystemCaller, error) {
	contract, err := bindMicroPaySystem(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MicroPaySystemCaller{contract: contract}, nil
}

// NewMicroPaySystemTransactor creates a new write-only instance of MicroPaySystem, bound to a specific deployed contract.
func NewMicroPaySystemTransactor(address common.Address, transactor bind.ContractTransactor) (*MicroPaySystemTransactor, error) {
	contract, err := bindMicroPaySystem(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MicroPaySystemTransactor{contract: contract}, nil
}

// NewMicroPaySystemFilterer creates a new log filterer instance of MicroPaySystem, bound to a specific deployed contract.
func NewMicroPaySystemFilterer(address common.Address, filterer bind.ContractFilterer) (*MicroPaySystemFilterer, error) {
	contract, err := bindMicroPaySystem(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MicroPaySystemFilterer{contract: contract}, nil
}

// bindMicroPaySystem binds a generic wrapper to an already deployed contract.
func bindMicroPaySystem(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MicroPaySystemABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MicroPaySystem *MicroPaySystemRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MicroPaySystem.Contract.MicroPaySystemCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MicroPaySystem *MicroPaySystemRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.MicroPaySystemTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MicroPaySystem *MicroPaySystemRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.MicroPaySystemTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MicroPaySystem *MicroPaySystemCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MicroPaySystem.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MicroPaySystem *MicroPaySystemTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MicroPaySystem *MicroPaySystemTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.contract.Transact(opts, method, params...)
}

// AllMinersOfPool is a free data retrieval call binding the contract method 0xdd7b0995.
//
// Solidity: function AllMinersOfPool(address pool) constant returns(bytes32[])
func (_MicroPaySystem *MicroPaySystemCaller) AllMinersOfPool(opts *bind.CallOpts, pool common.Address) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "AllMinersOfPool", pool)
	return *ret0, err
}

// AllMinersOfPool is a free data retrieval call binding the contract method 0xdd7b0995.
//
// Solidity: function AllMinersOfPool(address pool) constant returns(bytes32[])
func (_MicroPaySystem *MicroPaySystemSession) AllMinersOfPool(pool common.Address) ([][32]byte, error) {
	return _MicroPaySystem.Contract.AllMinersOfPool(&_MicroPaySystem.CallOpts, pool)
}

// AllMinersOfPool is a free data retrieval call binding the contract method 0xdd7b0995.
//
// Solidity: function AllMinersOfPool(address pool) constant returns(bytes32[])
func (_MicroPaySystem *MicroPaySystemCallerSession) AllMinersOfPool(pool common.Address) ([][32]byte, error) {
	return _MicroPaySystem.Contract.AllMinersOfPool(&_MicroPaySystem.CallOpts, pool)
}

// AllMyPools is a free data retrieval call binding the contract method 0x8fd83f76.
//
// Solidity: function AllMyPools(address userAddress) constant returns(address[])
func (_MicroPaySystem *MicroPaySystemCaller) AllMyPools(opts *bind.CallOpts, userAddress common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "AllMyPools", userAddress)
	return *ret0, err
}

// AllMyPools is a free data retrieval call binding the contract method 0x8fd83f76.
//
// Solidity: function AllMyPools(address userAddress) constant returns(address[])
func (_MicroPaySystem *MicroPaySystemSession) AllMyPools(userAddress common.Address) ([]common.Address, error) {
	return _MicroPaySystem.Contract.AllMyPools(&_MicroPaySystem.CallOpts, userAddress)
}

// AllMyPools is a free data retrieval call binding the contract method 0x8fd83f76.
//
// Solidity: function AllMyPools(address userAddress) constant returns(address[])
func (_MicroPaySystem *MicroPaySystemCallerSession) AllMyPools(userAddress common.Address) ([]common.Address, error) {
	return _MicroPaySystem.Contract.AllMyPools(&_MicroPaySystem.CallOpts, userAddress)
}

// CurMsgHash is a free data retrieval call binding the contract method 0xc3debfa2.
//
// Solidity: function CurMsgHash() constant returns(bytes32)
func (_MicroPaySystem *MicroPaySystemCaller) CurMsgHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "CurMsgHash")
	return *ret0, err
}

// CurMsgHash is a free data retrieval call binding the contract method 0xc3debfa2.
//
// Solidity: function CurMsgHash() constant returns(bytes32)
func (_MicroPaySystem *MicroPaySystemSession) CurMsgHash() ([32]byte, error) {
	return _MicroPaySystem.Contract.CurMsgHash(&_MicroPaySystem.CallOpts)
}

// CurMsgHash is a free data retrieval call binding the contract method 0xc3debfa2.
//
// Solidity: function CurMsgHash() constant returns(bytes32)
func (_MicroPaySystem *MicroPaySystemCallerSession) CurMsgHash() ([32]byte, error) {
	return _MicroPaySystem.Contract.CurMsgHash(&_MicroPaySystem.CallOpts)
}

// MinerData is a free data retrieval call binding the contract method 0x071b8916.
//
// Solidity: function MinerData(bytes32 ) constant returns(uint256 ID, address payer, address poolAddr, bytes32 subAddr, uint256 GTN, bytes2 zone)
func (_MicroPaySystem *MicroPaySystemCaller) MinerData(opts *bind.CallOpts, arg0 [32]byte) (struct {
	ID       *big.Int
	Payer    common.Address
	PoolAddr common.Address
	SubAddr  [32]byte
	GTN      *big.Int
	Zone     [2]byte
}, error) {
	ret := new(struct {
		ID       *big.Int
		Payer    common.Address
		PoolAddr common.Address
		SubAddr  [32]byte
		GTN      *big.Int
		Zone     [2]byte
	})
	out := ret
	err := _MicroPaySystem.contract.Call(opts, out, "MinerData", arg0)
	return *ret, err
}

// MinerData is a free data retrieval call binding the contract method 0x071b8916.
//
// Solidity: function MinerData(bytes32 ) constant returns(uint256 ID, address payer, address poolAddr, bytes32 subAddr, uint256 GTN, bytes2 zone)
func (_MicroPaySystem *MicroPaySystemSession) MinerData(arg0 [32]byte) (struct {
	ID       *big.Int
	Payer    common.Address
	PoolAddr common.Address
	SubAddr  [32]byte
	GTN      *big.Int
	Zone     [2]byte
}, error) {
	return _MicroPaySystem.Contract.MinerData(&_MicroPaySystem.CallOpts, arg0)
}

// MinerData is a free data retrieval call binding the contract method 0x071b8916.
//
// Solidity: function MinerData(bytes32 ) constant returns(uint256 ID, address payer, address poolAddr, bytes32 subAddr, uint256 GTN, bytes2 zone)
func (_MicroPaySystem *MicroPaySystemCallerSession) MinerData(arg0 [32]byte) (struct {
	ID       *big.Int
	Payer    common.Address
	PoolAddr common.Address
	SubAddr  [32]byte
	GTN      *big.Int
	Zone     [2]byte
}, error) {
	return _MicroPaySystem.Contract.MinerData(&_MicroPaySystem.CallOpts, arg0)
}

// MinerNoOfPool is a free data retrieval call binding the contract method 0x072d3e0d.
//
// Solidity: function MinerNoOfPool(address pool) constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemCaller) MinerNoOfPool(opts *bind.CallOpts, pool common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "MinerNoOfPool", pool)
	return *ret0, err
}

// MinerNoOfPool is a free data retrieval call binding the contract method 0x072d3e0d.
//
// Solidity: function MinerNoOfPool(address pool) constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemSession) MinerNoOfPool(pool common.Address) (*big.Int, error) {
	return _MicroPaySystem.Contract.MinerNoOfPool(&_MicroPaySystem.CallOpts, pool)
}

// MinerNoOfPool is a free data retrieval call binding the contract method 0x072d3e0d.
//
// Solidity: function MinerNoOfPool(address pool) constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemCallerSession) MinerNoOfPool(pool common.Address) (*big.Int, error) {
	return _MicroPaySystem.Contract.MinerNoOfPool(&_MicroPaySystem.CallOpts, pool)
}

// PartOfMiners is a free data retrieval call binding the contract method 0x962cfa78.
//
// Solidity: function PartOfMiners(address pool, uint256 start, uint256 end) constant returns(bytes32[])
func (_MicroPaySystem *MicroPaySystemCaller) PartOfMiners(opts *bind.CallOpts, pool common.Address, start *big.Int, end *big.Int) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "PartOfMiners", pool, start, end)
	return *ret0, err
}

// PartOfMiners is a free data retrieval call binding the contract method 0x962cfa78.
//
// Solidity: function PartOfMiners(address pool, uint256 start, uint256 end) constant returns(bytes32[])
func (_MicroPaySystem *MicroPaySystemSession) PartOfMiners(pool common.Address, start *big.Int, end *big.Int) ([][32]byte, error) {
	return _MicroPaySystem.Contract.PartOfMiners(&_MicroPaySystem.CallOpts, pool, start, end)
}

// PartOfMiners is a free data retrieval call binding the contract method 0x962cfa78.
//
// Solidity: function PartOfMiners(address pool, uint256 start, uint256 end) constant returns(bytes32[])
func (_MicroPaySystem *MicroPaySystemCallerSession) PartOfMiners(pool common.Address, start *big.Int, end *big.Int) ([][32]byte, error) {
	return _MicroPaySystem.Contract.PartOfMiners(&_MicroPaySystem.CallOpts, pool, start, end)
}

// PoolData is a free data retrieval call binding the contract method 0xf1b0e36d.
//
// Solidity: function PoolData(address ) constant returns(uint256 ID, address mainAddr, address payerAddr, uint256 GTN, bytes shortName, bytes email, bytes url)
func (_MicroPaySystem *MicroPaySystemCaller) PoolData(opts *bind.CallOpts, arg0 common.Address) (struct {
	ID        *big.Int
	MainAddr  common.Address
	PayerAddr common.Address
	GTN       *big.Int
	ShortName []byte
	Email     []byte
	Url       []byte
}, error) {
	ret := new(struct {
		ID        *big.Int
		MainAddr  common.Address
		PayerAddr common.Address
		GTN       *big.Int
		ShortName []byte
		Email     []byte
		Url       []byte
	})
	out := ret
	err := _MicroPaySystem.contract.Call(opts, out, "PoolData", arg0)
	return *ret, err
}

// PoolData is a free data retrieval call binding the contract method 0xf1b0e36d.
//
// Solidity: function PoolData(address ) constant returns(uint256 ID, address mainAddr, address payerAddr, uint256 GTN, bytes shortName, bytes email, bytes url)
func (_MicroPaySystem *MicroPaySystemSession) PoolData(arg0 common.Address) (struct {
	ID        *big.Int
	MainAddr  common.Address
	PayerAddr common.Address
	GTN       *big.Int
	ShortName []byte
	Email     []byte
	Url       []byte
}, error) {
	return _MicroPaySystem.Contract.PoolData(&_MicroPaySystem.CallOpts, arg0)
}

// PoolData is a free data retrieval call binding the contract method 0xf1b0e36d.
//
// Solidity: function PoolData(address ) constant returns(uint256 ID, address mainAddr, address payerAddr, uint256 GTN, bytes shortName, bytes email, bytes url)
func (_MicroPaySystem *MicroPaySystemCallerSession) PoolData(arg0 common.Address) (struct {
	ID        *big.Int
	MainAddr  common.Address
	PayerAddr common.Address
	GTN       *big.Int
	ShortName []byte
	Email     []byte
	Url       []byte
}, error) {
	return _MicroPaySystem.Contract.PoolData(&_MicroPaySystem.CallOpts, arg0)
}

// Pools is a free data retrieval call binding the contract method 0xae075993.
//
// Solidity: function Pools() constant returns(address[])
func (_MicroPaySystem *MicroPaySystemCaller) Pools(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "Pools")
	return *ret0, err
}

// Pools is a free data retrieval call binding the contract method 0xae075993.
//
// Solidity: function Pools() constant returns(address[])
func (_MicroPaySystem *MicroPaySystemSession) Pools() ([]common.Address, error) {
	return _MicroPaySystem.Contract.Pools(&_MicroPaySystem.CallOpts)
}

// Pools is a free data retrieval call binding the contract method 0xae075993.
//
// Solidity: function Pools() constant returns(address[])
func (_MicroPaySystem *MicroPaySystemCallerSession) Pools() ([]common.Address, error) {
	return _MicroPaySystem.Contract.Pools(&_MicroPaySystem.CallOpts)
}

// PoolsIndex is a free data retrieval call binding the contract method 0x6d4347da.
//
// Solidity: function PoolsIndex(uint256 ) constant returns(address)
func (_MicroPaySystem *MicroPaySystemCaller) PoolsIndex(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "PoolsIndex", arg0)
	return *ret0, err
}

// PoolsIndex is a free data retrieval call binding the contract method 0x6d4347da.
//
// Solidity: function PoolsIndex(uint256 ) constant returns(address)
func (_MicroPaySystem *MicroPaySystemSession) PoolsIndex(arg0 *big.Int) (common.Address, error) {
	return _MicroPaySystem.Contract.PoolsIndex(&_MicroPaySystem.CallOpts, arg0)
}

// PoolsIndex is a free data retrieval call binding the contract method 0x6d4347da.
//
// Solidity: function PoolsIndex(uint256 ) constant returns(address)
func (_MicroPaySystem *MicroPaySystemCallerSession) PoolsIndex(arg0 *big.Int) (common.Address, error) {
	return _MicroPaySystem.Contract.PoolsIndex(&_MicroPaySystem.CallOpts, arg0)
}

// RecoveredAddr is a free data retrieval call binding the contract method 0xdb30b27f.
//
// Solidity: function RecoveredAddr() constant returns(address)
func (_MicroPaySystem *MicroPaySystemCaller) RecoveredAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "RecoveredAddr")
	return *ret0, err
}

// RecoveredAddr is a free data retrieval call binding the contract method 0xdb30b27f.
//
// Solidity: function RecoveredAddr() constant returns(address)
func (_MicroPaySystem *MicroPaySystemSession) RecoveredAddr() (common.Address, error) {
	return _MicroPaySystem.Contract.RecoveredAddr(&_MicroPaySystem.CallOpts)
}

// RecoveredAddr is a free data retrieval call binding the contract method 0xdb30b27f.
//
// Solidity: function RecoveredAddr() constant returns(address)
func (_MicroPaySystem *MicroPaySystemCallerSession) RecoveredAddr() (common.Address, error) {
	return _MicroPaySystem.Contract.RecoveredAddr(&_MicroPaySystem.CallOpts)
}

// TokenBalance is a free data retrieval call binding the contract method 0xb6e2b395.
//
// Solidity: function TokenBalance(address userAddress) constant returns(uint256, uint256, uint256)
func (_MicroPaySystem *MicroPaySystemCaller) TokenBalance(opts *bind.CallOpts, userAddress common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _MicroPaySystem.contract.Call(opts, out, "TokenBalance", userAddress)
	return *ret0, *ret1, *ret2, err
}

// TokenBalance is a free data retrieval call binding the contract method 0xb6e2b395.
//
// Solidity: function TokenBalance(address userAddress) constant returns(uint256, uint256, uint256)
func (_MicroPaySystem *MicroPaySystemSession) TokenBalance(userAddress common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _MicroPaySystem.Contract.TokenBalance(&_MicroPaySystem.CallOpts, userAddress)
}

// TokenBalance is a free data retrieval call binding the contract method 0xb6e2b395.
//
// Solidity: function TokenBalance(address userAddress) constant returns(uint256, uint256, uint256)
func (_MicroPaySystem *MicroPaySystemCallerSession) TokenBalance(userAddress common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _MicroPaySystem.Contract.TokenBalance(&_MicroPaySystem.CallOpts, userAddress)
}

// UserData is a free data retrieval call binding the contract method 0x5a903303.
//
// Solidity: function UserData(address , address ) constant returns(uint256 nonce, uint256 tokenBalance, uint256 remindPacket, uint256 expiration, uint256 epoch, uint256 claimedAmount, uint256 claimedMicNonce)
func (_MicroPaySystem *MicroPaySystemCaller) UserData(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	Nonce           *big.Int
	TokenBalance    *big.Int
	RemindPacket    *big.Int
	Expiration      *big.Int
	Epoch           *big.Int
	ClaimedAmount   *big.Int
	ClaimedMicNonce *big.Int
}, error) {
	ret := new(struct {
		Nonce           *big.Int
		TokenBalance    *big.Int
		RemindPacket    *big.Int
		Expiration      *big.Int
		Epoch           *big.Int
		ClaimedAmount   *big.Int
		ClaimedMicNonce *big.Int
	})
	out := ret
	err := _MicroPaySystem.contract.Call(opts, out, "UserData", arg0, arg1)
	return *ret, err
}

// UserData is a free data retrieval call binding the contract method 0x5a903303.
//
// Solidity: function UserData(address , address ) constant returns(uint256 nonce, uint256 tokenBalance, uint256 remindPacket, uint256 expiration, uint256 epoch, uint256 claimedAmount, uint256 claimedMicNonce)
func (_MicroPaySystem *MicroPaySystemSession) UserData(arg0 common.Address, arg1 common.Address) (struct {
	Nonce           *big.Int
	TokenBalance    *big.Int
	RemindPacket    *big.Int
	Expiration      *big.Int
	Epoch           *big.Int
	ClaimedAmount   *big.Int
	ClaimedMicNonce *big.Int
}, error) {
	return _MicroPaySystem.Contract.UserData(&_MicroPaySystem.CallOpts, arg0, arg1)
}

// UserData is a free data retrieval call binding the contract method 0x5a903303.
//
// Solidity: function UserData(address , address ) constant returns(uint256 nonce, uint256 tokenBalance, uint256 remindPacket, uint256 expiration, uint256 epoch, uint256 claimedAmount, uint256 claimedMicNonce)
func (_MicroPaySystem *MicroPaySystemCallerSession) UserData(arg0 common.Address, arg1 common.Address) (struct {
	Nonce           *big.Int
	TokenBalance    *big.Int
	RemindPacket    *big.Int
	Expiration      *big.Int
	Epoch           *big.Int
	ClaimedAmount   *big.Int
	ClaimedMicNonce *big.Int
}, error) {
	return _MicroPaySystem.Contract.UserData(&_MicroPaySystem.CallOpts, arg0, arg1)
}

// Decimal is a free data retrieval call binding the contract method 0x76809ce3.
//
// Solidity: function decimal() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemCaller) Decimal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "decimal")
	return *ret0, err
}

// Decimal is a free data retrieval call binding the contract method 0x76809ce3.
//
// Solidity: function decimal() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemSession) Decimal() (*big.Int, error) {
	return _MicroPaySystem.Contract.Decimal(&_MicroPaySystem.CallOpts)
}

// Decimal is a free data retrieval call binding the contract method 0x76809ce3.
//
// Solidity: function decimal() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemCallerSession) Decimal() (*big.Int, error) {
	return _MicroPaySystem.Contract.Decimal(&_MicroPaySystem.CallOpts)
}

// MinersUnderPool is a free data retrieval call binding the contract method 0x5c4265cb.
//
// Solidity: function minersUnderPool(address , uint256 ) constant returns(bytes32)
func (_MicroPaySystem *MicroPaySystemCaller) MinersUnderPool(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "minersUnderPool", arg0, arg1)
	return *ret0, err
}

// MinersUnderPool is a free data retrieval call binding the contract method 0x5c4265cb.
//
// Solidity: function minersUnderPool(address , uint256 ) constant returns(bytes32)
func (_MicroPaySystem *MicroPaySystemSession) MinersUnderPool(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _MicroPaySystem.Contract.MinersUnderPool(&_MicroPaySystem.CallOpts, arg0, arg1)
}

// MinersUnderPool is a free data retrieval call binding the contract method 0x5c4265cb.
//
// Solidity: function minersUnderPool(address , uint256 ) constant returns(bytes32)
func (_MicroPaySystem *MicroPaySystemCallerSession) MinersUnderPool(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _MicroPaySystem.Contract.MinersUnderPool(&_MicroPaySystem.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MicroPaySystem *MicroPaySystemCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MicroPaySystem *MicroPaySystemSession) Owner() (common.Address, error) {
	return _MicroPaySystem.Contract.Owner(&_MicroPaySystem.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MicroPaySystem *MicroPaySystemCallerSession) Owner() (common.Address, error) {
	return _MicroPaySystem.Contract.Owner(&_MicroPaySystem.CallOpts)
}

// PoolsUsedByUser is a free data retrieval call binding the contract method 0x41b0012b.
//
// Solidity: function poolsUsedByUser(address , uint256 ) constant returns(address)
func (_MicroPaySystem *MicroPaySystemCaller) PoolsUsedByUser(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "poolsUsedByUser", arg0, arg1)
	return *ret0, err
}

// PoolsUsedByUser is a free data retrieval call binding the contract method 0x41b0012b.
//
// Solidity: function poolsUsedByUser(address , uint256 ) constant returns(address)
func (_MicroPaySystem *MicroPaySystemSession) PoolsUsedByUser(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _MicroPaySystem.Contract.PoolsUsedByUser(&_MicroPaySystem.CallOpts, arg0, arg1)
}

// PoolsUsedByUser is a free data retrieval call binding the contract method 0x41b0012b.
//
// Solidity: function poolsUsedByUser(address , uint256 ) constant returns(address)
func (_MicroPaySystem *MicroPaySystemCallerSession) PoolsUsedByUser(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _MicroPaySystem.Contract.PoolsUsedByUser(&_MicroPaySystem.CallOpts, arg0, arg1)
}

// Setting is a free data retrieval call binding the contract method 0xc0c9ce30.
//
// Solidity: function setting() constant returns(uint256 MBytesPerToken, uint256 Duration, uint256 MinPoolGuarantee, uint256 MinMinerGuarantee)
func (_MicroPaySystem *MicroPaySystemCaller) Setting(opts *bind.CallOpts) (struct {
	MBytesPerToken    *big.Int
	Duration          *big.Int
	MinPoolGuarantee  *big.Int
	MinMinerGuarantee *big.Int
}, error) {
	ret := new(struct {
		MBytesPerToken    *big.Int
		Duration          *big.Int
		MinPoolGuarantee  *big.Int
		MinMinerGuarantee *big.Int
	})
	out := ret
	err := _MicroPaySystem.contract.Call(opts, out, "setting")
	return *ret, err
}

// Setting is a free data retrieval call binding the contract method 0xc0c9ce30.
//
// Solidity: function setting() constant returns(uint256 MBytesPerToken, uint256 Duration, uint256 MinPoolGuarantee, uint256 MinMinerGuarantee)
func (_MicroPaySystem *MicroPaySystemSession) Setting() (struct {
	MBytesPerToken    *big.Int
	Duration          *big.Int
	MinPoolGuarantee  *big.Int
	MinMinerGuarantee *big.Int
}, error) {
	return _MicroPaySystem.Contract.Setting(&_MicroPaySystem.CallOpts)
}

// Setting is a free data retrieval call binding the contract method 0xc0c9ce30.
//
// Solidity: function setting() constant returns(uint256 MBytesPerToken, uint256 Duration, uint256 MinPoolGuarantee, uint256 MinMinerGuarantee)
func (_MicroPaySystem *MicroPaySystemCallerSession) Setting() (struct {
	MBytesPerToken    *big.Int
	Duration          *big.Int
	MinPoolGuarantee  *big.Int
	MinMinerGuarantee *big.Int
}, error) {
	return _MicroPaySystem.Contract.Setting(&_MicroPaySystem.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_MicroPaySystem *MicroPaySystemCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_MicroPaySystem *MicroPaySystemSession) Token() (common.Address, error) {
	return _MicroPaySystem.Contract.Token(&_MicroPaySystem.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_MicroPaySystem *MicroPaySystemCallerSession) Token() (common.Address, error) {
	return _MicroPaySystem.Contract.Token(&_MicroPaySystem.CallOpts)
}

// UsersOfPool is a free data retrieval call binding the contract method 0x9f8bee88.
//
// Solidity: function usersOfPool(address pool) constant returns(address[])
func (_MicroPaySystem *MicroPaySystemCaller) UsersOfPool(opts *bind.CallOpts, pool common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "usersOfPool", pool)
	return *ret0, err
}

// UsersOfPool is a free data retrieval call binding the contract method 0x9f8bee88.
//
// Solidity: function usersOfPool(address pool) constant returns(address[])
func (_MicroPaySystem *MicroPaySystemSession) UsersOfPool(pool common.Address) ([]common.Address, error) {
	return _MicroPaySystem.Contract.UsersOfPool(&_MicroPaySystem.CallOpts, pool)
}

// UsersOfPool is a free data retrieval call binding the contract method 0x9f8bee88.
//
// Solidity: function usersOfPool(address pool) constant returns(address[])
func (_MicroPaySystem *MicroPaySystemCallerSession) UsersOfPool(pool common.Address) ([]common.Address, error) {
	return _MicroPaySystem.Contract.UsersOfPool(&_MicroPaySystem.CallOpts, pool)
}

// UsersUnderPool is a free data retrieval call binding the contract method 0x4e7ffecc.
//
// Solidity: function usersUnderPool(address , uint256 ) constant returns(address)
func (_MicroPaySystem *MicroPaySystemCaller) UsersUnderPool(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "usersUnderPool", arg0, arg1)
	return *ret0, err
}

// UsersUnderPool is a free data retrieval call binding the contract method 0x4e7ffecc.
//
// Solidity: function usersUnderPool(address , uint256 ) constant returns(address)
func (_MicroPaySystem *MicroPaySystemSession) UsersUnderPool(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _MicroPaySystem.Contract.UsersUnderPool(&_MicroPaySystem.CallOpts, arg0, arg1)
}

// UsersUnderPool is a free data retrieval call binding the contract method 0x4e7ffecc.
//
// Solidity: function usersUnderPool(address , uint256 ) constant returns(address)
func (_MicroPaySystem *MicroPaySystemCallerSession) UsersUnderPool(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _MicroPaySystem.Contract.UsersUnderPool(&_MicroPaySystem.CallOpts, arg0, arg1)
}

// Vm is a free data retrieval call binding the contract method 0x3a768463.
//
// Solidity: function vm() constant returns(uint256 versionOfPoolData, uint256 versionOfMinerData, uint256 versionOfUserData, uint256 versionOfSysSetting)
func (_MicroPaySystem *MicroPaySystemCaller) Vm(opts *bind.CallOpts) (struct {
	VersionOfPoolData   *big.Int
	VersionOfMinerData  *big.Int
	VersionOfUserData   *big.Int
	VersionOfSysSetting *big.Int
}, error) {
	ret := new(struct {
		VersionOfPoolData   *big.Int
		VersionOfMinerData  *big.Int
		VersionOfUserData   *big.Int
		VersionOfSysSetting *big.Int
	})
	out := ret
	err := _MicroPaySystem.contract.Call(opts, out, "vm")
	return *ret, err
}

// Vm is a free data retrieval call binding the contract method 0x3a768463.
//
// Solidity: function vm() constant returns(uint256 versionOfPoolData, uint256 versionOfMinerData, uint256 versionOfUserData, uint256 versionOfSysSetting)
func (_MicroPaySystem *MicroPaySystemSession) Vm() (struct {
	VersionOfPoolData   *big.Int
	VersionOfMinerData  *big.Int
	VersionOfUserData   *big.Int
	VersionOfSysSetting *big.Int
}, error) {
	return _MicroPaySystem.Contract.Vm(&_MicroPaySystem.CallOpts)
}

// Vm is a free data retrieval call binding the contract method 0x3a768463.
//
// Solidity: function vm() constant returns(uint256 versionOfPoolData, uint256 versionOfMinerData, uint256 versionOfUserData, uint256 versionOfSysSetting)
func (_MicroPaySystem *MicroPaySystemCallerSession) Vm() (struct {
	VersionOfPoolData   *big.Int
	VersionOfMinerData  *big.Int
	VersionOfUserData   *big.Int
	VersionOfSysSetting *big.Int
}, error) {
	return _MicroPaySystem.Contract.Vm(&_MicroPaySystem.CallOpts)
}

// BuyPacket is a paid mutator transaction binding the contract method 0x51f922ac.
//
// Solidity: function BuyPacket(address user, uint256 tokenNo, address poolAddr) returns()
func (_MicroPaySystem *MicroPaySystemTransactor) BuyPacket(opts *bind.TransactOpts, user common.Address, tokenNo *big.Int, poolAddr common.Address) (*types.Transaction, error) {
	return _MicroPaySystem.contract.Transact(opts, "BuyPacket", user, tokenNo, poolAddr)
}

// BuyPacket is a paid mutator transaction binding the contract method 0x51f922ac.
//
// Solidity: function BuyPacket(address user, uint256 tokenNo, address poolAddr) returns()
func (_MicroPaySystem *MicroPaySystemSession) BuyPacket(user common.Address, tokenNo *big.Int, poolAddr common.Address) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.BuyPacket(&_MicroPaySystem.TransactOpts, user, tokenNo, poolAddr)
}

// BuyPacket is a paid mutator transaction binding the contract method 0x51f922ac.
//
// Solidity: function BuyPacket(address user, uint256 tokenNo, address poolAddr) returns()
func (_MicroPaySystem *MicroPaySystemTransactorSession) BuyPacket(user common.Address, tokenNo *big.Int, poolAddr common.Address) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.BuyPacket(&_MicroPaySystem.TransactOpts, user, tokenNo, poolAddr)
}

// ChangeBasicInfo is a paid mutator transaction binding the contract method 0x404d39ba.
//
// Solidity: function ChangeBasicInfo(address mainAddr, bytes name, bytes email, bytes url) returns()
func (_MicroPaySystem *MicroPaySystemTransactor) ChangeBasicInfo(opts *bind.TransactOpts, mainAddr common.Address, name []byte, email []byte, url []byte) (*types.Transaction, error) {
	return _MicroPaySystem.contract.Transact(opts, "ChangeBasicInfo", mainAddr, name, email, url)
}

// ChangeBasicInfo is a paid mutator transaction binding the contract method 0x404d39ba.
//
// Solidity: function ChangeBasicInfo(address mainAddr, bytes name, bytes email, bytes url) returns()
func (_MicroPaySystem *MicroPaySystemSession) ChangeBasicInfo(mainAddr common.Address, name []byte, email []byte, url []byte) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.ChangeBasicInfo(&_MicroPaySystem.TransactOpts, mainAddr, name, email, url)
}

// ChangeBasicInfo is a paid mutator transaction binding the contract method 0x404d39ba.
//
// Solidity: function ChangeBasicInfo(address mainAddr, bytes name, bytes email, bytes url) returns()
func (_MicroPaySystem *MicroPaySystemTransactorSession) ChangeBasicInfo(mainAddr common.Address, name []byte, email []byte, url []byte) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.ChangeBasicInfo(&_MicroPaySystem.TransactOpts, mainAddr, name, email, url)
}

// ChangePool is a paid mutator transaction binding the contract method 0xd50301b9.
//
// Solidity: function ChangePool(address from, address to, bytes32 subAddr) returns()
func (_MicroPaySystem *MicroPaySystemTransactor) ChangePool(opts *bind.TransactOpts, from common.Address, to common.Address, subAddr [32]byte) (*types.Transaction, error) {
	return _MicroPaySystem.contract.Transact(opts, "ChangePool", from, to, subAddr)
}

// ChangePool is a paid mutator transaction binding the contract method 0xd50301b9.
//
// Solidity: function ChangePool(address from, address to, bytes32 subAddr) returns()
func (_MicroPaySystem *MicroPaySystemSession) ChangePool(from common.Address, to common.Address, subAddr [32]byte) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.ChangePool(&_MicroPaySystem.TransactOpts, from, to, subAddr)
}

// ChangePool is a paid mutator transaction binding the contract method 0xd50301b9.
//
// Solidity: function ChangePool(address from, address to, bytes32 subAddr) returns()
func (_MicroPaySystem *MicroPaySystemTransactorSession) ChangePool(from common.Address, to common.Address, subAddr [32]byte) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.ChangePool(&_MicroPaySystem.TransactOpts, from, to, subAddr)
}

// Claim is a paid mutator transaction binding the contract method 0xa0e81a28.
//
// Solidity: function Claim(address user, address pool, uint256 credit, uint256 amount, uint256 micNonce, uint256 cn, bytes signature) returns()
func (_MicroPaySystem *MicroPaySystemTransactor) Claim(opts *bind.TransactOpts, user common.Address, pool common.Address, credit *big.Int, amount *big.Int, micNonce *big.Int, cn *big.Int, signature []byte) (*types.Transaction, error) {
	return _MicroPaySystem.contract.Transact(opts, "Claim", user, pool, credit, amount, micNonce, cn, signature)
}

// Claim is a paid mutator transaction binding the contract method 0xa0e81a28.
//
// Solidity: function Claim(address user, address pool, uint256 credit, uint256 amount, uint256 micNonce, uint256 cn, bytes signature) returns()
func (_MicroPaySystem *MicroPaySystemSession) Claim(user common.Address, pool common.Address, credit *big.Int, amount *big.Int, micNonce *big.Int, cn *big.Int, signature []byte) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.Claim(&_MicroPaySystem.TransactOpts, user, pool, credit, amount, micNonce, cn, signature)
}

// Claim is a paid mutator transaction binding the contract method 0xa0e81a28.
//
// Solidity: function Claim(address user, address pool, uint256 credit, uint256 amount, uint256 micNonce, uint256 cn, bytes signature) returns()
func (_MicroPaySystem *MicroPaySystemTransactorSession) Claim(user common.Address, pool common.Address, credit *big.Int, amount *big.Int, micNonce *big.Int, cn *big.Int, signature []byte) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.Claim(&_MicroPaySystem.TransactOpts, user, pool, credit, amount, micNonce, cn, signature)
}

// DestroyPool is a paid mutator transaction binding the contract method 0x2e0ee593.
//
// Solidity: function DestroyPool(address poolAddr) returns()
func (_MicroPaySystem *MicroPaySystemTransactor) DestroyPool(opts *bind.TransactOpts, poolAddr common.Address) (*types.Transaction, error) {
	return _MicroPaySystem.contract.Transact(opts, "DestroyPool", poolAddr)
}

// DestroyPool is a paid mutator transaction binding the contract method 0x2e0ee593.
//
// Solidity: function DestroyPool(address poolAddr) returns()
func (_MicroPaySystem *MicroPaySystemSession) DestroyPool(poolAddr common.Address) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.DestroyPool(&_MicroPaySystem.TransactOpts, poolAddr)
}

// DestroyPool is a paid mutator transaction binding the contract method 0x2e0ee593.
//
// Solidity: function DestroyPool(address poolAddr) returns()
func (_MicroPaySystem *MicroPaySystemTransactorSession) DestroyPool(poolAddr common.Address) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.DestroyPool(&_MicroPaySystem.TransactOpts, poolAddr)
}

// JoinPool is a paid mutator transaction binding the contract method 0x41bad6c2.
//
// Solidity: function JoinPool(address pool, uint256 tokenNo, bytes32 subAddr, bytes2 zone) returns()
func (_MicroPaySystem *MicroPaySystemTransactor) JoinPool(opts *bind.TransactOpts, pool common.Address, tokenNo *big.Int, subAddr [32]byte, zone [2]byte) (*types.Transaction, error) {
	return _MicroPaySystem.contract.Transact(opts, "JoinPool", pool, tokenNo, subAddr, zone)
}

// JoinPool is a paid mutator transaction binding the contract method 0x41bad6c2.
//
// Solidity: function JoinPool(address pool, uint256 tokenNo, bytes32 subAddr, bytes2 zone) returns()
func (_MicroPaySystem *MicroPaySystemSession) JoinPool(pool common.Address, tokenNo *big.Int, subAddr [32]byte, zone [2]byte) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.JoinPool(&_MicroPaySystem.TransactOpts, pool, tokenNo, subAddr, zone)
}

// JoinPool is a paid mutator transaction binding the contract method 0x41bad6c2.
//
// Solidity: function JoinPool(address pool, uint256 tokenNo, bytes32 subAddr, bytes2 zone) returns()
func (_MicroPaySystem *MicroPaySystemTransactorSession) JoinPool(pool common.Address, tokenNo *big.Int, subAddr [32]byte, zone [2]byte) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.JoinPool(&_MicroPaySystem.TransactOpts, pool, tokenNo, subAddr, zone)
}

// RegPool is a paid mutator transaction binding the contract method 0x4581a00b.
//
// Solidity: function RegPool(uint256 tokenNo, address poolAddr, bytes name, bytes email, bytes url) returns()
func (_MicroPaySystem *MicroPaySystemTransactor) RegPool(opts *bind.TransactOpts, tokenNo *big.Int, poolAddr common.Address, name []byte, email []byte, url []byte) (*types.Transaction, error) {
	return _MicroPaySystem.contract.Transact(opts, "RegPool", tokenNo, poolAddr, name, email, url)
}

// RegPool is a paid mutator transaction binding the contract method 0x4581a00b.
//
// Solidity: function RegPool(uint256 tokenNo, address poolAddr, bytes name, bytes email, bytes url) returns()
func (_MicroPaySystem *MicroPaySystemSession) RegPool(tokenNo *big.Int, poolAddr common.Address, name []byte, email []byte, url []byte) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.RegPool(&_MicroPaySystem.TransactOpts, tokenNo, poolAddr, name, email, url)
}

// RegPool is a paid mutator transaction binding the contract method 0x4581a00b.
//
// Solidity: function RegPool(uint256 tokenNo, address poolAddr, bytes name, bytes email, bytes url) returns()
func (_MicroPaySystem *MicroPaySystemTransactorSession) RegPool(tokenNo *big.Int, poolAddr common.Address, name []byte, email []byte, url []byte) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.RegPool(&_MicroPaySystem.TransactOpts, tokenNo, poolAddr, name, email, url)
}

// RetireFromPool is a paid mutator transaction binding the contract method 0xc480c0d3.
//
// Solidity: function RetireFromPool(bytes32 subAddr) returns()
func (_MicroPaySystem *MicroPaySystemTransactor) RetireFromPool(opts *bind.TransactOpts, subAddr [32]byte) (*types.Transaction, error) {
	return _MicroPaySystem.contract.Transact(opts, "RetireFromPool", subAddr)
}

// RetireFromPool is a paid mutator transaction binding the contract method 0xc480c0d3.
//
// Solidity: function RetireFromPool(bytes32 subAddr) returns()
func (_MicroPaySystem *MicroPaySystemSession) RetireFromPool(subAddr [32]byte) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.RetireFromPool(&_MicroPaySystem.TransactOpts, subAddr)
}

// RetireFromPool is a paid mutator transaction binding the contract method 0xc480c0d3.
//
// Solidity: function RetireFromPool(bytes32 subAddr) returns()
func (_MicroPaySystem *MicroPaySystemTransactorSession) RetireFromPool(subAddr [32]byte) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.RetireFromPool(&_MicroPaySystem.TransactOpts, subAddr)
}

// TestClaimSig is a paid mutator transaction binding the contract method 0xd9233624.
//
// Solidity: function TestClaimSig(address contractAddr, address tokenAddr, address user, address pool, uint256 amount, uint256 micNonce, uint256 cn, bytes signature) returns()
func (_MicroPaySystem *MicroPaySystemTransactor) TestClaimSig(opts *bind.TransactOpts, contractAddr common.Address, tokenAddr common.Address, user common.Address, pool common.Address, amount *big.Int, micNonce *big.Int, cn *big.Int, signature []byte) (*types.Transaction, error) {
	return _MicroPaySystem.contract.Transact(opts, "TestClaimSig", contractAddr, tokenAddr, user, pool, amount, micNonce, cn, signature)
}

// TestClaimSig is a paid mutator transaction binding the contract method 0xd9233624.
//
// Solidity: function TestClaimSig(address contractAddr, address tokenAddr, address user, address pool, uint256 amount, uint256 micNonce, uint256 cn, bytes signature) returns()
func (_MicroPaySystem *MicroPaySystemSession) TestClaimSig(contractAddr common.Address, tokenAddr common.Address, user common.Address, pool common.Address, amount *big.Int, micNonce *big.Int, cn *big.Int, signature []byte) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.TestClaimSig(&_MicroPaySystem.TransactOpts, contractAddr, tokenAddr, user, pool, amount, micNonce, cn, signature)
}

// TestClaimSig is a paid mutator transaction binding the contract method 0xd9233624.
//
// Solidity: function TestClaimSig(address contractAddr, address tokenAddr, address user, address pool, uint256 amount, uint256 micNonce, uint256 cn, bytes signature) returns()
func (_MicroPaySystem *MicroPaySystemTransactorSession) TestClaimSig(contractAddr common.Address, tokenAddr common.Address, user common.Address, pool common.Address, amount *big.Int, micNonce *big.Int, cn *big.Int, signature []byte) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.TestClaimSig(&_MicroPaySystem.TransactOpts, contractAddr, tokenAddr, user, pool, amount, micNonce, cn, signature)
}

// ChangeSetting is a paid mutator transaction binding the contract method 0xb75777c6.
//
// Solidity: function changeSetting(uint256 price, uint256 duration, uint256 pGTN, uint256 mGTN) returns()
func (_MicroPaySystem *MicroPaySystemTransactor) ChangeSetting(opts *bind.TransactOpts, price *big.Int, duration *big.Int, pGTN *big.Int, mGTN *big.Int) (*types.Transaction, error) {
	return _MicroPaySystem.contract.Transact(opts, "changeSetting", price, duration, pGTN, mGTN)
}

// ChangeSetting is a paid mutator transaction binding the contract method 0xb75777c6.
//
// Solidity: function changeSetting(uint256 price, uint256 duration, uint256 pGTN, uint256 mGTN) returns()
func (_MicroPaySystem *MicroPaySystemSession) ChangeSetting(price *big.Int, duration *big.Int, pGTN *big.Int, mGTN *big.Int) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.ChangeSetting(&_MicroPaySystem.TransactOpts, price, duration, pGTN, mGTN)
}

// ChangeSetting is a paid mutator transaction binding the contract method 0xb75777c6.
//
// Solidity: function changeSetting(uint256 price, uint256 duration, uint256 pGTN, uint256 mGTN) returns()
func (_MicroPaySystem *MicroPaySystemTransactorSession) ChangeSetting(price *big.Int, duration *big.Int, pGTN *big.Int, mGTN *big.Int) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.ChangeSetting(&_MicroPaySystem.TransactOpts, price, duration, pGTN, mGTN)
}

// Emergency is a paid mutator transaction binding the contract method 0xcaa6fea4.
//
// Solidity: function emergency() returns()
func (_MicroPaySystem *MicroPaySystemTransactor) Emergency(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MicroPaySystem.contract.Transact(opts, "emergency")
}

// Emergency is a paid mutator transaction binding the contract method 0xcaa6fea4.
//
// Solidity: function emergency() returns()
func (_MicroPaySystem *MicroPaySystemSession) Emergency() (*types.Transaction, error) {
	return _MicroPaySystem.Contract.Emergency(&_MicroPaySystem.TransactOpts)
}

// Emergency is a paid mutator transaction binding the contract method 0xcaa6fea4.
//
// Solidity: function emergency() returns()
func (_MicroPaySystem *MicroPaySystemTransactorSession) Emergency() (*types.Transaction, error) {
	return _MicroPaySystem.Contract.Emergency(&_MicroPaySystem.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MicroPaySystem *MicroPaySystemTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MicroPaySystem.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MicroPaySystem *MicroPaySystemSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.TransferOwnership(&_MicroPaySystem.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MicroPaySystem *MicroPaySystemTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.TransferOwnership(&_MicroPaySystem.TransactOpts, newOwner)
}

// MicroPaySystemPoolClaimIterator is returned from FilterPoolClaim and is used to iterate over the raw logs and unpacked data for PoolClaim events raised by the MicroPaySystem contract.
type MicroPaySystemPoolClaimIterator struct {
	Event *MicroPaySystemPoolClaim // Event containing the contract specifics and raw log

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
func (it *MicroPaySystemPoolClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MicroPaySystemPoolClaim)
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
		it.Event = new(MicroPaySystemPoolClaim)
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
func (it *MicroPaySystemPoolClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MicroPaySystemPoolClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MicroPaySystemPoolClaim represents a PoolClaim event raised by the MicroPaySystem contract.
type MicroPaySystemPoolClaim struct {
	Pool       common.Address
	User       common.Address
	Packet     *big.Int
	Tonken     *big.Int
	MicrNonce  *big.Int
	ClaimNonce *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPoolClaim is a free log retrieval operation binding the contract event 0xe249550089b34fd8690cecdc0bd02ac66196461e455ee12c3024e64b17d7f355.
//
// Solidity: event PoolClaim(address indexed pool, address indexed user, uint256 packet, uint256 tonken, uint256 micrNonce, uint256 claimNonce)
func (_MicroPaySystem *MicroPaySystemFilterer) FilterPoolClaim(opts *bind.FilterOpts, pool []common.Address, user []common.Address) (*MicroPaySystemPoolClaimIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MicroPaySystem.contract.FilterLogs(opts, "PoolClaim", poolRule, userRule)
	if err != nil {
		return nil, err
	}
	return &MicroPaySystemPoolClaimIterator{contract: _MicroPaySystem.contract, event: "PoolClaim", logs: logs, sub: sub}, nil
}

// WatchPoolClaim is a free log subscription operation binding the contract event 0xe249550089b34fd8690cecdc0bd02ac66196461e455ee12c3024e64b17d7f355.
//
// Solidity: event PoolClaim(address indexed pool, address indexed user, uint256 packet, uint256 tonken, uint256 micrNonce, uint256 claimNonce)
func (_MicroPaySystem *MicroPaySystemFilterer) WatchPoolClaim(opts *bind.WatchOpts, sink chan<- *MicroPaySystemPoolClaim, pool []common.Address, user []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MicroPaySystem.contract.WatchLogs(opts, "PoolClaim", poolRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MicroPaySystemPoolClaim)
				if err := _MicroPaySystem.contract.UnpackLog(event, "PoolClaim", log); err != nil {
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

// ParsePoolClaim is a log parse operation binding the contract event 0xe249550089b34fd8690cecdc0bd02ac66196461e455ee12c3024e64b17d7f355.
//
// Solidity: event PoolClaim(address indexed pool, address indexed user, uint256 packet, uint256 tonken, uint256 micrNonce, uint256 claimNonce)
func (_MicroPaySystem *MicroPaySystemFilterer) ParsePoolClaim(log types.Log) (*MicroPaySystemPoolClaim, error) {
	event := new(MicroPaySystemPoolClaim)
	if err := _MicroPaySystem.contract.UnpackLog(event, "PoolClaim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MicroPaySystemPoolDestroyIterator is returned from FilterPoolDestroy and is used to iterate over the raw logs and unpacked data for PoolDestroy events raised by the MicroPaySystem contract.
type MicroPaySystemPoolDestroyIterator struct {
	Event *MicroPaySystemPoolDestroy // Event containing the contract specifics and raw log

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
func (it *MicroPaySystemPoolDestroyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MicroPaySystemPoolDestroy)
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
		it.Event = new(MicroPaySystemPoolDestroy)
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
func (it *MicroPaySystemPoolDestroyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MicroPaySystemPoolDestroyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MicroPaySystemPoolDestroy represents a PoolDestroy event raised by the MicroPaySystem contract.
type MicroPaySystemPoolDestroy struct {
	Pool common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPoolDestroy is a free log retrieval operation binding the contract event 0x1409788ba1874c653fc722a785d59a23c177dc897c9602825f543d5d5fcf5d77.
//
// Solidity: event PoolDestroy(address indexed pool, uint256 time)
func (_MicroPaySystem *MicroPaySystemFilterer) FilterPoolDestroy(opts *bind.FilterOpts, pool []common.Address) (*MicroPaySystemPoolDestroyIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _MicroPaySystem.contract.FilterLogs(opts, "PoolDestroy", poolRule)
	if err != nil {
		return nil, err
	}
	return &MicroPaySystemPoolDestroyIterator{contract: _MicroPaySystem.contract, event: "PoolDestroy", logs: logs, sub: sub}, nil
}

// WatchPoolDestroy is a free log subscription operation binding the contract event 0x1409788ba1874c653fc722a785d59a23c177dc897c9602825f543d5d5fcf5d77.
//
// Solidity: event PoolDestroy(address indexed pool, uint256 time)
func (_MicroPaySystem *MicroPaySystemFilterer) WatchPoolDestroy(opts *bind.WatchOpts, sink chan<- *MicroPaySystemPoolDestroy, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _MicroPaySystem.contract.WatchLogs(opts, "PoolDestroy", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MicroPaySystemPoolDestroy)
				if err := _MicroPaySystem.contract.UnpackLog(event, "PoolDestroy", log); err != nil {
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

// ParsePoolDestroy is a log parse operation binding the contract event 0x1409788ba1874c653fc722a785d59a23c177dc897c9602825f543d5d5fcf5d77.
//
// Solidity: event PoolDestroy(address indexed pool, uint256 time)
func (_MicroPaySystem *MicroPaySystemFilterer) ParsePoolDestroy(log types.Log) (*MicroPaySystemPoolDestroy, error) {
	event := new(MicroPaySystemPoolDestroy)
	if err := _MicroPaySystem.contract.UnpackLog(event, "PoolDestroy", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MicroPaySystemPoolRegisterIterator is returned from FilterPoolRegister and is used to iterate over the raw logs and unpacked data for PoolRegister events raised by the MicroPaySystem contract.
type MicroPaySystemPoolRegisterIterator struct {
	Event *MicroPaySystemPoolRegister // Event containing the contract specifics and raw log

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
func (it *MicroPaySystemPoolRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MicroPaySystemPoolRegister)
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
		it.Event = new(MicroPaySystemPoolRegister)
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
func (it *MicroPaySystemPoolRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MicroPaySystemPoolRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MicroPaySystemPoolRegister represents a PoolRegister event raised by the MicroPaySystem contract.
type MicroPaySystemPoolRegister struct {
	Pool common.Address
	Gtn  *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPoolRegister is a free log retrieval operation binding the contract event 0xdebedec91d819c2c41c9c69a68386a87efefd2cdd9d5031d5e1a24e7d3100c26.
//
// Solidity: event PoolRegister(address indexed pool, uint256 gtn)
func (_MicroPaySystem *MicroPaySystemFilterer) FilterPoolRegister(opts *bind.FilterOpts, pool []common.Address) (*MicroPaySystemPoolRegisterIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _MicroPaySystem.contract.FilterLogs(opts, "PoolRegister", poolRule)
	if err != nil {
		return nil, err
	}
	return &MicroPaySystemPoolRegisterIterator{contract: _MicroPaySystem.contract, event: "PoolRegister", logs: logs, sub: sub}, nil
}

// WatchPoolRegister is a free log subscription operation binding the contract event 0xdebedec91d819c2c41c9c69a68386a87efefd2cdd9d5031d5e1a24e7d3100c26.
//
// Solidity: event PoolRegister(address indexed pool, uint256 gtn)
func (_MicroPaySystem *MicroPaySystemFilterer) WatchPoolRegister(opts *bind.WatchOpts, sink chan<- *MicroPaySystemPoolRegister, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _MicroPaySystem.contract.WatchLogs(opts, "PoolRegister", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MicroPaySystemPoolRegister)
				if err := _MicroPaySystem.contract.UnpackLog(event, "PoolRegister", log); err != nil {
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

// ParsePoolRegister is a log parse operation binding the contract event 0xdebedec91d819c2c41c9c69a68386a87efefd2cdd9d5031d5e1a24e7d3100c26.
//
// Solidity: event PoolRegister(address indexed pool, uint256 gtn)
func (_MicroPaySystem *MicroPaySystemFilterer) ParsePoolRegister(log types.Log) (*MicroPaySystemPoolRegister, error) {
	event := new(MicroPaySystemPoolRegister)
	if err := _MicroPaySystem.contract.UnpackLog(event, "PoolRegister", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MicroPaySystemUserRechargeIterator is returned from FilterUserRecharge and is used to iterate over the raw logs and unpacked data for UserRecharge events raised by the MicroPaySystem contract.
type MicroPaySystemUserRechargeIterator struct {
	Event *MicroPaySystemUserRecharge // Event containing the contract specifics and raw log

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
func (it *MicroPaySystemUserRechargeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MicroPaySystemUserRecharge)
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
		it.Event = new(MicroPaySystemUserRecharge)
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
func (it *MicroPaySystemUserRechargeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MicroPaySystemUserRechargeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MicroPaySystemUserRecharge represents a UserRecharge event raised by the MicroPaySystem contract.
type MicroPaySystemUserRecharge struct {
	From    common.Address
	To      common.Address
	Tokens  *big.Int
	Packets *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUserRecharge is a free log retrieval operation binding the contract event 0xcc8dce95919b1ee4bc195f48f53ce25fc697737c5022951361962fa21fd25195.
//
// Solidity: event UserRecharge(address indexed from, address indexed to, uint256 tokens, uint256 packets)
func (_MicroPaySystem *MicroPaySystemFilterer) FilterUserRecharge(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MicroPaySystemUserRechargeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MicroPaySystem.contract.FilterLogs(opts, "UserRecharge", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MicroPaySystemUserRechargeIterator{contract: _MicroPaySystem.contract, event: "UserRecharge", logs: logs, sub: sub}, nil
}

// WatchUserRecharge is a free log subscription operation binding the contract event 0xcc8dce95919b1ee4bc195f48f53ce25fc697737c5022951361962fa21fd25195.
//
// Solidity: event UserRecharge(address indexed from, address indexed to, uint256 tokens, uint256 packets)
func (_MicroPaySystem *MicroPaySystemFilterer) WatchUserRecharge(opts *bind.WatchOpts, sink chan<- *MicroPaySystemUserRecharge, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MicroPaySystem.contract.WatchLogs(opts, "UserRecharge", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MicroPaySystemUserRecharge)
				if err := _MicroPaySystem.contract.UnpackLog(event, "UserRecharge", log); err != nil {
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

// ParseUserRecharge is a log parse operation binding the contract event 0xcc8dce95919b1ee4bc195f48f53ce25fc697737c5022951361962fa21fd25195.
//
// Solidity: event UserRecharge(address indexed from, address indexed to, uint256 tokens, uint256 packets)
func (_MicroPaySystem *MicroPaySystemFilterer) ParseUserRecharge(log types.Log) (*MicroPaySystemUserRecharge, error) {
	event := new(MicroPaySystemUserRecharge)
	if err := _MicroPaySystem.contract.UnpackLog(event, "UserRecharge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MicroPaySystemUserRefundIterator is returned from FilterUserRefund and is used to iterate over the raw logs and unpacked data for UserRefund events raised by the MicroPaySystem contract.
type MicroPaySystemUserRefundIterator struct {
	Event *MicroPaySystemUserRefund // Event containing the contract specifics and raw log

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
func (it *MicroPaySystemUserRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MicroPaySystemUserRefund)
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
		it.Event = new(MicroPaySystemUserRefund)
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
func (it *MicroPaySystemUserRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MicroPaySystemUserRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MicroPaySystemUserRefund represents a UserRefund event raised by the MicroPaySystem contract.
type MicroPaySystemUserRefund struct {
	User    common.Address
	Pool    common.Address
	Tokens  *big.Int
	Packets *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUserRefund is a free log retrieval operation binding the contract event 0x3e76e2c548d3990883a9eaa236878db2c29d1e90bed321c625673773416e9d58.
//
// Solidity: event UserRefund(address indexed user, address indexed pool, uint256 tokens, uint256 packets)
func (_MicroPaySystem *MicroPaySystemFilterer) FilterUserRefund(opts *bind.FilterOpts, user []common.Address, pool []common.Address) (*MicroPaySystemUserRefundIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _MicroPaySystem.contract.FilterLogs(opts, "UserRefund", userRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &MicroPaySystemUserRefundIterator{contract: _MicroPaySystem.contract, event: "UserRefund", logs: logs, sub: sub}, nil
}

// WatchUserRefund is a free log subscription operation binding the contract event 0x3e76e2c548d3990883a9eaa236878db2c29d1e90bed321c625673773416e9d58.
//
// Solidity: event UserRefund(address indexed user, address indexed pool, uint256 tokens, uint256 packets)
func (_MicroPaySystem *MicroPaySystemFilterer) WatchUserRefund(opts *bind.WatchOpts, sink chan<- *MicroPaySystemUserRefund, user []common.Address, pool []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _MicroPaySystem.contract.WatchLogs(opts, "UserRefund", userRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MicroPaySystemUserRefund)
				if err := _MicroPaySystem.contract.UnpackLog(event, "UserRefund", log); err != nil {
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

// ParseUserRefund is a log parse operation binding the contract event 0x3e76e2c548d3990883a9eaa236878db2c29d1e90bed321c625673773416e9d58.
//
// Solidity: event UserRefund(address indexed user, address indexed pool, uint256 tokens, uint256 packets)
func (_MicroPaySystem *MicroPaySystemFilterer) ParseUserRefund(log types.Log) (*MicroPaySystemUserRefund, error) {
	event := new(MicroPaySystemUserRefund)
	if err := _MicroPaySystem.contract.UnpackLog(event, "UserRefund", log); err != nil {
		return nil, err
	}
	return event, nil
}
