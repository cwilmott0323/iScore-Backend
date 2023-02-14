package main

import (
	"flag"
	"fmt"
	//"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"iScore-api/api"
	"log"
)

func main() {
	env := flag.String("env", "dev", "enter env")
	flag.Parse()

	if *env == "prod" {
		fmt.Println("inside prod")
		err := godotenv.Load("prod.env")
		//lambda.Start(api.Run())
		if err != nil {
			log.Fatalf("Error getting env, %v", err)
		}
	} else {
		err := godotenv.Load("dev.env")
		fmt.Println("inside dev")
		err = api.Run()
		if err != nil {
			log.Fatalf("Error: , %v", err)
		}
		if err != nil {
			log.Fatalf("Error getting env, %v", err)
		}
	}
}
