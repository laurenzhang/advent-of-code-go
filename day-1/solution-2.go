package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/deckarep/golang-set"
)

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
	frequency := 0
	frequencyChanges := make([]int, 0)
	frequenciesSeen := mapset.NewSet()


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
        	frequencyChanges = append(frequencyChanges, frequencyChange)
        } else if line[0] == '-' {
        	frequencyChanges = append(frequencyChanges, frequencyChange * -1)
        }
    }

    for {
    	for _, frequencyChange := range frequencyChanges {
	    	frequency += frequencyChange
	        
	        if frequenciesSeen.Contains(frequency) {
	        	fmt.Println("First repeated frequency:", frequency)
	        	return
	        }

	        frequenciesSeen.Add(frequency)
    	}
    }
}
