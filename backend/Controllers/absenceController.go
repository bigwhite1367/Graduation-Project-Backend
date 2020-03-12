package controllers

import (
	"github.com/gin-gonic/gin"
    "github.com/appleboy/gin-jwt/v2"
    database "ncu_14_project/backend/Database"
    "fmt"
    "net/http"
    _ "time"
    model "ncu_14_project/backend/Models"
)

func GetAbsenceList(c *gin.Context){
	db, err := database.Connect()
    claims := jwt.ExtractClaims(c)
    user, _ := c.Get(identityKey)
    privilege := user.(*model.User).Privilege
    userID := claims[identityKey].(string)

    if err != nil {
        fmt.Println(err.Error())
    }

    m := make(map[string]string)
    m["1"] = "SELECT CourseID, StudentID, AbsenceDate FROM absence WHERE StudentID=?"
    m["0"] = "SELECT CourseID, StudentID, AbsenceDate FROM absence WHERE CourseID IN (SELECT CourseID FROM teachercourse WHERE TeacherID=?)"
    rows, err := db.Query(m[privilege],userID)
    defer rows.Close()

    if err != nil{
        fmt.Println(err.Error())
    }

    var absence model.Absence
    list := []model.Absence{}

    for rows.Next() {
    	err := rows.Scan(&absence.CourseID, &absence.StudentID, &absence.AbsenceDate)
    	
    	if err != nil {
            fmt.Println(err.Error())
        }
        
        list = append(list, absence)
    }

     c.JSON(http.StatusOK, list)

}
