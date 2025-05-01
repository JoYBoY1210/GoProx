package proxy

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

const proxy = "http://172.31.2.3:8080"
const HTTPSproxy = "https://172.31.2.3:8080"

func Enable() {
	cmd := exec.Command("netsh", "winhttp", "set", "proxy", proxy)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error enabling proxy:", err)
		return
	}
	setGitproxy()
	setNpmproxy()
	fmt.Println("Proxy enabled:", proxy)
	time.Sleep(2 * time.Second)

}

func Disable() {
	cmd := exec.Command("netsh", "winhttp", "reset", "proxy")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error disabling proxy:", err)
		return
	}
	removeGitProxy()
	removeNpmProxy()
	fmt.Println("Proxy disabled")
}

func Status() {
	cmd := exec.Command("netsh", "winhttp", "show", "proxy")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error getting proxy status:", err)
		return
	}
	fmt.Println("Proxy status:", string(output))
	GitStatus()
	NpmStatus()
}

func GitStatus() {
	cmd := exec.Command("git", "config", "--global", "--get", "http.proxy")
	httpOut, _ := cmd.Output()

	cmd = exec.Command("git", "config", "--global", "--get", "https.proxy")
	httpsOut, _ := cmd.Output()

	if len(httpOut) == 0 && len(httpsOut) == 0 {
		fmt.Println("Git proxy: not set")
	} else {
		fmt.Println("Git proxy settings:")
		if len(httpOut) > 0 {
			fmt.Println("  http.proxy:", strings.TrimSpace(string(httpOut)))
		}
		if len(httpsOut) > 0 {
			fmt.Println("  https.proxy:", strings.TrimSpace(string(httpsOut)))
		}
	}
}

func NpmStatus() {
	cmd := exec.Command("npm", "config", "get", "proxy")
	proxyOut, _ := cmd.Output()

	cmd = exec.Command("npm", "config", "get", "https-proxy")
	httpsOut, _ := cmd.Output()

	if strings.Contains(string(proxyOut), "null") && strings.Contains(string(httpsOut), "null") {
		fmt.Println("NPM proxy: not set")
	} else {
		fmt.Println("NPM proxy settings:")
		if !strings.Contains(string(proxyOut), "null") {
			fmt.Println("  proxy:", strings.TrimSpace(string(proxyOut)))
		}
		if !strings.Contains(string(httpsOut), "null") {
			fmt.Println("  https-proxy:", strings.TrimSpace(string(httpsOut)))
		}
	}
}

func setGitproxy() {
	cmd := exec.Command("git", "config", "--global", "http.proxy", proxy)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to set git proxy: ", err)
		return
	}

	cmd = exec.Command("git", "config", "--global", "https.proxy", HTTPSproxy)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Failed to set git proxy: ", err)
		return
	}

	fmt.Println("Git proxy set")

}

func removeGitProxy() {
	cmd := exec.Command("git", "config", "--global", "--unset", "http.proxy")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to rmeove git proxy: ", err)
		return
	}

	cmd = exec.Command("git", "config", "--global", "--unset", "https.proxy")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Failed to rmeove git proxy: ", err)
		return
	}

	fmt.Println("Git proxy removed")
}

func setNpmproxy() {
	cmd := exec.Command("npm", "config", "set", "proxy", proxy)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Faield to set npm proxy")
		return
	}

	cmd = exec.Command("npm", "config", "set", "https-proxy", HTTPSproxy)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Faield to set npm proxy")
		return
	}

	fmt.Println("NPM proxy set")
}
func removeNpmProxy() {
	cmd := exec.Command("npm", "config", "delete", "proxy")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to delete npm proxy:", err)
	}
	cmd = exec.Command("npm", "config", "delete", "https-proxy")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Failed to delete npm proxy:", err)
	}
	fmt.Println("npm proxy removed.")
}
