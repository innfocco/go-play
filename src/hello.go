package main

import (
	"encoding/json"
	"fmt"
	"model"
)

func main() {

	messages := []*model.Message{
		&model.Message{"Julio", "Olá tudo bem?", false},
		&model.Message{"Taise", "Olá tudo bem?", false},
		&model.Message{"Emily", "Olá tudo bem?", false},
		&model.Message{"João", "Olá tudo bem?", false},
		&model.Message{"Antônio", "Olá tudo bem?", false},
	}

	for _, msg := range messages {
		markAsReaded(msg)
	}
	fmt.Printf("\n\n\n")
	for _, msg := range messages {
		data := printMessageAsJson(msg)
		var newMsg model.Message
		jsonParseMessage(data, &newMsg)
		fmt.Println(newMsg.From)
	}

}

func printMessageAsJson(msg *model.Message) string {
	e, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		return "ERROR"
	}
	ret := string(e)
	fmt.Println(ret)
	return ret
}

func jsonParseMessage(data string, msg *model.Message) {
	jsonData := []byte(data)
	err := json.Unmarshal(jsonData, msg)
	if err != nil {
		fmt.Println(err)
	}
}

func markAsReaded(r *model.Message) {
	r.Readed = true
}
