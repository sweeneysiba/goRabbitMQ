package controller

import (
	"encoding/json"
	"flag"
	"fmt"
	"goRabbitMQ/models"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var (
	amqpURL = flag.String("amqp", "amqp://guest:guest@127.0.0.1:5672/", "amqp uri")
)

//PublishOffers ...
func PublishOffers(c *gin.Context) {
	w, r := c.Writer, c.Request

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to connect to rabbitmq server"})
		return
	}
	logrus.WithFields(logrus.Fields{
		"amqp": "amqp://guest:guest@127.0.0.1:5672/",
	}).Info("RabbitMQ connect")

	defer conn.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request | please check the request body"})
		return
	}
	var inputCheck models.OffersInput
	var inputCheckmapForLogs map[string]interface{}

	err = json.Unmarshal(body, &inputCheck)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request | please check the request body"})
		return
	}

	logrus.WithFields(inputCheckmapForLogs).Info("RabbitMQ will publish data to queue")

	ch, err := conn.Channel()
	fmt.Println("ch---------------------", ch)
	fmt.Println("err---------------------", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to create a channel"})
		return
	}
	defer ch.Close()

	queueName := "offer_reciver"

	q, err := ch.QueueDeclarePassive(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logrus.WithFields(logrus.Fields{
		"Queue":      queueName,
		"Passive":    false,
		"Durable":    false,
		"AutoDelete": false,
		"Exclusive":  false,
		"NoWait":     false,
		"Arguments":  nil,
	}).Info("RabbitMQ Queue delcared with arguments")

	// fmt.Println("Queuing Mail Scan : ", string(body))
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logrus.WithFields(logrus.Fields{
		"exchange":    "",
		"routing key": q.Name,
		"mandatory":   false,
		"immediate":   false,
		"ContentType": "application/json",
		"Body":        []byte(body),
	}).Info("RabbitMQ Queue Published")

	c.JSON(http.StatusOK, gin.H{"message": "published the message through rabbitmq"})
	return
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

//QueueConsumer ...
func QueueConsumer() {
	conn, err := amqp.Dial(*amqpURL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	queueName := "offer_reciver"

	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var request models.OffersInput
			json.Unmarshal([]byte(d.Body), &request)
			var requestMap map[string]interface{}
			json.Unmarshal([]byte(d.Body), &requestMap)
			fmt.Println("============= NEW  Request ===========")
			logrus.WithFields(requestMap).Info("RabbitMQ Queue QueueConsumer with request")
			// PrettyPrint(request)
			d.Ack(true)
			for _, reqqq := range request.Offers {
				err = models.StoreHotel(reqqq.Hotel)
				if err != nil {
					panic(err)
				}
				err = models.StoreRoom(reqqq.Room)
				if err != nil {
					panic(err)
				}
				err = models.StoreRatePlan(reqqq.RatePlan)
				if err != nil {
					panic(err)
				}
			}

			d.Ack(true)
		}
	}()
	log.Printf(" [*] Waiting for Scan Request logs. To exit press CTRL+C")
	<-forever
}

//PrettyPrint ...
func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}
