package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	Id string
	Mm string
}
func main() {
	var qr int
	var zh,mm string
	Allperson := make(map[string]Person)
	nonce:="wzxnb"
	key:="wzx"
	fp, _ := os.OpenFile("user1.json",os.O_RDWR|os.O_APPEND,0644)
	defer fp.Close()
	bytes, _ := ioutil.ReadAll(fp)
	deal(string(bytes),key,nonce)
	json.Unmarshal(bytes, &Allperson)
	for {
		fmt.Println("登录输入1，注册输入2，退出输入3")
		fmt.Scan(&qr)
		if qr == 1 {
			dl(Allperson)
		}
		if qr == 2 {
			zhbo:=true
			for zhbo{
				zhbo=false
				zh, mm = sr()
				for _,value:= range Allperson{
					if value.Id==zh{
						zhbo=true
						fmt.Println("账号已存在，请重新输入：")
					}
				}
			}

			Allperson[zh]=Person{zh,mm}
			fmt.Println(Allperson)
			fmt.Println("数据已保存，退出可实现注册")
		}
		if qr == 3 {
			if (len(zh) > 0 && len(mm) > 0) {
				os.Truncate("user.json",0)
				fp.Seek(0,0)
				h,_:=json.Marshal(Allperson)
				g:=add(string(h),key,nonce)
				fp.WriteString(g)
				fmt.Println("成功注册")
				break
			}
		}
	}

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
func dl(mp map[string]Person)int{
	var zh,mm string
	zhbo:=false
	mmbo:=false
	for !(zhbo&&mmbo) {
		zh, mm = sr()
		for _, value := range mp {
			if value.Id == zh {
				zhbo = true
				if value.Mm==mm{
					mmbo = true
				}
			}
		}
		if !zhbo{
			fmt.Println("账号错误")
		}else if !mmbo{
			fmt.Println("密码错误")
		}
	}
	fmt.Println("登录成功")
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

	block, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(block)

	plaintext, _ := aesgcm.Open(nil, nonce, ciphertext, nil)


	return string(plaintext)
}
