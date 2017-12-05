package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

type slackMessage struct {
	Text string `json:"text"`
}

var receivedMessage slackMessage
var url string
var log = logrus.New()

func main() {

	url = os.Getenv("WEBHOOK_URL")
	fmt.Println("URL:>", url)

	http.HandleFunc("/message", messageHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("Getting message")
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var t slackMessage
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Info("The messages is ", t)
	receivedMessage = t
	status, err := sendToSlack(t, url)
	w.WriteHeader(status)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	} else {
		fmt.Fprintf(w, "OK")
	}

}

func sendToSlack(message slackMessage, url string) (int, error) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(message)
	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	req.Header.Set("Content-Type", "application/json")

	log.Info("Posting ", req.URL, req.Method, req.Body)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	r, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, fmt.Errorf("Something went wrong posting to slack: %s", r)
	}
	log.Info("The response status is ", resp.Status)
	log.Info("The response is ", r)
	defer resp.Body.Close()
	return http.StatusOK, nil
}
