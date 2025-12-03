# Advent of Code solutions 2022-2025 in Go

### Run a specific day

`go run . -y=2024 -d=13`

### Download input for a specific day

Put the value of your session cookie in a `session` file in the root directory

run `go run . -y=2024 -d=13 -dl`

### Generate a template for a solution
`go run . -y=2024 -d13 -t`

### Generating solution registry

`go generate`

### Full preparation for a new solution

- Creates a solution.go template file in the right directory
- Downloads the input file to that directory (don't forget to put your session cookie value into `./session`)
- Refreshes the registry so you can immediately run your solution

```bash
# Generate and download all the things
go run . -y=2025 -d=4 -dl -t && go generate

# run it
go run . -y=2025 -d4
```
