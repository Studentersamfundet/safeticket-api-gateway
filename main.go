package main

import (
	"fmt"
	"encoding/xml"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"crypto/sha256"
	"strings"
	"encoding/hex"
)

type STcustomfield struct {
	STname string `xml:" name,omitempty" json:"name,omitempty"`
	STvalue string `xml:" value,omitempty" json:"value,omitempty"`
}

type STcustomfields struct {
	STcustomfield []STcustomfield `xml:" customfield,omitempty" json:"customfield,omitempty"`
}

type STorder struct {
	STaddress string `xml:" address,omitempty" json:"address,omitempty"`
	STaddress2 string `xml:" address2,omitempty" json:"address2,omitempty"`
	STcardfee string `xml:" cardfee,omitempty" json:"cardfee,omitempty"`
	STcity string `xml:" city,omitempty" json:"city,omitempty"`
	STcountry string `xml:" country,omitempty" json:"country,omitempty"`
	STcustomfields *STcustomfields `xml:" customfields,omitempty" json:"customfields,omitempty"`
	STemail string `xml:" email,omitempty" json:"email,omitempty"`
	STevent string `xml:" event,omitempty" json:"event,omitempty"`
	STmailshipping string `xml:" mailshipping,omitempty" json:"mailshipping,omitempty"`
	STname string `xml:" name,omitempty" json:"name,omitempty"`
	STnewsletter string `xml:" newsletter,omitempty" json:"newsletter,omitempty"`
	STorderid string `xml:" orderid,omitempty" json:"orderid,omitempty"`
	STpaymentsequencenumber string `xml:" paymentsequencenumber,omitempty" json:"paymentsequencenumber,omitempty"`
	STshippingfee string `xml:" shippingfee,omitempty" json:"shippingfee,omitempty"`
	STstatus string `xml:" status,omitempty" json:"status,omitempty"`
	STtelephone string `xml:" telephone,omitempty" json:"telephone,omitempty"`
	STtickets *STtickets `xml:" tickets,omitempty" json:"tickets,omitempty"`
	STtime string `xml:" time,omitempty" json:"time,omitempty"`
	STtotal string `xml:" total,omitempty" json:"total,omitempty"`
	STvatamount string `xml:" vatamount,omitempty" json:"vatamount,omitempty"`
	STvouchercode *STvouchercode `xml:" vouchercode,omitempty" json:"vouchercode,omitempty"`
	STzip string `xml:" zip,omitempty" json:"zip,omitempty"`
}

type STorders struct {
	STorder []STorder `xml:" order,omitempty" json:"order,omitempty"`
}

type STroot struct {
	STorders *STorders `xml:" orders,omitempty" json:"orders,omitempty"`
}

type STseat struct {
	STrow string `xml:" row,omitempty" json:"row,omitempty"`
	STseat string `xml:" seat,omitempty" json:"seat,omitempty"`
}

type STticket struct {
	STarrived string `xml:" arrived,omitempty" json:"arrived,omitempty"`
	STcustomfields *STcustomfields `xml:" customfields,omitempty" json:"customfields,omitempty"`
	STprice string `xml:" price,omitempty" json:"price,omitempty"`
	STseat *STseat `xml:" seat,omitempty" json:"seat,omitempty"`
	STticketfee string `xml:" ticketfee,omitempty" json:"ticketfee,omitempty"`
	STticketid string `xml:" ticketid,omitempty" json:"ticketid,omitempty"`
	STticketnumber string `xml:" ticketnumber,omitempty" json:"ticketnumber,omitempty"`
}

type STtickets struct {
	STticket []*STticket `xml:" ticket,omitempty" json:"ticket,omitempty"`
}

type STvouchercode struct {
	STcode string `xml:" code,omitempty" json:"code,omitempty"`
	STcount string `xml:" count,omitempty" json:"count,omitempty"`
}

type APIResponse struct {
	TicketExists bool
	OrderDetails STorder
}

const AccessToken string = "YOUR_ACCESS_TOKEN" // Access token for pulling data from the API Gateway, generate something secure for this
const STUser string = "SAFETICKET_API_USER" // Safeticket API user
const STSecret string = "SAFETICKET_API_SECRET" // Secret for the API user

func main() {
	// Handle HTTP requests to the defined url.
	http.HandleFunc("/aaulan", APIRequest)

	// Start the HTTP server and listen to the defined port.
	http.ListenAndServe(":8080", nil)
}

// GetAPIResponse returns resulting XML data from the Safeticket API.
func GetAPIResponse(APIURL string, user string, secret string, events string, t1 time.Time, t2 time.Time) *http.Response {
	version := "1"
	sequenceNumber := ""
	timeFrom := t1.Format("2006-01-02")
	timeTo := t2.Format("2006-01-02")

	// Hash input data as required by Safeticket.
	hashstring := []string{version, user, events, sequenceNumber, timeFrom, timeTo, secret}
	hasher := sha256.New()
	hasher.Write([]byte(strings.Join(hashstring, ":")))
	sha := hex.EncodeToString(hasher.Sum(nil))

	// Define the parameters that should be sent to the Safeticket API.
	form := url.Values{}
	form.Add("version", version)
	form.Add("user", user)
	form.Add("events", events)
	form.Add("sequence_number", sequenceNumber)
	form.Add("t1", timeFrom)
	form.Add("t2", timeTo)
	form.Add("sha", sha)

	// Create a new HTTP request and get the response from the Safeticket API.
	hc := http.Client{}
	req, err := http.NewRequest("POST", APIURL, strings.NewReader(form.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := hc.Do(req)
	if err != nil {
		fmt.Println("Error performing HTTP request", err)
	}

	return resp
}

// APIRequest handles incoming requests and returns appropriate information as JSON.
func APIRequest(w http.ResponseWriter, r *http.Request) {

	// Check if the entered key for authentication matches the one we've set.
	// If it does not match, return and throw an error.
	if r.FormValue("access_token") != AccessToken {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Wrong Access Token!")
		return
	}

	// Set correct content type and define the name of the ticketNumber url value.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ticketNumber := r.FormValue("ticketnumber")

	// Get the raw XML data from the Safeticket API
	xmlData := GetAPIResponse("https://studentersamfundet.safeticket.dk/api/orders", STUser, STSecret, "25975,26422", time.Date(2016, time.September, 06, 0, 0, 0, 0, time.UTC), time.Date(2016, time.October, 15, 0, 0, 0, 0, time.UTC))
	
	// Read XML data from the response into a byte array and unmarshal it to the structs defined above.
	var o STorders
	b, _ := ioutil.ReadAll(xmlData.Body)
	xml.Unmarshal(b, &o)

	// Check if an order with the entered ticketNumber exists and set orderExists and orderIndex accordingly.
	// This only works for orders with one ticket for now, no reason to waste compute time when you can only buy one per order anyways ;)
	var orderExists bool
	var orderIndex int
	for index,element := range o.STorder {
		if element.STtickets.STticket[0].STticketnumber == ticketNumber {
			orderIndex = index
			orderExists = true
		}
	}

	// Set the contents of the APIResponse struct based on whether the order exists or not.
	// Also set an appropriate HTTP status for each case.
	var response APIResponse
	if orderExists == true {
		w.WriteHeader(http.StatusOK)
		response.TicketExists = true
		response.OrderDetails = o.STorder[orderIndex]
	} else {
		w.WriteHeader(http.StatusNotFound)
		response.TicketExists = false
		response.OrderDetails = STorder{}
	}

	// Marshal the APIResponse defined above and write the resulting JSON to the ResponseWriter.
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(json)
}
