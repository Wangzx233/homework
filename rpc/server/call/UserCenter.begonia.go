// Code generated by Begonia. DO NOT EDIT.
// versions:
// 	Begonia v1.0.2
// source: server\server.go
// begonia client file

package call

import (
	"github.com/MashiroC/begonia/app"
	appClient "github.com/MashiroC/begonia/app/client"
	"github.com/MashiroC/begonia/app/coding"
)

var (
	UserCenterService appClient.Service

	_UserCenterServiceSayHello appClient.RemoteFunSync

	_UserCenterServiceSayHelloInSchema = `
{
			"namespace":"begonia.func.SayHello",
			"type":"record",
			"name":"In",
			"fields":[
				{"name":"F1","type":"string","alias":"name"}

			]
		}`
	_UserCenterServiceSayHelloOutSchema = `
{
			"namespace":"begonia.func.SayHello",
			"type":"record",
			"name":"Out",
			"fields":[
				{"name":"F1","type":"string"}

			]
		}`
	_UserCenterServiceSayHelloInCoder  coding.Coder
	_UserCenterServiceSayHelloOutCoder coding.Coder

	_UserCenterServiceRegister appClient.RemoteFunSync

	_UserCenterServiceRegisterInSchema = `
{
			"namespace":"begonia.func.Register",
			"type":"record",
			"name":"In",
			"fields":[
				{"name":"F1","type":"string","alias":"username"}
,{"name":"F2","type":"string","alias":"password"}

			]
		}`
	_UserCenterServiceRegisterOutSchema = `
{
			"namespace":"begonia.func.Register",
			"type":"record",
			"name":"Out",
			"fields":[
				{"name":"F1","type":"boolean"}
,{"name":"F1","type":"string"}

			]
		}`
	_UserCenterServiceRegisterInCoder  coding.Coder
	_UserCenterServiceRegisterOutCoder coding.Coder

	_UserCenterServiceLogin appClient.RemoteFunSync

	_UserCenterServiceLoginInSchema = `
{
			"namespace":"begonia.func.Login",
			"type":"record",
			"name":"In",
			"fields":[
				{"name":"F1","type":"string","alias":"username"}
,{"name":"F2","type":"string","alias":"password"}

			]
		}`
	_UserCenterServiceLoginOutSchema = `
{
			"namespace":"begonia.func.Login",
			"type":"record",
			"name":"Out",
			"fields":[
				{"name":"F1","type":"boolean"}
,{"name":"F1","type":"int"}

			]
		}`
	_UserCenterServiceLoginInCoder  coding.Coder
	_UserCenterServiceLoginOutCoder coding.Coder

	_UserCenterServiceChangePassword appClient.RemoteFunSync

	_UserCenterServiceChangePasswordInSchema = `
{
			"namespace":"begonia.func.ChangePassword",
			"type":"record",
			"name":"In",
			"fields":[
				{"name":"F1","type":"string","alias":"username"}
,{"name":"F2","type":"string","alias":"newPassword"}

			]
		}`
	_UserCenterServiceChangePasswordOutSchema = `
{
			"namespace":"begonia.func.ChangePassword",
			"type":"record",
			"name":"Out",
			"fields":[
				
			]
		}`
	_UserCenterServiceChangePasswordInCoder  coding.Coder
	_UserCenterServiceChangePasswordOutCoder coding.Coder
)

type _UserCenterServiceSayHelloIn struct {
	F1 string
}

type _UserCenterServiceSayHelloOut struct {
	F1 string
}

type _UserCenterServiceRegisterIn struct {
	F1 string
	F2 string
}

type _UserCenterServiceRegisterOut struct {
	F1 bool
	F2 string
}

type _UserCenterServiceLoginIn struct {
	F1 string
	F2 string
}

type _UserCenterServiceLoginOut struct {
	F1 bool
	F2 int
}

type _UserCenterServiceChangePasswordIn struct {
	F1 string
	F2 string
}

type _UserCenterServiceChangePasswordOut struct {
}

