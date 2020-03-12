package models

type Course struct{
    CourseID          string  `json:"courseID"`
    CourseName        string  `json:"courseName"`
    CoursePlace       string  `json:"coursePlace"`
    CourseAttendance  float64 `json:"courseAttendance"`
    CourseStartTime   string  `json:"courseStartTime"`
    CourseEndTime     string  `json:"courseEndTime"`
    CourseDay         string  `json:"courseDay"`
    Population        int    `json:"population"`
}

type Curriculum struct{
    Class    string  `json:"Class"`
    Mon      string  `json:"Mon"`
    Tue      string  `json:"Tue"`
    Wed      string  `json:"Wed"`
    Thu      string  `json:"Thu"`
    Fri      string  `json:"Fri"`
	/*
	"Class":"一",
    "Mon":"演算法",
    "Tue":"",
    "Wed":"",
    "Thu":"",
    "Fri":""
    */
}

