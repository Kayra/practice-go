// Type assertions and type switches are used to determine the underlying type of an interface value.

// A type assertion provides a way to recover the concrete type from an interface variable.
// A type switch provides a way to handle multiple types safely in a single action which is helpful when an interface variable might represent more than one concrete type.

package main

import "fmt"

type Media interface {
	Play()
}

type Audio struct {
	duration int
}

func (a Audio) Play() {
	fmt.Println("Playing audio for", a.duration, "seconds")
}

type Video struct {
	frameRate int
}

func (v Video) Play() {
	fmt.Println("Playing video at", v.frameRate, "fps")
}

func type_inference() {

	// var mediaPlayer Media = Audio{ duration: 120 }
	var mediaPlayer Media = Video{frameRate: 60}

	// Type assertion
	if _, ok := mediaPlayer.(Audio); ok {
		fmt.Println("is audio mediaPlayer")
	} else {
		fmt.Println("is not audio mediaPlayer")
	}

	// Type switch
	switch m := mediaPlayer.(type) {
	case Audio:
		fmt.Println("Should play audio")
		m.Play()
	case Video:
		fmt.Println("Should play video")
		m.Play()
	default:
		fmt.Println("Nothing to play")
	}
}
