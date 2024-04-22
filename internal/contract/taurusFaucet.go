// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// TaurusFaucetMetaData contains all meta data concerning the TaurusFaucet contract.
var TaurusFaucetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"FaucetDripped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ONE_DAY_SECONDS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"drip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040526201518060015534801561001757600080fd5b5061003461002961003960201b60201c565b61004160201b60201c565b610105565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6109c0806101146000396000f3fe6080604052600436106100595760003560e01c80635c16e15e1461006557806365bcac14146100a2578063715018a6146100cd5780638da5cb5b146100e45780639e353a1e1461010f578063f2fde38b1461013857610060565b3661006057005b600080fd5b34801561007157600080fd5b5061008c600480360381019061008791906105ae565b610161565b60405161009991906107b9565b60405180910390f35b3480156100ae57600080fd5b506100b7610179565b6040516100c491906107b9565b60405180910390f35b3480156100d957600080fd5b506100e261017f565b005b3480156100f057600080fd5b506100f9610193565b60405161010691906106f5565b60405180910390f35b34801561011b57600080fd5b50610136600480360381019061013191906105d7565b6101bc565b005b34801561014457600080fd5b5061015f600480360381019061015a91906105ae565b610379565b005b60026020528060005260406000206000915090505481565b60015481565b6101876103fd565b610191600061047b565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6101c46103fd565b61020c600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461053f565b61024b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024290610759565b60405180910390fd5b60008273ffffffffffffffffffffffffffffffffffffffff1682604051610271906106e0565b60006040518083038185875af1925050503d80600081146102ae576040519150601f19603f3d011682016040523d82523d6000602084013e6102b3565b606091505b50509050806102f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102ee90610799565b60405180910390fd5b42600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507f5c6f8a83250ac55dd524ad7bdf5ad3a59d3ccbe7c1827b0c4ec3615459fd9ae9838360405161036c929190610710565b60405180910390a1505050565b6103816103fd565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156103f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e890610739565b60405180910390fd5b6103fa8161047b565b50565b61040561057c565b73ffffffffffffffffffffffffffffffffffffffff16610423610193565b73ffffffffffffffffffffffffffffffffffffffff1614610479576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047090610779565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000428211156105525760009050610577565b600082116105635760019050610577565b600154824261057291906107f0565b101590505b919050565b600033905090565b6000813590506105938161095c565b92915050565b6000813590506105a881610973565b92915050565b6000602082840312156105c057600080fd5b60006105ce84828501610584565b91505092915050565b600080604083850312156105ea57600080fd5b60006105f885828601610584565b925050602061060985828601610599565b9150509250929050565b61061c81610824565b82525050565b600061062f6026836107df565b915061063a8261088f565b604082019050919050565b6000610652601f836107df565b915061065d826108de565b602082019050919050565b60006106756020836107df565b915061068082610907565b602082019050919050565b60006106986013836107df565b91506106a382610930565b602082019050919050565b60006106bb6000836107d4565b91506106c682610959565b600082019050919050565b6106da81610856565b82525050565b60006106eb826106ae565b9150819050919050565b600060208201905061070a6000830184610613565b92915050565b60006040820190506107256000830185610613565b61073260208301846106d1565b9392505050565b6000602082019050818103600083015261075281610622565b9050919050565b6000602082019050818103600083015261077281610645565b9050919050565b6000602082019050818103600083015261079281610668565b9050919050565b600060208201905081810360008301526107b28161068b565b9050919050565b60006020820190506107ce60008301846106d1565b92915050565b600081905092915050565b600082825260208201905092915050565b60006107fb82610856565b915061080683610856565b92508282101561081957610818610860565b5b828203905092915050565b600061082f82610836565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f48617320636c61696d656420696e20746865206c617374203234686f75727300600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f4661696c6564206472697070696e672041786300000000000000000000000000600082015250565b50565b61096581610824565b811461097057600080fd5b50565b61097c81610856565b811461098757600080fd5b5056fea2646970667358221220e931ab4824309a5d03c1b00256786b82eded1cc8e2efe735c91ec84180212ea564736f6c63430008020033",
}

// TaurusFaucetABI is the input ABI used to generate the binding from.
// Deprecated: Use TaurusFaucetMetaData.ABI instead.
var TaurusFaucetABI = TaurusFaucetMetaData.ABI

// TaurusFaucetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TaurusFaucetMetaData.Bin instead.
var TaurusFaucetBin = TaurusFaucetMetaData.Bin

// DeployTaurusFaucet deploys a new Ethereum contract, binding an instance of TaurusFaucet to it.
func DeployTaurusFaucet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TaurusFaucet, error) {
	parsed, err := TaurusFaucetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TaurusFaucetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TaurusFaucet{TaurusFaucetCaller: TaurusFaucetCaller{contract: contract}, TaurusFaucetTransactor: TaurusFaucetTransactor{contract: contract}, TaurusFaucetFilterer: TaurusFaucetFilterer{contract: contract}}, nil
}

// TaurusFaucet is an auto generated Go binding around an Ethereum contract.
type TaurusFaucet struct {
	TaurusFaucetCaller     // Read-only binding to the contract
	TaurusFaucetTransactor // Write-only binding to the contract
	TaurusFaucetFilterer   // Log filterer for contract events
}

// TaurusFaucetCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaurusFaucetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaurusFaucetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaurusFaucetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaurusFaucetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaurusFaucetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaurusFaucetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaurusFaucetSession struct {
	Contract     *TaurusFaucet     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaurusFaucetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaurusFaucetCallerSession struct {
	Contract *TaurusFaucetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TaurusFaucetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaurusFaucetTransactorSession struct {
	Contract     *TaurusFaucetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TaurusFaucetRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaurusFaucetRaw struct {
	Contract *TaurusFaucet // Generic contract binding to access the raw methods on
}

// TaurusFaucetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaurusFaucetCallerRaw struct {
	Contract *TaurusFaucetCaller // Generic read-only contract binding to access the raw methods on
}

// TaurusFaucetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaurusFaucetTransactorRaw struct {
	Contract *TaurusFaucetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaurusFaucet creates a new instance of TaurusFaucet, bound to a specific deployed contract.
func NewTaurusFaucet(address common.Address, backend bind.ContractBackend) (*TaurusFaucet, error) {
	contract, err := bindTaurusFaucet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaurusFaucet{TaurusFaucetCaller: TaurusFaucetCaller{contract: contract}, TaurusFaucetTransactor: TaurusFaucetTransactor{contract: contract}, TaurusFaucetFilterer: TaurusFaucetFilterer{contract: contract}}, nil
}

// NewTaurusFaucetCaller creates a new read-only instance of TaurusFaucet, bound to a specific deployed contract.
func NewTaurusFaucetCaller(address common.Address, caller bind.ContractCaller) (*TaurusFaucetCaller, error) {
	contract, err := bindTaurusFaucet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaurusFaucetCaller{contract: contract}, nil
}

// NewTaurusFaucetTransactor creates a new write-only instance of TaurusFaucet, bound to a specific deployed contract.
func NewTaurusFaucetTransactor(address common.Address, transactor bind.ContractTransactor) (*TaurusFaucetTransactor, error) {
	contract, err := bindTaurusFaucet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaurusFaucetTransactor{contract: contract}, nil
}

// NewTaurusFaucetFilterer creates a new log filterer instance of TaurusFaucet, bound to a specific deployed contract.
func NewTaurusFaucetFilterer(address common.Address, filterer bind.ContractFilterer) (*TaurusFaucetFilterer, error) {
	contract, err := bindTaurusFaucet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaurusFaucetFilterer{contract: contract}, nil
}

// bindTaurusFaucet binds a generic wrapper to an already deployed contract.
func bindTaurusFaucet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaurusFaucetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaurusFaucet *TaurusFaucetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaurusFaucet.Contract.TaurusFaucetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaurusFaucet *TaurusFaucetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaurusFaucet.Contract.TaurusFaucetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaurusFaucet *TaurusFaucetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaurusFaucet.Contract.TaurusFaucetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaurusFaucet *TaurusFaucetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaurusFaucet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaurusFaucet *TaurusFaucetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaurusFaucet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaurusFaucet *TaurusFaucetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaurusFaucet.Contract.contract.Transact(opts, method, params...)
}

