# GO-DisplayImage
Go语言显示路径图片的程序
package main：声明该文件属于main包，这是一个可执行程序的入口包。

import语句：导入所需的包。fmt用于格式化输出，image和image/jpeg用于处理图像文件，os用于操作文件。

func main()：程序的入口函数。

filePath := "YourImagePath"：将图片文件路径存储在变量filePath中。请替换为实际的图片文件路径。

file, err := os.Open(filePath)：使用os.Open函数打开指定路径的图片文件，并将返回的文件对象存储在变量file中。同时，将可能的错误存储在变量err中。

defer file.Close()：使用defer关键字延迟执行文件关闭操作，确保在程序退出前关闭打开的文件。

img, _, err := image.Decode(file)：使用image.Decode函数解码打开的图片文件，并将解码后的图像对象存储在变量img中。同时，将可能的错误存储在变量err中。

fmt.Printf("图片尺寸：%dx%d\n", img.Bounds().Dx(), img.Bounds().Dy())：使用fmt.Printf函数打印图片的尺寸信息。img.Bounds()返回一个image.Rectangle对象，可以通过Dx()和Dy()方法获取图像的宽度和高度。
