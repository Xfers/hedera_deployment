package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/joho/godotenv"
)

func CreateSupplyKeyListDeployMainnetXsgd() hedera.Key {

	keys := make([]hedera.PublicKey, 5)

	keys[0], _ = hedera.PublicKeyFromString("8ad0b83774cbb3f417df7972965b3fad1b2a686bf43f0222db10f090c31b343f") // Ian
	keys[1], _ = hedera.PublicKeyFromString("958b0b08c92219c541e5ebee5a8a250a4b1ed974734c0505a70b500c2df05c42") // Gwenda
	keys[2], _ = hedera.PublicKeyFromString("49a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a2") // Lum jun chi
	keys[3], _ = hedera.PublicKeyFromString("2a192c1728c1e4881c1e8e8847d0729fdabff726ea6b1bba6f7f1a574f0509ce") // David
	keys[4], _ = hedera.PublicKeyFromString("cf68de9adc91c0e06e0fd4be6baafd9b00be588a9900777eae52c0ed2288f6b7") // Victor
	//keys[5], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_SUPPLY_PUBLIC_KEY_2"))               // Aymeric / Mihir

	keyStructure := hedera.KeyListWithThreshold(2).AddAllPublicKeys(keys)

	fmt.Printf("The key list is %v\n", keyStructure)
	return keyStructure
}

func CreateBlacklisterKeyListDeployMainnetXsgd() hedera.Key {

	keys := make([]hedera.PublicKey, 7)

	keys[0], _ = hedera.PublicKeyFromString("3c7e801b043466e90b461b41e3c9de1d4c142787c75e5595cfd3701ca559d6a3") // Tianwei
	keys[1], _ = hedera.PublicKeyFromString("76036ad9338c79d4625143e3dca6ddb6dbc9377257d896c2cb4445ed891db03b") // Samson
	keys[2], _ = hedera.PublicKeyFromString("cf68de9adc91c0e06e0fd4be6baafd9b00be588a9900777eae52c0ed2288f6b7") // Victor
	keys[3], _ = hedera.PublicKeyFromString("8ad0b83774cbb3f417df7972965b3fad1b2a686bf43f0222db10f090c31b343f") // Ian
	keys[4], _ = hedera.PublicKeyFromString("958b0b08c92219c541e5ebee5a8a250a4b1ed974734c0505a70b500c2df05c42") // Gwenda
	keys[5], _ = hedera.PublicKeyFromString("49a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a2") // Lum jun chi
	keys[6], _ = hedera.PublicKeyFromString("2a192c1728c1e4881c1e8e8847d0729fdabff726ea6b1bba6f7f1a574f0509ce") // David
	//keys[7], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_BLACKLISTER_PUBLIC_KEY_1"))          // Aymeric / Mihir

	keyStructure := hedera.KeyListWithThreshold(2).AddAllPublicKeys(keys)

	fmt.Printf("The key list is %v\n", keyStructure)
	return keyStructure
}

func CreateKycKeyListDeployMainnetXsgd() hedera.Key {

	keys := make([]hedera.PublicKey, 7)

	keys[0], _ = hedera.PublicKeyFromString("3c7e801b043466e90b461b41e3c9de1d4c142787c75e5595cfd3701ca559d6a3") // Tianwei
	keys[1], _ = hedera.PublicKeyFromString("76036ad9338c79d4625143e3dca6ddb6dbc9377257d896c2cb4445ed891db03b") // Samson
	keys[2], _ = hedera.PublicKeyFromString("cf68de9adc91c0e06e0fd4be6baafd9b00be588a9900777eae52c0ed2288f6b7") // Victor
	keys[3], _ = hedera.PublicKeyFromString("8ad0b83774cbb3f417df7972965b3fad1b2a686bf43f0222db10f090c31b343f") // Ian
	keys[4], _ = hedera.PublicKeyFromString("958b0b08c92219c541e5ebee5a8a250a4b1ed974734c0505a70b500c2df05c42") // Gwenda
	keys[5], _ = hedera.PublicKeyFromString("49a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a2") // Lum jun chi
	keys[6], _ = hedera.PublicKeyFromString("2a192c1728c1e4881c1e8e8847d0729fdabff726ea6b1bba6f7f1a574f0509ce") // David
	//keys[7], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_KYC_PUBLIC_KEY_1"))                  // Aymeric / Mihir

	keyStructure := hedera.KeyListWithThreshold(2).AddAllPublicKeys(keys)

	fmt.Printf("The key list is %v\n", keyStructure)
	return keyStructure
}

func CreateFeeKeyListDeployMainnetXsgd() hedera.Key {

	keys := make([]hedera.PublicKey, 4)

	keys[0], _ = hedera.PublicKeyFromString("3c7e801b043466e90b461b41e3c9de1d4c142787c75e5595cfd3701ca559d6a3")     // Tianwei
	keys[1], _ = hedera.PublicKeyFromString("cf68de9adc91c0e06e0fd4be6baafd9b00be588a9900777eae52c0ed2288f6b7")     // Victor
	keys[2], _ = hedera.PublicKeyFromString("8ad0b83774cbb3f417df7972965b3fad1b2a686bf43f0222db10f090c31b343f")     // Ian
	keys[3], _ = hedera.PublicKeyFromString("key:db1e32279e9b153c47d37393ca976e82b9f0feeb6077b5fd26b81bbd50262779") // Yuenshe
	//keys[4], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_FEE_PUBLIC_KEY_3"))                      // Aymeric

	keyStructure := hedera.KeyListWithThreshold(3).AddAllPublicKeys(keys)

	fmt.Printf("The key list is %v\n", keyStructure)
	return keyStructure
}

