package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Coin struct {
	gorm.Model
	CoinId     string `json:"id"`
	CoinSymbol string `json:"symbol"`
	CoinName   string `json:"name"`
}

func refresh(w http.ResponseWriter, r *http.Request) {
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

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pagina de bienvenida ....")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/refresh", refresh)
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
