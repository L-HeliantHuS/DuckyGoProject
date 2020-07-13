package model

import (
	"fmt"
	"log"
)

func AddData() {
	for i := 0; i < 1000; i++ {
		db := Test{
			Title:  "Hello" + string(i),
			Number: i,
			Alias:  "Alias" + string(i),
		}

		if err := DB.Create(&db).Error; err != nil {
			log.Fatal(err)
		}

		fmt.Println("create success")

	}

}
