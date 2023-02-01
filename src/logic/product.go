package logic

import (
	"github.com/cureadumitru86/tt-go-api/src/data"
	"gorm.io/gorm"
)

// InsertProduct - Adds a product to the associated database table using the Product model
func InsertProduct(db *gorm.DB, name string, price float32) (uint, error) {
	id, err := data.AddProduct(db, name, price)
	if err != nil {
		return 0, err
	}

	return id, nil
}
