package minappapi

// PostFeedback 提交反馈
func PostFeedback(openID, formID, problem string) bool {

	feedback := Feedback{OpenID: openID, Problem: problem, FormID: formID}

	DB().Create(&feedback)
	if feedback.ID > 0 {
		return true
	}
	return false
}
