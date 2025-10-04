package model

import "time"

type User struct {
	ID        string    `json:"id"`
	Auth0ID   string    `json:"auth0Id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserProfile struct {
	ID          string    `json:"id"`
	UserID      string    `json:"userId"`
	DisplayName string    `json:"displayName"`
	Bio         string    `json:"bio"`
	AvatarURL   string    `json:"avatarUrl"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type UserProfileUpdate struct {
	DisplayName string `json:"displayName"`
	Bio         string `json:"bio"`
	AvatarURL   string `json:"avatarUrl"`
}

type UserData struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserDataCreate struct {
	Content string `json:"content"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code,omitempty"`
}
