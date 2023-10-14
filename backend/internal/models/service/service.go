package service

type Service struct {
	ServicesForBusinesses struct {
		SMEservices struct {
			ServiceActivity   string `json:"serviceActivity"`
			ServiceCapability string `json:"serviceCapability"`
		} `json:"SMEservices"`
		BusinessFinancing struct {
			ServiceActivity   string `json:"serviceActivity"`
			ServiceCapability string `json:"serviceCapability"`
		} `json:"businessFinancing"`
		CorporateAccounts struct {
			ServiceActivity   string `json:"serviceActivity"`
			ServiceCapability string `json:"serviceCapability"`
		} `json:"corporateAccounts"`
		CorporateCreditCards struct {
			ServiceActivity   string `json:"serviceActivity"`
			ServiceCapability string `json:"serviceCapability"`
		} `json:"corporateCreditCards"`
		TransactionAndPaymentServices struct {
			ServiceActivity   string `json:"serviceActivity"`
			ServiceCapability string `json:"serviceCapability"`
		} `json:"transactionAndPaymentServices"`
		InternationalOperationsServices struct {
			ServiceActivity   string `json:"serviceActivity"`
			ServiceCapability string `json:"serviceCapability"`
		} `json:"internationalOperationsServices"`
		LiquidityAndFinancialRiskManagement struct {
			ServiceActivity   string `json:"serviceActivity"`
			ServiceCapability string `json:"serviceCapability"`
		} `json:"liquidityAndFinancialRiskManagement"`
	} `json:"servicesForBusinesses"`
	ServicesForIndividuals struct {
		MortgageLoans struct {
			ServiceActivity   string `json:"serviceActivity"`
			ServiceCapability string `json:"serviceCapability"`
		} `json:"mortgageLoans"`
		LoansAndCredits struct {
			ServiceActivity   string `json:"serviceActivity"`
			ServiceCapability string `json:"serviceCapability"`
		} `json:"loansAndCredits"`
		DepositsAndSavings struct {
			ServiceActivity   string `json:"serviceActivity"`
			ServiceCapability string `json:"serviceCapability"`
		} `json:"depositsAndSavings"`
		InvestmentServices struct {
			ServiceActivity   string `json:"serviceActivity"`
			ServiceCapability string `json:"serviceCapability"`
		} `json:"investmentServices"`
		BankAccountsAndCards struct {
			ServiceActivity   string `json:"serviceActivity"`
			ServiceCapability string `json:"serviceCapability"`
		} `json:"bankAccountsAndCards"`
		OnlineBankingAndMobileApp struct {
			ServiceActivity   string `json:"serviceActivity"`
			ServiceCapability string `json:"serviceCapability"`
		} `json:"onlineBankingAndMobileApp"`
	} `json:"servicesForIndividuals"`
}
