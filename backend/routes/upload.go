package routes

import (
	"backend/server/types"
	"backend/server/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: /upload/")
	var R types.Recording
	err := json.NewDecoder(r.Body).Decode(&R)
	if err != nil {
		log.Println(err.Error())
		util.RespondWithError(w, http.StatusBadRequest, "invalid fields")
		return
	}

	log.Println("Recieved upload request")

	filename, err := writeToDisk(R)
	if err != nil {
		log.Println(err.Error())
		util.RespondWithError(w, http.StatusInternalServerError, "unable to write to disk")
		return
	}

	log.Println("Saved to", "./data/"+filename+".json")

	util.RespondWithJSON(w, http.StatusAccepted, fmt.Sprintf("File %s saved", filename))
}

func writeToDisk(R types.Recording) (string, error) {
	filename := "recording_" + time.Now().Format("2006_01_02_15:04:05")
	file, err := json.MarshalIndent(R, "", " ")
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile("./data/"+filename+".json", file, 0644)
	if err != nil {
		return "", err
	}

	return filename, nil
}
