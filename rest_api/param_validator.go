package rest_api

import (
	"fmt"
	"time"
)

// CheckCreateBurgerParams validates request parameters
func CheckCreateBurgerParams(req CreateBurgerRequest) error {
	if req.Name == "" {
		return fmt.Errorf("Missing burger name")
	}

	for _, p := range req.Places {
		if p.BurgerInfo.Price == 0 {
			return fmt.Errorf("Missing burger price")
		}
		if p.BurgerInfo.Rating == 0 {
			return fmt.Errorf("Missing burger rating")
		}
		if p.BurgerInfo.Supply <= 0 {
			return fmt.Errorf("Missing burger supply")
		}
		if p.BurgerInfo.Date == "" {
			return fmt.Errorf("Missing burger date")
		}
		_, err := time.Parse("01-2006", p.BurgerInfo.Date)
		if err != nil {
			return fmt.Errorf("Date wrong format. Should be '01-2006'")
		}
	}

	return nil
}
