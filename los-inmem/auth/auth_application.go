package auth

type AuthApp struct {
	repository *Repository
}

func NewApp(repository *Repository) *AuthApp {
	return &AuthApp{
		repository: repository,
	}
}
