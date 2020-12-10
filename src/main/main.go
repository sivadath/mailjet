package main

import (
	"encoding/json"
	"mailServer"
	"os"
	"samples"
	"fmt"
)

type Person struct {
	Name string
	Mail string
}

type Sample struct {
	From Person
	To []Person
	Cc []Person
	Bcc []Person
	Subject string
	Data string
}

func main() {
	var sampleMail Sample
	err := json.Unmarshal([]byte(samples.Sample1), &sampleMail)
	if err != nil {
		panic("Wrong sample data")
	}
	MS := mailServer.GetMailServer("")
	MS.AddAuthenticationToken(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))
	MS.AddFrom(sampleMail.From.Mail, sampleMail.From.Name)
	for _, bcc := range sampleMail.Bcc {
		MS.AddRecipientBccList(bcc.Mail, bcc.Name)
	}
	for _, to := range sampleMail.To {
		MS.AddRecipientToList(to.Mail, to.Name)
	}
	for _, cc := range sampleMail.Cc {
		MS.AddRecipientCcList(cc.Mail, cc.Name)
	}
	MS.AddSubject(sampleMail.Subject)
	MS.AddContent(sampleMail.Data)
	fmt.Println(MS.SendMail())

}