func init() {
	app.ServiceAppMode = app.Ast

	bService, err := BegoniaCli.Service("Echo")
	if err != nil {
		panic(err)
	}

	_UserCenterServiceSayHello, err = bService.FuncSync("SayHello")

	_UserCenterServiceSayHelloInCoder, err = coding.NewAvro(_UserCenterServiceSayHelloInSchema)
	if err != nil {
		panic(err)
	}
	_UserCenterServiceSayHelloOutCoder, err = coding.NewAvro(_UserCenterServiceSayHelloOutSchema)
	if err != nil {
		panic(err)
	}

	_UserCenterServiceRegister, err = bService.FuncSync("Register")

	_UserCenterServiceRegisterInCoder, err = coding.NewAvro(_UserCenterServiceRegisterInSchema)
	if err != nil {
		panic(err)
	}
	_UserCenterServiceRegisterOutCoder, err = coding.NewAvro(_UserCenterServiceRegisterOutSchema)
	if err != nil {
		panic(err)
	}

	_UserCenterServiceLogin, err = bService.FuncSync("Login")

	_UserCenterServiceLoginInCoder, err = coding.NewAvro(_UserCenterServiceLoginInSchema)
	if err != nil {
		panic(err)
	}
	_UserCenterServiceLoginOutCoder, err = coding.NewAvro(_UserCenterServiceLoginOutSchema)
	if err != nil {
		panic(err)
	}

	_UserCenterServiceChangePassword, err = bService.FuncSync("ChangePassword")

	_UserCenterServiceChangePasswordInCoder, err = coding.NewAvro(_UserCenterServiceChangePasswordInSchema)
	if err != nil {
		panic(err)
	}
	_UserCenterServiceChangePasswordOutCoder, err = coding.NewAvro(_UserCenterServiceChangePasswordOutSchema)
	if err != nil {
		panic(err)
	}

}

func SayHello(name string) (F1 string, err error) {
	var in _UserCenterServiceSayHelloIn
	in.F1 = name

	b, err := _UserCenterServiceSayHelloInCoder.Encode(in)
	if err != nil {
		panic(err)
	}

	begoniaResTmp, err := _UserCenterServiceSayHello(b)
	if err != nil {
		return
	}

	var out _UserCenterServiceSayHelloOut
	err = _UserCenterServiceSayHelloOutCoder.DecodeIn(begoniaResTmp.([]byte), &out)
	if err != nil {
		panic(err)
	}

	F1 = out.F1

	return
}

func Register(username string, password string) (F1 bool, F2 string, err error) {
	var in _UserCenterServiceRegisterIn
	in.F1 = username
	in.F2 = password

	b, err := _UserCenterServiceRegisterInCoder.Encode(in)
	if err != nil {
		panic(err)
	}

	begoniaResTmp, err := _UserCenterServiceRegister(b)
	if err != nil {
		return
	}

	var out _UserCenterServiceRegisterOut
	err = _UserCenterServiceRegisterOutCoder.DecodeIn(begoniaResTmp.([]byte), &out)
	if err != nil {
		panic(err)
	}

	F1 = out.F1

	F2 = out.F2

	return
}

func Login(username string, password string) (F1 bool, F2 int, err error) {
	var in _UserCenterServiceLoginIn
	in.F1 = username
	in.F2 = password

	b, err := _UserCenterServiceLoginInCoder.Encode(in)
	if err != nil {
		panic(err)
	}

	begoniaResTmp, err := _UserCenterServiceLogin(b)
	if err != nil {
		return
	}

	var out _UserCenterServiceLoginOut
	err = _UserCenterServiceLoginOutCoder.DecodeIn(begoniaResTmp.([]byte), &out)
	if err != nil {
		panic(err)
	}

	F1 = out.F1

	F2 = out.F2

	return
}

func ChangePassword(username string, newPassword string) (err error) {
	var in _UserCenterServiceChangePasswordIn
	in.F1 = username
	in.F2 = newPassword

	b, err := _UserCenterServiceChangePasswordInCoder.Encode(in)
	if err != nil {
		panic(err)
	}

	begoniaResTmp, err := _UserCenterServiceChangePassword(b)
	if err != nil {
		return
	}

	var out _UserCenterServiceChangePasswordOut
	err = _UserCenterServiceChangePasswordOutCoder.DecodeIn(begoniaResTmp.([]byte), &out)
	if err != nil {
		panic(err)
	}

	return
}
