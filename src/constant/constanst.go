package constant

const (
	// User
	AdminRoleName      string = "admin"
	DefaultRoleName    string = "default"
	DefaultUserName    string = "admin"
	RedisOtpDefaultKey string = "otp"
	MarketDefaultKey 	string = "market"

	// Claims
	AuthorizationHeaderKey string = "Authorization"
	UserIdKey              string = "UserId"
	FirstNameKey           string = "FirstName"
	LastNameKey            string = "LastName"
	UsernameKey            string = "Username"
	EmailKey               string = "Email"
	MobileNumberKey        string = "MobileNumber"
	RolesKey               string = "Roles"
	ExpireTimeKey          string = "Exp"

	// Order
	LimitOrder  = "limit"
	MarketOrder = "market"
	BuyAction   = "buy"
	SellAction  = "sell"
)