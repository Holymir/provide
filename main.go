package main

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go"
)

func main() {
	client := setupclient()
	privKey1, err := cryptocreate(client)
	if err != nil {
		panic(err)
	}
	privKey2, err := cryptocreate(client)
	if err != nil {
		panic(err)
	}
	privKey3, err := cryptocreate(client)
	if err != nil {
		panic(err)
	}
	k := hedera.KeyListWithThreshold(1)
	k = k.Add(privKey1.PublicKey()).Add(privKey2.PublicKey()).Add(privKey3.PublicKey())
	fmt.Println(k.String())
	txID, err := hedera.NewTopicCreateTransaction().SetSubmitKey(k).Execute(client)
	if err != nil {
		panic(err)
	}
	fmt.Println(txID)
	fmt.Println(privKey1.String())
	fmt.Println(privKey2.String())
	fmt.Println(privKey3.String())
}
func cryptocreate(client *hedera.Client) (hedera.PrivateKey, error) {
	privateKey, _ := hedera.GeneratePrivateKey()
	fmt.Println(privateKey.String())
	publicKey := privateKey.PublicKey()
	newAccount, err := hedera.NewAccountCreateTransaction().
		SetKey(publicKey).
		SetInitialBalance(hedera.NewHbar(100)).
		Execute(client)
	if err != nil {
		panic(err)
	}
	fmt.Println(newAccount.TransactionID)
	return privateKey, nil
}
func setupclient() *hedera.Client {
	client := hedera.ClientForTestnet()
	accID, _ := hedera.AccountIDFromString("0.0.143980")
	pK, _ := hedera.PrivateKeyFromString("302e020100300506032b657004220420e08b73b549f6a7222e91bdf641d5576f2864e70071c7c2e13ffd102844afd107")
	client.SetOperator(accID, pK)
	return client
}
