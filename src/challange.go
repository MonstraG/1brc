package src

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ProcessFile(filePath string) *bytes.Buffer {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Failed to open file ", err)
	}

	stationMap := make(map[string][]float64)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		addOrUpdateStation(&stationMap, text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Failed to scan file ", err)
		return nil
	}

	buffer := bytes.Buffer{}
	buffer.WriteRune('{')

	totalStations := len(stationMap)
	commaCounter := 0

	for index, stationTemps := range stationMap {
		if commaCounter != 0 && commaCounter != totalStations {
			buffer.WriteString(", ")
			commaCounter += 1
		}
		minTemp := slices.Min(stationTemps)
		meanTemp := mean(&stationTemps)
		maxTemp := slices.Max(stationTemps)
		line := fmt.Sprintf("%v=%v/%v/%v", index, formatFloat(minTemp), formatFloat(meanTemp), formatFloat(maxTemp))
		buffer.WriteString(line)
	}

	buffer.WriteRune('}')
	buffer.WriteRune('\n')
	return &buffer
}

func formatFloat(value float64) float64 {
	return math.Ceil(value*100) / 100
}

func mean(slice *[]float64) float64 {
	sliceVal := *slice
	if len(sliceVal) == 0 {
		return 0
	}
	var sum float64
	for _, d := range sliceVal {
		sum += d
	}
	return sum / float64(len(sliceVal))
}

func addOrUpdateStation(stationMap *map[string][]float64, line string) {
	parts := strings.Split(line, ";")
	stationName := parts[0]
	stationTempString := parts[1]
	stationTemp, err := strconv.ParseFloat(stationTempString, 64)
	if err != nil {
		log.Fatalf("Failed to parse temp. stationName=%v, temp=%v, error=%v", stationName, stationTemp, err)
		return
	}

	if (*stationMap)[stationName] == nil {
		(*stationMap)[stationName] = make([]float64, 0)
	}
	(*stationMap)[stationName] = append((*stationMap)[stationName], stationTemp)
}
