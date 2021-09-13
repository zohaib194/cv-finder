package model

import (
	"os"
	"log"
	"fmt"
	"encoding/json"
)

var Filepath string

func ReadAllCVs() CVWrapperModel {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	fmt.Print(path + Filepath)
	data, err := os.ReadFile(Filepath)
	if err != nil {
        panic(err)
    }

	cv := CVWrapperModel{}
	json.Unmarshal([]byte(data), &cv)
	return cv;
}