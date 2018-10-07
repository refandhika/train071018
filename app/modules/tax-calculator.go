package modules

func FoodAndBeverageTax(price float64) float64 {
	// 10% Price
	return price * 10.0 / 100.0
}

func TobaccoTax(price float64) float64 {
	// 10 + ( 2% Price )
	return 10 + (2.0 / 100.0 * price)
}

func EntertainmentTax(price float64) float64 {
	// 0 < Price < 100 : Tax Free
	if price < 100 {
		return 0
	} else {
		// Price >= 100 : 1% (Price - 100)
		return (1.0 / 100.0) * (price - 100)
	}
}
