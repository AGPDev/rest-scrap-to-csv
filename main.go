package main

import "agporto/scrap-products/scrap"

func main() {
	scrap.Start()
	// scrap.GetFabricantes()
	// getCategories()

	// // var myCookie = &http.Cookie{
	// // 	Name:  "PHPSESSID",
	// // 	Value: "xxxxxxxxxxxxxx",
	// // }

	// bar := pb.New(registros).Prefix("Registros ")
	// bar.ShowCounters = true
	// bar.Add(registro)
	// bar.Start()

	// for p := page; p <= pages; p++ {
	// 	fmt.Println("Página: " + strconv.Itoa(p) + "/" + strconv.Itoa(pages))

	// 	// Make HTTP GET request
	// 	request, err := http.NewRequest("GET", "http://"+strconv.Itoa(p)+"&pages=50&yt0=Atualizar", nil)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	// request.AddCookie(myCookie)

	// 	client := &http.Client{}
	// 	response, err := client.Do(request)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer response.Body.Close()

	// 	// Get the response body as a string
	// 	document, err := goquery.NewDocumentFromReader(response.Body)
	// 	if err != nil {
	// 		log.Fatal("Error loading HTTP response body. ", err)
	// 	}

	// 	var urls []string
	// 	// Find all links and process them with the function
	// 	// defined earlier
	// 	document.Find("table tbody a").Each(func(index int, element *goquery.Selection) {
	// 		href, exists := element.Attr("href")
	// 		if exists {
	// 			urls = append(urls, href)
	// 		}
	// 	})

	// 	for i := 0; i < len(urls); i++ {
	// 		// Make HTTP GET request
	// 		request, err = http.NewRequest("GET", "http://"+urls[i], nil)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}

	// 		// request.AddCookie(myCookie)

	// 		client := &http.Client{}
	// 		response, err := client.Do(request)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		defer response.Body.Close()

	// 		// Get the response body as a string
	// 		document, err := goquery.NewDocumentFromReader(response.Body)
	// 		if err != nil {
	// 			log.Fatal("Error loading HTTP response body. ", err)
	// 		}

	// 		// Find all links and process them with the function
	// 		// defined earlier
	// 		document.Find("div.conteudo div.field span").Each(func(index int, element *goquery.Selection) {
	// 			axis := alf[index] + strconv.Itoa(registro)
	// 			text := element.Text()
	// 			xlsx.SetCellValue("Sheet1", axis, text)
	// 		})

	// 		registro++
	// 		bar.Increment()
	// 	}

	// 	//xlsx.SetActiveSheet(index)
	// 	err = xlsx.SaveAs("./file.xlsx")
	// 	if err != nil {
	// 		fmt.Printf(err.Error())
	// 	}
	// }
}

// func getCategories() {
// 	// var alf = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
// 	// 	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC",
// 	// 	"AD", "AE", "AF", "AG", "AH", "AI", "AJ", "AK", "AL", "AM", "AN", "AO", "AP", "AQ", "AR", "AS", "AT", "AU"}
// 	var page = 1
// 	var pages = 20
// 	var xlsx, _ = excelize.OpenFile("./acessorios.xlsx")
// 	// var myCookie = &http.Cookie{
// 	// 	Name:  "PHPSESSID",
// 	// 	Value: "xxxxxxxxxxxxxx",
// 	// }

// 	for p := page; p <= pages; p++ {
// 		fmt.Println("Página: " + strconv.Itoa(p) + "/" + strconv.Itoa(pages))

// 		// Make HTTP GET request
// 		request, err := http.NewRequest("GET", "https://br.ingrammicro.com/_layouts/CommerceServer/IM/search2.aspx#PNavDS=N:4294641002&t=pTab#"+strconv.Itoa(p), nil)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		// request.AddCookie(myCookie)

// 		client := &http.Client{}
// 		response, err := client.Do(request)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println(response)
// 		defer response.Body.Close()

// 		// Get the response body as a string
// 		document, err := goquery.NewDocumentFromReader(response.Body)
// 		if err != nil {
// 			log.Fatal("Error loading HTTP response body. ", err)
// 		}

// 		var urls []string
// 		// Find all links and process them with the function
// 		// defined earlier
// 		document.Find(".product-name a").Each(func(index int, element *goquery.Selection) {
// 			href, exists := element.Attr("href")
// 			if exists {
// 				urls = append(urls, href)
// 			}
// 		})

// 		for i := 0; i < len(urls); i++ {
// 			fmt.Println("Processando produto: " + strconv.Itoa(i))

// 			// Make HTTP GET request
// 			request, err = http.NewRequest("GET", "https://br.ingrammicro.com"+urls[i], nil)
// 			if err != nil {
// 				fmt.Println(err)
// 			}

// 			// request.AddCookie(myCookie)

// 			client := &http.Client{}
// 			response, err := client.Do(request)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			defer response.Body.Close()

// 			// Get the response body as a string
// 			document, err := goquery.NewDocumentFromReader(response.Body)
// 			if err != nil {
// 				log.Fatal("Error loading HTTP response body. ", err)
// 			}
// 			fmt.Println(document)
// 			// Find all links and process them with the function
// 			// defined earlier
// 			// document.Find(".content").Each(func(index int, element *goquery.Selection) {
// 			// 	axis := alf[0] + strconv.Itoa(i)
// 			// 	text := element.Find(".product-name span").Text()
// 			// 	xlsx.SetCellValue("Sheet1", axis, text)

// 			// 	fmt.Println("Produto: " + text)

// 			// 	axis = alf[1] + strconv.Itoa(i)
// 			// 	text = element.Find(".product-details-vendor").Text()
// 			// 	xlsx.SetCellValue("Sheet1", axis, text)

// 			// 	axis = alf[2] + strconv.Itoa(i)
// 			// 	text = element.Find(".category-table tbody tr td a").Text()
// 			// 	xlsx.SetCellValue("Sheet1", axis, text)

// 			// 	axis = alf[3] + strconv.Itoa(i)
// 			// 	text = element.Find(".product-detail-description span").Text()
// 			// 	xlsx.SetCellValue("Sheet1", axis, text)

// 			// 	axis = alf[4] + strconv.Itoa(i)
// 			// 	text, _ = element.Find("#extended-specifications table").Html()
// 			// 	xlsx.SetCellValue("Sheet1", axis, text)

// 			// 	axis = alf[5] + strconv.Itoa(i)
// 			// 	text, _ = element.Find(".change-cursor.mouseDown").Attr("href")
// 			// 	xlsx.SetCellValue("Sheet1", axis, text)
// 			// })

// 			// fmt.Println("Fim produto: " + strconv.Itoa(i))
// 		}

// 		//xlsx.SetActiveSheet(index)
// 		err = xlsx.SaveAs("./acessorios.xlsx")
// 		if err != nil {
// 			fmt.Printf(err.Error())
// 		}
// 	}
// }
