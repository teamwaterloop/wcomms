package wbinary

import (
	"math"
	"testing"
)

func TestSetBit(t *testing.T) {
	var buf = []byte{255, 255, 255, 255}
	setBits(buf, 629, 10, 10)
	if readBits(buf, 10, 19) != 629 {
		t.Fail()
	}
	buf = []byte{0, 0, 0, 0}
	setBits(buf, 629, 10, 10)
	if readBits(buf, 10, 19) != 629 {
		t.Fail()
	}
}

func TestReadBits(t *testing.T) {
	var buf = []byte{3, 172, 128, 99}
	for i := 1; i < 8; i++ {
		if readBits(buf, 0, uint(i)) != 3 {
			t.Fail()
		}
	}
	if readBits(buf, 0, 0) != 1 {
		t.Fail()
	}
	if readBits(buf, 1, 1) != 1 {
		t.Fail()
	}
	if readBits(buf, 1, 7) != 1 {
		t.Fail()
	}
	if readBits(buf, 1, 10) != 513 {
		t.Fail()
	}
	if readBits(buf, 0, 13) != 11267 {
		t.Fail()
	}
	if readBits(buf, 5, 15) != 1376 {
		t.Fail()
	}
	if readBits(buf, 0, 31) != 1669377027 {
		t.Fail()
	}
	if readBits(buf, 13, 30) != 203781 {
		t.Fail()
	}
	if readBits(buf, 30, 13) != 203781 {
		t.Fail()
	}
	if readBits(buf, 50, 51) != 0 {
		t.Fail()
	}
	if readBits(buf, 28, 50) != 6 {
		t.Fail()
	}
}

func TestReadSegments(t *testing.T) {
	var buf = []byte{218, 73, 85, 117}
	var shape = []uint{1, 5, 6, 6, 12, 2}
	var result = readSegments(buf, shape)
	var expected = []uint32{0, 13, 39, 20, 3413, 1}
	if len(expected) != len(result) {
		t.Errorf("Expected length %d but got %d", len(expected), len(result))
		t.Fail()
	}
	for i := 0; i < 6; i++ {
		if result[i] != expected[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
			t.Fail()
		}
	}
}

func TestReadSegmentsUnderFlow(t *testing.T) {
	var buf = []byte{218, 73, 85, 117}
	var shape = []uint{1, 5, 6, 6, 9}
	var result = readSegments(buf, shape)
	var expected = []uint32{0, 13, 39, 20, 341}
	if len(expected) != len(result) {
		t.Errorf("Expected length %d but got %d", len(expected), len(result))
		t.Fail()
	}
	for i := 0; i < 5; i++ {
		if result[i] != expected[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
			t.Fail()
		}
	}
}

func TestReadSegmentsOverFlow(t *testing.T) {
	var buf = []byte{218, 73, 85, 117}
	var shape = []uint{1, 5, 6, 6, 12, 12, 5}
	var result = readSegments(buf, shape)
	var expected = []uint32{0, 13, 39, 20, 3413, 1, 0}
	if len(expected) != len(result) {
		t.Errorf("Expected length %d but got %d", len(expected), len(result))
		t.Fail()
	}
	for i := 0; i < 7; i++ {
		if result[i] != expected[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
			t.Fail()
		}
	}
}

func TestWriteSegments(t *testing.T) {
	var buf = []byte{218, 73, 85, 117}
	var shape = []uint{1, 5, 6, 6, 12, 2}
	var result = readSegments(buf, shape)
	var nbuf = []byte{0, 0, 0, 0}
	writeSegments(nbuf, shape, result)
	for i := 0; i < 4; i++ {
		if nbuf[i] != buf[i] {
			t.Errorf("Expected %d but got %d at %d", buf[i], nbuf[i], i)
			t.Fail()
		}
	}
}

func TestDecodeFloat18(t *testing.T) {
	var result = decodeFloat18(231079)
	var expected float32 = -724.875
	if math.Abs(float64(result-expected)) > 0.00001 {
		t.Errorf("Expected %.2f but go %.2f", expected, result)
		t.Fail()
	}
	result = decodeFloat18(100980)
	expected = 846.5
	if math.Abs(float64(result-expected)) > 0.00001 {
		t.Errorf("Expected %.2f but go %.2f", expected, result)
		t.Fail()
	}
	result = decodeFloat18(97193)
	expected = 442.5625
	if math.Abs(float64(result-expected)) > 0.00001 {
		t.Errorf("Expected %.2f but go %.2f", expected, result)
		t.Fail()
	}
}

func TestEncodeFloat18(t *testing.T) {
	if encodeFloat18(-724.99) != 231079 {
		t.Fail()
	}
	if encodeFloat18(846.53) != 100980 {
		t.Fail()
	}
	if encodeFloat18(442.59) != 97193 {
		t.Fail()
	}
}

func TestReadPacket(t *testing.T) {
	var buf = []byte{178, 157, 26, 78, 167, 88, 234, 94}
	var packet = ReadPacket(buf)
	var expected = CommPacket{
		PacketType: State,
		PacketId:   54,
		Data1:      -724.875,
		Data2:      846.5,
		Data3:      442.5625,
	}
	if packet.PacketType != expected.PacketType {
		t.Errorf("Expected packet type [%d] but got [%d]", expected.PacketType, packet.PacketType)
		t.Fail()
	}
	if packet.PacketId != expected.PacketId {
		t.Errorf("Expected packet name [%s] but got [%s]", expected.PacketId, packet.PacketId)
		t.Fail()
	}
	if math.Abs(float64(packet.Data1-expected.Data1)) > 0.00001 {
		t.Errorf("Expected Data1 [%.2f] but got [%.2f]", expected.Data1, packet.Data1)
		t.Fail()
	}
	if math.Abs(float64(packet.Data2-expected.Data2)) > 0.00001 {
		t.Errorf("Expected Data2 [%.2f] but got [%.2f]", expected.Data2, packet.Data2)
		t.Fail()
	}
	if math.Abs(float64(packet.Data3-expected.Data3)) > 0.00001 {
		t.Errorf("Expected Data3 [%.2f] but got [%.2f]", expected.Data3, packet.Data3)
		t.Fail()
	}
}

func TestWritePacket(t *testing.T) {
	var buf = []byte{181, 199, 212, 174, 57, 109, 167, 155}
	var packet = ReadPacket(buf)
	var nbuf = WritePacket(packet)
	for i := 0; i < 8; i++ {
		if nbuf[i] != buf[i] {
			t.Errorf("Expected %d but got %d at %d", buf[i], nbuf[i], i)
			t.Fail()
		}
	}
}
