package iteration

// Repeat takes a character and returns a 5 times replicated value
func Repeat(s string, count int) (ret string) {
	for i := 0; i < count; i++ {
		ret += s
	}
	return
}
