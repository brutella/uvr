package main

import (
	"github.com/brutella/can"
	"github.com/brutella/canopen"
	"github.com/brutella/uvr"
	"log"
	"net"
)

func HandleCANopen(frame can.Frame) {
	log.Printf("%X % X\n", frame.ID, frame.Data)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	iface, err := net.InterfaceByName("can0")

	if err != nil {
		log.Fatal(err)
	}

	conn, err := can.NewReadWriteCloserForInterface(iface)

	if err != nil {
		log.Fatal(err)
	}

	bus := can.NewBus(conn)
    // bus.SubscribeFunc(HandleCANopen)
	go bus.ConnectAndPublish()

	nodeID := canopen.NodeID(0x10)
    uvrID := canopen.NodeID(0x1)
    
    c := uvr.NewClient(nodeID, bus)
	c.Connect(uvrID)

	solar := uvr.NewMixerHeating1()
	if value, err := c.Read(solar.Description); err == nil {
		log.Println("Description", value)
	}

	if value, err := c.Read(solar.StartDelay); err == nil {
		log.Println("Start Delay", value)
	}

	if value, err := c.Read(solar.RunOnTime); err == nil {
		log.Println("Run On Time", value)
	}

	if value, err := c.Read(solar.Mode); err == nil {
		log.Println("Mode", value)
	}

	if value, err := c.Read(canopen.NewObjectIndex(0x20aa, 0x9)); err == nil {
		log.Println("State", value)
	}

	if value, err := c.Read(solar.SpeedStage); err == nil {
		log.Println("Speed Stage", value)
	}
    
    c.Disconnect()
    bus.Disconnect()
}
