package atm

import (
	"github.com/team-for-more-tech-5-0/opti-bank-backend.git/internal/models/service"
)

type Services struct {
	Blind             service.Service `json:"blind"`
	QrRead            service.Service `json:"qrRead"`
	Wheelchair        service.Service `json:"wheelchair"`
	SupportsEur       service.Service `json:"supportsEur"`
	SupportsRub       service.Service `json:"supportsRub"`
	SupportsUsd       service.Service `json:"supportsUsd"`
	NfcForBankCards   service.Service `json:"nfcForBankCards"`
	SupportsChargeRub service.Service `json:"supportsChargeRub"`
}
