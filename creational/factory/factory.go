package factory

import (
	"errors"
	"fmt"
)

const (
	PaymentTypeCreditCard = iota
	PaymentTypePaypal
	PaymentTypeBankTransfer
)

type PaymentMethod interface {
	ProcessPayment(amount float64) string
}

type paymentType int8

type creditCard struct{}

func (c creditCard) ProcessPayment(amount float64) string {
	return processPayment("credit card", amount)
}

func processPayment(paymentName string, amount float64) string {
	return fmt.Sprintf("processed %s payment of amount %.2f", paymentName, amount)
}

type paypal struct{}

func (p paypal) ProcessPayment(amount float64) string {
	return processPayment("paypal", amount)
}

type bankTransfer struct{}

func (p bankTransfer) ProcessPayment(amount float64) string {
	return processPayment("bank transfer", amount)
}

var unsupportedPaymentErr = errors.New("unsupported payment type")

func CreatePaymentMethod(t paymentType) (PaymentMethod, error) {
	switch t {
	case PaymentTypeCreditCard:
		return creditCard{}, nil
	case PaymentTypePaypal:
		return paypal{}, nil
	case PaymentTypeBankTransfer:
		return bankTransfer{}, nil
	}
	return nil, unsupportedPaymentErr
}
