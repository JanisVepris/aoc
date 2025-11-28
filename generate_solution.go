package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type SolutionTemplateData struct {
	Package string
	Year    string
	Day     string
}

func GenerateSolution(year, day string) {
	dirPath := fmt.Sprintf("%s/day%s", year, day)

	err := os.MkdirAll(dirPath, 0o755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	data := SolutionTemplateData{
		Package: "day" + day,
		Year:    year,
		Day:     "day" + day,
	}

	tmpl, err := template.ParseFiles("tools/solution.go.tmpl")
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	var content strings.Builder
	err = tmpl.Execute(&content, data)
	if err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		return
	}

	filePath := fmt.Sprintf("%s/solution.go", dirPath)

	if _, err := os.Stat(filePath); err == nil {
		fmt.Printf("Warning: %s already exists, skipping...\n", filePath)
		return
	}

	err = os.WriteFile(filePath, []byte(content.String()), 0o644)
	if err != nil {
		fmt.Printf("Error writing template file: %v\n", err)
		return
	}

	fmt.Printf("✓ Created template at %s\n", filePath)
}
