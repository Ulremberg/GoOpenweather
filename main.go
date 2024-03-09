package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Weather struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`		
		Humidity float64 `json:"humidity"`
	}
}

const (
	LIMIT_LOW_HUMIDITY = 30.0
	LIMIT_HIGH_TEMP    = 25.0
	LIMIT_LOW_TEMP     = 10.0
)

var data Weather

func getStatus(temp, humidity float64) string {
	switch {
	case humidity < LIMIT_LOW_HUMIDITY:
		return "Low Humidity"
	case temp > LIMIT_HIGH_TEMP:
		return "High Temperature"
	case temp < LIMIT_LOW_TEMP:
		return "Intense Cold"
	default:
		return "No Risk"
	}
}


func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
  }

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>" + "Weather Application" + "<h1>"))
	w.Write([]byte("<h2>" + "Router:" + "<h2>"))
	w.Write([]byte("<h3>" + "/name/:name" + "<h3>"))
	w.Write([]byte("<h3>" + "/coords/:lat/:lon" + "<h3>"))
}

func WeatherDataCity(w http.ResponseWriter, r *http.Request) {
	// retreiving city name from request
	params := mux.Vars(r)
	city := params["city"]
	key := goDotEnvVariable("APIKEY")
	units := goDotEnvVariable("UNIT")
	url := goDotEnvVariable("URL")

	urlFinal := url + "q=" + city + "&appid=" + key + "&units=" + units + "&lang=pt_br" 	

	client := &http.Client{}

    req, err := http.NewRequest("GET", urlFinal, nil)
    if err != nil {
        fmt.Println("Error while retrieving site", err)
    }

    req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error while retrieving site", err)
    }

    defer resp.Body.Close()   
	


	// fetched data is in json format need to convert it into struct to process
	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		log.Fatal(err)
	}

	// assigning fetched data into variables
	
	city = data.Name
	temp := fmt.Sprintf("%.2f",data.Main.Temp);
	
	humidity := fmt.Sprintf("%.2f",data.Main.Humidity);

	status := getStatus(data.Main.Temp, data.Main.Humidity);
	
	// print statements for server
	w.Write([]byte("<h1>" + "Fetched Weather Data is " + "<h1>"))	
	w.Write([]byte("<h2>" + "city is : " + city + "<h2>"))
	w.Write([]byte("<h3>" + "temp is : " + temp + " celsius" + "<h3>"))	
	w.Write([]byte("<h3>" + "humidity is : " + humidity + " %" + "<h3>"))
	w.Write([]byte("<h3>" + "status is : " + status  + "<h3>"))
}

func WeatherDataCoords(w http.ResponseWriter, r *http.Request) {
	// retreiving city name from request
	params := mux.Vars(r)
	lat := params["lat"]
	lon := params["lon"]
	key := goDotEnvVariable("APIKEY")
	units := goDotEnvVariable("UNIT")
	url := goDotEnvVariable("URL")

	urlFinal := url + "lat=" + lat + "&lon=" + lon + "&appid=" + key + "&units=" + units + "&lang=pt_br" 	

	client := &http.Client{}

    req, err := http.NewRequest("GET", urlFinal, nil)
    if err != nil {
        fmt.Println("Error while retrieving site", err)
    }

    req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error while retrieving site", err)
    }

    defer resp.Body.Close()   
	


	// fetched data is in json format need to convert it into struct to process
	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		log.Fatal(err)
	}

	// assigning fetched data into variables
	
	city := data.Name;
	status := getStatus(data.Main.Temp, data.Main.Humidity);
	temp := fmt.Sprintf("%.2f",data.Main.Temp);
	
	humidity := fmt.Sprintf("%.2f",data.Main.Humidity);

	
	// print statements for server
	w.Write([]byte("<h1>" + "Fetched Weather Data is " + "<h1>"))	
	w.Write([]byte("<h2>" + "city is : " + city + "<h2>"))
	w.Write([]byte("<h3>" + "temp is : " + temp + " celsius" + "<h3>"))	
	w.Write([]byte("<h3>" + "humidity is : " + humidity + " %" + "<h3>"))
	w.Write([]byte("<h3>" + "status is : " + status  + "<h3>"))
}

func main()  {
	
	fmt.Println("Weather Application")

	// initialising router
	router := mux.NewRouter()
	// handler to route request
	router.HandleFunc("/", serveHome)
	router.HandleFunc("/name/{city}", WeatherDataCity)
	router.HandleFunc("/coords/{lat}/{lon}", WeatherDataCoords)
	http.ListenAndServe(":4000", router)
	
}