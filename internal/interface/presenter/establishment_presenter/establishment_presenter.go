package establishment_presenter

import (
	"embed"
	"github.com/skip2/go-qrcode"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"html/template"
	"io"
)

//go:embed templates/*
var templates embed.FS

type establishmentPresenter struct {
	tmpl *template.Template
}

func (e *establishmentPresenter) PresentQrCode(w io.Writer, link string) error {
	png, err := qrcode.Encode(link, qrcode.Medium, 256)
	if err != nil {
		return err
	}

	_, err = w.Write(png)
	return err
}

func (e *establishmentPresenter) PresentEstablishmentInfo(w io.Writer, establishment *models.Establishment, currencies []models.Currency) error {
	return e.tmpl.ExecuteTemplate(w, "establishment_info", map[string]any{
		"E":          establishment,
		"Currencies": currencies,
	})
}

func (e *establishmentPresenter) PresentEditEstablishmentPage(w io.Writer, establishment *models.Establishment, topMenuID int64) error {
	return e.tmpl.ExecuteTemplate(w, "edit_establishment_page", map[string]any{
		"Establishment": establishment,
		"TopMenuID":     topMenuID,
	})
}

func (e *establishmentPresenter) PresentEstablishments(w io.Writer, establishments []models.Establishment, currencies []models.Currency) error {
	return e.tmpl.ExecuteTemplate(w, "establishments_page", map[string]any{
		"Establishments": establishments,
		"Currencies":     currencies,
	})
}

func NewEstablishmentPresenter() presenter.EstablishmentPresenter {
	tmpl, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	return &establishmentPresenter{tmpl: tmpl}
}
