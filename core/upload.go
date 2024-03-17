package core

import (
	"encoding/json"
	"rabbitmq/mq"
)

// 删除更换文件到OSS
func upload() {
	ossPath := "oss/"
	data := mq.TransferData{
		FileHash:      fileMeta.FileShal,
		CurLocation:   fileMeta.Location,
		DestLocation:  ossPath,
		DestStoreType: cmn.StoreOSS,
	}
	pubData, _ := json.Marshal(data)
	suc := mq.Publish(cfg.TransExchangeName, cfg.TransOSSRoutingKey, pubData)
	if !suc {
		//TODD:加入重拾发送消息逻辑
	}
}
