package mq

// TransferData 转移对列中消息载体的结构格式
type TransferData struct {
	FileHash      string //文件hash
	CurLocation   string // 存在缓存里的文件地址
	DestLocation  string // 要转移的目标地址
	DestStoreType string // 将要被转移到哪种类型的存储里面
}
