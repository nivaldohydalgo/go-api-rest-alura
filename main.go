package main

import (
	"fmt"

	"github.com/nivaldohydalgo/go-api-rest/database"
	"github.com/nivaldohydalgo/go-api-rest/models"
	"github.com/nivaldohydalgo/go-api-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "História 1"},
		{Id: 2, Nome: "Nome 2", Historia: "História 2"},
	}

	database.ConectaComBancoDeDados()

	fmt.Println("Inicializando o servidor REST com Go!")
	routes.HandleRequest()

}
