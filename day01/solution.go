package main

import (
	"fmt"
	"io"
	"maps"
	"slices"
	"strconv"
	"strings"

	"github.com/IEP/aoc-2024/utils"
)

func Solve01(input string) (string, error) {
	lines, err := utils.NewLineReader(input)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	s1 := make([]int, 0, 100)
	s2 := make([]int, 0, 100)

	for {
		line, err := lines.Line()
		if err == io.EOF {
			break
		}
		pairs := strings.Split(line, "   ")
		left_, right_ := pairs[0], pairs[1]
		left, _ := strconv.Atoi(left_)
		right, _ := strconv.Atoi(right_)
		s1 = append(s1, left)
		s2 = append(s2, right)
	}
	slices.Sort(s1)
	slices.Sort(s2)

	n := len(s1)
	totalDistance := 0
	for i := 0; i < n; i++ {
		distance := s1[i] - s2[i]
		if distance < 0 {
			distance = -1 * distance
		}
		totalDistance += distance
	}

	return fmt.Sprint(totalDistance), nil
}

func Solve02(input string) (string, error) {
	lines, err := utils.NewLineReader(input)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	leftMap := make(map[int]int)
	rightMap := make(map[int]int)
	checkMap := make(map[int]bool)

	for {
		line, err := lines.Line()
		if err == io.EOF {
			break
		}
		pairs := strings.Split(line, "   ")
		left_, right_ := pairs[0], pairs[1]
		left, _ := strconv.Atoi(left_)
		right, _ := strconv.Atoi(right_)

		leftMap[left] += 1
		rightMap[right] += 1
		checkMap[left] = true
		checkMap[right] = true
	}

	totalScore := 0
	for num := range maps.Keys(checkMap) {
		leftFreq := leftMap[num]
		rightFreq := rightMap[num]
		totalScore += num * leftFreq * rightFreq
	}

	return fmt.Sprint(totalScore), nil
}
