package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

func main() {
	// 创建一个600x400的图像
	img := image.NewRGBA(image.Rect(0, 0, 600, 400))

	// 填充白色背景
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// 添加边框
	for x := 0; x < 600; x++ {
		img.Set(x, 0, color.Black)
		img.Set(x, 399, color.Black)
	}
	for y := 0; y < 400; y++ {
		img.Set(0, y, color.Black)
		img.Set(599, y, color.Black)
	}

	// 在图像上添加文字位置标记（模拟身份证布局）
	// 这只是创建一个简单的测试图像
	drawRect(img, 50, 50, 200, 30, color.RGBA{200, 200, 200, 255})   // 姓名区域
	drawRect(img, 50, 100, 100, 30, color.RGBA{200, 200, 200, 255})  // 性别区域
	drawRect(img, 200, 100, 100, 30, color.RGBA{200, 200, 200, 255}) // 民族区域
	drawRect(img, 50, 150, 250, 30, color.RGBA{200, 200, 200, 255})  // 出生日期区域
	drawRect(img, 50, 200, 400, 60, color.RGBA{200, 200, 200, 255})  // 住址区域
	drawRect(img, 50, 280, 300, 30, color.RGBA{200, 200, 200, 255})  // 身份证号区域

	// 保存图像
	file, err := os.Create("test_id.jpg")
	if err != nil {
		fmt.Println("创建文件失败:", err)
		return
	}
	defer file.Close()

	err = jpeg.Encode(file, img, &jpeg.Options{Quality: 90})
	if err != nil {
		fmt.Println("保存图像失败:", err)
		return
	}

	fmt.Println("测试图像 test_id.jpg 已创建")
}

func drawRect(img *image.RGBA, x, y, width, height int, c color.Color) {
	for i := x; i < x+width && i < 600; i++ {
		for j := y; j < y+height && j < 400; j++ {
			img.Set(i, j, c)
		}
	}
}
