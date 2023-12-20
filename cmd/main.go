package main

import (
	"fmt"
	"my_project/internal/sort_filt_service"
)

func main() {
	result := sort_filt_service.GetResultData()

	fmt.Println(result)
}
