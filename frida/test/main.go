package main

//
//import (
//	"bytes"
//	"fmt"
//	"github.com/frida/frida-go/frida"
//	"io"
//	"log"
//	"net/http"
//	"os"
//	"strings"
//)
//
//func exe(sign string, uid string) {
//	body := "appKey=we_sign_2.0&method=profile.call.lover.request&v=3.1&fm=json&tag=cp%2Fprofile&uid=" + uid + "&sign=" + sign + "&sessionid=3adef82e2316b068abbefe78c780e426"
//	request, err := http.NewRequest("POST", "https://sapi.y1s1.co/wecoreapi/was", bytes.NewBufferString(body))
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	request.Header.Set("uid", "40658163")
//	request.Header.Set("extra", "google Pixel 4")
//	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//	request.Header.Set("User-Agent", "Dalvik/2.1.0 (Linux; U; Android 10; google Pixel 4 Build/QQ3A.200805.001)")
//	request.Header.Set("token", "IdB85gsKXVNJqRyfgGEx/Mluq/EzpYfOGRAkOFVSbVv2auIgJozTUhldLUfLI+dhx5LQUudULTHnojmMx2nw3qXFlUc43arZDb53ZrasFvtiC0dJiLn3C/24F5NRyeSlZVCuqJfv+nZtBIqnbn/fYRFe1sdrGSGqrORbrovM9rM=")
//
//	// ... 其他的请求头设置
//
//	client := &http.Client{}
//
//	resp, err := client.Do(request)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	all, err := io.ReadAll(resp.Body)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	fmt.Println(string(all))
//
//}
//
//func main() {
//
//	device := frida.USBDevice()
//
//	attach, err := device.Attach("恋爱物语", nil)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	file, _ := os.ReadFile("/root/cyvk/github/PycharmProjects/pythonProject/frida/rpc.js")
//
//	createScript, err := attach.CreateScript(string(file))
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	createScript.On("message", func(message string) {
//		fmt.Println(message)
//	})
//
//	if err := createScript.Load(); err != nil {
//		fmt.Println("Error loading script:", err)
//		os.Exit(1)
//	}
//
//	dd, _ := os.ReadFile("/root/cyvk/ManDown/XChecks/test/uuid25")
//
//	for _, s2 := range strings.Split(string(dd), "\n") {
//		call := createScript.ExportsCall("getsign", s2)
//		fmt.Println(call)
//		s := call.(string)
//		exe(s, s2)
//	}
//
//}
