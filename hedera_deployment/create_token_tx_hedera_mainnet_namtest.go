package main

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

func CreateSupplyKeyListDeployMainnetNamTest() hedera.Key {

	keys := make([]hedera.PublicKey, 2)

	keys[0], _ = hedera.PublicKeyFromString("49a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a2") // Lum jun chi
	keys[1], _ = hedera.PublicKeyFromString("8ad0b83774cbb3f417df7972965b3fad1b2a686bf43f0222db10f090c31b343f") // Ian

	keyStructure := hedera.KeyListWithThreshold(2).AddAllPublicKeys(keys)

	fmt.Printf("The key list is %v\n", keyStructure)
	return keyStructure
}

func CreateAdminPublicKeyListDeployMainnetNamTest() hedera.Key {

	keys := make([]hedera.PublicKey, 2)

	keys[0], _ = hedera.PublicKeyFromString("49a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a2") // Lum jun chi
	keys[1], _ = hedera.PublicKeyFromString("2a192c1728c1e4881c1e8e8847d0729fdabff726ea6b1bba6f7f1a574f0509ce") // David

	keyStructure := hedera.KeyListWithThreshold(2).AddAllPublicKeys(keys)

	fmt.Printf("The key list is %v\n", keyStructure)
	return keyStructure
}

func create_token_tx_hedera_mainnet_namtest() {

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

	treasuryAccountId, err := hedera.AccountIDFromString("0.0.950692") // Ian
	if err != nil {
		panic(err)
	}

	//Create your mainnet client
	client := hedera.ClientForMainnet()
	client.SetOperator(deployerAccountId, deployerPrivateKey)
	//Create the transaction and freeze the unsigned transaction

	//blacklisterPublicKey, _ := hedera.PublicKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_BLACKLISTER_PUBLIC_KEY"))
	//blacklisterPrivateKey, _ := hedera.PrivateKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_BLACKLISTER_PRIVATE_KEY"))

	//supplyPublicKey, _ := hedera.PublicKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_SUPPLY_PUBLIC_KEY"))
	//supplyPrivateKey, _ := hedera.PrivateKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_SUPPLY_PRIVATE_KEY"))

	//fmt.Printf("Blacklister pbk: %v\n Blacklister pvk: %v\n========\n", blacklisterPublicKey, blacklisterPrivateKey)
	//fmt.Printf("Supply pbk: %v\n Supply pvk: %v\n=========\n", supplyPublicKey, supplyPrivateKey)

	// CONFIGURE VALID START TIME
	var validStartTime time.Time = time.Now().Add(time.Minute * time.Duration(65))

	transactionId := hedera.NewTransactionIDWithValidStart(deployerAccountId, validStartTime)

	tokenCreateTransaction, err := hedera.NewTokenCreateTransaction().
		SetTransactionID(transactionId).
		SetTokenName("NAMTEST").                                     // token name
		SetTokenSymbol("NAMTEST").                                   // token symbol
		SetTreasuryAccountID(treasuryAccountId).                     // treasury account id
		SetAdminKey(CreateAdminPublicKeyListDeployMainnetNamTest()). // owner
		SetSupplyKey(CreateSupplyKeyListDeployMainnetNamTest()).     // minter/burner, without them , can't mint or burn
		//SetFreezeKey(CreateBlacklisterKeyListDeployMainnetXsgd()). // blacklister
		//SetWipeKey(CreateWiperKeyListDeployMainnetXsgd()).         // wiper
		//SetKycKey(CreateKycKeyListDeployMainnetXsgd()).            // KYC
		//SetFeeScheduleKey(CreateFeeKeyListDeployMainnetXsgd()).    // Fee
		SetInitialSupply(0). // total supply
		//SetDecimals(6).                                            // decimal
		//SetMaxTransactionFee(hedera.NewHbar(30)).                  // Change the default max transaction fee
		FreezeWith(client)
	// scheduled transaction doesn't neeed to be freeze

	//adminPrivateKey1, _ := hedera.PrivateKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_OWNER_PRIVATE_KEY_1"))
	//adminPrivateKey2, _ := hedera.PrivateKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_OWNER_PRIVATE_KEY_2"))

	// NON-SCHEDULED TRANSACTION APPROACH
	// var temp *hedera.TokenCreateTransaction
	// send txbytes to airgap machine to sign and collect signature
	//temp = tokenCreateTransaction.Sign(adminPrivateKey1).Sign(adminPrivateKey2).Sign(treasuryKey)
	txBytes, err := tokenCreateTransaction.ToBytes()
	fmt.Printf("%#v\n", hex.EncodeToString(txBytes))
}

//{"transaction": "tx_hex_string"}
func main() {
	create_token_tx_hedera_mainnet_namtest()
}
