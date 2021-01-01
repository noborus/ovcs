package cmd

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Terminal pager client",
	Long: `Terminal pager client.
Pipe to the server..`,
	Run: func(cmd *cobra.Command, args []string) {
		client()
	},
}

func client() {
	conn, err := net.Dial("unix", SockAddr)
	if err != nil {
		log.Printf("error: %v\n", err)
		return
	}
	defer conn.Close()
	var reader = bufio.NewReader(os.Stdin)

	for {
		buf, isPrefix, err := reader.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) || errors.Is(err, io.ErrClosedPipe) {
				break
			}
			log.Printf("error: %v\n", err)
			return
		}
		conn.Write(buf)
		if isPrefix {
			continue
		}
		conn.Write([]byte("\n"))
	}
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
