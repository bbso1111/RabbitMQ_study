package mq

import (
	"github.com/streadway/amqp"
	"log"
	"rabbitmq/config"
)

var conn *amqp.Connection //mq的连接对象
var channel *amqp.Channel //进行消息的发布和接收

// initChannel 初始化一个通道
func initChannel() bool {
	//1.判断channel是否已经创建
	if channel != nil {
		return true
	}
	//2.获得rabbitmq的一个连接
	conn, err := amqp.Dial(config.RabbitURL)
	if err != nil {
		log.Panicln(err.Error())
		return false
	}
	//3.打开一个channel，用于消息的发布和接收等
	channel, err = conn.Channel()
	if err != nil {
		log.Panicln(err.Error())
		return false
	}

	return true
}

// Publish 发送消息
func Publish(exchange, routingKey string, msg []byte) bool {
	//1.判断channel是否正常
	if !initChannel() {
		return false
	}
	//2.执行消息发布动作
	err := channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		},
	)
	if err != nil {
		log.Panicln(err.Error())
		return false
	}
	return true
}
