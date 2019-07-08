package main

import (
	"LearningGo/DataAnalysis/CsvReader"
	"fmt"
)

func main() {
	var records map[string]CsvReader.Record

	records = make(map[string]CsvReader.Record)
	records["FirstRecord"] = CsvReader.Record{Value:"First record"}

	document := CsvReader.Document{Title: "THis is a title", Records: records }

	//I dont know if this is how it should be done
	document.GetStringFromPath("DataAnalysis/test.csv")
	result, err := document.GetRecordPointerFromID("FirstRecord")

	HandleError(err)
	fmt.Println(result.Value)
	fmt.Println(document.FileDetails)
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
