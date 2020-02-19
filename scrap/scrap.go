package scrap

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
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
	productDetails := Product{}
	records := [][]string{
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

		productList := ProductList{}
		page := 1

		for {
			fmt.Println("Processando página: " + strconv.Itoa(page))

			_, err := client.
				SetResult(&productList).
				Get(fmt.Sprintf(productListURL, category.URL, page))
			if err != nil {
				log.Fatal(err)
			}

			for _, product := range productList.Products {
				_, err := client.
					SetResult(&productDetails).
					Get(fmt.Sprintf(productDetailsURL, product.ID))
				if err != nil {
					log.Fatal(err)
				}

				categoryName := "Default Category/" + category.Name
				if category.Name != productDetails.Category {
					categoryName += "/" + productDetails.Category
				}

				description := strings.ReplaceAll(productDetails.Description, "\n", "<br>")

				records = append(records, []string{
					productDetails.EAN,
					"Default",
					"simple",
					categoryName,
					productDetails.Name,
					" ",
					// "1",
					// "4",
					"has_options=0,required_options=0,manufacturer=" + productDetails.Manufacturer,
					description,
				})
			}

			if (productList.Total / 33) > page {
				page++
			} else {
				break
			}

			if len(records) >= 5 {
				break
			}
		}

		if len(records) >= 5 {
			break
		}

	}

	csvWriter.WriteAll(records)
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
