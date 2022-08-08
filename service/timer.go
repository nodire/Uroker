package service

var (
	lessonStartTime string
	lessonCountdown float64
	lessonStage     uint8
)

type timer struct {
	lessonStartTime string
	lessonCountdown float64
	lessonStage     uint8
}

func previousLesson() timer {
	if clock.After(l1End) && clock.Before(l2End) {
		lessonStartTime = "09:00:00"
		lessonCountdown = clock.Sub(l1Start).Minutes()
		lessonStage = 1
	} else if clock.After(l2End) && clock.Before(l3End) {
		lessonStartTime = "10:35:00"
		lessonCountdown = clock.Sub(l2Start).Minutes()
		lessonStage = 2
	} else if clock.After(l3End) && clock.Before(l4End) {
		lessonStartTime = "12:25:00"
		lessonCountdown = clock.Sub(l3Start).Minutes()
		lessonStage = 3
	} else if clock.After(l4End) && clock.Before(l5End) {
		lessonStartTime = "14:00:00"
		lessonCountdown = clock.Sub(l4Start).Minutes()
		lessonStage = 4
	} else if clock.After(l5End) && clock.Before(l6End) {
		lessonStartTime = "15:50:00"
		lessonCountdown = clock.Sub(l5Start).Minutes()
		lessonStage = 5
	} else {
		lessonStartTime = "00:00:00"
		lessonStage = 0
	}
	return timer{
		lessonStartTime: lessonStartTime,
		lessonCountdown: lessonCountdown,
		lessonStage:     lessonStage,
	}
}

func currentLesson() timer {
	if clock.After(l1Start) && clock.Before(l1End) {
		lessonStartTime = "09:00:00"
		lessonCountdown = l1End.Sub(clock).Minutes()
		lessonStage = 1
	} else if clock.After(l2Start) && clock.Before(l2End) {
		lessonStartTime = "10:35:00"
		lessonCountdown = l2End.Sub(clock).Minutes()
		lessonStage = 2
	} else if clock.After(l3Start) && clock.Before(l3End) {
		lessonStartTime = "12:25:00"
		lessonCountdown = l3End.Sub(clock).Minutes()
		lessonStage = 3
	} else if clock.After(l4Start) && clock.Before(l4End) {
		lessonStartTime = "14:00:00"
		lessonCountdown = l4End.Sub(clock).Minutes()
		lessonStage = 4
	} else if clock.After(l5Start) && clock.Before(l5End) {
		lessonStartTime = "15:50:00"
		lessonCountdown = l5End.Sub(clock).Minutes()
		lessonStage = 5
	} else if clock.After(l6Start) && clock.Before(l6End) {
		lessonStartTime = "17:25:00"
		lessonCountdown = l6End.Sub(clock).Minutes()
		lessonStage = 6
	} else {
		lessonStartTime = "00:00:00"
		lessonCountdown = 0
		lessonStage = 0
	}
	return timer{
		lessonStartTime: lessonStartTime,
		lessonCountdown: lessonCountdown,
		lessonStage:     lessonStage,
	}
}

func nextLesson() timer {
	if clock.After(l1Start) && clock.Before(l1End) {
		lessonStartTime = "10:35:00"
		lessonCountdown = l2Start.Sub(clock).Minutes()
		lessonStage = 2
	} else if clock.After(l2Start) && clock.Before(l2End) {
		lessonStartTime = "12:25:00"
		lessonCountdown = l3Start.Sub(clock).Minutes()
		lessonStage = 3
	} else if clock.After(l3Start) && clock.Before(l3End) {
		lessonStartTime = "14:00:00"
		lessonCountdown = l4Start.Sub(clock).Minutes()
		lessonStage = 4
	} else if clock.After(l4Start) && clock.Before(l4End) {
		lessonStartTime = "15:50:00"
		lessonCountdown = l5Start.Sub(clock).Minutes()
		lessonStage = 5
	} else if clock.After(l5Start) && clock.Before(l5End) {
		lessonStartTime = "17:25:00"
		lessonCountdown = 0
		lessonStage = 6
	}
	return timer{
		lessonStartTime: lessonStartTime,
		lessonCountdown: lessonCountdown,
		lessonStage:     lessonStage,
	}
}
