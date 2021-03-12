package helpers

// Burger entity
type Burger struct {
	Name  string  `json:"burger_name"`
	Place []Place `json:"place"`
}

// Place entity
type Place struct {
	Name       string `json:"place_name"`
	BurgerInfo Info   `json:"burger_info"`
}

// Info entity
type Info struct {
	Name   string  `json:"name,omitempty"`
	Price  float64 `json:"price"`
	Supply int     `json:"supply"`
	Date   string  `json:"date"`
	Rating float64 `json:"rating"`
}
