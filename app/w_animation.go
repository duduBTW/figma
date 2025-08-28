package app

import (
	"sort"

	"github.com/dudubtw/figma/fmath"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type AnimatedProp struct {
	Name               string
	Base               float32
	SortedKeyframes    [][2]float32
	KeyframesMap       Float32Map
	KeyFrameCurveStart Float32Vec2Map
	KeyFrameCurveEnd   Float32Vec2Map

	InputValue string
}

func NewAnimatedProp(defaultValue float32, name string) AnimatedProp {
	return AnimatedProp{
		Base:               defaultValue,
		SortedKeyframes:    [][2]float32{},
		KeyframesMap:       map[float32]float32{},
		Name:               name,
		InputValue:         EMPTY,
		KeyFrameCurveStart: map[float32]rl.Vector2{},
		KeyFrameCurveEnd:   map[float32]rl.Vector2{},
	}
}

func (prop *AnimatedProp) SortedKeyframesTimeline() []float32 {
	result := []float32{}
	for _, keyframe := range prop.SortedKeyframes {
		result = append(result, keyframe[0])
	}
	return result
}

func (prop *AnimatedProp) InsertKeyframe(key, value float32) {
	prop.SortedKeyframes = append(prop.SortedKeyframes, [2]float32{key, value})
	sort.Slice(prop.SortedKeyframes, func(i, j int) bool {
		return prop.SortedKeyframes[i][0] < prop.SortedKeyframes[j][0]
	})
	prop.KeyframesMap[key] = value
}

// Sets value
func (prop *AnimatedProp) Set(key, value float32) {
	if len(prop.KeyframesMap) == 0 {
		prop.Base = value
		return
	}

	for index, kf := range prop.SortedKeyframes {
		if kf[0] == key {
			prop.SortedKeyframes[index][1] = value
			prop.KeyframesMap[key] = value
			return
		}
	}

	prop.InsertKeyframe(key, value)
}

// Sets value to the current selected frame
func (prop *AnimatedProp) SetCurrent(value float32) {
	key := float32(Apk.Workplace.SelectedFrame)
	prop.Set(key, value)
}

// MAYBE CACHE THIS SO YOU DONT HAVE TO RUN EVERY TIME IF NOTHING CHANGED
func (animatedProp AnimatedProp) KeyFramePosition(selectedFrame int) float32 {
	prop := animatedProp.Base
	keyframes := animatedProp.SortedKeyframes

	if len(keyframes) >= 2 {
		if selectedFrame <= int(keyframes[0][0]) {
			prop = keyframes[0][1]
		} else if selectedFrame >= int(keyframes[len(keyframes)-1][0]) {
			prop = keyframes[len(keyframes)-1][1]
		} else {
			var framesAround [2][2]float32
			for index, keyframe := range keyframes {
				if keyframe[0] > float32(selectedFrame) {
					framesAround = [2][2]float32{keyframes[index-1], keyframe}
					break
				}
			}
			framePos := fmath.InverseLerp(framesAround[0][0], framesAround[1][0], float32(selectedFrame))

			key := framesAround[0][0]
			startOffset := animatedProp.KeyFrameCurveStart[key]
			endOffset := animatedProp.KeyFrameCurveEnd[key]
			if startOffset.X != 0 || startOffset.Y != 0 || endOffset.X != 0 || endOffset.Y != 0 {
				var min float64 = 9999
				var max float64 = -9999
				for _, frame := range animatedProp.SortedKeyframes {
					value := float64(frame[1])
					if value > max {
						max = value
					}

					if value < min {
						min = value
					}
				}
				y := NewLinear().Domain(float64(Apk.Workplace.FramesRect.Y), float64(Apk.Workplace.FramesRect.Y+Apk.Workplace.FramesRect.Height)).Range(min, max)

				pos := rl.NewVector2(Apk.Workplace.GetXTimelineFrame(Apk.Workplace.FramesRect, framesAround[0][0]), float32(y.Invert(float64(framesAround[0][1]))))
				nextFramePos := rl.NewVector2(Apk.Workplace.GetXTimelineFrame(Apk.Workplace.FramesRect, framesAround[1][0]), float32(y.Invert(float64(framesAround[1][1]))))
				prop = float32(y.Scale(float64(CubicBezierPoint(
					pos,
					rl.NewVector2(pos.X+startOffset.X, pos.Y+startOffset.Y),
					rl.NewVector2(nextFramePos.X+endOffset.X, nextFramePos.Y+endOffset.Y),
					nextFramePos,
					framePos,
				).Y)))
			} else {
				prop = fmath.Lerp(framesAround[0][1], framesAround[1][1], fmath.Clamp(framePos, 0, 1))
			}

		}

	} else if len(keyframes) == 1 {
		prop = keyframes[0][1]
	}

	return prop
}
