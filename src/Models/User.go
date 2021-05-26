package Models

import (
    "WoW-Assistant/src/Config"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

type UserService struct {}

// MySQL DB functions

// GetAllUsers Fetch all user data
func GetAllUsers(user *[]User) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}

	return err
}

// CreateUser ... Insert New data
func CreateUser(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}

// GetUserByID ... Fetch only one user by Id
func GetUserByID(user *User, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}

	return nil
}

// UpdateUser ... Update user
func UpdateUser(user *User, id string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}

// DeleteUser ... Delete user
func DeleteUser(user *User, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}

////=================================================
//// Mongo DB functions
//
//// Create is to register new user
//func (userService *UserService) Create(user *(UserMongo)) error {
//    conn := Config.GetConnection()
//    defer conn.Session.Close()
//
//    doc := mogo.NewDoc(UserMongo{}).(*(UserMongo))
//    err := doc.FindOne(bson.M{"email": user.Email}, doc)
//    if err == nil {
//        return errors.New("Already Exist")
//    }
//    userModel := mogo.NewDoc(user).(*(UserMongo))
//    err = mogo.Save(userModel)
//    if vErr, ok := err.(*mogo.ValidationError); ok {
//        return vErr
//    }
//    return err
//}
//
//// Delete a user from DB
//func (userService *UserService) Delete(email string) error {
//    user, _ := userService.FindByEmail(email)
//    conn := Config.GetConnection()
//    defer conn.Session.Close()
//    err := user.Remove()
//    return err
//}
//
//// Find user
//func (userService *UserService) Find(user *(UserMongo)) (*UserMongo, error) {
//    conn := Config.GetConnection()
//    defer conn.Session.Close()
//
//    doc := mogo.NewDoc(UserMongo{}).(*(UserMongo))
//    err := doc.FindOne(bson.M{"email": user.Email}, doc)
//
//    if err != nil {
//        return nil, err
//    }
//    return doc, nil
//}
//
//// Find user from email
//func (userService *UserService) FindByEmail(email string) (*UserMongo, error) {
//    conn := Config.GetConnection()
//    defer conn.Session.Close()
//
//    user := new(UserMongo)
//    user.Email = email
//    return userService.Find(user)
//}
