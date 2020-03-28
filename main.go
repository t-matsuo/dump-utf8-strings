package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: " + os.Args[0] + " textfile")
		os.Exit(0)
	}

	fp, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("cannot open file :", os.Args[1])
		os.Exit(0)
	}
	defer fp.Close()

	buf1 := make([]byte, 1)
	buf2 := make([]byte, 2)
	buf3 := make([]byte, 3)
	buf4 := make([]byte, 4)

	var multi_byte2 int = 0
	var multi_byte3 int = 0
	var multi_byte4 int = 0
	var isDecodeError int = 0

	for {
		n, err := fp.Read(buf1)

		if n == 0 {
			break
		}
		if err != nil {
			break
		}

		fmt.Printf("%x", buf1[:n])
		fmt.Print(" ")

		// handle first byte
		if multi_byte2 == 0 && multi_byte3 == 0 && multi_byte4 == 0 {
			switch buf1[0] {
			case byte(0x00):
				fmt.Println("[NULL] (^@)")
				continue
			case byte(0x08):
				fmt.Println("[BS] (^H)")
				continue
			case byte(0x0a):
				fmt.Println("[LF] (^J)")
				continue
			case byte(0x0d):
				fmt.Println("[CR] (^M)")
				continue
			case byte(0x1b):
				fmt.Println("[ESC] (^[)")
				continue
			case byte(0x20):
				fmt.Println("[SPACE]")
				continue
			default:
			}

			switch {
			case is2bytesChar(buf1[0]):
				copy(buf2, buf1)
				multi_byte2 = 1
				fmt.Println("[2>]")
				continue
			case is3bytesChar(buf1[0]):
				copy(buf3, buf1)
				multi_byte3 = 1
				fmt.Println("[3>]")
				if buf1[0] == byte(0xef) {
					isDecodeError = 1
				}
				continue
			case is4bytesChar(buf1[0]):
				copy(buf4, buf1)
				multi_byte4 = 1
				fmt.Println("[4>]")
				continue
			default:
			}

			fmt.Println(string(buf1[:n]))
			continue
		}

		// handle multibytes chars
		if multi_byte2 != 0 {
			buf4[multi_byte2] = buf1[0]
			multi_byte2++
			if multi_byte2 == 2 {
				fmt.Println(string(buf2))
				multi_byte2 = 0
				continue
			}
			fmt.Println("[->]")
			continue
		}
		if multi_byte3 != 0 {
			buf3[multi_byte3] = buf1[0]
			multi_byte3++
			if isDecodeError == 1 {
				if buf1[0] == byte(0xbf) {
					isDecodeError++
				} else {
					isDecodeError = 0
				}
			}
			if isDecodeError == 2 {
				if buf1[0] == byte(0xbd) {
					fmt.Println("[Decode Error]")
					multi_byte3 = 0
					isDecodeError = 0
					continue
				}
			}
			if multi_byte3 == 3 {
				fmt.Println(string(buf3))
				multi_byte3 = 0
				continue
			}
			fmt.Println("[->]")
			continue
		}
		if multi_byte4 != 0 {
			buf4[multi_byte4] = buf1[0]
			multi_byte4++
			if multi_byte4 == 4 {
				fmt.Println(string(buf4))
				multi_byte4 = 0
				continue
			}
			fmt.Println("[->]")
			continue
		}
	}
	fp.Close()
}

// 2bytes char begins with \cx \dx
func is2bytesChar(b byte) bool {
	var mask byte = 0xF0 // =11110000
	var i = uint8(b & mask)
	// get upper 4bit
	i = i >> 4
	if i == 0x0c || i == 0x0d {
		return true
	}
	return false
}

// 3bytes char begins with \xeX
func is3bytesChar(b byte) bool {
	var mask byte = 0xF0 // =11110000
	var i = uint8(b & mask)
	// get upper 4bit
	i = i >> 4
	if i == 0x0e {
		return true
	}
	return false
}

// 4bytes char begins with \xfX
func is4bytesChar(b byte) bool {
	var mask byte = 0xF0 // =11110000
	var i = uint8(b & mask)
	// get upper 4bit
	i = i >> 4
	if i == 0x0f {
		return true
	}
	return false
}
