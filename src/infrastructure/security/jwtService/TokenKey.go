package jwtService

// Ключи, лежащие внутри токена
const (
	RoleTokenKey       string = "role"  // Роль пользователя
	ScopeTokenKey      string = "scope" // Доступные области/разрешения пользователя
	userIdTokenKey     string = "id"    // Идентификатор пользователя
	expirationTokenKey string = "exp"   // Время истечения срока действия токена
)
