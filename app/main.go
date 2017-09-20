package main

import (
  "fmt"
  "bytes"
  "log"
	"net/http"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
//var router *gin.Engine

const (
  MongoServerAddr = "mongodb://mng:lau@ds163053.mlab.com:63053/pulsedb"
  DebugMode = true
  MonotonicBehavior = true
)

type OigModel struct {
	LASTNAME string `bson:"LASTNAME"`
	FIRSTNAME string `bson:"FIRSTNAME"`  
	MIDNAME string `bson:"MIDNAME"`   
	BUSNAME string `bson:"BUSNAME"`  
	GENERAL string `bson:"GENERAL"`   
	SPECIALTY string `bson:"SPECIALTY"`  
	UPIN string `bson:"UPIN"`   
	NPI int `bson:"NPI"`
	DOB int `bson:"DOB"`   
	ADDRESS string `bson:"ADDRESS"`
	CITY string `bson:"CITY"`
	STATE string `bson:"STATE"`
	ZIP int `bson:"ZIP"`
	EXCLTYPE string `bson:"EXCLTYPE"`
	EXCLDATE int `bson:"EXCLDATE"`
	REIN int `bson:"REIN"`
	WAIVER int `bson:"WAIVER"`
	WVRSTATE string `bson:"WVRSTATE"`       
}

var (
  mgoSession *mgo.Session
  err error
  databaseName = "pulsedb"
  buf    bytes.Buffer
  logger = log.New(&buf, "MGO INFO: ", log.Lshortfile)
)

func main() {
  router := initializeRoutes()
  router.Run()
}

func getSession() *mgo.Session {

  if mgoSession == nil {
    var err error
    mgoSession, err = mgo.Dial(MongoServerAddr)
    mgo.SetDebug(DebugMode)
    mgo.SetLogger(logger)
    gin.SetMode(gin.ReleaseMode)
    if err != nil {
        panic(err)
    }
  }
  //defer mgoSession.Close()

  // Optional. Switch the session to a monotonic behavior.
 // mgoSession.SetMode(gin.ReleaseMode, true) //mgo.Monotonic, MonotonicBehavior)
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
func SearchPerson (q interface{}, skip int, limit int) (searchResults []OigModel, searchErr string) {
  searchErr     = ""
  searchResults = []OigModel{}
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

func GetPersonByName (lastName string, firstName string, skip int, limit int) (searchResults []OigModel, searchErr string) {
  searchResults, searchErr = SearchPerson(bson.M{
    "LASTNAME": bson.RegEx{"^"+lastName, "i"},
    "FIRSTNAME": bson.RegEx{"^"+firstName, "i"}}, skip, limit)
  return
}

func search(d *gin.Context) {
 
    firstname := d.Query("firstname")
    lastname := d.Query("lastname")
    dob := d.Query("dob")
  
    oig := *mgoSession.DB("pulsedb").C("oigs")
    sam := *mgoSession.DB("pulsedb").C("sams")
    ofac := *mgoSession.DB("pulsedb").C("oigs")
  
    result := OigModel{}
    err := oig.Find(bson.M{"LASTNAME": "JOHNSON"}).One(&result)
    if err != nil {
      fmt.Println("find error: ", err)
    }
    
    OIGCounter, err := oig.Count() 
     if err != nil {
       fmt.Println("oig error: ", err)
     }
     SAMCounter, err := sam.Count() 
     if err != nil {
       fmt.Println("sam error: ", err)
     }
     OFACCounter, err := ofac.Count() 
     if err != nil {
       fmt.Println("ofac error: ", err)
     }
  
    d.JSON(200, gin.H{"Result": result,"LastName": result.LASTNAME, "OIG": OIGCounter,"SAM": SAMCounter,"OFAC": OFACCounter})
    fmt.Println("OIGcounter:", OIGCounter)


    d.JSON(200, gin.H{"firstname": firstname,"lastname": lastname, "dob": dob})
    fmt.Println(firstname + " " + lastname + " " + dob)
}

func initializeRoutes() *gin.Engine {
  r := gin.Default()

	//r.LoadHTMLGlob("templates/**/*.html")
	
	r.GET("/", func(c *gin.Context) {
		c.HTML(404, "404 - Page Not Found", nil)
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(404, "404 - Page Not Found", nil)
	})

	r.GET("/ping", func(c *gin.Context) {
	//	initDb(c)
	//	c.JSON(200, gin.H{"message": "pong",})
	})
  
  r.GET("/search", func(c *gin.Context) {
    searchResults, err := GetPersonByName( c.Query("lastname"), c.Query("firstname"), 0, 0)
    if err != "" {
      fmt.Print(err)
    } 
      fmt.Print(searchResults)
      c.JSON(200, gin.H{"status": "ok"})
      c.JSON(200, gin.H{"data": searchResults})
  })
  
	admin := r.Group("/admin")
	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html", nil)
	})

	admin.GET("/dashboard", func(c *gin.Context) {

		c.HTML(http.StatusNotFound, "404 - Page Not Found", nil)
	})

	r.Static("/public", "./public")

	return r
}
