// Copyright © 2018 Nicolas Mailhot <nim@fedoraproject.org>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import "os"
import "net"
import "strings"
import "fmt"
import "bytes"
import "bufio"
import "io"

func main() {
  toInstall := strings.Join(os.Args[1:], " ")
  if toInstall == "" {
    fmt.Printf("Usage: mock-install <something>\n")
    os.Exit(0)
  }
  socketPath := os.Getenv("PM_REQUEST_SOCKET")
  if socketPath == "" {
    fmt.Printf("PM_REQUEST_SOCKET environment variable not set, doing nothing\n")
    os.Exit(0)
  }
  sock, err := net.Dial("unix", socketPath)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error opening PM_REQUEST_SOCKET (%v): %v\n", socketPath, err)
    os.Exit(1)
  }
  defer sock.Close()
  _, err = sock.Write([]byte("install " + toInstall + "\n"))
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error writing PM_REQUEST_SOCKET (%v): %v\n", socketPath, err)
    os.Exit(1)
  }
  mockResult   := ""
  mockResponse := make([]byte, 262144)
  for {
    n, err := sock.Read(mockResponse)
    if err == io.EOF {
      break
    }
    if err != nil {
      fmt.Fprintf(os.Stderr, "Error reading PM_REQUEST_SOCKET (%v): %+v\n", socketPath, err)
      os.Exit(1)
    }
    scanner := bufio.NewScanner(bytes.NewReader(mockResponse[:n]))
    if mockResult == "" {
      scanner.Scan()
      mockResult = scanner.Text()
    }
    for scanner.Scan() {
      fmt.Println(scanner.Text())
    }
  }
  if mockResult != "ok" {
    os.Exit(1)
  }
}
