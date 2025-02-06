package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	dict "labproj/entities/dictionary"
	"net/http"
	"slices"
	"strconv"
)

func main() {
	biomaterials := []dict.Biomaterial{
		{"Венозная кровь", "Кровь из вены"},
	}
	supplies := []dict.Supply{
		{"Вакуумная пробирка для забора венозной крови", "Гранат Био Тех", biomaterials[0], 5.0, 5},
	}
	integerIndicators := []dict.IntegerIndicator{
		{"Эритроциты (RBC)", "10^12/л", 6.0, 4.5},
		{"Гематокрит (HCT)", "%", 48.0, 40.0},
		{"Гемоглобин (HGB)", "г/л", 180.0, 130.0},
		{"Лейкоциты (WBC)", "10^9/л", 9.0, 4.0},
		{"Тромбоциты (PLT)", "10^9/л", 400.0, 150.0},
		{"Эозинофилы", "%", 0.0, 0.0},
		{"Лимфоциты", "%", 37.0, 19.0},
		{"Моноциты", "%", 11.0, 3.0},
		{"СОЭ по Панченкову", "мм/ч", 10.0, 2.0},
	}
	binaryIndicators := make([]dict.BinaryIndicator, 0)
	stringIndicators := make([]dict.StringIndicator, 0)
	services := []dict.Service{
		{"Забор венозной крови", 350.0},
	}
	tests := []dict.Test{
		{"Общий анализ крови", []string{"ОАК"}, integerIndicators[0:5], binaryIndicators[0:], stringIndicators[0:], services[0:], supplies[0:], false, 200.0},
		{"Лейкоцитарная формула", []string{"Лейкоформула"}, integerIndicators[5:8], binaryIndicators[0:], stringIndicators[0:], services[0:], supplies[0:], false, 150.0},
		{"СОЭ", []string{"Сахар", "Диабет"}, integerIndicators[8:], binaryIndicators[0:], stringIndicators[0:], services[0:], supplies[0:], false, 100.0},
	}

	r := gin.Default()
	r.GET("/tests/:id", func(c *gin.Context) {
		GetTest(c, tests)
	})
	err := r.Run()
	if err != nil {
		return
	}

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
}

func GetTest(c *gin.Context, tests []dict.Test) {
	index, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusNotAcceptable)
		return
	}
	if index >= int64(len(tests)) || index < 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	test := tests[index]
	c.JSON(200, test)
}
