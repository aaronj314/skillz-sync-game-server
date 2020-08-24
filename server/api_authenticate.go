// //
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"context"
	"errors"
	"math/rand"
	"regexp"
	"time"

	"github.com/aaron-skillz/sync-server-go/api"
	"github.com/dgrijalva/jwt-go"
)

var (
	invalidCharsRegex = regexp.MustCompilePOSIX("([[:cntrl:]]|[[:space:]])+")
	emailRegex        = regexp.MustCompile("^.+@.+\\..+$")
)

type SessionTokenClaims struct {
	UserId    string            `json:"uid,omitempty"`
	Username  string            `json:"usn,omitempty"`
	Vars      map[string]string `json:"vrs,omitempty"`
	ExpiresAt int64             `json:"exp,omitempty"`
}

func (stc *SessionTokenClaims) Valid() error {
	// Verify expiry.
	if stc.ExpiresAt <= time.Now().UTC().Unix() {
		vErr := new(jwt.ValidationError)
		vErr.Inner = errors.New("Token is expired")
		vErr.Errors |= jwt.ValidationErrorExpired
		return vErr
	}
	return nil
}

func (s *ApiServer) AuthenticateCustom(ctx context.Context, in *api.AuthenticateCustomRequest) (*api.Session, error) {
	//// Before hook.
	//if fn := s.runtime.BeforeAuthenticateCustom(); fn != nil {
	//	beforeFn := func(clientIP, clientPort string) error {
	//		result, err, code := fn(ctx, s.logger, "", "", nil, 0, clientIP, clientPort, in)
	//		if err != nil {
	//			return status.Error(code, err.Error())
	//		}
	//		if result == nil {
	//			// If result is nil, requested resource is disabled.
	//			s.logger.Warn("Intercepted a disabled resource.", zap.Any("resource", ctx.Value(ctxFullMethodKey{}).(string)))
	//			return status.Error(codes.NotFound, "Requested resource was not found.")
	//		}
	//		in = result
	//		return nil
	//	}
	//
	//	// Execute the before function lambda wrapped in a trace for stats measurement.
	//	err := traceApiBefore(ctx, s.logger, s.metrics, ctx.Value(ctxFullMethodKey{}).(string), beforeFn)
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	//
	//if in.Account == nil || in.Account.Id == "" {
	//	return nil, status.Error(codes.InvalidArgument, "Custom ID is required.")
	//} else if invalidCharsRegex.MatchString(in.Account.Id) {
	//	return nil, status.Error(codes.InvalidArgument, "Custom ID invalid, no spaces or control characters allowed.")
	//} else if len(in.Account.Id) < 6 || len(in.Account.Id) > 128 {
	//	return nil, status.Error(codes.InvalidArgument, "Custom ID invalid, must be 6-128 bytes.")
	//}
	//
	//username := in.Username
	//if username == "" {
	//	username = generateUsername()
	//} else if invalidCharsRegex.MatchString(username) {
	//	return nil, status.Error(codes.InvalidArgument, "Username invalid, no spaces or control characters allowed.")
	//} else if len(username) > 128 {
	//	return nil, status.Error(codes.InvalidArgument, "Username invalid, must be 1-128 bytes.")
	//}
	//
	//create := in.Create == nil || in.Create.Value
	//
	//dbUserID, dbUsername, created, err := AuthenticateCustom(ctx, s.logger, s.db, in.Account.Id, username, create)
	//if err != nil {
	//	return nil, err
	//}
	//
	//token, exp := generateToken(s.config, dbUserID, dbUsername, in.Account.Vars)
	//session := &api.Session{Created: created, Token: token}
	//
	//// After hook.
	//if fn := s.runtime.AfterAuthenticateCustom(); fn != nil {
	//	afterFn := func(clientIP, clientPort string) error {
	//		return fn(ctx, s.logger, dbUserID, dbUsername, in.Account.Vars, exp, clientIP, clientPort, session, in)
	//	}
	//
	//	// Execute the after function lambda wrapped in a trace for stats measurement.
	//	traceApiAfter(ctx, s.logger, s.metrics, ctx.Value(ctxFullMethodKey{}).(string), afterFn)
	//}
	//
	//return session, nil
	return nil, errors.New("Not supported")
}

