package scrap

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

// ManufacturerList json content
type ManufacturerList struct {
	Manufacturers []Manufacturer `json:"Fabricantes"`
}

// Manufacturer json content
type Manufacturer struct {
	Name string `json:"NomeFabricante"`
}

// ProductList json content
type ProductList struct {
	Products []Product `json:"Produtos"`
	Total    int       `json:"TotalProdutos"`
}

// ProductDetails json content
type ProductDetails struct {
	Product Product `json:"Produto"`
}

// Product json content
type Product struct {
	ID           string `json:"IdProduto"`
	EAN          string `json:"EAN"`
	Name         string `json:"NomeProduto"`
	Manufacturer string `json:"NomeFabricante"`
	Category     string `json:"NomeCategoria"`
	Description  string `json:"DescricaoCurta"`
	Picture      string `json:"FotoPrincipal"`
}

// CategoryList json content
type CategoryList struct {
	Categories []Category `json:"Categorias"`
}

// Category json content
type Category struct {
	ID   string `json:"IdCategoria"`
	Name string `json:"NomeCategoria"`
	URL  string `json:"NomeCategoriaURL"`
}

// Start scrap
func Start() {
	client := GetRestClient()
	productListURL := "Produto/GetProdutosCategoria?NomeCategoriaURL=%s&PaginaAtual=%d&TamanhoPagina=33"
	productDetailsURL := "Produto/GetProduto?IdProduto=%s"
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

	csvFile, err := os.OpenFile("products.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()
	csvWriter := csv.NewWriter(csvFile)

	categoryList := getCategories()
	for _, category := range categoryList.Categories {
		fmt.Println("Processando categoria: " + category.Name)

		ids := []string{}
		productList := ProductList{}
		page := 1

		_, err := rest.SetResult(&).Get(fmt.Sprintf(productListURL, category.Name, 1))
		for p := 1; p <= pages; p++ {

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
	}

	csvWriter.WriteAll(record)
}

// GetFabricantes ...
func GetFabricantes() {
	list := ManufacturerList{}
	url := "/Fabricante/GetFabricantes"
	client := GetRestClient()

	csvFile, err := os.OpenFile("fabricantes.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	records := [][]string{}

	_, err = client.SetResult(&list).Get(url)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range list.Manufacturers {
		records = append(records, []string{
			f.Name,
		})
	}

	csvWriter.WriteAll(records)
}

func getCategories() CategoryList {
	list := CategoryList{}
	client := GetRestClient()

	_, err := client.SetResult(&list).Get("Categoria/GetCategorias")
	if err != nil {
		log.Fatal(err)
	}

	return list
}

func downloadFile(filepath string, url string) error {

    // Get the data
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Create the file
    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    // Write the body to file
    _, err = io.Copy(out, resp.Body)
    return err
}