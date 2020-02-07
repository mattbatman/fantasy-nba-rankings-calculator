package sheets

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/mattbatman/fantasy-nba-rankings-calculator/config-parser"
	"github.com/mattbatman/fantasy-nba-rankings-calculator/utility"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

// From google sheets quickstart: https://developers.google.com/sheets/api/quickstart/go
// Retrieve a token, saves the token, then returns the generated client.
func getClient(c *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "sheets-auth/token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(c)
		saveToken(tokFile, tok)
	}
	return c.Client(context.Background(), tok)
}

// From google sheets quickstart: https://developers.google.com/sheets/api/quickstart/go
// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(c *oauth2.Config) *oauth2.Token {
	authURL := c.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := c.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// From google sheets quickstart: https://developers.google.com/sheets/api/quickstart/go
// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// From google sheets quickstart: https://developers.google.com/sheets/api/quickstart/go
// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// writeDataAsJSON takes JSON-esque data, marshals it, and writes it to file.
func writeDataAsJSON(data []map[string]interface{}) {
	rankingsJSON, _ := json.Marshal(data)

	if _, err := os.Stat("data"); os.IsNotExist(err) {
		os.MkdirAll("data", 0700)
	}

	err := ioutil.WriteFile("data/sheets-data.json", rankingsJSON, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// FetchSheetsData is the driver logic form include -mode="fetch".
// Similar to main() from google sheets quickstart: https://developers.google.com/sheets/api/quickstart/go
func FetchSheetsData() {
	appConfig, configErr := config.GetConfig()

	if configErr != nil {
		fmt.Printf("Config File Error: %s\n", configErr)
		os.Exit(1)
	}

	b, err := ioutil.ReadFile("sheets-auth/credentials.json")

	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	sheetsConfig, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")

	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client := getClient(sheetsConfig)

	srv, err := sheets.New(client)

	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	spreadsheetID := appConfig.SpreadsheetID

	readRange := appConfig.SpreadsheetRange

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()

	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		k, d := resp.Values[0], resp.Values[1:]
		zippedData := utility.ZipSheetsData(k, d)
		writeDataAsJSON(zippedData)
	}
}
