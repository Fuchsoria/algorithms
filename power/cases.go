package exponentiation

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Case struct {
	name   string
	input  float64
	inputN float64
	output float64
}

func parseInt(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}

	r := bufio.NewReader(file)
	line, _, err := r.ReadLine()
	if err != nil {
		return 0, err
	}

	num, err := strconv.Atoi(string(line))

	return num, nil
}

func parseFloat(path string) (float64, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}

	r := bufio.NewReader(file)
	line, _, err := r.ReadLine()
	if err != nil {
		return 0, err
	}

	num, err := strconv.ParseFloat(string(line), 64)

	return num, nil
}

func parseTwoFloats(path string) (float64, float64, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, 0, err
	}

	r := bufio.NewReader(file)
	line, _, err := r.ReadLine()
	if err != nil {
		return 0, 0, err
	}

	lineTwo, _, err := r.ReadLine()
	if err != nil {
		return 0, 0, err
	}

	num, err := strconv.ParseFloat(string(line), 64)
	numTwo, err := strconv.ParseFloat(string(lineTwo), 64)

	return num, numTwo, nil
}

func ParseCases() []Case {
	cases := []Case{}

	i := 0
	for {
		inputPath := fmt.Sprintf("./cases/test.%d.in", i)
		outputPath := fmt.Sprintf("./cases/test.%d.out", i)

		inputNum, inputNumTwo, err := parseTwoFloats(inputPath)
		if err != nil {
			break
		}

		outputNum, err := parseFloat(outputPath)
		if err != nil {
			break
		}

		cases = append(cases, Case{
			name:   fmt.Sprintf("test.%d", i),
			input:  inputNum,
			inputN: inputNumTwo,
			output: outputNum,
		})

		i++
	}

	return cases
}
