package pubsub

import (
	"os"
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"
)

//Consume consumes the messages from the queues and passes it as map of chan of amqp.Delivery
func (c *Connection) Consume() (map[string]<-chan amqp.Delivery, error) {
	m := make(map[string]<-chan amqp.Delivery)
	for _, q := range c.queues {
		deliveries, err := c.channel.Consume(q, "", false, false, false, false, nil)
		if err != nil {
			return nil, err
		}
		m[q] = deliveries
	}
	return m, nil
}

//HandleConsumedDeliveries handles the consumed deliveries from the queues. Should be called only for a consumer connection
func (c *Connection) HandleConsumedDeliveries(q string, delivery <-chan amqp.Delivery, fn func(Connection, string, <-chan amqp.Delivery)) {
	for {
		go fn(*c, q, delivery)
		if err := <-c.err; err != nil {
			c.Reconnect()
			deliveries, err := c.Consume()
			if err != nil {
				panic(err) //raising panic if consume fails even after reconnecting
			}
			delivery = deliveries[q]
		}
	}
}

func (c *Connection) CrewBasicHandleConsumedDeliveries(rtry bool, q string, delivery <-chan amqp.Delivery, router func(string, []byte) error) {
	var (
		pubConn *Connection
		maxRtry = 3
	)
	pubConn = GetConnection(os.Getenv("RMQ_PRODUCER_NAME"))
	if os.Getenv("RMQ_MAX_RETRY") != "" {
		mxrtry, err := strconv.Atoi(os.Getenv("RMQ_MAX_RETRY"))
		if err == nil {
			maxRtry = mxrtry
		}
	}
	for {
		go crewBasicMessageHandler(pubConn, maxRtry, rtry, q, delivery, router)
		if err := <-c.err; err != nil {
			c.Reconnect()
			deliveries, err := c.Consume()
			if err != nil {
				panic(err) //raising panic if consume fails even after reconnecting
			}
			delivery = deliveries[q]
		}
	}
}

func crewBasicMessageHandler(pubConn *Connection, maxRtry int, rtry bool, q string, deliveries <-chan amqp.Delivery, router func(string, []byte) error) {
	var err error
	if rtry {
		for d := range deliveries {
			deliverCount := d.Headers["x-amqp-delivery-count"].(int32)
			m := Message{
				Headers:     amqp.Table{"x-amqp-delivery-count": deliverCount + 1},
				Queue:       q,
				ContentType: d.ContentType,
				Body: MessageBody{
					Data: d.Body,
				},
			}

			err = router(q, d.Body)

			if err != nil && deliverCount >= int32(maxRtry) {
				d.Nack(false, false)
			} else if err != nil {
				pubConn.Publish(m)
				d.Nack(false, false)
			} else {
				d.Ack(false)
			}
		}
	} else {
		for d := range deliveries {
			err = router(q, d.Body)
			if err != nil {
				d.Nack(false, false)
			} else {
				d.Ack(false)
			}
		}
	}

}
