package codegen

import (
	"github.com/vitali-fedulov/images"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadFile(t *testing.T) {
	res := LoadFile(DefaultFileName)
	if len(res) != 4 {
		t.Error("Codes number not equal 4 in file codes.txt")
	}

	code := res[1]
	if code != "046070283942871910240606501" {
		t.Error("2 Code not equal expected")
	}
}

func TestScanGtin(t *testing.T) {
	codes := LoadFile(DefaultFileName)
	code := codes[0]
	gtin := ScanGtin(code)
	if gtin != "04607028394287" {
		t.Error("Gtin for 1 code not expected")
	}
}

func TestScanSerial(t *testing.T) {
	codes := LoadFile(DefaultFileName)
	code := codes[0]
	serial := ScanSerial(code)
	if serial != "1910240606480" {
		t.Error("Serial for 1 code not expected")
	}
}

func TestCreateTail(t *testing.T) {
	tail := CreateTail()
	if tail != DefaultTail {
		t.Error("Tail not equal default tail value")
	}
}

func TestCreateCode(t *testing.T) {
	codes := LoadFile(DefaultFileName)
	code := codes[0]
	data := CreateCode(code)
	if data != "0104607028394287211910240606480910bad92YntF8eliG9/1KohpDEKWt8/Q7Zm1LkcMJaLdWml5IoOf" {
		t.Error("Code value created from 1 code not expected")
	}

}

func TestGeneratedImage(t *testing.T) {
	GenerateAndSave(DefaultFileName)
	genFileNames := make([]string, 0)
	genFileNames = append(genFileNames,
		"dm_1.png",
		"dm_2.png",
		"dm_3.png",
		"dm_4.png")
	for _, name := range genFileNames {

		filePath := filepath.FromSlash(DefaultOutputDir + "/" + name)

		imgA, err := images.Open(filePath)
		if err != nil {
			panic(err)
		}
		imgB, err := images.Open("test/test_" + name)
		if err != nil {
			panic(err)
		}

		masks := images.Masks()

		hA, imgSizeA := images.Hash(imgA, masks)
		hB, imgSizeB := images.Hash(imgB, masks)

		if !images.Similar(hA, hB, imgSizeA, imgSizeB) {
			t.Error("Generated png not expected")
		}
		errR := os.Remove(filePath)
		if errR != nil {
			log.Fatal("Error remove generated file")
		}
	}
}
