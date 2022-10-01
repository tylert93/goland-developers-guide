//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func getStats(servers []string, statuses map[string]int) {
	var (
		onlineServers, offlineServers, maintenanceServers, retiredServers int
	)
	
	fmt.Println("Number of servers:", len(servers))

	for _, status := range statuses {
		switch {
		case status == Online:
			onlineServers += 1
		case status == Offline:
			offlineServers += 1
		case status == Maintenance:
			maintenanceServers += 1
		case status == Retired:
			retiredServers += 1
		}
	}

	fmt.Println("Online servers:", onlineServers)
	fmt.Println("Offline servers:", offlineServers)
	fmt.Println("Maintenance servers:", maintenanceServers)
	fmt.Println("Retired servers:", retiredServers)
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverStatuses := make(map[string]int)

	for _, server := range servers {
		serverStatuses[server] = Online
	}

	getStats(servers, serverStatuses)

	fmt.Println("\n----- UPDATING SERVERS -----")
	serverStatuses["darkstar"] = Retired
	serverStatuses["aiur"] = Offline

	getStats(servers, serverStatuses)

	fmt.Println("\n----- UPDATING SERVERS -----")
	for server := range serverStatuses {
		serverStatuses[server] = Maintenance
	}

	getStats(servers, serverStatuses)
}
