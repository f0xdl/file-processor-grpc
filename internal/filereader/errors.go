package filereader

import "fmt"

type wrongFileNameError struct {
	path string
}

func WrongFileNameError(path string) *wrongFileNameError {
	return &wrongFileNameError{path: path}
}

func (e *wrongFileNameError) Error() string {
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
