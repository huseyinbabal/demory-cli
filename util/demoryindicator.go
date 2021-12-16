package util

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

const Frames = `⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏`
const ClearLine = "\r\033[K"
const ProgressBarWidth = 30

type LoadingIndicator struct {
	startTime   time.Time
	position    int
	isActive    bool
	message     string
	totalSteps  int
	currentStep int
}

func NewLoadingIndicator(message string, totalSteps int) *LoadingIndicator {
	s := &LoadingIndicator{
		message:    message,
		totalSteps: totalSteps,
	}
	return s
}

func (s *LoadingIndicator) SetStep(message string, currentStep int) {
	s.currentStep = currentStep
	s.message = message
}

func (s *LoadingIndicator) Start() {
	s.isActive = true
	s.startTime = time.Now()
	go func() {
		for s.isActive {
			white := color.New(color.Bold, color.FgHiWhite)
			indicator := color.New(color.Bold)
			hiBlack := color.New(color.FgHiBlack)
			fmt.Printf("%s %s %s %s ", ClearLine+s.getProgressBar(), indicator.Sprint(s.next()), white.Sprint(s.message), hiBlack.Sprint(s.getPastTime()))
			time.Sleep(200 * time.Millisecond)
		}
	}()
}

func (s *LoadingIndicator) Stop() string {
	s.isActive = false
	fmt.Printf("%s", ClearLine)
	return s.getPastTime()
}

func (s *LoadingIndicator) getProgressBar() string {
	white := color.New(color.FgHiWhite)
	fillColor := color.New(color.FgHiCyan)
	var progressBar = fillColor.Sprintf("[")
	filledSteps := (ProgressBarWidth * s.currentStep) / s.totalSteps
	for i := 0; i < ProgressBarWidth; i++ {
		if filledSteps > i {
			progressBar += fillColor.Sprintf("=")
		} else if filledSteps == i {
			progressBar += fillColor.Sprintf(">")
		} else {
			progressBar += white.Sprintf("-")
		}
	}
	progressBar += fillColor.Sprintf("] %d/%d", s.currentStep, s.totalSteps)
	return progressBar
}

func (s *LoadingIndicator) getPastTime() string {
	since := time.Since(s.startTime)
	minutes := int(since.Seconds() / 60)
	seconds := int(since.Seconds()) % 60

	if minutes == 0 {
		return fmt.Sprintf("%ds", seconds)
	} else {
		return fmt.Sprintf("%dm%ds", minutes, seconds)
	}
}

func (s *LoadingIndicator) next() string {
	r := []rune(Frames)[(s.position)%len([]rune(Frames))]
	s.position++
	return string(r)
}
