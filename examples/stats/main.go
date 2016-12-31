package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/dseevr/go-oanda-streaming-api/client"
)

type PairStat struct {
	name  string
	count int64
	// TODO: extend with more stats
	//       e.g., lowest bid, highest ask, biggest spread, etc.
}

type PairStatsByCount []*PairStat

// Go is truly one of the worst mainstream languages
// I want to strangle whoever thought this was a good idea
func (slice PairStatsByCount) Len() int {
	return len(slice)
}
func (slice PairStatsByCount) Less(i, j int) bool {
	return slice[i].count > slice[j].count
}
func (slice PairStatsByCount) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func main() {
	pairStats := map[string]*PairStat{}
	totalTicks := int64(0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		tick := &client.Tick{}

		err := json.Unmarshal(scanner.Bytes(), tick)
		if err != nil {
			log.Fatalln(err)
		}

		// don't assume heartbeats were removed
		if !tick.IsTradeable() {
			continue
		}

		symbol := tick.Symbol()

		ps, found := pairStats[symbol]
		if !found {
			ps = &PairStat{name: symbol}
			pairStats[symbol] = ps
		}

		ps.count++
		totalTicks++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	//
	// output stats
	//
	fmt.Printf("Total ticks: %d\n", totalTicks)
	fmt.Printf("Unique pairs: %d\n", len(pairStats))
	fmt.Println("")

	//
	// show pair names alphabetically
	//
	names := []string{}
	for name, _ := range pairStats {
		names = append(names, name)
	}

	sort.Strings(names)

	fmt.Println("All pairs, sorted alphabetically:")
	for _, name := range names {
		fmt.Println(name)
	}
	fmt.Println("")

	//
	// show pairs by tick count
	//
	byTickCount := PairStatsByCount{}
	for _, value := range pairStats {
		byTickCount = append(byTickCount, value)
	}

	sort.Sort(byTickCount)

	for _, value := range byTickCount {
		fmt.Printf("%s: %d\n", value.name, value.count)
	}
}
