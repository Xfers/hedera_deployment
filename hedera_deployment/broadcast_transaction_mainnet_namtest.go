package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

//func broadcast_mint_transaction_mainnet(txBytes string, signature1 []byte, signature2 []byte, signature3 []byte) {
func broadcast_mint_transaction_mainnet_namtest(txBytes string, s1 string, s2 string) {
	var _txBytes []byte

	deployerAccountId, err := hedera.AccountIDFromString("0.0.1126201")
	if err != nil {
		panic(err)
	}

	deployerPublicKey, err := hedera.PublicKeyFromString("3a16c869ec6e4c4d96a78a04cfdddda79ccd5c464c21619f21076d462f643d7a")
	if err != nil {
		panic(err)
	}

	_ = deployerPublicKey

	deployerPrivateKey, err := hedera.PrivateKeyFromString("302e020100300506032b65700422042086ee233f7278fb32e9a51e1a1e9daa4555db51f29abcceb5f4d4acbb334a54ef")
	if err != nil {
		panic(err)
	}

	supplyPublicKey1, err := hedera.PublicKeyFromString("49a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a2") // Lum Jun Chi
	if err != nil {
		panic(err)
	}

	supplyPublicKey2, err := hedera.PublicKeyFromString("8ad0b83774cbb3f417df7972965b3fad1b2a686bf43f0222db10f090c31b343f") // Ian
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

	// if necessary , convert signature string to byte[]
	signature1, err := hex.DecodeString(s1)
	signature2, err := hex.DecodeString(s2)

	mintTokenTx = tx.(hedera.TokenMintTransaction)
	signedMintTokenTx = mintTokenTx.AddSignature(supplyPublicKey1, signature1).AddSignature(supplyPublicKey2, signature2)

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

	deployerAccountId, err := hedera.AccountIDFromString("0.0.1126201")
	if err != nil {
		panic(err)
	}

	deployerPublicKey, err := hedera.PublicKeyFromString("3a16c869ec6e4c4d96a78a04cfdddda79ccd5c464c21619f21076d462f643d7a")
	if err != nil {
		panic(err)
	}

	_ = deployerPublicKey

	deployerPrivateKey, err := hedera.PrivateKeyFromString("302e020100300506032b65700422042086ee233f7278fb32e9a51e1a1e9daa4555db51f29abcceb5f4d4acbb334a54ef")
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

//func broadcast_deploy_transaction_mainnet(txBytes string, signature1 []byte, signature2 []byte, signature3 []byte) {
func broadcast_deploy_transaction_mainnet_namtest(txBytes string, s1 string, s2 string, s3 string) {
	var _txBytes []byte

	deployerAccountId, err := hedera.AccountIDFromString("0.0.1126201")
	if err != nil {
		panic(err)
	}

	deployerPublicKey, err := hedera.PublicKeyFromString("3a16c869ec6e4c4d96a78a04cfdddda79ccd5c464c21619f21076d462f643d7a")
	if err != nil {
		panic(err)
	}

	_ = deployerPublicKey

	deployerPrivateKey, err := hedera.PrivateKeyFromString("302e020100300506032b65700422042086ee233f7278fb32e9a51e1a1e9daa4555db51f29abcceb5f4d4acbb334a54ef")
	if err != nil {
		panic(err)
	}

	adminPublicKey1, _ := hedera.PublicKeyFromString("49a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a2") // Lum Jun Chi
	adminPublicKey2, _ := hedera.PublicKeyFromString("2a192c1728c1e4881c1e8e8847d0729fdabff726ea6b1bba6f7f1a574f0509ce") // David

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

	// if necessary , convert signature string to byte[]
	signature1, err := hex.DecodeString(s1)
	signature2, err := hex.DecodeString(s2)
	signature3, err := hex.DecodeString(s3)

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
	tx_bytes := "0af5012af2010aed010a140a0c0881a7a49a0610c0f68b8101120418b9de44120218181880d0acf30e22020878ea01c5010a074e414d5445535412074e414d544553542a0418a4833a324e2a4c080212480a22122049a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a20a2212202a192c1728c1e4881c1e8e8847d0729fdabff726ea6b1bba6f7f1a574f0509ce524e2a4c080212480a22122049a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a20a2212208ad0b83774cbb3f417df7972965b3fad1b2a686bf43f0222db10f090c31b343f720418b9de447a0508d0c8e10312000af5012af2010aed010a140a0c0881a7a49a0610c0f68b8101120418b9de44120218181880d0acf30e22020878ea01c5010a074e414d5445535412074e414d544553542a0418a4833a324e2a4c080212480a22122049a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a20a2212202a192c1728c1e4881c1e8e8847d0729fdabff726ea6b1bba6f7f1a574f0509ce524e2a4c080212480a22122049a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a20a2212208ad0b83774cbb3f417df7972965b3fad1b2a686bf43f0222db10f090c31b343f720418b9de447a0508d0c8e10312000af5012af2010aed010a140a0c0881a7a49a0610c0f68b8101120418b9de44120218181880d0acf30e22020878ea01c5010a074e414d5445535412074e414d544553542a0418a4833a324e2a4c080212480a22122049a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a20a2212202a192c1728c1e4881c1e8e8847d0729fdabff726ea6b1bba6f7f1a574f0509ce524e2a4c080212480a22122049a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a20a2212208ad0b83774cbb3f417df7972965b3fad1b2a686bf43f0222db10f090c31b343f720418b9de447a0508d0c8e10312000af5012af2010aed010a140a0c0881a7a49a0610c0f68b8101120418b9de44120218181880d0acf30e22020878ea01c5010a074e414d5445535412074e414d544553542a0418a4833a324e2a4c080212480a22122049a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a20a2212202a192c1728c1e4881c1e8e8847d0729fdabff726ea6b1bba6f7f1a574f0509ce524e2a4c080212480a22122049a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a20a2212208ad0b83774cbb3f417df7972965b3fad1b2a686bf43f0222db10f090c31b343f720418b9de447a0508d0c8e10312000af5012af2010aed010a140a0c0881a7a49a0610c0f68b8101120418b9de44120218181880d0acf30e22020878ea01c5010a074e414d5445535412074e414d544553542a0418a4833a324e2a4c080212480a22122049a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a20a2212202a192c1728c1e4881c1e8e8847d0729fdabff726ea6b1bba6f7f1a574f0509ce524e2a4c080212480a22122049a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a20a2212208ad0b83774cbb3f417df7972965b3fad1b2a686bf43f0222db10f090c31b343f720418b9de447a0508d0c8e10312000af5012af2010aed010a140a0c0881a7a49a0610c0f68b8101120418b9de44120218181880d0acf30e22020878ea01c5010a074e414d5445535412074e414d544553542a0418a4833a324e2a4c080212480a22122049a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a20a2212202a192c1728c1e4881c1e8e8847d0729fdabff726ea6b1bba6f7f1a574f0509ce524e2a4c080212480a22122049a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a20a2212208ad0b83774cbb3f417df7972965b3fad1b2a686bf43f0222db10f090c31b343f720418b9de447a0508d0c8e10312000af5012af2010aed010a140a0c0881a7a49a0610c0f68b8101120418b9de44120218181880d0acf30e22020878ea01c5010a074e414d5445535412074e414d544553542a0418a4833a324e2a4c080212480a22122049a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a20a2212202a192c1728c1e4881c1e8e8847d0729fdabff726ea6b1bba6f7f1a574f0509ce524e2a4c080212480a22122049a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a20a2212208ad0b83774cbb3f417df7972965b3fad1b2a686bf43f0222db10f090c31b343f720418b9de447a0508d0c8e10312000af5012af2010aed010a140a0c0881a7a49a0610c0f68b8101120418b9de44120218181880d0acf30e22020878ea01c5010a074e414d5445535412074e414d544553542a0418a4833a324e2a4c080212480a22122049a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a20a2212202a192c1728c1e4881c1e8e8847d0729fdabff726ea6b1bba6f7f1a574f0509ce524e2a4c080212480a22122049a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a20a2212208ad0b83774cbb3f417df7972965b3fad1b2a686bf43f0222db10f090c31b343f720418b9de447a0508d0c8e10312000af5012af2010aed010a140a0c0881a7a49a0610c0f68b8101120418b9de44120218181880d0acf30e22020878ea01c5010a074e414d5445535412074e414d544553542a0418a4833a324e2a4c080212480a22122049a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a20a2212202a192c1728c1e4881c1e8e8847d0729fdabff726ea6b1bba6f7f1a574f0509ce524e2a4c080212480a22122049a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a20a2212208ad0b83774cbb3f417df7972965b3fad1b2a686bf43f0222db10f090c31b343f720418b9de447a0508d0c8e1031200"

	broadcast_deploy_transaction_mainnet_namtest(tx_bytes,
		"a",
		"a",
		"a")
	//broadcast_mint_transaction_mainnet_namtest(tx_bytes,
	//	"a",
	//	"a")
	//[]byte{})
}
