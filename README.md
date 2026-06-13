# wechat-qy
> 易于使用的企业微信（原微信企业号）通用 SDK (Golang)

[![Go Reference](https://pkg.go.dev/badge/github.com/shengbox/wechat-qy.svg)](https://pkg.go.dev/github.com/shengbox/wechat-qy)
[![Go Report Card](https://goreportcard.com/badge/github.com/shengbox/wechat-qy)](https://goreportcard.com/report/github.com/shengbox/wechat-qy)

## 特性

* **高并发安全**：Token 管理机制内置读写锁与双重校验（Double-Checked Locking），原生支持高并发后端服务。
* **第三方应用支持**：完整支持第三方应用套件接口，支持基于应用套件级别的调用和企业号单独调用。
* **独立的服务商管理**：支持独立的 `provider_access_token` 获取与自动续期，完美支持 License 账号许可等服务商 API。
* **连接复用与超时安全**：内置合理的 HTTP 超时控制，统一 Resty 连接池复用，杜绝 TCP 连接泄露（TIME_WAIT 堆积）。
* **可插拔日志系统**：支持注入自定义结构化日志组件（如 Zap、Logrus、Slog 等）。
* **Access Token 自动续期**：Token 超期或失效导致接口调用错误时，自动刷新并重试一次当前调用的 API。
* **加解密支持**：提供被动接收消息（事件）的安全解密解析方法，以及生成被动响应消息的方法。

## 安装

```bash
$ go get -u github.com/shengbox/wechat-qy
```

## 使用手册

### 1. 企业自建应用开发（企业自用模式）

适用于企业自己开发内部应用：

```go
package main

import (
	"fmt"
	"time"

	"github.com/shengbox/wechat-qy/api"
	"github.com/shengbox/wechat-qy/base"
)

func main() {
	// 初始化 API 实例
	wechatAPI := api.New("CORP_ID", "CORP_SECRET", "TOKEN", "ENCODING_AES_KEY")

	// [可选] 配置 HTTP 超时时间（默认 10 秒）
	wechatAPI.Client.SetTimeout(5 * time.Second)

	// [可选] 配置自定义的全局日志（需实现 base.Logger 接口）
	// base.SetLogger(myCustomLogger)

	// 创建用户
	err := wechatAPI.CreateUser(&api.User{
		UserID:        "zhangsan",
		Name:          "张三",
		DepartmentIds: []int64{1, 2},
		Mobile:        "13800000000",
	})
	if err != nil {
		fmt.Printf("创建用户失败: %v\n", err)
	}

	// 创建被动消息解析器（用于处理回调事件）
	recvMsgHandler := wechatAPI.NewRecvMsgHandler()
	
	// 解析回调模式接收的消息/事件
	// data, err := recvMsgHandler.Parse(body, signature, timestamp, nonce)
}
```

### 2. 第三方应用开发（服务商模式）

适用于为其他企业提供 SaaS 应用的第三方服务商：

```go
package main

import (
	"fmt"

	"github.com/shengbox/wechat-qy/suite"
)

func main() {
	// 初始化 Suite 实例
	wechatSuite := suite.New("SUITE_ID", "SUITE_SECRET", "SUITE_TOKEN", "SUITE_ENCODING_AES_KEY")

	// [重要] 如果需要调用 License 账号许可、注册定制化等服务商 API，请设置服务商凭证信息
	wechatSuite.SetProvider("PROVIDER_CORPID", "PROVIDER_SECRET")

	// 获取企业的账号许可激活详情 (需要 provider_access_token)
	info, err := wechatSuite.GetActiveInfoByUser("AUTH_CORPID", "USER_ID")
	if err != nil {
		fmt.Printf("获取许可详情失败: %v\n", err)
	} else {
		fmt.Printf("激活码: %s\n", info.ActiveCode)
	}

	// 获取企业号的永久授权码
	authInfo, err := wechatSuite.GetPermanentCode("AUTH_CODE")

	// 创建基于特定授权企业的 API 实例 (继承套件的 Token 管理机制)
	apiInstance := wechatSuite.NewAPI("AUTH_CORPID", authInfo.PermanentCode)
	
	// 后续调用与自建应用 API 实例一致
	// apiInstance.CreateUser(...)
}
```

## 贡献与开发

### 运行单元测试
项目已针对高并发 Token 刷新、加解密安全性、接口反序列化崩溃等核心场景编写了完整的单元测试，可以使用 Race 检测器运行：

```bash
$ go test -race -v ./...
```

## License
MIT
