package storage

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type URL struct {
	ID          uint   `gorm:"primaryKey`
	ShortURL    string `gorm:"unique"`
	OriginalURL string
	CreatedAt   string
	VisionCount int
	DeviceStats JSONMap `gorm:"type:json"`
}

type JSONMap map[string]int

func (j JSONMap) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan десериализует строку JSON из базы в JSONMap
func (j *JSONMap) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, j)
}

func ConnectDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("url_shortener.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&URL{})
	return db, nil
}
