package main

import (
	"fmt"
	"os"
	"strings"
	"net/http"
	"io/ioutil"
	"strconv"
	"time"
)

func ShowHelp() {
	fmt.Println("----------------------")
	fmt.Println(" HTTP CHECK TOOL 1.0")
	fmt.Println("----------------------")

	fmt.Println("usage:\n")
	fmt.Println("http-check <url> code <http-code> - check http response code")
	fmt.Println("return:")
	fmt.Println("  0 - if match")
	fmt.Println("  1 - if return code is different\n")

	fmt.Println("http-check <url> substring <substring> - look for <substring> in response")
	fmt.Println("return:")
	fmt.Println("  0 - if match")
	fmt.Println("  1 - if not found")
}

func ArrayContains(arr []string, element string) bool {
	result := false

	for _, value := range arr {
		if value == element {
			result = true
			break
		}
	}

	return result
}

func ResultOK() {
	fmt.Print("0")
}

func ResultFailure() {
	fmt.Print("1")
}

func main() {
	if len(os.Args) == 4 {
		url := os.Args[1] // <url>
		cmd := os.Args[2] // <cmd>
		data := os.Args[3] // <substring\code>

		CODE_COMMAND := "code"
		SUBSTRING_COMMAND := "substring"
		DIAG_COMMAND := "diag"

		diagMode := false

		if cmd == DIAG_COMMAND {
			diagMode = true
		}

		if ! ArrayContains([]string{CODE_COMMAND, SUBSTRING_COMMAND, DIAG_COMMAND}, cmd) {
			fmt.Printf("unknown command '%s'", cmd)
			os.Exit(1)
		}

		if (! strings.HasPrefix(url, "http://")) && (! strings.HasPrefix(url, "https://"))  {
			url = "http://" + url
		}

		timeout := time.Duration(2 * time.Second)

		client := http.Client{
			Timeout: timeout,
		}

		response, err := client.Get(url)

		if err != nil {
			if diagMode {
				fmt.Printf("%s", err)

			} else {
				ResultFailure()
			}

			os.Exit(1)

		} else {
			defer response.Body.Close()

			if cmd == CODE_COMMAND {
				code, err := strconv.Atoi(data)

				if err != nil {
					if diagMode {
						fmt.Printf("invalid code value: '%s'", data)

					} else {
						ResultFailure()
					}

					os.Exit(1)
				}

				if response.StatusCode == code {
					if diagMode {


					} else {
						ResultOK()
					}

				} else {
					ResultFailure()
				}

			} else if cmd == SUBSTRING_COMMAND || diagMode {
				contents, err := ioutil.ReadAll(response.Body)

				if err != nil {
					if diagMode {
						fmt.Printf("error: %s", err)

					} else {
						ResultFailure()
					}

					os.Exit(1)
				}

				if diagMode {
					fmt.Printf("response code: %d\n", response.StatusCode)

					fmt.Println("start of response (128):")
					fmt.Println(string(contents)[:128])
				} else {
					if strings.Contains(string(contents), data) {
						ResultOK()

					} else {
						ResultFailure()
					}
				}
			}
		}

	} else {
		ShowHelp()
	}
}