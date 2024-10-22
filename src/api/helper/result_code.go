package helper

type ResultCode int

const (
    // Success Codes
    Success         ResultCode = 200   // Request successful

    // Client Errors (4xx)
    ValidationError ResultCode = 400   // Validation error
    AuthError       ResultCode = 401   // Authentication error
    ForbiddenError  ResultCode = 403   // Access forbidden
    NotFoundError   ResultCode = 404   // Resource not found
    LimiterError    ResultCode = 429   // Too many requests (rate limit)
    OtpLimiterError ResultCode = 42901 // Too many OTP requests (rate limit for OTP)

    // Server Errors (5xx)
    CustomRecovery  ResultCode = 500   // Custom recovery error
    InternalError   ResultCode = 500   // Internal server error
)
