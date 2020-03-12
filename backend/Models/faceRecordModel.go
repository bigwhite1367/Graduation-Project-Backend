package models

import "database/sql"

type FaceRecord struct{
	CourseID      string            `json:"courseID"`
	Created_at    string            `json:"created_at"`
	CourseName    string            `json:"courseName"`
	Status        string            `json:"status"`
	Total         sql.NullString    `json:"total"`   
}