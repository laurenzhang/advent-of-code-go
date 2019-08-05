package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
	frequency := 0

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(bufio.NewReader(file))
	scanner.Split(bufio.ScanLines)

    for scanner.Scan() {
        line := scanner.Text()

        frequencyChange, err := strconv.Atoi(line[1:])
        if err != nil {
			panic(err)
		}
        
        if line[0] == '+' {
        	frequency += frequencyChange
        } else if line[0] == '-' {
        	frequency -= frequencyChange
        }
    }

    fmt.Println("Frequency:", frequency)
}
