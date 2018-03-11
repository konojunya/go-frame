package frame

import (
	"fmt"
	"reflect"
)

var (
	frameHorizontal   = "━"
	frameTopLeft      = "┏"
	frameTopCenter    = "┳"
	frameTopRight     = "┓"
	frameBottomLeft   = "┗"
	frameBottomRight  = "┛"
	frameBottomCenter = "┻"
	frameMiddleLeft   = "┣"
	frameMiddleRight  = "┫"
	frameMiddleCenter = "╋"
	padding           = 1
)

// Print 表示する
func Print(structure interface{}) {
	rt, rv := reflect.TypeOf(structure), reflect.ValueOf(structure)
	frameNameLenMax := 0
	valueMaxLen := 0

	valuesMap := make(map[string]string)

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		frameName := field.Tag.Get("frame")
		if frameName == "" {
			frameName = field.Name
		}
		value := fmt.Sprint(rv.Field(i).Interface())

		if len(frameName) > frameNameLenMax {
			frameNameLenMax = len(frameName)
		}
		if len(value) > valueMaxLen {
			valueMaxLen = len(value)
		}

		valuesMap[frameName] = value

	}

	var count int

	fmt.Println(generateTopLine(frameNameLenMax, valueMaxLen))
	for k, v := range valuesMap {
		count++
		fmt.Println(fmt.Sprintf(frameMiddleLeft + createElementString(k, frameNameLenMax-len(k)) + frameMiddleCenter + createElementString(v, valueMaxLen-len(v)) + frameMiddleRight))
		if count == len(valuesMap) {
			break
		}
		fmt.Println(generateMiddleLine(frameNameLenMax, valueMaxLen))
	}
	fmt.Println(generateBottomLine(frameNameLenMax, valueMaxLen))
}

func createElementString(key string, diff int) string {
	var paddingStr string
	for i := 0; i < padding; i++ {
		paddingStr += " "
	}

	var diffStr string
	for i := 0; i < diff; i++ {
		diffStr += " "
	}

	line := paddingStr + key + diffStr + paddingStr
	return line
}

func generateMiddleLine(frameNameLenMax, valueMaxLen int) string {
	var line string

	line += frameMiddleLeft

	for i := 0; i < frameNameLenMax+padding*2; i++ {
		line += frameHorizontal
	}

	line += frameMiddleCenter

	for i := 0; i < valueMaxLen+padding*2; i++ {
		line += frameHorizontal
	}

	line += frameMiddleRight

	return line
}

func generateTopLine(frameNameLenMax, valueMaxLen int) string {
	var line string

	line += frameTopLeft

	for i := 0; i < frameNameLenMax+padding*2; i++ {
		line += frameHorizontal
	}

	line += frameTopCenter

	for i := 0; i < valueMaxLen+padding*2; i++ {
		line += frameHorizontal
	}

	line += frameTopRight

	return line
}

func generateBottomLine(frameNameLenMax, valueMaxLen int) string {
	var line string

	line += frameBottomLeft

	for i := 0; i < frameNameLenMax+padding*2; i++ {
		line += frameHorizontal
	}

	line += frameBottomCenter

	for i := 0; i < valueMaxLen+padding*2; i++ {
		line += frameHorizontal
	}

	line += frameBottomRight

	return line
}
