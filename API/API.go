package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
Interface address：
http://test.growatt.com/v1/plant/energy
The interface requires parameters：
plant_id:2
start_date:”2015-05-03”
end_date:”2015-05-06”
time_unit:year
page:1
perpage:10
*/

type API struct {
}

func Start() {
	router := http.NewServeMux()
	api := &API{}

	router.HandleFunc("/", api.getJson)
	//Server
	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	//Run Server
	server.ListenAndServe()
}

func (a *API) getJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	switch r.Method {
	case "GET": //READ
		a.showWebConfig(w, r)
	default:
		log.Println("Request welcome route")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"Welcome to root"}`))
	}
}

func (a *API) showWebConfig(w http.ResponseWriter, r *http.Request) {
	device := r.URL.Query().Get("device")
	var err error
	switch device {
	case "Streamax":
		webconfig, err = a.SDK[0].GetConfig()
		if err != nil {
			fmt.Fprintf(w, "%+v", err)
			return
		}
	case "Dahua":
		webconfig, err = a.SDK[1].GetConfig()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%+v", err)
			return
		}
	}
	fmt.Fprintf(w, "%+v", webconfig)
}
