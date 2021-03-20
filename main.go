package main

import (
  "WoW-Assistant/src/Routes"
  "WoW-Assistant/src/Wow"

  // "WoW-Assistant/src/Wow"
  "fmt"
)

var err error

func main() {
	//Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	//if err != nil {
	//	fmt.Println("Status: ", err)
	//}
  //
	//defer Config.DB.Close()
	//Config.DB.AutoMigrate(&Models.User{})

	r := Routes.SetupRouter()

  // {
  //    id: 110241996,
  //    item: {id: 172092, context: ""},
  //    buyout: 47,
  //    quantity: 21000,
  //    time left: VERY_LONG
  // }

  auctions := Wow.WowAuctions("aegwynn")
  fmt.Println(auctions.Auctions[0])

  // running
	r.Run()
}
