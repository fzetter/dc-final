package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// POST Workload
func CreateWorkload(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
			"workload_id": "xxx",
			"filter": "xxx",
			"workload_name": "xxx",
      "status": "xxx",
      "running_jobs": 0,
      "filtered_images": "xxx",
	})
}

// GET Workload
func GetWorkload(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
			"workload_id": c.Param("workload_id"),
			"filter": "xxx",
			"workload_name": "xxx",
      "status": "xxx",
      "running_jobs": 0,
      "filtered_images": "xxx",
	})
}
