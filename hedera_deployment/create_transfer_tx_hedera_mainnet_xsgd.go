package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/joho/godotenv"
)

func create_transfer_tx_hedera_mainnet_xsgd() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		panic(fmt.Errorf("Unable to load environment variables from .env file. Error:\n%v\n", err))
	}

	senderAccountId, err := hedera.AccountIDFromString(os.Getenv("HEDERA_XSGD_MAINNET_SENDER_ACCOUNT_ID"))
	receiverAccountId, err := hedera.AccountIDFromString(os.Getenv("HEDERA_XSGD_MAINNET_RECEIVER_ACCOUNT_ID"))

	if err != nil {
		panic(err)
	}

	senderPublicKey, err := hedera.PublicKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_SENDER_PUBLIC_KEY_1"))
	if err != nil {
		panic(err)
	}

	_ = senderPublicKey
	senderPrivateKey, err := hedera.PrivateKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_SENDER_PRIVATE_KEY_1"))
	if err != nil {
		panic(err)
	}

	//Create your mainnet client
	client := hedera.ClientForMainnet()
	client.SetOperator(senderAccountId, senderPrivateKey)
	//Create the transaction and freeze the unsigned transaction

	token_id, err := hedera.TokenIDFromString(os.Getenv("HEDERA_XSGD_MAINNET_TOKEN_ID"))

	var amount int64 = 100000

	// RECEIVER ACCOUNT_ID MUST BE ASSOCIATED WITH THE TOKEN_ID BEFORE TRIGGERING THIS TRANSFER TRANSACTION
	tokenTransferTransaction, err := hedera.NewTransferTransaction().
		AddTokenTransfer(token_id, senderAccountId, -amount).
		AddTokenTransfer(token_id, receiverAccountId, amount).
		FreezeWith(client)

	if err != nil {
		panic(err)
	}

	//txResponse, err := tokenTransferTransaction.Sign(senderPrivateKey).Execute(client)

	if err != nil {
		panic(err)
	}

	txBytes, err := tokenTransferTransaction.ToBytes()
	fmt.Printf("%#v\n", hex.EncodeToString(txBytes))

}
func main() {
	create_transfer_tx_hedera_mainnet_xsgd()
}
