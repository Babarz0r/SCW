package main

import (
	"log"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	servers := []string{"nats://nats.mnq.fr-par.scw.cloud:4222"}

	nc, err := nats.Connect(strings.Join(servers, ","), nats.UserCredentials("admin_crendentials.creds"))

	if err != nil {
		log.Fatalln(err)
	}

	defer nc.Close()

	nc.Subscribe("foo", func(msg *nats.Msg) {
		log.Println("Request received:", string(msg.Data))

		msg.Respond([]byte("Here you go!"))
	})

	reply, err := nc.Request("foo", []byte("Give me data"), 4*time.Second)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Got Reply:", string(reply.Data))
}
