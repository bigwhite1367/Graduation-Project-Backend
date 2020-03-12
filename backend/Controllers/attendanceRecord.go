package controllers

import(
	"github.com/gin-gonic/gin"
    _ "github.com/appleboy/gin-jwt/v2"
	database "ncu_14_project/backend/Database"
	"fmt"
    _ "strconv"
    "net/http"
    model "ncu_14_project/backend/Models"
)

func InsertLeave(c *gin.Context) {
    
    var attendance model.Attendance
    var studentID string
	db ,err := database.Connect()
	
	if err != nil{
    	fmt.Println(err)
    }
	insertStmt, inserterr := db.Prepare("INSERT INTO attendancerecord (CourseID, StudentID, AttendanceDate, AttendanceRate) VALUES (?,?,?,?)")
    updateStmt, updateerr := db.Prepare("UPDATE attendancerecord SET AttendanceRate=2 WHERE CourseID=? AND StudentID=? AND AttendanceDate=?")
    
    if inserterr != nil{
    	fmt.Println(inserterr)
    }

    if updateerr != nil{
    	fmt.Println(updateerr)
    }

   // defer insertStmt.Close()
    //defer updateStmt.Close()


    if err:=c.ShouldBind(&attendance); err != nil{
    	fmt.Println(err)
    }

    row := db.QueryRow("SELECT StudentID FROM attendancerecord WHERE CourseID=? AND StudentID=? AND AttendanceDate=?", attendance.CourseID, attendance.StudentID, attendance.AttendanceDate)
  
    if err := row.Scan(&studentID); err != nil{
        _ , err:=insertStmt.Exec(attendance.CourseID, attendance.StudentID, attendance.AttendanceDate, 2)
        
        if err !=nil {
            fmt.Println(err)
        }

    }else {
    	_ , err:=updateStmt.Exec(attendance.CourseID, attendance.StudentID, attendance.AttendanceDate)

    	 if err !=nil {
            fmt.Println(err)   
        }

    }

    checkAbsence(attendance.CourseID, attendance.StudentID, attendance.AttendanceDate)
    attendance.AttendanceRate = 2
    
    c.JSON(http.StatusOK, attendance)


}

func checkAbsence(courseID string, studentID string, absenceDate string) bool{
    db ,err := database.Connect()
    
    if err != nil{
        fmt.Println(err)
    }

    deleteStmt, deleteerr := db.Prepare("DELETE FROM absence WHERE CourseID=? AND StudentID=? AND AbsenceDate LIKE ?")
    if deleteerr != nil{
        fmt.Println(deleteerr)
    }

    row:= db.QueryRow("SELECT AttendanceDate FROM absence WHERE CourseID=? AND StudentID=? AND AbsenceDate LIKE ?", courseID, studentID, absenceDate+"%")

    var absence model.Absence
    err = row.Scan(&absence.CourseID, &absence.StudentID, &absence.AbsenceDate)

    if err != nil {
        _, err := deleteStmt.Exec(courseID, studentID, absenceDate+"%")
        if err !=nil {
            fmt.Println(err)   
        }

        return true
    } 

    return false

}








