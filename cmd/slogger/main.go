package main

import (
	"bufio"
	"flag"
	"os"
	"strings"
	"fmt"

	sl "github.com/eshu0/simplelogger"
	//kitlog "github.com/go-kit/kit/log"
)

func main() {

	// this is the dummy logger object
	//&sl.SimpleLogger{}
	filename := flag.String("filename", "slogger.log", "Filename out - defaults to slogger.log")
	session := flag.String("session", "123", "Session - defaults to 123")

	flag.Parse()

	log := sl.NewSimpleLogger(*filename, *session)

	// lets open a flie log using the session
	log.OpenAllChannels()

	//defer the close till the shell has closed
	defer log.CloseAllChannels()

	reader := bufio.NewReader(os.Stdin)

	for {
		// read the string input
		text, readerr := reader.ReadString('\n')

		if readerr != nil {
			log.LogDebugf("main()", "Reading input has provided following err '%s'", readerr.Error())
			break
			// break out for loop
		}

		log.LogDebugf("main()", "input was: '%s'", text)

		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		if strings.ToLower(text) == "quit" || strings.ToLower(text) == "exit" {
			fmt.Println("bye bye")
			break
		} else {

			inputs := strings.Split(text, " ")

			if(len(inputs) == 1 ){

				if(strings.ToLower(inputs[0]) == "sessionids"){
					sessionids := log.GetSessionIDs()
					for SessionID := range sessionids {
						fmt.Println(fmt.Sprintf("'%s'", SessionID))
						log.LogInfof("main()", "Get Session ID: '%s'", SessionID)
					}
				}else if (strings.ToLower(inputs[0]) == "sessions") {
						for _,channel := range log.GetChannels() {
							fmt.Println(fmt.Sprintf("'%s'", channel.GetSessionID()))
							log.LogInfof("main()", "Get Session: '%s'", channel.GetSessionID())
						}
				}

			}else{

				if(len(inputs) >= 2){
						fmt.Println(fmt.Sprintf("Logged to '%s' with %s", inputs[0], inputs[1]))
						if(strings.ToLower(inputs[0]) == "debug"){
							log.LogDebugf("main()", "'%s'", inputs[1])
						}else if(strings.ToLower(inputs[0]) == "info"){
							log.LogInfof("main()", "'%s'", inputs[1])
						}else if(strings.ToLower(inputs[0]) == "error"){
							log.LogError("main()", "'%s'", inputs[1])
						}else if(strings.ToLower(inputs[0]) == "warn"){
							log.LogWarnf("main()", "'%s'", inputs[1])
						}else if(strings.ToLower(inputs[0]) == "add" && strings.ToLower(inputs[1]) == "session"){
							//logger := kitlog.NewLogfmtLogger(f1)
							//logger = kitlog.With(logger, "session_id", inputs[2], "ts", kitlog.DefaultTimestampUTC)
							//log.AddLog(logger)
						}
				}else{
					fmt.Println(fmt.Sprintf("'%s' was split but only had %d inputs", text, len(inputs)))
					log.LogDebugf("main()", "'%s' was split but only had %d inputs", text, len(inputs))
				}
			}

		}
	}

}