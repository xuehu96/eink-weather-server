package piximg

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/mattn/go-runewidth"
	"log"
	"regexp"
	"time"
)

func Create(w int, h int) *gg.Context {
	dc := gg.NewContext(w, h)
	return dc
}

// DrawDate 画日期
func DrawDate(dc *gg.Context) {
	// 加载字体
	if err := dc.LoadFontFace("./static/sarasa-mono-sc-nerd-regular.ttf", 28); err != nil {
		panic(err)
	}
	// 设置颜色
	dc.SetRGB255(255, 0, 0) // 设置画笔颜色为红色
	dc.SetLineWidth(1)
	dc.DrawLine(0, 36, 200, 36)
	dc.DrawLine(0, 163, 200, 163)
	dc.DrawLine(100, 163, 100, 200)
	dc.StrokePreserve()
	// 显示日期
	timeNow := time.Now()
	var WeekDayMap = map[string]string{
		"Monday":    "周一",
		"Tuesday":   "周二",
		"Wednesday": "周三",
		"Thursday":  "周四",
		"Friday":    "周五",
		"Saturday":  "周六",
		"Sunday":    "周日",
	}
	str := timeNow.Format("1月02日 ") + WeekDayMap[timeNow.Weekday().String()]
	sWidth, sHeight := dc.MeasureString(str)
	dc.DrawString(str, (200-sWidth)/2, sHeight+5)

	if err := dc.LoadFontFace("./static/SourceHanSansCN-Regular.ttf", 12); err != nil {
		panic(err)
	}
	str = timeNow.Format("15:04")

	dc.DrawString(str, 2, 52)

}

// DrawWeatherIcon 画天气图标
func DrawWeatherIcon(dc *gg.Context, key string) {
	img, err := gg.LoadImage(".\\static\\icons\\80\\" + key + ".png")
	if err != nil {
		fmt.Println(err)
		return
	}

	dc.DrawImage(img, 3, 40)
}

var stripAnsiEscapeRegexp = regexp.MustCompile(`(\x9B|\x1B\[)[0-?]*[ -/]*[@-~]`)

func stripAnsiEscape(s string) string {
	return stripAnsiEscapeRegexp.ReplaceAllString(s, "")
}
func realLength(s string) int {
	return runewidth.StringWidth(stripAnsiEscape(s))
}

func DrawWeatherDescribe(dc *gg.Context, str string) {
	if err := dc.LoadFontFace("./static/SourceHanSansCN-Regular.ttf", 18); err != nil {
		panic(err)
	}
	if realLength(str) > 20 {
		strRune := []rune(str)
		str = string(strRune[:10]) + "\n" + string(strRune[11:])
	}
	dc.DrawStringWrapped(str, 200/2, 140, 0.5, 0.5, 200-10, 1.3, gg.AlignCenter)
}

func DrawForecast(dc *gg.Context, d string, t string) {
	if err := dc.LoadFontFace("./static/LanaPixel.ttf", 12); err != nil {
		panic(err)
	}
	dc.DrawString(d, 90, 100)
	dc.DrawString(t, 90, 116)
}
func DrawTemp(dc *gg.Context, tmp string, f string) {
	if err := dc.LoadFontFace("./static/SourceHanSansCN-Regular.ttf", 18); err != nil {
		panic(err)
	}
	dc.DrawString("℃", 168, 80)

	if err := dc.LoadFontFace("./static/sarasa-mono-sc-nerd-regular.ttf", 36); err != nil {
		panic(err)
	}
	dc.DrawString(tmp, 90, 77)

	if err := dc.LoadFontFace("./static/SourceHanSansCN-Regular.ttf", 12); err != nil {
		panic(err)
	}
	sw_, _ := dc.MeasureString(f)

	dc.DrawString(f, 196-sw_, 52)
}

func DrawBottom(dc *gg.Context) {
	// 加载字体
	if err := dc.LoadFontFace("./static/LanaPixel.ttf", 12); err != nil {
		panic(err)
	}
	dc.DrawString("今天", 2, 180)
	dc.DrawString("还剩", 2, 195)

	if err := dc.LoadFontFace("./static/sarasa-mono-sc-nerd-regular.ttf", 26); err != nil {
		panic(err)
	}

	h, m, _ := time.Now().Clock()
	past := h*60 + m
	remain := 1440 - past
	str := fmt.Sprintf("%.1f%%", float32(remain*100)/1440)
	dc.DrawString(str, 30, 193)
}

func DrawCountdown(dc *gg.Context, txt string, date string) {
	// 加载字体
	if err := dc.LoadFontFace("./static/LanaPixel.ttf", 12); err != nil {
		panic(err)
	}
	dc.DrawString(txt, 107, 180)
	dc.DrawString("还有", 107, 195)

	if err := dc.LoadFontFace("./static/sarasa-mono-sc-nerd-regular.ttf", 26); err != nil {
		panic(err)
	}

	day_, err := time.ParseInLocation("20060102", date, time.Local)
	if err != nil {
		log.Println("1." + err.Error())
		dc.DrawString("???", 137, 193)
		return
	}
	now := time.Now()
	str := fmt.Sprintf("%d", day_.Sub(now)/(time.Duration(24)*time.Hour))
	dc.DrawString(str, 137, 193)
	if err := dc.LoadFontFace("./static/sarasa-mono-sc-nerd-regular.ttf", 18); err != nil {
		panic(err)
	}
	sw_, _ := dc.MeasureString(str)
	ss_, _ := dc.MeasureString("天")
	dc.DrawString("天", 134+sw_+ss_, 190)

}
