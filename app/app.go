package app

import (
	"bytes"
	"fmt"
	"image/png"
	"os"
	"time"

	"github.com/yeqown/go-qrcode"
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
