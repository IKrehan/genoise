package main

import (
	"genoise/pkg/noise"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatal("Missing Arguments")
	}

	widthInput, heightInput, fileName := os.Args[1], os.Args[2], os.Args[3]

	width, err := strconv.Atoi(widthInput)
	if err != nil {
		log.Fatalf("Value is not an integer: %v", widthInput)
	}

	height, err := strconv.Atoi(heightInput)
	if err != nil {
		log.Fatalf("Value is not an integer: %v", heightInput)
	}

	noise := noise.New(noise.Hash)
	err = noise.GenerateImage(fileName, width, height)
	if err != nil {
		log.Fatal(err.Error())
	}
}
