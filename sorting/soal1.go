package sorting

import (
	"fmt"
	"go-magicsoft/utils"
)

func Soal1() {
	input := []int{1, 4, 5, 6, 8, 2}
	_ , max := utils.FindMinAndMax(input)

	fmt.Println(input)
	fmt.Println("OUTPUT: Vertical Barcharts")
	utils.Visualization(max,input)
}

