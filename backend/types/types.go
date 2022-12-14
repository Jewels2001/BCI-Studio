package types

type Ch struct {
	Data [][5]float64 `json:"hist"`
}

type Recording struct {
	Ch0     Ch  `json:"ch0"`
	Ch1     Ch  `json:"ch1"`
	Ch2     Ch  `json:"ch2"`
	Ch3     Ch  `json:"ch3"`
	Samples int `json:"samples"`
}

type FileDirectory struct {
	Files []string `json:"files"`
}

type Message struct {
	Message string `json:"message"`
}
