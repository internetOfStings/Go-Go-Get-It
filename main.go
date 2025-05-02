package main

import (
	"fmt"
	"os"
	"os/exec"
	"log"
	"strings"
)

func main() {

	// Programs to be used.
	audio := "mpv"
	video := "mpv"
	image := "sxiv"
	office := "libreoffice"
	pdf := "evince"
	web := "firefox"

	// Command flags.
	audioFlags := "-no-video"
	videoFlags := "-fullscreen"
	imageFlags := "-f"
	officeFlags := ""
	pdfFlags := "--fullscreen"
	webFlags := "" 


	fileFormats := map[string]string{
		"mp3":  "audio",
		"wav":  "audio",
		"aac":  "audio",
		"flac": "audio",
		"ogg":  "audio",
		"wma":  "audio",
		"aiff": "audio",
		"m4a":  "audio",
		"ape":  "audio",
		"mp4":  "video",
		"avi":  "video",
		"mkv":  "video",
		"mov":  "video",
		"wmv":  "video",
		"flv":  "video",
		"webm": "video",
		"3gp":  "video",
		"3g2":  "video",
		"mpeg": "video",
		"ogv":  "video",
		"jpeg": "image",
		"jpg":  "image",
		"png":  "image",
		"ico":  "image",
		"doc":  "office",
		"docx": "office",
		"docm": "office",
		"xls":  "office",
		"xlsx": "office",
		"xlsm": "office",
		"ppt":  "office",
		"pptx": "office",
		"pps":  "office",
		"ppsx": "office",
		"pptm": "office",
		"ppsm": "office",
		"odt":  "office",
		"ods":  "office",
		"odp":  "office",
		"html": "web",
		"htm":  "web",
		"pdf":	"pdf",
	}

	// Check if the correct number of arguments is provided.
	if len(os.Args) < 2 {
		fmt.Println("Enter an argumet")
		return
	} else if len(os.Args) > 2 {
		fmt.Println("Enter less argumets: Only 1 is allowed.")
		return
	} else {

		argument := os.Args[1]

		fileExt := fileExtension(argument)

		fileFormat := fileFormats[fileExt]

		switch fileFormat {
		case "audio":
			runCommand(audio, audioFlags, argument)
		case "video":
			runCommand(video, videoFlags, argument)
		case "image":
			runCommand(image, imageFlags, argument)
		case "office":
			runCommand(office, officeFlags, argument)
		case "web":
			runCommand(web, webFlags, argument)
		case "pdf":
			runCommand(pdf, pdfFlags, argument)
		default:
			fmt.Println("Invalid File Extensions")
		}

	}
}

func runCommand(command, argument1, argument2 string) {
	cmd := exec.Command(command, argument1, argument2)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func fileExtension(file string) string { // check for files that have no extension?
    s := strings.Split(file, ".")  
    return strings.ToLower(s[len(s)-1])
}