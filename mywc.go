package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func wc_l(filePath string) int64 {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(strings.NewReader(string(data)))
	scanner := bufio.NewScanner(reader)
	var cmp int = 0
	for scanner.Scan() {
		cmp++
	}
	return int64(cmp)
}
func wc_c(filePath string) int64 {

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	count := int64(len(data))
	return int64(count)
}
func wc_w(filePath string) int64 {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	Scanner := bufio.NewScanner(file)
	Scanner.Split(bufio.ScanWords)
	var cmp int = 0
	for Scanner.Scan() {
		cmp++
	}
	return int64(cmp)
}
func main() {
	if len(os.Args) > 2 {
		filePath := os.Args[2]
		fileParts := strings.Split(filePath, "/")
		fileName := fileParts[len(fileParts)-1]
		param := os.Args[1]
		switch param {
		case "-l":
			fmt.Printf("%d %s", wc_l(filePath), fileName)

			break
		case "-c":
			{
				fmt.Printf("%d %s", wc_c(filePath), fileName)
				break
			}
		case "-w":
			{
				fmt.Printf("%d %s", wc_w(filePath), fileName)

				break
			}

		}

	} else {
		filePath := os.Args[1]
		fileParts := strings.Split(filePath, "/")
		fileName := fileParts[len(fileParts)-1]
		fmt.Printf("%d %d %d %s \n", wc_l(filePath), wc_w(filePath), wc_c(filePath), fileName)

		return
	}

}
