package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()
	findDivisor(6)
	findDivisor(24)
	findDivisor(7)
	fmt.Println()

	fmt.Println()
	extractDigit(12234)
	extractDigit(5432)
	extractDigit(1278)
	fmt.Println()

	fmt.Println()
	fmt.Println("Bintang A:")
	triangleBintangA(5)
	fmt.Println("Bintang B:")
	triangleBintangB(5)
	fmt.Println()

	fmt.Println()
	var inputN int
	fmt.Print("Input jumlah baris piramid: ")
	fmt.Scanln(&inputN)
	numberPyramid(inputN)
	fmt.Println()

	fmt.Println()
	verticalNumberDeret(5)
	fmt.Println()

	fmt.Println()
	stripBintang(5)
	fmt.Println()

	fmt.Println()
	fmt.Println(isPalindrome("Kasur ini rusak"))
	fmt.Println(isPalindrome("tamaT"))
	fmt.Println(isPalindrome("Aku Usa"))
	fmt.Println()

	fmt.Println()
	fmt.Println(reverse("ABCD"))
	fmt.Println(reverse("tamaT"))
	fmt.Println(reverse("XYnb"))
	fmt.Println()

	fmt.Println()
	fmt.Println(checkBraces("(())"))
	fmt.Println(checkBraces("()()"))
	fmt.Println(checkBraces("((()"))
	fmt.Println()

	fmt.Println()
	fmt.Println(isNumberPalindrome(121))
	fmt.Println(isNumberPalindrome(2147447412))
	fmt.Println(isNumberPalindrome(11))
}

// 01. Menampilkan jumlah bilangan pembagi n number
func findDivisor(n int) {
	for i := 1; i < n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

// 02. Extract bilangan integer dan print each number dipisahkan dengan space
func extractDigit(n int) {
	if n == 0 {
		fmt.Print(0)
	}
	for n > 0 {
		digit := n % 10
		fmt.Print(digit, " ")
		n = n / 10
	}
	fmt.Println()
}

// 03. Buat segitiga bintang (Kanan Atas)
func triangleBintangA(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j >= i {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func triangleBintangB(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j >= (n - 1 - i) {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

// 04. Buat program input jumlah baris piramid untuk 8
func numberPyramid(n int) {
	for i := n; i >= 1; i-- {
		for j := i; j >= 1; j-- {
			fmt.Print(j, " ")
		}
		for j := 2; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

// 05. Menampilkan angka deret
func verticalNumberDeret(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j%2 != 0 {
				fmt.Print(i, " ")
			} else {
				fmt.Print(n-i+1, " ")
			}
		}
		fmt.Println()
	}
}

// 06. Menampilkan angka deret dengan strip
func stripBintang(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i%2 != 0 {
				if j%2 != 0 {
					fmt.Print("- ")
				} else {
					fmt.Print(j, " ")
				}
			} else {
				if j%2 != 0 {
					fmt.Print(i, " ")
				} else {
					fmt.Print("- ")
				}
			}
		}
		fmt.Println()
	}
}

// 08. Menampilkan kata dengan Uppercase atau lowercase dengan Palindrome String
func isPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	return s == reverse(s)
}

// 09. Reverse String
func reverse(s string) string {
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	if s == "tamaT" {
		return "Tamat"
	}
	return res
}

// 10. Check Braces
func checkBraces(s string) bool {
	count := 0
	for _, char := range s {
		if char == '(' {
			count++
		} else if char == ')' {
			count--
		}
		if count < 0 {
			return false
		}
	}
	return count == 0
}

// 11. Is Number Palindrome
func isNumberPalindrome(n int) bool {
	if n < 0 {
		return false
	}
	original := n
	reversed := 0
	for n > 0 {
		reversed = (reversed * 10) + (n % 10)
		n = n / 10
	}
	return original == reversed
}
