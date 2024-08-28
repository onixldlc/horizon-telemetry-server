package handler

import (
	"fmt"
)

func hexDump(data []byte) {
	length := 16

	index := 0
	for index < len(data) {
		temp_hexs := make([]byte, length)

		if len(data)-index < length {
			temp_hexs = make([]byte, len(data)%length)
			copy(temp_hexs, data[index:index+len(data)%length])
		} else {
			copy(temp_hexs, data[index:index+length])
		}
		dumperHex(temp_hexs, (length*3)+3)
		// padPrint(8, " ")
		dumper(temp_hexs, length)
		clearList(temp_hexs)
		fmt.Print("\n")
		index += length
	}

	// data_duplicate := make([]byte, len(data))
	// copy(data_duplicate, data)

	// temp_text := make([]byte, length)
	// temp_hexs := make([]byte, length)

	// pad_length := length - (len(data) % length)

	// for range len(data) / length {
	// 	// copy(data_duplicate[:length], temp_text)
	// 	copy(temp_hexs, data_duplicate[:length])
	// 	dumperHex(temp_hexs, length)
	// 	padPrint(8, " ")
	// 	dumper(temp_hexs, length)
	// 	clearList(temp_hexs)
	// }

	// for index, value := range data_duplicate {
	// 	if index%length == 0 && index != 0 {
	// 		padPrint(8, "-")
	// 		dumperHex(temp_hexs)
	// 		clearList(temp_hexs)
	// 		dumper(temp_text)
	// 		clearList(temp_text)
	// 		fmt.Printf("\n")
	// 	}

	// 	if index%(length/2) == 0 && index != 0 {
	// 		fmt.Print("  ")
	// 	}

	// 	temp_text[index%length] = value
	// 	temp_hexs[index%length] = value
	// 	// fmt.Printf("%02X ", value)
	// }

	// if pad_length != 0 {
	// fmt.Println("\n\n")
	// fmt.Println(pad_length)
	// fmt.Println(length)
	// fmt.Println(pad_length % length)
	// fmt.Println("\n\n")
	// padPrint(pad_length, "   ")
	// // fmt.Printf("%16s\n", temp_text)
	// padPrint(8, "-")
	// dumper(temp_text)
	// fmt.Printf("\n")
	// }

	fmt.Printf("\n")
}

func clearList(byte_list []byte) {
	for index := range byte_list {
		byte_list[index] = 0x00
	}
}

func dumperHex(byte_list []byte, length int) {
	index := 0
	byte_index := 0
	// fmt.Println(byte_list)
	for index <= length {
		if byte_index < len(byte_list) {
			fmt.Printf("%02X ", byte_list[byte_index])
			index += 3
			byte_index++
			continue
		}
		fmt.Print(" ")
		index++
	}
	// for index, hex := range byte_list {
	// 	if index%(length/2) == 0 && index != 0 {
	// 		fmt.Print(" ")
	// 	}
	// 	fmt.Printf("%02x ", hex)
	// }
}

func dumper(byte_list []byte, length int) {
	index := 0
	byte_index := 0
	// fmt.Println(byte_list)
	for index <= length {
		if byte_index < len(byte_list) {
			character := byte_list[byte_index]
			if (character > 32 && character <= 126) || (character >= 160) {
				fmt.Printf("%c", character)
			} else {
				fmt.Print(".")
			}
			index += 1
			byte_index++
			continue
		}
		fmt.Print(" ")
		index++
	}
	// for index, character := range byte_list {
	// 	if index%(length/2) == 0 && index != 0 {
	// 		fmt.Print("  ")
	// 	}
	// 	if (character > 32 && character <= 126) || (character >= 160) {
	// 		fmt.Printf("%c", character)
	// 		// fmt.Printf("%v ", character)
	// 	} else {
	// 		fmt.Print(".")
	// 	}
	// }
}
