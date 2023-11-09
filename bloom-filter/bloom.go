package main

import (
	"encoding/binary"
	"fmt"
	"github.com/bits-and-blooms/bitset"
	"hash"
	"hash/fnv"
	"math"
)

type Bloom interface {
	Add(key []byte) Bloom
	Check(key []byte) (bool, float64)
	PrintStats()
	SetHashFunction(hash.Hash)
	Reset()
	FillRatio() float64
	EstimatedFillRatio() float64
}

func K(e float64) uint {
	return uint(math.Ceil(math.Log2(1 / e)))
}

func M(n uint, p, e float64) uint {
	// m =~ n / ((log(p)*log(1-p))/abs(log e))
	return uint(math.Ceil(float64(n) / ((math.Log(p) * math.Log(1-p)) / math.Abs(math.Log(e)))))
}

type BloomFunction struct {
	// h is the hash function
	h hash.Hash

	// the total number of bits for the bloom filter.
	// will be divided into k partitions, or slices. So each partition contains Math.ceil(m/k) bits.
	// m =~ n / ((log(p)*log(1-p))/abs(log e))
	m uint

	// the number of hash values used to set and test bits.
	k uint

	// the size of the partition, or slice.
	// s = m / k
	s uint

	// the fill ratio of the filter partitions. It's mainly used to calculate m at the start.
	// p is not checked when new items are added. So if the fill ratio goes above p, the likelihood
	// of false positives (error rate) will increase.
	//
	// By default, we use the fill ratio of p = 0.5
	p float64

	// the desired error rate of the bloom filter. The lower the e, the higher the k.
	// By default, we use the error rate of e = 0.1% = 0.001.
	e float64

	// number of elements the filter is predicted to hold
	n uint

	// b is the set of bit array holding the bloom filters. There will be k b's.
	b *bitset.BitSet

	// number of items we have added to the filter
	c uint

	// the list of bits
	bs []uint
}

var _ Bloom = (*BloomFunction)(nil)

func NewBloomFilter(n uint) Bloom {
	var (
		p float64 = 0.5
		e float64 = 0.001
		k uint    = K(e)
		m uint    = M(n, p, e)
	)

	return &BloomFunction{
		h:  fnv.New64(),
		n:  n,
		p:  p,
		e:  e,
		k:  k,
		m:  m,
		b:  bitset.New(m),
		bs: make([]uint, k),
	}
}

func (f *BloomFunction) SetHashFunction(h hash.Hash) {
	f.h = h
}

func (f *BloomFunction) Reset() {
	f.k = K(f.e)
	f.m = M(f.n, f.p, f.e)
	f.b = bitset.New(f.m)
	f.bs = make([]uint, f.k)

	if f.h == nil {
		f.h = fnv.New64()
	} else {
		f.h.Reset()
	}
}

func (f *BloomFunction) EstimatedFillRatio() float64 {
	return 1 - math.Exp((-float64(f.c)*float64(f.k))/float64(f.m))
}

func (f *BloomFunction) FillRatio() float64 {
	return float64(f.b.Count()) / float64(f.m)
}

func (f *BloomFunction) Add(item []byte) Bloom {
	f.bits(item)
	for _, v := range f.bs[:f.k] {
		f.b.Set(v)
	}
	f.c++
	return f
}

func (f *BloomFunction) Check(item []byte) (bool, float64) {
	f.bits(item)
	for _, v := range f.bs[:f.k] {
		if !f.b.Test(v) {
			return false, 0
		}
	}
	fPosProb := math.Pow(1.0-math.Pow(1.0-1.0/float64(f.m), float64(f.k*f.n)), float64(f.k))

	return true, fPosProb
}

func (f *BloomFunction) Count() uint {
	return f.c
}

func (f *BloomFunction) PrintStats() {
	fmt.Printf("m = %d, n = %d, k = %d, s = %d, p = %f, e = %f\n", f.m, f.n, f.k, f.s, f.p, f.e)
	fmt.Println("Total items:", f.c)
	c := f.b.Count()
	fmt.Printf("Total bits set: %d (%.1f%%)\n", c, float32(c)/float32(f.m)*100)
}

func (f *BloomFunction) bits(item []byte) {
	f.h.Reset()
	f.h.Write(item)
	s := f.h.Sum(nil)
	a := binary.BigEndian.Uint32(s[4:8])
	b := binary.BigEndian.Uint32(s[0:4])

	for i := range f.bs[:f.k] {
		f.bs[i] = (uint(a) + uint(b)*uint(i)) % f.m
	}
}
