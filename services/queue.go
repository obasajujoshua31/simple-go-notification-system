package services



import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"simple-go-notification-system/config"
)

const (
	emailConfirmation = "email-confirmation"
	duration = true
	autoDelete = false
	exclusive = false
	nowait = false
	preFetchCount = 1
	preFetchSize = 0
	global = false
	noLocal = false
	autoAck = false
	consumer = ""
	)

type Queue struct {}

func CreateConnection(config config.AppConfig) (*amqp.Connection, error) {
	conn, err := amqp.Dial(config.RabbitMQURL)

	if err != nil {
		return nil, err
	}
	return  conn, nil
}


func CreateChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	channel, err := conn.Channel()

	if err != nil {
		return nil, err
	}

	return channel, nil

}


func CreateQueue (channel *amqp.Channel) (amqp.Queue,  error) {
	queue, err := channel.QueueDeclare(emailConfirmation, duration, autoDelete, exclusive, nowait, nil )

	if err != nil {
		return amqp.Queue{}, err
	}

	err = channel.Qos(preFetchCount, preFetchSize, global)

	if err != nil {
		return amqp.Queue{}, err
	}

	return queue, nil
}

func ReadQueue(channel *amqp.Channel, queue amqp.Queue) (<-chan amqp.Delivery, error) {
	msgChannel, err := channel.Consume(queue.Name, consumer, autoAck, exclusive, noLocal, nowait, nil)
	if err != nil {
		return nil, err
	}

	return msgChannel, nil
}



func ConsumeMessages(delivery <-chan amqp.Delivery, client MailClient) error {

	log.Printf("Consumer ready, PID: %d", os.Getpid() )

	for d := range delivery {
		mailMessage := &Message{}

		err := json.Unmarshal(d.Body, mailMessage)

		if err != nil {
			return err
		}


		if err := d.Ack(false); err != nil {
			return errors.New("Unable to acknowledge message " + err.Error())
		}

		fmt.Printf("Message Acknowledged successfully :%s", mailMessage.ToName)

		message := NewEmailMessage(mailMessage)

		err = client.SendEmailMessage(message)
		if err != nil {
			return err
		}

		fmt.Printf("Message Sent successfully to: %s", mailMessage.ToName)
	}
	return nil
}