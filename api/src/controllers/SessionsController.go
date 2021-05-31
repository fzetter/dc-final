package controllers

import (
  "fmt"
  "time"
  "strings"
  "net/http"
	"github.com/gin-gonic/gin"
  "github.com/fzetter/dc-final/api/src/utils"
  "github.com/fzetter/dc-final/api/src/models"
)

// Login
func Login(c *gin.Context) {

  var body utils.Authentication
  c.BindJSON(&body)

	val, err := models.Login(&body)

	if err != nil {
		fmt.Println(err.Error())
    c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, val)
	}

}

// Logout
func Logout(c *gin.Context) {

  user, _ := c.Get("User")
  account := user.(*utils.ClaimsStruct)

  token, _ := c.Get("Token")
  jwtToken, _ := token.(string)

  // Revoke Token
  utils.Tokens = utils.Remove(utils.Tokens, jwtToken)

  val := utils.MessageStruct {
    Message: "Bye " + account.User + ", your token has been revoked",
  }

	c.JSON(http.StatusOK, val)
}

// Status
func Status(c *gin.Context) {
  user, _ := c.Get("User")
  account := user.(*utils.ClaimsStruct)

  currTime := time.Now().String()
  splitTime := strings.Split(currTime, ".")
  time := splitTime[0]

  val := utils.StatusStruct {
    User: account.User,
    System_Name: "DPIP System",
    Server_Time: time,
    Active_Workloads: utils.Workers.ActiveWorkers,
  }

  c.JSON(http.StatusOK, val)

}
