package csv2

import (
	"log"
	"os"
	"stm32pin/model"

	"github.com/tushar2708/altcsv"
)

func CreateCSVFile(fileName string, balls []model.Ball) {
	headers := []string{"Pin", "Unit", "Type", "Name"}
	fileWtr, _ := os.Create(fileName)
	csvWtr := altcsv.NewWriter(fileWtr)
	csvWtr.Quote = '"'      // use | as "quote"
	csvWtr.AllQuotes = true // surround each field with '|'
	csvWtr.Write(headers)
	for _, v := range balls {
		row := []string{v.Pin, "", "passive", v.Name}
		if err := csvWtr.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
	csvWtr.Flush()
	fileWtr.Close()
}
