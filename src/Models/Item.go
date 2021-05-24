package Models

import (
  "WoW-Assistant/src/Config"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
)

// GetAllItems Fetch all Items data
func GetAllItems(item *[]Item) (err error) {
  if err = Config.DB.Find(item).Error; err != nil {
    return err
  }

  return err
}

// CreateItem ... Insert New Data
func CreateItem(item *Item) (err error) {
  if err = Config.DB.Create(item).Error; err != nil {
    return err
  }

  return err
}

// GetItemByID ... Fetch only one item by Id
func GetItemByID(item *Item, id string) (err error) {
  if err = Config.DB.Where("id = ?", id).First(item).Error; err != nil {
    return err
  }

  return err
}

// UpdateItem ... Update item
func UpdateItem(item *Item, id string) (err error) {
  fmt.Println(item)
  Config.DB.Save(item)
  return err
}

// DeleteItem ... Delete Item
func DeleteItem(item *Item, id string) (err error) {
  Config.DB.Where("id = ?", id).Delete(item)

  return err
}


