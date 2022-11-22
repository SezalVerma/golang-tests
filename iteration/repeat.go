package iteration

func Repeat(char string, repeat int) (result string) {
	for i := 1; i <= repeat; i++ {
		result += char
	}
	return
}
