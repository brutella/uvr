// uvrinflux queries an UVR1611 via can bus and prints the data to the console.
// The output is formatted in the influxdb [line protocol][1].
//
// [1]: https://docs.influxdata.com/influxdb/latest/write_protocols/line_protocol_tutorial/
package main

import (
	"flag"
	"fmt"
	"github.com/brutella/can"
	"github.com/brutella/uvr"
	"log"
	"strings"
)

const (
	MaxOutletsCount = 13
	MaxInletsCount  = 16
)

func collect(client *uvr.Client) (map[string]string, error) {
	values := map[string]interface{}{}

	var desc interface{}
	var val interface{}
	var err error

	for i := uint8(1); i <= MaxOutletsCount; i++ {
		out := uvr.NewOutlet(i)

		if desc, err = client.Read(out.Description); err != nil {
			return nil, err
		}

		str := strings.TrimSpace(desc.(string))

		if strings.HasSuffix(str, uvr.DescriptionUnused) {
			continue
		}

		if val, err = client.Read(out.State); err != nil {
			return nil, err
		}

		values[str] = val
	}

	for i := uint8(1); i <= MaxInletsCount; i++ {
		in := uvr.NewInlet(i)

		if desc, err = client.Read(in.Description); err != nil {
			return nil, err
		}

		str := strings.TrimSpace(desc.(string))

		if strings.HasSuffix(str, uvr.DescriptionUnused) {
			continue
		}

		if val, err = client.Read(in.Value); err != nil {
			return nil, err
		}

		values[str] = val
	}

	var dict = map[string]string{}
	for key, value := range values {
		if str, ok := value.(string); ok == true {
			if v, err := uvr.StringToBool(str); err == nil {
				if v {
					dict[key] = "true"
				} else {
					dict[key] = "false"
				}
			} else {
				dict[key] = str
			}
		} else if v, ok := value.(float32); ok == true {
			dict[key] = fmt.Sprintf("%.2f", v)
		}
	}

	return dict, nil
}

func main() {
	var (
		clientId = flag.Int("client_id", 16, "id of the client; range from [1...254]")
		serverId = flag.Int("server_id", 1, "id of the server to which the client connects to: range from [1...254]")
		iface    = flag.String("if", "can0", "name of the can network interface")
	)

	flag.Parse()

	bus, err := can.NewBusForInterfaceWithName(*iface)

	if err != nil {
		log.Fatal(err)
	}
	go bus.ConnectAndPublish()

	nodeID := uint8(*clientId)
	uvrID := uint8(*serverId)

	c := uvr.NewClient(nodeID, bus)
	c.Connect(uvrID)

	if dict, err := collect(c); err == nil {
		var pairs = []string{}
		for key, value := range dict {
			pairs = append(pairs, fmt.Sprintf("%s=%s", strings.Replace(key, " ", "_", -1), value))
		}
		fmt.Printf("uvr,client_id=%d,server_id=%d,if=%s %s", *clientId, *serverId, *iface, strings.Join(pairs, ","))
	}

	c.Disconnect(uvrID)
	bus.Disconnect()
}
