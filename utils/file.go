package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// ReadTemplate 读取文件并进行填充,返回填充的字节
func ReadTemplate(filePath string, availableData any) ([]byte, error) {
	tmpl, err := template.New(filepath.Base(filePath)).
		// Funcs(availableFunctions).
		ParseFiles(filePath)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, availableData); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// ReadIndexfile 按照行读取文件,返回读取完毕的切片
func ReadIndexfile(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("读取的项目索引文件不存在: %s", err)
	}
	var projectList = make([]string, 0)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		projectList = append(projectList, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return projectList, nil
}

// IsExist 判断文件或文件夹是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}

func WriterFile(path, filename, data string) error {
	filePath := path + filename + ".yaml"
	if IsExist(filePath) {
		return fmt.Errorf("文件已存在 %s", filePath)
	}
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	if _, err := write.WriteString(data); err != nil {
		return fmt.Errorf("文件 %s 写入失败：%s", filePath, err)
	}

	//Flush将缓存的文件真正写入到文件中
	write.Flush()
	return nil
}
