// This will define the models used for the DB interaction.
package data

import (
	"gorm.io/gorm"
)

// Product - GORM Model defining the product data entity and associated database table
type Product struct {
	gorm.Model
	Name  string
	Price float32
}

// AddProduct - Inserts product into the database
func AddProduct(db *gorm.DB, name string, price float32) (uint, error) {
	p := Product{Name: name, Price: price}
	res := db.Create(&p)
	if res.Error != nil {
		return 0, res.Error
	}
	return p.ID, nil

}
