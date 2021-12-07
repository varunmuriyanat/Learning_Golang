package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	type AccountAttributes struct {
		AccountClassification   string   `json:"account_classification,omitempty"`
		AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
		AccountNumber           string   `json:"account_number,omitempty"`
		AlternativeNames        []string `json:"alternative_name,omitempty"`
		BankID                  string   `json:"bank_id,omitempty"`
		BankIDCode              string   `json:"bank_id_code,omitempty"`
		BaseCurrency            string   `json:"base_currency,omitempty"`
		Bic                     string   `json:"bic,omitempty"`
		Country                 *string  `json:"country,omitempty"`
		Iban                    string   `json:"iban,omitempty"`
		JointAccount            *bool    `json:"joint_account,omitempty"`
		Name                    []string `json:"name,omitempty"`
		SecondaryIdentification string   `json:"secondary_identification,omitempty"`
		Status                  *string  `json:"status,omitempty"`
		Switched                *bool    `json:"switched,omitempty"`
	}

	type Account struct {
		Attributes     *AccountAttributes `json:"attributes,omitempty"`
		ID             string             `json:"id,omitempty"`
		OrganisationID string             `json:"organisation_id,omitempty"`
		Type           string             `json:"type,omitempty"`
		Version        *int64             `json:"version,omitempty"`
	}

	type Data struct {
		Data Account `json:"data"`
	}

	accountNames := []string{"Varun Muriyanat"}
	alternativeNames := []string{"VM"}
	accountMatchingOptOut := false
	jointAccount := false
	country := "GB"

	accountAttributes := AccountAttributes{
		AccountClassification:   "Personal",
		AccountMatchingOptOut:   &accountMatchingOptOut,
		AlternativeNames:        alternativeNames,
		BankIDCode:              "GBDSC",
		BaseCurrency:            "GBP",
		BankID:                  "4003000",
		Bic:                     "NWBKGB22",
		Country:                 &country,
		Name:                    accountNames,
		SecondaryIdentification: "A1B2C3D4",
		JointAccount:            &jointAccount,
	}

	content := Account{
		Type:           "accounts",
		ID:             "bd27e265-9605-4b4b-a0e5-3003ea9cc4dc",
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Attributes:     &accountAttributes,
	}

	dataObj := Data{
		Data: content,
	}

	b, err := json.Marshal(dataObj)
	responseBody := bytes.NewBuffer(b)
	resp, err := http.Post("http://localhost:8080/v1/organisation/accounts", "application/json", responseBody)

	//handle error
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	// read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// convert the body to string
	sb := string(body)
	log.Printf(sb)
}
