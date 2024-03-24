package models

// Customer struct represents customer data
type Customer struct {
	CustomerID  string
	FirstName   string
	LastName    string
	FullName    string
	Email       string
	CellNumber  string
	IsSMS       bool
	IsEmail     bool
	IsSMSSent   bool
	IsEmailSent bool
}
