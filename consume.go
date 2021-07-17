package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Coin struct {
	coinId     string `json:"id"`
	coinSymbol string `json:"symbol"`
	coinName   string `json:"name"`
}

func main() {
	response, err := http.Get("https://api.coingecko.com/api/v3/coins/list")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data []Coin
	err = json.Unmarshal([]byte(responseData), &data)

	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("%x - %x \n", i, data[i])
	}
	fmt.Println(len(data))

}
