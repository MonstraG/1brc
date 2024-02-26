package main

import (
	"1brc/src"
	"fmt"
)

func main() {
	buffer := src.ProcessFile("weather_stations.csv")
	if buffer != nil {
		fmt.Println(buffer.String())
	}
}
