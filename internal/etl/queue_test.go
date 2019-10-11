// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package etl

import (
	"testing"
	"time"
)

func TestNewTaskID(t *testing.T) {
	// Verify that the task ID is the same within taskIDChangeInterval and changes
	// afterwards.
	const (
		module  = "mod"
		version = "ver"
	)

	tm := time.Now().Truncate(taskIDChangeInterval)
	id1 := newTaskID(module, version, tm)
	id2 := newTaskID(module, version, tm.Add(taskIDChangeInterval/2))
	if id1 != id2 {
		t.Error("wanted same task ID, got different")
	}
	id3 := newTaskID(module, version, tm.Add(taskIDChangeInterval+1))
	if id1 == id3 {
		t.Error("wanted different task ID, got same")
	}
}
