package entities

import (
	"fmt"
	dict "labproj/entities/dictionary"
	"slices"
)

type MedicalDictionary struct {
	Biomaterials []dict.Biomaterial
	Supplies     []dict.Supply
	Indicators   []dict.Indicator
	Services     []dict.Service
	Tests        []dict.Test
	Equipment    []dict.Eqiupment
}

func TemplateMedicalDictionary() MedicalDictionary {
	equipment := []dict.Eqiupment{
		{Name: "Test System", Id: 1},
	}
	biomaterials := []dict.Biomaterial{
		{FullName: "Венозная кровь", ShortName: "Кровь из вены", Id: 1},
	}
	supplies := []dict.Supply{
		{Name: "Вакуумная пробирка для забора венозной крови", Supplier: "Гранат Био Тех", Biomaterial: biomaterials[0], Volume: 5.0, TestCapacity: 5, Id: 1},
	}
	indicators := []dict.Indicator{
		{Name: "Эритроциты (RBC)", Measure: "10^12/л", HighReference: 6.0, LowReference: 4.5, Id: 1},
		{Name: "Гематокрит (HCT)", Measure: "%", HighReference: 48.0, LowReference: 40.0, Id: 2},
		{Name: "Гемоглобин (HGB)", Measure: "г/л", HighReference: 180.0, LowReference: 130.0, Id: 3},
		{Name: "Лейкоциты (WBC)", Measure: "10^9/л", HighReference: 9.0, LowReference: 4.0, Id: 4},
		{Name: "Тромбоциты (PLT)", Measure: "10^9/л", HighReference: 400.0, LowReference: 150.0, Id: 5},
		{Name: "Эозинофилы", Measure: "%", HighReference: 0.0, LowReference: 0.0, Id: 6},
		{Name: "Лимфоциты", Measure: "%", HighReference: 37.0, LowReference: 19.0, Id: 7},
		{Name: "Моноциты", Measure: "%", HighReference: 11.0, LowReference: 3.0, Id: 8},
		{Name: "СОЭ по Панченкову", Measure: "мм/ч", HighReference: 10.0, LowReference: 2.0, Id: 9},
	}
	services := []dict.Service{
		{Name: "Забор венозной крови", Price: 350.0, Id: 1},
	}
	tests := []dict.Test{
		{Name: "Общий анализ крови", Aliases: []string{"ОАК"}, Indicators: indicators[0:5], Services: services[0:], Cases: supplies[0:], IsSeparated: false, Price: 200.0, Id: 1},
		{Name: "Лейкоцитарная формула", Aliases: []string{"Лейкоформула"}, Indicators: indicators[5:8], Services: services[0:], Cases: supplies[0:], IsSeparated: false, Price: 150.0, Id: 2},
		{Name: "СОЭ", Aliases: []string{"Сахар", "Диабет"}, Indicators: indicators[8:], Services: services[0:], Cases: supplies[0:], IsSeparated: false, Price: 100.0, Id: 3},
	}
	return MedicalDictionary{biomaterials, supplies, indicators, services, tests, equipment}
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
