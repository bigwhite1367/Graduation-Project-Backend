package route

import(
    _"net/http"
    "github.com/gin-gonic/gin"
    controllers "ncu_14_project/backend/Controllers"
    util "ncu_14_project/backend/Util"
)


func Serve(){
    router := gin.New()

    router.Use(gin.Logger())
    router.Use(gin.Recovery())

    router.POST("/login", util.AuthMiddleware.LoginHandler)//login

    auth := router.Group("auth")
    auth.GET("/refresh_token", util.AuthMiddleware.RefreshHandler)//refresh token

    auth.Use(util.AuthMiddleware.MiddlewareFunc())
    
    //

    get := router.Group("/GET")

    get.Use(util.AuthMiddleware.MiddlewareFunc())
    {
	    get.GET("/course/:courseID", controllers.GetCourse) //GET course information
        get.GET("/curriculum", controllers.GetCurriculum)//GET student curriculum
        get.GET("/faceRecord", controllers.GetFaceRecord)//GET the face records by course
        get.GET("/user",controllers.GetUser)//GET UserName and privilege
        get.GET("/absenceList", controllers.GetAbsenceList)//GET list of absence of all course from a professor
    }

    post := router.Group("/POST")

    post.Use(util.AuthMiddleware.MiddlewareFunc())
    {
        post.POST("/courseRecord", controllers.InsertCourseRecord) //POST course start and end time
        post.POST("/leave", controllers.InsertLeave)  //INSERT or UPDATE student leave record 
        post.POST("/picture")
    }

    router.Run(":8080")
}