package Controllers

import (
  "WoW-Assistant/src/Models"
  "github.com/gin-gonic/gin"
  "net/http"
)

// GetItems ... Get all Items
func GetItems(c *gin.Context) {
  var item []Models.Item
  err := Models.GetAllItems(&item)
  if err != nil {
    c.AbortWithStatus(http.StatusNotFound)
  } else {
    c.JSON(http.StatusOK, item)
  }
}
