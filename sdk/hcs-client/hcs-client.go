package client

import (
	"fmt"
	"github.com/hashgraph/hedera-sdk-go"
)

type HCSClient struct {
	Client *hedera.Client
}

func (hcsc *HCSClient) CreateTopic(publicKeys []string, memo string, maxFee float64) (topicID string, err error) {
	transaction := hedera.NewTopicCreateTransaction().
		SetAdminKey(hcsc.Client.GetOperatorPublicKey()).
		SetTransactionMemo(memo).
		SetMaxTransactionFee(hedera.HbarFrom(maxFee, hedera.HbarUnits.Hbar))

	if len(publicKeys) > 0 {
		k := hedera.KeyListWithThreshold(1)

		for i := 0; i < len(publicKeys); i++ {
			publicKey, err := hedera.PublicKeyFromString(publicKeys[i])
			if err != nil {
				return "", err
			}
			k = k.Add(publicKey)
		}
		transaction.SetSubmitKey(k)
	}

	txResponse, err := transaction.Execute(hcsc.Client)
	if err != nil {
		return "", err
	}


	transactionReceipt, err := txResponse.GetReceipt(hcsc.Client)

	if err != nil {
		return "", err
	}

	newTopicID := *transactionReceipt.TopicID

	return newTopicID.String(), nil
}

func (hcsc *HCSClient) SubmitMessage(topicID string, message []byte, memo string) (err error) {
	topicIDFromString, err := hedera.TopicIDFromString(topicID)
	if err != nil {
		return err
	}

	transaction := hedera.NewTopicMessageSubmitTransaction().
		SetTopicID(topicIDFromString).
		SetTransactionMemo(memo).
		SetMessage(message)

	txResponse, err := transaction.Execute(hcsc.Client)
	if err != nil {
		return err
	}

	// TODO: handle the status
	txReceipt, err := txResponse.GetReceipt(hcsc.Client)
	if err != nil {
		return err
	}

	fmt.Println(txReceipt.Status)
	return nil
}

func (hcsc *HCSClient) SubscribeToTopic(topicID string, callback func(hedera.TopicMessage)) (err error) {
	topicIDFromString, err := hedera.TopicIDFromString(topicID)
	if err != nil {
		return err
	}

	_, err = hedera.NewTopicMessageQuery().
		SetTopicID(topicIDFromString).
		Subscribe(hcsc.Client, callback)

	if err != nil {
		return err
	}

	return nil
}

func (hcsc *HCSClient) GetTopicInfo(topicID string) (memo string, err error){
	topicIDFromString, err := hedera.TopicIDFromString(topicID)
	if err != nil {
		panic(err)
		return "", err
	}

	topicInfo, err := hedera.NewTopicInfoQuery().
		SetTopicID(topicIDFromString).
		Execute(hcsc.Client)
	if err != nil {
		panic(err)
		return "", err
	}

	fmt.Println(topicInfo)
	return topicInfo.Memo, nil
}


func NewHCSClient(client *hedera.Client) *HCSClient {
	return &HCSClient{
		Client: client,
	}
}
