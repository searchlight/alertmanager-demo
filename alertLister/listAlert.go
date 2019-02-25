package main

import (
	"encoding/json"
	"github.com/tamalsaha/go-oneliners"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//Alert Structure
type Alert struct {
	Annotations  map[string]string `json:"annotations"`
	Labels       map[string]string `json:"labels"`
	StartsAt     time.Time         `json:"startsAt"`
	EndsAt       time.Time         `json:"endsAt"`
	Fingerprint  string            `json:"fingerprint"`
	GeneratorURL string            `json:"generatorURL"`
	Receivers    []struct {
		Name string `json:"name"`
	} `json:"receivers"`
	Status struct {
		InhibitedBy []string `json:"inhibitedBy"`
		SilencedBy  []string `json:"silencedBy"`
		State       string   `json:"state"`
	} `json:"status"`
}

type Matcher struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	IsRegex bool   `json:"isRegex"`
}
// silenced alert structure
type SilencedAlert struct {
	ID        string     `json:"id"`
	Matchers  []*Matcher `json:"matchers"`
	StartsAt  time.Time  `json:"startsAt"`
	EndsAt    time.Time  `json:"endsAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	CreatedBy string     `json:"createdBy"`
	Comment   string     `json:"comment,omitempty"`
	Status    struct{
		State string `json:"state"`
	}     `json:"status"`
}

// Demo database
var AlertStorage map[string]Alert
var SilencedAlertStorage map[string]SilencedAlert

// Initialize demo database
func init() {
	AlertStorage = make(map[string]Alert)
	SilencedAlertStorage = make(map[string]SilencedAlert)
}

// Updates demo database for storing alerts
func UpdateAlertStorage(body io.ReadCloser)  {

	data, err := ioutil.ReadAll(body)
	if err != nil {
		log.Printf("ioutil.ReadAll() failed with '%s'\n", err)
	}

	var alerts []Alert
	err = json.Unmarshal(data, &alerts)
	if err != nil {
		log.Printf("json.Unmarshal() failed with '%s'\n", err)
	}

	for _, value := range alerts {
		AlertStorage[value.Labels["alertname"]] = value

	}
}
// Updates demo database for storing silenced alerts
func UpdateSilencedAlertStorage(body io.ReadCloser)  {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		log.Printf("ioutil.ReadAll() failed with '%s'\n",err)
	}
	var alerts []SilencedAlert
	err = json.Unmarshal(data, &alerts)
	if err != nil {
		log.Printf("json.Unmarshal() failed with '%s'\n", err)
	}
	for _,value := range alerts {
		SilencedAlertStorage[value.ID]=value
	}
}

func main() {
	client := &http.Client{}
	client.Timeout = time.Second * 10

	//Get alert list from api/v2/alerts end point
	res, err := client.Get("http://localhost:9093/api/v2/alerts")
	if err != nil {
		log.Printf("client.Get() failed with '%s'\n", err)
	}
	UpdateAlertStorage(res.Body)
	oneliners.PrettyJson(AlertStorage)

	// Get silenced alert list from api/v2/silences end point
	res, err = client.Get("http://localhost:9093/api/v2/silences")
	if err != nil {
		log.Printf("client.Get() failed with '%s'\n",err)
	}
	UpdateSilencedAlertStorage(res.Body)
	oneliners.PrettyJson(SilencedAlertStorage)

}
