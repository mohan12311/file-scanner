package main

import (
	"os"
	"testing"
)

func TestSaveFile(t *testing.T) {
	type args struct {
		saveFileName     string
		scannedFileNames []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "EmptyFileName",
			args:    args{saveFileName: "", scannedFileNames: []string{"file1.txt", "file2.txt"}},
			wantErr: true,
		},
		{
			name:    "EmptyScannedFileNames",
			args:    args{saveFileName: "output.txt", scannedFileNames: []string{}},
			wantErr: false,
		},
		{
			name:    "ValidFileNameAndScannedFileNames",
			args:    args{saveFileName: "output.txt", scannedFileNames: []string{"file1.txt", "file2.txt"}},
			wantErr: false,
		},
		{
			name:    "InvalidPath",
			args:    args{saveFileName: "/invalid_path/output.txt", scannedFileNames: []string{"file1.txt"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SaveFile(tt.args.saveFileName, tt.args.scannedFileNames)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			// 파일이 생성되었는지 확인하고 삭제
			if err == nil && tt.args.saveFileName != "" {
				if _, err := os.Stat(tt.args.saveFileName); os.IsNotExist(err) {
					t.Errorf("SaveFile() failed to create file %v", tt.args.saveFileName)
				} else {
					os.Remove(tt.args.saveFileName)
				}
			}
		})
	}
}
