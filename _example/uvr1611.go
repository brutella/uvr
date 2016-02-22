package main

import (
	"github.com/brutella/can"
	"github.com/brutella/canopen"
	"github.com/brutella/uvr"
	"log"
	"net"
	"time"
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
	bus.SubscribeFunc(HandleCANopen)
	go bus.ConnectAndPublish()

	nodeID := canopen.NodeID(0x10)
	// Produce heartbeat
	stop := canopen.ProduceHeartbeat(nodeID, canopen.Operational, bus, time.Second*10)
	defer func() {
		stop <- struct{}{}
	}()

	// Wait
	<-time.After(time.Second * 1)

	// Connect to UVR1611
	id := canopen.NodeID(0x1)
	uvr.ConnectTo(id, nodeID, bus)

	solar := uvr.NewMixerHeating1()
	if value, err := uvr.ReadFromIndex(solar.Description, nodeID, bus); err == nil {
		log.Println("Description", value)
	}

	if value, err := uvr.ReadFromIndex(solar.StartDelay, nodeID, bus); err == nil {
		log.Println("Start Delay", value)
	}

	if value, err := uvr.ReadFromIndex(solar.RunOnTime, nodeID, bus); err == nil {
		log.Println("Run On Time", value)
	}

	if value, err := uvr.ReadFromIndex(solar.Mode, nodeID, bus); err == nil {
		log.Println("Mode", value)
	}

	if value, err := uvr.ReadFromIndex(canopen.NewObjectIndex(0x20aa, 0x9), nodeID, bus); err == nil {
		log.Println("State", value)
	}

	if value, err := uvr.ReadFromIndex(solar.SpeedStage, nodeID, bus); err == nil {
		log.Println("Speed Stage", value)
	}

	// 0x2 + WRITE (String) + 0x91 (Datatype)
	written := []byte{0x2, 0x57, 0x52, 0x49, 0x54, 0x45, 0x91}
	if err := uvr.WriteToIndex(solar.State, written, nodeID, bus); err != nil {
		log.Fatal(err)
	}
}
