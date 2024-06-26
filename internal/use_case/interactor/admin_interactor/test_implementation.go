package admin_interactor

import (
	"context"
	"io"

	"github.com/stretchr/testify/mock"
)

type TestImplementation struct {
	mock.Mock
}

func (t *TestImplementation) RestorePasswordPage(ctx context.Context, w io.Writer) error {
	args := t.Called(ctx, w)
	return args.Error(0)
}

func (t *TestImplementation) RestorePassword(ctx context.Context, w io.Writer, email, pass, passRepeat string) (int64, error) {
	args := t.Called(ctx, w, email, pass, passRepeat)
	return int64(args.Int(0)), args.Error(1)
}

func (t *TestImplementation) VerifyRestorePasswordPage(ctx context.Context, w io.Writer, id int64) error {
	args := t.Called(ctx, w, id)
	return args.Error(0)
}

func (t *TestImplementation) VerifyRestorePassword(ctx context.Context, w io.Writer, id int64, verificationCode string) error {
	args := t.Called(ctx, w, id, verificationCode)
	return args.Error(0)
}

func (t *TestImplementation) Auth(context context.Context, tokenString string) (int64, error) {
	args := t.Called(context, tokenString)
	return int64(args.Int(0)), args.Error(1)
}

func (t *TestImplementation) AdminPage(context context.Context, w io.Writer, id int64) error {
	args := t.Called(context, w, id)
	return args.Error(0)
}

func (t *TestImplementation) ProfilePage(context context.Context, w io.Writer, id int64) error {
	args := t.Called(context, w, id)
	return args.Error(0)
}

func (t *TestImplementation) LoginPage(context context.Context, w io.Writer) error {
	args := t.Called(context, w)
	return args.Error(0)
}

func (t *TestImplementation) HandleLogin(context context.Context, w io.Writer, email, password string) (string, error) {
	args := t.Called(context, w, email, password)
	return args.String(0), args.Error(1)
}

func (t *TestImplementation) RegisterPage(context context.Context, w io.Writer) error {
	args := t.Called(context, w)
	return args.Error(0)
}

func (t *TestImplementation) HandleRegister(context context.Context, w io.Writer, name, email, password, passRepeat string) (int64, error) {
	args := t.Called(context, w, name, email, password, passRepeat)
	return int64(args.Int(0)), args.Error(1)
}

func (t *TestImplementation) VerifyEmailPage(context context.Context, w io.Writer, id int64) error {
	args := t.Called(context, w, id)
	return args.Error(0)
}

func (t *TestImplementation) HandleVerifyEmail(context context.Context, w io.Writer, id int64, verificationCode string) (string, error) {
	args := t.Called(context, w, id, verificationCode)
	return args.String(0), args.Error(1)
}

func (t *TestImplementation) ChangePassword(ctx context.Context, password, passRepeat string, id int64) error {
	args := t.Called(ctx, password, passRepeat, id)
	return args.Error(0)
}

func (t *TestImplementation) EditName(ctx context.Context, email, name string, id int64) (bool, error) {
	args := t.Called(ctx, email, name, id)
	return args.Bool(0), args.Error(1)
}

func (t *TestImplementation) VerifyEmailEditPage(ctx context.Context, w io.Writer, id int64, email string) error {
	args := t.Called(ctx, w, id, email)
	return args.Error(0)
}

func (t *TestImplementation) EditEmail(ctx context.Context, email, verificationCode string, id int64) error {
	args := t.Called(ctx, email, verificationCode, id)
	return args.Error(0)
}

func NewTestImplementation() *TestImplementation {
	return &TestImplementation{}
}
