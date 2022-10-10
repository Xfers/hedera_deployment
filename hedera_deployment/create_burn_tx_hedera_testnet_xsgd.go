package main

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

func create_burn_tx_hedera_testnet_xsgd() {
	//err := godotenv.Load("../.env")
	//if err != nil {
	//	panic(fmt.Errorf("Unable to load environment variables from .env file. Error:\n%v\n", err))
	//}

	//Grab your testnet account ID and private key from the .env file
	myAccountId, err := hedera.AccountIDFromString("0.0.45949322")
	// fmt.Println(os.Getenv("MY_ACCOUNT_ID"), os.Getenv("MY_PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}

	myPublicKey, err := hedera.PublicKeyFromString("302a300506032b657003210009fc854d1005984c386206cc68382900a0f9647d7a0ddd164d27f9dcd767f1f6")
	if err != nil {
		panic(err)
	}

	_ = myPublicKey
	myPrivateKey, err := hedera.PrivateKeyFromString("302e020100300506032b657004220420c26e9b65de501380d105634579077492fd59b8614b0aedb55410b20a47469f28")
	if err != nil {
		panic(err)
	}

	//Print your testnet account ID and private key to the console to make sure there was no error
	fmt.Printf("The account ID is = %v\n", myAccountId)
	fmt.Printf("The private key is = %v\n", myPrivateKey)

	//Create your testnet client
	client := hedera.ClientForTestnet()
	client.SetOperator(myAccountId, myPrivateKey)
	//Create the transaction and freeze the unsigned transaction

	//	supplierPrivateKey1, err := hedera.PrivateKeyFromString("302e020100300506032b657004220420386c65217f58c9d6df6d92ff52f386dfe199b8ed661a00e7431db21646c769df")
	//	supplierPrivateKey2, err := hedera.PrivateKeyFromString("302e020100300506032b6570042204209d3e65325b92e4dc130cb59bb12cd8756b2b79b4aa5657ed82655fc4f4e9262d")
	//	supplierPrivateKey3, err := hedera.PrivateKeyFromString("302e020100300506032b657004220420851425f000a82a60de4a5c3a089837e3bad2948ec20c6adb8f297036741450cb")

	//blacklisterPublicKey, blacklisterPrivateKey := getRandomKeypair()
	//supplyPublicKey, supplyPrivateKey := getRandomKeypair()

	treasuryAccountId := myAccountId
	treasuryKey := myPrivateKey
	fmt.Printf("Treasury id: %v\n Treasury pvk: %v\n========\n", treasuryAccountId, treasuryKey)

	var validStartTime time.Time = time.Now() /*.Add(time.Minute * time.Duration(2))*/

	transactionId := hedera.NewTransactionIDWithValidStart(myAccountId, validStartTime)

	token_id, err := hedera.TokenIDFromString("0.0.47996109")
	tokenBurnTransaction, err := hedera.NewTokenBurnTransaction().
		SetTransactionID(transactionId).
		SetTokenID(token_id).
		SetAmount(50000000).
		FreezeWith(client)

	if err != nil {
		panic(err)
	}

	//Sign with the admin private key of the token, sign with the token treasury private key, sign with the client operator private key and submit the transaction to a Hedera network
	//txResponse, err := tokenBurnTransaction.Sign(supplierPrivateKey1).Sign(supplierPrivateKey2).Sign(supplierPrivateKey3).Execute(client)

	txBytes, err := tokenBurnTransaction.ToBytes()
	fmt.Printf("%#v\n", hex.EncodeToString(txBytes))

	if err != nil {
		panic(err)
	}

}
