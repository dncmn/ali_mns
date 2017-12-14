package ali_mns

import (
	"encoding/xml"
)

type MessageTopicSendRequest struct {
	XMLName           xml.Name    `xml:"Message"`
	MessageBody       Base64Bytes `xml:"MessageBody"`
	MessageTag        string      `xml:"MessageTag"`
	MessageAttributes string      `xml:"MessageAttributes"`
}

type MessageTopicSendResponse struct {
	MessageResponse
	MessageId      string `xml:"MessageId" json:"message_id"`
	MessageBodyMD5 string `xml:"MessageBodyMD5" json:"message_body_md5"`
}

type TopicAttribute struct {
	XMLName                xml.Name `xml:"Topic" json:"-"`
	TopicName              string   `xml:"TopicName,omitempty" json:"topic_name,omitempty"`
	LoggingEnabled         bool     `xml:"LoggingEnabled,omitempty" json:"logging_enabled,omitempty"`
	MaxMessageSize         int32    `xml:"MaximumMessageSize,omitempty" json:"maximum_message_size,omitempty"`
	MessageCount           int32    `xml:"MessageCount,omitempty" json:"message_count,omitempty"`
	MessageRetentionPeriod int32    `xml:"MessageRetentionPeriod,omitempty" json:"message_retention_period,omitempty"`
	CreateTime             int64    `xml:"CreateTime,omitempty" json:"create_time,omitempty"`
	LastModifyTime         int64    `xml:"LastModifyTime,omitempty" json:"last_modify_time,omitempty"`
}
