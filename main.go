package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/icd10codes", getICD10Code)

	router.Run("localhost:8080")
}

func getICD10Code(c *gin.Context) {
	resp, err := http.Get(BASE_URL + "/search?sf=code,name&terms=tuberc")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var response []interface{}
	json.Unmarshal(body, &response)
	fmt.Println(response[0])
}
