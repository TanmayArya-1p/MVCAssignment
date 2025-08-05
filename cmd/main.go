package main

import (
	"fmt"
	"inorder/pkg/models"
	"inorder/pkg/types"
)

func main() {
	it, err := models.GetAllUsers(types.Page{
		Limit:  types.DefaultLimit,
		Offset: types.DefaultOffset,
	})
	if err != nil {
		fmt.Println("HERER")
		panic(err)
	}
	for _, item := range it {
		fmt.Println(item)
	}
}
