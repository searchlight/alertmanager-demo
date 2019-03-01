package main


import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)
// Alert structure holds the alerting information
type Alert struct {
	Labels map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	StartsAt string `json:"startsAt,omitempty"`
	EndsAt string `json:"endsAt,omitempty"`
	GeneratorURL string `json:"generatorURL"`

}

var Alerts []Alert

//Demo alerts created
func CreateAlerts()  {
	Alerts = append(Alerts,Alert{
		Labels: map[string]string{
			"alertname":"First Alert",
			"job":"Do nothing",
		},
		Annotations: map[string]string{
			"annotation1":"annotationValue1",
		},
		GeneratorURL: "http://localhost:9093",
	})

	Alerts = append(Alerts,Alert{
		Labels: map[string]string{
			"alertname":"Second Alert",
			"job":"Do something",
		},
		Annotations: map[string]string{
			"annotation2":"annotationValue2",
		},
		GeneratorURL: "http://localhost:9094",
	})

	Alerts = append(Alerts,Alert{
		Labels: map[string]string{
			"alertname":"Final Alert",
			"job":"take action",
		},
		Annotations: map[string]string{
			"annotation3":"annotationValue3",
		},
		GeneratorURL: "http://localhost:9095",
	})
}

func main()  {
	//Creating alerts
	CreateAlerts()
	//creating http client
	client := &http.Client{}
	client.Timeout = time.Second * 10

	//Converting structure to json
	value, err := json.Marshal(Alerts)
	if err != nil {
		log.Fatalf("json.Marshal() failed with %s\n",err)
	}
	body := bytes.NewReader(value)

	//Making POST request to alertManager
	resp, err := client.Post("http://localhost:9093/api/v1/alerts","application/json",body)
	if err != nil {
		log.Fatalf("client.Post() failed with %s\n",err)
	}
	defer resp.Body.Close()
	fmt.Printf("http.Post() returned status code %d\n",resp.StatusCode)


}
