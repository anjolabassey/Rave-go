package rave

import (
	"fmt"
	"testing"
	"time"
)

func TestTransfer_InitiateSingleTransfer(t *testing.T) {
	payloads := []SinglePaymentData{
		{
			AccountBank:   "044",
			AccountNumber: "0690000044",
			Amount:        500,
			SecKey:        "FLWSECK-xxxxxxxxxxxxxxxxxxxxx-X",
			Narration:     "New transfer",
			Currency:      "NGN",
			Reference:     time.Now().String(),
		},
		{
			AccountBank:     "044",
			AccountNumber:   "0690000031",
			Amount:          500,
			SecKey:          "FLWSECK-xxxxxxxxxxxxxxxxxxxxx-X",
			Narration:       "New transfer",
			Currency:        "NGN",
			Reference:       time.Now().String(),
			BeneficiaryName: "Kwame Adew",
		},
	}

	for _, payload := range payloads {
		err, response := Transfer{r}.InitiateSingleTransfer(payload)
		if err != nil {
			t.Fatalf("An error occurred while testing single transfer: %v", err)
		}
		if response["status"] != "success" {
			t.Fatalf("Single transfer wasn't successful: %v", response)
		}
	}
}

func TestTransfer_InitiateBulkTransfer(t *testing.T) {
	payloads := []BulkPaymentData{
		{
			SecKey: "FLWSECK-xxxxxxxxxxxxxxxxxxxxx-X",
			Title:  "May Staff Salary",
			BulkData: []map[string]string{
				{
					"Bank":           "044",
					"Account Number": "0690000032",
					"Amount":         "500",
					"Currency":       "NGN",
					"Narration":      "Bulk transfer 1",
					"reference":      time.Now().String(),
				},
				{
					"Bank":           "044",
					"Account Number": "0690000034",
					"Amount":         "500",
					"Currency":       "NGN",
					"Narration":      "Bulk transfer 2",
					"reference":      time.Now().String(),
				},
			},
		},
	}

	for _, payload := range payloads {
		err, response := Transfer{r}.InitiateBulkTransfer(payload)
		if err != nil {
			t.Fatalf("An error occurred while testing single transfer: %v", err)
		}
		if response["status"] != "success" {
			t.Fatalf("Single transfer wasn't successful: %v", response)
		}
	}
}

func TestTransfer_FetchAllTransfers(t *testing.T) {
	err, response := Transfer{r}.FetchAllTransfers("")
	if err != nil {
		t.Fatalf("An error occurred while testing single transfer: %v", err)
	}
	if response["status"] != "success" {
		t.Fatalf("Single transfer wasn't successful: %v", response)
	}
}

func TestTransfer_FetchTransfer(t *testing.T) {
	reference := "kkkkkkkkkkkkk"
	err, response := Transfer{r}.FetchTransfer(reference)
	if err != nil {
		t.Fatalf("An error occurred while testing single transfer: %v", err)
	}
	if response["status"] != "success" {
		t.Fatalf("Single transfer wasn't successful: %v", response)
	}
}

func TestGetBulkTransferStatus(t *testing.T) {
	batchIDs := [2]string{"634", "635"}

	for _, batchID := range batchIDs {
		error, response := Transfer{
			r,
		}.GetBulkTransferStatus(batchID)
		if error != nil {
			t.Fatalf("Transfer failed with error %v", error)
		}
		if response["status"] != "success" {
			t.Fatalf("Transfer status: %v, Details: %v", response["status"], response)
		}
	}
}

func TestGetTransferFees(t *testing.T) {
	currencies := [4]string{"NGN", "USD", "GHS", "KES"}

	for _, currency := range currencies {
		error, response := Transfer{
			r,
		}.GetTransferFee(currency)
		if error != nil {
			t.Fatalf("Transfer failed with error %v", error)
		}
		if response["status"] != "success" {
			t.Fatalf("Transfer status: %v, Details: %v", response["status"], response)
		}
	}
}

func TestGetRaveBalance(t *testing.T) {
	currencies := [4]string{"NGN", "USD", "GHS", "KES"}

	for _, currency := range currencies {
		error, response := Transfer{
			r,
		}.GetRaveBalance(currency)
		if error != nil {
			t.Fatalf("Transfer failed with error %v", error)
		}
		if response["status"] != "success" {
			t.Fatalf("Transfer status: %v, Details: %v", response["status"], response)
		}
	}
}

func TestResolveAccount(t *testing.T) {
	payloads := []AccountResolveData{
		{
			RecipientAccount: "0690000034",
			DestBankCode:     "044",
			PublicKey:        r.GetPublicKey(),
		},
		{
			RecipientAccount: "0690000034",
			DestBankCode:     "044",
			PublicKey:        r.GetPublicKey(),
			Currency:         "NGN",
			Country:          "NG",
		},
	}

	for _, payload := range payloads {
		error, response := Transfer{
			r,
		}.ResolveAccount(payload)
		fmt.Println(response)
		if error != nil {
			t.Fatalf("Transfer failed with error %v", error)
		}
		if response["status"] != "success" {
			t.Fatalf("Transfer status: %v, Details: %v", response["status"], response)
		}
	}
}
