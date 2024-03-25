package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// import "io"
// import "bufio"
func main() {
	currentLine := 0
	fmt.Println("Hello, World!")

	dat, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("ERR")
	}
	lines := strings.Fields(string(dat))

	BluePrint := 0
	for k := 0; k < len(lines); k++ {

		for i := 0; i < len(lines[k]); i++ {
			numfound := false
			num := ""
			numLength := 0
			for lines[k][i] >= '0' && lines[k][i] <= '9' {
				num += string(lines[k][i])
				numfound = true
				numLength += 1
				i++
				if i == len(lines[k]) {
					break
				}
			}
			if numfound {
				numValid := false
				for j := 0; j < numLength+2; j++ {
					if k-1 >= 0 && (i-j) >= 0 && (i-j) < len(lines[k-1]) && lines[k-1][i-j] != '.' && !(lines[k-1][i-j] >= '0' && lines[k-1][i-j] <= '9') {
						numValid = true
						break
					}
				}
				if !numValid {
					for j := 0; j < numLength+2; j++ {
						if k+1 < len(lines) && (i-j) >= 0 && (i-j) < len(lines[k+1]) && lines[k+1][i-j] != '.' && !(lines[k+1][i-j] >= '0' && lines[k+1][i-j] <= '9') {
							numValid = true
							break
						}
					}
					if (i+1) < len(lines[k]) && lines[k][i] != '.' && !(lines[k][i] >= '0' && lines[k][i] <= '9') {
						numValid = true
					}
					if (i-numLength-1) >= 0 && lines[k][i-numLength-1] != '.' && !(lines[k][i-numLength-1] >= '0' && lines[k][i-numLength-1] <= '9') {
						numValid = true
					}
				}
				if numValid {
					numI, err := strconv.Atoi(num)
					if err != nil {
					}
					if currentLine != k {
						currentLine = k
						fmt.Print("\n")
					}
					fmt.Print(numI, " ")
					BluePrint += numI

				}
			}

		}
	}
	fmt.Print("\n", BluePrint)
	//fmt.Print("\n", dat[19876-2] - '0')
	//fmt.Print("\n", dat[19876-142-0] - '.')

}
