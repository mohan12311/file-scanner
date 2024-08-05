package main

import "errors"

var (
	ErrorReading    = errors.New("error Reading Directory")
	ErrorCreateFile = errors.New("error Saving To File")
	ErrorWriteFile  = errors.New("error Writing File list")
)
