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
	//// Lb4 pubKey: 302a300506032b6570032100424d2b3f9ec5d189bf24515ced88cdef28725b2fa32eb31023c551c56eb76f2f
	//topicName := "Private Topic 2"
	//maxFee := 5.0
	//topicID, err := hcsClient.CreateTopic([]string{
	//	"302a300506032b6570032100cf97438ddf5769e6dcd674e9fbf22fd30d10e6b727e9cca4edb1c69e1c555a5c",
	//	"302a300506032b6570032100424d2b3f9ec5d189bf24515ced88cdef28725b2fa32eb31023c551c56eb76f2f",
	//}, topicName, maxFee)
	//
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
	topicID := "0.0.160549"
	err := hcsClient.SubmitMessage(topicID, []byte("Only Lb4 and Vnc can submit"), "zdr bepce")
	if err != nil {
		panic(err)
	}

	//time.Sleep(20 * time.Second)

	memo, _ := hcsClient.GetTopicInfo(topicID)

	fmt.Println(memo)


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
