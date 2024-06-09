package factory

import "testing"

func TestPaymentFactory(t *testing.T) {
	t.Parallel()
	var (
		pm  PaymentMethod
		got string
		err error
	)
	for _, tc := range []struct {
		paymentType paymentType
		amount      float64
		want        string
	}{
		{
			paymentType: PaymentTypeCreditCard,
			amount:      4.0,
			want:        "processed credit card payment of amount 4.00",
		},
		{
			paymentType: PaymentTypePaypal,
			amount:      1.9,
			want:        "processed paypal payment of amount 1.90",
		},
		{
			paymentType: PaymentTypeBankTransfer,
			amount:      14.11,
			want:        "processed bank transfer payment of amount 14.11",
		},
	} {
		pm, err = CreatePaymentMethod(tc.paymentType)
		if err != nil {
			t.Fatalf("want nil error, got %v", err)
		}

		if got = pm.ProcessPayment(tc.amount); got != tc.want {
			t.Errorf("got '%s', want '%s'", got, tc.want)
		}
	}
}

func TestPaymentFactory_UnsupportedPayment(t *testing.T) {
	pm, err := CreatePaymentMethod(-1)
	if err != unsupportedPaymentErr {
		t.Errorf("want error %v, got %v", unsupportedPaymentErr, err)
	}
	if pm != nil {
		t.Errorf("want nil paymentm method, got %v", pm)
	}
}
