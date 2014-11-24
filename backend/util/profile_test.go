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

func TestString(t *testing.T) {
	Prof.String()
}
