package user

import (
	"errors"
	"strings"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func (v *User)ValidateUser() []error{
	var err []error

	if len(v.Name)<6{
		err=append(err, errors.New("siz kiritgan ism 6 ta belgidan kam?"))
	}

	if v.Age<=0 || v.Age>=120{
		err=append(err, errors.New("yosh 0 dan kotta va 120 dan kichik bulishi shart?"))
	}

	if v.Name==""{
		err=append(err, errors.New("ism bo'sh bo'lishi kerak emas?"))
	}

	if v.Email==""{
		err=append(err, errors.New("email bo'sh bo'lishi kerak emas?"))
	}

	if !strings.Contains(v.Email,".com"){
		err=append(err, errors.New("email no'to'g'ri formatda?"))
	}

	return err
}
