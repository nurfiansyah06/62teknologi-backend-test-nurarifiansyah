package entity

type Business struct {
	BusinessId int    `gorm:"primary_key" json:"business_id"`
	Name       string `json:"name"`
	Location   string `json:"location"`
	Latitude   int    `json:"latitude"`
	Longitude  int    `json:"longitude"`
	Categories string `json:"categories"`
	ImageLink  string `json:"image_link"`
}