package api

import (
	"errors"
	"fmt"
	"strings"
)

// Common errors
var (
	ErrNotAuthenticated = errors.New("not authenticated - run 'gh auth login' first")
	ErrNotFound         = errors.New("resource not found")
	ErrRateLimited      = errors.New("API rate limit exceeded")
)

// APIError wraps GitHub API errors with additional context
type APIError struct {
	Operation string
	Resource  string
	Err       error
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%s %s: %v", e.Operation, e.Resource, e.Err)
}

func (e *APIError) Unwrap() error {
	return e.Err
}

// IsNotFound checks if an error indicates a resource was not found
func IsNotFound(err error) bool {
	if errors.Is(err, ErrNotFound) {
		return true
	}
	if err == nil {
		return false
	}
	// Check for GraphQL "not found" patterns
	msg := err.Error()
	return strings.Contains(msg, "Could not resolve") ||
		strings.Contains(msg, "NOT_FOUND")
}

// IsRateLimited checks if an error indicates rate limiting
func IsRateLimited(err error) bool {
	if errors.Is(err, ErrRateLimited) {
		return true
	}
	if err == nil {
		return false
	}
	msg := err.Error()
	return strings.Contains(msg, "rate limit") ||
		strings.Contains(msg, "RATE_LIMITED")
}

// IsAuthError checks if an error indicates authentication issues
func IsAuthError(err error) bool {
	if errors.Is(err, ErrNotAuthenticated) {
		return true
	}
	if err == nil {
		return false
	}
	msg := err.Error()
	return strings.Contains(msg, "401") ||
		strings.Contains(msg, "authentication") ||
		strings.Contains(msg, "not authenticated")
}

// WrapError wraps an API error with operation context
func WrapError(operation, resource string, err error) error {
	if err == nil {
		return nil
	}

	// Check for specific error types and wrap accordingly
	if IsRateLimited(err) {
		return &APIError{
			Operation: operation,
			Resource:  resource,
			Err:       ErrRateLimited,
		}
	}

	if IsNotFound(err) {
		return &APIError{
			Operation: operation,
			Resource:  resource,
			Err:       ErrNotFound,
		}
	}

	return &APIError{
		Operation: operation,
		Resource:  resource,
		Err:       err,
	}
}
