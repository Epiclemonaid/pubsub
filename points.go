package twitchpubsub

import (
	"encoding/json"
//	"errors"
	"fmt"
//	"strings"
	"time"
)

const pointsActionTopicPrefix = "channel-points-channel-v1."

// ModerationAction describes an incoming "Moderation" action coming from Twitch's PubSub servers
type PointsAction struct {
	Timestamp   time.Time `json:"timestamp"`
	Redemption  string `json:"redemption"`
	User_input  string `json:"user_input"`
	Status      string `json:"status"`
	Channel_id  string `json:"channel_id"`
	Redeemed_at string `json:"redeemed_at"`
	Reward      string `json"reward"`
	
}

type outerPointsAction struct {
	Data PointsAction `json:"data"`
}

func parsePointsAction(bytes []byte) (*PointsAction, error) {
	data := &outerPointsAction{}
	err := json.Unmarshal(bytes, data)
	if err != nil {
		return nil, err
	}

	return &data.Data, nil
}

//func parseChannelIDFromPointsTopic(topic string) (string, error) {
//	parts := strings.Split(topic, ".")
//	if len(parts) != 3 {
//		return "", errors.New("Unable to parse channel ID from Points topic")
//	}
//
//	return parts[2], nil
//}

// ModerationActionTopic returns a properly formatted moderation action topic string with the given user and channel ID arguments
func PointsActionTopic(channelID string) string {
	const f = `channel-points-channel-v1.%s`
	return fmt.Sprintf(f, channelID)
}
