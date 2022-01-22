package libs

import (
	"encoding/binary"
	// "math"
)

// Uint64ToBytes convert uint64 to bytes
func Uint64ToBytes(v uint64) (b []byte) {
	b = make([]byte, 8)
	binary.LittleEndian.PutUint64(b, v)
	return b
}

// Int64ToBytes convert int64 to bytes
func Int64ToBytes(v int64) (b []byte) {
	b = make([]byte, 8)
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
	b[6] = byte(v >> 48)
	b[7] = byte(v >> 56)
	return
}

// Uint32ToBytes convert uint32 to bytes
func Uint32ToBytes(v uint32) (b []byte) {
	b = make([]byte, 4)
	binary.LittleEndian.PutUint32(b, v)
	return
}

// Int32ToBytes convert int32 to bytes
func Int32ToBytes(v int32) (b []byte) {
	b = make([]byte, 4)
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	return
}

// Uint16ToBytes convert uint16 to bytes
func Uint16ToBytes(v uint16) (b []byte) {
	b = make([]byte, 2)
	binary.LittleEndian.PutUint16(b, v)
	return b
}

// Int16ToBytes convert int16 to bytes
func Int16ToBytes(v int16) (b []byte) {
	b = make([]byte, 2)
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	return
}

// BytesToInt64 bytes to int64
func BytesToInt64(b []byte) (v int64) {
	return int64(binary.LittleEndian.Uint64(b))
}

// BytesToUint64 bytes to uint64
func BytesToUint64(b []byte) (v uint64) {
	return binary.LittleEndian.Uint64(b)
}
