package utils

import "fmt"

func FindMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func Visualization(max int, arr []int){
	for i := max; i > 0; i-- {
		for j := 0; j < len(arr); j++ {
			if i <= arr[j] {
				fmt.Print(" | ")		
			}else{
				fmt.Printf("   ")		
			}
		}
		fmt.Println("")
	}
	for i := 0; i < len(arr); i++ {
		fmt.Printf(" %d ",arr[i])		
	}
	fmt.Println("")
}
