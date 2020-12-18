package main

import (
	"fmt"
	//"fmt"
	hcs "github.com/Holymir/provide/sdk/hcs-client"
	"github.com/hashgraph/hedera-sdk-go"
	"github.com/joho/godotenv"
	"os"
	//"time"
)

func main() {
	//topicId := "0.0.159186"
	godotenv.Load()
	client := SetupClient(os.Getenv("ACCOUNT_ID"), os.Getenv("PRIVATE_KEY"))

	hcsClient := hcs.NewHCSClient(client)

	// TODO: Create Topic
	//topicID, err := hcsClient.CreateTopic([]string{"302a300506032b6570032100424d2b3f9ec5d189bf24515ced88cdef28725b2fa32eb31023c551c56eb76f2f"}, "topic memo new", 5)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("The new topic ID is %v\n", topicID)

	// TODO: Subscribe
	//err := hcsClient.SubscribeToTopic("0.0.159186", func(message hedera.TopicMessage) {
	//	fmt.Println(message.ConsensusTimestamp.String(), " received topic message:", string(message.Contents))
	//})

	//
	//if err != nil {
	//	panic(err)
	//}

	// TODO: Submit message
	//err := hcsClient.SubmitMessage("0.0.159256", []byte("Only Lb4 and Vnc can submit"), "zdr bepce")
	//if err != nil {
	//	panic(err)
	//}

	//time.Sleep(20 * time.Second)
	//
	submitKey, _ := hcsClient.GetTopicInfo("0.0.159256")

	fmt.Println(submitKey)


}

func SetupClient(accountID, privateKey string) *hedera.Client {
	client := hedera.ClientForTestnet()
	accID, err := hedera.AccountIDFromString(accountID)

	if err != nil {
		panic(err)
	}

	pK, err := hedera.PrivateKeyFromString(privateKey)

	if err != nil {
		panic(err)
	}

	client.SetOperator(accID, pK)
	return client
}
