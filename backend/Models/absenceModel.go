package models

type Absence struct{
	CourseID       string  `json:"courseID"`
	StudentID      string  `json:"studentID"`
	AbsenceDate    string  `json:"absenceDate"` 
}