package models

// Smartphone model structure for smartphone
type Smartphone struct {
	Id                      int64
	Price                   int
	Name, CountryOrigin, Os string
}

// CreateSmartphoneCMD
type CreateSmartphoneCMD struct {
	Price         int    `json:"price"`
	Name          string `json:"name"`
	CountryOrigin string `json:"country_origin"`
	Os            string `json:"os"`
}
