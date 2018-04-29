package repeat

func Repeat(c string) string {
	repeated := ""

	for i := 0; i < 5; i++ {
		repeated += c
	}

	return repeated
}
