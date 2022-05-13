package color

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	Black = 30 + iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	LightGray
)

const (
	DarkGray = 90 + iota
	LightRed
	LightGreen
	LightYellow
	LightBlue
	LightCyan
	LightMagenta
	White
)

const (
	BBlack = 40 + iota
	BRed
	BGreen
	BYellow
	BBlue
	BMagenta
	BCyan
	BLightGray
)

const (
	BDarkGray = 100 + iota
	BLightRed
	BLightGreen
	BLightYellow
	BLightBlue
	BLightMagenta
	BLightCyan
	BWhite
)

const (
	Reset = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

const colorFormat = "\u001B[%dm"

const colorDefault = "\u001B[0m"

const sign = len("<c?red?>")

// 映射表
var colorSign = map[string]int{
	"black":         Black,
	"red":           Red,
	"green":         Green,
	"yellow":        Yellow,
	"blue":          Blue,
	"magenta":       Magenta,
	"cyan":          Cyan,
	"lightgray":     LightGray,
	"darkgray":      DarkGray,
	"lighteed":      LightRed,
	"lightgreen":    LightGreen,
	"lightyellow":   LightYellow,
	"lightblue":     LightBlue,
	"lightcyan":     LightCyan,
	"lightmagenta":  LightMagenta,
	"white":         White,
	"bblack":        BBlack,
	"bred":          BRed,
	"bgreen":        BGreen,
	"byellow":       BYellow,
	"bblue":         BBlue,
	"bmagenta":      BMagenta,
	"bcyan":         BCyan,
	"blightgray":    BLightGray,
	"bdarkgray":     BDarkGray,
	"blightred":     BLightRed,
	"blightgreen":   BLightGreen,
	"blightyellow":  BLightYellow,
	"blightblue":    BLightBlue,
	"blightmagenta": BLightMagenta,
	"blightcyan":    BLightCyan,
	"bwhite":        BWhite,
	"reset":         Reset,
	"bold":          Bold,
	"faint":         Faint,
	"italic":        Italic,
	"underline":     Underline,
	"blinkslow":     BlinkSlow,
	"blinkrapid":    BlinkRapid,
	"reversevideo":  ReverseVideo,
	"concealed":     Concealed,
	"crossedout":    CrossedOut,
}

func Coat(s string, color int) string {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf(colorFormat, color) + s + colorDefault)
	return buf.String()
}

func CoatMany(s string, colors ...int) string {
	buf := bytes.Buffer{}
	for i := range colors {
		buf.WriteString(fmt.Sprintf(colorFormat, colors[i]))
	}
	buf.WriteString(s + colorDefault)
	return buf.String()
}

func CoatFormat(s string) string {
	n := len(s)
	switch {
	case n == 0:
		return ""
	case n == 1:
		return s
	case n <= sign:
		return s
	}
	tempM := make(map[string]struct{})
	for i := 0; i < len(s)-2; i++ {
		if s[i] == '<' && s[i+1] == 'c' && s[i+2] == '?' {
			var begin, end = i + 3, i + 3
			for y := begin + 2; y < len(s)-2; y++ {
				if s[y] == '?' && s[y+1] == '>' {
					end = y
					break
				}
			}
			if _, ok := tempM[s[begin:end]]; !ok {
				tempM[s[begin:end]] = struct{}{}
				s = strings.ReplaceAll(s, s[begin-3:end+2], fmt.Sprintf(colorFormat, colorSign[strings.ToLower(s[begin:end])]))
			}
		}
	}
	return s
}
