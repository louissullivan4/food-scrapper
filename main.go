package main

import (
	"fmt"
	"regexp"
	"strconv"
	"github.com/gocolly/colly"
)

type FoodItem struct{
	name, serving, calories string
}

func main() {
	foodArray := scrapeFood()
	fmt.Println(foodArray)	
}

func scrapeFood() []FoodItem {
	c := colly.NewCollector()

	foodURL := "https://www.calories.info/food/"

	categoriesArray := []string{"fruit-juices", "alcoholic-drinks-beverages", "baking-ingredients", 
	"beef-veal", "beer", "cakes-pies", "candy-sweets", "canned-fruit", "cereal-products", "cheese", "cold-cuts-lunch-meat", "cream-cheese", "dishes-meals",
	"fast-food", "fish-seafood", "fruits", "herbs-spices", "ice-cream", "legumes", "meat", "milk-dairy-products", "non-alcoholic-drinks-beverages",
	"nuts-seeds", "oatmeal-muesli-cereals", "offal-giblets", "oils-fats", "pasta-noodles", "pastries-breads-rolls", "pizza", "pork", "potato-products", 
	"poultry-fowl", "sauces-dressings", "sausage", "sliced-cheese", "soda-soft-drinks", "spreads", "tropical-exotic-fruits", "vegetable-oils", "vegetables",
	"venison-game", "wine", "yogurt"}

	var foodArray []FoodItem

	for i := 0; i < len(categoriesArray); i++ {
		var foodItem FoodItem
		c.OnHTML("tr", func(e *colly.HTMLElement){
			// skip the title values
			if ((e.ChildText("td.serving")) != "ServingServingServing") {
				foodItem.name = e.ChildText("a")
				foodItem.serving = "100"
				foodItem.calories = e.ChildText("td.kcal")
				foodArray = append(foodArray, foodItem)
			}
		})
		c.Visit(foodURL + categoriesArray[i])
	}
	return foodArray
}