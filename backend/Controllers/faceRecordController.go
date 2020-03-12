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

func GetFaceRecord(c *gin.Context){
    
    db, err := database.Connect()

    if err != nil {
        fmt.Println(err.Error())
    }

    claims := jwt.ExtractClaims(c)
    user, _ := c.Get(identityKey)
    privilege := user.(*model.User).Privilege

    userID := claims[identityKey].(string)

    /*query := "SELECT facerecord.created_at created_at,facerecord.period period,facerecord.status status, course.CourseName, course.CourseID FROM facerecord INNER JOIN course ON (facerecord.CourseID=course.CourseID) WHERE course.CourseID IN "
    student := "(SELECT CourseID FROM studentcourse WHERE StudentID=?)"
    teacher := "(SELECT CourseID FROM teachercourse WHERE TeacherID=?)"
    */
    m := make(map[string]string)

    m["1"] = "SELECT facerecord.created_at created_at,facerecord.period period,facerecord.status status, course.CourseName, course.CourseID FROM facerecord INNER JOIN course ON (facerecord.CourseID=course.CourseID) WHERE course.CourseID IN (SELECT CourseID FROM studentcourse WHERE StudentID=?)"
    m["0"] = "SELECT facerecord.created_at created_at,facerecord.period period,facerecord.status status, course.CourseName, course.CourseID FROM facerecord INNER JOIN course ON (facerecord.CourseID=course.CourseID) WHERE course.CourseID IN (SELECT CourseID FROM teachercourse WHERE TeacherID=?)"

    rows, err := db.Query(m[privilege], userID)
	if err != nil {
	    fmt.Println(err.Error())
    }
    defer rows.Close()

    var faceRecord model.FaceRecord
    list := []model.FaceRecord{}

    
    for rows.Next(){
    	err := rows.Scan(&faceRecord.Created_at, &faceRecord.Total,&faceRecord.Status, &faceRecord.CourseName, &faceRecord.CourseID)
    	
        if err != nil {
            fmt.Println(err.Error())
        }
        
        list = append(list, faceRecord)
	
    }

    c.JSON(http.StatusOK, list)
}

