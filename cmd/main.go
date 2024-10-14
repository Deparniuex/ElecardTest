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
	tasks, err := api.GetTasks(viper.GetString("API_URL"), &entity.Body{
		Key:    viper.GetString("API_KEY"),
		Method: "GetTasks",
		Params: nil,
	})
	fmt.Printf("%v", tasks)
	if err != nil {
		panic(err)
	}
	rectangles := calculate.CalcRectangles(tasks)
	fmt.Println("Debug rectangles")
	fmt.Println(rectangles)
}

func setupConfig(path string) error {

	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	return viper.ReadInConfig()
}
