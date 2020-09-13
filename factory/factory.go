package factory

import (
	"github.com/SrmHitter9062/go-factory/fault"
	"github.com/SrmHitter9062/go-factory/shutdown"
	"github.com/SrmHitter9062/go-factory/start"
)

// MsgInterface for all kind of messages
type MsgInterface interface {
	Process(msgType string, data []byte) error
}

// Factory object
type Factory struct{}

// GetFactory method for getting dynamic message object
func GetFactory(msgType string, packat []byte) (MsgInterface, error) {
	switch msgType {
	case "start":
		startHandle := start.New(packat)
		return startHandle, nil
	case "shutdown":
		shutDownHandle := shutdown.New(packat)
		return shutDownHandle, nil
	case "fault":
		faultHanle := fault.New(packat)
		return faultHanle, nil
	default:
		startHandle := start.New(packat)
		return startHandle, nil
	}

}
