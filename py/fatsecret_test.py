from fatsecret import Fatsecret

fs = Fatsecret('c9af633c94e4439ebbbd78a61f2ba95a', '1d21e139848042428ddd49f1bbd50db7')

foods = fs.foods_search('big tasty')

# foods = fs.food_get(61798376)

print(foods)