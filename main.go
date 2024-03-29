package main

import (
	"bufio"
	"encoding/json"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"os"
	"rabbitmq/config"
	"rabbitmq/mq"
)

// 转移函数
func main() {
	log.Println("开始监听转移任务队列")
	mq.StartConsume(config.TransExchangeName, "transfer_oss", ProcessTransfer)
}

// ProcessTransfer 处理文件转移的真正逻辑
func ProcessTransfer(msg []byte) bool {
	//1.解析msg
	pubData := mq.TransferData{}
	err := json.Unmarshal(msg, &pubData)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	//2.根据临时存储文件路径，创建文件句柄
	filed, err := os.Open(pubData.CurLocation)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	//3.通过文件句柄将文件内容读出来并且上传道OSS
	err = oss.Bucket{}.PutObject(
		pubData.DestLocation,
		bufio.NewReader(filed))
	if err != nil {
		log.Println(err.Error())
		return false
	}
	//4.更新文件的存储路径到文件表
	suc := dblayer.UpdateFileLocation(
		pubData.FileHash,
		pubData.DestLocation)
	if !suc {
		return false
	}
	return true
}
