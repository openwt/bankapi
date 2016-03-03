package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yageek/euroconv/Godeps/_workspace/src/github.com/codegangsta/negroni"
	"github.com/yageek/euroconv/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/yageek/euroconv/Godeps/_workspace/src/github.com/nabeken/negroni-auth"
	"github.com/yageek/euroconv/Godeps/_workspace/src/github.com/unrolled/render"
	"github.com/yageek/euroconv/eurobank"
)

var router *mux.Router
var ren *render.Render

func init() {
	ren = render.New(render.Options{})

	router = mux.NewRouter()

	router.HandleFunc("/dayrate", func(w http.ResponseWriter, req *http.Request) {

		cacheRate, err := eurobank.GetDayRate()

		if err != nil {
			log.Println("The eurobank query failed:", err)
			http.Error(w, "Err", http.StatusInternalServerError)
			return
		}

		ren.JSON(w, http.StatusOK, cacheRate)

	})

	router.HandleFunc("/dayrate90", func(w http.ResponseWriter, req *http.Request) {

		rates, err := eurobank.Get90DayRates()
		if err != nil {
			log.Println("Could not fetch data:", err)
			http.Error(w, "err", http.StatusInternalServerError)
			return
		}

		ren.JSON(w, http.StatusOK, rates)
	})
}

func main() {

	n := negroni.Classic()

	n.Use(auth.Basic("foo", "bar"))
	n.UseHandler(router)
	n.Run(":" + os.Getenv("PORT"))
}
