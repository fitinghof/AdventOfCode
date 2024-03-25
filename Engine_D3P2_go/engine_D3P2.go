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
	//currentLine := 0
	fmt.Println("Hello, World!")

	dat, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("ERR")
	}
	lines := strings.Fields(string(dat))

	BluePrint := 0
	//loop through lines
	for line := 0; line < len(lines); line++ {
		//loop through ch
		for ch := 0; ch < len(lines[line]); ch++ {
			if lines[line][ch] == '*' {
				num1 := 0
				num1Pos := 0
				num1Init := false
				num2 := 0
				pairFound := false
				for row := 0; row < 3; row++ {
					for col := 0; col < 3; col++ {
						if line+row-1 < len(lines) && line+row-1 >= 0 && ch+col-1 <= len(lines[line]) && ch+col-1 >= 0 && lines[line+row-1][ch+col-1] <= '9' && lines[line+row-1][ch+col-1] >= '0' {
							relPos := 0
							for j := 0; lines[line+row-1][ch+col-1+j] <= '9' && lines[line+row-1][ch+col-1+j] >= '0'; j-- {
								relPos = j
								if ch+col-1+j == 0 {
									break
								}
							}
							temp := ""
							tempPos := 0
							for j := 0; ch+col-1+relPos+j < len(lines[line]) && lines[line+row-1][ch+col-1+relPos+j] <= '9' && lines[line+row-1][ch+col-1+relPos+j] >= '0'; j++ {
								temp += string(lines[line+row-1][ch+col-1+relPos+j])
								tempPos = (line+row-1)*140 + (ch + col - 1 + relPos + j)
							}
							if !num1Init {
								num1, _ = strconv.Atoi(temp)
								fmt.Print("num1:", num1, "\n")
								num1Pos = tempPos
								num1Init = true
							} else if tempPos != num1Pos {
								num2, _ = strconv.Atoi(temp)
								fmt.Print("num2:", num2, "\n")
								BluePrint += num1 * num2
								pairFound = true
								break
							}

						}
					}
					if pairFound {
						break
					}
				}
			}


		}
	}
	fmt.Print("\n", BluePrint)
}
