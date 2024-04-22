package internal

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/axiomesh/faucet/global"
	"github.com/axiomesh/faucet/internal/contract"
)

func sendTxAxm(c *Client, toAddr string, amount float64) (string, error) {
	c.axiomLock.Lock()
	defer c.axiomLock.Unlock()
	client := c.axiomClient

	fromAddress := c.axiomAuth.From
	// 余额查询
	balanceNow, err := client.BalanceAt(context.Background(), common.HexToAddress(toAddr), nil)
	if err != nil {
		c.logger.Error(err)
		return "", err
	}
	limit := floatToEtherBigInt(c.Config.Axiom.ClaimLimit)
	if balanceNow.Cmp(limit) >= 0 {
		return "", fmt.Errorf(global.EnoughTokenMsg)
	}

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		c.logger.Error(err)
		return "", err
	}

	value := floatToEtherBigInt(amount)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return "", err
	}
	taurusFaucet, err := contract.NewTaurusFaucet(common.HexToAddress(c.Config.Axiom.FaucetAddr), c.axiomClient)
	if err != nil {
		return "", err
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(contract.TaurusFaucetABI)))
	if err != nil {
		return "", err
	}

	input, err := contractAbi.Pack("drip", common.HexToAddress(toAddr), value)
	if err != nil {
		return "", err
	}
	contractAddress := common.HexToAddress(c.Config.Axiom.FaucetAddr)

	msg := ethereum.CallMsg{
		From: fromAddress,
		To:   &contractAddress,
		Data: input,
	}
	_, err = client.CallContract(context.Background(), msg, nil)
	if err != nil {
		c.logger.Error(err)
		return "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(c.axiomPrivateKey, chainId)
	if err != nil {
		return "", err
	}
	gasTipCap, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		return "", err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(c.Config.Axiom.GasLimit)
	auth.GasFeeCap = new(big.Int).Mul(gasPrice, big.NewInt(2))
	auth.GasTipCap = gasTipCap

	tx, err := taurusFaucet.Drip(auth, common.HexToAddress(toAddr), value)
	if err != nil {
		c.logger.Error(err)
		return "", err
	}

	c.logger.Infof("axm tx sent: %s", tx.Hash().Hex())

	return tx.Hash().Hex(), nil
}

func checkBalance(c *Client, toAddr string) (bool, error) {
	client := c.axiomClient
	// 余额查询
	balanceNow, err := client.BalanceAt(context.Background(), common.HexToAddress(toAddr), nil)
	if err != nil {
		c.logger.Error(err)
		return false, err
	}
	limit := floatToEtherBigInt(c.Config.Axiom.ClaimLimit)
	if balanceNow.Cmp(limit) >= 0 {
		return false, fmt.Errorf(global.EnoughTokenMsg)
	}
	return true, nil
}

func floatToEtherBigInt(value float64) *big.Int {
	decimalMultiplier := new(big.Int)
	decimalMultiplier.Exp(big.NewInt(10), big.NewInt(18), nil)

	valueAsBigFloat := new(big.Float).SetFloat64(value)
	valueAsBigFloat.Mul(valueAsBigFloat, new(big.Float).SetInt(decimalMultiplier))

	etherBigInt := new(big.Int)
	valueAsBigFloat.Int(etherBigInt)

	return etherBigInt
}
