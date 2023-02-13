package pubsub

import (
	"fmt"

	"github.com/crewdible/go-lib/stringlib"
	amqp "github.com/rabbitmq/amqp091-go"
)

//Publish publishes a request to the amqp queue
func (c *Connection) Publish(m Message) error {
	select { //non blocking channel - if there is no error will go to default where we do nothing
	case err := <-c.err:
		if err != nil {
			c.Reconnect()
		}
	default:
	}

	p := amqp.Publishing{
		Headers:     m.Headers,
		ContentType: m.ContentType,
		// CorrelationId: m.CorrelationID, // NOT USED YET
		Body: m.Body.Data,
		// ReplyTo:       m.ReplyTo, // NOT USED YET
	}
	if err := c.channel.Publish(c.exchange, m.Queue, false, false, p); err != nil {
		return fmt.Errorf("Error in Publishing: %s", err)
	}
	return nil
}

func (c *Connection) CrewBasicPublish(qname string, msg interface{}, retry bool) error {
	var rtry int
	if retry {
		rtry = 0
	} else {
		rtry = 100
	}

	msgStr, err := stringlib.StructToJsonString(msg)
	if err != nil {
		return err
	}

	m := Message{
		Headers:     amqp.Table{"x-amqp-delivery-count": int32(rtry)},
		Queue:       qname,
		ContentType: "application/json",
		Body: MessageBody{
			Data: []byte(msgStr),
		},
	}

	err = c.Publish(m)

	return err
}
