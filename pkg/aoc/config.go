package aoc

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AoCConfig struct {
	ProjectPath string
	SessionId   string
}

func GetAoCConfig() AoCConfig {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading the .env file")
	}

	return AoCConfig{
		SessionId:   os.Getenv("AOC_SESSION"),
		ProjectPath: os.Getenv("PWD"),
	}
}

func SetSessionId(id string) {
	env, err := godotenv.Unmarshal(fmt.Sprintf("AOC_SESSION=%v", id))

	if err != nil {
		log.Fatal("error writing the session into the environment")
	}

	err = godotenv.Write(env, "./.env")

	if err != nil {
		log.Fatal("error writing the session into the environment")
	}
}
