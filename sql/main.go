package main

import (
	"fmt"
	"net/http"

	"github.com/eleoterio/cursodego/sql/manipulador"
	"github.com/eleoterio/cursodego/sql/repositorio"
)

func init() {
	//fmt.Println("Vamos começar a subir o servidor...")
}

func main() {

	err := repositorio.AbreConexaoComBancoDeDadosSQL()
	if err != nil {
		fmt.Println("Parando a carga do servidor. Erro ao abrir o banco de dados: ", err.Error())
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá Mundo!")
	})
	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	http.HandleFunc("/local/", manipulador.Local)
	fmt.Println("Estou escutando, comandante...")
	http.ListenAndServe(":8181", nil)
}
