package structures

type Admin struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	UserName  string `json:"UserName"`
	Email     string `json:"Email"`
	Contact   string `json:"Contact"`
	Contact2  string `json:"Contact2"`
	Password  string `json:"Password"`
}

type Owner struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Contact   string `json:"Contact"`
	Address   string `json:"Address"`
	Password  string `json:"Password"`
}

type Seeker struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Contact   string `json:"Contact"`
	Password  string `json:"Password"`
}

type Tenant struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Contact   string `json:"Contact"`
	HouseRef  string `json:"HouseRef"`
	Password  string `json:"Password"`
}
