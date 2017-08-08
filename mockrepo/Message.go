package mockrepo

import (
	"fmt"
	. "github.com/paduvi/MockRemoteService/models"
	"time"
)

var currentId int

var messages Messages

// Give us some seed data
func init() {
	go CreateMessage(Message{
		Title:   "Write presentation",
		Content: "How to use Goroutine in Golang",
	})
	go CreateMessage(Message{
		Title:   "Host meetup",
		Content: "I want to invite you to come our Golang Meetup",
	})
}

func FindMessage(id int) Message {
	for _, t := range messages {
		if t.Id == id {
			return t
		}
	}
	// return empty Message if not found
	return Message{}
}

func ListMessage() Messages {
	return messages
}

func CreateMessage(t Message) Message {
	currentId += 1
	t.Id = currentId
	t.CreatedAt = time.Now()
	messages = append(messages, t)
	return t
}

func DestroyMessage(id int) error {
	for i, t := range messages {
		if t.Id == id {
			messages = append(messages[:i], messages[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Message with id of %d to delete", id)
}
