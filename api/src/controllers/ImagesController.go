package controllers

import (
  "fmt"
  "strings"
  "net/http"
  "path/filepath"
	"github.com/gin-gonic/gin"
  "github.com/fzetter/dc-final/api/src/models"
)

// Upload Image
func UploadImage(c *gin.Context) {

  // File Upload
  file, err := c.FormFile("data")
  workload_id := c.PostForm("workload_id")
  img_type := c.PostForm("type")

  // No File Received
  if err != nil {
      c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "No File Received"})
      return
  }

	val, err := models.UploadImage(c, file, workload_id, img_type)

	if err != nil {
		fmt.Println(err.Error())
    c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, val)
	}

}

// Get Image
func GetImage(c *gin.Context) {
  
  file := strings.Split(c.Param("image_id"), "_")
  targetPath := filepath.Join("images/" + file[0], file[1])
  //c.FileAttachment(targetPath, "filename.png")
  c.File(targetPath)

}
