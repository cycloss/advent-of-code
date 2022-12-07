package utils

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func ExitFatal(format string, args ...interface{}) {
	errorMessage := fmt.Sprintf(format, args...)
	_, file, line, ok := runtime.Caller(1)
	var lineInfo string
	if ok {
		lineInfo = fmt.Sprintf("%s:%d", file, line)
	} else {
		lineInfo = "unknown filename and line num"
	}
	log.Fatalf("fatal error %s: %s", lineInfo, errorMessage)
}

func MustGetEnv(key string) string {
	var temp = os.Getenv(key)
	if temp == "" {
		ExitFatal("no value found in environment for key: %s", key)
	}
	return temp
}

// FileWithLineNum returns the file name and line number of the function that it was called by.
// It can be used to determine the approximate location of an error
func FileWithLineNum() string {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		return fmt.Sprintf("%s:%d", file, line)
	} else {
		return "could not find filename and line num"
	}
}

func ErrorfWithLineNum(format string, a ...any) error {
	// cant call FileWithLineNum as this function will be the caller
	_, file, line, ok := runtime.Caller(1)
	var lineNum string
	if ok {
		lineNum = fmt.Sprintf("%s:%d", file, line)
	} else {
		lineNum = "could not find filename and line num"
	}

	errMsg := fmt.Sprintf(format, a...)
	return fmt.Errorf("%s - %s", lineNum, errMsg)
}

func WrapErrorWithLineNum(err error) error {
	_, file, line, ok := runtime.Caller(1)
	var lineNum string
	if ok {
		lineNum = fmt.Sprintf("%s:%d", file, line)
	} else {
		lineNum = "could not find filename and line num"
	}
	return fmt.Errorf("%s - %v", lineNum, err)
}
