package main

import (
	"flag"
	//"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"iScore-api/api"
	"log"
)

func main() {
	env := flag.String("env", "prod", "enter env")
	flag.Parse()

	if *env == "prod" {

		err := godotenv.Load("prod.env")
		//lambda.Start(api.Run())
		err = api.Run()
		if err != nil {
			log.Fatalf("Error getting env, %v", err)
		}
	} else {
		err := godotenv.Load("dev.env")

		err = api.Run()
		if err != nil {
			log.Fatalf("Error: , %v", err)
		}
		if err != nil {
			log.Fatalf("Error getting env, %v", err)
		}
	}
}
