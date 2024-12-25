package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

var BeijingLoc = time.FixedZone("Beijing Time", 8*60*60)
var FilePath = ""

// Config func to get env value
func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env.sample")
	if err != nil {
		err = godotenv.Load("../.env.sample")
		if err != nil {
			err = godotenv.Load("../../.env.sample")
			if err != nil {
				err = godotenv.Load("./.env.sample")
			}
		}
	}

	return os.Getenv(key)
}
func ConfigPKS(key string) string {
	// load .env file
	err := godotenv.Load("output.txt")
	if err != nil {
		err = godotenv.Load("../config/output.txt")
		if err != nil {
			err = godotenv.Load("../../config/output.txt")
			if err != nil {
				err = godotenv.Load("./config/output.txt")
			}
		}
	}

	return os.Getenv(key)
}
func SetPKSConfig(path string, oValue string, aValue string) error {
	file, err := os.OpenFile(path+"/output.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("打开文件通道出错:", err)
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件通道出错:", err)
		}
	}(file) // 确保在函数结束时关闭文件

	// 写入数据到文件
	_, err = file.WriteString("OPK_SECRET=" + oValue + "\nAPK_SECRET=" + aValue)
	if err != nil {
		fmt.Println("写入数据出错:", err)
		return err
	}
	fmt.Println("写入成功。")
	return nil
}
func EditPKSConfig(path string, key string, value string) error {
	filePath := path + "/output.txt" // 文件路径
	lineNumber := 0
	newContent := key + "=" + value // 新的行内容
	if key == "OPK_SECRET" {
		lineNumber = 1
	} else if key == "APK_SECRET" {
		lineNumber = 2
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("打开文件通道出错:", err)
		return err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Println("关闭文件通道出错:", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var lines []string
	i := 1

	for scanner.Scan() {
		if i == lineNumber {
			lines = append(lines, newContent)
		} else {
			lines = append(lines, scanner.Text())
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("扫描原文件出错:", err)
		return err
	}

	output, err := os.Create(filePath)
	if err != nil {
		fmt.Println("创建写入通道出错:", err)
		return err
	}
	defer func(output *os.File) {
		err = output.Close()
		if err != nil {
			fmt.Println("关闭写入通道出错:", err)
		}
	}(output) // 确保在函数结束时关闭文件

	writer := bufio.NewWriter(output)
	for _, line := range lines {
		_, err := fmt.Fprintln(writer, line)
		if err != nil {
			return err
		}
	}

	err = writer.Flush()
	if err != nil {
		return err
	}
	fmt.Println("写入成功。")
	return nil
}
func GetAbsolutePath() {
	// 获取当前工作目录
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("获取绝对路径失败" + err.Error())
	}
	projectName := "resource-go-0508"
	fs := strings.Split(wd, projectName)
	FilePath = filepath.Join(fs[0], projectName, "/config/")

}
