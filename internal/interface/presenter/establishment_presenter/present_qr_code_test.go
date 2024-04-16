package establishment_presenter

import (
	"github.com/skip2/go-qrcode"
	"os"
	"reflect"
	"testing"
)

func TestEstablishmentPresenter_PresentQrCode(t *testing.T) {
	golden, err := os.ReadFile("templates/qr_code_golden.png")
	if err != nil {
		t.Error(err)
	}

	png, err := qrcode.Encode("link", qrcode.Medium, 256)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(golden, png) {
		t.Error("results don't match")
	}
}

func generateGoldenQrCode() {
	f, err := os.OpenFile("templates/qr_code_golden.png", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	png, err := qrcode.Encode("link", qrcode.Medium, 256)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(png)
	if err != nil {
		panic(err)
	}
}
