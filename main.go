package main

import (
	"github.com/globalsign/mgo"
	"github.com/manveru/faker"
)

type data struct {
	PostCode           string
	City               string
	Country            string
	CityPrefix         string
	CitySuffix         string
	CompanyBs          string
	CompanyCatchPhrase string
	CompanyName        string
	CompanySuffix      string
	DomainName         string
	DomainSuffix       string
	Email              string
	FirstName          string
	FreeEmail          string
	JobTitle           string
	LastName           string
	Name               string
	NamePrefix         string
	NameSuffix         string
	PhoneNumber        string
	SecondaryAddress   string
	State              string
	StateAbbr          string
	StreetAddress      string
	StreetName         string
	StreetSuffix       string
	URL                string
	UserName           string
	Latitude           float64
	Longitude          float64
}

func main() {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	collection := session.DB("fakeData").C("test")
	for i := 0; i < 2000000; i++ {
		fake, _ := faker.New("en")
		collection.Insert(data{
			PostCode:           fake.PostCode(),
			City:               fake.City(),
			Country:            fake.Country(),
			CityPrefix:         fake.CityPrefix(),
			CitySuffix:         fake.CitySuffix(),
			CompanyBs:          fake.CompanyBs(),
			CompanyCatchPhrase: fake.CompanyCatchPhrase(),
			CompanyName:        fake.CompanyName(),
			CompanySuffix:      fake.CompanySuffix(),
			DomainName:         fake.DomainName(),
			DomainSuffix:       fake.DomainSuffix(),
			Email:              fake.Email(),
			FirstName:          fake.FirstName(),
			FreeEmail:          fake.FreeEmail(),
			JobTitle:           fake.JobTitle(),
			LastName:           fake.LastName(),
			Name:               fake.Name(),
			NamePrefix:         fake.NamePrefix(),
			NameSuffix:         fake.NameSuffix(),
			PhoneNumber:        fake.PhoneNumber(),
			SecondaryAddress:   fake.SecondaryAddress(),
			State:              fake.State(),
			StateAbbr:          fake.StateAbbr(),
			StreetAddress:      fake.StreetAddress(),
			StreetName:         fake.StreetName(),
			StreetSuffix:       fake.StreetSuffix(),
			URL:                fake.URL(),
			UserName:           fake.UserName(),
			Latitude:           fake.Latitude(),
			Longitude:          fake.Longitude(),
		})
	}
}
