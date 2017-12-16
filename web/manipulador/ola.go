package manipulador

import (
	"fmt"
	"net/http"
	"time"

	"github.com/eleoterio/cursodego/web/model"
)

//Ola é o manipulador da requisição a rota /ola
func Ola(w http.ResponseWriter, r *http.Request) {
	hora := time.Now().Format("15:04:05")
	pagina := model.Pagina{}
	pagina.Hora = hora
	if err := Modelos.ExecuteTemplate(w, "ola.html", pagina); err != nil {
		http.Error(w, "Erro na renderização.", http.StatusInternalServerError)
		fmt.Println("[Ola] Erro: ", err.Error())
	}
}
