package Middlewares

import (
    "WoW-Assistant/src/Models"
    "WoW-Assistant/src/Utils"
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "net/http"
    "strings"
)

//Authentication is for auth middleware
func Authentication() gin.HandlerFunc {
   return func(c *gin.Context) {
       authHeader := c.Request.Header.Get("Authentication")
       if len(authHeader) == 0 {
           c.AbortWithStatusJSON(400, gin.H{
               "error": "Authentication header is missing",
           })
           return
       }
       temp := strings.Split(authHeader, "Bearer")
       if len(temp) < 2 {
           c.AbortWithStatusJSON(400, gin.H{"error": "Invalid token"})
           return
       }
       tokenString := strings.TrimSpace(temp[1])
       token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
           // if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
           // 	return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
           // }
           secretKey := Utils.EnvVar("TOKEN_KEY", "")
           return []byte(secretKey), nil
       })

       if err != nil {
           c.AbortWithStatusJSON(401, gin.H{
               "error": err.Error(),
           })
           return
       }
       if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
           email := claims["email"].(string)
           userservice := Models.UserService{}
           user, err := userservice.FindByEmail(email)
           if err != nil {
               c.AbortWithStatusJSON(402, gin.H{
                   "error": "User not found",
               })
               return
           }
           c.Set("user", user)
           c.Next()
       } else {
           c.AbortWithStatusJSON(400, gin.H{
               "error": "Token is not valid",
           })
           return
       }
   }
}

// ErrorHandler is for global errors
func ErrorHandler(c *gin.Context) {
   c.Next()

   if len(c.Errors) > 0 {
       c.JSON(http.StatusBadRequest, gin.H{
           "errors": c.Errors,
       })
   }
}
