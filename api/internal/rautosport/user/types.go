package user

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	DNI     string `json:"dni"`
	BDay    string `json:"bday"`
	Vehicle Vehicle
}

type Vehicle struct {
	Model  string
	Plates string
}
