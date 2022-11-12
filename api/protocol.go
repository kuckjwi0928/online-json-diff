package api

type Request interface {
	Validate() error
}

type DiffTargetRequest struct {
	OriginURL  string `form:"originURL" binding:"required"`
	CompareURL string `form:"compareURL" binding:"required"`
	HeaderKeys string `form:"headerKeys"`
}

func (d *DiffTargetRequest) Validate() error {
	// validate
	return nil
}

type DiffTargetResponse struct {
	Left  string `json:"left"`
	Right string `json:"right"`
}
