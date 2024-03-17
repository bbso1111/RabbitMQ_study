package mq

import (
	"log"
)

var done chan bool

// StartConsume 开始监听对列，获取消息
func StartConsume(qName, cName string, callback func(msg []byte) bool) {
	//1.通过channel.Consume获得消息信道
	msgs, err := channel.Consume(qName, cName, true, false, false, false, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	//2.循环获取对列的消息
	done = make(chan bool)

	go func() {
		for msg := range msgs {
			//3.调用callback方法来处理新的消息
			procssSuc := callback(msg.Body)
			if !procssSuc {
				//TODD:将任务写道另一个对列，用于异常情况的重试
			}
		}
	}()

	//done没有新的消息过来，则会一直堵塞
	<-done

	//关闭rabbitmq通道
	channel.Close()
}
