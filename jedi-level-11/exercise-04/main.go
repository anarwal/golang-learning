package main

import (
	"fmt"
	"log"
)

// Starting with: https://play.golang.org/p/wlEM1tgfQD,
// use the sqrt.Error struct as a value of type error.
// If you would like, use these numbers for your solution
//   - lat "50.2289 N"
//   - long "99.4656 W"

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("\n math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		sqrerror := fmt.Errorf("\n can not find sqrt of negative number: %v", f)
		return 0, sqrtError{"50.2289 N", "99.4656 W", sqrerror}
	}
	return 42, nil
}
