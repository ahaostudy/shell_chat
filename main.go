package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"shell_chat/config"
	"shell_chat/entity"
	"shell_chat/utils"
	"strings"
	"syscall"
	"time"
)

func main() {
	fmt.Print(utils.ColorString("Welcome to chat shell!\n\n", config.Cfg.Colors.Yellow))
	config.InitConfig()
	var messages entity.Messages
	var ipt string
	var question string
	for {
		// input
		ipt = utils.ReadLine(utils.ColorString("user", config.Cfg.Colors.Purple) + ": ")
		if len(ipt) == 0 {
			continue
		}
		// clear
		if ipt == "clear" {
			messages = entity.Messages{}
			fmt.Print(utils.ColorString("system", config.Cfg.Colors.Blue) + ": messages are emptied\n\n")
			continue
		}
		// multi-line input
		question = ipt
		if len(ipt) >= len(config.Cfg.Prefix) && ipt[:len(config.Cfg.Prefix)] == config.Cfg.Prefix {
			question = question[len(config.Cfg.Prefix):] + "\n"
			for {
				ipt = utils.ReadLine("")
				if len(ipt) >= len(config.Cfg.Suffix) && ipt[len(ipt)-len(config.Cfg.Suffix):] == config.Cfg.Suffix {
					question += ipt[:len(ipt)-len(config.Cfg.Suffix)]
					break
				}
				question += ipt + "\n"
			}
		}
		messages.AddUser(question)

		// request
		start := time.Now()
		body, err := utils.ChatStream(messages)

		if err != nil {
			fmt.Println(utils.ColorString("error", config.Cfg.Colors.Red)+": request openai chat error:", err.Error())
			return
		}
		fmt.Println(utils.ColorString("time", config.Cfg.Colors.Blue)+":", time.Now().Sub(start).String())

		c := make(chan os.Signal, 1)
		o := make(chan bool)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		done := make(chan bool)

		// read stream
		go func() {
			answer := ""
			scanner := bufio.NewScanner(body)
			fmt.Print(utils.ColorString("ai", config.Cfg.Colors.Green) + ": ")
			for scanner.Scan() {
				select {
				case <-done:
					return
				default:
					fields := strings.SplitN(scanner.Text(), ": ", 2)
					if len(fields) == 2 {
						if fields[1] == "[DONE]" {
							break
						}
						var d entity.OpenaiChatStreamResponse
						err := json.Unmarshal([]byte(fields[1]), &d)
						if err != nil {
							fmt.Println(utils.ColorString("error", config.Cfg.Colors.Red)+": json un marshal error:", err.Error())
							break
						}
						answer += d.Choices[0].Delta.Content
						fmt.Print(d.Choices[0].Delta.Content)
					}
				}
			}
			if err := scanner.Err(); err != nil {
				fmt.Println(utils.ColorString("error", config.Cfg.Colors.Red)+": scanner error:", err.Error())
			}

			// add messages
			_ = body.Close()
			messages.AddAssistant(answer)
			o <- true
		}()

		// listen ctrl+c
		go func() {
			for {
				select {
				case <-c:
					done <- true
					o <- true
				case <-o:
					return
				}
			}
		}()

		<-o
		fmt.Print("\n\n")
	}
}
