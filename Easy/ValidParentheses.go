package Easy

func ValidParentheses(s string) bool {
	var openP, openB, openC, closeP, closeB, closeC int
	for index := range len(s) {
		switch s[index] {
		case '(':
			openP++
		case ')':
			closeP++
		case '[':
			openB++
		case ']':
			closeB++
		case '{':
			openC++
		case '}':
			closeC++
		}
	}
	return openP == closeP && openB == closeB && closeC == closeC
}
