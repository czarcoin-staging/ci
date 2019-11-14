// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

// +build ignore

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"storj.io/storj/private/memory"
)

var ignoreFolder = map[string]bool{
	".build":       true,
	".git":         true,
	"node_modules": true,
	"coverage":     true,
	"dist":         true,
}

func main() {
	const fileSizeLimit = 650 * memory.KB

	var failed int

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if info.IsDir() && ignoreFolder[info.Name()] {
			return filepath.SkipDir
		}

		size := memory.Size(info.Size())
		if size > fileSizeLimit {
			failed++
			fmt.Printf("%v (%v)\n", path, size)
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	if failed > 0 {
		fmt.Printf("some files were over size limit %v\n", fileSizeLimit)
		os.Exit(1)
	}
}
