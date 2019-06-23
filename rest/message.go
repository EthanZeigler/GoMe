package rest

type Message struct {
	MemberNickname string `json:"nickname"`
	Text string `json:"text"`
	ImageURL string `json:"image_url"`
	_AttachmentJSON string 
}

type MessageAttachment struct {
	
}