package mailServer

import (
	mailJet "github.com/mailjet/mailjet-apiv3-go"
)


type MailJet struct {
	Key string
	Secret string
	*mailJet.Client
	From mailJet.RecipientV31
	To []mailJet.RecipientV31
	Cc []mailJet.RecipientV31
	Bcc []mailJet.RecipientV31
	mailJet.InfoMessagesV31
}

func (MJ *MailJet) AddFrom(mail, name string) {
	MJ.From.Name = name
	MJ.From.Email = mail
}

func (MJ *MailJet) AddAuthenticationToken(key, secret string) {
		MJ.Client = mailJet.NewMailjetClient(key,secret)
}

func (MJ *MailJet) AddRecipientToList(toMail string, name string) {
	if MJ.To == nil {
		MJ.To = []mailJet.RecipientV31{}
	}
	MJ.To = append(MJ.To, mailJet.RecipientV31{Email: toMail,
		Name:name,
	})
}

func (MJ *MailJet) AddRecipientCcList(ccMail string, name string) {
	if MJ.Cc == nil {
		MJ.Cc = []mailJet.RecipientV31{}
	}
	MJ.Cc = append(MJ.To, mailJet.RecipientV31{Email: ccMail,
		Name:name,
	})
}

func (MJ *MailJet) AddRecipientBccList(bccMail string, name string) {
	if MJ.Bcc == nil {
		MJ.Bcc = []mailJet.RecipientV31{}
	}
	MJ.Bcc = append(MJ.To, mailJet.RecipientV31{Email: bccMail,
		Name:name,
	})
}

func (MJ *MailJet) AddSubject(sub string) {
	MJ.InfoMessagesV31.Subject = sub
}

func (MJ *MailJet) AddContent(data string) {
	MJ.TextPart = data
}

func (MJ *MailJet) SendMail() (response interface{}) {
	MJ.InfoMessagesV31.From = &MJ.From
	To := mailJet.RecipientsV31(MJ.To)
	MJ.InfoMessagesV31.To = &To
	Cc := mailJet.RecipientsV31(MJ.Cc)
	MJ.InfoMessagesV31.Cc = &Cc
	Bcc := mailJet.RecipientsV31(MJ.Bcc)
	MJ.InfoMessagesV31.Bcc = &Bcc

	res, err := MJ.Client.SendMailV31(
		&mailJet.MessagesV31{Info:[]mailJet.InfoMessagesV31 {MJ.InfoMessagesV31}})
	if err != nil {
		panic(err)
	}
	return res
}

func GetMailJetClient() MailServer{
	return &MailJet{
		InfoMessagesV31:mailJet.InfoMessagesV31{CustomID:"Codemax Demo"},
	}
}