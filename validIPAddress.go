package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//res := validIPAddress("0.16.254.1")
	res := validIPAddress("-0:0db8:85a3:0000:0000:8a2e:0370:7334")

	fmt.Println(res)
}

func validIPAddress(IP string) string {
	if !strings.Contains(IP, ".") && !strings.Contains(IP, ":") {
		return "Neither"
	} else if strings.Contains(IP, ".") {
		return validIPv4(IP)
	} else if strings.Contains(IP, ":") {
		return validIPv6(IP)
	} else {
		return "unknown"
	}
}

func validIPv4(IP string) string {
	val := strings.Split(IP, ".")
	if len(val) != 4 {
		return "Neither"
	}
	var a int64

	for _, r := range val {
		if matched, err := regexp.MatchString("^[0-9]+$", r); !matched || err != nil {
			return "Neither"
		}

		if len(r) > 1 && strings.HasPrefix(r, "0") {
			return "Neither"
		}
		a, _ = strconv.ParseInt(r, 10, 32)
		if a < 0 || a > 255 {
			return "Neither"
		}
	}
	return "IPv4"
}

func validIPv6(IP string) string {
	val := strings.Split(IP, ":")
	if len(val) != 8 {
		return "Neither"
	}
	var err error
	for _, r := range val {
		if len(r) > 4 {
			return "Neither"
		}
		if matched, err := regexp.MatchString("^[0-9A-Za-z]+$", r); !matched || err != nil {
			return "Neither"
		}
		if _, err = strconv.ParseInt(r, 16, 32); err != nil {

			return "Neither"
		}
	}
	return "IPv6"
}
