package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var url = "http://127.0.0.1:9000/car/add"

//var url = "http://127.0.0.1:9000/api/car"
var brand = []string{"ford", "toyota", "mazda", "honda", "audi", "bmw", "benz"}
var color = []string{"red", "black", "white", "blue", "green", "blue"}
var city = []string{"taipei", "new taipei", "yilan", "kaohsiung", "taichung", "hsinchu", "taitung"}
var waitgroup sync.WaitGroup

func body() {
	brandStr := `{"brand":"` + brand[getRand(len(brand))] + `"`
	colorStr := `,"color":"` + color[getRand(len(color))] + `"`
	cityStr := `,"city":"` + city[getRand(len(city))] + `"}`
	jsonStr := []byte(brandStr + colorStr + cityStr)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		waitgroup.Done()
	} else {
		defer resp.Body.Close()
		waitgroup.Done()
	}
}

func getRand(len int) int {
	t := time.Now().UnixNano()
	r1 := rand.New(rand.NewSource(t))

	return r1.Int() % len
}

func main() {
	len := 1000

	t1 := time.Now()
	for i := 0; i < len; i++ {
		waitgroup.Add(1)
		go body()
	}
	waitgroup.Wait()
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)

}
