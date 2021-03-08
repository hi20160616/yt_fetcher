// The CLI for data manage
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"github.com/hi20160616/yt_fetcher/internal/biz"
	"github.com/hi20160616/yt_fetcher/internal/data"
	"github.com/hi20160616/yt_fetcher/internal/pkg/db/mysql"
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
			addOrUpdateChannel(cid)
			fmt.Println("done.\n------------------------------------------------------")
			menu()
			continue
		case "2":
			fmt.Println(">> Update channels")
			log.Println("Start!")
			updateChannel(false)
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
				updateChannel(true)
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
			delChannel(cid)
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

func addOrUpdateChannel(id string) error {
	dc, err := mysql.NewDBCase()
	if err != nil {
		return err
	}
	defer dc.Close()

	c := &pb.Channel{Id: id}
	// get info from source
	c, err = data.GetChannelFromSource(c)
	// storage
	return mysql.InsertOrUpdateChannel(dc, c)
}

func delChannel(id string) error {
	fr := data.NewFetcherRepo()
	return fr.DelChannel(&pb.Channel{Id: id})
}

func updateChannel(greedy bool) error {
	fr := data.NewFetcherRepo()
	fc := biz.NewFetcherCase(fr)

	// 1. get cids from database
	cs := &pb.Channels{}
	cs, err := fc.GetChannels(cs)
	if err != nil {
		return err
	}
	// 2. for range cids, get vids from video pages where cid is
	fc.SetGreedy(greedy)
	err = fc.UpdateChannels(cs, fc.GetGreedy())
	return nil
}
