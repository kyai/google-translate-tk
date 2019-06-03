package tk

import (
	"fmt"
	"strconv"
	"strings"
)

func GetTK(s, tkk string) string {
	a := ascii(s)
	e := make([]uint32, 0)
	for f, g := 0, 0; g < len(a); g++ {
		m := uint32(a[g])
		if 128 > m {
			ff := f
			f++
			e = fill(e, ff, m)
		} else {
			if 2048 > m {
				ff := f
				f++
				e = fill(e, ff, m>>6|192)
			} else {
				if 55296 == (m&64512) && g+1 < len(a) && 56320 == int(a[g+1])&64512 {
					g++
					m = 65536 + ((m & 1023) << 10) + (uint32(a[g]) & 1023)
					ff := f
					f++
					e = fill(e, ff, m>>18|240)
					ff = f
					f++
					e = fill(e, ff, m>>12&63|128)
				} else {
					ff := f
					f++
					e = fill(e, ff, m>>12|224)
					ff = f
					f++
					e = fill(e, ff, m>>6&63|128)
					ff = f
					f++
					e = fill(e, ff, m&63|128)
				}
			}
		}
	}

	ka := suint32(strings.Split(tkk, ".")[0])
	kb := suint32(strings.Split(tkk, ".")[1])

	Sb := "+-a^+6"
	Zb := "+-3^+b+-f"

	aa := ka
	for f := 0; f < len(e); f++ {
		aa += e[f]
		aa = r(aa, Sb)
	}
	aa = r(aa, Zb)
	aa ^= kb
	aa %= 1E6

	return fmt.Sprintf("%v.%v", aa, aa^ka)
}

func r(a uint32, b string) uint32 {
	t := "a"
	Yb := "+"
	for c := 0; c < len(b)-2; c += 3 {
		d := string(b[c+2])
		var dd uint32
		if d >= t {
			dd = uint32(d[0]) - 87
		} else {
			dd = suint32(d)
		}
		if string(b[c+1]) == Yb {
			dd = a >> dd
		} else {
			dd = a << dd
		}
		if string(b[c]) == Yb {
			a = a + dd&4294967295
		} else {
			a = a ^ dd
		}
	}
	return a
}

func suint32(s string) uint32 {
	i, _ := strconv.Atoi(s)
	return uint32(i)
}

func ascii(s string) []int32 {
	ints := make([]int32, 0)
	for _, v := range strings.Split(s, "") {
		var i64 int64
		if len(v) == 1 {
			i64 = int64(v[0])
		} else {
			quoted := strconv.QuoteToASCII(v)
			unquoted := quoted[1 : len(quoted)-1]
			unquoted = strings.Replace(unquoted, "\\u", "", 1)
			i64, _ = strconv.ParseInt(unquoted, 16, 32)
		}
		ints = append(ints, int32(i64))
	}
	return ints
}

func fill(slice []uint32, index int, value uint32) []uint32 {
	for {
		if len(slice) > index {
			break
		}
		slice = append(slice, 0)
	}
	slice[index] = value
	return slice
}
