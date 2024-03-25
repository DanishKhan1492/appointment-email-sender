package utility

import (
	"appointment-notification-sender/main/src/models"
	"bytes"
	"encoding/csv"
	"github.com/xuri/excelize/v2"
	"io"
	"strings"
	"sync"
)

func ParseCSVFromMemory(data []byte) ([]models.Customer, error) {
	reader := csv.NewReader(bytes.NewReader(data))
	// Assuming the CSV file has a header row
	_, err := reader.Read()
	if err != nil {
		return nil, err
	}

	customersCh := make(chan models.Customer)
	errCh := make(chan error)
	var wg sync.WaitGroup

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		wg.Add(1)
		go func(record []string) {
			defer wg.Done()
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
			customersCh <- customer
		}(record)
	}

	go func() {
		wg.Wait()
		close(customersCh)
		close(errCh)
	}()

	var customers []models.Customer
	for customer := range customersCh {
		customers = append(customers, customer)
	}

	if err := <-errCh; err != nil {
		return nil, err
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

	customersCh := make(chan models.Customer)
	errCh := make(chan error)
	var wg sync.WaitGroup

	for i, row := range rows {
		if i == 0 {
			continue
		}
		wg.Add(1)
		go func(row []string) {
			defer wg.Done()
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
			customersCh <- customer
		}(row)
	}

	go func() {
		wg.Wait()
		close(customersCh)
		close(errCh)
	}()

	var customers []models.Customer
	for customer := range customersCh {
		customers = append(customers, customer)
	}

	if err := <-errCh; err != nil {
		return nil, err
	}

	return customers, nil
}
