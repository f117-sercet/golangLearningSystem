package main

import "fmt"

func main() {

	var scores [3][5]float64

	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			fmt.Println("请输入学生成绩\n", i+1, j+1)
			fmt.Scanln(&scores[i][j])
		}

	}
}
