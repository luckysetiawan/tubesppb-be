package response

type GeneralResponse struct {
	Status  uint        `form:"status" json:"status"`
	Message string      `form:"message" json:"message"`
	Data    interface{} `form:"data" json:"data"`
}
