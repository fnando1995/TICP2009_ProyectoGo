package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"html/template"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Coin struct {
	gorm.Model
	CoinId     string `json:"id"`
	CoinSymbol string `json:"symbol"`
	CoinName   string `json:"name"`
}

func updatedatabase(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Consultando API Gecko para actualizar monedas")
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
	json.Unmarshal([]byte(responseData), &data)

	fmt.Println("Limpiando base de datos...")
	_, file := os.Stat("database/coins.db")
	if !os.IsNotExist(file) {
		e := os.Remove("database/coins.db")
		if e != nil {
			log.Fatal(e)
		}
	}
	db, err := gorm.Open(sqlite.Open("database/coins.db"), &gorm.Config{})
	if err != nil {
		panic("Conexion a base de datos fallida...")
	}
	time.Sleep(1 * time.Second)
	db.AutoMigrate(&Coin{})

	fmt.Println("Agregando informacion...")
	for i := 0; i < len(data); i++ {
		time.Sleep(1 * time.Millisecond)
		db.Create(&Coin{CoinId: data[i].CoinId, CoinSymbol: data[i].CoinSymbol, CoinName: data[i].CoinName})
	}
	fmt.Println("Informacion actualizada...")
	fmt.Fprintf(w, "Se actualizo la base de datos ...")
}

func exchange(w http.ResponseWriter, r *http.Request) {
	getData := func(coin1, convertion string) float64 {
		url := "https://api.coingecko.com/api/v3/simple/price?ids=" + coin1 + "&vs_currencies=" + convertion
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		var data1 map[string]map[string]float64
		err = json.NewDecoder(resp.Body).Decode(&data1)
		if err != nil {
			panic(err)
		}
		return data1[coin1][convertion]
	}
	fmt.Println("Endpoint Hit: exchangeMid")
	coin1 := strings.Split(r.URL.Query()["coin1"][0], ",")[0]
	coin2 := strings.Split(r.URL.Query()["coin2"][0], ",")[0]
	convertion := strings.Split(r.URL.Query()["mid"][0], ",")[0]
	numerador := getData(coin1, convertion)
	denominador := getData(coin2, convertion)
	// fmt.Fprint(w, )
	fmt.Fprint(w, "<h1>", "1 ", coin1, " es equivalente a  ", numerador/denominador, " ", coin2, "</h1>")

}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENTER HOME")
	t1, _ := template.ParseFiles("templates/home.html")
	t1.Execute(w, nil)
}

func update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENTER UPDATE")
	t2, _ := template.ParseFiles("templates/update.html")
	t2.Execute(w, nil)
	fmt.Println("FINISH UPDATE")
}

func handleRequests() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/update", update)
	http.HandleFunc("/updatedatabase", updatedatabase)
	http.HandleFunc("/exchange", exchange)

	port := ":9000"
	log.Println("Listening on port ", port)
	http.ListenAndServe(port, nil)
}

func main() {
	handleRequests()
}
