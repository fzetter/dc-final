package routes

import (
  "fmt"
  "time"
  "net/http"
  "strings"
  "github.com/dgrijalva/jwt-go"
  "github.com/gin-gonic/gin"
  "github.com/fzetter/dc-final/api/src/utils"
  "github.com/fzetter/dc-final/api/src/controllers"
)

// ******
// ROUTES
// ******

func Init(app *gin.Engine) *gin.Engine {

 // About
 aboutRoutes := app.Group("/about")
 {
  aboutRoutes.GET("/", controllers.About)
 }

 // Images
 imageRoutes := app.Group("/images").Use(Authorization)
 {
  imageRoutes.POST("/", controllers.UploadImage)
  imageRoutes.GET("/:image_id", controllers.GetImage)
 }

 // Sessions
 app.POST("/login", controllers.Login)
 app.DELETE("/logout", Authorization, controllers.Logout)
 app.GET("/status", Authorization, controllers.Status)

  // Workloads
  workloadRoutes := app.Group("/workloads").Use(Authorization)
  {
   workloadRoutes.POST("/", controllers.CreateWorkload)
   workloadRoutes.GET("/", controllers.GetWorkloads)
   workloadRoutes.GET("/:workload_id", controllers.GetWorkload)
  }

 return app
}

// *****
// AUTH
// *****

func Authorization(c *gin.Context) {

  // Retrieve Token From Header
  jwtFromHeader := c.Request.Header.Get("Authorization")
  splitToken := strings.Split(jwtFromHeader, "Bearer ")
  reqToken := splitToken[1]

  // Validate Token And Verify Signature
  token, _ := jwt.ParseWithClaims(
    reqToken,
    &utils.ClaimsStruct{},
    func(token *jwt.Token) (interface{}, error) {
        return utils.JWTKey, nil
    },
  )

  // Parse Data and Verify It Hasn't Been Tempered With
  claims, ok := token.Claims.(*utils.ClaimsStruct)
  if !ok {
    fmt.Println("Couldn't parse claims")
    c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Couldn't parse claims"})
    return
  }

  // Check Expiry Time
  if claims.ExpiresAt < time.Now().UTC().Unix() {
    fmt.Println("JWT is expired")
    c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "JWT is expired"})
    return
  }

  // Check If Token Is Valid
  tokenValid := false
  for _, curr := range utils.Tokens {
    if reqToken == curr {
        tokenValid = true
    }
  }

  if tokenValid == false {
    fmt.Println("Token Is Revoked")
    c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Token Is Revoked"})
    return
  }

  // Save User and Token In Context
  c.Set("User", claims)
  c.Set("Token", reqToken)
  c.Next()

}
