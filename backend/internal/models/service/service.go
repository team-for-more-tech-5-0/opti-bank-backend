package service

type ServicesForBusinesses struct {
	SMEServices                         Service `json:"SMEservices"`
	BusinessFinancing                   Service `json:"businessFinancing"`
	CorporateAccounts                   Service `json:"corporateAccounts"`
	CorporateCreditCards                Service `json:"corporateCreditCards"`
	TransactionAndPaymentServices       Service `json:"transactionAndPaymentServices"`
	InternationalOperationsServices     Service `json:"internationalOperationsServices"`
	LiquidityAndFinancialRiskManagement Service `json:"liquidityAndFinancialRiskManagement"`
}

type ServicesForIndividuals struct {
	MortgageLoans             Service `json:"mortgageLoans"`
	LoansAndCredits           Service `json:"loansAndCredits"`
	DepositsAndSavings        Service `json:"depositsAndSavings"`
	InvestmentServices        Service `json:"investmentServices"`
	BankAccountsAndCards      Service `json:"bankAccountsAndCards"`
	OnlineBankingAndMobileApp Service `json:"onlineBankingAndMobileApp"`
}

type Service struct {
	ServiceActivity   string `json:"serviceActivity"`
	ServiceCapability string `json:"serviceCapability"`
}

type Services struct {
	ServiceForBusinesses   ServicesForBusinesses  `json:"servicesForBusinesses"`
	ServicesForIndividuals ServicesForIndividuals `json:"servicesForIndividuals"`
}
