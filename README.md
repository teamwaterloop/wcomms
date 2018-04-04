# WComms

Communication packet protocol for sending and serializing pod data across multiple servers

## Usage

#### type CommPacketJson

``` go
type CommPacketJson struct {
	Time int64     `json:"time"`
	Type string    `json:"type"`
	Id   uint8     `json:"name"`
	Data []float32 `json:"data"`
}
```

#### func PacketEncodeJson

``` go
func PacketEncodeJson(packet *wbin.CommPacket) ([]byte, error)
```

Serializes binary communication packet to JSON string/bytes

#### func PacketDecodeJson

``` go
func PacketDecodeJson(encoded []byte) (*wbin.CommPacket, error)
```

Deserializes JSON string/bytes to binary communication packet
