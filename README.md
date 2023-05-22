# Scrap some food data for SteW
## Basics
This code uses the package Colly to scrap food names, serving size and calories from https://www.calories.info/food/,
put them into a FoodItem struct outlined below.
```
type FoodItem struct{
	fid int
	name, serving, calories string
}
```
```
 fid int - Identification number of food in database
```
```
 name string - Name of food item
```
```
 serving string - Stored as 100 grams or milliliters depending on food type (grams for pizza, milliliters for orange juice)
```
```
 calories string - Total calories in the food item per serving variable (i.e. 100 grams/milliliters)
```
Each FoodItem is then inserted into the postgres database.

## Why?
This food is intended for personal use in a home assistant called steW. Tracking calories with be an important component of this home assistant and so data
is needed for daily insertion of food eaten.

## Future Work
- Scrap foods from more sources but ensure not to make duplicates
- Add marcos for all foods in database (protein, carbs, fats and possibly mineral, sodium, etc)
- Add functionality to update database with personal food items (may not be done in this repo)

## Important
db.json excluded from this repo due to sensitive nature of passwords (etc) to connect to database.
Example of db.json for a postgres database:
```
{
    "host": "localhost",
    "port": "5000",
    "user": "user1",
    "password": "<password>",
    "dbname": "foods"
}
```
