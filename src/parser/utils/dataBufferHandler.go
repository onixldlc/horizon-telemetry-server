package utils

import "encoding/binary"

const INT_32_SIZE = 4
const UINT_32_SIZE = 4
const UINT_16_SIZE = 2
const FLOAT_SIZE = 4

type DataBuffer struct {
	data []byte
	pos  int
	len  int
}

func NewDataBuffer(data []byte) *DataBuffer {
	return &DataBuffer{
		data: data,
		pos:  0,
		len:  len(data),
	}
}

// func (buff *DataBuffer) add(data byte) {
// 	buff.data = append(buff.data, data)
// }

func (buff *DataBuffer) save() {
	buff.data = buff.data[buff.pos:]
}

func (buff *DataBuffer) read(length int) []byte {
	data := buff.data[buff.pos : buff.pos+length]
	buff.pos += length
	return data
}

func (buff *DataBuffer) dump() []byte {
	data := buff.read(len(buff.data))
	return data
}

func (buff *DataBuffer) getSignedInt() int32 {
	data := buff.read(INT_32_SIZE)
	return int32(binary.BigEndian.Uint32(data))
}

func (buff *DataBuffer) getUnsignedInt() uint32 {
	data := buff.read(UINT_32_SIZE)
	return uint32(binary.BigEndian.Uint32(data))
}

func (buff *DataBuffer) getSignedShort() int16 {
	data := buff.read(UINT_16_SIZE)
	return int16(binary.BigEndian.Uint16(data))
}

func (buff *DataBuffer) getUnsignedShort() uint16 {
	data := buff.read(UINT_16_SIZE)
	return uint16(binary.BigEndian.Uint16(data))
}

func (buff *DataBuffer) getFloat() float32 {
	data := buff.read(UINT_16_SIZE)
	return float32(binary.BigEndian.Uint32(data))
}

func (buff *DataBuffer) Parse(format string) []interface{} {
	result := make([]interface{}, 0)
	for letter := range format {
		switch letter {
		case 'i':
			result = append(result, buff.getSignedInt())
		case 'I':
			result = append(result, buff.getUnsignedInt())
		case 'h':
			result = append(result, buff.getSignedShort())
		case 'H':
			result = append(result, buff.getUnsignedShort())
		case 'f':
			result = append(result, buff.getFloat())
		}
	}
	return result
}
