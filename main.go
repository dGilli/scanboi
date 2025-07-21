package main

import (
    "fmt"
    "os"
    "os/exec"
)

func checkClamavUpdate() error {
    fmt.Println("Checking ClamAV database...")
    cmd := exec.Command("freshclam", "--quiet")
    err := cmd.Run()
    if err != nil {
        return fmt.Errorf("ClamAV database is outdated. Please rebuild the image to update libraries.")
    }
    fmt.Println("ClamAV database is up to date.")
    return nil
}

func clamavScan(directory string) {
    fmt.Printf("Running ClamAV scan on %s...\n", directory)
    cmd := exec.Command("clamscan", "-r", "--bell", "-i", directory)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
}

func rkhunterCheck() {
    fmt.Println("Running rkhunter check...")
    cmd := exec.Command("rkhunter", "--check")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <directory_to_scan>")
        os.Exit(1)
    }

    directory := os.Args[1]

    if err := checkClamavUpdate(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    clamavScan(directory)
    rkhunterCheck()
}

