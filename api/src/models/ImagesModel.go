package models

import (
  "fmt"
  "errors"
  "strconv"
  "path/filepath"
  "mime/multipart"
  "github.com/gin-gonic/gin"
  "github.com/satori/go.uuid"
  "github.com/fzetter/dc-final/api/src/utils"
)

// GET Workload
func UploadImage(c *gin.Context, file *multipart.FileHeader, workload_id string, img_type string) (message utils.UploadImageStruct, err error) {

  // ** INPUT VALIDATION
  if (img_type != "original" && img_type != "filtered") {
    return utils.UploadImageStruct{}, errors.New("Error: Invalid Type")
  }

  // ** VERIFY WORKLOAD UUID
  found := false
  uid := uuid.Must(uuid.NewV4()).String()
  workloadName := ""
  imageFilename := ""
  fileExt := filepath.Ext(file.Filename)

  for index, element := range utils.Workloads {
    if element.Workload_Id == workload_id {
      if (element.Status != "scheduling") { return utils.UploadImageStruct{}, errors.New("Error: Cannot Upload to Busy Workload") }
      found = true
      workloadName = element.Workload_Name
      imageFilename = workloadName + "_" + img_type[0:1] + "-" + uid + fileExt
      utils.Workloads[index].Filtered_Images = append(element.Filtered_Images, imageFilename)
    }
  }
  if (found == false) { return utils.UploadImageStruct{}, errors.New("Error: Workload UUID Not Found") }

  // ** SAVE IMAGE
  err = c.SaveUploadedFile(file, "images/" + workloadName + "/" + img_type[0:1] + "-" + uid + fileExt)
	if err != nil {
    fmt.Printf("Error Uploading File", err)
    return utils.UploadImageStruct{}, errors.New("Error: Could Not Upload Image")
  }

  // ** RESPONSE
  res := utils.UploadImageStruct {
    Workload_Id: workload_id,
    Image_Id: imageFilename,
    Type: img_type, // original, filtered
    Filename: file.Filename,
    Extension: fileExt,
    Size: strconv.FormatInt(int64(file.Size/1000), 10) + "kb",
  }

  return res, nil

}
