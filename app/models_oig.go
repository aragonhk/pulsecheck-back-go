package main

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