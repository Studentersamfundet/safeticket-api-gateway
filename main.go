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

const AccessKey string = "1337"
const STUser string = "Lasse"
const STSecret string = "b0acd427e2a99895c1dfe0c6d42b7ce071a88419cb76e271f8a18efe9e4bbf7b"

func main() {

	/*
	for _, order := range o.STorder {
		fmt.Println(order.STemail.Text)
		fmt.Println(order.STtickets.STticket[0].STticketnumber.Text)
	}
	*/

	http.HandleFunc("/aaulan", APIRequest)

	http.ListenAndServe(":8080", nil)
}

func GetAPIResponse(APIURL string, user string, events string, t1 time.Time, t2 time.Time, secret string) *http.Response {
	version := "1"
	sequenceNumber := ""
	timeFrom := t1.Format("2006-01-02")
	timeTo := t2.Format("2006-01-02")
	hashstring := []string{version, user, events, sequenceNumber, timeFrom, timeTo, secret}
	hasher := sha256.New()
	hasher.Write([]byte(strings.Join(hashstring, ":")))
	sha := hex.EncodeToString(hasher.Sum(nil))

	form := url.Values{}
	form.Add("version", version)
	form.Add("user", user)
	form.Add("events", events)
	form.Add("sequence_number", sequenceNumber)
	form.Add("t1", timeFrom)
	form.Add("t2", timeTo)
	form.Add("sha", sha)

	hc := http.Client{}
	req, err := http.NewRequest("POST", APIURL, strings.NewReader(form.Encode()))
	//req.Header.Add("Content-Type", "application/xml; charset=utf-8")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := hc.Do(req)
	if err != nil {
		fmt.Println("Error performing HTTP request", err)
	}

	return resp
}

func APIRequest(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("key") != AccessKey {
		fmt.Fprint(w, "Wrong Access Key!")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	xmlData := GetAPIResponse("https://studentersamfundet.safeticket.dk/api/orders", STUser, "23791", time.Date(2016, time.June, 01, 0, 0, 0, 0, time.UTC), time.Date(2016, time.September, 01, 0, 0, 0, 0, time.UTC), STSecret)
	
	b, _ := ioutil.ReadAll(xmlData.Body)
	var o STorders
	xml.Unmarshal(b, &o)
	json, err := json.Marshal(o)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(json)
}