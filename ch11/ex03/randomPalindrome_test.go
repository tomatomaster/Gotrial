// Copyright © 2017 Ryutarou Ono.

package main

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

func TestNotPlindrome(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000000; i++ {
		rs := randomNotPlindrome(rng)
		if IsPalindrome(rs) {
			t.Errorf("Actual %v \n", rs)
		}
	}
}

func TestRandomPalindromeStartWithPunctuation(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000000; i++ {
		rs := randomPalindromeStartWithPunctuation(rng)
		if !IsPalindrome(rs) {
			t.Errorf("Actual %v \n", rs)
		}
	}
}

func randomPalindromeStartWithPunctuation(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		if i == 0 {
			r := rune(rng.Intn(0x2E))
			runes[i] = r
			r = rune(rng.Intn(0x2E))
			runes[n-1-i] = r
			continue
		}
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func randomNotPlindrome(rng *rand.Rand) string {
	n := rng.Intn(25) + 3
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(0x41 + rng.Intn(10))
		runes[i] = r
		for {
			r = rune(int(r) + 10)
			if unicode.IsLetter(r) {
				break
			}
		}
		runes[n-1-i] = r
	}
	return string(runes)
}
