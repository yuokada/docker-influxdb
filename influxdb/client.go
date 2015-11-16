package main

import (
	// "io"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const RANGE = 10000

// curl -i -XPOST 'http://localhost:8086/write?db=ysar'
// --data-binary 'cpu,host=server01,region=us-west value=39.64 1440481430000000000'

func makeDummyData() bytes.Buffer {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	value := r.Float64() * 100.0
	ts := time.Now().Add(time.Duration(r.Intn(RANGE*0.1)) * time.Second)
	var buffer bytes.Buffer
	record := fmt.Sprintf("cpu,host=server%02d value=%2f %d", 100+r.Intn(10), value, ts.UnixNano())
	buffer.WriteString(record)
	return buffer
}

func main() {
	for i := 0; i <= 100; i++ {
		d := makeDummyData()
		log.Println("DEBUG: ", &d)
		// resp, err := http.Post("http://graph01.data.cc1.ynwm.yahoo.co.jp:8086/write?db=ysar", "application/x-www-form-urlencoded", &buffer)
		// resp, err := http.Post("http://graph01.data.cc1.ynwm.yahoo.co.jp:8086/write?db=ysar", "application/x-www-form-urlencoded", &d)
		resp, err := http.Post("http://192.168.59.103:8086/write?db=ysar", "application/x-www-form-urlencoded", &d)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		if string(body) != "" {
			log.Print(string(body))
		}
	}
}
