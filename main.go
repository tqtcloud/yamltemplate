package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/tqtcloud/yamltemplate/cmd"
	"os"
	"path/filepath"
	"text/template"
)

var projectList = make([]string, 0)

// 读取文件并进行填充
func Read(filePath string) ([]byte, error) {
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

var availableData = map[string]string{
	"Namespace":          "prod-lebei-html",
	"IngressClass":       "ack-nginx-lebei",
	"Url":                "lccx-qrcode.36bike.com",
	"TlsSecretName":      "36bike",
	"TttpPath":           "/",
	"BackendServiceName": "lebei-html-qrcode-analysis",
	"ServicePort":        "80",
}

// 模板函数
// var availableFunctions = template.FuncMap{
// 	"GeneratePassword": GeneratePasswordFunc,
// }

// func GeneratePasswordFunc() (string, error) {
// 	return "123", nil
// }

func ReadIndexfile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("读取的项目索引文件不存在: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		projectList = append(projectList, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
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

func main() {
	cmd.Execute()
	//if err := ReadIndexfile("./template/ingress-index.txt"); err != nil {
	//	log.Println(err)
	//	os.Exit(-1)
	//}
	//fmt.Printf("存在索引：%s \n", projectList)
	//for _, value := range projectList {
	//	availableData["Url"] = value
	//	byte, err := Read("template/ingress-template.yaml")
	//	if err != nil {
	//		log.Printf("模板生成失败：%s \n", err)
	//	}
	//	if err := WriterFile("yaml/", value, string(byte)); err != nil {
	//		log.Printf("文件生成失败; ERROR: %s \n", err)
	//		continue
	//	}
	//}

}
