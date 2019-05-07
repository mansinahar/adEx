package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// Config represents the app config
type Config struct {
	port int
}

func main() {
	config := new(Config)

	// Parse command line flags
	parseFlags(config)

	// Setup logger
	logger := log.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(log.InfoLevel)

	// Setup http router
	router := setupRouter()

	// Setup http server
	s := &http.Server{
		Addr:         ":" + strconv.Itoa(config.port),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Infof("Starting adEx server at port %d...", config.port)

	// Start listening and serving http requests
	log.Fatal(s.ListenAndServe())
}

// parseFlags parses command line flags
func parseFlags(config *Config) {
	flag.IntVar(&config.port, "port", 5000, "server listen address")
	flag.Parse()
}

// setupRouter sets up the http router
func setupRouter() *httprouter.Router {
	router := httprouter.New()
	// Setup health check endpoint
	router.GET("/health", health)
	// Setup bidders endpoint
	router.GET("/bidders", getBidders)

	return router
}

// health is an http handler for the /health - health check endpoint
func health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Hello World!"))
}

// getBidders is an http handler for the /bidders endpoint
func getBidders(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	biddersJSON, err := json.Marshal(GetAllBidders())
	if err != nil {
		err = fmt.Errorf("Error marshaling bidder response: %s", err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		err = fmt.Errorf("Error marshaling bidder response: %s", err.Error())
                log.Error(err)
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
	}
	w.Write(biddersJSON)
}
