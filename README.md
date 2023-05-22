# Scrap some food data for SteW
## Basics
This code uses the package Colly to scrap food names, serving size and calories from https://www.calories.info/food/,
put them into a FoodItem struct and then insert each food into our personal Postgres database. This food is intended for personal
use for a home assistant called steW. Tracking calories with be an important compoent of this home assistant and so data
is needed for daily insertion of food eaten.

## Future Work
-> Scrap foods from more sources but ensure not to make duplicates
-> Add marcos for all foods in database (protein, carbs, fats and possibly mineral, sodium, etc)
-> Add functionality to update database with personal food items (may not be done in this repo)

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