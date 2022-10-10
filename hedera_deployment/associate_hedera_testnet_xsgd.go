package main

import (
	"fmt"
	"time"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

func associate_hedera_testnet_xsgd() {
	//err := godotenv.Load("../.env")
	//if err != nil {
	//	panic(fmt.Errorf("Unable to load environment variables from .env file. Error:\n%v\n", err))
	//}

	//Grab your testnet account ID and private key from the .env file
	senderAccountId, err := hedera.AccountIDFromString("0.0.45949322")
	receiverAccountId, err := hedera.AccountIDFromString("0.0.47984526")
	// fmt.Println(os.Getenv("MY_ACCOUNT_ID"), os.Getenv("MY_PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}

	senderPublicKey, err := hedera.PublicKeyFromString("302a300506032b657003210009fc854d1005984c386206cc68382900a0f9647d7a0ddd164d27f9dcd767f1f6")
	if err != nil {
		panic(err)
	}

	_ = senderPublicKey
	senderPrivateKey, err := hedera.PrivateKeyFromString("302e020100300506032b657004220420c26e9b65de501380d105634579077492fd59b8614b0aedb55410b20a47469f28")
	if err != nil {
		panic(err)
	}

	receiverPublicKey, err := hedera.PublicKeyFromString("5c97a1e4dff419af615f1085ed533e35514e347342ad7fbc79820fb234d6e9e2")
	if err != nil {
		panic(err)
	}
	_ = receiverPublicKey

	receiverPrivateKey, err := hedera.PrivateKeyFromString("302e020100300506032b657004220420831e80289e120ecdf442c614e74e1bf8df85380b0dc7334f6f0a7943fa6d6d4d")
	if err != nil {
		panic(err)
	}

	//Print your testnet account ID and private key to the console to make sure there was no error
	fmt.Printf("The account ID is = %v\n", senderAccountId)
	fmt.Printf("The private key is = %v\n", senderPrivateKey)

	//Create your testnet client
	client := hedera.ClientForTestnet()
	client.SetOperator(senderAccountId, senderPrivateKey)
	//Create the transaction and freeze the unsigned transaction

	var validStartTime time.Time = time.Now() /*.Add(time.Minute * time.Duration(2))*/

	transactionId := hedera.NewTransactionIDWithValidStart(senderAccountId, validStartTime)

	token_id, err := hedera.TokenIDFromString("0.0.47996109")
	tokenTransferTransaction, err := hedera.NewTokenAssociateTransaction().
		SetTransactionID(transactionId).
		SetAccountID(receiverAccountId).
		AddTokenID((token_id)).
		FreezeWith(client)

	if err != nil {
		panic(err)
	}

	//Sign with the admin private key of the token, sign with the token treasury private key, sign with the client operator private key and submit the transaction to a Hedera network
	txResponse, err := tokenTransferTransaction.Sign(receiverPrivateKey).Execute(client)

	if err != nil {
		panic(err)
	}

	//Request the receipt of the transaction
	receipt, err := txResponse.GetReceipt(client)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", receipt)
}

func main() {
	associate_hedera_testnet_xsgd()
}
