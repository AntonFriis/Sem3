package main

import (
	"fmt"
	"strconv"
)

func main() {
	var P, K, H, T [13]bool
	var s string
	fmt.Scanln(&s)
	for i := 0; i < len(s); i += 3 {
		var c = s[i : i+3]
		var suit = string(c[0])
		var number, _ = strconv.Atoi(c[1:3])
		number--
		if suit == "P" {
			if P[number] == true {
				fmt.Println("GRESKA")
				return
			} else {
				P[number] = true
			}
		} else if suit == "K" {
			if K[number] == true {
				fmt.Println("GRESKA")
				return
			} else {
				K[number] = true
			}
		} else if suit == "H" {
			if H[number] == true {
				fmt.Println("GRESKA")
				return
			} else {
				H[number] = true
			}
		} else if suit == "T" {
			if T[number] == true {
				fmt.Println("GRESKA")
				return
			} else {
				T[number] = true
			}
		}
	}

	var p, k, h, t = 13, 13, 13, 13
	for i := 0; i < 13; i++ {
		if P[i] == true {
			p--
		}
		if K[i] == true {
			k--
		}
		if H[i] == true {
			h--
		}
		if T[i] == true {
			t--
		}
	}
	fmt.Println(p, k, h, t)
}
