package main

import "math"

func Degree2Radian(deg float64) float64 {
	return deg * math.Pi / 180.0
}

func Radian2Degree(rad float64) float64 {
	return rad * 180.0 / math.Pi
}

// 時・分を受け取り時針の角度[degree]を返す
func HourHandAngle(h, m int) float64 {
	return 30.0*float64(h) + float64(m)/2.0
}

// 分を受け取り分針の角度[degree]を返す
func MinuteHandAngle(m int) float64 {
	return 6.0 * float64(m)
}

// 時針と分針がなす小さい方の角[degree]を返す
func AcuteAngle(h, m int) float64 {
	diff := HourHandAngle(h, m) - MinuteHandAngle(m)
	if diff < 0 {
		diff *= -1
	}
	if diff > 180.0 {
		diff = 360.0 - diff
	}
	return diff
}
