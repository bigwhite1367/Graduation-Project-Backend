package controllers

import(
	"fmt"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	models "ncu_14_project/backend/Models"
	database "ncu_14_project/backend/Database"
)

var identityKey = "id"

func PayloadFunc(data interface{}) jwt.MapClaims{
	if v, ok := data.(*models.User); ok {
			return jwt.MapClaims{
				identityKey: v.UserID,
			}
		}
		return jwt.MapClaims{}
}

func IdentityHandler(c *gin.Context) interface{} {
	 db, err := database.Connect()

    if err != nil {
        fmt.Println(err.Error())
    }

    claims := jwt.ExtractClaims(c)
    var user models.User
    user.UserID = claims[identityKey].(string)

    rows, err := db.Query("SELECT UserName, Privilege FROM user WHERE UserID=?", user.UserID)

    defer rows.Close()
    if err != nil {
	    fmt.Println(err.Error())
	}

	for rows.Next(){
		err :=rows.Scan(&user.UserName, &user.Privilege)

		if err != nil {
	        fmt.Println(err.Error())
	    }
	}



    return &models.User{
	        UserID:    user.UserID,
	        UserName:  user.UserName,
	        Privilege: user.Privilege,
    }
}
func Authenticator(c *gin.Context)  (interface{}, error){

	 db, err := database.Connect()

    if err != nil {
        fmt.Println(err.Error())
    }

    var loginVals models.Login
    var auth models.Login
    var user models.User
	if err := c.ShouldBind(&loginVals); err != nil {
	    return "", jwt.ErrMissingLoginValues
	}
		
    userID := loginVals.UserID
	password := loginVals.Password
        
    auth.UserID = userID

    rows, err := db.Query("SELECT UserName, Password, Privilege FROM user WHERE UserID=?", userID)
    if err != nil {
	    fmt.Println(err.Error())
    }	
        
    defer rows.Close()
    for rows.Next() {
        err := rows.Scan(&user.UserName, &auth.Password, &user.Privilege)

        if err != nil {
	        fmt.Println(err.Error())
	    }
        
    }

	if (userID == auth.UserID && password == auth.Password) {
		user.UserID = userID
		return &models.User{
			UserID:    user.UserID,
			UserName:  user.UserName,
			Privilege: user.Privilege,
		}, nil
	}

	return nil, jwt.ErrFailedAuthentication
}

func Authorizator (data interface{}, c *gin.Context) bool {
    
    /*if _, ok := data.(*models.User); ok {
        return true
    }
    */
    db, err := database.Connect()

    if err != nil {
        fmt.Println(err.Error())
    }

    
	if v, ok := data.(*models.User); ok {

	    row:= db.QueryRow("SELECT UserName FROM user WHERE UserID=?", v.UserID)
	 
	    err:=row.Scan(&v.UserName)
		if(err == nil){	
			return true
		}
	    
    }
	return false
}






