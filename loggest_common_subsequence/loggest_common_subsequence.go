package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// getLCS は、再調共通部分列を取得する。
func getLCS(x, y string) int{
	m := len(x)
	n := len(y)

	c := make([][]int, m+1)

	maxLength := 0
	x = addSpaceToHead(x)
	y = addSpaceToHead(y)

	for i := 0; i <= m; i++ {
		c[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++{
			if x[i] == y[j] {
				c[i][j] = c[i-1][j-1] +1
			} else {
				c[i][j] = max(c[i-1][j], c[i][j-1])
			}
			maxLength = max(maxLength, c[i][j])
		}
	}
	return maxLength
}

func addSpaceToHead(s string)string {
	return fmt.Sprintf(" %s", s)
}

func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}

var sc = bufio.NewScanner(os.Stdin)

func scanToInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func scanToText() string{
	sc.Scan()
	return sc.Text()
}


func main() {
	n := scanToInt()
	for i:= 0; i < n; i++ {
		s1, s2 := scanToText(), scanToText()
		ans := getLCS(s1, s2)
		fmt.Println(ans)
	}
}
