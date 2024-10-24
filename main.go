package main

import (
	"log"  // 导入日志包，用于记录日志
	"os"   // 导入操作系统包，用于文件操作
	"time" // 导入时间包，用于计时

	"qrcode" // 导入二维码模块
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile) // 创建一个日志记录器，输出到标准输出，日志格式包含日期、时间和文件名
var imagePath = "./img/qrcode.jpg"                                      // 图片路径

func main() {
	startAt := time.Now() // 记录程序开始时间

	// 打开二维码图片文件
	fi, err := os.Open(imagePath)
	if !check(err) { // 检查是否有错误
		return
	}
	defer fi.Close() // 在函数结束时关闭文件

	// 解码二维码图片
	qrMatrix, err := qrcode.Decode(fi)
	if !check(err) { // 检查是否有错误
		return
	}

	logger.Println("解码后的内容:", qrMatrix.Content)
	logger.Println("运行时间:", time.Since(startAt))
}

// 检查错误的辅助函数
func check(err error) bool {
	if err != nil {
		logger.Println("报错：", err) // 如果有错误，记录日志
	}
	return err == nil // 返回是否没有错误
}
