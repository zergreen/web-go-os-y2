package info

type userRequest struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	District  string `json:"district"`
}
type userList struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type checkuser struct {
	Id int `json:"id"`
}

type geographyCode struct {
	GeographyId int `json:"geography_id"`
	Population int    `json:"population"`
}

type statusDetail struct {
	Status string `json:"status"`
}

type address struct {
	NameTh     string `json:"name_th"`
	NameGeo    string `json:"name_geo"`
	Population int    `json:"population"`
}

type searchInfo struct {
	Id         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	District   string `json:"district"`
	Province   string `json:"province"`
	Region     string `json:"region"`
	Population int    `json:"population"`
}
type showAllGeography struct {
	Id         int    `json:"id"`
	NameGeo    string `json:"name_geo"`
	Population int    `json:"population"`
}
