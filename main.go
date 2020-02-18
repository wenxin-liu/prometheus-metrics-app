package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

var (
	BoringCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "boring_count",
	})
	SpecialCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "special_count",
	})
	RequestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "request_count",
	})
	IncrementalCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "incremental_count",
	})
)

type MyDefaultHandler struct {

}

type SpecialHandler struct {

}


func (m MyDefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	log.Printf("Serving request for path %s", r.URL.Path)
	w.Write([]byte(r.URL.Path))
	BoringCounter.Inc()
	RequestCounter.Inc()

}

func (s SpecialHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	log.Println("Special handler is called")
	w.Write([]byte("Special"))
	SpecialCounter.Inc()
	RequestCounter.Inc()
}

func Increment(){
	for {
		IncrementalCounter.Inc()
		time.Sleep(time.Second * 2)
	}
}


func main() {
	go Increment()
	mux := http.NewServeMux()
	mux.Handle("/special", SpecialHandler{})
	mux.Handle("/", MyDefaultHandler{})
	mux.Handle("/metrics", promhttp.Handler())
	server := http.Server{Addr: "127.0.0.1:8080", Handler: mux}
	server.ListenAndServe()
}


