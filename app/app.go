package app

import (
	"log"
	"net/http"
	"os"

	"github.com/freischarler/hexpattern/domain"
	"github.com/freischarler/hexpattern/service"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

//USER- HandlerAdapter -> IPortService -> (Domain LOGIC) -> IPortRepository -> DBAdapter -> DB

const defaultAddr = "localhost"
const defaultPort = "9000"

func Start() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		println("Set default port")
		serverPort = defaultPort
	}

	addr := os.Getenv("IP")
	if addr == "" {
		addr = defaultAddr
	}
	router := mux.NewRouter()

	//wiring
	ch := BeerHandlers{service.NewBeerService(domain.NewBeerRepositoryDb())}

	//define routes
	router.HandleFunc("/beers", ch.getAllBeers).Methods(http.MethodGet)
	router.HandleFunc("/beers", ch.postBeer).Methods(http.MethodPost)
	router.HandleFunc("/beers/{beer_id:[0-9]+}", ch.getOneBeer).Methods(http.MethodGet)
	router.HandleFunc("/beers/{beer_id:[0-9]+}/boxprice", ch.getBoxBeer).Methods(http.MethodGet)

	//starting server
	log.Fatal(http.ListenAndServe(addr+":"+serverPort, router))
}
