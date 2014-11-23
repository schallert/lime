package log_test

import (
	"fmt"
	"github.com/limetext/lime/backend/log"
	"testing"
)

func TestNewLogWriter(t *testing.T) {
	l := log.NewLogWriter(func(str string) { fmt.Print("yo") })
	if l == nil {
		t.Error("NewLogWriter produced a nil")
	}
}
