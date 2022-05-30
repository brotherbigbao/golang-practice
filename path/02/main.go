package main

import (
	"fmt"
	"os"
)

func main() {
	const envDotenv = "DOTENV_PATH"
	env, isSuccess := os.LookupEnv(envDotenv)
	if !isSuccess {
		fmt.Println("not get env")
		return
	}
	fmt.Println(env)
}
