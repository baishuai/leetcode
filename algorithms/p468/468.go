package p468

import (
	"strconv"
	"strings"
)

/**
Write a function to check whether an input string is a valid IPv4 address or IPv6 address or neither.

IPv4 addresses are canonically represented in dot-decimal notation, which consists of four decimal numbers, each ranging from 0 to 255, separated by dots ("."), e.g.,172.16.254.1;

Besides, leading zeros in the IPv4 is invalid. For example, the address 172.16.254.01 is invalid.

IPv6 addresses are represented as eight groups of four hexadecimal digits, each group representing 16 bits. The groups are separated by colons (":"). For example, the address 2001:0db8:85a3:0000:0000:8a2e:0370:7334 is a valid one. Also, we could omit some leading zeros among four hexadecimal digits and some low-case characters in the address to upper-case ones, so 2001:db8:85a3:0:0:8A2E:0370:7334 is also a valid IPv6 address(Omit leading zeros and using upper cases).

However, we don't replace a consecutive group of zero value with a single empty group using two consecutive colons (::) to pursue simplicity. For example, 2001:0db8:85a3::8A2E:0370:7334 is an invalid IPv6 address.

Besides, extra leading zeros in the IPv6 is also invalid. For example, the address 02001:0db8:85a3:0000:0000:8a2e:0370:7334 is invalid.

Note: You may assume there is no extra space or special characters in the input string.

**/

func validV4(v4s []string) bool {
	if v4s == nil || len(v4s) != 4 {
		return false
	}
	for _, v := range v4s {
		//0...255
		vi, err := strconv.ParseUint(v, 10, 0)
		if err != nil || vi > 255 || len(v) != len(strconv.FormatUint(vi, 10)) {
			return false
		}
	}
	return true
}

func validV6(v6s []string) bool {
	if v6s == nil || len(v6s) != 8 {
		return false
	}
	for _, v := range v6s {
		//0...255
		vi, err := strconv.ParseUint(v, 16, 0)
		if err != nil || vi > 0xFFFF || len(v) > 4 {
			return false
		}
	}
	return true
}

func validIPAddress(IP string) string {
	if strings.Contains(IP, ".") {
		v4 := strings.Split(IP, ".")
		if validV4(v4) {
			return "IPv4"
		}
	} else if strings.Contains(IP, ":") {
		v6 := strings.Split(IP, ":")
		if validV6(v6) {
			return "IPv6"
		}
	}
	return "Neither"
}
