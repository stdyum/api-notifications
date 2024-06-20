package templates

import (
	"context"
	"errors"
)

type Templates interface {
	GetBodyByID(ctx context.Context, templateID string, vars map[string]string) (Template, error)
}

type templates struct {
	templates map[string]Template
}

func NewTemplates(t any) Templates {
	return &templates{
		templates: mockTemplates(),
	}
}

// GetBodyByID TODO
func (t *templates) GetBodyByID(_ context.Context, templateID string, _ map[string]string) (Template, error) {
	template, ok := t.templates[templateID]
	if !ok {
		return Template{}, errors.New("template not found")
	}

	return template, nil
}

func mockTemplates() map[string]Template {
	return map[string]Template{
		"sign_up_email_verification": {
			Subject: "Добро пожаловать",
			Content: "Добро пожаловать на Studyum. Вы получили данное письмо так как вы зарегистрировались на https://studyum.net",
		},
	}
}
