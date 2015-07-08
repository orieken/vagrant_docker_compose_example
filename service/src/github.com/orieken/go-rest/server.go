package main

import (
	"os"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/orieken/go-rest/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()

	uc := controllers.NewCarController(getSession())

	r.GET("/test", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Winning!\n")
	})

	r.GET("/car/:id", uc.GetCar)

	r.POST("/car", uc.CreateCar)

	r.DELETE("/car/:id", uc.DeleteCar)

	http.ListenAndServe("localhost:7000", r)
}

func getSession() *mgo.Session {
	uri := os.Getenv("DB_PORT_27017_TCP_ADDR")
	if uri == "" {
		fmt.Println("no connection string provided")
		os.Exit(1)
	}

	s, err:= mgo.Dial(uri)

	if err != nil {
		panic(err)
	}
	return s
}