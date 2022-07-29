package pubsub

import (
	"fmt"

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
		// Headers:       amqp.Table{"type": m.Body.Type}, // NOT USED YET
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
