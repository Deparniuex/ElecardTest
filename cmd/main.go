package main

import (
	"ElecardTest/internal/api"
	"ElecardTest/internal/calculate"
	"ElecardTest/internal/entity"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	err := setupConfig("config")
	if err != nil {
		panic(err)
	}
	url := viper.GetString("API_URL")
	key := viper.GetString("API_KEY")
	tasks, err := api.GetTasks(url, &entity.Body{
		Key:    key,
		Method: "GetTasks",
		Params: nil,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Tasks: ")
	for _, task := range tasks.Result {
		fmt.Println(task)
	}

	rectangles := calculate.CalcRectangles(tasks)
	results, err := api.CheckResults(url, &entity.Body{
		Key:    key,
		Method: "CheckResults",
		Params: rectangles,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Rectangles: ")
	for _, rectangle := range rectangles {
		fmt.Println(rectangle)
	}
	fmt.Println("Results: ")
	for _, result := range results.Result {
		fmt.Println(result)
	}
}

func setupConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	return viper.ReadInConfig()
}
