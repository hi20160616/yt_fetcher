// The CLI for data manage
// TODO: remote control, edit channel individual
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hi20160616/yt_fetcher/internal/job"
)

func menu() {
	fmt.Println("[!] Options here:")
	fmt.Println("[1] Add or update Channel by id")
	fmt.Println("[2] Update Channels")
	fmt.Println("[3] Update Channels Greedy!")
	fmt.Println("[4] Delete Channel by id")
	fmt.Println("[5] Update Thumbnails")
	fmt.Println("[q] Quit")
	fmt.Printf(">> Input your choice: ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cid := ""
	menu()
	for scanner.Scan() {
		switch scanner.Text() {
		case "1":
			fmt.Printf(">> Input Channel id: ")
			scanner.Scan()
			cid = scanner.Text()
			fmt.Printf(">> ADD or UPDATE channel by id: `%s`\n", cid)
			job.AddOrUpdateChannel(cid)
			fmt.Println("done.\n------------------------------------------------------")
			menu()
			continue
		case "2":
			fmt.Println(">> Update channels")
			log.Println("Start!")
			err := job.UpdateChannels(false)
			if err != nil {
				log.Println(err)
				continue
			}
			log.Println("done.")
			fmt.Println("------------------------------------------------------")
			menu()
			continue
		case "3":
			fmt.Println(">> Update channels Greedy!")
			fmt.Println(">> This should be slowly but fully rewrite the records!")
			fmt.Println(">> Are you sure?![y/N]")
			fmt.Printf(">> ")
			scanner.Scan()
			answer := scanner.Text()
			if strings.ToLower(strings.TrimSpace(answer)) == "y" {
				log.Println("Greedy Update Start!")
				err := job.UpdateChannels(true)
				if err != nil {
					fmt.Println(err)
					return
				}
				log.Println("done")
				fmt.Println("------------------------------------------------------")
			}
			menu()
			continue
		case "4":
			fmt.Printf(">> Input Channel id: ")
			scanner.Scan()
			cid = scanner.Text()
			fmt.Printf(">> DEL channel by id: `%s`\n", cid)
			job.DelChannel(cid)
			fmt.Println("done.\n------------------------------------------------------")
			menu()
			continue
		case "5":
			fmt.Println(">> Update Thumbnails")
			log.Println("Start!")
			err := job.UpdateThumbnails()
			if err != nil {
				log.Println(err)
				continue
			}
			log.Println("done.")
			fmt.Println("------------------------------------------------------")
			menu()
			continue
		case "q":
			fmt.Println("Bye!")
			return
		default:
			menu()
			continue
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
