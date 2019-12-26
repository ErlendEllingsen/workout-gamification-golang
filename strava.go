package main

import "os"
import "strings"
import "net/http"
import "io/ioutil"
import "encoding/json"
import "log"

type StravaActivity struct {
	ID          int64   `json:"id"`
	Type        string  `json:"type"`
	Distance    float32 `json:"distance"`
	AverageTemp int32   `json:"average_temp"`
}

func (sa StravaActivity) ToActivityEntry() ActivityEntry {
	if strings.ToLower(sa.Type) == "run" {
		return ActivityEntry{
			activity:    Activities.running,
			distance:    sa.Distance / 1000, // in metres
			outsiteTemp: float32(sa.AverageTemp),
		}
	}
	return ActivityEntry{}
}

type StravaOauthResponse struct {
	AccessToken string `json:"access_token"`
}

func getNewToken() string {
	// STRAVA DETS
	clientID := "41925"

	// Read client secret
	clientSecretB, _ := ioutil.ReadFile("strava-app-secret.txt")
	clientSecret := string(clientSecretB)

	// Read user token
	userTokenB, _ := ioutil.ReadFile("strava-user-secret.txt")
	userToken := string(userTokenB)

	reqStr := "https://www.strava.com/oauth/token?client_id=" + clientID + "&client_secret=" + clientSecret + "&code=" + userToken + "&grant_type=authorization_code"
	client := &http.Client{}
	req, err := http.NewRequest("POST", reqStr, nil)
	resp, err := client.Do(req)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var oauthResp StravaOauthResponse
	if err := json.Unmarshal(b, &oauthResp); err != nil {
		panic(err)
	}

	token := oauthResp.AccessToken

	// Store oauth2-token for further use
	ioutil.WriteFile("strava-oauth2-token.txt", []byte(token), 0644)
	return token
}

func getActivitiesFromStrava(retry bool) []StravaActivity {
	token := ""
	client := &http.Client{}

	// Check if oauth2 temp token exists
	_, err := os.Stat("strava-oauth2-token.txt")
	if os.IsNotExist(err) {
		log.Print("Key did not exist, fetching new oaut2-token")
		token = getNewToken()
	} else {
		tokenBytes, _ := ioutil.ReadFile("strava-oauth2-token.txt")
		token = string(tokenBytes)
	}

	// Test oauth2 temp token, check for error
	req, err := http.NewRequest("GET", "https://www.strava.com/api/v3/athlete", nil)
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)

	if resp.StatusCode != 200 {

		// Prevent recursive iteration deeper than one level
		if retry {
			log.Fatal("Unable to fetch valid oauth2-token for user. Check user secret or re-authenticate.")
			os.Exit(1)
		}

		log.Print("Statuscode was NOT 200, fetching new oauth2-token")

		// Token might have expired
		token = getNewToken()
		return getActivitiesFromStrava(true)
	}

	log.Print("Statuscode was 200 OK")

	// Fetch activities

	reqStr := "https://www.strava.com/api/v3/athlete/activities?after=1561939200"
	reqActs, _ := http.NewRequest("GET", reqStr, nil)
	reqActs.Header.Add("Authorization", "Bearer "+token)
	respActs, _ := client.Do(reqActs)

	b, err := ioutil.ReadAll(respActs.Body)
	if err != nil {
		panic(err)
	}

	return parseActivities(b)
}

func parseActivities(input []byte) []StravaActivity {
	var acts []StravaActivity
	if err := json.Unmarshal(input, &acts); err != nil {
		panic(err)
	}
	return acts
}
