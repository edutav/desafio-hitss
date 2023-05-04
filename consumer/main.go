package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/streadway/amqp"
)

type Clients struct {
	Nome           string `json:"nome"`
	Sobrenome      string `json:"sobrenome"`
	Contato        string `json:"contato"`
	Endereco       string `json:"endereco"`
	DataNascimento string `json:"data_nascimento"`
	CPF            string `json:"cpf"`
}

func main() {
	fmt.Println("RabbitMQ in Golang")

	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	fmt.Println("Successfully connected to RabbitMQ instance")

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	msgs, err := channel.Consume(
		"cliente_queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			client := Clients{}
			err := json.Unmarshal(msg.Body, &client)
			if err != nil {
				log.Printf("Failed to decode message body: %v", err)
				continue
			}
			if err := CreateClient(&client); err != nil {
				log.Printf("Failed to save client: %v", err)
				continue
			}
		}
	}()

	log.Println("Waiting for messages...")
	<-forever
}

func CreateClient(client *Clients) error {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO clientes(nome, sobrenome, contato, endereco, data_nascimento, cpf) VALUES($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(client.Nome, client.Sobrenome, client.Contato, client.Endereco, client.DataNascimento, client.CPF)
	if err != nil {
		return err
	}

	return nil
}
