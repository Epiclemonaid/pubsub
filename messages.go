package twitchpubsub

import (
	"strings"
	"fmt"
)
// Base TODO: Refactor
type Base struct {
	Type string `json:"type"`
}

type messageType = int

const (
	messageTypeUnknown messageType = iota
	messageTypeModerationAction
	messageTypeBitsEvent
	messageTypePointsEvent
	messageTypePointsAction
)

func getMessageType(topic string) messageType {
	if strings.HasPrefix(topic, moderationActionTopicPrefix) {
		return messageTypeModerationAction
	}
	if strings.HasPrefix(topic, bitsEventTopicPrefix) {
		return messageTypeBitsEvent
	}
	if strings.HasPrefix(topic, pointsActionTopicPrefix) {
		return messageTypePointsAction
	}
	if strings.HasPrefix(topic, pointsEventTopicPrefix) {
		return messageTypePointsEvent
	}
	if topic == "reward-redeemed" {
		fmt.Println("reward redeemed type")
	}

	return messageTypeUnknown
}
