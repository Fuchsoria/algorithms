package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Case struct {
	name   string
	input  int
	output int
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

func ParseCases() []Case {
	cases := []Case{}

	i := 0
	for {
		inputPath := fmt.Sprintf("./cases/test.%d.in", i)
		outputPath := fmt.Sprintf("./cases/test.%d.out", i)

		inputNum, err := parseInt(inputPath)
		if err != nil {
			break
		}

		outputNum, err := parseInt(outputPath)
		if err != nil {
			break
		}

		cases = append(cases, Case{
			name:   fmt.Sprintf("test.%d", i),
			input:  inputNum,
			output: outputNum,
		})

		i++
	}

	return cases
}
