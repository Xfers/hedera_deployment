package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

func broadcast_mint_transaction_mainnet(txBytes string, signature1 []byte, signature2 []byte, signature3 []byte) {
	var _txBytes []byte

	deployerAccountId, err := hedera.AccountIDFromString(os.Getenv("HEDERA_XSGD_MAINNET_DEPLOYER_ACCOUNT_ID"))
	if err != nil {
		panic(err)
	}

	deployerPublicKey, err := hedera.PublicKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_DEPLOYER_PUBLIC_KEY"))
	if err != nil {
		panic(err)
	}

	_ = deployerPublicKey

	deployerPrivateKey, err := hedera.PrivateKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_DEPLOYER_PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}

	supplyPublicKey1, err := hedera.PublicKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_SUPPLY_PUBLIC_KEY1"))
	if err != nil {
		panic(err)
	}

	supplyPublicKey2, err := hedera.PublicKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_SUPPLY_PUBLIC_KEY2"))
	if err != nil {
		panic(err)
	}

	client := hedera.ClientForMainnet()
	client.SetOperator(deployerAccountId, deployerPrivateKey)

	_txBytes, err = hex.DecodeString(txBytes)
	if err != nil {
		fmt.Println("error with txBytes")
	}

	var tx interface{}
	//var createTokenTx hedera.TokenCreateTransaction
	var mintTokenTx hedera.TokenMintTransaction
	var signedMintTokenTx *hedera.TokenMintTransaction

	tx, err = hedera.TransactionFromBytes(_txBytes)
	if err != nil {
		fmt.Println("error with txBytes")
		panic(err)
	}

	mintTokenTx = tx.(hedera.TokenMintTransaction)
	signedMintTokenTx = mintTokenTx.AddSignature(supplyPublicKey1, signature1).AddSignature(supplyPublicKey2, signature1)

	response, err := signedMintTokenTx.Execute(client)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", "response")
	fmt.Println(response)
	receipt, err := response.GetReceipt(client)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", "receipt")
	fmt.Printf("%+v\n", receipt)

	//Get the token ID from the receipt
	tokenId := *receipt.TokenID

	fmt.Printf("The new token ID is %v\n", tokenId)

}

func broadcast_burn_transaction_mainnet(txBytes string, signature1 []byte, signature2 []byte, signature3 []byte) {
	var _txBytes []byte

	deployerAccountId, err := hedera.AccountIDFromString(os.Getenv("HEDERA_XSGD_MAINNET_DEPLOYER_ACCOUNT_ID"))
	if err != nil {
		panic(err)
	}

	deployerPublicKey, err := hedera.PublicKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_DEPLOYER_PUBLIC_KEY"))
	if err != nil {
		panic(err)
	}

	_ = deployerPublicKey

	deployerPrivateKey, err := hedera.PrivateKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_DEPLOYER_PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}

	supplyPublicKey1, err := hedera.PublicKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_SUPPLY_PUBLIC_KEY1"))
	if err != nil {
		panic(err)
	}

	supplyPublicKey2, err := hedera.PublicKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_SUPPLY_PUBLIC_KEY2"))
	if err != nil {
		panic(err)
	}

	client := hedera.ClientForMainnet()
	client.SetOperator(deployerAccountId, deployerPrivateKey)

	_txBytes, err = hex.DecodeString(txBytes)
	if err != nil {
		fmt.Println("error with txBytes")
	}

	var tx interface{}
	//var createTokenTx hedera.TokenCreateTransaction
	var burnTokenTx hedera.TokenBurnTransaction
	var signedBurnTokenTx *hedera.TokenBurnTransaction

	tx, err = hedera.TransactionFromBytes(_txBytes)
	if err != nil {
		fmt.Println("error with txBytes")
		panic(err)
	}

	burnTokenTx = tx.(hedera.TokenBurnTransaction)
	signedBurnTokenTx = burnTokenTx.AddSignature(supplyPublicKey1, signature1).AddSignature(supplyPublicKey2, signature1)

	response, err := signedBurnTokenTx.Execute(client)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", "response")
	fmt.Println(response)
	receipt, err := response.GetReceipt(client)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", "receipt")
	fmt.Printf("%+v\n", receipt)

	//Get the token ID from the receipt
	tokenId := *receipt.TokenID

	fmt.Printf("The new token ID is %v\n", tokenId)

}

