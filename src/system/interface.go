package system

import (
	"io"
	"math/rand"
	"os"
	"time"
)

// Dir : return True if input string path is a directory
func Dir(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	fileStat, err := file.Stat()
	if err != nil {
		return false
	}
	return fileStat.IsDir()
}

// MakeRange : return a range array between input int(s) min and max
func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// RandString : return a (input int) n-long random string
func RandString(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), SystemLetterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), SystemLetterIdxMax
		}
		if idx := int(cache & SystemLetterIdxMask); idx < len(SystemLetterBytes) {
			b[i] = SystemLetterBytes[idx]
			i--
		}
		cache >>= SystemLetterIdxBits
		remain--
	}

	return string(b)
}

// FileExists : return True if input string path points to a valid file
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// FileTouch : create file in input string path
func FileTouch(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}

// FileCopy : copy file from input string pathFrom to input string pathTo
func FileCopy(pathFrom string, pathTo string) error {
	pathFromOpen, err := os.Open(pathFrom)
	if err != nil {
		return err
	}
	defer pathFromOpen.Close()

	pathToOpen, err := os.Create(pathTo)
	if err != nil {
		return err
	}

	if _, err := io.Copy(pathToOpen, pathFromOpen); err != nil {
		pathToOpen.Close()
		return err
	}

	return pathToOpen.Close()
}