func (s *ApiServer) AuthenticateDevice(ctx context.Context, in *api.AuthenticateDeviceRequest) (*api.Session, error) {
	// Before hook.
	//if fn := s.runtime.BeforeAuthenticateDevice(); fn != nil {
	//	beforeFn := func(clientIP, clientPort string) error {
	//		result, err, code := fn(ctx, s.logger, "", "", nil, 0, clientIP, clientPort, in)
	//		if err != nil {
	//			return status.Error(code, err.Error())
	//		}
	//		if result == nil {
	//			// If result is nil, requested resource is disabled.
	//			s.logger.Warn("Intercepted a disabled resource.", zap.Any("resource", ctx.Value(ctxFullMethodKey{}).(string)))
	//			return status.Error(codes.NotFound, "Requested resource was not found.")
	//		}
	//		in = result
	//		return nil
	//	}
	//
	//	// Execute the before function lambda wrapped in a trace for stats measurement.
	//	err := traceApiBefore(ctx, s.logger, s.metrics, ctx.Value(ctxFullMethodKey{}).(string), beforeFn)
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	//
	//if in.Account == nil || in.Account.Id == "" {
	//	return nil, status.Error(codes.InvalidArgument, "Device ID is required.")
	//} else if invalidCharsRegex.MatchString(in.Account.Id) {
	//	return nil, status.Error(codes.InvalidArgument, "Device ID invalid, no spaces or control characters allowed.")
	//} else if len(in.Account.Id) < 10 || len(in.Account.Id) > 128 {
	//	return nil, status.Error(codes.InvalidArgument, "Device ID invalid, must be 10-128 bytes.")
	//}
	//
	//username := in.Username
	//if username == "" {
	//	username = generateUsername()
	//} else if invalidCharsRegex.MatchString(username) {
	//	return nil, status.Error(codes.InvalidArgument, "Username invalid, no spaces or control characters allowed.")
	//} else if len(username) > 128 {
	//	return nil, status.Error(codes.InvalidArgument, "Username invalid, must be 1-128 bytes.")
	//}
	//
	//create := in.Create == nil || in.Create.Value
	//
	//dbUserID, dbUsername, created, err := AuthenticateDevice(ctx, s.logger, s.db, in.Account.Id, username, create)
	//if err != nil {
	//	return nil, err
	//}
	//
	//token, exp := generateToken(s.config, dbUserID, dbUsername, in.Account.Vars)
	//session := &api.Session{Created: created, Token: token}
	//
	//// After hook.
	//if fn := s.runtime.AfterAuthenticateDevice(); fn != nil {
	//	afterFn := func(clientIP, clientPort string) error {
	//		return fn(ctx, s.logger, dbUserID, dbUsername, in.Account.Vars, exp, clientIP, clientPort, session, in)
	//	}
	//
	//	// Execute the after function lambda wrapped in a trace for stats measurement.
	//	traceApiAfter(ctx, s.logger, s.metrics, ctx.Value(ctxFullMethodKey{}).(string), afterFn)
	//}
	//
	//return session, nil
	return nil, errors.New("Not supported")
}

