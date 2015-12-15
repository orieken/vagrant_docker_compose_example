package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/orieken/go-rest/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	CarController struct {
		session *mgo.Session
	}
)

func NewCarController(s *mgo.Session) *CarController {
	return &CarController{s}
}


func (cc CarController) GetCar(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	c := models.Car{}

	if err := cc.session.DB("go_rest").C("cars").FindId(oid).One(&c); err != nil {
		w.WriteHeader(404)
		return
	}

	cj, _ := json.Marshal(c)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", cj)
}

func (cc CarController) CreateCar(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c := models.Car{}

	json.NewDecoder(r.Body).Decode(&c)

	c.Id = bson.NewObjectId()

	cc.session.DB("go_rest").C("cars").Insert(c)

	cj, _ := json.Marshal(c)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", cj)
}

func (cc CarController) DeleteCar(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := cc.session.DB("go_rest").C("cars").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(200)
}