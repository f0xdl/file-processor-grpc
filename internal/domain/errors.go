package domain

import "fmt"

type WrongFileNameErr struct {
	path string
}

func WrongFileNameError(path string) *WrongFileNameErr {
	return &WrongFileNameErr{path: path}
}

func (e *WrongFileNameErr) Error() string {
	return fmt.Sprintf("wrong file name in path: '%s'", e.path)
}

// type fileNotFoundError struct {
// 	path string
// }

// func FileNotFoundError(path string) *fileNotFoundError {
// 	return &fileNotFoundError{path: path}
// }

// func (e *fileNotFoundError) Error() string {
// 	return "file not found: " + e.path
// }
