package taskmq

import (
	"encoding/json"

	"github.com/go-irain/logger"

	"github.com/streadway/amqp"
)

//AmqpConfig 消息队列链接配置
type AmqpConfig struct {
	Addr, ExchangeName, RoutingKey, ExchangeType string
}

// AmqpCollector
var AmqpCollector AmqpConfig
var publishChan = make(chan string)

// wangxiaobo@parkingwang.com
func initSendMQ(amqpCnf AmqpConfig, publishChan chan string) error {
	connection, err := amqp.Dial(amqpCnf.Addr)
	if err != nil {
		logger.Error("Dial error: ", err)
		return err
	}
	channel, err := connection.Channel()
	if err != nil {
		logger.Error("Channel error: ", err)
		return err
	}
	if err := channel.ExchangeDeclare(
		amqpCnf.ExchangeName, // name
		amqpCnf.ExchangeType, // type
		true,                 // durable
		false,                // auto-deleted
		false,                // internal
		false,                // noWait
		nil,                  // arguments
	); err != nil {
		logger.Error("Exchange Declare error : ", err)
		return err
	}
	go publishHandler(amqpCnf, channel, connection, publishChan)
	return nil
}

// wangxiaobo@parkingwang.com
func reconnect(amqpCnf AmqpConfig, channel **amqp.Channel, connection **amqp.Connection) {
	var err error
	*connection, err = amqp.Dial(amqpCnf.Addr)
	if err != nil {
		logger.Error("Dial error: ", err)
		return
	}
	*channel, err = (*connection).Channel()
	if err != nil {
		logger.Error("Channel error: ", err)
		return
	}
	if err := (*channel).ExchangeDeclare(
		amqpCnf.ExchangeName, // name
		amqpCnf.ExchangeType, // type
		true,                 // durable
		false,                // auto-deleted
		false,                // internal
		false,                // noWait
		nil,                  // arguments
	); err != nil {
		logger.Error("Exchange Declare error : ", err.Error())
		return
	}
}

// wangxiaobo@parkingwang.com
func publishHandler(amqpCnf AmqpConfig, channel *amqp.Channel, connection *amqp.Connection, publishChan chan string) {
	for {
		data := <-publishChan
		if data != "" {
			err := channel.Publish(
				amqpCnf.ExchangeName, // publish to an exchange
				amqpCnf.RoutingKey,   // routing to 0 or more queues
				false,                // mandatory
				false,                // immediate
				amqp.Publishing{
					Headers:         amqp.Table{},
					ContentType:     "text/plain",
					ContentEncoding: "",
					Body:            []byte(data),
					DeliveryMode:    amqp.Persistent, // 1=non-persistent, 2=persistent
					Priority:        0,               // 0-9
					// a bunch of application/implementation-specific fields
				},
			)
			if err != nil {
				logger.Error("Exchange Publish error: ", err.Error())
				channel.Close()
				connection.Close()
				reconnect(amqpCnf, &channel, &connection)
				continue
			}
		}
	}
}

/* ---------------------------------------------------------- */

// InitMQ 初始化发送消息队列
// wangxiaobo@parkingwang.com
func InitMQ(cfg map[string]string) error {
	AmqpCollector = AmqpConfig{
		Addr:         cfg["addr"],
		ExchangeName: cfg["exchange_name"],
		RoutingKey:   cfg["routing_key"],
		ExchangeType: cfg["exchange_type"],
	}

	err := initSendMQ(AmqpCollector, publishChan)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

// Publish 发送消息
// wangxiaobo@parkingwang.com
func Publish(msgMap map[string]interface{}) {
	data, err := json.Marshal(msgMap)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Info("publish bill msg := ", string(data))
	publishChan <- string(data)
}
