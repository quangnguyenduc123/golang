package main

import (
	"fmt"
	"strings"
)

type Email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email Email
}

func (e *EmailBuilder) From(from string) *EmailBuilder {
	// we can do validation here
	if !strings.Contains(from, "@") {
		panic("Email not validate")
	}
	e.email.from = from
	return e
}

func (e *EmailBuilder) To(to string) *EmailBuilder {
	e.email.to = to
	return e
}

func (e *EmailBuilder) Subject(subject string) *EmailBuilder {
	e.email.subject = subject
	return e
}

func (e *EmailBuilder) Body(body string) *EmailBuilder {
	e.email.body = body
	return e
}

func newEmailBuilder() *EmailBuilder{
	return &EmailBuilder{Email{}}
}
func main() {
	e := newEmailBuilder()
	e.From("123").To("123")
	fmt.Println(e.email)
}