func broadcast_deploy_transaction_mainnet(txBytes string, signature1 []byte, signature2 []byte, signature3 []byte) {
	var _txBytes []byte

	deployerAccountId, err := hedera.AccountIDFromString(os.Getenv("HEDERA_XSGD_MAINNET_DEPLOYER_ACCOUNT_ID"))
	if err != nil {
		panic(err)
	}

	deployerPublicKey, err := hedera.PublicKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_DEPLOYER_PUBLIC_KEY"))
	if err != nil {
		panic(err)
	}

	_ = deployerPublicKey

	deployerPrivateKey, err := hedera.PrivateKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_DEPLOYER_PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}

	adminPublicKey1, _ := hedera.PublicKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_ADMIN_PUBLIC_KEY_1"))
	adminPublicKey2, _ := hedera.PublicKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_ADMIN_PUBLIC_KEY_2"))

	treasurerPublicKey, _ := hedera.PublicKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_TREASURER_PUBLIC_KEY"))

	client := hedera.ClientForMainnet()
	client.SetOperator(deployerAccountId, deployerPrivateKey)

	_txBytes, err = hex.DecodeString(txBytes)
	if err != nil {
		fmt.Println("error with txBytes")
	}

	var tx interface{}
	//var createTokenTx hedera.TokenCreateTransaction
	var createTokenTx hedera.TokenCreateTransaction
	var signedCreateTokenTx *hedera.TokenCreateTransaction

	tx, err = hedera.TransactionFromBytes(_txBytes)
	if err != nil {
		fmt.Println("error with txBytes")
		panic(err)
	}

	createTokenTx = tx.(hedera.TokenCreateTransaction)
	signedCreateTokenTx = createTokenTx.AddSignature(adminPublicKey1, signature1).AddSignature(adminPublicKey2, signature2).AddSignature(treasurerPublicKey, signature3)

	response, err := signedCreateTokenTx.Execute(client)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", "response")
	fmt.Println(response)
	receipt, err := response.GetReceipt(client)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", "receipt")
	fmt.Printf("%+v\n", receipt)

	//Get the token ID from the receipt
	tokenId := *receipt.TokenID

	fmt.Printf("The new token ID is %v\n", tokenId)
}

func broadcast_transfer_transaction_mainnet(txBytes string) {
	var _txBytes []byte

	deployerAccountId, err := hedera.AccountIDFromString(os.Getenv("HEDERA_XSGD_MAINNET_DEPLOYER_ACCOUNT_ID"))
	if err != nil {
		panic(err)
	}

	deployerPublicKey, err := hedera.PublicKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_DEPLOYER_PUBLIC_KEY"))
	if err != nil {
		panic(err)
	}

	_ = deployerPublicKey

	deployerPrivateKey, err := hedera.PrivateKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_DEPLOYER_PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}

	client := hedera.ClientForMainnet()
	client.SetOperator(deployerAccountId, deployerPrivateKey)

	_txBytes, err = hex.DecodeString(txBytes)
	if err != nil {
		fmt.Println("error with txBytes")
	}

	var signedTx interface{}
	var signedTransferTx *hedera.TransferTransaction

	signedTx, err = hedera.TransactionFromBytes(_txBytes)
	if err != nil {
		panic(err)
	}
	signedTransferTx = signedTx.(*hedera.TransferTransaction)
	response, err := signedTransferTx.Execute(client)
	if err != nil {
		panic(err)
	}

	//Get the transaction ID
	txId := response.TransactionID
	fmt.Println("The transaction ID ", txId)

	fmt.Println(response)
}

func broadcast_update_transaction_mainnet(txBytes string) {
	var _txBytes []byte

	deployerAccountId, err := hedera.AccountIDFromString(os.Getenv("HEDERA_XSGD_MAINNET_DEPLOYER_ACCOUNT_ID"))
	if err != nil {
		panic(err)
	}

	deployerPublicKey, err := hedera.PublicKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_DEPLOYER_PUBLIC_KEY"))
	if err != nil {
		panic(err)
	}

	_ = deployerPublicKey

	deployerPrivateKey, err := hedera.PrivateKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_DEPLOYER_PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}

	client := hedera.ClientForMainnet()
	client.SetOperator(deployerAccountId, deployerPrivateKey)

	_txBytes, err = hex.DecodeString(txBytes)
	if err != nil {
		fmt.Println("error with txBytes")
	}

	var signedTx interface{}
	var signedTokenUpdateTx *hedera.TokenUpdateTransaction

	signedTx, err = hedera.TransactionFromBytes(_txBytes)
	if err != nil {
		panic(err)
	}
	signedTokenUpdateTx = signedTx.(*hedera.TokenUpdateTransaction)
	response, err := signedTokenUpdateTx.Execute(client)
	if err != nil {
		panic(err)
	}

	//Get the transaction ID
	txId := response.TransactionID
	fmt.Println("The transaction ID ", txId)

	fmt.Println(response)
}

func main() {
	//panic: invalid compressed ecsda public key length: 0 bytes
	broadcast_deploy_transaction_mainnet("txbytes")
}
