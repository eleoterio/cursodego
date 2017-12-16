package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/eleoterio/cursodego/post/model"
)

func main() {
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	usuario := model.Usuario{}
	usuario.ID = 1
	usuario.Nome = "Eleoterio"

	conteudoEnviar, err := json.Marshal(usuario)
	if err != nil {
		fmt.Println("[main]  Erro: ", err.Error())
		return
	}

	request, err := http.NewRequest("POST", "https://requestb.in/1j5s01o1", bytes.NewBuffer(conteudoEnviar))
	if err != nil {
		fmt.Println("[main]  Erro: ", err.Error())
		return
	}

	request.SetBasicAuth("fizz", "buzz")
	request.Header.Set("content-type", "application/json; charset=utf-8")
	resposta, err := cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro: ", err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro: ", err.Error())
			return
		}
		fmt.Println(" ")
		fmt.Println(string(corpo))
	}
}