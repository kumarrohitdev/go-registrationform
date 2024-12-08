package main

import (
	"log"
	"net/http"
	"context"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

//Creating the type of user
type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

var client *mongo.Client
var collection *mongo.Collection


func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func init() {
	mongoURI :="mongodb://localhost:27017"
	
	var err error
	client, err =mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err !=nil{
		log.Fatal(err)
	}


	//Connecting to the mongoDB server
	err = client.Connect(context.TODO())
	if err!= nil{
        log.Fatal(err)
    }

	collection=client.Database("mydatabasego").Collection("users")
}
func main(){
	router :=gin.Default()
	router.Use(cors.Default())

	router.GET("/",HomeRoute)
	router.POST("/register",RegisterUserRoute)
	router.Run(":8080")
}
func HomeRoute(c *gin.Context){

	c.JSON(http.StatusOK,gin.H{
		"status":true,
		"message":"This is my Hello Message for World!",
	})
}

func RegisterUserRoute(c *gin.Context){
	var user User

	//taking data from body
	if err := c.ShouldBindJSON(&user); err !=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"Status":false,
			"message":"Invalid Request",
		})
		return
	}
	hashPassword, err := HashPassword(user.Password)
if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
        "status":  false,
        "message": "Error hashing password",
    })
    return
}

_, insertErr := collection.InsertOne(context.TODO(), bson.M{
    "name":     user.Name,
    "email":    user.Email,
    "password": hashPassword,
})

if insertErr != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
        "status":  false,
        "message": "Failed to insert data",
    })
    return
}

// Return success response
c.JSON(http.StatusOK, gin.H{
    "status":  true,
    "message": "User successfully registered",
    "data":    user,
})
}