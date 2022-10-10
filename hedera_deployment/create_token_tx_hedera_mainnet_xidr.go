package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/joho/godotenv"
)

func CreateSupplyKeyListDeployMainnetXidr() hedera.Key {

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

func CreateBlacklisterKeyListDeployMainnetXidr() hedera.Key {

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

func CreateKycKeyListDeployMainnetXidr() hedera.Key {

	keys := make([]hedera.PublicKey, 8)

	keys[0], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_KYC_PUBLIC_KEY_1"))
	keys[1], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_KYC_PUBLIC_KEY_2"))
	keys[2], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_KYC_PUBLIC_KEY_3"))
	keys[3], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_KYC_PUBLIC_KEY_4"))
	keys[4], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_KYC_PUBLIC_KEY_5"))
	keys[5], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_KYC_PUBLIC_KEY_6"))
	keys[6], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_KYC_PUBLIC_KEY_7"))
	keys[7], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_KYC_PUBLIC_KEY_8"))

	keyStructure := hedera.KeyListWithThreshold(2).AddAllPublicKeys(keys)

	fmt.Printf("The key list is %v\n", keyStructure)
	return keyStructure
}

func CreateFeeKeyListDeployMainnetXidr() hedera.Key {

	keys := make([]hedera.PublicKey, 5)

	keys[0], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_FEE_PUBLIC_KEY_1"))
	keys[1], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_FEE_PUBLIC_KEY_2"))
	keys[2], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_FEE_PUBLIC_KEY_3"))
	keys[3], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_FEE_PUBLIC_KEY_4"))
	keys[4], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_FEE_PUBLIC_KEY_5"))

	keyStructure := hedera.KeyListWithThreshold(3).AddAllPublicKeys(keys)

	fmt.Printf("The key list is %v\n", keyStructure)
	return keyStructure
}

func CreateWiperKeyListDeployMainnetXidr() hedera.Key {

	keys := make([]hedera.PublicKey, 8)

	keys[0], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_WIPER_PUBLIC_KEY_1"))
	keys[1], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_WIPER_PUBLIC_KEY_2"))
	keys[2], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_WIPER_PUBLIC_KEY_3"))
	keys[3], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_WIPER_PUBLIC_KEY_4"))
	keys[4], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_WIPER_PUBLIC_KEY_5"))
	keys[5], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_WIPER_PUBLIC_KEY_6"))
	keys[6], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_WIPER_PUBLIC_KEY_7"))
	keys[7], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_WIPER_PUBLIC_KEY_8"))

	keyStructure := hedera.KeyListWithThreshold(2).AddAllPublicKeys(keys)

	fmt.Printf("The key list is %v\n", keyStructure)
	return keyStructure
}

func CreateAdminPublicKeyListDeployMainnetXidr() hedera.Key {

	keys := make([]hedera.PublicKey, 5)

	keys[0], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_ADMIN_PUBLIC_KEY_1"))
	keys[1], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_ADMIN_PUBLIC_KEY_2"))
	keys[2], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_ADMIN_PUBLIC_KEY_3"))
	keys[3], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_ADMIN_PUBLIC_KEY_4"))
	keys[4], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_ADMIN_PUBLIC_KEY_5"))

	keyStructure := hedera.KeyListWithThreshold(2).AddAllPublicKeys(keys)

	fmt.Printf("The key list is %v\n", keyStructure)
	return keyStructure
}

func create_token_tx_hedera_mainnet_xidr() {
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

	treasuryAccountId, err := hedera.AccountIDFromString(os.Getenv("HEDERA_XIDR_MAINNET_TREASURER_ACCOUNT_ID"))
	if err != nil {
		panic(err)
	}

	//Create your mainnet client
	client := hedera.ClientForMainnet()
	client.SetOperator(deployerAccountId, deployerPrivateKey)
	//Create the transaction and freeze the unsigned transaction

	//blacklisterPublicKey, _ := hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_BLACKLISTER_PUBLIC_KEY"))
	//blacklisterPrivateKey, _ := hedera.PrivateKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_BLACKLISTER_PRIVATE_KEY"))

	//supplyPublicKey, _ := hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_SUPPLY_PUBLIC_KEY"))
	//supplyPrivateKey, _ := hedera.PrivateKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_SUPPLY_PRIVATE_KEY"))

	//fmt.Printf("Blacklister pbk: %v\n Blacklister pvk: %v\n========\n", blacklisterPublicKey, blacklisterPrivateKey)
	//fmt.Printf("Supply pbk: %v\n Supply pvk: %v\n=========\n", supplyPublicKey, supplyPrivateKey)

	// CONFIGURE VALID START TIME
	var validStartTime time.Time = time.Now() /*.Add(time.Minute * time.Duration(2))*/

	transactionId := hedera.NewTransactionIDWithValidStart(deployerAccountId, validStartTime)

	tokenCreateTransaction, err := hedera.NewTokenCreateTransaction().
		SetTransactionID(transactionId).
		SetTokenName("xidr").                                      // token name
		SetTokenSymbol("xidr").                                    // token symbol
		SetTreasuryAccountID(treasuryAccountId).                   // treasury account id
		SetAdminKey(CreateAdminPublicKeyListDeployMainnetXidr()).  // owner
		SetSupplyKey(CreateSupplyKeyListDeployMainnetXidr()).      // minter/burner, without them , can't mint or burn
		SetFreezeKey(CreateBlacklisterKeyListDeployMainnetXidr()). // blacklister
		SetWipeKey(CreateWiperKeyListDeployMainnetXidr()).         // wiper
		SetKycKey(CreateKycKeyListDeployMainnetXidr()).            // KYC
		SetFeeScheduleKey(CreateFeeKeyListDeployMainnetXidr()).    // Fee
		SetInitialSupply(0).                                       // total supply
		//SetDecimals(6).                                            // decimal
		//SetMaxTransactionFee(hedera.NewHbar(30)).                  // Change the default max transaction fee
		FreezeWith(client)
	// scheduled transaction doesn't neeed to be freeze

	//adminPrivateKey1, _ := hedera.PrivateKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_OWNER_PRIVATE_KEY_1"))
	//adminPrivateKey2, _ := hedera.PrivateKeyFromString(os.Getenv("HEDERA_XIDR_MAINNET_OWNER_PRIVATE_KEY_2"))

	// NON-SCHEDULED TRANSACTION APPROACH
	// var temp *hedera.TokenCreateTransaction
	// send txbytes to airgap machine to sign and collect signature
	//temp = tokenCreateTransaction.Sign(adminPrivateKey1).Sign(adminPrivateKey2).Sign(treasuryKey)
	txBytes, err := tokenCreateTransaction.ToBytes()
	fmt.Printf("%#v\n", hex.EncodeToString(txBytes))
}

//{"transaction": "tx_hex_string"}
func main() {
	create_token_tx_hedera_mainnet_xidr()
}
