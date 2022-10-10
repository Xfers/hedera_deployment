package main

import (
	"encoding/hex"
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

func create_transfer_tx_hedera_testnet_xsgd() {
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

	//Print your testnet account ID and private key to the console to make sure there was no error
	fmt.Printf("The account ID is = %v\n", senderAccountId)
	fmt.Printf("The private key is = %v\n", senderPrivateKey)

	//Create your testnet client
	client := hedera.ClientForTestnet()
	client.SetOperator(senderAccountId, senderPrivateKey)
	//Create the transaction and freeze the unsigned transaction

	token_id, err := hedera.TokenIDFromString("0.0.47996109")

	// RECEIVER ACCOUNT_ID MUST BE ASSOCIATED WITH THE TOKEN_ID BEFORE TRIGGERING THIS TRANSFER TRANSACTION
	tokenTransferTransaction, err := hedera.NewTransferTransaction().
		AddTokenTransfer(token_id, senderAccountId, -10000000).
		AddTokenTransfer(token_id, receiverAccountId, 10000000).
		FreezeWith(client)

	if err != nil {
		panic(err)
	}

	//Sign with the admin private key of the token, sign with the token treasury private key, sign with the client operator private key and submit the transaction to a Hedera network
	//txResponse, err := tokenTransferTransaction.Sign(senderPrivateKey).Execute(client)

	if err != nil {
		panic(err)
	}

	txBytes, err := tokenTransferTransaction.ToBytes()
	fmt.Printf("%#v\n", hex.EncodeToString(txBytes))

}
func main() {
	create_transfer_tx_hedera_testnet_xsgd()
}
