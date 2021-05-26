package Controllers
//
//import (
//    "WoW-Assistant/src/Models"
//    "github.com/gin-gonic/gin"
//    "golang.org/x/crypto/bcrypt"
//    "log"
//)
//
//// AuthController is for auth logic
//type AuthController struct{}
//
//// Login is to process login request
//func (auth *AuthController) Login(c *gin.Context) {
//
//    var loginInfo Models.UserMongo
//    if err := c.ShouldBindJSON(&loginInfo); err != nil {
//        c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
//        return
//    }
//
//    // TODO
//    userService := Models.UserService{}
//    user, errf := userService.Find(&loginInfo)
//    if errf != nil {
//        c.AbortWithStatusJSON(401, gin.H{"error": "Not Found"})
//        return
//    }
//
//    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password))
//    if err != nil {
//        c.AbortWithStatusJSON(402, gin.H{"error": "Email or password is invalid"})
//        return
//    }
//
//    token, err := user.GetJwtToken()
//    if err != nil {
//        c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
//        return
//    }
//
//    c.JSONP(200, gin.H{
//        "token": token,
//    })
//}
//
//// Profile is to provide current user info
//func (auth *AuthController) Profile(c *gin.Context) {
//    user := c.MustGet("user").(*(Models.User))
//
//    c.JSON(200, gin.H{
//        "user_name": user.Name,
//        "email": user.Email,
//    })
//}
//
//// Signup is for user signup
//func (auth *AuthController) Signup(c *gin.Context) {
//    type signupInfo struct {
//        Email       string  `json:"email" binding:"required"`
//        Password    string  `json:"password" binding:"required"`
//        Name        string  `json:"name"`
//    }
//
//    var info signupInfo
//    if err := c.ShouldBindJSON(&info); err != nil {
//        c.AbortWithStatusJSON(401, gin.H{"error": "Please input all fields"})
//        return
//    }
//
//    user := Models.UserMongo{}
//    user.Email = info.Email
//    hash, err := bcrypt.GenerateFromPassword([]byte(info.Password), bcrypt.MinCost)
//    if err != nil {
//        log.Fatal(err)
//        return
//    }
//
//    user.Password = string(hash)
//    user.Name = info.Name
//    userService := Models.UserService{}
//    err = userService.Create(&user)
//    if err != nil {
//        c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
//    } else {
//        c.JSON(200, gin.H{"result": "ok"})
//    }
//
//    return
//}
