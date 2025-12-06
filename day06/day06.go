package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := [][]string{}
	clearLines := []string{}
	emptyColumns := []int{}
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("failed to open file: %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		clearLines = append(clearLines, line)
		lines = append(lines, strings.Fields(line))
	}
	re := regexp.MustCompile(`[*+]`)
	matches := re.FindAllStringIndex(clearLines[len(clearLines)-1], -1)
	for _, element := range matches {
		index := element[0] - 1
		if index >= 0 {
			emptyColumns = append(emptyColumns, index)
		}
	}
	part1 := int64(0)
	part2 := int64(0)
	for i := 0; i < len(lines[0]); i++ {
		sign := strings.TrimSpace(lines[len(lines)-1][i])
		temp1, _ := strconv.ParseInt(lines[0][i], 10, 64)
		tempArray := make([]string, len(lines))
		largest := len(lines[0][i])
		tempArray[0] = lines[0][i]
		for j := 1; j < len(lines)-1; j++ {
			tempArray[j] = lines[j][i]
			if len(lines[j][i]) > largest {
				largest = len(lines[j][i])
			}
			switch sign {
			case "+":
				x, _ := strconv.ParseInt(lines[j][i], 10, 64)
				temp1 += x
			case "*":
				x, _ := strconv.ParseInt(lines[j][i], 10, 64)
				temp1 *= x
			}
		}
		part1 += temp1
	}
	fmt.Println(part1)

	tempNumbers := []int64{}
	tempSign := '+'
	for i := len(clearLines[0]) - 1; i >= -1; i-- {
		if slices.Contains(emptyColumns, i) || i == -1 {
			switch tempSign {
			case '+':
				temp2 := tempNumbers[0]
				for num := 1; num < len(tempNumbers); num++ {
					temp2 += tempNumbers[num]
				}
				part2 += temp2
				tempNumbers = []int64{}
			case '*':
				temp2 := tempNumbers[0]
				for num := 1; num < len(tempNumbers); num++ {
					temp2 *= tempNumbers[num]
				}
				part2 += temp2
				tempNumbers = []int64{}
			default:
				fmt.Println("wrong sign ")
				fmt.Println(tempSign)
			}
		} else {
			r := []rune(clearLines[len(clearLines)-1])
			if i < len(r) {
				tempSign = r[i]
			}
			tempStr := ""
			for j := 0; j < len(clearLines)-1; j++ {
				r := []rune(clearLines[j])
				tempStr += string(r[i])
			}
			tempStr = strings.ReplaceAll(tempStr, " ", "")
			x, _ := strconv.ParseInt(tempStr, 10, 64)
			tempNumbers = append(tempNumbers, x)
		}
	}
	fmt.Println(part2)

}
