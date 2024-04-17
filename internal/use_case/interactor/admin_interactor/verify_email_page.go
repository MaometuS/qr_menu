package admin_interactor

import (
	"context"
	crypto_rand "crypto/rand"
	"encoding/binary"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"io"
	"math/rand"
	"strconv"
)

func (a *adminInteractor) VerifyEmailPage(context context.Context, w io.Writer, id int64) error {
	err := sendVerificationCode(context, a.profileRepository, a.emailRepository, id, "")
	if err != nil {
		return err
	}

	err = a.presenter.VerifyPage(w, id)
	if err != nil {
		return err
	}

	return nil
}

func generateCode() (string, error) {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		return "", err
	}
	rnd := rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(b[:]))))
	verificationCode := strconv.FormatInt(100000+rnd.Int63n(999999-100000), 10)

	return verificationCode, nil
}

func sendVerificationCode(
	context context.Context,
	repo repository.ProfileRepository,
	mailRepo repository.EmailRepository,
	id int64,
	email string,
) error {
	verificationCode, err := generateCode()
	if err != nil {
		return err
	}

	err = repo.SetVerificationCode(context, id, verificationCode)
	if err != nil {
		return err
	}

	profile, err := repo.GetOne(context, id)
	if err != nil {
		return err
	}

	if email == "" {
		email = profile.Email
	}

	err = mailRepo.SendMail(email, "Код для верификации", verificationCode)
	if err != nil {
		return err
	}

	return nil
}
