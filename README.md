# gifCaptcha
> gif 图像验证码
# 丰富自定义设置
* 图片大小
* 多颜色
* 文字模式
* 文字数量
* 干扰强度
# 示例代码
````
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
````
# 示例效果
## 黑白
![code](code.gif)
## 彩色
![code2](code2.gif)

# 感谢
https://github.com/afocus/captcha
