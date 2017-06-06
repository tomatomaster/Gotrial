package main

import "fmt"

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
	s.Add(127)
	s.Add(128)
	s.Add(27)
	s.Add(17)
	s.Add(12)
	s.Remove(128)
	s.Remove(1)
	s.Remove(17)
	fmt.Println(s.Len())
	s.Remove(127)
	s.Clear()
	fmt.Println(s.Len())
	c := s.Copy()
	c.Add(1)
	fmt.Printf("Clen %v oriLen %v\n", c.Len(), s.Len())

	fmt.Printf("%v %v", s.Len(), s.Has(127))
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
