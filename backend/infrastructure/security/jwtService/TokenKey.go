package jwtService

// TODO переписать на ENUM
// Ключи, лежащие внутри токена
const (
	RoleTokenKey       string = "spec"  // Роль пользователя
	ScopeTokenKey      string = "scope" // Доступные области/разрешения пользователя
	userIdTokenKey     string = "id"    // Идентификатор пользователя
	expirationTokenKey string = "exp"   // Время истечения срока действия токена
)
