package models

import (
  "os"
  "errors"
  "github.com/satori/go.uuid"
  "github.com/fzetter/dc-final/api/src/utils"
)

// POST Workload
func CreateWorkload(body *utils.CreateWorkloadStruct) (message utils.WorkloadStruct, err error) {

  // ** INPUT VALIDATION
  if (body.Filter != "grayscale" && body.Filter != "blur") {
    return utils.WorkloadStruct{}, errors.New("Error: Invalid Parameter")
  }

  // ** CHECK WORKLOAD NAME AVAILABILITY
  for _, element := range utils.Workloads {
    if element.Workload_Name == body.Workload_Name {
      return utils.WorkloadStruct{}, errors.New("Error: Workload Name Is Taken")
    }
  }

  // ** CREATE WORKLOAD
  _ = os.Mkdir("images/" + body.Workload_Name, 0755)

  res := utils.WorkloadStruct {
    Workload_Id: uuid.Must(uuid.NewV4()).String(),
    Filter: body.Filter,
    Workload_Name: body.Workload_Name,
    Status: "scheduling", // scheduling, running, completed
    Running_Jobs: 0,
    Filtered_Images: []string{},
  }

  utils.Workloads = append(utils.Workloads, res)

  return res, nil

}

// GET Workload
func GetWorkload(workload_id string) (message utils.WorkloadStruct, err error) {

  // ** SEARCH WORKLOAD BY UUID
  for _, element := range utils.Workloads {
    if element.Workload_Id == workload_id {
      return element, nil
    }
  }

  return utils.WorkloadStruct{}, errors.New("Error: Workload Not Found")

}
