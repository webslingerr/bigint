package bigint

func reverse(str string) (res string) {
	for _, v := range str {
		res = string(v)+res
	}
	return
}
