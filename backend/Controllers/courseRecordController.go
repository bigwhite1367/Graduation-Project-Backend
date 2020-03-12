package controllers

import (
    "github.com/gin-gonic/gin"
    database "ncu_14_project/backend/Database"
    "fmt"
    "net/http"
    "time"
    model "ncu_14_project/backend/Models"
)

func InsertCourseRecord(c *gin.Context){

    db, err := database.Connect()

    if err != nil {
        fmt.Println(err.Error())
    }


	stmt, dberr := db.Prepare("INSERT INTO courserecord (CourseID, CourseStartTime, CourseEndTime, CourseDate) VALUES (?,?,?,?)")

    defer stmt.Close()

	dt := time.Now()
	if dberr != nil {
	    fmt.Print(dberr.Error())
    }

    var courseRecord model.CourseRecord
	jsonerr := c.BindJSON(&courseRecord)

    if jsonerr != nil {
	    fmt.Println(jsonerr.Error())
    }

	rows, rowerr := db.Query("SELECT CourseEndTime FROM course WHERE CourseID = ?", courseRecord.CourseID)
	defer rows.Close()

	if rowerr != nil {
	    fmt.Println(rowerr.Error())
    }

	rows.Next()
	rows.Scan(&courseRecord.CourseEndTime)

    
	courseRecord.CourseDate = dt.Format("2006-01-02")
	_, dberr = stmt.Exec(courseRecord.CourseID, courseRecord.CourseStartTime, courseRecord.CourseEndTime,  courseRecord.CourseDate)

   	

    c.JSON(http.StatusOK, courseRecord)
    
}

func GetCourseRecord(c *gin.Context){
    //get user data by userid
}

func UpdateCourseRecord(c *gin.Context){

}

func DeleteCourseRecord(c *gin.Context){
	
}