func (s *ApiServer) AuthenticateEmail(ctx context.Context, in *api.AuthenticateEmailRequest) (*api.Session, error) {
	//// Before hook.
	//if fn := s.runtime.BeforeAuthenticateEmail(); fn != nil {
	//	beforeFn := func(clientIP, clientPort string) error {
	//		result, err, code := fn(ctx, s.logger, "", "", nil, 0, clientIP, clientPort, in)
	//		if err != nil {
	//			return status.Error(code, err.Error())
	//		}
	//		if result == nil {
	//			// If result is nil, requested resource is disabled.
	//			s.logger.Warn("Intercepted a disabled resource.", zap.Any("resource", ctx.Value(ctxFullMethodKey{}).(string)))
	//			return status.Error(codes.NotFound, "Requested resource was not found.")
	//		}
	//		in = result
	//		return nil
	//	}
	//
	//	// Execute the before function lambda wrapped in a trace for stats measurement.
	//	err := traceApiBefore(ctx, s.logger, s.metrics, ctx.Value(ctxFullMethodKey{}).(string), beforeFn)
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	//
	//email := in.Account
	//if email == nil {
	//	return nil, status.Error(codes.InvalidArgument, "Email address and password is required.")
	//}
	//
	//var attemptUsernameLogin bool
	//if email.Email == "" {
	//	// Password was supplied, but no email. Perhaps the user is attempting to login with username/password.
	//	attemptUsernameLogin = true
	//} else if invalidCharsRegex.MatchString(email.Email) {
	//	return nil, status.Error(codes.InvalidArgument, "Invalid email address, no spaces or control characters allowed.")
	//} else if !emailRegex.MatchString(email.Email) {
	//	return nil, status.Error(codes.InvalidArgument, "Invalid email address format.")
	//} else if len(email.Email) < 10 || len(email.Email) > 255 {
	//	return nil, status.Error(codes.InvalidArgument, "Invalid email address, must be 10-255 bytes.")
	//}
	//
	//if len(email.Password) < 8 {
	//	return nil, status.Error(codes.InvalidArgument, "Password must be at least 8 characters long.")
	//}
	//
	//username := in.Username
	//if username == "" {
	//	// If no username was supplied and the email was missing.
	//	if attemptUsernameLogin {
	//		return nil, status.Error(codes.InvalidArgument, "Username is required when email address is not supplied.")
	//	}
	//
	//	// Email address was supplied, we are allowed to generate a username.
	//	username = generateUsername()
	//} else if invalidCharsRegex.MatchString(username) {
	//	return nil, status.Error(codes.InvalidArgument, "Username invalid, no spaces or control characters allowed.")
	//} else if len(username) > 128 {
	//	return nil, status.Error(codes.InvalidArgument, "Username invalid, must be 1-128 bytes.")
	//}
	//
	//var dbUserID string
	//var created bool
	//var err error
	//
	//if attemptUsernameLogin {
	//	// Attempting to log in with username/password. Create flag is ignored, creation is not possible here.
	//	dbUserID, err = AuthenticateUsername(ctx, s.logger, s.db, username, email.Password)
	//} else {
	//	// Attempting email authentication, may or may not create.
	//	cleanEmail := strings.ToLower(email.Email)
	//	create := in.Create == nil || in.Create.Value
	//
	//	dbUserID, username, created, err = AuthenticateEmail(ctx, s.logger, s.db, cleanEmail, email.Password, username, create)
	//}
	//if err != nil {
	//	return nil, err
	//}
	//
	//token, exp := generateToken(s.config, dbUserID, username, in.Account.Vars)
	//session := &api.Session{Created: created, Token: token}
	//
	//// After hook.
	//if fn := s.runtime.AfterAuthenticateEmail(); fn != nil {
	//	afterFn := func(clientIP, clientPort string) error {
	//		return fn(ctx, s.logger, dbUserID, username, in.Account.Vars, exp, clientIP, clientPort, session, in)
	//	}
	//
	//	// Execute the after function lambda wrapped in a trace for stats measurement.
	//	traceApiAfter(ctx, s.logger, s.metrics, ctx.Value(ctxFullMethodKey{}).(string), afterFn)
	//}
	//
	//return session, nil
	return nil, errors.New("Not supported")
}

func generateToken(config Config, userID, username string, vars map[string]string) (string, int64) {
	exp := time.Now().UTC().Add(time.Duration(config.GetSession().TokenExpirySec) * time.Second).Unix()
	return generateTokenWithExpiry(config, userID, username, vars, exp)
}

func generateTokenWithExpiry(config Config, userID, username string, vars map[string]string, exp int64) (string, int64) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &SessionTokenClaims{
		UserId:    userID,
		Username:  username,
		Vars:      vars,
		ExpiresAt: exp,
	})
	signedToken, _ := token.SignedString([]byte(config.GetSession().EncryptionKey))
	return signedToken, exp
}

func generateUsername() string {
	const usernameAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 10)
	for i := range b {
		b[i] = usernameAlphabet[rand.Intn(len(usernameAlphabet))]
	}
	return string(b)
}
