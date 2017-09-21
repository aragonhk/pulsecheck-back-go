package database

import (
	"fmt"
	"bytes"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/aragonhk/pulsecheck-back-go/app/models"
)

const (
	MongoServerAddr = "mongodb://mng:lau@ds163053.mlab.com:63053/pulsedb"
	databaseName = "pulsedb"
	DebugMode = true
	MonotonicBehavior = true
  )

var (
  mgoSession *mgo.Session
  err error
  buf    bytes.Buffer
  logger = log.New(&buf, "MGO INFO: ", log.Lshortfile)
)

func getSession() *mgo.Session {
	
	if mgoSession == nil {
	var err error
	mgoSession, err = mgo.Dial(MongoServerAddr)
	mgo.SetDebug(DebugMode)
	mgo.SetLogger(logger)
	
	if err != nil {
		panic(err)
	}
	}

	// Optional. Switch the session to a monotonic behavior.
	mgoSession.SetMode(mgo.Monotonic, MonotonicBehavior)
	fmt.Print(&buf)
	return mgoSession.Clone()
}
	
func withCollection(collection string, s func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB(databaseName).C(collection)
	return s(c)
}
//The skip and limit parameters are optional in that if skip is set to zero, 
//it is effectively asking for all the results, and, similarly, if limit is set to an integer less than zero,
// it is ignored in the query that gets invoked inside the withCollection() function. 
func SearchPerson (q interface{}, skip int, limit int) (searchResults []models.OigModel, searchErr string) {
	searchErr     = ""
	searchResults = []models.OigModel{}
	query := func(c *mgo.Collection) error {
		fn := c.Find(q).Skip(skip).Limit(limit).All(&searchResults)
		if limit < 0 {
			fn = c.Find(q).Skip(skip).All(&searchResults)
		}
		return fn
	}
	search := func() error {
	return withCollection("oigs", query)
	}
	err := search()
	if err != nil {
		searchErr = "Database Error"
		fmt.Println("searchErr:" + searchErr)
	}
	return
}

func GetPersonByName (lastName string, firstName string, skip int, limit int) (searchResults []models.OigModel, searchErr string) {
	searchResults, searchErr = SearchPerson(bson.M{
	  "LASTNAME": bson.RegEx{"^"+lastName, "i"},
	  "FIRSTNAME": bson.RegEx{"^"+firstName, "i"}}, skip, limit)
	return
}