package transformers

import (
	"time"

	"github.com/chrislentz/example-go-restful-api/sqlc"
)

// Transform item struct
type UserTransformer struct {
	sqlc.User
}

// Transform collection struct
type UsersTransformer struct {
	Users []sqlc.User
}

type UserResponse struct {
	Uuid      string    `json:"uuid"`
	Name      string    `json:"name"`
	Github    string    `json:"github"`
	Twitter   string    `json:"twitter"`
	Mastodon  string    `json:"mastodon"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Transform item method
func (instance *UserTransformer) Transform() UserResponse {

	output := UserResponse{
		Uuid:      instance.Uuid,
		Name:      instance.Name,
		Github:    instance.Github.String,
		Twitter:   instance.Twitter.String,
		Mastodon:  instance.Mastodon.String,
		CreatedAt: instance.CreatedAt,
		UpdatedAt: instance.UpdatedAt,
	}

	return output
}

// Transform collection method
func (instance *UsersTransformer) Transform() []UserResponse {
	output := []UserResponse{}

	for _, user := range instance.Users {
		transformer := UserTransformer{user}
		output = append(output, transformer.Transform())
	}

	return output
}
