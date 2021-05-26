package Models

import (
    "WoW-Assistant/src/Config"
    "context"
    "errors"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

type ItemService struct {}

// GetAllItems Fetch all Items data
func GetAllItems(item *[]ItemSQL) (err error) {
  if err = Config.DB.Find(item).Error; err != nil {
    return err
  }

  return err
}

// CreateItem ... Insert New Data
func CreateItem(item *ItemSQL) (err error) {
  if err = Config.DB.Create(item).Error; err != nil {
    return err
  }

  return err
}

// GetItemByID ... Fetch only one item by Id
func GetItemByID(item *ItemSQL, id string) (err error) {
  if err = Config.DB.Where("id = ?", id).First(item).Error; err != nil {
    return err
  }

  return err
}

// UpdateItem ... Update item
func UpdateItem(item *ItemSQL, id string) (err error) {
  fmt.Println(item)
  Config.DB.Save(item)
  return err
}

// DeleteItem ... Delete Item
func DeleteItem(item *ItemSQL, id string) (err error) {
  Config.DB.Where("id = ?", id).Delete(item)

  return err
}

//=================================================
// Mongo DB functions

// Create is to register new user
func Create(item []interface{}, realmName string) error {
    ctx := context.Background()
    fmt.Println("Created context")

    coll := Config.GetConnectionItem(realmName)
    fmt.Println("Got connection")

    _, err := coll.InsertMany(ctx, item)
    if err != nil {
        errors.New("can't insert that row")
    }

    //err = doc.FindOne(bson.M{"auction_id": item.AuctionID}, doc)
    //if err == nil {
    //   return errors.New("already exist")
    //}
    //itemModel := mogo.NewDoc(item).(*ItemMongo)
    //err = mogo.Save(itemModel)
    //if vErr, ok := err.(*mogo.ValidationError); ok {
    //    return vErr
    //}
    return err
}

//// Delete a user from DB
//func (itemService *ItemService) Delete(auctionID int) error {
//    item, _ := itemService.FindByAuctionID(auctionID)
//    conn := Config.GetConnection()
//    defer conn.Session.Close()
//    err := item.Remove()
//    return err
//}
//
//// Find user
//func (itemService *ItemService) Find(item *(ItemMongo)) (*ItemMongo, error) {
//    conn := Config.GetConnection()
//    defer conn.Session.Close()
//
//    doc := mogo.NewDoc(UserMongo{}).(*(ItemMongo))
//    err := doc.FindOne(bson.M{"auction_id": item.AuctionID}, doc)
//
//    if err != nil {
//        return nil, err
//    }
//    return doc, nil
//}
//
//// Find user from Auction ID
//func (itemService *ItemService) FindByAuctionID(auctionID int) (*ItemMongo, error) {
//    conn := Config.GetConnection()
//    defer conn.Session.Close()
//
//    item := new(ItemMongo)
//    item.AuctionID = auctionID
//    return itemService.Find(item)
//}
