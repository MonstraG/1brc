package src

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func TestChallenge(t *testing.T) {
	files, err := os.ReadDir("./testcases")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".out") {
			continue
		}

		ValidateTestCase(t, "./testcases/"+filename)
	}
}

func ValidateTestCase(t *testing.T, file string) {
	inputFile := file
	outputFile := strings.Replace(file, ".txt", ".out", 1)

	result := ProcessFile(inputFile)

	output, err := os.ReadFile(outputFile)
	if err != nil {
		log.Fatal("Couldn't read inputFile, ", err)
	}

	got := result.String()
	want := string(output)
	if got != want {
		errorMsg := fmt.Sprintf(" failed, got=%v, want=%v", got, want)
		t.Error(file, errorMsg)
	}
}
