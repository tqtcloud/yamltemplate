package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tqtcloud/yamltemplate/utils"
	"os"
)

var (
	namespace          string
	ingressClass       string
	url                string
	tlsSecretName      string
	httpPath           string
	backendServiceName string
	servicePort        string
)

func newavailableData() map[string]string {
	return map[string]string{
		"Namespace":          namespace,
		"IngressClass":       ingressClass,
		"Url":                url,
		"TlsSecretName":      tlsSecretName,
		"TttpPath":           httpPath,
		"BackendServiceName": backendServiceName,
		"ServicePort":        servicePort,
	}
}

var txt = `
	go run main.go    ingress -n prod-lebei-html  -t 36bike -b lebei-html-qrcode-analysis
	或
	yaml-generator   ingress -n prod-lebei-html  -t 36bike -b lebei-html-qrcode-analysis 
`
var IngressCmd = &cobra.Command{
	Use:     "ingress",
	Short:   "ingress 生成模块",
	Long:    "ingress 生成模块",
	Example: txt,
	RunE: func(cmd *cobra.Command, args []string) error {
		availableData := newavailableData()
		ingressList, err := utils.ReadIndexfile("./template/ingress-index.txt")
		if err != nil {
			log.Println(err)
			os.Exit(-1)
		}
		fmt.Printf("存在索引：%s \n", ingressList)
		for _, value := range ingressList {
			availableData["Url"] = value
			byte, err := utils.ReadTemplate("template/ingress-template.yaml", availableData)
			if err != nil {
				log.Printf("模板生成失败：%s ", err)
			}
			if err := utils.WriterFile("yaml/ingress/", value, string(byte)); err != nil {
				log.Printf("文件生成失败; ERROR: %s ", err)
				continue
			}
		}
		return nil
	},
}

func init() {
	IngressCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "名称空间")
	IngressCmd.PersistentFlags().StringVarP(&ingressClass, "ingressClass", "c", "ack-nginx-lebei", "ingress Class")
	IngressCmd.PersistentFlags().StringVarP(&url, "url", "u", "", "ingress url 地址,默认读取文件：template/ingress-index.txt")
	IngressCmd.PersistentFlags().StringVarP(&tlsSecretName, "tlsSecretName", "t", "36bike", "https 证书的Secret Name")
	IngressCmd.PersistentFlags().StringVarP(&httpPath, "httpPath", "p", "/", "ingress  path路径")
	IngressCmd.PersistentFlags().StringVarP(&backendServiceName, "backendServiceName", "b", "", "后端服务名")
	IngressCmd.PersistentFlags().StringVarP(&servicePort, "servicePort", "s", "80", "后端服务端口")
	RootCmd.AddCommand(IngressCmd)
}
