package main

import (
	"flag"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"shorturl/conf"
	"shorturl/service"
	"sync"
	"time"
)

var (
	confFile string
	svr      *service.Service
)

var mutex = &sync.Mutex{}

func init() {
	flag.StringVar(&confFile, "c", "conf/shorturl.conf", "set config path")
}

func main() {
	flag.Parse()
	conf.InitConfig(confFile)
	run()
}

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleEcho).Methods("GET")
	muxRouter.HandleFunc("/gen_short_url", handleGenShortUrl).Methods("POST")
	muxRouter.HandleFunc("/get_long_url", handleGetLongUrl).Methods("GET")
	return muxRouter
}

func handleGenShortUrl(w http.ResponseWriter, r *http.Request) {
	longurl := r.PostFormValue("longurl")
	if longurl != "" {
		io.WriteString(w, svr.GenShortUrl(longurl)+"\n")
		return
	}
	io.WriteString(w, "noting post\n")
}

func handleGetLongUrl(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if len(query) != 0 {
		_, isExist := query["short"]
		if isExist {
			shorturl := query["short"][0]
			io.WriteString(w, svr.GetOriginUrl(shorturl)+"\n")
			return
		}
	}
	io.WriteString(w, "noting get\n")
}

func handleEcho(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Path)
}

func run() error {
	mux := makeMuxRouter()
	httpAddr := "8088"
	log.Println("Listening on ", httpAddr)
	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil

}
