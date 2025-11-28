package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"
)

type Solution struct {
	Year       string
	Day        string
	DayNum     string
	ImportPath string
	ImportName string
}

type SolutionWithKey struct {
	Solution
	Key string
}

type TemplateData struct {
	Solutions       []SolutionWithKey
	SolutionsByYear map[string][]Solution
}

func main() {
	solutions := []Solution{}
	yearPattern := regexp.MustCompile(`^20\d{2}$`)
	dayPattern := regexp.MustCompile(`^day(\d{2})$`)

	entries, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		yearName := entry.Name()
		if !yearPattern.MatchString(yearName) {
			continue
		}

		dayEntries, err := os.ReadDir(yearName)
		if err != nil {
			continue
		}

		for _, dayEntry := range dayEntries {
			if !dayEntry.IsDir() {
				continue
			}

			dayName := dayEntry.Name()
			matches := dayPattern.FindStringSubmatch(dayName)
			if matches == nil {
				continue
			}

			dayNum := matches[1]
			solutionPath := filepath.Join(yearName, dayName, "solution.go")

			if _, err := os.Stat(solutionPath); err == nil {
				solutions = append(solutions, Solution{
					Year:       yearName,
					Day:        dayName,
					DayNum:     dayNum,
					ImportPath: fmt.Sprintf("janisvepris/aoc25/%s/%s", yearName, dayName),
					ImportName: fmt.Sprintf("%s_%s", dayName, yearName),
				})
			}
		}
	}

	sort.Slice(solutions, func(i, j int) bool {
		if solutions[i].Year != solutions[j].Year {
			return solutions[i].Year < solutions[j].Year
		}
		return solutions[i].Day < solutions[j].Day
	})

	// Build template data
	var solutionsWithKeys []SolutionWithKey
	solutionsByYear := make(map[string][]Solution)

	for _, sol := range solutions {
		key := fmt.Sprintf("%s-%s", sol.Year, sol.DayNum)
		solutionsWithKeys = append(solutionsWithKeys, SolutionWithKey{
			Solution: sol,
			Key:      key,
		})

		// Group by year
		solutionsByYear[sol.Year] = append(solutionsByYear[sol.Year], sol)
	}

	data := TemplateData{
		Solutions:       solutionsWithKeys,
		SolutionsByYear: solutionsByYear,
	}

	// Parse and execute template
	tmpl, err := template.ParseFiles("tools/registry.tmpl")
	if err != nil {
		panic(err)
	}

	var output strings.Builder
	err = tmpl.Execute(&output, data)
	if err != nil {
		panic(err)
	}

	// Write result
	err = os.WriteFile("registry.go", []byte(output.String()), 0o644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Generated registry with %d solutions\n", len(solutionsWithKeys))
}
