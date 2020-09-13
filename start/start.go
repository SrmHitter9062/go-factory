package start

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/SrmHitter9062/go-factory/base"
)

type Message struct {
	base.Message
	Msg       string    `json:"msg"`
	MsgType   string    `json:"msg_type"`
	StartTime time.Time `json:"start_time,omitempty"`
}

// New function to get object of start Message
func New(packat []byte) *Message {
	var msg *Message
	err := json.Unmarshal(packat, &msg)
	if err != nil {
		return nil
	}
	return msg
}

// Process the start message
func (m *Message) Process(msgType string, data []byte) error {
	fmt.Printf("event type :%s, message is: %s\n", m.MsgType, m.Msg)
	// some code
	return nil
}
