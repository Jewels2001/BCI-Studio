package routes

import (
	"backend/server/types"
	"backend/server/util"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func GetFiles(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: /files/")
	var FD types.FileDirectory
	err := filepath.Walk("./data/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return nil
		}

		if !info.IsDir() && filepath.Ext(path) == ".json" {
			path = strings.TrimPrefix(path, "data/")
			path = strings.TrimSuffix(path, ".json")
			FD.Files = append(FD.Files, path)
		}

		return nil
	})
	if err != nil {
		log.Println("error reading data directory")
		util.RespondWithError(w, http.StatusInternalServerError, "error reading filenames from disk")
		return
	}

	util.RespondWithJSON(w, http.StatusAccepted, FD)
}

func Download(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")
	log.Println("Endpoint hit: /download?filename=" + filename)

	file, err := os.Open("./data/" + filename + ".json")
	if err != nil {
		log.Println(err.Error())
		util.RespondWithError(w, http.StatusBadRequest, "file not found")
		return
	}

	var R types.Recording
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err.Error())
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	if err = json.Unmarshal(bytes, &R); err != nil {
		log.Println(err.Error())
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	util.RespondWithJSON(w, http.StatusAccepted, R)
}
