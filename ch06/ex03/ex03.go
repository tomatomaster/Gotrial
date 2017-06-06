package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

//
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)                             //word = bisSize
	return word < len(s.words) && s.words[word]&(1<<bit) != 0 //words[2] ならば words[1] words[0]は11111111...1111
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64) //128 word(2) bit(0) [1 0 0]
	for word >= len(s.words) {    //まだwordの場所にbitが何もセットされていないのであれば
		s.words = append(s.words, 0) //000000...0000の状態にセットする
	}
	s.words[word] |= 1 << bit //bitの場所のbitを1に立てる
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func main() {
	var s IntSet
	var t IntSet
	svals := []int{1, 100, 20}
	tvals := []int{10, 100, 20}
	s.AddAll(svals...)
	t.AddAll(tvals...)
	s.DifferenceWith(&t)
	fmt.Println(s)
}

func (s *IntSet) Len() int {
	var sum int
	for _, w := range s.words {
		sum += countBit(w)
	}
	return sum
}

func countBit(w uint64) int {
	var num int
	for ; w != 0; w &= w - 1 {
		num++
	}
	return num
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word > len(s.words) {
		return
	}
	s.words[word] &^= 1 << bit //&^ビットクリア AND NOT指定ビットが立っていたらクリアする
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() *IntSet {
	c := make([]uint64, len(s.words), cap(s.words))
	copy(c, s.words)
	return &IntSet{words: c}
}

func (s *IntSet) AddAll(inputs ...int) {
	for _, i := range inputs {
		s.Add(i)
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	sword := len(s.words)
	tword := len(t.words)
	m := min(sword, tword)
	for i := 0; i <= len(s.words); i++ {
		if i <= m {
			s.words[i] &= t.words[i]
		}
	}
}

func min(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	min := min(len(s.words), len(t.words)) - 1
	for i := 0; i <= len(s.words); i++ {
		if i <= min {
			s.words[i] ^= t.words[i]
		} else {
			s.words[i] &= ^uint64(0)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
