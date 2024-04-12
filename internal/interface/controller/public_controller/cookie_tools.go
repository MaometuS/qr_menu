package public_controller

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"net/http"
	"time"
)

func getSelectedItems(r *http.Request) (map[int64][]int64, error) {
	selectedItems := make(map[int64][]int64)
	cook, _ := r.Cookie("selected_items")

	if cook == nil {
		return selectedItems, nil
	}

	mp := make(map[int64][]int64)
	if len(cook.Value) == 0 {
		return mp, nil
	}

	destination, err := base64.URLEncoding.DecodeString(cook.Value)
	if err != nil {
		return nil, err
	}

	err = gob.NewDecoder(bytes.NewBuffer(destination)).Decode(&mp)
	if err != nil {
		return nil, err
	}

	return mp, nil
}

func setSelectedItems(source map[int64][]int64) (*http.Cookie, error) {
	buf := bytes.NewBuffer([]byte{})
	err := gob.NewEncoder(buf).Encode(source)
	if err != nil {
		return nil, err
	}

	res := base64.URLEncoding.EncodeToString(buf.Bytes())

	cookie := &http.Cookie{
		Name:     "selected_items",
		Value:    res,
		Path:     "/",
		Expires:  time.Now().AddDate(1, 0, 0),
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	return cookie, nil
}
