package parser

import (
	"fmt"
	"forza_telemetry/parser/utils"
	"strings"
)

type CARDASH struct {
	buffer *utils.DataBuffer
}

func NewDataCarDash(data []byte) *CARDASH {
	return &CARDASH{
		buffer: utils.NewDataBuffer(data),
	}
}

func (cardash CARDASH) ParseCarDashData() string {
	jsonListConverted := make([]string, 0)
	for key, value := range utils.CarDashDictionary {
		unpacked := cardash.buffer.Parse(value)[0]
		jsonListConverted = append(jsonListConverted, fmt.Sprintf("\"%s\":\"%s\"", key, unpacked))
	}
	jsonListConverted = append(jsonListConverted, "\"packet_type\":\"car_dash\"")
	jsonConverted := "{" + strings.Join(jsonListConverted, ",") + "}"
	return jsonConverted
}
