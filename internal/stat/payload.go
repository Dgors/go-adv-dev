package stat

type GetStatResponse struct {
	Period string `json:"period"`
	Sum    int64  `json:"sum"`
}
