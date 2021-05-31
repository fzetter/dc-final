package utils

import (
  "github.com/dgrijalva/jwt-go"
)

// JWT CLAIMS

type ClaimsStruct struct {
    User string `json:"username" binding:"required"`
    Email string `json:"email" binding:"required"`
    Role string `json:"role" binding:"required"`
    jwt.StandardClaims
}

// IMAGES
type UploadImageStruct struct {
    Workload_Id string `json:"workload_id" binding:"required"`
    Image_Id string `json:"image_id" binding:"required"`
    Filename string `json:"filename" binding:"required"`
    Size string `json:"size" binding:"required"`
    Extension string `json:"extension" binding:"required"`
    Type string `json:"type" binding:"required"`
}

// SESSIONS

type Authentication struct {
    User string `json:"user" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type MessageStruct struct {
    Message string `json:"message" binding:"required"`
    Token string `json:"token"`
    Time string `json:"time"`
}

type UserStruct struct {
    User string `json:"user" binding:"required"`
    Email string `json:"email" binding:"required"`
    Role string `json:"role" binding:"required"`
    Password string `json:"password" binding:"required"`
    Token string `json:"token" binding:"required"`
}

// STATUS

type StatusStruct struct {
    User string `json:"user" binding:"required"`
    System_Name string `json:"system_name" binding:"required"`
    Server_Time string `json:"server_time" binding:"required"`
    Active_Workloads int `json:"active_workloads" binding:"required"`
}

// WORKLOADS

type WorkloadStruct struct {
    Workload_Id string `json:"workload_id" binding:"required"`
    Filter string `json:"filter" binding:"required"`
    Workload_Name string `json:"workload_name" binding:"required"`
    Status string `json:"status" binding:"required"`
    Running_Jobs int `json:"running_jobs" binding:"required"`
    Filtered_Images []string `json:"filtered_images" binding:"required"`
}

type CreateWorkloadStruct struct {
    Filter string `json:"filter" binding:"required"`
    Workload_Name string `json:"workload_name" binding:"required"`
}
