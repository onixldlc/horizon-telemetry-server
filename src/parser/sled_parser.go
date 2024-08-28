package parser

import (
	"fmt"
	"forza_telemetry/parser/utils"
	"strings"
)

type SLED struct {
	buffer *utils.DataBuffer
}

func NewDataSled(data []byte) *SLED {
	return &SLED{
		buffer: utils.NewDataBuffer(data),
	}
}

func (sled SLED) ParseSledData() string {
	jsonListConverted := make([]string, 0)
	for key, value := range utils.SledDictionary {
		unpacked := sled.buffer.Parse(value)[0]
		jsonListConverted = append(jsonListConverted, fmt.Sprintf("\"%s\":\"%s\"", key, unpacked))
	}
	jsonListConverted = append(jsonListConverted, "\"packet_type\":\"sled\"")
	jsonConverted := "{" + strings.Join(jsonListConverted, ",") + "}"
	return jsonConverted
}
