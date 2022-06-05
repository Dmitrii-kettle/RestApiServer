package RestApiServer

type Users struct {
	Id  int `json:"-"`
	NSP struct {
		Name       string `json:"name"`
		Sirname    string `json:"sirname"`
		Patronymic string `json:"patronymic"`
	}
	Number    int    `json:"number"`
	Email     string `json:"email"`
	Birthdate string `json:"birthdate"`
}
