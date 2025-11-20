package entities

import (
	"fmt"
	dict "labproj/entities/dictionary"
	"slices"
)

type MedicalDictionary struct {
	Biomaterials      []dict.Biomaterial
	Supplies          []dict.Supply
	IntegerIndicators []dict.IntegerIndicator
	BinaryIndicators  []dict.BinaryIndicator
	StringIndicators  []dict.StringIndicator
	Services          []dict.Service
	Tests             []dict.Test
}

func TemplateMedicalDictionary() MedicalDictionary {
	biomaterials := []dict.Biomaterial{
		{FullName: "Венозная кровь", ShortName: "Кровь из вены", Id: 1},
	}
	supplies := []dict.Supply{
		{Name: "Вакуумная пробирка для забора венозной крови", Supplier: "Гранат Био Тех", Biomaterial: biomaterials[0], Volume: 5.0, TestCapacity: 5, Id: 1},
	}
	integerIndicators := []dict.IntegerIndicator{
		{Name: "Эритроциты (RBC)", Measure: "10^12/л", High: 6.0, Low: 4.5, Id: 1},
		{Name: "Гематокрит (HCT)", Measure: "%", High: 48.0, Low: 40.0, Id: 2},
		{Name: "Гемоглобин (HGB)", Measure: "г/л", High: 180.0, Low: 130.0, Id: 3},
		{Name: "Лейкоциты (WBC)", Measure: "10^9/л", High: 9.0, Low: 4.0, Id: 4},
		{Name: "Тромбоциты (PLT)", Measure: "10^9/л", High: 400.0, Low: 150.0, Id: 5},
		{Name: "Эозинофилы", Measure: "%", High: 0.0, Low: 0.0, Id: 6},
		{Name: "Лимфоциты", Measure: "%", High: 37.0, Low: 19.0, Id: 7},
		{Name: "Моноциты", Measure: "%", High: 11.0, Low: 3.0, Id: 8},
		{Name: "СОЭ по Панченкову", Measure: "мм/ч", High: 10.0, Low: 2.0, Id: 9},
	}
	binaryIndicators := make([]dict.BinaryIndicator, 0)
	stringIndicators := make([]dict.StringIndicator, 0)
	services := []dict.Service{
		{Name: "Забор венозной крови", Price: 350.0, Id: 1},
	}
	tests := []dict.Test{
		{Name: "Общий анализ крови", Aliases: []string{"ОАК"}, IntegerIndicators: integerIndicators[0:5], BinaryIndicators: binaryIndicators[0:], StringIndicators: stringIndicators[0:], Services: services[0:], Cases: supplies[0:], IsSeparated: false, Price: 200.0, Id: 1},
		{Name: "Лейкоцитарная формула", Aliases: []string{"Лейкоформула"}, IntegerIndicators: integerIndicators[5:8], BinaryIndicators: binaryIndicators[0:], StringIndicators: stringIndicators[0:], Services: services[0:], Cases: supplies[0:], IsSeparated: false, Price: 150.0, Id: 2},
		{Name: "СОЭ", Aliases: []string{"Сахар", "Диабет"}, IntegerIndicators: integerIndicators[8:], BinaryIndicators: binaryIndicators[0:], StringIndicators: stringIndicators[0:], Services: services[0:], Cases: supplies[0:], IsSeparated: false, Price: 100.0, Id: 3},
	}
	return MedicalDictionary{biomaterials, supplies, integerIndicators, binaryIndicators, stringIndicators, services, tests}
}

func (m MedicalDictionary) Calculate(tests []dict.Test) float64 {
	sum := 0.0
	consume := make(map[dict.Supply]int)
	var served []dict.Service

	for _, test := range tests {
		supply := test.Cases[0]

		if _, ok := consume[supply]; !ok {
			consume[supply] = 0
		}

		if test.IsSeparated {
			consume[supply] += supply.TestCapacity
		} else {
			consume[supply] += 1
		}

		for _, service := range test.Services {
			if !slices.Contains(served, service) {
				served = append(served, service)
			}
		}

		sum += test.Price
	}
	for _, service := range served {
		sum += service.Price
	}

	fmt.Println("Анализы на сумму: ", sum, " рублей")
	fmt.Println("\nДля взятия потребуется:")
	for supply, consumption := range consume {
		count := consumption / supply.TestCapacity
		if consumption%supply.TestCapacity > 0 {
			count += 1
		}
		fmt.Println(supply.Name, "\t", supply.Volume, " мл", "\t", count)
	}
	return sum
}
