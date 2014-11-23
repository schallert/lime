// Copyright 2014 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package log_test

import (
	"code.google.com/p/log4go"
	"github.com/limetext/lime/backend/log"
	"sync"
	"testing"
)

type testlogger func(string)

func (l testlogger) LogWrite(rec *log4go.LogRecord) {
	l(rec.Message)
}

func (l testlogger) Close() {}

func TestGlobalLog(t *testing.T) {
	var wg sync.WaitGroup
	log.Global.Close()
	log.Global.AddFilter("globaltest", log.FINEST, testlogger(func(str string) {
		if str != "Testing: hello world" {
			t.Errorf("got: %s", str)
		}
		wg.Done()
	}))
	wg.Add(1)
	log.Info("Testing: %s %s", "hello", "world")
	wg.Wait()
}

func TestLogf(t *testing.T) {
	l := log.NewLogger()
	l.Logf(log.FINEST, "sometest")
	l.Logf(log.FINE, "sometest")
	l.Logf(log.DEBUG, "sometest")
	l.Logf(log.TRACE, "sometest")
	l.Logf(log.INFO, "sometest")
	l.Logf(log.WARNING, "sometest")
	l.Logf(log.ERROR, "sometest")
	l.Logf(log.CRITICAL, "sometest")
	l.Logf(999, "sometest")
}