func CreateWiperKeyListDeployMainnetXsgd() hedera.Key {

	keys := make([]hedera.PublicKey, 7)

	keys[0], _ = hedera.PublicKeyFromString("3c7e801b043466e90b461b41e3c9de1d4c142787c75e5595cfd3701ca559d6a3") // Tianwei
	keys[1], _ = hedera.PublicKeyFromString("76036ad9338c79d4625143e3dca6ddb6dbc9377257d896c2cb4445ed891db03b") // Samson
	keys[2], _ = hedera.PublicKeyFromString("cf68de9adc91c0e06e0fd4be6baafd9b00be588a9900777eae52c0ed2288f6b7") // Victor
	keys[3], _ = hedera.PublicKeyFromString("8ad0b83774cbb3f417df7972965b3fad1b2a686bf43f0222db10f090c31b343f") // Ian
	keys[4], _ = hedera.PublicKeyFromString("958b0b08c92219c541e5ebee5a8a250a4b1ed974734c0505a70b500c2df05c42") // Gwenda
	keys[5], _ = hedera.PublicKeyFromString("49a2206ffeb21ca5d5a215d07c4d4af8792c2f436928f0b18619aab3eafce7a2") // Lum jun chi
	keys[6], _ = hedera.PublicKeyFromString("2a192c1728c1e4881c1e8e8847d0729fdabff726ea6b1bba6f7f1a574f0509ce") // David
	keys[7], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_WIPER_PUBLIC_KEY_1"))                // Aymeric / Mihir

	keyStructure := hedera.KeyListWithThreshold(2).AddAllPublicKeys(keys)

	fmt.Printf("The key list is %v\n", keyStructure)
	return keyStructure
}

func CreateAdminPublicKeyListDeployMainnetXsgd() hedera.Key {

	keys := make([]hedera.PublicKey, 5)

	keys[0], _ = hedera.PublicKeyFromString("3c7e801b043466e90b461b41e3c9de1d4c142787c75e5595cfd3701ca559d6a3")     // Tianwei
	keys[1], _ = hedera.PublicKeyFromString("cf68de9adc91c0e06e0fd4be6baafd9b00be588a9900777eae52c0ed2288f6b7")     // Victor
	keys[2], _ = hedera.PublicKeyFromString("d4491e38016c67bdd1b7f2d3bad725048ac6bd05851d9a24317833362ef78921")     // Tianyao
	keys[3], _ = hedera.PublicKeyFromString("key:db1e32279e9b153c47d37393ca976e82b9f0feeb6077b5fd26b81bbd50262779") // Yuenshe
	keys[4], _ = hedera.PublicKeyFromString("76036ad9338c79d4625143e3dca6ddb6dbc9377257d896c2cb4445ed891db03b")     // Samson

	keyStructure := hedera.KeyListWithThreshold(2).AddAllPublicKeys(keys)

	fmt.Printf("The key list is %v\n", keyStructure)
	return keyStructure
}

func create_token_tx_hedera_mainnet_xsgd() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		panic(fmt.Errorf("Unable to load environment variables from .env file. Error:\n%v\n", err))
	}

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

	treasuryAccountId, err := hedera.AccountIDFromString(os.Getenv("HEDERA_XSGD_MAINNET_TREASURER_ACCOUNT_ID"))
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
	var validStartTime time.Time = time.Now() /*.Add(time.Minute * time.Duration(2))*/

	transactionId := hedera.NewTransactionIDWithValidStart(deployerAccountId, validStartTime)

	tokenCreateTransaction, err := hedera.NewTokenCreateTransaction().
		SetTransactionID(transactionId).
		SetTokenName("XSGD").                                      // token name
		SetTokenSymbol("XSGD").                                    // token symbol
		SetTreasuryAccountID(treasuryAccountId).                   // treasury account id
		SetAdminKey(CreateAdminPublicKeyListDeployMainnetXsgd()).  // owner
		SetSupplyKey(CreateSupplyKeyListDeployMainnetXsgd()).      // minter/burner, without them , can't mint or burn
		SetFreezeKey(CreateBlacklisterKeyListDeployMainnetXsgd()). // blacklister
		SetWipeKey(CreateWiperKeyListDeployMainnetXsgd()).         // wiper
		SetKycKey(CreateKycKeyListDeployMainnetXsgd()).            // KYC
		SetFeeScheduleKey(CreateFeeKeyListDeployMainnetXsgd()).    // Fee
		SetInitialSupply(0).                                       // total supply
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
	create_token_tx_hedera_mainnet_xsgd()
}
