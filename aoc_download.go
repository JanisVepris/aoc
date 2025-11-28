package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func DownloadInput(year, day string) {
	session := ""

	data, err := os.ReadFile("session")
	if err != nil {
		panic(err)
	}
	session = strings.TrimSpace(string(data))

	dayNum := strings.TrimLeft(day, "0")
	if dayNum == "" {
		dayNum = "0"
	}
	url := "https://adventofcode.com/" + year + "/day/" + dayNum + "/input"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic("Failed to download input data: " + resp.Status)
	}

	inputData, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	filePath := fmt.Sprintf("%s/day%s/input.txt", year, day)
	dirPath := fmt.Sprintf("%s/day%s", year, day)

	err = os.MkdirAll(dirPath, 0o755)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filePath, inputData, 0o644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("✓ Downloaded input file at %s\n", filePath)
}
