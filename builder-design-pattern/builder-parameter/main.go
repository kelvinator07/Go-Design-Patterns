package main

import (
	"fmt"
	"os"
	"strings"
)

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		fmt.Println("Sender Email not valid")
		os.Exit(0)
	}
	b.email.from = from
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	if !strings.Contains(to, "@") {
		fmt.Println("Recipient Email not valid")
		os.Exit(0)
	}
	b.email.to = to
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subject = subject
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}

func sendMailImpl(email *email) {

}

// This forces this user to use the builder
type build func(*EmailBuilder)

func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendMailImpl(&builder.email)
}


func main() {
	SendEmail(func(eb *EmailBuilder) {
		eb.From("foo@bar.com").
			To("bar@baz.com").
			Subject("Meeting").
			Body("Hello, do you want to meet?")
	})
}
