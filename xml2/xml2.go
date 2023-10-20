package xml2

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"stm32pin/model"
)

// our struct which contains the complete
// array of all Users in the file
type Balls struct {
	XMLName xml.Name `xml:"Mcu"`
	Balls   []Ball   `xml:"Pin"`
}

// the user struct, this contains our
// Type attribute, our user's name and
// a social struct which will contain all
// our social links
type Ball struct {
	XMLName xml.Name `xml:"Pin"`
	Name    string   `xml:"Name,attr"`
	Pin     string   `xml:"Position,attr"`
}

func ReadXML(fileName string) []model.Ball {
	data := []model.Ball{}
	// Open our xmlFile
	xmlFile, err := os.Open(fileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Users array
	var users Balls
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &users)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(users.Balls); i++ {
		//fmt.Println(users.Balls[i].Pin + " " + users.Balls[i].Name)
		temp := model.Ball{
			Pin:  users.Balls[i].Pin,
			Name: users.Balls[i].Name,
		}
		data = append(data, temp)
	}
	return data
}
