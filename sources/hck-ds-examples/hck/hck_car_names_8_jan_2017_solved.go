package main

import "fmt"

func main() {
	var i, t, count, prevCount uint8

	var s string
	fmt.Scanf("%d", &t)

	for i = 0; i < t; i++ {
		fmt.Scanf("%s", &s)

		n := 0
		ok := true
		m := make(map[byte]byte)
		for k := 0; k < len(s); k++ {
			ch := s[k]
			if n > 1 {
				_, found := m[s[k]]
				if found {
					ok = false
					break
				}
			}
			m[s[k]] = 65

			j := k
			count = 0
			for j < len(s) && byte(ch) == s[j] {
				j++
				count += 1
			}
			k = j - 1

			n += 1
			if n > 3 {
				ok = false
				break
			}
			if n > 1 {
				if count != prevCount {
					ok = false
					break
				}
			}
			prevCount = count
		}
		if ok && n == 3 {
			fmt.Println("OK")
		} else {
			fmt.Println("Not OK")
		}
	}
}
