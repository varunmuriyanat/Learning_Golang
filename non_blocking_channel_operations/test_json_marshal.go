package main

import (
    "encoding/json"
    "fmt"
    "os"
)

func main() {
    type AccountAttributes struct {
        Country                   string            `json:"country"`
        BaseCurrency              string            `json:"base_currency"`
        BankId                    string            `json:"bank_id"`
        BankIdCode                string            `json:"bank_id_code"`
        Bic                       string            `json:"bic"`
        Name                      []string          `json:"name"`
        AlternativeNames          []string          `json:"alternative_name"`
        AccountClassification     string            `json:"account_classification"`
        JointAccount              bool              `json:"joint_account"`
        AccountMatchingOptOut     bool              `json:"account_matching_opt_out"`
        SecondaryIdentification   string            `json:"secondary_identification"`
    }

    type Account struct {
        Type1                     string            `json:"type"`
        Id                        string            `json:"id"`
        OrganisationId            string            `json:"organization_id"`
        Attributes                AccountAttributes `json:"attributes"`
    }

    type Data struct {
        Data                      Account           `json:"data"`
    }

    accountNames      := []string{"Varun Muriyanat"}
    alternativeNames  := []string{"VM"}

    accountAttributes := AccountAttributes {
        Country:                  "GB",
        BaseCurrency:             "GBP",
        BankId:                   "4003000",
        BankIdCode:               "GBDSC",
        Bic:                      "NWBKGB22",
        Name:                     accountNames,
        AlternativeNames:         alternativeNames,
        AccountClassification:    "Personal",
        JointAccount:             false,
        AccountMatchingOptOut:    false,
        SecondaryIdentification:  "A1B2C3D4",
    }

    content := Account {
        Type1:                    "accounts",
        Id:                       "bd27e265-9605-4b4b-a0e5-3003ea9cc4dc",
        OrganisationId:           "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
        Attributes:               accountAttributes,
    }

    dataObj := Data {
        Data:                     content,
    }

    b, err := json.Marshal(dataObj)
    if err != nil {
        fmt.Println("error:", err)
    }
    os.Stdout.Write(b)
}

