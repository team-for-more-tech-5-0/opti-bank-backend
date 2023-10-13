package atm

import (
	"github.com/Spawn4real/hackthon_more.tech_5.0.git/internal/models/service"
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

func NewAtmServices(blind, QrRead, Wheelchair, SupportsEur, SupportsRub, SupportsUsd, NfcForBankCards, SupportsChargeRub service.Service) *Services {
	return &Services{
		Blind:             blind,
		QrRead:            QrRead,
		Wheelchair:        Wheelchair,
		SupportsEur:       SupportsEur,
		SupportsRub:       SupportsRub,
		SupportsUsd:       SupportsUsd,
		NfcForBankCards:   NfcForBankCards,
		SupportsChargeRub: SupportsChargeRub,
	}
}
