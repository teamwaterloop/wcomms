package wjson

import (
	wbin "github.com/teamwaterloop/wcomms/wbinary"
	"time"
	"encoding/json"
)

type CommPacketJson struct {
	Time int64     `json:"time"`
	Type string    `json:"type"`
	Id   uint8     `json:"name"`
	Data []float32 `json:"data"`
}

func CurrentTimeMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func PacketEncodeJson(packet *wbin.CommPacket) ([]byte, error) {
	packetJson := &CommPacketJson{
		Time: CurrentTimeMs(),
		Type: wbin.TypeToString(packet.PacketType),
		Id:   packet.PacketId,
		Data: []float32{packet.Data1, packet.Data2, packet.Data3},
	}
	return json.Marshal(packetJson)
}

func PacketDecodeJson(encoded []byte) (*wbin.CommPacket, error) {
	packetJson := &CommPacketJson{}
	err := json.Unmarshal(encoded, packetJson)
	packet := &wbin.CommPacket{
		PacketType: wbin.StringToType(packetJson.Type),
		PacketId:   packetJson.Id,
		Data1:      packetJson.Data[0],
		Data2:      packetJson.Data[1],
		Data3:      packetJson.Data[2],
	}
	return packet, err
}
