package main

func SetWay(mp *[8][7]int, i int, j int) bool {

	// 找到通路
	if mp[6][5] == 2 {

		return true
	} else {
		// 继续找
		if mp[i][j] == 0 {

			mp[i][j] = 2
			if SetWay(mp, i-1, j) {
				return true
			}

		} else if SetWay(mp, i+1, j) {
			return true
		} else if SetWay(mp, i+1, j-1) {
			return true
		} else if SetWay(mp, i+1, j+1) {
			return true
		} else {
			mp[i][j] = 3
			return false
		}
	}
	return false
}

func main() {

	// 1：阻碍
	//  0：没走过
	// 2：能走通
	// 3：曾经走过，没通
	var Map [8][7]int

	for i := 0; i < 7; i++ {
		Map[0][i] = 1
		Map[7][i] = 1
	}

	for i := 0; i < 8; i++ {
		Map[i][0] = 1
		Map[i][6] = 1
	}

	SetWay(&Map, 1, 1)

}
