package main

import (
    "WoW-Assistant/src/Config"
    "WoW-Assistant/src/Models"
    "WoW-Assistant/src/Routes"
    "fmt"
    // "WoW-Assistant/src/Wow"
    "github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status: ", err)
	}
  //
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	Config.DB.AutoMigrate(&Models.ItemSQL{})
	fmt.Println("AutoMigrate complete")

	r := Routes.SetupRouter()

  // {
  //    id: 110241996,
  //    item: {id: 172092, context: ""},
  //    buyout: 47,
  //    quantity: 21000,
  //    time left: VERY_LONG
  // }

  //auctions := Wow.Auctions("aegwynn")
  //fmt.Println(auctions.Auctions[0])

  // running
  r.Run()
}
