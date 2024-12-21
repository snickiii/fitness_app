package utils

import (
	"encoding/json"
	"fitness_back/DTO"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func FetchFoodDataFromAPI(searchExp string) (*DTO.FindResponse, error) {

	authorizationToken := os.Getenv("AUTHORIZTION_TOKEN")
	apiBaseURL := os.Getenv("API_URL")
	url := fmt.Sprintf("%s&search_expression=%s", apiBaseURL, searchExp)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Add("Authorization", "Bearer "+authorizationToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var apiResponse struct {
		Foods struct {
			Food []struct {
				FoodDescription string `json:"food_description"`
				FoodID          string `json:"food_id"`
				FoodName        string `json:"food_name"`
				FoodType        string `json:"food_type"`
				FoodURL         string `json:"food_url"`
			} `json:"food"`
		} `json:"foods"`
	}

	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal API response: %v", err)
	}

	var findResponse DTO.FindResponse
	for _, foodItem := range apiResponse.Foods.Food {
		foodID, err := strconv.Atoi(foodItem.FoodID)
		if err != nil {
			log.Printf("Invalid food_id %s, skipping item", foodItem.FoodID)
			continue
		}

		portion, calories, fat, carbs, protein, err := parseNutritionInfo(foodItem.FoodDescription)
		if err != nil {
			log.Printf("Error parsing nutrition info for %s: %v", foodItem.FoodName, err)
			continue
		}

		meal := DTO.FatsecretMeal{
			FoodDescription: foodItem.FoodDescription,
			FoodID:          uint(foodID),
			FoodName:        foodItem.FoodName,
			Portion:         portion,
			Calories:        calories,
			Fat:             fat,
			Carbs:           carbs,
			Protein:         protein,
		}

		findResponse.FindResponse = append(findResponse.FindResponse, meal)
	}

	return &findResponse, nil
}
