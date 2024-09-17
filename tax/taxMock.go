package main

func main() {
	calcItemCostIncludingTax(10, "", "'")
}

func calcItemCostIncludingTax(price float64, productType string, sellerCity string, buyerCity string, isEcommerce bool) float64 {
	var city = sellerCity
	if isEcommerce {
		city = buyerCity
	}
	stateTax := GetStateTaxByCity(city, productType) //California has a 7.25%
	cityTax := GetCityTax(city, productType)         //Los Angeles additional 2.25% for some products
	federalTax := GetFederalTax(city, productType)   //gasoline etc
	//but must pass productType cause some products have Exemptions per state/city
	//and some have district taxes for transportation etc

	return price * (100 + stateTax + cityTax + federalTax) / 100
}

func GetStateTaxByCity(city, productType) float64 {
	return 1
}

func GetCityTax(city, productType) float64 {
	return 1
}

func GetFederalTax(city, productType) float64 {
	return 1
}

func calcItemTaxIsraeliStyle(price float64, productType string, city string) float64 {
	var taxRate = 17.    //from db…
	if city == "Eilat" { //in list of tax free cities
		return 0
	} else if productType == "vegetable" || productType == "fruit" { //list of tax free types
		return 0
	} else {
		return price - price/(100+taxRate)
	}
}
