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

const (
	defaultName      = "arif"
	defaultCharacter = "octavia"
	defaultRunTime   = 0
	maxRunTime       = 3600
)

func LoadVars() *Vars {
	return &Vars{
		Name:      os.Getenv("NAME"),
		Character: os.Getenv("CHARACTER"),
		Seconds:   os.Getenv("RUN_TIME"),
	}
}

func ConvertToInt(str string) (int, error) {
	if str == "" {
		return 0, fmt.Errorf("empty string")
	}
	sec, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("invalid integer %q: %w", str, err)
	}
	return sec, nil

}

func NormalizeVars(vars *Vars) *Vars {
	normalized := *vars
	if normalized.Name == "" {
		normalized.Name = defaultName
	}
	if normalized.Character == "" {
		normalized.Character = defaultCharacter
	}
	if normalized.Seconds == "" {
		normalized.Seconds = strconv.Itoa(defaultRunTime)
	}
	return &normalized
}

func ValidateRunTime(sec int) error {
	if sec < 0 {
		return fmt.Errorf("RUN_TIME must be >= 0, got %d", sec)
	}
	if sec > maxRunTime {
		return fmt.Errorf("RUN_TIME must be <= %d, got %d", maxRunTime, sec)
	}
	return nil
}

func main() {
	vars := NormalizeVars(LoadVars())

	sec, err := ConvertToInt(vars.Seconds)
	if err != nil {
		fmt.Printf("Invalid RUN_TIME: %v\n", err)
		os.Exit(1)
	}
	if err := ValidateRunTime(sec); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Hello,"+" %s"+"!", vars.Name)
	fmt.Printf("\nFav Character:"+" %s \n", vars.Character)
	fmt.Printf("\rscript will run for %d seconds", sec)

	for i := sec; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Printf("\rscript will exit after %d seconds   ", i)
	}
	fmt.Printf("\rExited!                           \n")
}
