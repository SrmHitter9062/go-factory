package main

import (
	"fmt"

	"github.com/SrmHitter9062/go-factory/factory"
)

func init() {
	fmt.Println("init function")
}

func main() {
	fmt.Println("In main function")
	startEvent := `{"msg": "device started","msg_type": "start"}`
	faultEvent := `{"msg": "device threw error","msg_type": "fault"}`
	shutDownEvent := `{"msg": "device shutdown","msg_type": "shutdown"}`
	mapData := map[string][]byte{
		"start":    []byte(startEvent),
		"fault":    []byte(faultEvent),
		"shutdown": []byte(shutDownEvent),
	}
	// send message for processing
	fmt.Println("Sending the messages for processing")
	for msgType, msg := range mapData {
		factoryObj, err := factory.GetFactory(msgType, msg)
		if err != nil {
			fmt.Println("Error in gettig factory object, message type:", msgType)
		}
		ferr := factoryObj.Process(msgType, msg)
		if ferr != nil {
			fmt.Println("Error in processing the message:", msgType)
		}

	}
}
