package main

import (
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 3 {
		return
	}
	if len(arg[1]) != 1 {
		return
	}
	a, err := Atoi23(arg[0])
	b, err1 := Atoi23(arg[2])
	if a == 9223372036854775807 || b == 9223372036854775807 || a == -9223372036854775807 || b == -9223372036854775807 {
		return
	}
	s := arg[1][0]
	if err != "" || err1 != "" {
		return
	}
	res := 0
	if s == '+' {
		res = a + b
	} else if s == '-' {
		res = a - b
	} else if s == '*' {
		res = a * b
	} else if s == '/' {
		if b == 0 {
			PrintStr("No division by 0")
			return
		}
		res = a / b
	} else if s == '%' {
		if b == 0 {
			PrintStr("No modulo by 0")
			return
		}
		res = a % b
	} else {
		return
	}
	PrintNbr(res)
}

func PrintStr(s string) {
	s += "\n"
	os.Stdout.Write([]byte(s))
}

func Atoi23(s string) (int, string) {
	if len(s) == 0 {
		return 0, ""
	}
	res := 0
	sign := 0
	m := 1
	if s[0] == '-' {
		sign = -1
	} else if s[0] == '+' {
		sign = 1
	}
	end := 0
	if sign != 0 {
		end = 1
	}
	for i := len(s) - 1; i >= end; i-- {
		if rune(s[i])-'0' > 9 || rune(s[i])-'0' < 0 {
			return 0, "err"
		}
		res += m * int(rune(s[i])-'0')
		m *= 10
	}
	if sign == 0 {
		sign = 1
	}
	return res * sign, ""
}

func PrintNbr(n int) {
	ans := ""
	if n >= 0 {
		if n == 0 {
			ans += "0"
		}
		for n > 0 {
			last := n % 10
			ans = string('0'+rune(last)) + ans
			n = n / 10
		}
		PrintStr(ans)
	} else {
		for n <= -1 {
			last := n % 10
			last *= -1
			ans = string('0'+rune(last)) + ans
			n = n / 10
		}
		ans = "-" + ans
		PrintStr(ans)
	}
}
