package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

func CreateSupplyKeyListUpdateMainnetXidr() hedera.Key {

	keys := make([]hedera.PublicKey, 6)

	keys[0], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_SUPPLY_PUBLIC_KEY_1"))
	keys[1], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_SUPPLY_PUBLIC_KEY_2"))
	keys[2], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_SUPPLY_PUBLIC_KEY_3"))
	keys[3], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_SUPPLY_PUBLIC_KEY_4"))
	keys[4], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_SUPPLY_PUBLIC_KEY_5"))
	keys[5], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_SUPPLY_PUBLIC_KEY_6"))

	keyStructure := hedera.KeyListWithThreshold(2).AddAllPublicKeys(keys)

	fmt.Printf("The key list is %v\n", keyStructure)
	return keyStructure
}

func CreateBlacklisterKeyListUpdateMainnetXidr() hedera.Key {

	keys := make([]hedera.PublicKey, 8)

	keys[0], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_BLACKLISTER_PUBLIC_KEY_1"))
	keys[1], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_BLACKLISTER_PUBLIC_KEY_2"))
	keys[2], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_BLACKLISTER_PUBLIC_KEY_3"))
	keys[3], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_BLACKLISTER_PUBLIC_KEY_4"))
	keys[4], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_BLACKLISTER_PUBLIC_KEY_5"))
	keys[5], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_BLACKLISTER_PUBLIC_KEY_6"))
	keys[6], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_BLACKLISTER_PUBLIC_KEY_7"))
	keys[7], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_BLACKLISTER_PUBLIC_KEY_8"))

	keyStructure := hedera.KeyListWithThreshold(2).AddAllPublicKeys(keys)

	fmt.Printf("The key list is %v\n", keyStructure)
	return keyStructure
}

func create_update_tx_hedera_keys_mainnet_xidr() {

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

	//Grab your testnet account ID and private key from the .env file
	ownerAccountId, err := hedera.AccountIDFromString(os.Getenv("HEDERA_XIDR_MAINNET_OWNER_ACCOUNT_ID"))
	// fmt.Println(os.Getenv("MY_ACCOUNT_ID"), os.Getenv("MY_PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}

	ownerPublicKey, err := hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_OWNER_PUBLIC_KEY"))
	if err != nil {
		panic(err)
	}

	_ = ownerPublicKey

	ownerPrivateKey, err := hedera.PrivateKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_OWNER_PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}

	//Print your testnet account ID and private key to the console to make sure there was no error
	fmt.Printf("The account ID is = %v\n", ownerAccountId)
	fmt.Printf("The private key is = %v\n", ownerPrivateKey)

	//Create your mainnet client
	client := hedera.ClientForMainnet()
	client.SetOperator(deployerAccountId, deployerPrivateKey)
	//Create the transaction and freeze the unsigned transaction

	tokenID, err := hedera.TokenIDFromString(os.Getenv("HEDERA_XIDR_MAINNET_TOKEN_ID"))
	var validStartTime time.Time = time.Now() /*.Add(time.Minute * time.Duration(2))*/

	transactionId := hedera.NewTransactionIDWithValidStart(deployerAccountId, validStartTime)

	//Create the transaction and freeze for manual signing
	tokenUpdateTransaction, err := hedera.NewTokenUpdateTransaction().
		SetTransactionID(transactionId).
		SetTokenID(tokenID).
		SetSupplyKey(CreateSupplyKeyListUpdateMainnetXidr()).      // mint and burn
		SetFreezeKey(CreateBlacklisterKeyListUpdateMainnetXidr()). // blacklister
		SetTokenName("XIDR").                                      // token name
		SetTokenSymbol("XIDR").                                    // token symbol
		//SetTreasuryAccountID((treasuryAccountId).
		SetMaxTransactionFee(hedera.NewHbar(30)). // Change the default max transaction fee
		FreezeWith(client)

	if err != nil {
		panic(err)
	}

	// [Important]
	// Sign with the admin private key of the token, sign with the client operator private key and submit the transaction to a Hedera network
	txResponse, err := tokenUpdateTransaction.Sign(ownerPrivateKey).Execute(client)

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
}

func main() {
	create_update_tx_hedera_keys_mainnet_xidr()
}
