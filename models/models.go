package models

type Config struct {
	DbURI string `json:"connectUriDb"`
	LogName string `json:"logName"`
	Port string `json:"port"`
}


type Menu struct {
	Id_menu int `gorm:"column:id_menu;primary_key" json:"id_menu"`
	Name  string  `gorm:"column:name" json:"name"`
	Price float64 `gorm:"column:sena" json:"sena"`
	Size  float64 `gorm:"column:razmer" json:"razmer"`
	// Id_vid int `gorm:"column:id_vid" json:"id_vid"`
}

func (Menu) TableName() string {
	return "menu"
}

type Request struct{
	Pizzas []Pizza `json:"pizzas"`
	Date string `json:"date"`
	TotalSum float64 `json:"totalSum"`
}

type Pizza struct{
	Name string `json:"name"`
	Count float64 `json:"count"`
	Sum float64 `json:"sum"`
}

type Trans struct{
	Name string `gorm:"column:id_tovar"`
	Count float64 `gorm:"column:kol-vo"`
	Sum float64 `gorm:"column:summa"`
	Date string `gorm:"column:date"`
}