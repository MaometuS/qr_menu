package admin_interactor

import (
	"context"
	crypto_rand "crypto/rand"
	"crypto/tls"
	"encoding/binary"
	"fmt"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"io"
	"math/rand"
	"strconv"

	"gopkg.in/gomail.v2"
)

func (a *adminInteractor) VerifyEmailPage(context context.Context, w io.Writer, id int64) error {
	err := sendVerificationCode(context, a.profileRepository, id, "")
	if err != nil {
		return err
	}

	err = a.presenter.VerifyPage(w, id)
	if err != nil {
		return err
	}

	return nil
}

func sendVerificationCode(context context.Context, repo repository.ProfileRepository, id int64, email string) error {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		return err
	}
	rnd := rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(b[:]))))
	verificationCode := strconv.FormatInt(100000+rnd.Int63n(999999-100000), 10)

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

	//auth := smtp.PlainAuth("", "test@qoob.uz", "jGRj&u2fb+pu", "server3.ahost.uz")
	//err = smtp.SendMail("server3.ahost.uz:587", auth, "test@qoob.uz", []string{"maometusu@gmail.com"}, []byte(verificationCode))
	//if err != nil {
	//	return err
	//}

	m := gomail.NewMessage()

	m.SetHeader("From", "test@qoob.uz")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Код для верификации")
	m.SetBody("text/plain", verificationCode)

	d := gomail.NewDialer("server3.ahost.uz", 587, "test@qoob.uz", "jGRj&u2fb+pu")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return nil
}
