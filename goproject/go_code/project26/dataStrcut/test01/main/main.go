package main

import "fmt"

type ValNode struct {
	row int
	col int
	val int
}

func main() {

	//稀疏数组

	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	var sparseArr []ValNode

	for i, v := range chessMap {

		for j, v2 := range v {

			if v2 != 0 {
				//创建一个ValNode值接节点
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}

		}
	}

	// 输出数组
	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d %d %d", i, valNode.row, valNode.col, valNode.val)
	}

}
