package chap2

import (
	"fmt"
	"os"
)

func writeFile() {
	f, err := os.Create("./chap2/create.txt")
	if err != nil {
		fmt.Println("よみこめへん")
	}
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	wannaBeWriteText := "目が乾燥してるんだけどめっちゃ、、、"
	data := []byte(wannaBeWriteText)
	count, err := f.Write(data)
	if err != nil {
		fmt.Println(err)
		fmt.Println("fail to write file")
	}

	fmt.Printf("write %d bytes\n", count)
}
