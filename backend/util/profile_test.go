package util

import (
	"testing"
)

func TestEnter(t *testing.T) {
	profiler := &Profiler{}
	pt := profiler.Enter("some name")

	if &pt == nil {
		t.Error("Returned a nil ProfileToken")
	}
}

func TestExit(t *testing.T) {
	ptoken := &ProfToken{}
	ptoken.Exit()
}

func TestLess(t *testing.T) {
	ps := &prsorter{nil, func(i, j int) bool {
		return i < j
	}}
	result := ps.Less(4, 5)
	if !result {
		t.Error("Less returned incorrect result for 4 < 5")
	}
}

func TestLen(t *testing.T) {
	ps := &prsorter{nil, func(i, j int) bool {
		return i < j
	}}
	if ps.Len() != 0 {
		t.Error("Incorrect length of Data (nil data should be 0 len)")
	}
}

func TestSwap(t *testing.T) {
	pr0 := &ProfileResult{"pr1", *&ProfileEntry{}}
	pr1 := &ProfileResult{"pr2", *&ProfileEntry{}}
	prList := []ProfileResult{*pr0, *pr1}

	ps := &prsorter{prList, func(i, j int) bool { return i < j }}
	ps.Swap(0, 1)

	if ps.data[0] != *pr1 || ps.data[1] != *pr0 {
		t.Error("prsorter swapped incorrectly!")
	}
}
