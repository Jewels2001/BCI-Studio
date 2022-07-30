package util

import (
	"backend/server/types"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ErrHandle(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func DisplayRecording(H types.Recording) {
	fmt.Println("HISTORY=======================================")
	for i := 0; i < H.Samples; i++ {
		fmt.Println("Sample", i)
		fmt.Println("Ch0:", H.Ch0.Data[i][0], H.Ch0.Data[i][1], H.Ch0.Data[i][2], H.Ch0.Data[i][3])
		fmt.Println("Ch1:", H.Ch1.Data[i][0], H.Ch1.Data[i][1], H.Ch1.Data[i][2], H.Ch1.Data[i][3])
		fmt.Println("Ch2:", H.Ch2.Data[i][0], H.Ch2.Data[i][1], H.Ch2.Data[i][2], H.Ch2.Data[i][3])
		fmt.Println("Ch3:", H.Ch3.Data[i][0], H.Ch3.Data[i][1], H.Ch3.Data[i][2], H.Ch3.Data[i][3])
		fmt.Println("")
	}
}
