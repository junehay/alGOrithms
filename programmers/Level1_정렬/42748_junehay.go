/*
	Date : 2020-11-05
*/
package main

import (
	"fmt"
	"sort"
)

func solution(array []int, commands [][]int) []int {
    answer := make([]int,len(commands))
	for i, value := range commands {
        c := make([]int, len(array))
	    copy(c, array)
        start := value[0]
        end := value[1]
        target := value[2]
        cut := c[start-1:end]
        sort.Ints(cut)
        answer[i] = cut[target-1]
	}
    return answer
}

func main() {
	array := []int{1, 5, 2, 6, 3, 7, 4}
	commands := [][]int{{2, 5, 3}, {4, 4, 1}, {1, 7, 3}}
	fmt.Println(solution(array, commands)) // answer : 	[5, 6, 3]
}