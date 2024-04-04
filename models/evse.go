package models

type Evse struct {
	Uid    string `json:"uid"`
	Status int    `json:"status"`
}

type EvseResponse struct {
	Uid    string `json:"uid"`
	Status string `json:"status"`
}
