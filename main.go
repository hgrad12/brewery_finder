package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type configurationData struct {
	From       string `json:"from"`
	To         string `json:"to"`
	AccountSID string `json:"account_sid"`
	AuthToken  string `json:"auth_token"`
}

var patterns = map[string]string{
	"phone":   "^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\\s\\./0-9]*$",
	"zipCode": "^\\d{5}(?:[-\\s]\\d{4})?$",
	"latLong": "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?),\\s*[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$",
}

const openBreweryDB_repository string = "https://api.openbrewerydb.org/breweries"

type Brewery struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Brewery_Type    string `json:"brewery_type"`
	Street          string `json:"street"`
	Address_2       string `json:"address_2"`
	Address_3       string `json:"address_3"`
	City            string `json:"city"`
	County_Province string `json:"county_province"`
	State           string `json:"state"`
	Postal_Code     string `json:"postal_code"`
	Country         string `json:"country"`
	Longitude       string `json:"longitude"`
	Latitude        string `json:"latitude"`
	Phone           string `json:"phone"`
	Website_URL     string `json:"website_url"`
	Updated_At      string `json:"updated_at"`
	Created_at      string `json:"created_at"`
}

var Breweries []Brewery

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Brewery Listing!")
	fmt.Println("Endpoint Hit: listing")
}

func handleRequests() {
	http.HandleFunc(openBreweryDB_repository, homePage)

	//return all breweries
	http.HandleFunc("/breweries", returnAllBreweries)
	log.Fatal(http.ListenAndServe(":1000", nil))
}

func returnAllBreweries(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllBreweries")
	json.NewEncoder(w).Encode(Breweries)
}

/*
	Keep-Alive: re-use the same underlying TCP connection when sending multiple HTTP Requests/Responses.
*/
func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

func validate(message string) string {
	/*
		Determine if message contains a zipcode
	*/
	r, _ := regexp.Compile(patterns["zipCode"])
	if r.MatchString(message) {
		return "by_postal"
	}

	/*
		Determine if message contains a latitude longitude coordinate
	*/
	r, _ = regexp.Compile(patterns["latLong"])
	if r.MatchString(message) {
		return "by_dist"
	}

	/*
		Name is the default determination
	*/
	return "by_name"
}

func main() {
	fmt.Println("Start of something great!")
	// handleRequests()

	if len(os.Args) == 1 {
		fmt.Println("No message included")
		os.Exit(1)
	}

	message := strings.Join(os.Args[1:], " ")
	fmt.Println(message)
	var config configurationData

	abspath, _ := filepath.Abs("config.json")
	jsonFile, err := os.Open(abspath)

	if err != nil {
		fmt.Println("Config file does not exist")
		os.Exit(1)
	}

	defer jsonFile.Close()

	data, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(data, &config)

	choice := validate(message)
	fmt.Println(choice)

	rep := fmt.Sprintf("%v?%v=%v&per_page=3&sort=name:desc", openBreweryDB_repository, choice, strings.ReplaceAll(strings.ToLower(message), " ", "_"))
	fmt.Println(rep)
	response, err := http.Get(rep)

	//from := config.From
	//to := config.To
	//accountSid = config.AccountSID
	//authToken := config.AuthToken

	//resp, err := http.Get(openBreweryDB_repository)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}

/*
	Regular Expressions

	Phone Number:

		All						-		/^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\s\./0-9]*$/g
		9999999999				-		/\d{10}/g
		(999) 999-9999			-


	State:

		CO						-		/\b\w{2}\b/g
		Denver, CO				-		/([A-Za-z]+(?: [A-Za-z]+)*),? ([A-Za-z]{2})/g


	Zip Code:

		All						-		/^\d{5}(?:[-\s]\d{4})?$/g
		11111					-		/\b\d{5}\b/g


	Latitude Longitude

		All						-		^[-+]?([1-8]?\d(\.\d+)?|90(\.0+)?),\s*[-+]?(180(\.0+)?|((1[0-7]\d)|([1-9]?\d))(\.\d+)?)$
*/
