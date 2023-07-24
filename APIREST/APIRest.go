package main

import(
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux" //go get -u github.com/gorilla/mux
	)
	
type WebConfig struct {
  Address IPAddr `json: "Address"`
  Port uint16    `json: "Port"`
  Slot uint8     `json: "Slot"`
}

type IP [4]byte

type IPAddr struct {
  IP IP
}

func (pIP IP) IPAddrToCharArr(dwArrSize uint8) [dwArrSize]C.Char {
  var ArrChar [dwArrSize]C.Char
  for i:=range ArrChar {
    ArrChar[i]=pIP[i]
  }
}

var webConf []WebConfig
	
func GetNewServer(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(webConf)
}

func CreateServer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var server WebConfig
	_ = json.NewDecoder(r.Body).Decode(&server)
	server.Address = params["Address"]
	server.Port = params["Port"]
	server.Slot = params["Slot"]
	webConf = append(webConf, server)
	json.NewEncoder(w).Encode(webConf)
}

func DeleteServer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index,item := range webConf {
		if item.Address == params["Address"] {
			webConf = append(webConf[:index], webConf[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(webConf)
}

func main(){
	router := mux.NewRouter()
	webConf = webConf.append(webConf, WebConfig{Address: "moovsec.intelbras.com.br", Port: 5020, Slot: 2})
	router.HandleFunc("/server",GetNewServer).Methods("GET")
	router.HandleFunc("/server/{id}",CreateServer).Methods("POST")
	router.HandleFunc("/server/{id}",DeleteServer).Methods("DELETE")	
	log.Fatal(http.ListenAndServe(":8000", router))
}
