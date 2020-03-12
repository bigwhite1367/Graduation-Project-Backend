package models


type Attendance struct{
	CourseID          string   `form:"courseID"                                 json:"courseID"`
	StudentID         string   `form:"studentID"                                json:"studentID"`
	AttendanceDate    string   `form:"attendanceDate" time_format:"2006-01-02"  json:"attendanceDate"`
	AttendanceRate    float64  `form:"attendanceRate"                           json:"attendanceRate"`
}