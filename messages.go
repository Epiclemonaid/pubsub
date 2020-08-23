package twitchpubsub

import "strings"

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

	return messageTypeUnknown
}
