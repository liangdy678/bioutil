
package utils

import (
	"bufio"
	"compress/gzip"
	"io"

	"os"
)

// 以下为字符串处理函数
func DealFromFile(file string, handle func(s string) bool) error {
	fi, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fi.Close()
	return DealFromReader(fi, handle)
}

func DealFromGzip(file string, handle func(s string) bool) error {
	fi, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fi.Close()

	gzfi, err := gzip.NewReader(fi)
	if err != nil {
		return err
	}
	defer gzfi.Close()

	return DealFromReader(gzfi, handle)

}

func DealFromReader(fi io.Reader, handle func(s string) bool) error {

	scanner := newScanner(fi)

	for scanner.Scan() {
		data := scanner.Text()
		ok := handle(data)
		if !ok {
			break
		}
	}

	return scanner.Err()

}

// 以下为字节流处理函数
func DealByteFromGzip(file string, handle func(b []byte) bool) error {
	fi, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fi.Close()
	gzfi, err := gzip.NewReader(fi)
	if err != nil {
		return err
	}
	defer gzfi.Close()
	return DealByteFromReader(gzfi, handle)
}

func DealByteFromFile(file string, handle func(b []byte) bool) error {
	fi, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fi.Close()
	return DealByteFromReader(fi, handle)
}

func DealByteFromReader(fi io.Reader, handle func(b []byte) bool) error {

	scanner := newScanner(fi)

	for scanner.Scan() {
		ok := handle(scanner.Bytes())
		if !ok {
			break
		}
	}
	return scanner.Err()

}

func newScanner(fi io.Reader) *bufio.Scanner {
	// bufio.Scanner 默认使用 bufio.MaxScanTokenSize 作为缓存大小，
	scanner := bufio.NewScanner(fi)

	// 这里可设置一个较大的缓存大小，以便一次性读取整个文件内容。
	buf := make([]byte, 0, bufio.MaxScanTokenSize)
	scanner.Buffer(buf, bufio.MaxScanTokenSize*2)

	return scanner
}
