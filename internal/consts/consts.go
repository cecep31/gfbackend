package consts

// Context key constants for storing values in context
const (
	// CtxUserID is the key for storing user ID in context
	CtxUserID = "ctx_user_id"
	// CtxUsername is the key for storing username in context
	CtxUsername = "ctx_username"
)

// JWT related constants
const (
	// JWTTokenType is the token type for JWT
	JWTTokenType = "Bearer"
	// JWTContextKey is the key for JWT token in context
	JWTContextKey = "jwt_token"
)
