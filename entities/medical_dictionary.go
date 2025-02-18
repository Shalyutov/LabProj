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
		{"Венозная кровь", "Кровь из вены", 1},
	}
	supplies := []dict.Supply{
		{"Вакуумная пробирка для забора венозной крови", "Гранат Био Тех", biomaterials[0], 5.0, 5, 1},
	}
	integerIndicators := []dict.IntegerIndicator{
		{"Эритроциты (RBC)", "10^12/л", 6.0, 4.5, 1},
		{"Гематокрит (HCT)", "%", 48.0, 40.0, 2},
		{"Гемоглобин (HGB)", "г/л", 180.0, 130.0, 3},
		{"Лейкоциты (WBC)", "10^9/л", 9.0, 4.0, 4},
		{"Тромбоциты (PLT)", "10^9/л", 400.0, 150.0, 5},
		{"Эозинофилы", "%", 0.0, 0.0, 6},
		{"Лимфоциты", "%", 37.0, 19.0, 7},
		{"Моноциты", "%", 11.0, 3.0, 8},
		{"СОЭ по Панченкову", "мм/ч", 10.0, 2.0, 9},
	}
	binaryIndicators := make([]dict.BinaryIndicator, 0)
	stringIndicators := make([]dict.StringIndicator, 0)
	services := []dict.Service{
		{"Забор венозной крови", 350.0, 1},
	}
	tests := []dict.Test{
		{"Общий анализ крови", []string{"ОАК"}, integerIndicators[0:5], binaryIndicators[0:], stringIndicators[0:], services[0:], supplies[0:], false, 200.0, 1},
		{"Лейкоцитарная формула", []string{"Лейкоформула"}, integerIndicators[5:8], binaryIndicators[0:], stringIndicators[0:], services[0:], supplies[0:], false, 150.0, 2},
		{"СОЭ", []string{"Сахар", "Диабет"}, integerIndicators[8:], binaryIndicators[0:], stringIndicators[0:], services[0:], supplies[0:], false, 100.0, 3},
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
