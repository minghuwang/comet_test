package main

const (
	Control_Method_RoundAbout   = "Round_About"
	Control_Method_StopSign     = "Stop_Sign"
	Control_Method_TrafficLight = "Traffic_Light"
)

func ControlMethod(s int32, w int32, e int32, n int32) map[string]int32 {
	scores := make(map[string]int32)
	total := s + e + n + w
	if total >= 20 {
		scores[Control_Method_RoundAbout] = 50
		scores[Control_Method_StopSign] = 20
		scores[Control_Method_TrafficLight] = 90
	} else if total >= 10 && total < 20 {
		scores[Control_Method_RoundAbout] = 75
		scores[Control_Method_StopSign] = 30
		scores[Control_Method_TrafficLight] = 75
	} else {
		scores[Control_Method_RoundAbout] = 90
		scores[Control_Method_StopSign] = 40
		scores[Control_Method_TrafficLight] = 30
	}
	// item 5
	if s+e != n+w {
		scores[Control_Method_RoundAbout] += 10
	}
	return scores
}
