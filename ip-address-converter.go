
package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	// rgb(0,0,0) == 0x000000 ,

	// rgb(255,255,255) == FFFFFF

	//address := "181.61.208.136"
	address := "50.35.65.141"
	parts := strings.Split(address, ".")
	if len(parts) != 4 {
		log.Println("INVALID IP ADDRESS:" + address)
		return
	}
	segments := []int{0, 0, 0, 0}
	for i := 0; i < 4; i++ {
		segment, err := strconv.Atoi(parts[i])

		if err != nil {
			log.Println("Invalid segment:" + parts[i] + " on client ipaddress:" + address)
		}
		segments[i] = segment
	}
	decimalAddress := segments[3] + segments[2]*256 + segments[1]*256*256 + segments[0]*256*256*256
	fmt.Println(decimalAddress)

}
