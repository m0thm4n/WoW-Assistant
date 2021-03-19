package Wow

import (
  "fmt"
  "github.com/FuzzyStatic/blizzard/v2"
  "context"
  "github.com/FuzzyStatic/blizzard/v2/wowgd"
  "log"
)



func getClient() *blizzard.Client {
  ctx := context.Background()

  blizz := blizzard.NewClient("d2a18a7c4f364725822fc2ef3740ecc5", "Bv5t7l2AifIWsgESnosKhUbYnlRVFLeG", blizzard.US, blizzard.EnUS)

  err := blizz.AccessTokenRequest(ctx)
  if err != nil {
    fmt.Println(err)
  }

  return blizz
}

// WowAuctions gets all auctions from a realm
func WowAuctions(realmToGet string) *wowgd.AuctionHouse {
  ctx := context.Background()

  client := getClient()

  realm, _, err := client.WoWRealm(ctx, realmToGet)
  if err != nil {
    log.Fatal(err)
  }

  wowAuctions, _, err := client.WoWAuctions(ctx, realm.ID)
  if err != nil {
    log.Fatal(err)
  }

  return wowAuctions
}
