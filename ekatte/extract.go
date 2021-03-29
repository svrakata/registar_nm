package ekatte

import (
	"os"
	"path/filepath"
)

const ekatteURL = "http://www.nsi.bg/sites/default/files/files/EKATTE/Ekatte.zip"
const zipName = "ekatte.zip"

func getPaths() (string, error) {
	dirName := "ekatte_data"
	here, err := filepath.Abs("./")

	if err != nil {
		return "", err
	}

	dirPath := filepath.Join(here, dirName)

	_, err = os.Stat(dirPath)

	if err != nil {
		if os.IsNotExist(err) {
			// the directory doesnt exist
			// create one

			err = os.MkdirAll(dirPath, os.ModePerm)

			if err != nil {
				return "", err
			}

		} else {
			// some other problem with the directory (permisions or)
			return "", err
		}
	}

	return dirPath, nil
}

// DownloadEkatte download
func DownloadEkatte() error {
	dirPath, err := getPaths()

	if err != nil {
		return err
	}

	// the dir is created; continue with downloading the file
	err = DownloadFile(filepath.Join(dirPath, zipName), ekatteURL)

	if err != nil {
		return err
	}

	return nil
}

// UnzipEkatte extracts the ekatte.zip file in
func UnzipEkatte(srcDir string, dest string) error {
	ekatteFileName := "Ekatte_xlsx.zip"

	err := Unzip(filepath.Join(srcDir, zipName), dest)

	if err != nil {
		return err
	}

	err = Unzip(filepath.Join(srcDir, ekatteFileName), dest)

	if err != nil {
		return err
	}

	return nil
}

// ExtractData downloads and unzips the ekatte data
func ExtractData() error {
	dirPath, err := getPaths()

	if err != nil {
		return err
	}

	err = DownloadEkatte()

	if err != nil {
		return err
	}

	err = UnzipEkatte(dirPath, dirPath)

	if err != nil {
		return err
	}

	return nil
}
