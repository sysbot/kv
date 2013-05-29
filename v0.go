// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//TODO -> kv.go

package kv

import (
	"os"

	"github.com/cznic/exp/lldb"
)

func open00(name string, in *DB) (db *DB, err error) {
	db = in
	if db.alloc, err = lldb.NewAllocator(lldb.NewInnerFiler(db.filer, 16), &lldb.Options{}); err != nil {
		return nil, &os.PathError{Op: "dbm.Open", Path: name, Err: err}
	}

	db.alloc.Compress = true
	return db, db.boot()
}
