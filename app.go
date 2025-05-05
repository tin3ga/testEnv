package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Vars struct {
	Name      string
	Character string
	Seconds   string
}

func LoadVars() (*Vars, error) {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	vars := &Vars{
		Name:      os.Getenv("NAME"),
		Character: os.Getenv("CHARACTER"),
		Seconds:   os.Getenv("RUN_TIME"),
	}

	return vars, nil
}

func ConvertToInt(str string) (int, error) {
	if str == "" {
		return -1, fmt.Errorf("empty string")
	}
	sec, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error converting seconds to int:", err)
		os.Exit(1)
	}
	return sec, nil

}

func main() {
	vars, err := LoadVars()

	if err != nil {
		os.Exit(1)
	}

	if *vars == (Vars{}) {
		vars = &Vars{
			Name:      "arif",
			Character: "octavia",
			Seconds:   "0",
		}
	}

	fmt.Printf("Hello,"+" %s"+"!", vars.Name)
	fmt.Printf("\nFav Character:"+" %s \n", vars.Character)

	sec, err := ConvertToInt(vars.Seconds)
	if err != nil || sec < 0 {
		fmt.Println("Error converting seconds:", err)
		os.Exit(1)
	}

	fmt.Printf("\rscript will run for %d seconds", sec)

	for i := sec; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Printf("\rscript will exit after %d seconds   ", i)
	}
	fmt.Printf("\rExited!                           \n")
	os.Exit(0)

}
