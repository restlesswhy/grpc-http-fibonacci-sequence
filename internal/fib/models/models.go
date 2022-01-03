package models

type FibSeq struct {
	Seq map[int32]string `json:"seq"`
}

type Interval struct {
	From int32 `json:"from,omitempty"`
	To   int32 `json:"to,omitempty"`
}
