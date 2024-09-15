package Easy

func LongestCommonPrefix(input []string) string {
	temp := input[0]
	for _, str := range input {
		i := 0
		for ; i < len(str) && i < len(temp) && temp[i] == str[i]; i++ {
		}
		temp = temp[:i]
	}
	return temp

}
