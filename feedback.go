package minappapi

// PostFeedback 提交反馈
func PostFeedback(openID, formID, answer string) bool {

	feedback := Feedback{OpenID: openID, Answer: answer, FormID: formID}

	DB().Create(&feedback)
	if feedback.ID > 0 {
		return true
	}
	return false
}
