package main

import (
	"os"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/orieken/go-rest/controllers"
	"gopkg.in/mgo.v2"
	"time"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	r := httprouter.New()

	uc := controllers.NewUserController(getSession())

	r.GET("/test", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Winning!\n")
	})

	r.GET("/user/:id", uc.GetUser)

	r.POST("/user", uc.CreateUser)

	r.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe("localhost:7000", r)
}

type Ping struct {
	Id   bson.ObjectId `bson:"_id"`
	Time time.Time     `bson:"time"`
}

func getSession() *mgo.Session {
	uri := os.Getenv("DATABASE_PORT_27017_TCP_ADDR")

	session, _ := mgo.Dial(uri)
	db := session.DB("go_rest")

	ping := Ping{
		Id:   bson.NewObjectId(),
		Time: time.Now(),
	}

	db.C("pings").Insert(ping)

	// get all records
	pings := []Ping{}
	db.C("pings").Find(nil).All(&pings)

	fmt.Println(pings)

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