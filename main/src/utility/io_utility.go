package utility

import (
	"appointment-notification-sender/main/src/models"
	"bytes"
	"encoding/csv"
	"github.com/xuri/excelize/v2"
	"io"
	"log"
	"strings"
)

func ParseCSVFromMemory(data []byte) ([]models.Customer, error) {
	reader := csv.NewReader(bytes.NewReader(data))
	// Assuming the CSV file has a header row
	_, err := reader.Read()
	if err != nil {
		return nil, err
	}

	var customers []models.Customer

	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Parse CSV record
		customer := models.Customer{
			CustomerID: record[0],
			FirstName:  record[1],
			LastName:   record[2],
			FullName:   record[3],
			Email:      record[4],
			CellNumber: record[5],
			IsSMS:      strings.ToUpper(record[6]) == "Y",
			IsEmail:    strings.ToUpper(record[7]) == "Y",
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func ParseXLSXFromMemory(data []byte) ([]models.Customer, error) {
	xlFile, err := excelize.OpenReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	rows, err := xlFile.GetRows("Sheet1")
	if err != nil {
		return nil, err
	}

	var customers []models.Customer

	for i, row := range rows {
		if i == 0 {
			continue
		}
		log.Println(row[3])
		// Parse XLSX row
		customer := models.Customer{
			CustomerID: row[0],
			FirstName:  row[1],
			LastName:   row[2],
			FullName:   row[3],
			Email:      row[4],
			CellNumber: row[5],
			IsSMS:      strings.ToUpper(row[6]) == "Y",
			IsEmail:    strings.ToUpper(row[7]) == "Y",
		}
		customers = append(customers, customer)
	}

	return customers, nil
}
