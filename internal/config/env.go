package config

import (
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {

    DB_PORT string
    DB_HOST string
    DB_USER string
    DB_NAME string
    DB_PASSWORD string
}

// NewEnv initializes and returns a new instance of the Env struct.
func NewEnv() *Env {
	env := Env{}
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	err := envMapToStruct(&env)
	if err != nil {
		log.Fatalf("Error populating Env struct: %v", err)
	}

	return &env
}

// envMapToStruct populates a struct with values from environment variables using reflection.
func envMapToStruct(envStruct interface{}) error {
	// Iterate through the fields of the struct and set their values from environment variables
	structValue := reflect.ValueOf(envStruct).Elem()
	structType := structValue.Type()

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		envKey := field.Tag.Get("mapstructure")

		if envKey != "" {
			envValue := os.Getenv(envKey)
			fieldType := field.Type

			switch fieldType.Kind() {
			case reflect.String:
				structValue.Field(i).SetString(envValue)
			case reflect.Int:
				intValue, err := strconv.Atoi(envValue)
				if err != nil {
					return err
				}
				structValue.Field(i).SetInt(int64(intValue))
				// Add cases for other supported data types as needed
			}
		}
	}

	return nil
}
