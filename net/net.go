package net

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func Download(url, dest string) error {
	if _, err := os.Stat(dest); err == nil {
		return nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.Status != "200" {
		msg, _ := ioutil.ReadAll(resp.Body)
		return errors.Errorf("unable to download: %s", string(msg))
	}

	if err := os.MkdirAll(filepath.Dir(dest), 0750); err != nil {
		return err
	}

	out, err := os.Create(filepath.Join(dest))
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err = io.Copy(out, resp.Body); err != nil {
		return err
	}
	return nil
}
