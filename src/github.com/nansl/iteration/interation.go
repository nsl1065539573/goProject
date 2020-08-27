package iteration

func Repeat(str string, n int) string {
	var res string
	for i := 0; i < n; i++ {
		res += str
	}
	return res
}
