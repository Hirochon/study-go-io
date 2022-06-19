package chap5

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	Fsize = 10000
)

func Chap5() {
	f, err := os.Open("./chap5/sample.txt")
	if err != nil {
		fmt.Println("みすってる！", err.Error())
		return
	}
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	ReadOS(f, 100)
}

// ReadOS はサイズがFsizeのファイルをnbyteごと読む関数
func ReadOS(r io.Reader, n int) {
	data := make([]byte, n)

	t := Fsize / n
	for i := 0; i < t; i++ {
		r.Read(data)
	}
}

func ScanOS() {
	f, _ := os.Open("./chap5/sample2.txt")
	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		line := sc.Text()
		fmt.Println(line)
	}
}
