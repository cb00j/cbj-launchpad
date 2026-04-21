package utils

import (
	"crypto/md5"
	"errors"
	"fmt"

	"io"
	"math/rand"
	"os"
	"path"

	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"gopkg.in/ini.v1"
)

// 时间戳转换成日期
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// 日期转换成时间戳 2020-05-02 15:04:05
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// 获取时间戳
func GetUnix() int64 {
	return time.Now().Unix()
}

// GetUnixNano 获取时间戳,单位纳秒
func GetUnixNano() int64 {
	return time.Now().UnixNano()
}

// 获取当前的日期
func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

// 获取年月日
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

func MD5(str string) string {
	h := md5.New()
	_, err := io.WriteString(h, str)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

// 表示把string转换成int
func Int(str string) (int, error) {
	n, err := strconv.Atoi(str)
	return n, err
}

// Float 表示把string转换成float64
func Float(str string) (float64, error) {
	n, err := strconv.ParseFloat(str, 64)
	return n, err
}

func Mul(price float64, num int) float64 {
	return price * float64(num)
}

// 表示把int转换成string
func String(n int) string {
	str := strconv.Itoa(n)
	return str
}

func UploadFile(c *gin.Context, name string) (dest string, err error) {
	file, err := c.FormFile(name)
	if err != nil {
		return "", err
	}
	// 获取文件后缀
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".gif":  true,
	}

	if !allowExtMap[extName] {
		return "", errors.New("文件后缀名不合法")
	}

	// 创建图片保存目录  e.g. static/upload/20260310
	day := GetDay()
	dir := "./static/upload/" + day
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}

	// 生成文件名称和文件保存的目录
	fileName := strconv.FormatInt(GetUnixNano(), 10) + extName
	dest = path.Join(dir, fileName)

	err = c.SaveUploadedFile(file, dest)

	if err != nil {
		return "", err
	}

	return dest, nil
}

// 获取Oss的状态
func GetOssStatus() int {
	config, iniErr := ini.Load("./conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}
	ossStatus, _ := Int(config.Section("oss").Key("status").String())
	return ossStatus
}

func Sub(a int, b int) int {
	return a - b
}

// Substr截取字符串
func Substr(str string, start int, end int) string {
	rs := []rune(str)
	rl := len(rs)
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = 0
	}

	if end < 0 {
		end = rl
	}
	if end > rl {
		end = rl
	}
	if start > end {
		start, end = end, start
	}

	return string(rs[start:end])
}

//生成随机数

func GetRandomNum() string {
	var str string

	for i := 0; i < 4; i++ {
		current := rand.Intn(10)

		str += String(current)
	}
	return str
}
