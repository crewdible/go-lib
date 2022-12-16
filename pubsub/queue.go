package pubsub

import (
	"fmt"
	"regexp"
)

func GenerateMultipleQueues(qName string, qTot int) []string {
	queues := []string{}
	for i := 0; i < qTot; i++ {
		queue := fmt.Sprintf("%s-%d", qName, i)
		queues = append(queues, queue)
	}

	return queues

}

func MultipleQueuesCheck(baseName, qName string) (bool, error) {
	good := fmt.Sprintf("^%s-[0-9]*$", baseName)
	return regexp.MatchString(good, qName)
}
