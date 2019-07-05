package CsvReader

import (
	"github.com/pkg/errors"
	"io/ioutil"
)

type Document struct {
	//Map of records
	Title 		string
	Records 	map[string]Record	//Need to think about this one
	FileDetails string
}

//GetRecordFromID will return a pointer to the requested record from the ID passed
//Pointer returned because we dont know how big these records will actually be
func (document Document) GetRecordPointerFromID(id string) (*Record, error) {
	result := document.Records[id]
	//Panics if a record cant be found
	if result == (Record{}) {
		//Returning the error up to the caller
		return nil, errors.New("No record found")
	}

	return &result, nil
}

//ReadFileIn will read the file in and return the result of the load
func (document Document) ReadFileIn(path string) bool {
	fileResult, err := getFileString(path)
	if err != nil {
		return false
	}
	document.FileDetails = fileResult
	return true
}


func getFileString(pathToFile string) (string, error) {
	csv, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		return "", errors.New("File couldnt be opened or found")
	}
	return string(csv), err
}