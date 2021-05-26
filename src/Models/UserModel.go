package Models

import (
    "WoW-Assistant/src/Utils"
    "github.com/dgrijalva/jwt-go"
    "github.com/goonode/mogo"
    "time"
)

type User struct {
	Id		uint	`json:"id"`
	Name	string	`json:"name"`
	Email	string	`json:"email"`
	Phone	string	`json:"phone"`
	Address	string	`json:"address"`
}

//User struct is to handle user data
type UserMongo struct {
    mogo.DocumentModel `bson:",inline" coll:"users"`
    Email              string `idx:"{email},unique" json:"email" binding:"required"`
    Password           string `json:"password" binding:"required"`
    Name               string `json:"name"`
    // CreatedAt          *time.Time
    // UpdatedAt          *time.Time
    VerifiedAt *time.Time
}

// GetJwtToken returns jwt token with user email claims
func (user *UserMongo) GetJwtToken() (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "email": string(user.Email),
    })

    secretKey := Utils.EnvVar("TOKEN_KEY", "")
    tokenString, err := token.SignedString([]byte(secretKey))
    return tokenString, err
}

func init() {
    mogo.ModelRegistry.Register(UserMongo{})
}

func (u *User) TableName() string {
	return "user"
}
