package tree

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readFile(fileName string) (arr []int) {
	arr = make([]int, 0)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		num, err := strconv.Atoi(line)

		if nil == err {
			arr = append(arr, num)
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}

	return arr
}

var (
	dataArr1kk        []int = readFile("cases/1000000n-data.txt")
	dataArr100k       []int = dataArr1kk[:100_000]
	dataArr10k        []int = dataArr1kk[:10_000]
	dataArr1k         []int = dataArr1kk[:1_000]
	dataArrSorted1kk  []int = readFile("cases/1000000n-sorted-data.txt")
	dataArrSorted100k []int = dataArrSorted1kk[:100_000]
	dataArrSorted10k  []int = dataArrSorted1kk[:10_000]
	dataArrSorted1k   []int = dataArrSorted1kk[:1_000]
)
