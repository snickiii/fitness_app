package utils

import (
	"fmt"
	"regexp"
	"strconv"
)

func parseNutritionInfo(description string) (portion string, calories, fat, carbs, protein float64, err error) {
	re := regexp.MustCompile(`Per\s+([\d/]+(\s*\w+)?)\s*-\s*Calories:\s*(\d+)kcal\s*\|\s*Fat:\s*(\d+\.?\d*)g\s*\|\s*Carbs:\s*(\d+\.?\d*)g\s*\|\s*Protein:\s*(\d+\.?\d*)g`)

	matches := re.FindStringSubmatch(description)
	if len(matches) != 7 {
		return "", 0, 0, 0, 0, fmt.Errorf("failed to parse description: %s", description)
	}

	portion = matches[1]
	calories, _ = strconv.ParseFloat(matches[3], 8)
	fat, _ = strconv.ParseFloat(matches[4], 8)
	carbs, _ = strconv.ParseFloat(matches[5], 8)
	protein, _ = strconv.ParseFloat(matches[6], 8)

	return portion, calories, fat, carbs, protein, nil
}
