package mailServer


type MailServer interface {
	AddAuthenticationToken(key, secret string)
	AddRecipientToList(toMail string, name string)
	AddRecipientCcList(ccMail string, name string)
	AddRecipientBccList(bccMail string, name string)
	AddFrom(mail, name string)
	AddSubject(sub string)
	AddContent(data string)
	SendMail()(response interface{})
}

func GetMailServer(ServerProvider string) MailServer {
	switch ServerProvider {
	default:
		//Later when we have multiple providers we can add those cases
		return GetMailJetClient()
	}
}