package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/devopsfaith/bloomfilter/rpc/client"
)

func main() {
	krakendRevokerPort := os.Getenv("KRAKEND_CE_PORT_1")
	server := flag.String("server", "krakend_ce:"+krakendRevokerPort, "ip:port of the remote bloomfilter to connect to")
	key := flag.String("key", "jti", "the name of the claim to inspect for revocations")
	port := flag.Int("port", 8080, "port to expose the service")
	flag.Parse()

	c, err := client.New(*server)
	if err != nil {
		log.Println("unable to create the rpc client:" +krakendRevokerPort, err.Error())
		return
	}
	defer c.Close()

	http.HandleFunc("/check/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		subject := *key + "-" + r.FormValue(*key)
		res := c.Check([]byte(subject))
		log.Printf("checking [%s] %s => %v", *key, subject, res)
		fmt.Fprintf(w, "%v", res)
	})

	http.HandleFunc("/", func(rw http.ResponseWriter, _ *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		jsonData := []byte(`{"status":true, "message":"Berhasil melakukan logout", "data":[]}`)
		rw.Write(jsonData)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}