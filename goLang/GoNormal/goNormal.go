package normal

func GetData(x int) []int {
	var res []int
	for i := 0; i < x; i++ {
		res = append(res, i*2)
	}
	return res
}

func PrintData(x []int) {
	var res int
	for id, val := range x {

		res = res + id + val
	}

}

func CompareData(a []byte, b []byte) bool {
	res := true
	for i := range a {
		if a[i] != b[i] {
			res = false
			break
		}
	}
	return res
}
