package model

import (
	"encoding/json"
	"log"
	"os"
)


func ReadAllCVs() CVWrapperModel {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	data, err := os.ReadFile(path + "/db/cv.json");
	if err != nil {
        panic(err)
    }

	cv := CVWrapperModel{}
	json.Unmarshal([]byte(data), &cv)
	return cv;
}