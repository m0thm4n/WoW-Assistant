package Wow

import (
    "WoW-Assistant/src/Config"
    "context"
    "fmt"
    "github.com/FuzzyStatic/blizzard/v2"
    "github.com/FuzzyStatic/blizzard/v2/wowgd"
    "log"
    "sync"
)

// GetClient gets a blizzard client
func GetClient() *blizzard.Client {
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

// Auctions gets all auctions from a realm
func Auctions(realmToGet string, client *blizzard.Client) *wowgd.AuctionHouse {
	ctx := context.Background()

    realm, _, err := client.WoWRealm(ctx, realmToGet)
    if err != nil {
        log.Fatal("Failed to get realm")
    }


    wowAuctions, _, err := client.WoWAuctions(ctx, realm.ID)
    if err != nil {
        log.Fatal("Failed to get auction")
    }

    return wowAuctions
}

func GetItem(id int, client *blizzard.Client, wg sync.WaitGroup) *wowgd.Item {
	ctx := context.Background()

    item, _, err := client.WoWItem(ctx, id)

    if err != nil {
        ignoreError("Got 404 back from API skipping", err)
    }

    return item
}

func ignoreError(val interface{}, err error) interface{} {
	return val
}
