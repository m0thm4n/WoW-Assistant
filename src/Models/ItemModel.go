package Models

type Item struct {
  ID uint
  ItemID int
  Name string
  SubclassName string
  SubclassID int
  QualityType string
  QualityName string
  Level int
  RequiredLevel int
  ItemClassName string
  ItemClassID int
  InventoryTypeName string
  InventoryType string
  PurchasePrice int
  SellPrice int
  MaxCount int
  IsEquippable bool
  IsStackable bool
}

func (item *Item) TableName() string {
  return "item"
}
