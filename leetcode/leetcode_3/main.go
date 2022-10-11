package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("vim-go")
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

// func lenSubset(s string) int {

// }

func lengthOfLongestSubstring(s string) int {

	l := len(s)
	memo := make(map[int][]byte)
	max := 0
	for i := 0; i < len(s); i++ {
		result := lenReqTwo(s, "", l, i, memo)
		if max < result {
			max = result
		}
	}
	return max
}

func cropStr(s string, memo map[int][]byte) (string, string) {
	if len(s) == 0 {
		return "", ""
	}
	currlen := len(s)
	v, ok := memo[0]
	if !ok {
		bytes := []byte(s)
		memo[0] = bytes
		return string(bytes[1:]), string(bytes[0])
	}

	return string(v[:currlen-1]), string(v[len(v)-currlen])
}

func lenReqTwo(s, seq string, initialLength, offset int, memo map[int][]byte) int {
	if offset > len(s) {
		return 0
	}
	if len(s) == initialLength {
		s = strings.Join(strings.Split(s, "")[offset:], "")
	}

	if len(s) == 0 {
		return 0
	}

	cropped, last := cropStr(s, memo)
	if strings.Contains(seq, last) {
		return 0
	}
	return lenReqTwo(cropped, seq+last, initialLength, offset, memo) + 1
}
