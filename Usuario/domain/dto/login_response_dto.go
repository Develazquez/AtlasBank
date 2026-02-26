package dto

import (
	"database/sql/driver"
	"encoding/json"
)

type CardDTO struct {
	NameCard   string `json:"name_card"`   
	NumCard    string `json:"num_card"`   
	Expires    string `json:"expires"`     
	CardHolder string `json:"card_holder"` 
}

type RecentlyInfDTO struct {
	Concept     string  `json:"concept"`      
	TypeInf     string  `json:"type_inf"`     
	DayTransfer string  `json:"day_transfer"` 
	Mount       float64 `json:"mount"`        
}


type UserDashboardDTO struct {
	ID          string          `json:"id"`           
	Name        string          `json:"name"`         
	Wallet      float64         `json:"wallet"`       
	Card        CardDTO         `json:"card"`         
	RecentlyInf *RecentlyInfDTO `json:"recently_inf"` 
}

func (c *CardDTO) Scan(value interface{}) error {
	bytes, _ := value.([]byte)
	return json.Unmarshal(bytes, &c)
}

func (c CardDTO) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (r *RecentlyInfDTO) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, _ := value.([]byte)
	return json.Unmarshal(bytes, &r)
}

func (r RecentlyInfDTO) Value() (driver.Value, error) {
	return json.Marshal(r)
}
