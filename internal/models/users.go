package models

//Users ...
type Users struct {
	Userid     int
	Nombre     string
	Apellido   string
	Correo     string
	Estado     int
	Contrasena string
}

//UsersTransform ...
type UsersTransform struct {
	UserID     int    `json:"userid"`
	Nombre     string `json:"nombre"`
	Apellido   string `json:"apellido"`
	Correo     string `json:"correo"`
	Estado     int    `json:"estado"`
	Contrasena string `json:"contrasena"`
}
