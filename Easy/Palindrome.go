package Easy

//Given an integer x, return true if x is a
//palindrome
//, and false otherwise.
//
//
//
//Example 1:
//
//Input: x = 121
//Output: true
//Explanation: 121 reads as 121 from left to right and from right to left.
//Example 2:
//
//Input: x = -121
//Output: false
//Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
//Example 3:
//
//Input: x = 10
//Output: false
//Explanation: Reads 01 from right to left. Therefore, it is not a palindrome.

func IsPalindrome(number int) bool {
	if number < 0 {
		return false

	}
	var digits int
	var copyNumber = number
	var copy2 = number
	for copyNumber != 0 {
		copyNumber = copyNumber / 10
		digits++
	}
	var array = make([]int, digits)

	for i := digits - 1; i >= 0; i-- {
		array[i] = copy2 % 10
		copy2 = copy2 / 10
	}
	var result int
	var index int
	for i := digits; i > 0; i-- {
		result = result + array[i-1]*Power(10, i-1)
		index++
	}

	if result == number {
		return true
	}
	return false

}

func Power(base int, power int) int {
	var result int = base
	if power == 0 {
		return 1
	}
	for i := 1; i < power; i++ {
		result = result * base
	}
	return result
}

func PalindromeFast(number int) bool {
	if number < 0 {
		return false
	}
	x := number
	reversed := 0
	for x != 0 {
		reversed = 10*reversed + x%10
		x /= 10
	}
	return reversed == number
}
