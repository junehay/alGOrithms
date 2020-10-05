/*
	Date : 2020-10-06
*/

func solution(answers []int) []int {
    solve := make([]int, 0, 3)
    person1 := []int{1,2,3,4,5}
    person2 := []int{2,1,2,3,2,4,2,5}
    person3 := []int{3,3,1,1,2,2,4,4,5,5}
    correct := []int{0,0,0}
    
    for index, answer := range answers {
        if answer == person1[index % len(person1)] {
			correct[0]++
		}
		if answer == person2[index % len(person2)] {
			correct[1]++
		}
		if answer == person3[index % len(person3)] {
			correct[2]++
		}	
    }
    
    max := correct[0]
	for _, v := range correct {
        if (max < v) {
           max = v
        }
	}
    
    for i, v := range correct {
        if (v == max) {
           solve = append(solve, i+1)
        }
    }
    
    return solve
}