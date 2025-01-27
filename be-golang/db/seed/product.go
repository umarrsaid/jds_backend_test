package seed

import (
	"log"

	"be-golang/app/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedProducts(db *gorm.DB) {
	// Define the products
	products := []entity.Product{
		{ID: uuid.New(), Price: 218, Department: "Outdoors", Product: "Salad"},
		{ID: uuid.New(), Price: 784, Department: "Games", Product: "Chips"},
		{ID: uuid.New(), Price: 652, Department: "Automotive", Product: "Bike"},
		{ID: uuid.New(), Price: 544, Department: "Tools", Product: "Tuna"},
		{ID: uuid.New(), Price: 643, Department: "Toys", Product: "Fish"},
		{ID: uuid.New(), Price: 100, Department: "Baby", Product: "Bacon"},
		{ID: uuid.New(), Price: 584, Department: "Clothing", Product: "Pants"},
		{ID: uuid.New(), Price: 265, Department: "Kids", Product: "Bacon"},
		{ID: uuid.New(), Price: 493, Department: "Health", Product: "Computer"},
		{ID: uuid.New(), Price: 877, Department: "Games", Product: "Soap"},
		{ID: uuid.New(), Price: 293, Department: "Beauty", Product: "Table"},
		{ID: uuid.New(), Price: 546, Department: "Books", Product: "Chair"},
		{ID: uuid.New(), Price: 906, Department: "Automotive", Product: "Salad"},
		{ID: uuid.New(), Price: 870, Department: "Computers", Product: "Car"},
		{ID: uuid.New(), Price: 443, Department: "Movies", Product: "Pants"},
		{ID: uuid.New(), Price: 810, Department: "Toys", Product: "Towels"},
		{ID: uuid.New(), Price: 33, Department: "Health", Product: "Car"},
		{ID: uuid.New(), Price: 803, Department: "Outdoors", Product: "Table"},
		{ID: uuid.New(), Price: 219, Department: "Computers", Product: "Fish"},
		{ID: uuid.New(), Price: 878, Department: "Garden", Product: "Cheese"},
		{ID: uuid.New(), Price: 556, Department: "Games", Product: "Keyboard"},
		{ID: uuid.New(), Price: 38, Department: "Clothing", Product: "Hat"},
		{ID: uuid.New(), Price: 381, Department: "Music", Product: "Keyboard"},
		{ID: uuid.New(), Price: 881, Department: "Outdoors", Product: "Shoes"},
		{ID: uuid.New(), Price: 327, Department: "Movies", Product: "Chicken"},
		{ID: uuid.New(), Price: 538, Department: "Sports", Product: "Salad"},
		{ID: uuid.New(), Price: 621, Department: "Shoes", Product: "Bike"},
		{ID: uuid.New(), Price: 849, Department: "Beauty", Product: "Chips"},
		{ID: uuid.New(), Price: 869, Department: "Games", Product: "Ball"},
		{ID: uuid.New(), Price: 840, Department: "Automotive", Product: "Towels"},
		{ID: uuid.New(), Price: 353, Department: "Garden", Product: "Hat"},
		{ID: uuid.New(), Price: 645, Department: "Sports", Product: "Sausages"},
		{ID: uuid.New(), Price: 304, Department: "Kids", Product: "Shoes"},
		{ID: uuid.New(), Price: 665, Department: "Home", Product: "Computer"},
		{ID: uuid.New(), Price: 179, Department: "Jewelery", Product: "Table"},
		{ID: uuid.New(), Price: 752, Department: "Books", Product: "Pizza"},
		{ID: uuid.New(), Price: 653, Department: "Kids", Product: "Chicken"},
		{ID: uuid.New(), Price: 307, Department: "Games", Product: "Sausages"},
		{ID: uuid.New(), Price: 740, Department: "Garden", Product: "Bike"},
		{ID: uuid.New(), Price: 111, Department: "Games", Product: "Bike"},
		{ID: uuid.New(), Price: 894, Department: "Industrial", Product: "Cheese"},
		{ID: uuid.New(), Price: 421, Department: "Clothing", Product: "Ball"},
		{ID: uuid.New(), Price: 819, Department: "Tools", Product: "Soap"},
		{ID: uuid.New(), Price: 929, Department: "Home", Product: "Bike"},
		{ID: uuid.New(), Price: 231, Department: "Shoes", Product: "Mouse"},
		{ID: uuid.New(), Price: 991, Department: "Tools", Product: "Gloves"},
		{ID: uuid.New(), Price: 211, Department: "Music", Product: "Car"},
		{ID: uuid.New(), Price: 408, Department: "Shoes", Product: "Cheese"},
		{ID: uuid.New(), Price: 137, Department: "Grocery", Product: "Table"},
		{ID: uuid.New(), Price: 505, Department: "Shoes", Product: "Salad"},
		{ID: uuid.New(), Price: 921, Department: "Automotive", Product: "Mouse"},
		{ID: uuid.New(), Price: 146, Department: "Beauty", Product: "Gloves"},
		{ID: uuid.New(), Price: 559, Department: "Automotive", Product: "Hat"},
		{ID: uuid.New(), Price: 96, Department: "Beauty", Product: "Tuna"},
		{ID: uuid.New(), Price: 82, Department: "Health", Product: "Bacon"},
		{ID: uuid.New(), Price: 285, Department: "Garden", Product: "Chips"},
		{ID: uuid.New(), Price: 977, Department: "Home", Product: "Keyboard"},
		{ID: uuid.New(), Price: 973, Department: "Kids", Product: "Car"},
		{ID: uuid.New(), Price: 950, Department: "Kids", Product: "Sausages"},
		{ID: uuid.New(), Price: 180, Department: "Industrial", Product: "Table"},
		{ID: uuid.New(), Price: 276, Department: "Baby", Product: "Gloves"},
		{ID: uuid.New(), Price: 72, Department: "Books", Product: "Bike"},
		{ID: uuid.New(), Price: 140, Department: "Movies", Product: "Tuna"},
		{ID: uuid.New(), Price: 916, Department: "Beauty", Product: "Salad"},
		{ID: uuid.New(), Price: 416, Department: "Garden", Product: "Hat"},
		{ID: uuid.New(), Price: 875, Department: "Clothing", Product: "Bike"},
		{ID: uuid.New(), Price: 51, Department: "Health", Product: "Pizza"},
		{ID: uuid.New(), Price: 779, Department: "Books", Product: "Cheese"},
		{ID: uuid.New(), Price: 773, Department: "Automotive", Product: "Chair"},
		{ID: uuid.New(), Price: 4, Department: "Games", Product: "Computer"},
		{ID: uuid.New(), Price: 355, Department: "Movies", Product: "Chicken"},
		{ID: uuid.New(), Price: 918, Department: "Baby", Product: "Car"},
		{ID: uuid.New(), Price: 13, Department: "Automotive", Product: "Car"},
		{ID: uuid.New(), Price: 5, Department: "Computers", Product: "Fish"},
		{ID: uuid.New(), Price: 576, Department: "Health", Product: "Ball"},
		{Price: 102, Department: "Garden", Product: "Salad", ID: uuid.New()},
		{Price: 169, Department: "Movies", Product: "Towels", ID: uuid.New()},
		{Price: 907, Department: "Outdoors", Product: "Car", ID: uuid.New()},
		{Price: 653, Department: "Automotive", Product: "Tuna", ID: uuid.New()},
		{Price: 485, Department: "Toys", Product: "Pants", ID: uuid.New()},
		{Price: 596, Department: "Tools", Product: "Chicken", ID: uuid.New()},
		{Price: 587, Department: "Toys", Product: "Pizza", ID: uuid.New()},
		{Price: 470, Department: "Clothing", Product: "Mouse", ID: uuid.New()},
		{Price: 764, Department: "Grocery", Product: "Car", ID: uuid.New()},
		{Price: 810, Department: "Clothing", Product: "Sausages", ID: uuid.New()},
		{Price: 782, Department: "Grocery", Product: "Towels", ID: uuid.New()},
		{Price: 395, Department: "Home", Product: "Pants", ID: uuid.New()},
		{Price: 758, Department: "Sports", Product: "Keyboard", ID: uuid.New()},
		{Price: 837, Department: "Games", Product: "Pizza", ID: uuid.New()},
		{Price: 267, Department: "Jewelery", Product: "Soap", ID: uuid.New()},
		{Price: 908, Department: "Kids", Product: "Keyboard", ID: uuid.New()},
		{Price: 924, Department: "Computers", Product: "Mouse", ID: uuid.New()},
		{Price: 644, Department: "Industrial", Product: "Chips", ID: uuid.New()},
		{Price: 866, Department: "Tools", Product: "Fish", ID: uuid.New()},
		{Price: 775, Department: "Jewelery", Product: "Chips", ID: uuid.New()},
		{Price: 584, Department: "Outdoors", Product: "Keyboard", ID: uuid.New()},
		{Price: 940, Department: "Games", Product: "Soap", ID: uuid.New()},
		{Price: 908, Department: "Home", Product: "Gloves", ID: uuid.New()},
		{Price: 226, Department: "Kids", Product: "Tuna", ID: uuid.New()},
		{Price: 979, Department: "Clothing", Product: "Bike", ID: uuid.New()},
	}

	// Insert the data into the database
	for _, product := range products {
		result := db.Create(&product)
		if result.Error != nil {
			log.Fatalf("Error inserting product: %v\n", result.Error)
		}
	}
}
