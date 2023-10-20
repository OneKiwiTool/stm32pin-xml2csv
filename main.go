package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"stm32pin/csv2"
	"stm32pin/xml2"
	"strings"
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

func fileNameWithoutExtension(fileName string) string {
	if pos := strings.LastIndexByte(fileName, '.'); pos != -1 {
		return fileName[:pos]
	}
	return fileName
}
func main() {

	xmlpath := ""
	path := flag.String("f", "", "file bom")
	flag.Parse()
	fmt.Println("Input:", *path)
	xmlpath = *path

	balls := xml2.ReadXML(xmlpath)
	csvpath := fileNameWithoutExtension(xmlpath)

	csvpath = csvpath + ".csv"
	csv2.CreateCSVFile(csvpath, balls)
	fmt.Println("Output:", csvpath)
}
