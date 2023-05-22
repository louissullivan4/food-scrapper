package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gocolly/colly"
)

type FoodItem struct{
	fid int
	name, serving, calories string
}

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

func main() {
	// set config values from db.json 
	config, _ := setupConfig()

	foodArray := scrapeFood()
	if len(foodArray) < 1 {
		fmt.Println("Error scraping food items")
	} else {
		// fmt.Println(foodArray)
		dbDone := insertDb(foodArray, config)
		if (dbDone) {
			fmt.Println("Database updated")
		} else {
			fmt.Println("Error updating database")
		}
	}
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
	countFid := 0

	for i := 0; i < len(categoriesArray); i++ {
		var foodItem FoodItem
		c.OnHTML("tr", func(e *colly.HTMLElement){
			// skip the title values
			if ((e.ChildText("td.serving")) != "ServingServingServing") {
				// convert int to string
				foodItem.fid = countFid
				foodItem.name = e.ChildText("a")
				foodItem.serving = "100"
				cals := e.ChildText("td.kcal")
				foodItem.calories = strings.Split(cals, "cal")[0]
				foodArray = append(foodArray, foodItem)
				countFid++
			}
		})
		c.Visit(foodURL + categoriesArray[i])
	}
	return foodArray
}

func insertDb(foodArray []FoodItem, config Config) bool {
	// Table created: CREATE table StewFood ( Fid INT PRIMARY KEY, Name TEXT, Serving TEXT, Calories TEXT,);
	// read in const from db.txt
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Dbname)
	fmt.Println(psqlconn)
	return true
}

func setupConfig() (Config, error) {
	var config Config

	data, err := ioutil.ReadFile("db.json")
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}