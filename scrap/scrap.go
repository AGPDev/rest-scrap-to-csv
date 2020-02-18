package scrap

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

// Fabricantes json content
type Fabricantes struct {
	Fabricantes []Fabricante `json:"Fabricantes"`
}

// Fabricante json content
type Fabricante struct {
	Nome string `json:"NomeFabricante"`
}

// ProductList json content
type ProductList struct {
	Produtos []Produto `json:"Produtos"`
}

// ProductDetails json content
type ProductDetails struct {
	Produto Produto `json:"Produto"`
}

// Produto json content
type Produto struct {
	ID            string `json:"IdProduto"`
	EAN           string `json:"EAN"`
	Nome          string `json:"NomeProduto"`
	Marca         string `json:"NomeFabricante"`
	Categoria     string `json:"NomeCategoria"`
	Descricao     string `json:"DescricaoCurta"`
	FotoPrincipal string `json:"FotoPrincipal"`
}

var alf = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC",
	"AD", "AE", "AF", "AG", "AH", "AI", "AJ", "AK", "AL", "AM", "AN", "AO", "AP", "AQ", "AR", "AS", "AT", "AU"}

func getCategories() {

}

// Start scrap
func Start() {
	var list ProductList
	var details ProductDetails

	categoryURL := "/Produto/GetProdutosCategoria?NomeCategoriaURL=%s&PaginaAtual=%d&TamanhoPagina=%d"
	productURL := "/Produto/GetProduto?IdProduto=%s"
	pages := 1
	categoryName := "acessorios"
	ids := []string{}
	rest := GetRestClient()

	csvFile, err := os.OpenFile("products.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	// csvWriter.UseCRLF = true

	record := [][]string{
		{
			"sku",
			"attribute_set_code",
			"product_type",
			"categories",
			"name",
			"price",
			// "product_online",
			// "visibility",
			"additional_attributes",
			"description",
		},
	}

	for p := 1; p <= pages; p++ {
		_, err := rest.SetResult(&list).Get(fmt.Sprintf(categoryURL, categoryName, p, pages))
		if err != nil {
			log.Fatal(err)
		}

		for _, row := range list.Produtos {
			ids = append(ids, row.ID)
		}

		for _, id := range ids {
			_, err := rest.SetResult(&details).Get(fmt.Sprintf(productURL, id))
			if err != nil {
				log.Fatal(err)
			}

			description := strings.ReplaceAll(details.Produto.Descricao, "\n", "<br>")
			record = append(record, []string{
				details.Produto.EAN,
				"Default",
				"simple",
				"Default Category/Acessórios/" + details.Produto.Categoria,
				details.Produto.Nome,
				" ",
				// "1",
				// "4",
				"has_options=0,required_options=0,manufacturer=" + details.Produto.Marca,
				description,
			})
		}
	}

	csvWriter.WriteAll(record)
}

// GetFabricantes ...
func GetFabricantes() {
	var list Fabricantes
	var url = "/Fabricante/GetFabricantes"
	var rest = GetRestClient()

	csvFile, err := os.OpenFile("fabricantes.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	records := [][]string{}

	_, err = rest.SetResult(&list).Get(url)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range list.Fabricantes {
		records = append(records, []string{
			f.Nome,
		})
	}

	csvWriter.WriteAll(records)
}
