package utils

import (
  "time"
  "github.com/dgrijalva/jwt-go"
  "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func Remove(slice []string, val string) []string {
    for i, curr := range slice {
        if curr == val {
            return append(slice[:i], slice[i+1:]...)
        }
    }
    return slice
}

func AdminAuth() (string) {

  claims := ClaimsStruct{
      User: Users[0].User,
      Email: Users[0].Email,
      Role: Users[0].Role,
      StandardClaims: jwt.StandardClaims{
          ExpiresAt: time.Now().Add(60 * time.Minute).Unix(),
          Issuer: "DC",
      },
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  signedToken, _ := token.SignedString(JWTKey)
  Tokens = append(Tokens, signedToken)
  jwt_token := signedToken
  return jwt_token

}
