package authen

import (
	"context"
	"errors"
	"testing"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

type registerUseCaseStub struct {
	called bool
	input  models.CreateUsersRequest
	err    error
}

func (s *registerUseCaseStub) RegisterUser(_ context.Context, input models.CreateUsersRequest) error {
	s.called = true
	s.input = input
	return s.err
}

type loginUseCaseStub struct {
	called   bool
	email    string
	password string
	token    string
	err      error
}

func (s *loginUseCaseStub) Login(_ context.Context, email string, password string) (string, error) {
	s.called = true
	s.email = email
	s.password = password
	return s.token, s.err
}

type verifyUseCaseStub struct {
	changePasswordCalled     bool
	sendVerificationCalled   bool
	verifyEmailCalled        bool
	verifyOTPViaEmailCalled  bool
	verifyPhoneCalled        bool
	changePasswordErr        error
	sendVerificationEmailErr error
	verifyEmailErr           error
	verifyOTPViaEmailErr     error
	verifyPhoneErr           error
}

func (s *verifyUseCaseStub) ChangePassword(_ context.Context, _ string, _ string) error {
	s.changePasswordCalled = true
	return s.changePasswordErr
}

func (s *verifyUseCaseStub) SendVerificationEmail(_ context.Context, _ string) error {
	s.sendVerificationCalled = true
	return s.sendVerificationEmailErr
}

func (s *verifyUseCaseStub) VerifyEmail(_ context.Context, _, _ string) error {
	s.verifyEmailCalled = true
	return s.verifyEmailErr
}

func (s *verifyUseCaseStub) VerifyOTP_viaEmail(_ context.Context, _, _ string) error {
	s.verifyOTPViaEmailCalled = true
	return s.verifyOTPViaEmailErr
}

func (s *verifyUseCaseStub) VerifyPhone(_ context.Context, _, _ string) error {
	s.verifyPhoneCalled = true
	return s.verifyPhoneErr
}

func TestAuthServiceDelegatesCalls(t *testing.T) {
	ctx := context.Background()
	registerStub := &registerUseCaseStub{}
	loginStub := &loginUseCaseStub{token: "access-token"}
	verifyStub := &verifyUseCaseStub{}

	svc := NewAuthService(registerStub, loginStub, verifyStub)

	if err := svc.Register(ctx, models.CreateUsersRequest{Email: "test@example.com"}); err != nil {
		t.Fatalf("Register() unexpected error: %v", err)
	}
	if !registerStub.called {
		t.Fatalf("Register() did not delegate to RegisterUseCase")
	}

	token, err := svc.Login(ctx, "test@example.com", "password")
	if err != nil {
		t.Fatalf("Login() unexpected error: %v", err)
	}
	if token != "access-token" || !loginStub.called {
		t.Fatalf("Login() did not delegate correctly, token=%q called=%v", token, loginStub.called)
	}

	if err := svc.ChangePassword(ctx, "uid", "new-pass"); err != nil {
		t.Fatalf("ChangePassword() unexpected error: %v", err)
	}
	if err := svc.SendVerificationEmail(ctx, "test@example.com"); err != nil {
		t.Fatalf("SendVerificationEmail() unexpected error: %v", err)
	}
	if err := svc.VerifyEmail(ctx, "test@example.com", "123456"); err != nil {
		t.Fatalf("VerifyEmail() unexpected error: %v", err)
	}
	if err := svc.VerifyOTP_viaEmail(ctx, "test@example.com", "123456"); err != nil {
		t.Fatalf("VerifyOTP_viaEmail() unexpected error: %v", err)
	}
	if err := svc.VerifyPhone(ctx, "0123456789", "123456"); err != nil {
		t.Fatalf("VerifyPhone() unexpected error: %v", err)
	}

	if !verifyStub.changePasswordCalled || !verifyStub.sendVerificationCalled || !verifyStub.verifyEmailCalled || !verifyStub.verifyOTPViaEmailCalled || !verifyStub.verifyPhoneCalled {
		t.Fatalf("verification methods were not fully delegated")
	}
}

func TestAuthServicePropagatesErrors(t *testing.T) {
	expectedErr := errors.New("expected")
	ctx := context.Background()
	registerStub := &registerUseCaseStub{err: expectedErr}
	loginStub := &loginUseCaseStub{err: expectedErr}
	verifyStub := &verifyUseCaseStub{
		changePasswordErr:        expectedErr,
		sendVerificationEmailErr: expectedErr,
		verifyEmailErr:           expectedErr,
		verifyOTPViaEmailErr:     expectedErr,
		verifyPhoneErr:           expectedErr,
	}

	svc := NewAuthService(registerStub, loginStub, verifyStub)

	if err := svc.Register(ctx, models.CreateUsersRequest{}); !errors.Is(err, expectedErr) {
		t.Fatalf("Register() should propagate error")
	}
	if _, err := svc.Login(ctx, "", ""); !errors.Is(err, expectedErr) {
		t.Fatalf("Login() should propagate error")
	}
	if err := svc.ChangePassword(ctx, "", ""); !errors.Is(err, expectedErr) {
		t.Fatalf("ChangePassword() should propagate error")
	}
	if err := svc.SendVerificationEmail(ctx, ""); !errors.Is(err, expectedErr) {
		t.Fatalf("SendVerificationEmail() should propagate error")
	}
	if err := svc.VerifyEmail(ctx, "", ""); !errors.Is(err, expectedErr) {
		t.Fatalf("VerifyEmail() should propagate error")
	}
	if err := svc.VerifyOTP_viaEmail(ctx, "", ""); !errors.Is(err, expectedErr) {
		t.Fatalf("VerifyOTP_viaEmail() should propagate error")
	}
	if err := svc.VerifyPhone(ctx, "", ""); !errors.Is(err, expectedErr) {
		t.Fatalf("VerifyPhone() should propagate error")
	}
}
