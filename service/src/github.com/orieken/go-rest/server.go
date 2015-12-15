package main


import (
//	"os"
//	"time"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
//	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
//	"github.com/orieken/go-rest/controllers"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func Test(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Winning!\n")

}

func main() {
	router := httprouter.New()


	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/test", Test)

//	carController := controllers.NewCarController(getSession())
//	Car routes
//	router.GET("/car/:id", carController.GetCar)
//	router.POST("/car/new", carController.CreateCar)
//	router.DELETE("/car/:id", carController.DeleteCar)

	log.Fatal(http.ListenAndServe("0.0.0.0:49200", router))
}

//type Ping struct {
//	Id   bson.ObjectId `bson:"_id"`
//	Time time.Time     `bson:"time"`
//}

//func getSession() *mgo.Session {
//	uri := os.Getenv("DATABASE_PORT_27017_TCP_ADDR")
//
//	session, _ := mgo.Dial(uri)
//	db := session.DB("go_rest")
//
//	ping := Ping{
//		Id:   bson.NewObjectId(),
//		Time: time.Now(),
//	}
//
//	db.C("pings").Insert(ping)
//
//	// get all records
//	pings := []Ping{}
//	db.C("pings").Find(nil).All(&pings)
//
//	fmt.Println(pings)
//
//	if uri == "" {
//		fmt.Println("no connection string provided")
//		os.Exit(1)
//	}
//
//	s, err:= mgo.Dial(uri)
//
//	if err != nil {
//		panic(err)
//	}
//	return s
//}