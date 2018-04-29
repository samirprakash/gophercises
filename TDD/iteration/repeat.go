package iteration

const count = 5

// Repeat takes a character and returns a 5 times replicated value
func Repeat(s string) (ret string) {
	for i := 0; i < count; i++ {
		ret += s
	}
	return
}
