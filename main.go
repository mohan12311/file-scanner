package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type SafeFileNames struct {
	mu    sync.Mutex
	names []string
}

func NewSafeFileNames() *SafeFileNames {
	return &SafeFileNames{names: make([]string, 0)}
}

func (s *SafeFileNames) Add(name string) {
	s.mu.Lock()
	s.names = append(s.names, name)
	s.mu.Unlock()
}

func main() {
	currentDirOnly := flag.Bool("c", false, "Scan only the current directory (no recursive search)")
	outputPtr := flag.String("o", "output.txt", "Output file name")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the directory to scan: ")
	fmt.Println("Current directory: .")
	dir, _ := reader.ReadString('\n')
	dir = strings.TrimSpace(dir)
	dir = filepath.Clean(dir)

	var wg sync.WaitGroup
	sfn := NewSafeFileNames()

	wg.Add(1)
	err := ScanFileNames(dir, &wg, sfn, *currentDirOnly)
	wg.Wait()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = SaveFile(*outputPtr, sfn.names)
	if err != nil {
		fmt.Println(err)
	}
}
