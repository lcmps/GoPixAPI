package app

import (
	"bytes"
	"fmt"
	"image/png"
	"os"
	"time"
	"unicode"

	"github.com/yeqown/go-qrcode"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func SaveImage(fileName string, imgData []byte) error {
	fileName = fmt.Sprintf("./pages/assets/img/qrs/%s.png", fileName)
	img, err := png.Decode(bytes.NewReader(imgData))
	if err != nil {
		return err
	}
	out, err := os.Create(fileName)
	if err != nil {
		return err
	}

	err = png.Encode(out, img)
	if err != nil {
		return err
	}
	return nil
}

func GenerateQR(fgHex, bgHex, pasteCode string) ([]byte, error) {
	var byt bytes.Buffer
	qop := qrcode.Config{
		EncMode: qrcode.EncModeByte,
		EcLevel: qrcode.ErrorCorrectionMedium,
	}

	q, err := qrcode.NewWithConfig(pasteCode,
		&qop,
		qrcode.WithFgColorRGBHex(fgHex),
		qrcode.WithBgColorRGBHex(bgHex),
		qrcode.WithBuiltinImageEncoder(qrcode.PNG_FORMAT),
	)
	if err != nil {
		return nil, err
	}
	err = q.SaveTo(&byt)
	if err != nil {
		return nil, err
	}

	return byt.Bytes(), nil
}

func fileNameGenerator(k string) string {
	t := time.Now()
	today := t.Format("2006-01-02-T15")
	name := fmt.Sprintf(`%s_%s`, today, k)
	return name
}

func normalizeInput(input string) (string, error) {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	res, _, err := transform.String(t, input)
	if err != nil {
		fmt.Println(err.Error())
		return input, err
	}

	fmt.Println("input normalize to " + res)
	return res, nil
}
