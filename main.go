package main

import (
	"fmt"
	"log"

	"github.com/taylormonacelli/lemondrop"
)

func main() {
	regions, err := lemondrop.GetAllAwsRegions()
	if err != nil {
		log.Fatal(err)
	}

	for _, region := range regions {
		fmt.Println(*region.RegionName)
	}
}
