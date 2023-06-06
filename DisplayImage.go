package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

func main() {
	filePath := "YourImagePath" // 图片文件路径

	file, err := os.Open(filePath) // 打开图片文件
	if err != nil {
		fmt.Println("无法打开图片文件:", err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file) // 解码图片
	if err != nil {
		fmt.Println("无法解码图片:", err)
		return
	}

	fmt.Printf("图片尺寸：%dx%d\n", img.Bounds().Dx(), img.Bounds().Dy()) // 打印图片尺寸
}
