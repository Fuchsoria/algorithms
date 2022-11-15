package bits

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type Case struct {
	Name   string
	Input  int
	Count  int
	Output uint64
}

type Cases = []Case

var (
	casesKing  Cases = ReadCases("./cases/king")
	casesHorse Cases = ReadCases("./cases/horse")
)

func ReadCases(path string) Cases {
	var index int
	var cases Cases

	for {
		fileInput, err := os.Open(filepath.Join(path, fmt.Sprintf("test.%d.in", index)))
		if err != nil {
			break
		}
		inputReader := bufio.NewReader(fileInput)
		inputLine, _, err := inputReader.ReadLine()
		if err != nil {
			break
		}
		fileOutput, err := os.Open(filepath.Join(path, fmt.Sprintf("test.%d.out", index)))
		if err != nil {
			break
		}
		outputReader := bufio.NewReader(fileOutput)
		countLine, _, err := outputReader.ReadLine()
		if err != nil {
			break
		}
		outputLine, _, err := outputReader.ReadLine()
		if err != nil {
			break
		}

		input, err := strconv.Atoi(string(inputLine))
		if err != nil {
			break
		}

		count, err := strconv.Atoi(string(countLine))
		if err != nil {
			break
		}

		output, err := strconv.ParseUint(string(outputLine), 0, 64)
		if err != nil {
			break
		}

		cases = append(cases, Case{fmt.Sprintf("Position %d", input), input, count, output})

		index++
	}

	return cases
}
