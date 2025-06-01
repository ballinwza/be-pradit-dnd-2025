package graph

import user_service "github.com/ballinwza/be-pradit-dnd-2025/internal/outbound/user/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	UserService *user_service.UserService
}


