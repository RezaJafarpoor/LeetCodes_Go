package Easy

//Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
//
//Symbol       Value
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
//For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
//
//Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
//
//I can be placed before V (5) and X (10) to make 4 and 9.
//X can be placed before L (50) and C (100) to make 40 and 90.
//C can be placed before D (500) and M (1000) to make 400 and 900.
//Given a roman numeral, convert it to an integer.
//
//
//
//Example 1:
//
//Input: s = "III"
//Output: 3
//Explanation: III = 3.
//Example 2:
//
//Input: s = "LVIII"
//Output: 58
//Explanation: L = 50, V= 5, III = 3.
//Example 3:
//
//Input: s = "MCMXCIV"
//Output: 1994
//Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
//
//
//Constraints:
//
//1 <= s.length <= 15
//s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
//It is guaranteed that s is a valid roman numeral in the range [1, 3999].

func RomanToInt(input string) int {
	symbols := make(map[byte]int)
	symbols['I'] = 1
	symbols['V'] = 5
	symbols['X'] = 10
	symbols['L'] = 50
	symbols['C'] = 100
	symbols['D'] = 500
	symbols['M'] = 1000
	var result int
	for i := range input {
		result += symbols[input[i]]
		if i != 0 {
			if symbols[input[i-1]] < symbols[input[i]] {
				result -= 2 * symbols[input[i-1]]
			}
		}
	}

	return result
}

func RomanToIntImprovment(input string) int {

	var result int
	for i := range input {
		result += roman(input[i])
		if i != 0 {
			if roman(input[i-1]) < roman(input[i]) {
				result -= 2 * roman(input[i-1])
			}
		}
	}

	return result
}

// some time using a switch case instead of a map is faster

func roman(input byte) int {
	switch input {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}

// if else consuming less memory than switch case in this scenario
func romanIf(input byte) int {

	if input == 'I' {
		return 1
	} else if input == 'V' {
		return 5
	} else if input == 'X' {
		return 10
	} else if input == 'L' {
		return 50
	} else if input == 'C' {
		return 100
	} else if input == 'D' {
		return 500
	} else if input == 'M' {
		return 1000
	}
	return 0

}
