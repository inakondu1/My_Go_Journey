package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {

    sucess := successRate/100.0

    result := float64(productionRate) * sucess
    return result
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    success := successRate/100.0

    result1 := float64(productionRate) * success
    v := int(result1)

    return v/60
	//panic("CalculateWorkingCarsPerMinute not implemented")
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	//panic("CalculateCost not implemented")
    great := (((carsCount/10) * 95000) + ((carsCount % 10) * 10000)) 
    result := uint(great)
    return result
}
