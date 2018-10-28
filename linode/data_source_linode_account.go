package linode

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/linode/linodego"
)

func dataSourceLinodeAccount() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLinodeAccountRead,
		Schema: map[string]*schema.Schema{
			"email": {
				Type:        schema.TypeString,
				Description: "The email address for this Account, for account management communications, and may be used for other communications as configured.",
				Computed:    true,
			},
			"first_name": {
				Type:        schema.TypeString,
				Description: "The first name of the person associated with this Account.",
				Computed:    true,
			},
			"last_name": {
				Type:        schema.TypeString,
				Description: "The last name of the person associated with this Account.",
				Computed:    true,
			},
			"company": {
				Type:        schema.TypeString,
				Description: "The company name associated with this Account.",
				Computed:    true,
			},
			"address_1": {
				Type:        schema.TypeString,
				Description: "First line of this Account's billing address.",
				Computed:    true,
			},
			"address_2": {
				Type:        schema.TypeString,
				Description: "Second line of this Account's billing address.",
				Computed:    true,
			},
			"phone": {
				Type:        schema.TypeString,
				Description: "The phone number associated with this Account.",
				Computed:    true,
			},
			"city": {
				Type:        schema.TypeString,
				Description: "The city for this Account's billing address.",
				Computed:    true,
			},
			"state": {
				Type:        schema.TypeString,
				Description: "If billing address is in the United States, this is the State portion of the Account's billing address. If the address is outside the US, this is the Province associated with the Account's billing address.",
				Computed:    true,
			},
			"country": {
				Type:        schema.TypeString,
				Description: "The two-letter country code of this Account's billing address.",
				Computed:    true,
			},
			"zip": {
				Type:        schema.TypeString,
				Description: "The zip code of this Account's billing address.",
				Computed:    true,
			},
			"credit_card": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "Credit Card information associated with this Account.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"last_four": {
							Type:        schema.TypeString,
							Description: "The last four digits of the credit card associated with this Account.",
							Computed:    true,
						},
						"expiry": {
							Type:        schema.TypeString,
							Description: "The expiration month and year of the credit card.",
							Computed:    true,
						},
					},
				},
			},
			"tax_id": {
				Type:        schema.TypeString,
				Description: "The tax identification number associated with this Account, for tax calculations in some countries. If the account is not based in a country that collects tax, this should be null.",
				Computed:    true,
			},
			"balance": {
				Type:        schema.TypeInt,
				Description: "This Account's balance, in US dollars.",
				Computed:    true,
			},
		},
	}
}

func dataSourceLinodeAccountRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)

	account, err := client.GetAccount(context.Background())
	if err != nil {
		return fmt.Errorf("Error getting account: %s", err)
	}

	d.SetId(account.Email)
	d.Set("email", account.Email)
	d.Set("first_name", account.FirstName)
	d.Set("last_name", account.LastName)
	d.Set("company", account.Company)

	d.Set("address_1", account.Address1)
	d.Set("address_2", account.Address2)
	d.Set("phone", account.Phone)
	d.Set("city", account.City)
	d.Set("state", account.State)
	d.Set("country", account.Country)
	d.Set("zip", account.Zip)

	if account.CreditCard == nil {
		d.Set("credit_card", nil)
	} else {
		if err := d.Set("credit_card", flattenAccountCreditCard(*account.CreditCard)); err != nil {
			return fmt.Errorf("Error parsing account credit card: %s", err)
		}
	}

	d.Set("tax_id", account.TaxID)
	d.Set("balance", account.Balance)

	return nil
}