package main

import (
	"fmt"
	"time"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

func create_delete_tx_hedera_mainnet_xidr() {

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

	//Print your testnet account ID and private key to the console to make sure there was no error
	fmt.Printf("The account ID is = %v\n", deployerAccountId)
	fmt.Printf("The private key is = %v\n", deployerPrivateKey)

	//Create your mainnet client
	client := hedera.ClientForMainnet()
	client.SetOperator(deployerAccountId, deployerPrivateKey)

	tokenID, err := hedera.TokenIDFromString("0.0.1126254") // XIDR

	var validStartTime time.Time = time.Now() /*.Add(time.Minute * time.Duration(2))*/

	transactionId := hedera.NewTransactionIDWithValidStart(deployerAccountId, validStartTime)

	//Create the transaction and freeze the unsigned transaction for manual signing
	transaction, err := hedera.NewTokenDeleteTransaction().
		SetTransactionID(transactionId).
		SetTokenID(tokenID).
		FreezeWith(client)

	if err != nil {
		panic(err)
	}

	//Sign with the admin private key of the account, submit the transaction to a Hedera network
	txResponse, err := transaction.Sign(deployerPrivateKey).Execute(client)

	if err != nil {
		panic(err)
	}

	//Request the receipt of the transaction
	receipt, err := txResponse.GetReceipt(client)

	if err != nil {
		panic(err)
	}

	//Get the transaction consensus status
	status := receipt.Status

	fmt.Printf("The transaction consensus status is %v\n", status)

	//v2.1.0
}
