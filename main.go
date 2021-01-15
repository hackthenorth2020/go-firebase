package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"

	jwtMiddleware "github.com/auth0/go-jwt-middleware"

	"google.golang.org/api/option"
)

/*
Go-Firebase Template
- HTTP REST JSON API --> DONE
- Parse and validate Google Firebase Auth Token --> DONE
- create valid routes that work with the Auth service --> DONE
- CRUD (Create, read, update, delete) operations on a SQL database --> TODO
- Get claims from auth token --> TODO
- Set claims on user --> TODO
*/

var (
	app *firebase.App
)

func main() {
	fmt.Println("Starting Server")
	r := gin.Default()

	app, err := initFirebase()
	if err != nil {
		panic(err)
	}

	authMiddleware := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			ctx := context.Background()
			idToken, _ := jwtMiddleware.FromAuthHeader(c.Request)

			// fmt.Println(c.Request, idToken)

			client, err := app.Auth(ctx)
			if err != nil {
				log.Printf("error getting Auth client: %v\n", err)
				c.AbortWithError(401, err)
				return
			}

			token, err := client.VerifyIDToken(ctx, idToken)
			if err != nil {
				log.Printf("error verifying ID token: %v\n", err)
				c.AbortWithError(401, err)
				return
			}

			log.Printf("Verified ID token: %v\n", token)
			c.Set("token", token)
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/auth", authMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": c.MustGet("token"),
		})
	})

	r.GET("/users", authMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"users": "[emily, allen, aman, harp]",
		})
	})

	r.Run(":8081") // listen and serve on 0.0.0.0:8081 (for windows "localhost:8081")
}

func initFirebase() (*firebase.App, error) {
	opt := option.WithCredentialsFile("secrets/vue-firebase-key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}
