package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

// 图片bmp文件头二进制读取
func main() {
	file, err := os.Open("./img.bmp")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var headA, headB byte
	binary.Read(file, binary.LittleEndian, &headA)
	binary.Read(file, binary.LittleEndian, &headB)

	var size uint32
	binary.Read(file, binary.LittleEndian, &size)

	fmt.Printf("%c %c\n", headA, headB)
	fmt.Println(size / 1000)
}
