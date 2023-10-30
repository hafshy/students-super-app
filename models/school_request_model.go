package models

type SchoolRequest struct {
    Name     string `json:"name"`
    Address  string `json:"address"`
    City     string `json:"city"`
    Country  string `json:"country"`
}
