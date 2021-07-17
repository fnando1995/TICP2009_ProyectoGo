// package main

// import (
//   "gorm.io/gorm"
//   "gorm.io/driver/sqlite"
// )

// type Product struct {
//   gorm.Model
//   coinId  string
//   coinSymbol string
//   coinName string
// }

package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type CoinS struct {
	gorm.Model
	coinId     string
	coinSymbol string
	coinName   string
}

func main() {

	db, err := gorm.Open(sqlite.Open("coins.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&CoinS{})

	// Create
	db.Create(&CoinS{coinId: "bitcoin", coinSymbol: "btc", coinName: "Bitcoin"})
}
