package routes

import (
	"backend/server/types"
	"backend/server/util"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
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
	vars := mux.Vars(r)
	filename := vars["filename"]
	log.Println("Endpoint hit: /download/" + filename)

}
