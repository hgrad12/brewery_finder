package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

/*
	struct that contains config.json values
*/
type configurationData struct {
	From       string `json:"from"`
	To         string `json:"to"`
	AccountSID string `json:"account_sid"`
	AuthToken  string `json:"auth_token"`
}

/*
	All messages returned are encapsulated in TwiML for Twilio to read
*/
type TwiML struct {
	XMLName xml.Name `xml:"Response"`
	Say     string   `xml:",omitempty"`
}

/*
	Map of patterns used in regexp objects
*/
var patterns = map[string]string{
	"phone":   "^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\\s\\./0-9]*$",
	"zipCode": "^\\d{5}(?:[-\\s]\\d{4})?$",
	"latLong": "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?),\\s*[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$",
}

/*
	Path to Open Brewery DB
*/
const openBreweryDB_repository string = "https://api.openbrewerydb.org/breweries"

/*
	struct to collect brewery(ies) and place in Breweries array
*/
type Brewery struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	BreweryType    string `json:"brewery_type"`
	Street         string `json:"street"`
	Address2       string `json:"address_2"`
	Address3       string `json:"address_3"`
	City           string `json:"city"`
	CountyProvince string `json:"county_province"`
	State          string `json:"state"`
	PostalCode     string `json:"postal_code"`
	Country        string `json:"country"`
	Longitude      string `json:"longitude"`
	Latitude       string `json:"latitude"`
	Phone          string `json:"phone"`
	WebsiteURL     string `json:"website_url"`
	UpdatedAt      string `json:"updated_at"`
	CreatedAt      string `json:"created_at"`
}

/*
	Query specific brewery(ies)
*/
func FindBrewery(w http.ResponseWriter, r *http.Request) {
	var Breweries []Brewery

	/*
		extract params from URL,
			body: the message sent from the user,
			number: the phone number belonging to the user
	*/
	body, bodyOK := r.URL.Query()["body"]
	number, numberOK := r.URL.Query()["phone_number"]

	if !bodyOK || len(body) == 0 || !numberOK || len(number) == 0 {
		log.Println("Params were not passed!")
		return
	}

	b := body[0]
	n := number[0]

	choice := Validate(b)

	rep := ConstructQuery(choice, b)

	/*
		execute http request to Open Brewery Database
	*/
	response, err := http.Get(rep)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	/*
		execute a Close() to prevent a resource leak
	*/
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&Breweries)

	if err != nil {
		panic(err)
	}

	msg := StructureBreweriesToString(Breweries)

	var config configurationData
	var jsonFile *os.File

	abspath, _ := filepath.Abs("./config.json")
	jsonFile, err = os.Open(abspath)

	if err != nil {
		log.Println("json can not be located!")
		return
	}

	data, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(data, &config)

	defer jsonFile.Close()

	ExecuteMessageToTwilio(config, n, msg)
}

/*
	constuct the Twilio client, that will execute the request to the
	Twilio console
*/
func ExecuteMessageToTwilio(config configurationData, phoneNumber, msg string) {
	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: config.AccountSID,
		Password: config.AuthToken,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo(phoneNumber)
	params.SetFrom(config.From)
	params.SetBody(msg)

	_, err := client.ApiV2010.CreateMessage(params)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SMS sent successfully!")
	}
}

/*
	Return a string contianing a query that will be executed to the Open
	Brewery Databse
*/
func ConstructQuery(choice, body string) string {
	query := "%v?%v=%v&per_page=3&sort=name:asc"

	msg := strings.ReplaceAll(strings.ToLower(body), " ", "_")

	query = fmt.Sprintf(query, openBreweryDB_repository, choice, msg)

	return query
}

/*
	Convert all elements of the array into a string
	If array is empty, return a "No Result Found" message
*/
func StructureBreweriesToString(breweries []Brewery) string {
	if len(breweries) == 0 {
		return "No Results Found"
	}

	str := ""

	for _, x := range breweries {
		phone := ""
		typeOfBrewery := ""

		if len(x.BreweryType) != 0 {
			breweryType := string(x.BreweryType)
			typeOfBrewery = strings.Replace(breweryType, breweryType[:1], strings.ToUpper(breweryType[:1]), 1)
		}

		if len(x.Phone) != 0 {
			phone += "(" + x.Phone[:3] + ") "
			phone += x.Phone[3:6] + "-"
			phone += x.Phone[6:]
		}

		str += fmt.Sprintf("%v\n%v\n%v\n%v, %v %v\n%v\n\n", x.Name, typeOfBrewery, x.Street, x.City, x.State, x.PostalCode, phone)
	}

	return str
}

/*
	Method to determine how to query the brewery list
*/
func Validate(message string) string {
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

	http.HandleFunc("/", FindBrewery)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
