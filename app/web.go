package app

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lcmps/gopix/models"
)

func Host() {
	r := gin.Default()
	r.LoadHTMLGlob("./pages/html/*")

	// Assets
	r.Static("/assets", "./pages/assets/")
	r.Static("/fvc/", "./pages/assets/img/favicon/")
	r.Static("/stored", "./pages/assets/img/qrs/")

	// Pages
	r.GET("/", home)
	r.GET("/p", pastaGenerator)

	// API's
	r.POST("/qr", serveQRCode())
	r.POST("/paste", servePaste())
	r.POST("/link", serveQRLink())

	gin.SetMode(gin.ReleaseMode)
	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func home(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "GoPix",
		},
	)
}

func pastaGenerator(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"pasta-index.html",
		gin.H{
			"title": "GoPix | Copy Pasta",
		},
	)
}

func serveQRCode() gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		var incReq models.APIRequestQRCode

		if err := ctx.ShouldBindJSON(&incReq); err == nil {

			incReq.Name, _ = normalizeInput(incReq.Name)
			incReq.City, _ = normalizeInput(incReq.City)
			incReq.Description, _ = normalizeInput(incReq.Description)

			paste, err := GeneratePaste(
				incReq.Amount,
				incReq.Name,
				incReq.City,
				incReq.Description,
				incReq.Transactionid,
				incReq.Pixkey,
			)
			if err != nil {
				ermsg := errors.New(err.Error())
				ctx.JSON(http.StatusBadRequest, ermsg)
				return
			}

			img, err := GenerateQR(
				incReq.Foregroundcolor,
				incReq.Backgroundcolor,
				paste,
			)
			if err != nil {
				ermsg := errors.New(err.Error())
				ctx.JSON(http.StatusBadRequest, ermsg)
				return
			}

			ctx.Data(http.StatusOK, "image/png", img)
			return
		}

	}
	return fn
}

func servePaste() gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		var incReq models.APIRequestQRCode

		if err := ctx.ShouldBindJSON(&incReq); err == nil {

			incReq.Name, _ = normalizeInput(incReq.Name)
			incReq.City, _ = normalizeInput(incReq.City)
			incReq.Description, _ = normalizeInput(incReq.Description)

			paste, err := GeneratePaste(
				incReq.Amount,
				incReq.Name,
				incReq.City,
				incReq.Description,
				incReq.Transactionid,
				incReq.Pixkey,
			)
			if err != nil {
				ermsg := errors.New(err.Error())
				ctx.JSON(http.StatusBadRequest, ermsg)
				return
			}

			ctx.Data(http.StatusOK, "text/plain", []byte(paste))
			return
		}

	}
	return fn
}

func serveQRLink() gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		var incReq models.APIRequestQRCode

		if err := ctx.ShouldBindJSON(&incReq); err == nil {

			incReq.Name, _ = normalizeInput(incReq.Name)
			incReq.City, _ = normalizeInput(incReq.City)
			incReq.Description, _ = normalizeInput(incReq.Description)

			paste, err := GeneratePaste(
				incReq.Amount,
				incReq.Name,
				incReq.City,
				incReq.Description,
				incReq.Transactionid,
				incReq.Pixkey,
			)
			if err != nil {
				ermsg := errors.New(err.Error())
				ctx.JSON(http.StatusBadRequest, ermsg)
				return
			}

			fln := fileNameGenerator(incReq.Pixkey)
			img, err := GenerateQR(
				incReq.Foregroundcolor,
				incReq.Backgroundcolor,
				paste,
			)
			if err != nil {
				ermsg := errors.New(err.Error())
				ctx.JSON(http.StatusBadRequest, ermsg)
				return
			}
			err = SaveImage(fln, img)
			if err != nil {
				ermsg := errors.New(err.Error())
				ctx.JSON(http.StatusBadRequest, ermsg)
				return
			}
			resp := models.PathResp{
				Path: "/stored/" + fln + ".png",
			}

			ctx.JSON(http.StatusOK, resp)
			return
		}

	}
	return fn
}
