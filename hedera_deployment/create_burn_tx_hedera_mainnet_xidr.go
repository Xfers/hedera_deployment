package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/joho/godotenv"
)

func create_burn_tx_hedera_mainnet_xidr() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		panic(fmt.Errorf("Unable to load environment variables from .env file. Error:\n%v\n", err))
	}

	deployerAccountId, err := hedera.AccountIDFromString(os.Getenv("HEDERA_XIDR_MAINNET_DEPLOYER_ACCOUNT_ID"))
	if err != nil {
		panic(err)
	}

	deployerPublicKey, err := hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_DEPLOYER_PUBLIC_KEY"))
	if err != nil {
		panic(err)
	}

	_ = deployerPublicKey

	deployerPrivateKey, err := hedera.PrivateKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_DEPLOYER_PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}

	client := hedera.ClientForMainnet()
	client.SetOperator(deployerAccountId, deployerPrivateKey)

	token_id, err := hedera.TokenIDFromString(os.Getenv("HEDERA_XIDR_MAINNET_TOKEN_ID")) // NEED TO CHANGE THIS TO XIDR MAINNET TOKEN ID
	var mint_amount uint64 = 200000000                                                   // NEED TO CHANGE THIS TO OUR NEED

	var validStartTime time.Time = time.Now() /*.Add(time.Minute * time.Duration(2))*/

	transactionId := hedera.NewTransactionIDWithValidStart(deployerAccountId, validStartTime)

	tokenBurnTransaction, err := hedera.NewTokenBurnTransaction().
		SetTransactionID(transactionId).
		SetTokenID(token_id).
		SetAmount(mint_amount). // 200 mil = 200 token
		FreezeWith(client)

	if err != nil {
		panic(err)
	}

	txBytes, err := tokenBurnTransaction.ToBytes()
	fmt.Printf("%#v\n", hex.EncodeToString(txBytes))
}
