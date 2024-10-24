# QRCode

二维码解析器

# Plan

1. 动态二值化:
2. 提升图片扫描的速度 OK
3. 修复标线取值 OK
4. 容错码纠正数据 OK
5. 数据编码方式
<br/>Numbert
<br/>alphanumeric OK
<br/>8-bit byte OK
<br/>Kanji
6. 识别各角度倾斜的二维码

# 启动

```go
切换阿里代理
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
```

- `go run main.go`
