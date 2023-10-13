package config

type ServiceType int

const (
	DebitCard ServiceType = iota
	CreditCard
	Mortgage
	SavingsAccount
	AutoLoan
	PersonalLoan
	Investments
	InsuranceProducts
	GovernmentServices
	PensionPlans
	Safes
	LettersOfCredit
	PawnShop
	CertificatesAndApplications
	BankruptcyServices
)
