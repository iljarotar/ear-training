package cmd

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// tableCmd represents the table command
var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "print sight singing table",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		min, _ := cmd.Flags().GetString("min")
		max, _ := cmd.Flags().GetString("max")
		length, _ := cmd.Flags().GetString("length")
		width, _ := cmd.Flags().GetString("width")
		title, _ := cmd.Flags().GetString("title")
		printTable(length, min, max, width, title)
	},
}

func init() {
	rootCmd.AddCommand(tableCmd)
	tableCmd.Flags().StringP("min", "m", "1", "set lowest note")
	tableCmd.Flags().StringP("max", "M", "8", "set highest note")
	tableCmd.Flags().StringP("length", "l", "100", "set table length")
	tableCmd.Flags().StringP("width", "w", "10", "set print width")
	tableCmd.Flags().StringP("title", "t", "sight singing table", "set table title")
}

func printTable(length, min, max, width, title string) {
	l, err := strconv.Atoi(length)
	if err != nil {
		fmt.Println("could not parse length argument; please provide an integer")
		return
	}

	low, err := strconv.Atoi(min)
	if err != nil {
		fmt.Println("could not parse min argument; please provide an integer")
		return
	}

	high, err := strconv.Atoi(max)
	if err != nil {
		fmt.Println("could not parse max argument; please provide an integer")
		return
	}

	w, err := strconv.Atoi(width)
	if err != nil {
		fmt.Println("could not parse print width; please provide an integer")
	}

	table := randomNotes(l, low, high)
	print(title, table, w)
}

func randomNotes(length, low, high int) []int {
	n := make([]int, length)

	rand.Seed(time.Now().UnixNano())
	for i := range n {
		r := rand.Intn(high-low+1) + low

		for i > 0 && n[i-1] == r {
			r = rand.Intn(high-low+1) + low
		}

		n[i] = r
	}

	return n
}

func print(title string, notes []int, width int) {
	for i := 0; i < width; i++ {
		fmt.Print("---")
	}
	fmt.Println()

	fmt.Printf("%s\n", title)

	for i := 0; i < width; i++ {
		fmt.Print("---")
	}
	fmt.Println()
	fmt.Println()

	c := 0
	for i := range notes {
		if c == width {
			fmt.Println()
			c = 0
		}

		space := " "
		if notes[i] < 10 && notes[i] >= 0 {
			space = "  "
		}

		fmt.Printf("%d%s", notes[i], space)
		c++
	}
	fmt.Println()
}
