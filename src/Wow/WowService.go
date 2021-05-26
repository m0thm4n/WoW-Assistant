package Wow

import (
	"context"
	"fmt"
	"log"

	"github.com/FuzzyStatic/blizzard/v2"
	"github.com/FuzzyStatic/blizzard/v2/wowgd"

	"WoW-Assistant/src/Config"
)

func getClient() *blizzard.Client {
	config := Config.LoadConfiguration("config.json")
	clientID := config.ClientID
	key := config.Key

	ctx := context.Background()

	blizz := blizzard.NewClient(clientID, key, blizzard.US, blizzard.EnUS)

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

func GetItem(id int) *wowgd.Item {
	ctx := context.Background()

	client := getClient()

	item, _, err := client.WoWItem(ctx, id)
	if err != nil {
		ignoreError("Got 404 back from API skipping", err)
	}

	return item
}

func ignoreError(val interface{}, err error) interface{} {
	return val
}
