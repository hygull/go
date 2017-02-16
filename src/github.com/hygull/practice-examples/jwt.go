package main

import (
    "fmt"
    jwt "github.com/dgrijalva/jwt-go"
    "io/ioutil"
    "time"
)

func main() {

    privateKey, err := ioutil.ReadFile("keys/app.rsa")
    if err != nil {
        fmt.Println("Error reading private key")
        return
    }

    t := jwt.New(jwt.GetSigningMethod("HS256"))
    t.Claims["AccessToken"] = "bar"
    t.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
    tokenString, err := t.SignedString(privateKey)

    fmt.Println("tokenString: " + tokenString)

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) ([]byte, error) {

        fmt.Printf("parsed token obj: %v\n", token)
        publicKey, err := ioutil.ReadFile("keys/app.rsa.pub")
        if err != nil {
            return nil, fmt.Errorf("Error reading public key")
        }

        return publicKey, nil
    })

    if err == nil && token.Valid {
        //Carry on
    } else {
        fmt.Printf("Token parse error: %v\n", err)
    }
}
Ru