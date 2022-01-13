package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

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

	m := r.URL.Path[1:]

	choice := validate(m)

	rep := fmt.Sprintf("%v?%v=%v&per_page=3&sort=name:asc", openBreweryDB_repository, choice, strings.ReplaceAll(strings.ToLower(m), " ", "_"))

	response, err := http.Get(rep)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&Breweries)

	if err != nil {
		panic(err)
	}

	msg := structureBreweriesToString(Breweries)

	twiml := TwiML{Say: msg}

	breweryXml, err := xml.Marshal(twiml)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")

	w.Write(breweryXml)
}

/*
	Convert all elements of the array into a string
	If array is empty, return a "No Result Found" message
*/
func structureBreweriesToString(breweries []Brewery) string {
	if len(breweries) == 0 {
		return "No Results Found"
	}

	str := ""

	for _, x := range breweries {
		str += fmt.Sprintf("%v\n%v\n%v\n%v, %v %v\n%v\n\n", x.Name, x.BreweryType, x.Street, x.City, x.State, x.PostalCode, x.Phone)
	}

	return str
}

/*
	Method to determine how to query the brewery list
*/
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

	http.HandleFunc("/", FindBrewery)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
