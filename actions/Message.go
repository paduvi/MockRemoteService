package actions

import "github.com/paduvi/MockRemoteService/mockrepo"
import . "github.com/paduvi/MockRemoteService/models"

func FindMessage(id int, done chan Message) {
	// return empty Message if not found
	done <- mockrepo.FindMessage(id)
}

func ListMessage(done chan Messages) {
	done <- mockrepo.ListMessage()
}

func CreateMessage(t Message, done chan Message) {
	done <- mockrepo.CreateMessage(t)
}

func DestroyMessage(id int, done chan error) {
	done <- mockrepo.DestroyMessage(id)
}
