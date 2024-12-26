package jwtService

type ContextKey string

// Ключи, лежащие внутри контекста
const (
	UserIdKey   ContextKey = "userId"   // Идентификатор пользователя
	UserRoleKey ContextKey = "userRole" // Роль пользователя
)
