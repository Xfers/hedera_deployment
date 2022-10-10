package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

func broadcast_mint_transaction_testnet(txBytes string, signature1 []byte, signature2 []byte, signature3 []byte) {
	var _txBytes []byte

	myAccountId, err := hedera.AccountIDFromString("0.0.45949322")

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

	supplyPublicKey1, err := hedera.PublicKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_SUPPLY_PUBLIC_KEY1"))
	if err != nil {
		panic(err)
	}

	supplyPublicKey2, err := hedera.PublicKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_SUPPLY_PUBLIC_KEY2"))
	if err != nil {
		panic(err)
	}

	client := hedera.ClientForTestnet()
	client.SetOperator(myAccountId, myPrivateKey)

	_txBytes, err = hex.DecodeString(txBytes)
	if err != nil {
		fmt.Println("error with txBytes")
	}

	var tx interface{}
	//var createTokenTx hedera.TokenCreateTransaction
	var mintTokenTx hedera.TokenMintTransaction
	var signedMintTokenTx *hedera.TokenMintTransaction

	tx, err = hedera.TransactionFromBytes(_txBytes)
	if err != nil {
		fmt.Println("error with txBytes")
		panic(err)
	}

	mintTokenTx = tx.(hedera.TokenMintTransaction)
	signedMintTokenTx = mintTokenTx.AddSignature(supplyPublicKey1, signature1).AddSignature(supplyPublicKey2, signature1)

	response, err := signedMintTokenTx.Execute(client)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", "response")
	fmt.Println(response)
	receipt, err := response.GetReceipt(client)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", "receipt")
	fmt.Printf("%+v\n", receipt)

	//Get the token ID from the receipt
	tokenId := *receipt.TokenID

	fmt.Printf("The new token ID is %v\n", tokenId)

}

func broadcast_burn_transaction_testnet(txBytes string, signature1 []byte, signature2 []byte, signature3 []byte) {
	var _txBytes []byte

	myAccountId, err := hedera.AccountIDFromString("0.0.45949322")

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

	supplyPublicKey1, err := hedera.PublicKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_SUPPLY_PUBLIC_KEY1"))
	if err != nil {
		panic(err)
	}

	supplyPublicKey2, err := hedera.PublicKeyFromString(os.Getenv("HEDERA_XSGD_MAINNET_SUPPLY_PUBLIC_KEY2"))
	if err != nil {
		panic(err)
	}

	client := hedera.ClientForTestnet()
	client.SetOperator(myAccountId, myPrivateKey)

	_txBytes, err = hex.DecodeString(txBytes)
	if err != nil {
		fmt.Println("error with txBytes")
	}

	var tx interface{}
	//var createTokenTx hedera.TokenCreateTransaction
	var burnTokenTx hedera.TokenBurnTransaction
	var signedBurnTokenTx *hedera.TokenBurnTransaction

	tx, err = hedera.TransactionFromBytes(_txBytes)
	if err != nil {
		fmt.Println("error with txBytes")
		panic(err)
	}

	burnTokenTx = tx.(hedera.TokenBurnTransaction)
	signedBurnTokenTx = burnTokenTx.AddSignature(supplyPublicKey1, signature1).AddSignature(supplyPublicKey2, signature1)

	response, err := signedBurnTokenTx.Execute(client)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", "response")
	fmt.Println(response)
	receipt, err := response.GetReceipt(client)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", "receipt")
	fmt.Printf("%+v\n", receipt)

	//Get the token ID from the receipt
	tokenId := *receipt.TokenID

	fmt.Printf("The new token ID is %v\n", tokenId)
}

