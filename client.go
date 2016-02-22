package uvr

import(
    "github.com/brutella/can"
    "github.com/brutella/canopen"
    "time"
)

type Client struct {
    id canopen.NodeID
    bus *can.Bus
    heartbeat chan <- struct{}
}

func NewClient(id canopen.NodeID, bus *can.Bus) *Client {
    return &Client{id, bus, nil,}
}

func (c *Client) Connect(id canopen.NodeID) error {
	c.heartbeat = canopen.ProduceHeartbeat(c.id, canopen.Operational, c.bus, time.Second*10)    
    return Connect(id, c.id, c.bus)
}

func (c *Client) Disconnect(id canopen.NodeID) error {
	c.heartbeat <- struct{}{}
    return Disconnect(id, c.id, c.bus)
}

func (c *Client) Read(i canopen.ObjectIndex) (interface{}, error) {
	return ReadFromIndex(i, c.id, c.bus)
}

func (c *Client) Write(b []byte, i canopen.ObjectIndex) error {
    return WriteToIndex(i, b, c.id, c.bus)
}