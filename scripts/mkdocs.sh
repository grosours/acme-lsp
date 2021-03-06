#!/bin/bash
# Copyright 2012 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# Generate command docs from usage help.
# Based on https://golang.org/src/cmd/go/mkalldocs.sh

go build -o main.latest
(
	echo '/*'
	./main.latest -help 2>&1
	echo '*/'
	echo 'package main'
) >doc.go
gofmt -w doc.go
rm main.latest
