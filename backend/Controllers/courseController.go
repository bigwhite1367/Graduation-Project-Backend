package controllers

import( 
	"github.com/gin-gonic/gin"
    "github.com/appleboy/gin-jwt/v2"
	database "ncu_14_project/backend/Database"
    "fmt"
    "strconv"
    "net/http"
    model "ncu_14_project/backend/Models"
)


func InsertCourse(c *gin.Context){

}

func GetCourse(c *gin.Context) {

    db, err := database.Connect()
    if err != nil {
        fmt.Println(err)
    }

    courseID := c.Param("courseID")
	rows, err := db.Query("SELECT CourseName, CoursePlace, CourseStartTime, CourseEndTime, CourseAttendance, CourseDay, Population FROM course WHERE CourseID=?", courseID)

    defer rows.Close()

    if err != nil {
	    fmt.Println(err)
    }
    
    

    var course model.Course
    for rows.Next() {
        course.CourseID = courseID
        err := rows.Scan(&course.CourseName, &course.CoursePlace, &course.CourseStartTime,  &course.CourseEndTime, &course.CourseAttendance, &course.CourseDay, &course.Population)

        if err != nil {
	        fmt.Println(err.Error())
	    }
    }
    rows.Close()

    c.JSON(http.StatusOK, course)

}

func GetCurriculum(c *gin.Context){

    db, err := database.Connect()

    if err != nil {
        fmt.Println(err.Error())
    }

    


    class := [13]string{"一","二","三","四","午","五","六","七","八","九","A","B","C"}
    claims := jwt.ExtractClaims(c)
    user, _ := c.Get(identityKey)
    privilege := user.(*model.User).Privilege
    var queryString string
    userID := claims[identityKey].(string)
    if privilege == "1" {
        queryString = "SELECT CourseName, CourseStartTime, CourseEndTime, CourseDay FROM course WHERE CourseID IN (SELECT CourseID FROM studentcourse WHERE StudentID=?)"
    }else if privilege == "0"{
        queryString = "SELECT CourseName, CourseStartTime, CourseEndTime, CourseDay FROM course WHERE CourseID IN (SELECT CourseID FROM teachercourse WHERE TeacherID=?)"
    }

    rows, err := db.Query(queryString, userID)

    defer rows.Close()
    if err != nil {
        fmt.Println(err.Error())
    }   
  

    var course model.Course
    
    list := [13]model.Curriculum{}

    for i, _ := range list{
        list[i].Class = class[i]
    }
    for rows.Next(){
        err := rows.Scan(&course.CourseName, &course.CourseStartTime, &course.CourseEndTime, &course.CourseDay)
        if err != nil {
            fmt.Println(err.Error())
        }

        startTime, _ := strconv.Atoi(course.CourseStartTime[0:2])
        endTime, _ := strconv.Atoi(course.CourseEndTime[0:2])

        for i:=startTime; i<endTime; i++ {
            switch course.CourseDay {
                case "1":
                    list[i-8].Mon = course.CourseName
                case "2":
                    list[i-8].Tue = course.CourseName
                case "3":
                    list[i-8].Wed = course.CourseName
                case "4":
                    list[i-8].Thu = course.CourseName
                case "5":
                    list[i-8].Fri = course.CourseName
            }  
        }
    }

    c.JSON(http.StatusOK, list)
}
