func broadcast_deploy_transaction_testnet(txBytes string, signature1 []byte, signature2 []byte, signature3 []byte) {
	var _txBytes []byte

	myAccountId, err := hedera.AccountIDFromString("0.0.45949322")

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

	adminPublicKey1, _ := hedera.PublicKeyFromString("302a300506032b657003210009fc854d1005984c386206cc68382900a0f9647d7a0ddd164d27f9dcd767f1f6")
	adminPublicKey2, _ := hedera.PublicKeyFromString("5c97a1e4dff419af615f1085ed533e35514e347342ad7fbc79820fb234d6e9e2")

	treasurerPublicKey := adminPublicKey1

	client := hedera.ClientForTestnet()
	client.SetOperator(myAccountId, myPrivateKey)

	_txBytes, err = hex.DecodeString(txBytes)
	if err != nil {
		fmt.Println("error with txBytes")
	}

	var tx interface{}
	//var createTokenTx hedera.TokenCreateTransaction
	var createTokenTx hedera.TokenCreateTransaction
	var signedCreateTokenTx *hedera.TokenCreateTransaction

	tx, err = hedera.TransactionFromBytes(_txBytes)
	if err != nil {
		fmt.Println("error with txBytes")
		panic(err)
	}

	createTokenTx = tx.(hedera.TokenCreateTransaction)
	signedCreateTokenTx = createTokenTx.AddSignature(adminPublicKey1, signature1).AddSignature(adminPublicKey2, signature2).AddSignature(treasurerPublicKey, signature3)

	response, err := signedCreateTokenTx.Execute(client)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", "response")
	fmt.Println(response)
	receipt, err := response.GetReceipt(client)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", "receipt")
	fmt.Printf("%+v\n", receipt)

	//Get the token ID from the receipt
	tokenId := *receipt.TokenID

	fmt.Printf("The new token ID is %v\n", tokenId)
}

func broadcast_transfer_transaction_testnet(txBytes string) {
	var _txBytes []byte

	myAccountId, err := hedera.AccountIDFromString("0.0.45949322")

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

	client := hedera.ClientForTestnet()
	client.SetOperator(myAccountId, myPrivateKey)

	_txBytes, err = hex.DecodeString(txBytes)
	if err != nil {
		fmt.Println("error with txBytes")
	}

	var signedTx interface{}
	var signedTransferTx *hedera.TransferTransaction

	signedTx, err = hedera.TransactionFromBytes(_txBytes)
	if err != nil {
		panic(err)
	}
	signedTransferTx = signedTx.(*hedera.TransferTransaction)
	response, err := signedTransferTx.Execute(client)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", "response")
	fmt.Println(response)
	receipt, err := response.GetReceipt(client)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", "receipt")
	fmt.Printf("%+v\n", receipt)

	//Get the token ID from the receipt
	tokenId := *receipt.TokenID

	fmt.Printf("The new token ID is %v\n", tokenId)

}

func broadcast_update_transaction_testnet(txBytes string) {
	var _txBytes []byte

	myAccountId, err := hedera.AccountIDFromString("0.0.45949322")

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

	client := hedera.ClientForTestnet()
	client.SetOperator(myAccountId, myPrivateKey)

	_txBytes, err = hex.DecodeString(txBytes)
	if err != nil {
		fmt.Println("error with txBytes")
	}

	var signedTx interface{}
	var signedTokenUpdateTx *hedera.TokenUpdateTransaction

	signedTx, err = hedera.TransactionFromBytes(_txBytes)
	if err != nil {
		panic(err)
	}
	signedTokenUpdateTx = signedTx.(*hedera.TokenUpdateTransaction)
	response, err := signedTokenUpdateTx.Execute(client)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", "response")
	fmt.Println(response)
	receipt, err := response.GetReceipt(client)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", "receipt")
	fmt.Printf("%+v\n", receipt)

	//Get the token ID from the receipt
	tokenId := *receipt.TokenID

	fmt.Printf("The new token ID is %v\n", tokenId)
}

