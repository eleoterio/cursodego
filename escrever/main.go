package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/eleoterio/cursodego/escrever/model"
)

func main() {

	arquivo, err := os.Open("city.csv")
	if err != nil {
		fmt.Println("[main]  Erro: ", err.Error())
		return
	}

	leitorCsv := csv.NewReader(arquivo)

	conteudo, err := leitorCsv.ReadAll()
	if err != nil {
		fmt.Println("[main] Erro: ", err.Error())
		return
	}

	arquivoJSON, err := os.Create("city.json")
	if err != nil {
		fmt.Println("[main] Erro: ", err.Error())
		return
	}

	escritor := bufio.NewWriter(arquivoJSON)
	escritor.WriteString("[\r\n")
	for _, linha := range conteudo {

		for indiceItem, item := range linha {
			dados := strings.Split(item, "/")
			cidade := model.Cidade{}
			cidade.Nome = dados[0]
			cidade.Estado = dados[1]

			fmt.Printf("Cidade: %+v\r\n", cidade)
			cidadeJSON, err := json.Marshal(cidade)

			if err != nil {
				fmt.Println("[main] erro item ", item, ". Erro: ", err.Error())
			}

			escritor.WriteString("  " + string(cidadeJSON))

			if (indiceItem + 1) < len(linha) {
				escritor.WriteString(",\r\n")
			}
		}
	}

	escritor.WriteString("\r\n]")
	escritor.Flush()
	arquivoJSON.Close()
	arquivo.Close()
}
