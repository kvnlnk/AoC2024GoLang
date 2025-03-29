package main

import (
	"fmt"
	"github.com/kvnlnk/AoC2024GoLang/Common"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data := Common.ReadInput("Day01/Input.txt")
	firstList, secondList := splitInput(data)
	sort.Ints(firstList)
	sort.Ints(secondList)
	fmt.Println(solvePartOne(firstList, secondList))
	fmt.Println(solvePartTwo(firstList, secondList))
}

func splitInput(data []string) ([]int, []int) {
	var firstList []int
	var secondList []int
	for i := 0; i < len(data); i++ {

		leftValue, err := strconv.Atoi(strings.Fields(data[i])[0])
		rightValue, err := strconv.Atoi(strings.Fields(data[i])[1])

		if err != nil {
			log.Fatalf("Error while converting from string to int")
		}
		firstList = append(firstList, leftValue)
		secondList = append(secondList, rightValue)
	}
	return firstList, secondList
}

func solvePartOne(firstList []int, secondList []int) string {
	sum := 0

	for i := 0; i < len(firstList); i++ {
		sum += int(math.Abs(float64(firstList[i] - secondList[i])))
	}

	return strconv.Itoa(sum)
}

func solvePartTwo(firstList []int, secondList []int) string {
	sum := 0
	countAppearancesMap := make(map[int]int)

	for i := 0; i < len(secondList); i++ {
		countAppearancesMap[secondList[i]]++
	}

	for i := 0; i < len(firstList); i++ {
		sum += firstList[i] * countAppearancesMap[firstList[i]]
	}
	return strconv.Itoa(sum)
}
