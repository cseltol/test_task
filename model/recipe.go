package model

import "time"

type Step struct {
	StepDuration    time.Duration `json:"stepDuration"`
	StepDescription string        `json:"stepDescription"`
}

type Recipe struct {
	Id          uint32        `json:"id"`
	CreatedById uint32        `json:"createdById"`
	Name        string        `json:"name"`
	Ingredients []string      `json:"ingredients"`
	Description string        `json:"description"`
	Steps       []Step        `json:"steps"`
	Duration    time.Duration `json:"duration"`
}
