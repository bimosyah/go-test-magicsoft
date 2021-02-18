package sorting

import (
	"fmt"
	"go-magicsoft/utils"
)

func Soal3() {
	input := []int{1, 4, 5, 6, 8, 2}

	fmt.Println("- Sorted array (descending)")
	fmt.Println("- Steps visualization")
	for i := 0; i < len(input)-1; i++ {
		_ , max := utils.FindMinAndMax(input)
		utils.Visualization(max,input)
		for j := 0; j < len(input)-1-i; j++ {	
			if input[j] < input[j+1] {
				var tmp = input[j]
				input[j] = input[j+1]
				input[j+1] = tmp
			}
		}
	}
}
