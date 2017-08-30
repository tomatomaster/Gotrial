// Copyright © 2017 Ryutarou Ono.

package ex01

type SampleKey struct {
	sam string
	fa  int
}

type Sample struct {
	s map[SampleKey]string
	a map[[5]string]string
}

func main() {
	factor := SampleKey{"factor", 12}
	mapfactor := make(map[SampleKey]string)
	mapfactor[factor] = "test"
	arrayfactor := make(map[[5]string]string)
	var a [5]string
	a[0] = "a"
	a[1] = "ab"
	a[2] = "abc"
	a[3] = "abcd"
	a[4] = "abcde"
	arrayfactor[a] = "test2"

	s := Sample{mapfactor, arrayfactor}
	Display("test", s)
}
