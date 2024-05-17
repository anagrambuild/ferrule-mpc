package comittee

type Epoch struct {
	Proof []byte `json:"proof"`
	Epoch uint64 `json:"epoch"`
}

type EpochSource interface {
	Epoch() *Epoch
	Expired(*Epoch) bool
}

type OpBlockEpochSource struct {
	epoch *Epoch
}

func (o *OpBlockEpochSource) Epoch() *Epoch {
	return o.epoch
}

func NewOpBlockEpochSource(
	opUrl string,
) *OpBlockEpochSource {
	return &OpBlockEpochSource{}
}

func (o *OpBlockEpochSource) Start() {

}
