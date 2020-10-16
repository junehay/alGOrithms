/*
	Date : 2020-10-16
*/
package main

import "fmt"

func solution(n int, lost []int, reserve []int) int {
    for reserveIndex, r := range reserve {
        for lostIndex, l := range lost {
            if r == l {
                reserve[reserveIndex] = 0
                lost[lostIndex] = 0
                goto PASS
            }
        }
        PASS:
    }  
    
    for _, r := range reserve {
        for i, l := range lost {
            switch {
                case r-1 == l:
                    lost = remove(lost, i)
                    goto NEXT
                case r == l:
                    lost = remove(lost, i)
                    goto NEXT
                case r+1 == l:
                    lost = remove(lost, i)
                    goto NEXT
            }
        }
        NEXT:
    }
    return n-len(lost)
}
func remove(slice []int, s int) []int {
    return append(slice[:s], slice[s+1:]...)
}

func main() {
	n := 5
	lost := []int{2,4}
	reserve := []int{1,3,5} // answer : 5
	fmt.Println(solution(n, lost, reserve))
}