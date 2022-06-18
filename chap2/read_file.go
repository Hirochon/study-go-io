package chap2

import (
	"fmt"
	"os"
)

func readFile() {
	fmt.Println("ほげ")
	f, err := os.Open("./chap2/text.txt")
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
	fmt.Println(f)

	data := make([]byte, 44)
	count, err := f.Read(data)
	if err != nil {
		fmt.Println(err)
		fmt.Println("fail to read file")
	}

	fmt.Printf("read %d bytes:\n", count)
	fmt.Println("文字列", string(data[:count]))
	fmt.Println("string(data)", string(data))
	fmt.Println("data[:count]", data[:count])
	fmt.Println("data", data)
}
