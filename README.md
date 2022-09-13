# yaml 配置生成文件

## 代码逻辑：

通过读取 `template `文件夹中的指定文件模板文件，例如ingress的：`ingress-template.yaml` 作为基准进行生成，其中按行读取文件`ingress-index.txt` 中的域名，生成最终的yaml文件到 `yaml/ingress` 目录下。

## 已实现：

- [x] ingress 批量生成
- [ ] confmap 批量生成
- [ ] service 批量生成
- [ ] deployment 批量生成

## 打包

> **注意不要用中文**

```bash
go build -o yaml-generator -ldflags \
"-X github.com/tqtcloud/yamltemplate/version.GIT_TAG='v0.0.1' -X github.com/tqtcloud/yamltemplate/version.BUILD_TIME='2022-09' -X github.com/tqtcloud/yamltemplate/version.GIT_COMMIT='init' -X github.com/tqtcloud/yamltemplate/version.GIT_BRANCH='master' -X github.com/tqtcloud/yamltemplate/version.GO_VERSION='v1.18.3'" main.go
```

### 预期输出

```bash
$ ./yaml-generator   -v
Version   : 'v0.0.1'
Build Time: '2022-09'
Git Branch: 'master'
Git Commit: 'init'
Go Version: 'v1.18.3'
```

## 使用

```bash
$ ./yaml-generator   -h 
k8s yaml 模板化生成工具

Usage:
  yaml-template [flags]
  yaml-template [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  ingress     ingress 生成模块

Flags:
  -h, --help      help for yaml-template
  -v, --version   the cmdb version

Use "yaml-template [command] --help" for more information about a command.
```

### ingress生成

```bash
$ ./yaml-generator   ingress -h
ingress 生成模块

Usage:
  yaml-template ingress [flags]

Examples:

        go run main.go    ingress -n prod-lebei-html  -t 36bike -b lebei-html-qrcode-analysis
        或
        yaml-generator   ingress -n prod-lebei-html  -t 36bike -b lebei-html-qrcode-analysis

Flags:
  -b, --backendServiceName string   后端服务名
  -h, --help                        help for ingress
  -p, --httpPath string             ingress  path路径 (default "/")
  -c, --ingressClass string         ingress Class (default "ack-nginx-lebei")
  -n, --namespace string            名称空间
  -s, --servicePort string          后端服务端口 (default "80")
  -t, --tlsSecretName string        https 证书的Secret Name (default "36bike")
  -u, --url string                  ingress url 地址,默认读取文件：template/ingress-index.txt

Global Flags:
  -v, --version   the cmdb version
```

运行后预期输出

```bash
$ ./yaml-generator   ingress -n prod-lebei-html  -t 36bike -b lebei-html-qrcode-analysis 
存在索引：[abmcx-qrcode.36bike.com bdcx-qrcode.36bike.com blcx-qrcode.36bike.com] 
```

如果存在问题报错

```bash
$ ./yaml-generator   ingress -n prod-lebei-html  -t 36bike -b lebei-html-qrcode-analysis 
存在索引：[abmcx-qrcode.36bike.com bdcx-qrcode.36bike.com blcx-qrcode.36bike.com] 
[2022-09-10 17:56:41] [info] [ingress.go:53 github.com/tqtcloud/yamltemplate/cmd.glob..func1] 文件生成失败; ERROR: 文件已存在 yaml/ingressabmcx-qrcode.36bike.com.yaml 
[2022-09-10 17:56:41] [info] [ingress.go:53 github.com/tqtcloud/yamltemplate/cmd.glob..func1] 文件生成失败; ERROR: 文件已存在 yaml/ingressbdcx-qrcode.36bike.com.yaml
[2022-09-10 17:56:41] [info] [ingress.go:53 github.com/tqtcloud/yamltemplate/cmd.glob..func1] 文件生成失败; ERROR: 文件已存在 yaml/ingressblcx-qrcode.36bike.com.yaml
```

