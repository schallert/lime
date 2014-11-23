// Copyright 2014 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package parser_test

import (
	"bytes"
	"github.com/limetext/lime/backend/render"
	"github.com/limetext/text"
	"github.com/quarnster/parser"
	"sort"
	"sync"
)

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
