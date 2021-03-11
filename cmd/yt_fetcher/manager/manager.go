// The CLI for data manage
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
			job.UpdateChannels(false)
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
				job.UpdateChannels(true)
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