func main() {
	//panic: invalid compressed ecsda public key length: 0 bytes
	broadcast_deploy_transaction_testnet("0a9e032a9b030a96030a140a0b089097999a0610d0b2a3281205188ac3f415120218041880d0acf30e22020878ea01ee020a12585347442d746573746e65742d4f637431321212585347442d746573746e65742d4f637431322080c2d72f2a05188ac3f415324a32480a22122009fc854d1005984c386206cc68382900a0f9647d7a0ddd164d27f9dcd767f1f60a2212205c97a1e4dff419af615f1085ed533e35514e347342ad7fbc79820fb234d6e9e2426e326c0a221220068f0eab7ae589debf5d0ddf7db1b0acf844c254ce472a4bddcc7b04f07d8dc40a221220f556b42ed6d5a7a95bf5c44369296c14984831eaa2010d6a3b0450fac5ec5d030a22122099945810ccac97cf9fc98877dad6028576cd901e967830a98ea5e2b10bdbeb01526e326c0a221220a2ca56b83a3e2011e1b81316553dfbd4e13777e2924d55edfff07651914945200a2212205832813ab4ff848a60dfa7ef55f8930855ceda7bf34f6b2a2bd5b97b092e14e80a221220cd3b39d33a3eea3dfdbfa86a7f938fb42be6f702a47799c13c3a3bddfe516b527205188ac3f4157a0508d0c8e10312000a9e032a9b030a96030a140a0b089097999a0610d0b2a3281205188ac3f415120218041880d0acf30e22020878ea01ee020a12585347442d746573746e65742d4f637431321212585347442d746573746e65742d4f637431322080c2d72f2a05188ac3f415324a32480a22122009fc854d1005984c386206cc68382900a0f9647d7a0ddd164d27f9dcd767f1f60a2212205c97a1e4dff419af615f1085ed533e35514e347342ad7fbc79820fb234d6e9e2426e326c0a221220068f0eab7ae589debf5d0ddf7db1b0acf844c254ce472a4bddcc7b04f07d8dc40a221220f556b42ed6d5a7a95bf5c44369296c14984831eaa2010d6a3b0450fac5ec5d030a22122099945810ccac97cf9fc98877dad6028576cd901e967830a98ea5e2b10bdbeb01526e326c0a221220a2ca56b83a3e2011e1b81316553dfbd4e13777e2924d55edfff07651914945200a2212205832813ab4ff848a60dfa7ef55f8930855ceda7bf34f6b2a2bd5b97b092e14e80a221220cd3b39d33a3eea3dfdbfa86a7f938fb42be6f702a47799c13c3a3bddfe516b527205188ac3f4157a0508d0c8e10312000a9e032a9b030a96030a140a0b089097999a0610d0b2a3281205188ac3f415120218041880d0acf30e22020878ea01ee020a12585347442d746573746e65742d4f637431321212585347442d746573746e65742d4f637431322080c2d72f2a05188ac3f415324a32480a22122009fc854d1005984c386206cc68382900a0f9647d7a0ddd164d27f9dcd767f1f60a2212205c97a1e4dff419af615f1085ed533e35514e347342ad7fbc79820fb234d6e9e2426e326c0a221220068f0eab7ae589debf5d0ddf7db1b0acf844c254ce472a4bddcc7b04f07d8dc40a221220f556b42ed6d5a7a95bf5c44369296c14984831eaa2010d6a3b0450fac5ec5d030a22122099945810ccac97cf9fc98877dad6028576cd901e967830a98ea5e2b10bdbeb01526e326c0a221220a2ca56b83a3e2011e1b81316553dfbd4e13777e2924d55edfff07651914945200a2212205832813ab4ff848a60dfa7ef55f8930855ceda7bf34f6b2a2bd5b97b092e14e80a221220cd3b39d33a3eea3dfdbfa86a7f938fb42be6f702a47799c13c3a3bddfe516b527205188ac3f4157a0508d0c8e1031200",
		[]byte{0x34, 0xf, 0x1e, 0x34, 0xa8, 0x33, 0x30, 0xd4, 0x19, 0xb8, 0xb3, 0x40, 0xdd, 0xa8, 0xf8, 0xac, 0x71, 0x95, 0x77, 0x74, 0x69, 0xa6, 0xce, 0xf0, 0x86, 0xa6, 0x32, 0xe, 0x87, 0x69, 0xc5, 0xf, 0xfb, 0xf9, 0xea, 0x46, 0x42, 0xec, 0x88, 0x2f, 0x1b, 0x18, 0x39, 0x72, 0x57, 0x6a, 0x3, 0x54, 0x6c, 0x26, 0x18, 0x1, 0x2d, 0x25, 0xae, 0x7b, 0x9d, 0xaf, 0xe, 0xd4, 0x19, 0xa4, 0xa5, 0x5},
		[]byte{0xc9, 0xbb, 0x71, 0xe9, 0xec, 0xac, 0x20, 0x34, 0x22, 0x78, 0x27, 0x22, 0x4a, 0x36, 0xac, 0xc1, 0x18, 0x32, 0x57, 0xad, 0x69, 0xda, 0x3, 0x4f, 0x85, 0xa4, 0xbd, 0x97, 0xb7, 0x50, 0xd6, 0x64, 0x73, 0xa3, 0xcd, 0x10, 0xf7, 0xfa, 0x91, 0xf3, 0x79, 0x98, 0xf0, 0x77, 0x5a, 0x90, 0x4f, 0x89, 0x3c, 0xb4, 0xd9, 0xa8, 0x14, 0x1d, 0x3f, 0xcb, 0xa5, 0x4b, 0x6, 0xb9, 0x7d, 0xce, 0x59, 0x4},
		[]byte{})
}
