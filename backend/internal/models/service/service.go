package service

import (
	"encoding/json"
	"errors"
)

type Services struct {
	ServiceForBusinesses   *ServicesForBusinesses  `json:"servicesForBusinesses"`
	ServicesForIndividuals *ServicesForIndividuals `json:"servicesForIndividuals"`
}

type ServicesForBusinesses struct {
	SMEservices                         Service `json:"SMEservices"`
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

type ServiceMap map[string]interface{}

func (p *Services) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	m, ok := i.(map[string]interface{})
	if !ok {
		return errors.New("Type assertion .(map[string]interface{}) failed.")
	}

	p.ServiceForBusinesses = &ServicesForBusinesses{}
	p.ServicesForIndividuals = &ServicesForIndividuals{}

	if servicesForBusinesses, ok := m["servicesForBusinesses"].(map[string]interface{}); ok {
		p.ServiceForBusinesses.SMEservices = Service{}
		p.ServiceForBusinesses.BusinessFinancing = Service{}
		p.ServiceForBusinesses.CorporateAccounts = Service{}
		p.ServiceForBusinesses.CorporateCreditCards = Service{}
		p.ServiceForBusinesses.TransactionAndPaymentServices = Service{}
		p.ServiceForBusinesses.InternationalOperationsServices = Service{}
		p.ServiceForBusinesses.LiquidityAndFinancialRiskManagement = Service{}

		if SMEservices, ok := servicesForBusinesses["SMEservices"].(map[string]interface{}); ok {
			p.ServiceForBusinesses.SMEservices.ServiceActivity = SMEservices["serviceActivity"].(string)
			p.ServiceForBusinesses.SMEservices.ServiceCapability = SMEservices["serviceCapability"].(string)
		}

		if businessFinancing, ok := servicesForBusinesses["businessFinancing"].(map[string]interface{}); ok {
			p.ServiceForBusinesses.BusinessFinancing.ServiceActivity = businessFinancing["serviceActivity"].(string)
			p.ServiceForBusinesses.BusinessFinancing.ServiceCapability = businessFinancing["serviceCapability"].(string)
		}

		if corporateAccounts, ok := servicesForBusinesses["corporateAccounts"].(map[string]interface{}); ok {
			p.ServiceForBusinesses.CorporateAccounts.ServiceActivity = corporateAccounts["serviceActivity"].(string)
			p.ServiceForBusinesses.CorporateAccounts.ServiceCapability = corporateAccounts["serviceCapability"].(string)
		}

		if corporateCreditCards, ok := servicesForBusinesses["corporateCreditCards"].(map[string]interface{}); ok {
			p.ServiceForBusinesses.CorporateCreditCards.ServiceActivity = corporateCreditCards["serviceActivity"].(string)
			p.ServiceForBusinesses.CorporateCreditCards.ServiceCapability = corporateCreditCards["serviceCapability"].(string)
		}

		if transactionAndPaymentServices, ok := servicesForBusinesses["transactionAndPaymentServices"].(map[string]interface{}); ok {
			p.ServiceForBusinesses.TransactionAndPaymentServices.ServiceActivity = transactionAndPaymentServices["serviceActivity"].(string)
			p.ServiceForBusinesses.TransactionAndPaymentServices.ServiceCapability = transactionAndPaymentServices["serviceCapability"].(string)
		}

		if internationalOperationsServices, ok := servicesForBusinesses["internationalOperationsServices"].(map[string]interface{}); ok {
			p.ServiceForBusinesses.InternationalOperationsServices.ServiceActivity = internationalOperationsServices["serviceActivity"].(string)
			p.ServiceForBusinesses.InternationalOperationsServices.ServiceCapability = internationalOperationsServices["serviceCapability"].(string)
		}

		if liquidityAndFinancialRiskManagement, ok := servicesForBusinesses["liquidityAndFinancialRiskManagement"].(map[string]interface{}); ok {
			p.ServiceForBusinesses.LiquidityAndFinancialRiskManagement.ServiceActivity = liquidityAndFinancialRiskManagement["serviceActivity"].(string)
			p.ServiceForBusinesses.LiquidityAndFinancialRiskManagement.ServiceCapability = liquidityAndFinancialRiskManagement["serviceCapability"].(string)
		}
	}

	if servicesForIndividuals, ok := m["servicesForIndividuals"].(map[string]interface{}); ok {
		p.ServicesForIndividuals.MortgageLoans = Service{}
		p.ServicesForIndividuals.LoansAndCredits = Service{}
		p.ServicesForIndividuals.DepositsAndSavings = Service{}
		p.ServicesForIndividuals.InvestmentServices = Service{}
		p.ServicesForIndividuals.BankAccountsAndCards = Service{}
		p.ServicesForIndividuals.OnlineBankingAndMobileApp = Service{}

		if mortgageLoans, ok := servicesForIndividuals["mortgageLoans"].(map[string]interface{}); ok {
			p.ServicesForIndividuals.MortgageLoans.ServiceActivity = mortgageLoans["serviceActivity"].(string)
			p.ServicesForIndividuals.MortgageLoans.ServiceCapability = mortgageLoans["serviceCapability"].(string)
		}

		if loansAndCredits, ok := servicesForIndividuals["loansAndCredits"].(map[string]interface{}); ok {
			p.ServicesForIndividuals.LoansAndCredits.ServiceActivity = loansAndCredits["serviceActivity"].(string)
			p.ServicesForIndividuals.LoansAndCredits.ServiceCapability = loansAndCredits["serviceCapability"].(string)
		}

		if depositsAndSavings, ok := servicesForIndividuals["depositsAndSavings"].(map[string]interface{}); ok {
			p.ServicesForIndividuals.DepositsAndSavings.ServiceActivity = depositsAndSavings["serviceActivity"].(string)
			p.ServicesForIndividuals.DepositsAndSavings.ServiceCapability = depositsAndSavings["serviceCapability"].(string)
		}

		if investmentServices, ok := servicesForIndividuals["investmentServices"].(map[string]interface{}); ok {
			p.ServicesForIndividuals.InvestmentServices.ServiceActivity = investmentServices["serviceActivity"].(string)
			p.ServicesForIndividuals.InvestmentServices.ServiceCapability = investmentServices["serviceCapability"].(string)
		}

		if bankAccountsAndCards, ok := servicesForIndividuals["bankAccountsAndCards"].(map[string]interface{}); ok {
			p.ServicesForIndividuals.BankAccountsAndCards.ServiceActivity = bankAccountsAndCards["serviceActivity"].(string)
			p.ServicesForIndividuals.BankAccountsAndCards.ServiceCapability = bankAccountsAndCards["serviceCapability"].(string)
		}

		if onlineBankingAndMobileApp, ok := servicesForIndividuals["onlineBankingAndMobileApp"].(map[string]interface{}); ok {
			p.ServicesForIndividuals.OnlineBankingAndMobileApp.ServiceActivity = onlineBankingAndMobileApp["serviceActivity"].(string)
			p.ServicesForIndividuals.OnlineBankingAndMobileApp.ServiceCapability = onlineBankingAndMobileApp["serviceCapability"].(string)
		}
	}

	return nil
}
