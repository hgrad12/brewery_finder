package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestFindBrewery(t *testing.T) {
	t.Run("Execute FindBrewery, with no params", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/upper", nil)
		w := httptest.NewRecorder()
		FindBrewery(w, req)
		res := w.Result()
		defer res.Body.Close()
		_, err := ioutil.ReadAll(res.Body)

		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}
	})

	t.Run("Execute FindBrewery, so that it returns no result", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/upper?body=Hello&phone_number=4438916412", nil)
		w := httptest.NewRecorder()
		FindBrewery(w, req)
		res := w.Result()
		defer res.Body.Close()
		data, err := ioutil.ReadAll(res.Body)

		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}

		intValue, _ := strconv.ParseInt(string(data), 0, 3)

		if intValue != 0 {
			t.Errorf("expected ABC got %v", intValue)
		}
	})

	t.Run("Execute FindBrewery, so that it return a result", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/upper?body=our_mutual_friend&phone_number=4438916412", nil)
		w := httptest.NewRecorder()
		FindBrewery(w, req)
		res := w.Result()
		defer res.Body.Close()
		data, err := ioutil.ReadAll(res.Body)

		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}

		intValue, _ := strconv.ParseInt(string(data), 0, 3)

		if intValue != 1 {
			t.Errorf("expected ABC got %v", intValue)
		}
	})

	t.Run("Execute FindBrewery, so that it returns multiple results", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/upper?body=80216&phone_number=4438916412", nil)
		w := httptest.NewRecorder()
		FindBrewery(w, req)
		res := w.Result()
		defer res.Body.Close()
		data, err := ioutil.ReadAll(res.Body)

		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}

		intValue, _ := strconv.ParseInt(string(data), 0, 3)

		if intValue <= 2 {
			t.Errorf("expected ABC got %v", intValue)
		}
	})

	t.Run("Execute FindBrewery, send results back to an unavailable number", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/upper?body=80216&phone_number=9999999999", nil)
		w := httptest.NewRecorder()
		FindBrewery(w, req)
		res := w.Result()
		defer res.Body.Close()
		data, err := ioutil.ReadAll(res.Body)

		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}

		intValue, _ := strconv.ParseInt(string(data), 0, 3)

		if intValue <= 2 {
			t.Errorf("expected ABC got %v", intValue)
		}
	})
}

func TestStructureBreweriesToString(t *testing.T) {
	var Breweries []Brewery
	got := StructureBreweriesToString(Breweries)
	want := "No Results Found"

	if want != got {
		t.Errorf("got %q want %q", got, want)
	}

	name := "Brewery Name"
	typeOfBrewery := "micro"
	street := "1234 Rainbow RD"
	city := "South Royalton"
	state := "VT"
	zip := "99999-1234"
	phone := "1111111111"
	var b Brewery
	b.Name = name
	b.BreweryType = typeOfBrewery
	b.Street = street
	b.City = city
	b.State = state
	b.PostalCode = zip
	b.Phone = phone
	Breweries = append(Breweries, b)

	got = StructureBreweriesToString(Breweries)

	typeOfBrewery = "Micro"
	phone = "(111) 111-1111"

	want = fmt.Sprintf("%v\n%v\n%v\n%v, %v %v\n%v\n\n", name, typeOfBrewery, street, city, state, zip, phone)

	if want != got {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestConstructQuery(t *testing.T) {
	t.Run("Constuct a zip code query", func(t *testing.T) {
		want := "https://api.openbrewerydb.org/breweries?by_postal=99999&per_page=3&sort=name:asc"
		got := ConstructQuery("by_postal", "99999")

		if want != got {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Construct a name query", func(t *testing.T) {
		want := "https://api.openbrewerydb.org/breweries?by_name=our_mutual_friend&per_page=3&sort=name:asc"
		got := ConstructQuery("by_name", "Our Mutual Friend")

		if want != got {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestValidate(t *testing.T) {
	t.Run("Validate if string is a zip code", func(t *testing.T) {
		want := "by_postal"
		got := Validate("11111")

		if want != got {
			t.Errorf("got %q want %q", got, want)
		}

		got = Validate("123")

		if want == got {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Validate if string is a name", func(t *testing.T) {
		want := "by_name"
		got := Validate("Something Brewery")

		if want != got {
			t.Errorf("got %q want %q", got, want)
		}

		got = Validate("12345-1234")

		if want == got {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Validate if string is a set of coordinates (lat/long)", func(t *testing.T) {
		want := "by_dist"
		got := Validate("47.1231231, 179.99999999")

		if want != got {
			t.Errorf("got %q want %q", got, want)
		}

		got = Validate("Latitude, Longitude")

		if want == got {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
