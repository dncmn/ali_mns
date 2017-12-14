package ali_mns

import (
	"fmt"
)

type AliMNSTopic interface {
	Name() string
	SendMessage(message MessageTopicSendRequest) (resp MessageSendResponse, err error)
}

type MNSTopic struct {
	name    string
	client  MNSClient
	decoder MNSDecoder
}

func NewMNSTopic(name string, client MNSClient, qps ...int32) AliMNSTopic {
	if name == "" {
		panic("ali_mns: topic name could not be empty")
	}
	topic := new(MNSTopic)
	topic.client = client
	topic.name = name
	topic.decoder = NewAliMNSDecoder()

	var attr TopicAttribute
	if _, err := send(client, topic.decoder, GET, nil, nil, "topics/"+name, &attr); err != nil {
		panic(err)
	}
	return topic
}
func (p *MNSTopic) Name() string {
	return p.name
}
func (p *MNSTopic) SendMessage(message MessageTopicSendRequest) (resp MessageSendResponse, err error) {
	_, err = send(p.client, p.decoder, POST, nil, message, fmt.Sprintf("topics/%s/%s", p.name, "messages"), &resp)
	return
}
