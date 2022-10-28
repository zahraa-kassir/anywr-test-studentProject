package dto

type SimpleData struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// -----------------------------------------------------------------
// re-construct with dto

func ReconstructSimpleData(count int, name string, email string) SimpleData {
	dt := SimpleData{
		Id:    count,
		Name:  name,
		Email: email,
	}
	return dt
}

// -----------------------------------------------------------------
