package main

import (
	"fmt"
	"strings"
)

var NATO = map[string]string{
	"A": "Alfa",
	"B": "Bravo",
	"C": "Charlie",
	"D": "Delta",
	"E": "Echo",
	"F": "Foxtrot",
	"G": "Golf",
	"H": "Hotel",
	"I": "India",
	"J": "Juliet",
	"K": "Kilo",
	"L": "Lima",
	"M": "Mike",
	"N": "November",
	"O": "Oscar",
	"P": "Papa",
	"Q": "Quebec",
	"R": "Romeo",
	"S": "Sierra",
	"T": "Tango",
	"U": "Uniform",
	"V": "Victor",
	"W": "Whiskey",
	"X": "Xray",
	"Y": "Yankee",
	"Z": "Zulu",
}

// <6 kyu> kata - if you can read this...
// https://www.codewars.com/kata/586538146b56991861000293/train/go
func ToNato(words string) string {
	letters := strings.Split(words, "")
	parts := []string{}

	for _, letter := range letters {
		if nato, ok := NATO[strings.ToUpper(letter)]; ok {
			parts = append(parts, nato)
		} else {
			if letter != " " {
				parts = append(parts, letter)
			}
		}
	}
	return strings.Join(parts, " ")
}

// <5 kyu> moving zeros to the end:
// https://www.codewars.com/kata/52597aa56021e91c93000cb0/go
// create an array of ints with the same length as the input, the default value will be zero
// only copy the non-zero numbers over
func MoveZeros(arr []int) []int {
	retval := make([]int, len(arr))
	idx := 0

	for _, v := range arr {
		if v != 0 {
			retval[idx] = v
			idx++
		}
	}
	return retval
}

// <5 kyu> Hamster me
// https://www.codewars.com/kata/595ddfe2fc339d8a7d000089/train/go
func HamsterMe(code, message string) string {
	toEncode := []rune(code)
	depth := 1

	encoder := map[rune]string{}
	for len(toEncode) > 0 {
		nextLevel := []rune{}
		for _, c := range toEncode {
			if _, ok := encoder[c]; !ok {
				codeval := c
				if depth > 1 {
					codeval = rune((((int(c) - depth + 27) - 'a') % 26) + 'a')
				}
				encoder[c] = fmt.Sprintf("%c%d", codeval, depth)
				nextLevel = append(nextLevel, (((c+1)-'a')%26)+'a')
			}
		}
		depth += 1
		toEncode = nextLevel
	}

	encodedMsg := ""
	for _, c := range message {
		encodedMsg += encoder[c]
	}
	return encodedMsg
}

// kata - human readable duration format
// https://www.codewars.com/kata/52742f58faf5485cae000b9a/train/go
func FormatDurationClever(seconds int64) string {
	// start with the smallest, work your way up, add separators as you go:
	timeList := []int64{60, 3600, 86400, 31536000}
	unitSize := []int64{1, 60, 3600, 86400, 31536000}
	unitNames := []string{"second", "minute", "hour", "day", "year"}

	output := ""
	usedAnd := false

	for i, v := range timeList {
		remainder := seconds % v
		seconds -= remainder
		units := remainder / unitSize[i]

		if remainder != 0 {
			modifier := ""
			if units > 1 {
				modifier = "s"
			}
			separator := ""
			if seconds > 0 {
				if !usedAnd {
					separator = " and "
					usedAnd = true
				} else {
					separator = ", "
				}
			}
			output = fmt.Sprintf("%s%d %s%s%s", separator, units, unitNames[i], modifier, output)
		}
		if seconds <= 0 {
			break
		}
	}
	return output
}

// <3 kyu> twice linear
// https://www.codewars.com/kata/5672682212c8ecf83e000050/train/go
// This solution is OK, but the clever ones are better. Next time.
func DblLinear(n int) int {
	sortedInsertFcn := func(val int, sorted []int) []int {
		idx := len(sorted) - 1
		if val > sorted[idx] {
			return append(sorted, val)
		}

		for {
			if val == sorted[idx] {
				return sorted
			}
			if val < sorted[idx] {
				idx -= 1
			} else {
				idx += 1
				sorted = append(sorted, 0)
				copy(sorted[idx+1:], sorted[idx:])
				sorted[idx] = val
				return sorted
			}
		}
	}

	index := 0
	sorted := []int{1}

	for {
		v := sorted[index]
		index += 1
		n1 := (v * 2) + 1
		n2 := (v * 3) + 1

		sorted = sortedInsertFcn(n1, sorted)
		sorted = sortedInsertFcn(n2, sorted)

		if len(sorted) > n && n1 >= sorted[n] {
			return sorted[n]
		}
	}
}
