package models

type NameData struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	National   string `json:"national"`
}

type NationalApiModel struct {
	National []CountryModel `json:"country"`
}

type CountryModel struct {
	CountryID string `json:"country_id"`
}
