package main

import (
	"fmt"
	opnsense "github.com/adminhead/gopnsense/pkg"
	"html"
	"os"
	"strings"
)

func main() {
	key := os.Getenv("OPNSENSE_KEY")
	secret := os.Getenv("OPNSENSE_SECRET")
	client := opnsense.NewAPIClient("https://localhost:10443", key, secret)
	client.SetTLSVerify(false)
	backups, err := client.Core().Backup().Backups()
	if err != nil {
		panic(err)
	}
	fmt.Println("Backups:")
	for _, backup := range backups.Items {
		fmt.Printf("Backup Time: %s\n", backup.Time)
		fmt.Printf("Backup ID: %s\n", backup.ID)
	}
	fmt.Println()

	diff, err := client.Core().Backup().Diff(backups.Items[0].ID, backups.Items[1].ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Diff:")
	fmt.Println(html.UnescapeString(strings.Join(diff.Items, "\n")))
	fmt.Println()
}
