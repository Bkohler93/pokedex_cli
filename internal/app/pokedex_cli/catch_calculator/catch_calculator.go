package catchcalculator

import (
	"math"
	"math/rand"
)


const (
	equationSlope float64 = -0.17132867 
	equationYIntercept float64 = 105.167832 
	maxNum = 100
)

func EvaluateCatchOutcome(baseExperience int) bool {

	catchChance := calculateCatchChance(baseExperience)

	num := rand.Intn(maxNum + 1)	

	return catchChance > num	
}


func calculateCatchChance(baseExperience int) int {

	catchChanceFloat := equationSlope * float64(baseExperience) + equationYIntercept 
	return int(math.Round(catchChanceFloat))
}