// ONEDAYSECONDS is a free data retrieval call binding the contract method 0x65bcac14.
//
// Solidity: function ONE_DAY_SECONDS() view returns(uint256)
func (_TaurusFaucet *TaurusFaucetCaller) ONEDAYSECONDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaurusFaucet.contract.Call(opts, &out, "ONE_DAY_SECONDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ONEDAYSECONDS is a free data retrieval call binding the contract method 0x65bcac14.
//
// Solidity: function ONE_DAY_SECONDS() view returns(uint256)
func (_TaurusFaucet *TaurusFaucetSession) ONEDAYSECONDS() (*big.Int, error) {
	return _TaurusFaucet.Contract.ONEDAYSECONDS(&_TaurusFaucet.CallOpts)
}

// ONEDAYSECONDS is a free data retrieval call binding the contract method 0x65bcac14.
//
// Solidity: function ONE_DAY_SECONDS() view returns(uint256)
func (_TaurusFaucet *TaurusFaucetCallerSession) ONEDAYSECONDS() (*big.Int, error) {
	return _TaurusFaucet.Contract.ONEDAYSECONDS(&_TaurusFaucet.CallOpts)
}

// LastClaim is a free data retrieval call binding the contract method 0x5c16e15e.
//
// Solidity: function lastClaim(address ) view returns(uint256)
func (_TaurusFaucet *TaurusFaucetCaller) LastClaim(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TaurusFaucet.contract.Call(opts, &out, "lastClaim", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastClaim is a free data retrieval call binding the contract method 0x5c16e15e.
//
// Solidity: function lastClaim(address ) view returns(uint256)
func (_TaurusFaucet *TaurusFaucetSession) LastClaim(arg0 common.Address) (*big.Int, error) {
	return _TaurusFaucet.Contract.LastClaim(&_TaurusFaucet.CallOpts, arg0)
}

// LastClaim is a free data retrieval call binding the contract method 0x5c16e15e.
//
// Solidity: function lastClaim(address ) view returns(uint256)
func (_TaurusFaucet *TaurusFaucetCallerSession) LastClaim(arg0 common.Address) (*big.Int, error) {
	return _TaurusFaucet.Contract.LastClaim(&_TaurusFaucet.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaurusFaucet *TaurusFaucetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaurusFaucet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaurusFaucet *TaurusFaucetSession) Owner() (common.Address, error) {
	return _TaurusFaucet.Contract.Owner(&_TaurusFaucet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaurusFaucet *TaurusFaucetCallerSession) Owner() (common.Address, error) {
	return _TaurusFaucet.Contract.Owner(&_TaurusFaucet.CallOpts)
}

// Drip is a paid mutator transaction binding the contract method 0x9e353a1e.
//
// Solidity: function drip(address _recipient, uint256 _amount) returns()
func (_TaurusFaucet *TaurusFaucetTransactor) Drip(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TaurusFaucet.contract.Transact(opts, "drip", _recipient, _amount)
}

// Drip is a paid mutator transaction binding the contract method 0x9e353a1e.
//
// Solidity: function drip(address _recipient, uint256 _amount) returns()
func (_TaurusFaucet *TaurusFaucetSession) Drip(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TaurusFaucet.Contract.Drip(&_TaurusFaucet.TransactOpts, _recipient, _amount)
}

// Drip is a paid mutator transaction binding the contract method 0x9e353a1e.
//
// Solidity: function drip(address _recipient, uint256 _amount) returns()
func (_TaurusFaucet *TaurusFaucetTransactorSession) Drip(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TaurusFaucet.Contract.Drip(&_TaurusFaucet.TransactOpts, _recipient, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaurusFaucet *TaurusFaucetTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaurusFaucet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaurusFaucet *TaurusFaucetSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaurusFaucet.Contract.RenounceOwnership(&_TaurusFaucet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaurusFaucet *TaurusFaucetTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaurusFaucet.Contract.RenounceOwnership(&_TaurusFaucet.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaurusFaucet *TaurusFaucetTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TaurusFaucet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaurusFaucet *TaurusFaucetSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaurusFaucet.Contract.TransferOwnership(&_TaurusFaucet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaurusFaucet *TaurusFaucetTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaurusFaucet.Contract.TransferOwnership(&_TaurusFaucet.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TaurusFaucet *TaurusFaucetTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaurusFaucet.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TaurusFaucet *TaurusFaucetSession) Receive() (*types.Transaction, error) {
	return _TaurusFaucet.Contract.Receive(&_TaurusFaucet.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TaurusFaucet *TaurusFaucetTransactorSession) Receive() (*types.Transaction, error) {
	return _TaurusFaucet.Contract.Receive(&_TaurusFaucet.TransactOpts)
}

// TaurusFaucetFaucetDrippedIterator is returned from FilterFaucetDripped and is used to iterate over the raw logs and unpacked data for FaucetDripped events raised by the TaurusFaucet contract.
type TaurusFaucetFaucetDrippedIterator struct {
	Event *TaurusFaucetFaucetDripped // Event containing the contract specifics and raw log

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
func (it *TaurusFaucetFaucetDrippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaurusFaucetFaucetDripped)
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
		it.Event = new(TaurusFaucetFaucetDripped)
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
func (it *TaurusFaucetFaucetDrippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaurusFaucetFaucetDrippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaurusFaucetFaucetDripped represents a FaucetDripped event raised by the TaurusFaucet contract.
type TaurusFaucetFaucetDripped struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFaucetDripped is a free log retrieval operation binding the contract event 0x5c6f8a83250ac55dd524ad7bdf5ad3a59d3ccbe7c1827b0c4ec3615459fd9ae9.
//
// Solidity: event FaucetDripped(address _recipient, uint256 _amount)
func (_TaurusFaucet *TaurusFaucetFilterer) FilterFaucetDripped(opts *bind.FilterOpts) (*TaurusFaucetFaucetDrippedIterator, error) {

	logs, sub, err := _TaurusFaucet.contract.FilterLogs(opts, "FaucetDripped")
	if err != nil {
		return nil, err
	}
	return &TaurusFaucetFaucetDrippedIterator{contract: _TaurusFaucet.contract, event: "FaucetDripped", logs: logs, sub: sub}, nil
}

// WatchFaucetDripped is a free log subscription operation binding the contract event 0x5c6f8a83250ac55dd524ad7bdf5ad3a59d3ccbe7c1827b0c4ec3615459fd9ae9.
//
// Solidity: event FaucetDripped(address _recipient, uint256 _amount)
func (_TaurusFaucet *TaurusFaucetFilterer) WatchFaucetDripped(opts *bind.WatchOpts, sink chan<- *TaurusFaucetFaucetDripped) (event.Subscription, error) {

	logs, sub, err := _TaurusFaucet.contract.WatchLogs(opts, "FaucetDripped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaurusFaucetFaucetDripped)
				if err := _TaurusFaucet.contract.UnpackLog(event, "FaucetDripped", log); err != nil {
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

// ParseFaucetDripped is a log parse operation binding the contract event 0x5c6f8a83250ac55dd524ad7bdf5ad3a59d3ccbe7c1827b0c4ec3615459fd9ae9.
//
// Solidity: event FaucetDripped(address _recipient, uint256 _amount)
func (_TaurusFaucet *TaurusFaucetFilterer) ParseFaucetDripped(log types.Log) (*TaurusFaucetFaucetDripped, error) {
	event := new(TaurusFaucetFaucetDripped)
	if err := _TaurusFaucet.contract.UnpackLog(event, "FaucetDripped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaurusFaucetOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TaurusFaucet contract.
type TaurusFaucetOwnershipTransferredIterator struct {
	Event *TaurusFaucetOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TaurusFaucetOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaurusFaucetOwnershipTransferred)
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
		it.Event = new(TaurusFaucetOwnershipTransferred)
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
func (it *TaurusFaucetOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaurusFaucetOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaurusFaucetOwnershipTransferred represents a OwnershipTransferred event raised by the TaurusFaucet contract.
type TaurusFaucetOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaurusFaucet *TaurusFaucetFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TaurusFaucetOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaurusFaucet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TaurusFaucetOwnershipTransferredIterator{contract: _TaurusFaucet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaurusFaucet *TaurusFaucetFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TaurusFaucetOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaurusFaucet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaurusFaucetOwnershipTransferred)
				if err := _TaurusFaucet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TaurusFaucet *TaurusFaucetFilterer) ParseOwnershipTransferred(log types.Log) (*TaurusFaucetOwnershipTransferred, error) {
	event := new(TaurusFaucetOwnershipTransferred)
	if err := _TaurusFaucet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
