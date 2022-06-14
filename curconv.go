package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	//   "log"
	//   "io/ioutil"
	// "net/url"
	"github.com/labstack/echo"
)

// I imported four important packages "net/http" to access the core go functionality, "fmt" for formatting our text,
//  "github.com/labstack/echo" to import the echo framework, "encoding/json" implements encoding and decoding of json

type Conversion struct {
	ResultInUSD float64 `json:"result"`
	// ProfileImage string `json:"profile_image"`
}

// type Queryvalue struct {
// 	From string `json:"from" query:"from"`
// 	To string `json:"to" query:"to"`
// 	AmountInRupee uint `json:"amount" query:"amount"`
// }

func main() {
	url := "https://api.apilayer.com/exchangerates_data/convert?to=USD&from=INR&amount=120000"
	e := echo.New()
	// 	e.POST("/amount", func(c echo.Context) error {

	// 		// postRequest := new(model.Queryvalue)
	// 		if err := c.Bind(postRequest); err != nil {
	// 		 	return err
	// 		}
	// 		params := url.{
	// 			"amount":
	// }
	// var resp bytes.Buffer
	// var b = []byte(
	// 	fmt.Sprintf(`{
	// 		"from": INR,
	// 		"to": USD,
	// 		"amount": %d,
	// 		}`, postRequest.AmountInRupee),
	// 	)
	// 	err := json.Indent(&resp, b, "", "  ")
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return c.JSONBlob(
	// 		http.StatusOK,
	// 		[]byte(
	// 			fmt.Sprintf("%s", resp.Bytes()),
	// 		),
	// }

	// 	)

	// if err ! = nil {
	// 	log.Fatal(err)
	// 	}
	// q := u.Query
	// q.Set("amount", "")
	e.GET("/convert", func(c echo.Context) error {
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Set("apikey", "hdDaP85PW72WhBERFHG1rJdVfGj0JrPZ")

		if err != nil {
			fmt.Println(err)
		}
		res, err := client.Do(req)
		if res.Body != nil {
			defer res.Body.Close()
		}
		var data Conversion
		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			fmt.Println(err)
		}

		return c.JSON(http.StatusOK, data)
		//   body, err := ioutil.ReadAll(res.Body)

		//   fmt.Println(string(body))
	})
	e.Logger.Fatal(e.Start(":1323"))
}
