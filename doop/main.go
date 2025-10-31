package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	arg1 := os.Args[1]
	operator := os.Args[2]
	arg3 := os.Args[3]

	a, ok1 := safeAtoi(arg1)
	b, ok2 := safeAtoi(arg3)

	if !ok1 || !ok2 {
		return
	}

	switch operator {
	case "+":
		result, overflow := addCheckOverflow(a, b)
		if overflow {
			return
		}
		itoa(result)
	case "-":
		result, overflow := subCheckOverflow(a, b)
		if overflow {
			return
		}
		itoa(result)
	case "*":
		result, overflow := mulCheckOverflow(a, b)
		if overflow {
			return
		}
		itoa(result)
	case "/":
		if b == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		result, overflow := divCheckOverflow(a, b)
		if overflow {
			return
		}
		itoa(result)
	case "%":
		if b == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		result := a % b
		itoa(result)
	default:
		return
	}
}

// safeAtoi converts string to int64 and checks for overflow
func safeAtoi(s string) (int64, bool) {
	var result int64
	var sign int64 = 1
	i := 0

	if len(s) > 0 {
		if s[0] == '+' {
			i++
		} else if s[0] == '-' {
			sign = -1
			i++
		}
	}

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		digit := int64(s[i] - '0')

		if result > (1<<63-1-digit)/10 {
			return 0, false
		}
		result = result*10 + digit
	}

	return result * sign, true
}

// itoa converts int64 to string and writes to stdout
func itoa(n int64) {
	if n == 0 {
		os.Stdout.WriteString("0\n")
		return
	}

	var buf [20]byte
	i := len(buf) - 1
	negative := n < 0

	if negative {
		n = -n
	}

	for n > 0 {
		buf[i] = byte('0' + n%10)
		n /= 10
		i--
	}

	if negative {
		buf[i] = '-'
		i--
	}

	os.Stdout.Write(buf[i+1:])
	os.Stdout.WriteString("\n")
}

// Overflow check functions
func addCheckOverflow(a, b int64) (int64, bool) {
	if b > 0 && a > (1<<63-1)-b {
		return 0, true
	}
	if b < 0 && a < (-1<<63)-b {
		return 0, true
	}
	return a + b, false
}

func subCheckOverflow(a, b int64) (int64, bool) {
	if b < 0 && a > (1<<63-1)+b {
		return 0, true
	}
	if b > 0 && a < (-1<<63)+b {
		return 0, true
	}
	return a - b, false
}

func mulCheckOverflow(a, b int64) (int64, bool) {
	if a == 0 || b == 0 {
		return 0, false
	}

	result := a * b
	if a == (1<<63-1) && b == -1 {
		return 0, true
	}
	if b == (1<<63-1) && a == -1 {
		return 0, true
	}
	if a == -1<<63 && b == -1 {
		return 0, true
	}
	if result/b != a {
		return 0, true
	}
	return result, false
}

func divCheckOverflow(a, b int64) (int64, bool) {
	if a == -1<<63 && b == -1 {
		return 0, true
	}
	return a / b, false
}
