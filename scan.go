package main

import (
	"os"
	"path/filepath"
	"sync"
)

/* 디렉토리를 받아서 읽어오는 재귀 함수 */
func ScanFileNames(rootDir string, wg *sync.WaitGroup, sfn *SafeFileNames, currentDirOnly bool) error {
	defer wg.Done()

	/*  디렉토리를 읽어옵니다. */
	entries, err := os.ReadDir(rootDir)
	if err != nil {
		return ErrorReading
	}

	/* []DirEntry 전개하기. */
	if currentDirOnly {
		for _, entry := range entries {
			name := entry.Name()
			if entry.IsDir() {
				name = name + "/"
			}
			sfn.Add(entry.Name())
		}
	} else {
		for _, entry := range entries {
			if entry.IsDir() {
				newPath := filepath.Join(rootDir, entry.Name())
				wg.Add(1)
				go ScanFileNames(newPath, wg, sfn, currentDirOnly)
			} else {
				fileName := rootDir + entry.Name()
				sfn.Add(fileName)
			}
		}
	}

	return nil
}
