package gifCaptcha

import (
	"github.com/labstack/echo"
	"image/color"
	"image/gif"
	"net/http"
	"testing"
)

func TestGifCaptcha(t *testing.T) {
	e := echo.New()
	gifCaptcha := New()
	gifCaptcha.SetFrontColor(color.Black,color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "测试")
	})
	e.GET("/img", func(c echo.Context) error {
		gifData, code := gifCaptcha.RangCaptcha()
		println(code)
		gif.EncodeAll(c.Response().Writer, gifData)
		return nil
	})
	e.Logger.Fatal(e.Start(":1323"))
}
