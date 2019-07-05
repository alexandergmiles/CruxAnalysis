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
	document.ReadFileIn("DataAnalysis/test.csv")

	result, err := document.GetRecordPointerFromID("FirstRecord")

	HandleError(err)
	fmt.Println(result.Value)
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
