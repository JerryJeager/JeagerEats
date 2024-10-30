package http

import "context"

type UserIDPathParam struct {
	UserID string `uri:"id" binding:"required,uuid_rfc4122"`
}

type RestaurantIDPathParam struct {
	UserID string `uri:"id" binding:"required,uuid_rfc4122"`
}

func GetUserID(ctx context.Context) (string, error) {
	return ctx.Value("user_id").(string), nil
}
func GetRole(ctx context.Context) (string, error) {
	return ctx.Value("role").(string), nil
}
