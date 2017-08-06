// Copyright © 2017 Ryutarou Ono.

package ex07

import (
	"math/rand"
	"testing"
	"time"

	"./intset"
	"./intset8"
)

func BenchmarkAdd(b *testing.B) {
	var x intset.IntSet
	seed := time.Now().UTC().UnixNano()
	rand := rand.New(rand.NewSource(seed))
	for i := 0; i < b.N; i++ {
		randomVal := rand.Intn(0xfffff)
		x.Add(randomVal)
	}
}

func BenchmarkAdd8(b *testing.B) {
	var x intset8.IntSet
	seed := time.Now().UTC().UnixNano()
	rand := rand.New(rand.NewSource(seed))
	for i := 0; i < b.N; i++ {
		randomVal := rand.Intn(0xfffff)
		x.Add(randomVal)
	}
}

func BenchmarkUnionWith(b *testing.B) {
	var x intset.IntSet
	var y intset.IntSet
	seed := time.Now().UTC().UnixNano()
	rand := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		randomValx := rand.Intn(0xfffff)
		randomValy := rand.Intn(0xfffff)
		x.Add(randomValx)
		y.Add(randomValy)
	}

	for i := 0; i < b.N; i++ {
		x.UnionWith(&y)
	}
}

func BenchmarkUnionWith8(b *testing.B) {
	var x intset8.IntSet
	var y intset8.IntSet
	seed := time.Now().UTC().UnixNano()
	rand := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		randomValx := rand.Intn(0xfffff)
		randomValy := rand.Intn(0xfffff)
		x.Add(randomValx)
		y.Add(randomValy)
	}

	for i := 0; i < b.N; i++ {
		x.UnionWith(&y)
	}
}

func BenchmarkHas(b *testing.B) {
	var x intset.IntSet
	seed := time.Now().UTC().UnixNano()
	rand := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		randomValx := rand.Intn(0xfffff)
		x.Add(randomValx)
	}

	for i := 0; i < b.N; i++ {
		x.Has(rand.Intn(0xfffff))
	}
}

func BenchmarkHas8(b *testing.B) {
	var x intset8.IntSet
	seed := time.Now().UTC().UnixNano()
	rand := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		randomValx := rand.Intn(0xfffff)
		x.Add(randomValx)
	}

	for i := 0; i < b.N; i++ {
		x.Has(rand.Intn(0xfffff))
	}
}
