package env

import "os"

func GetModelPath() string {
	env := os.Getenv("MODEL_PATH")
	if env == "" {
		return "models"
	}
	return env
}

func GetPort() string {
	env := os.Getenv("PORT")
	if env == "" {
		return ":3000"
	}
	return ":" + env
}
