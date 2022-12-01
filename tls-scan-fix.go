package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    // List all json files in the current directory
    files, err := filepath.Glob("*.json")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Loop over all the files
    for _, file := range files {
        // Open the input JSON file
        f, err := os.Open(file)
        if err != nil {
            fmt.Println(err)
            continue
        }
        defer f.Close()

        // Create a new scanner to read the input file
        scanner := bufio.NewScanner(f)

        // Read each line of the input file
        for scanner.Scan() {
            // Unmarshal the JSON data
            var data map[string]interface{}
            err := json.Unmarshal(scanner.Bytes(), &data)
            if err != nil {
                fmt.Println(err)
                continue
            }

            // Get the IP address from the data
            ip, ok := data["ip"].(string)
            if !ok {
                fmt.Println("IP not found")
                continue
            }

            // Get the certificate chain from the data
            certificateChain, ok := data["certificateChain"].([]interface{})
            if !ok {
                fmt.Println("certificateChain not found")
                continue
            }

            // Get the first certificate from the chain
            firstCertificate, ok := certificateChain[0].(map[string]interface{})
            if !ok {
                fmt.Println("first certificate not found")
                continue
            }

            // Get the subjectCN field from the first certificate
            subjectCN, ok := firstCertificate["subjectCN"].(string)
            if !ok {
                fmt.Println("subjectCN not found")
                continue
            }

            // Print the domain name and IP address to the console
            fmt.Println(subjectCN + " - " + ip)
        }

        // Check for errors while reading the input file
        if err := scanner.Err(); err != nil {
            fmt.Println(err)
        }
    }
}
