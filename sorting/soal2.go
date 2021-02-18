package sorting

import (
	"fmt"
	"go-magicsoft/utils"
)

func Soal2() {
	input := []int{1, 4, 5, 6, 8, 2}

	fmt.Println("- Sorted array (ascending)")
	fmt.Println("- Steps visualization")
	
	//var i = 0
	fmt.Println(input)
	for i := len(input)-1; i > 0; i-- {
		_ , max := utils.FindMinAndMax(input)
		utils.Visualization(max,input)
		
		if input[i] < input[i-1] {
			temp := input[i-1]
			input[i-1] = input[i]
			input[i] = temp
		}
	}
}
