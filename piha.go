package main

import (
	"bytes"
	"log"
	"math/rand"
	"os"
	"strconv"
	"text/template"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

// Credentials stores the secret keys needed
// for authentication against the Twitter API
type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

// NewCredentials creates a new Credentials
// struct reading the secrets from .env file
func NewCredentials() Credentials {
	return Credentials{
		ConsumerKey:       os.Getenv("TWITTER_API_KEY"),
		ConsumerSecret:    os.Getenv("TWITTER_API_KEY_SECRET"),
		AccessToken:       os.Getenv("TWITTER_ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"),
	}
}

// NewTwitterClient takes a Credentials struct and
// returns a twitter client
func NewTwitterClient(creds *Credentials) *twitter.Client {
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	return twitter.NewClient(httpClient)
}

// RantConfig stores the rant configuration
type RantConfig struct {
	User     string
	Days     string
	Suffix   string
	Template string
}

// NewRantConfig creates a new rant config struct
// reading the values from .env file
func NewRantConfig() RantConfig {
	return RantConfig{
		User:     os.Getenv("RANT_USER"),
		Days:     getRantDays(os.Getenv("RANT_DATE")),
		Suffix:   getRandomRantSuffix(),
		Template: os.Getenv("RANT_TEMPLATE"),
	}
}

// getRantDays computes the number of days
// since a given date
func getRantDays(date string) string {
	layout := "2006-01-02"
	since, err := time.Parse(layout, date)
	if err != nil {
		// fallback
		return "SEVERAL"
	}
	// count the number of days elapsed
	days := int(time.Now().Sub(since).Hours() / 24)
	return strconv.Itoa(days)
}

// getRandomRantSuffix returns a random rant
// suffix from a predfined list
func getRandomRantSuffix() string {
	// list of rant suffixes
	rants := []string{
		"Where is my money? 😡",
		"Are you proud of this? 😒",
		"Do you do this to all your customers? 🙄",
	}
	rand.Seed(time.Now().UnixNano())
	// pick a random rant from the list
	return rants[rand.Intn(len(rants))]
}

// NewRant creates a new rant from the template
// with a random rant suffix
func NewRant() (string, error) {
	config := NewRantConfig()
	tmpl, err := template.New("rant").Parse(config.Template)
	if err != nil {
		return "", err
	}
	// capture the template output to a var
	var rant bytes.Buffer
	err = tmpl.Execute(&rant, config)
	if err != nil {
		return "", err
	}
	return rant.String(), nil
}

func init() {
	// check if running locally
	localEnvFile := "local.env"
	if _, err := os.Stat(localEnvFile); err == nil {
		godotenv.Load(localEnvFile)
	}
	// load config
	godotenv.Load("config.env")
}

func main() {
	rant, err := NewRant()
	if err != nil {
		// no rant, no fun. exit.
		log.Fatal("Couldn't make a new rant")
	}

	// setup credentials and twitter client
	creds := NewCredentials()
	client := NewTwitterClient(&creds)

	// post the rant
	_, _, err = client.Statuses.Update(rant, nil)
	if err != nil {
		log.Fatal("Rant didn't go through!")
	}
	log.Print("Tweeted rant: ", rant)
}
