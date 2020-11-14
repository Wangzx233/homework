package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var str string
	var qr int
	var zh,mm string
	nonce := "37b8e8a308c354048d245f6d"
	key := "AES256Key-32Characters1234567890"

	//var mp = make(map[string]string)
	fp, _ := os.OpenFile("user.txt",os.O_RDWR|os.O_APPEND,0644)
	bytes, _ := ioutil.ReadAll(fp)
	str=string(bytes)
	str=deal(str,key,nonce)
	//json.Unmarshal(bytes, &p1)
	sz :=strings.Split(str,":")
	mp:=stom(sz)
	for {
		fmt.Println("登录输入1，注册输入2，退出输入3")
		fmt.Scan(&qr)
		if qr == 1 {
			dl(mp)
		}
		if qr == 2 {
			zh, mm = sr()
			zh=add(zh,key,nonce)
			mm=add(mm,key,nonce)
			fmt.Println("数据已保存，退出可实现注册")
		}
		if qr == 3 {
			if (len(zh) > 0 && len(mm) > 0) {


				fp.WriteString(zh)
				fp.WriteString(":")
				fp.WriteString(mm)
				fp.WriteString(":")
				fmt.Println("成功注册")
				break
			}
		}
	}
	defer fp.Close()
}
func stom(sz []string)map[string]string{
	long:=len(sz)
	mp:=make(map[string]string)
	for i:=0;i<long;i++ {
		if i%2==0{
			mp["id"]=sz[i]
		}else {
			mp["mm"]=sz[i]
		}
	}
	return mp
}
func sr()(string,string)  {
	var zh,mm string
	for {
		fmt.Println("请输入账号密码：")
		fmt.Scan(&zh)
		fmt.Scan(&mm)
		if len(mm)>=6{break}else {fmt.Println("密码长度不能小于6")}
	}
	return zh,mm
}
func dl(mp map[string]string)int{
	var zh,mm string
	bo:=false
	boz:=false
	bozm:=false
	for !bo {
		zh,mm=sr()
		for key, _ := range mp {
			if bozm {
				if mp[key] == mm {
					bo = true
					break
				}
				boz = true
				bozm = false
			} else {
				if key == "id" {
					if mp[key] == zh {
						bozm = true
					}
				}
			}
		}
		if (bo) {
			fmt.Println("账号密码正确")
		} else {
			if (boz) {
				fmt.Println("账号正确，密码错误")
			} else {
				fmt.Println("账号不存在")
			}
		}
	}
	return 2
}
func add(src, k, n string)string {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte(k)
	plaintext := []byte(src)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	nonce, _ := hex.DecodeString(n)

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	return fmt.Sprintf("%x", ciphertext)
}

func deal(src, k, n string) string {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte(k)
	ciphertext, _ := hex.DecodeString(src)

	nonce, _ := hex.DecodeString(n)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)


	return string(plaintext)
}
