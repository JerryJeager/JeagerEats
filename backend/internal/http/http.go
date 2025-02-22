package http

import "context"

type UserIDPathParam struct {
	UserID string `uri:"id" binding:"required,uuid_rfc4122"`
}

type RestaurantIDPathParam struct {
	ID string `uri:"id" binding:"required,uuid_rfc4122"`
}
type OrderIDPathParam struct {
	ID string `uri:"id" binding:"required,uuid_rfc4122"`
}
type MenuIDPathParam struct {
	ID string `uri:"id" binding:"required,uuid_rfc4122"`
}

func GetUserID(ctx context.Context) (string, error) {
	return ctx.Value("user_id").(string), nil
}
func GetRole(ctx context.Context) (string, error) {
	return ctx.Value("role").(string), nil
}
func GetRestaurantID(ctx context.Context) (string, error) {
	return ctx.Value("restaurant_id").(string), nil
}

type ErrorJson struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func GetErrorJson(err error, message string) *ErrorJson{
	return &ErrorJson{
		Message: message,
		Error: err.Error(),
	}
}