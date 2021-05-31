package controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/fzetter/dc-final/api/src/utils"
	"github.com/fzetter/dc-final/api/src/models"
)

// POST Workload
func CreateWorkload(c *gin.Context) {

	var body utils.CreateWorkloadStruct
  c.BindJSON(&body)

	val, err := models.CreateWorkload(&body)

	if err != nil {
		fmt.Println(err.Error())
    c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, val)
	}

}

// GET Workload
func GetWorkload(c *gin.Context) {

	val, err := models.GetWorkload(c.Param("workload_id"))

	if err != nil {
		fmt.Println(err.Error())
    c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, val)
	}

}
