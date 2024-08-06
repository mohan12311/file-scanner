package main

import (
	"os"
	"path/filepath"
	"sync"
	"testing"
)

func setupTestDir(t *testing.T, dir string, files []string) {
	t.Helper()
	// 디렉토리 생성
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	// 파일 생성
	for _, file := range files {
		f, err := os.Create(filepath.Join(dir, file))
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
		f.Close()
	}
}

func teardownTestDir(t *testing.T, dir string) {
	t.Helper()
	// 디렉토리 삭제
	if err := os.RemoveAll(dir); err != nil {
		t.Fatalf("Failed to remove test directory: %v", err)
	}
}

func TestScanFileNames(t *testing.T) {
	type args struct {
		rootDir        string
		wg             *sync.WaitGroup
		sfn            *SafeFileNames
		currentDirOnly bool
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		setup    func()
		teardown func()
	}{
		{
			name: "Valid directory with subdirectories",
			args: args{
				rootDir:        "testdata/valid",
				wg:             &sync.WaitGroup{},
				sfn:            NewSafeFileNames(),
				currentDirOnly: false,
			},
			wantErr: false,
			setup: func() {
				setupTestDir(t, "testdata/valid", []string{"file1.txt", "file2.txt"})
				setupTestDir(t, "testdata/valid/subdir", []string{"file3.txt"})
			},
			teardown: func() {
				teardownTestDir(t, "testdata/valid")
			},
		},
		{
			name: "Empty directory",
			args: args{
				rootDir:        "testdata/empty",
				wg:             &sync.WaitGroup{},
				sfn:            NewSafeFileNames(),
				currentDirOnly: false,
			},
			wantErr: false,
			setup: func() {
				setupTestDir(t, "testdata/empty", nil)
			},
			teardown: func() {
				teardownTestDir(t, "testdata/empty")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			defer func() {
				if tt.teardown != nil {
					tt.teardown()
				}
			}()

			tt.args.wg.Add(1)
			err := ScanFileNames(tt.args.rootDir, tt.args.wg, tt.args.sfn, tt.args.currentDirOnly)
			tt.args.wg.Wait()

			if (err != nil) != tt.wantErr {
				t.Errorf("ScanFileNames() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
