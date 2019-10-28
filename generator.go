package codegen

import (
	"bufio"
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/datamatrix"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

const DefaultFileName = "codes.txt"

const DefaultOutputDir = "codes"

const DefaultTail = "910bad92YntF8eliG9/1KohpDEKWt8/Q7Zm1LkcMJaLdWml5IoOf"

const GtinId = "01"

const SerialId = "21"

func LoadFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if e := file.Close(); e != nil {
			log.Fatal(e)
		}
	}()

	codes := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		codes = append(codes, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return codes
}

func ScanGtin(code string) string {
	return code[0:14]
}

func ScanSerial(code string) string {
	return code[14:27]
}

func CreateTail() string {
	return DefaultTail
}

func CreateCode(code string) string {
	gtin := ScanGtin(code)
	serial := ScanSerial(code)
	tail := CreateTail()
	return GtinId + gtin + SerialId + serial + tail
}

func GenerateAndSave(fileName string) {

	codes := LoadFile(fileName)
	for i, code := range codes {

		toEncode := CreateCode(code)

		dmCode, _ := datamatrix.Encode(toEncode)

		dmCode, _ = barcode.Scale(dmCode, 150, 150)

		fileName := fmt.Sprintf("%s%d%s", "dm_", i+1, ".png")

		if _, err := os.Stat(DefaultOutputDir); os.IsNotExist(err) {
			errM := os.Mkdir(DefaultOutputDir, os.ModePerm)
			if errM != nil {
				log.Fatal("Cannot create output dir")
			}
		}

		file, errC := os.Create(filepath.FromSlash(DefaultOutputDir + "/" + fileName))

		if errC != nil {
			log.Fatal(errC)
		}

		defer func() {
			e := file.Close()
			if e != nil {
				log.Fatal(e)
			}
		}()

		err := png.Encode(file, dmCode)

		if err != nil {
			log.Fatal(err)
		}
	}
}
