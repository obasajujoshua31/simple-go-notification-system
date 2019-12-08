package server

import (
	"simple-go-notification-system/config"
	"simple-go-notification-system/services"
)

func StartServer() error {
	appConfig, err := config.LoadEnv()

	if err != nil {
		return err
	}

	client := services.NewMailClient(appConfig)


	conn, err := services.CreateConnection(appConfig)

	if err != nil {
		return  err
	}

	defer conn.Close()

	channel, err := services.CreateChannel(conn)

	if err != nil {
		return  err
	}

	defer channel.Close()

	queue, err := services.CreateQueue(channel)

	if err != nil {
		return  err
	}

	delivery, err := services.ReadQueue(channel, queue)
	if err != nil {
		return err
	}

	stopChan := make(chan bool)

	go services.ConsumeMessages(delivery, client)

	<-stopChan





	if err != nil {
		return err
	}

	return nil
}
