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
	Timestamp   time.Time `json:"timestamp"`
	Redemption  string `json:"redemption"`
	Channel_id  string `json:"channel_id"`
	Redeemed_at string `json:"redeemed_at"`
	Reward      string `json"reward"`
	User_input  string `json:"user_input"`
	Status      string `json:"status"`
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
