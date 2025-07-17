package app

import (
	"fmt"
	"sort"

	"github.com/dudubtw/figma/fmath"
)

type AnimatedProp struct {
	Name            string
	Base            float32
	SortedKeyframes [][2]float32
	KeyframesMap    map[float32]float32

	InputValue string
}

func NewAnimatedProp(defaultValue float32, name string) AnimatedProp {
	return AnimatedProp{Base: defaultValue, SortedKeyframes: [][2]float32{}, KeyframesMap: map[float32]float32{}, Name: name, InputValue: EMPTY}
}

func (prop *AnimatedProp) InsertKeyframe(key, value float32) {
	fmt.Println("Inserting keyframe!", key, value, prop)
	prop.SortedKeyframes = append(prop.SortedKeyframes, [2]float32{key, value})
	sort.Slice(prop.SortedKeyframes, func(i, j int) bool {
		return prop.SortedKeyframes[i][0] < prop.SortedKeyframes[j][0]
	})
	prop.KeyframesMap[key] = value
}

func (prop *AnimatedProp) Set(value float32) {
	if len(prop.KeyframesMap) == 0 {
		prop.Base = value
		return
	}

	key := float32(Apk.State.SelectedFrame)
	for index, kf := range prop.SortedKeyframes {
		if kf[0] == key {
			prop.SortedKeyframes[index][1] = value
			prop.KeyframesMap[key] = value
			return
		}
	}

	prop.InsertKeyframe(key, value)
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
			prop = fmath.Lerp(framesAround[0][1], framesAround[1][1], fmath.Clamp(framePos, 0, 1))
		}

	} else if len(keyframes) == 1 {
		prop = keyframes[0][1]
	}

	return prop
}
