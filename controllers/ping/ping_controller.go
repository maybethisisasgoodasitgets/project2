package ping



import (

    "net/http"



    "github.com/gin-gonic/gin"

)



func Ping(c *gin.Context) { //Need to send the pointer for the library to know what to do

    c.String(http.StatusOK, "pong")



}