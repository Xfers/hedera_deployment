package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

func CreateAdminKeyListDeployTestnetXidr() hedera.Key {

	keys := make([]hedera.PublicKey, 2)

	keys[0], _ = hedera.PublicKeyFromString("302a300506032b657003210009fc854d1005984c386206cc68382900a0f9647d7a0ddd164d27f9dcd767f1f6")
	keys[1], _ = hedera.PublicKeyFromString("5c97a1e4dff419af615f1085ed533e35514e347342ad7fbc79820fb234d6e9e2")

	keyStructure := hedera.KeyListWithThreshold(2).AddAllPublicKeys(keys)

	fmt.Printf("The key list is %v\n", keyStructure)
	return keyStructure
}

func CreateSupplyKeyListDeployTestnetXidr() hedera.Key {

	keys := make([]hedera.PublicKey, 3)

	keys[0], _ = hedera.PublicKeyFromString("302a300506032b6570032100a2ca56b83a3e2011e1b81316553dfbd4e13777e2924d55edfff0765191494520")
	keys[1], _ = hedera.PublicKeyFromString("302a300506032b65700321005832813ab4ff848a60dfa7ef55f8930855ceda7bf34f6b2a2bd5b97b092e14e8")
	keys[2], _ = hedera.PublicKeyFromString("302a300506032b6570032100cd3b39d33a3eea3dfdbfa86a7f938fb42be6f702a47799c13c3a3bddfe516b52")

	keyStructure := hedera.KeyListWithThreshold(2).AddAllPublicKeys(keys)

	fmt.Printf("The key list is %v\n", keyStructure)
	return keyStructure
}

func CreateBlacklisterKeyListDeployTestnetXidr() hedera.Key {

	keys := make([]hedera.PublicKey, 3)

	keys[0], _ = hedera.PublicKeyFromString("302a300506032b6570032100068f0eab7ae589debf5d0ddf7db1b0acf844c254ce472a4bddcc7b04f07d8dc4")
	keys[1], _ = hedera.PublicKeyFromString("302a300506032b6570032100f556b42ed6d5a7a95bf5c44369296c14984831eaa2010d6a3b0450fac5ec5d03")
	keys[2], _ = hedera.PublicKeyFromString("302a300506032b657003210099945810ccac97cf9fc98877dad6028576cd901e967830a98ea5e2b10bdbeb01")

	keyStructure := hedera.KeyListWithThreshold(2).AddAllPublicKeys(keys)

	fmt.Printf("The key list is %v\n", keyStructure)
	return keyStructure
}

func CreateKycKeyListDeployTestnetxidr() hedera.Key {

	keys := make([]hedera.PublicKey, 2)

	keys[0], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_TESTNET_KYC_PUBLIC_KEY_1"))
	keys[1], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_TESTNET_KYC_PUBLIC_KEY_2"))

	keyStructure := hedera.KeyListWithThreshold(2).AddAllPublicKeys(keys)

	fmt.Printf("The key list is %v\n", keyStructure)
	return keyStructure
}

func CreateFeeKeyListDeployTestnetxidr() hedera.Key {

	keys := make([]hedera.PublicKey, 3)

	keys[0], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_TESTNET_FEE_PUBLIC_KEY_1"))
	keys[1], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_TESTNET_FEE_PUBLIC_KEY_2"))
	keys[2], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_TESTNET_FEE_PUBLIC_KEY_3"))

	keyStructure := hedera.KeyListWithThreshold(3).AddAllPublicKeys(keys)

	fmt.Printf("The key list is %v\n", keyStructure)
	return keyStructure
}

func CreateWiperKeyListDeployTestnetxidr() hedera.Key {

	keys := make([]hedera.PublicKey, 2)

	keys[0], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_TESTNET_WIPER_PUBLIC_KEY_1"))
	keys[1], _ = hedera.PublicKeyFromString(os.Getenv("HEDERA_XIDR_TESTNET_WIPER_PUBLIC_KEY_2"))

	keyStructure := hedera.KeyListWithThreshold(2).AddAllPublicKeys(keys)

	fmt.Printf("The key list is %v\n", keyStructure)
	return keyStructure
}

func create_token_tx_hedera_testnet_xidr() {
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

	adminPrivateKey1, err := hedera.PrivateKeyFromString("302e020100300506032b657004220420c26e9b65de501380d105634579077492fd59b8614b0aedb55410b20a47469f28")
	adminPrivateKey2, err := hedera.PrivateKeyFromString("302e020100300506032b657004220420831e80289e120ecdf442c614e74e1bf8df85380b0dc7334f6f0a7943fa6d6d4d")

	//blacklisterPublicKey, blacklisterPrivateKey := getRandomKeypair()
	//supplyPublicKey, supplyPrivateKey := getRandomKeypair()

	treasuryAccountId := myAccountId
	treasuryKey := myPrivateKey
	fmt.Printf("Treasury id: %v\n Treasury pvk: %v\n========\n", treasuryAccountId, treasuryKey)

	// CONFIGURE VALID START TIME
	var validStartTime time.Time = time.Now() /*.Add(time.Minute * time.Duration(2))*/

	transactionId := hedera.NewTransactionIDWithValidStart(myAccountId, validStartTime)

	tokenCreateTransaction, err := hedera.NewTokenCreateTransaction().
		SetTransactionID(transactionId).
		SetTokenName("xidr-testnet-Oct12").                        // token name
		SetTokenSymbol("xidr-testnet-Oct12").                      // token symbol
		SetTreasuryAccountID(treasuryAccountId).                   // treasury account id
		SetAdminKey(CreateAdminKeyListDeployTestnetXidr()).        // owner
		SetSupplyKey(CreateSupplyKeyListDeployTestnetXidr()).      // minter/burner
		SetFreezeKey(CreateBlacklisterKeyListDeployTestnetXidr()). // blacklister
		//SetWipeKey(CreateWiperKeyListDeployTestnetxidr()).         // wiper
		//SetKycKey(CreateKycKeyListDeployTestnetxidr()).            // KYC
		//SetFeeScheduleKey(CreateFeeKeyListDeployTestnetxidr()).    // Fee
		SetInitialSupply(100000000). // total supply
		//SetDecimals(6).                                            // decimal
		//SetMaxTransactionFee(hedera.NewHbar(30)).                  // Change the default max transaction fee
		FreezeWith(client)
		// scheduled transaction doesn't neeed to be freeze

		/*if err != nil {
			panic(err)
		}
		*/

	// NON-SCHEDULED TRANSACTION APPROACH
	var temp *hedera.TokenCreateTransaction
	temp = tokenCreateTransaction.Sign(adminPrivateKey1).Sign(adminPrivateKey2).Sign(treasuryKey)
	txBytes, err := temp.ToBytes()
	fmt.Printf("%#v\n", hex.EncodeToString(txBytes))
}

func main() {
	create_token_tx_hedera_testnet_xidr()
}
