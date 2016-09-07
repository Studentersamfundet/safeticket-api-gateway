package main

import (
	"fmt"
	"encoding/xml"
	//"encoding/json"
	//"os"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"crypto/sha256"
	"strings"
	//"encoding/base64"
	"encoding/hex"
	//"bytes"
)

type STaddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STaddress2 struct {
}

type STarrived struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STcardfee struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STcity struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STcode struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STcount struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STcountry struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STcustomfield struct {
	STname *STname `xml:" name,omitempty" json:"name,omitempty"`
	STvalue *STvalue `xml:" value,omitempty" json:"value,omitempty"`
}

type STcustomfields struct {
	STcustomfield *STcustomfield `xml:" customfield,omitempty" json:"customfield,omitempty"`
}

type STemail struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STevent struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STmailshipping struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STname struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STnewsletter struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STorder struct {
	STaddress *STaddress `xml:" address,omitempty" json:"address,omitempty"`
	STaddress2 *STaddress2 `xml:" address2,omitempty" json:"address2,omitempty"`
	STcardfee *STcardfee `xml:" cardfee,omitempty" json:"cardfee,omitempty"`
	STcity *STcity `xml:" city,omitempty" json:"city,omitempty"`
	STcountry *STcountry `xml:" country,omitempty" json:"country,omitempty"`
	STcustomfields *STcustomfields `xml:" customfields,omitempty" json:"customfields,omitempty"`
	STemail *STemail `xml:" email,omitempty" json:"email,omitempty"`
	STevent *STevent `xml:" event,omitempty" json:"event,omitempty"`
	STmailshipping *STmailshipping `xml:" mailshipping,omitempty" json:"mailshipping,omitempty"`
	STname *STname `xml:" name,omitempty" json:"name,omitempty"`
	STnewsletter *STnewsletter `xml:" newsletter,omitempty" json:"newsletter,omitempty"`
	STorderid *STorderid `xml:" orderid,omitempty" json:"orderid,omitempty"`
	STpaymentsequencenumber *STpaymentsequencenumber `xml:" paymentsequencenumber,omitempty" json:"paymentsequencenumber,omitempty"`
	STshippingfee *STshippingfee `xml:" shippingfee,omitempty" json:"shippingfee,omitempty"`
	STstatus *STstatus `xml:" status,omitempty" json:"status,omitempty"`
	STtelephone *STtelephone `xml:" telephone,omitempty" json:"telephone,omitempty"`
	STtickets *STtickets `xml:" tickets,omitempty" json:"tickets,omitempty"`
	STtime *STtime `xml:" time,omitempty" json:"time,omitempty"`
	STtotal *STtotal `xml:" total,omitempty" json:"total,omitempty"`
	STvatamount *STvatamount `xml:" vatamount,omitempty" json:"vatamount,omitempty"`
	STvouchercode *STvouchercode `xml:" vouchercode,omitempty" json:"vouchercode,omitempty"`
	STzip *STzip `xml:" zip,omitempty" json:"zip,omitempty"`
}

type STorderid struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STorders struct {
	STorder []STorder `xml:" order,omitempty" json:"order,omitempty"`
}

type STpaymentsequencenumber struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STprice struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STroot struct {
	STorders *STorders `xml:" orders,omitempty" json:"orders,omitempty"`
}

type STrow struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STseat struct {
	STrow *STrow `xml:" row,omitempty" json:"row,omitempty"`
	STseat *STseat `xml:" seat,omitempty" json:"seat,omitempty"`
	Text string `xml:",chardata" json:",omitempty"`
}

type STshippingfee struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STstatus struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STtelephone struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STticket struct {
	STarrived *STarrived `xml:" arrived,omitempty" json:"arrived,omitempty"`
	STcustomfields *STcustomfields `xml:" customfields,omitempty" json:"customfields,omitempty"`
	STprice *STprice `xml:" price,omitempty" json:"price,omitempty"`
	STseat *STseat `xml:" seat,omitempty" json:"seat,omitempty"`
	STticketfee *STticketfee `xml:" ticketfee,omitempty" json:"ticketfee,omitempty"`
	STticketid *STticketid `xml:" ticketid,omitempty" json:"ticketid,omitempty"`
	STticketnumber *STticketnumber `xml:" ticketnumber,omitempty" json:"ticketnumber,omitempty"`
}

type STticketfee struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STticketid struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STticketnumber struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STtickets struct {
	STticket []*STticket `xml:" ticket,omitempty" json:"ticket,omitempty"`
}

type STtime struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STtotal struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STvalue struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STvatamount struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type STvouchercode struct {
	STcode *STcode `xml:" code,omitempty" json:"code,omitempty"`
	STcount *STcount `xml:" count,omitempty" json:"count,omitempty"`
}

type STzip struct {
	Text string `xml:",chardata" json:",omitempty"`
}

func main() {
	xmlData := GetAPIResponse("https://studentersamfundet.safeticket.dk/api/orders", "Lasse", "23791", time.Date(2016, time.July, 01, 0, 0, 0, 0, time.UTC), time.Date(2016, time.September, 01, 0, 0, 0, 0, time.UTC), "b0acd427e2a99895c1dfe0c6d42b7ce071a88419cb76e271f8a18efe9e4bbf7b")

	b, _ := ioutil.ReadAll(xmlData.Body)

	var o STorders
	xml.Unmarshal(b, &o)

	for _, order := range o.STorder {
		fmt.Println(order.STemail.Text)
		fmt.Println(order.STtickets.STticket[0].STticketnumber.Text)
	}
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