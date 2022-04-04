package main

import (
	"bytes"
	"text/template"
)

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan condition pada template.
// Lengkapi function ExecuteToByteBuffer dan variabel textTemplate sehingga memiliki keluaran:
// 1. Jika saldo sama dengan 0, cetak "Akun Tony dengan Nomor 1002321 tidak memiliki saldo."
// 2. Jika saldo lebih dari 0, cetak "Akun Tony dengan Nomor 1002321 memiliki saldo sebesar $1000."

type Account struct {
	Name    string
	Number  string
	Balance int
}

func ExecuteToByteBuffer(account Account) ([]byte, error) {
	var textTemplate string
	if account.Balance == 0 {
		textTemplate = "Akun {{ .Name }} dengan Nomor {{ .Number }} tidak memiliki saldo."
	} else if account.Balance > 0 {
		textTemplate = "Akun {{ .Name }} dengan Nomor {{ .Number }} memiliki saldo sebesar ${{ .Balance }}."
	}

	tmp, err := template.New("Template_1").Parse(textTemplate)
	if err != nil {
		panic(err)
	}
	var b bytes.Buffer
	err = tmp.Execute(&b, account)
	if err != nil {
		panic(err)
	}
	return b.Bytes(), nil
}
