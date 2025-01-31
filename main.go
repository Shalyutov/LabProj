package main

import (
	"fmt"
	"labproj/dictionary"
	"slices"
)

func main() {
	biomaterials := []dictionary.Biomaterial{
		{"Венозная кровь", "Кровь из вены"},
		{"Капиллярная кровь", "Кровь из пальца"},
	}
	supplies := []dictionary.Supply{
		{"Пробирка для забора венозной крови", "Гранат Био Тех", biomaterials[0], 5.0, 5},
	}
	measures := []dictionary.Measure{
		{"vu/r", "10^-5"},
	}
	indicators := []dictionary.Indicator{
		{"Эритроциты", measures[0], 150.0, 10.0, true},
		{"Тромбоциты", measures[0], 100.0, 60.0, true},
		{"Гемоглобин", measures[0], 200.0, 100.0, true},
	}
	services := []dictionary.Service{
		{"Забор венозной крови", 350.0},
	}
	tests := []dictionary.Test{
		{"Общий анализ крови", []string{"ОАК"}, indicators[0:], services[0:], supplies[0:], false, 200.0},
		{"Анализ крови с лейкоформулой", []string{"Лейкоформула"}, indicators[0:], services[0:], supplies[0:], true, 500.0},
		{"Анализ крови на сахар", []string{"Сахар", "Диабет"}, indicators[0:], services[0:], supplies[0:], false, 100.0},
	}

	sum := 0.0
	consume := make(map[dictionary.Supply]int)
	var served []dictionary.Service

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
}
