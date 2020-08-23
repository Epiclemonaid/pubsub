package twitchpubsub

// Helper functions and structures for twitch bits events

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

const pointsEventTopicPrefix = "channel-points-channel-v1."

// BitsEvent describes an incoming "Bit" action coming from Twitch's PubSub servers
type PointsEvent struct {
	UserName         string    `json:"user_name"`
	ChannelName      string    `json:"channel_name"`
	UserID           string    `json:"user_id"`
	ChannelID        string    `json:"channel_id"`
	Time             time.Time `json:"time"`
	ChatMessage      string    `json:"chat_message"`
	BitsUsed         int       `json:"bits_used"`
	TotalBitsUsed    int       `json:"total_bits_used"`
	Context          string    `json:"context"`
	BadgeEntitlement struct {
		NewVersion      int `json:"new_version"`
		PreviousVersion int `json:"previous_version"`
	} `json:"badge_entitlement"`
}

type outerPointsEvent struct {
	Data PointsEvent `json:"data"`
}

func parsePointsEvent(bytes []byte) (*PointsEvent, error) {
	data := &outerPointsEvent{}
	err := json.Unmarshal(bytes, data)
	if err != nil {
		return nil, err
	}

	return &data.Data, nil
}

func parseChannelIDFromPointsTopic(topic string) (string, error) {
	parts := strings.Split(topic, ".")
	if len(parts) != 2 {
		return "", errors.New("Unable to parse channel ID from bits topic")
	}

	return parts[1], nil
}

// BitsEventTopic returns a properly formatted bits event topic string with the given channel ID argument
func PointsEventTopic(channelID string) string {
	const f = `channel-points-channel-v1.%s`
	return fmt.Sprintf(f, channelID)
}
