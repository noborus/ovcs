package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"github.com/noborus/ov/oviewer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	config oviewer.Config
)

// serverCmd represents the base command when called without any subcommands
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Terminal pager server",
	Long: `Terminal pager server.
Wait for unix domain socket on startup.`,
	Run: func(cmd *cobra.Command, args []string) {
		server()
	},
}

func receive(ov *oviewer.Root, l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		doc, err := oviewer.NewDocument()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%#v %#v\n", ov.Config.General.AlternateRows, doc.AlternateRows)
		doc.ReadAll(conn)
		ov.AddDocument(doc)
	}
}

func server() {
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	doc, err := oviewer.NewDocument()
	if err != nil {
		log.Fatal(err)
	}

	s := SockAddr
	doc.ReadAll(ioutil.NopCloser(bytes.NewBufferString(s)))

	ov, err := oviewer.NewOviewer(doc)
	if err != nil {
		log.Fatal(err)
	}
	ov.SetConfig(config)

	l, err := net.Listen("unix", SockAddr)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer l.Close()

	go receive(ov, l)

	if err := ov.Run(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(serverCmd)

	config = oviewer.NewConfig()
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ov.yaml)")

	rootCmd.PersistentFlags().StringVarP(&SockAddr, "socket", "p", SockAddr, "socket path ")

	serverCmd.PersistentFlags().BoolP("wrap", "w", true, "wrap mode")
	_ = viper.BindPFlag("general.Wrap", serverCmd.PersistentFlags().Lookup("wrap"))

	serverCmd.PersistentFlags().IntP("tab-width", "x", 8, "tab stop width")
	_ = viper.BindPFlag("general.TabWidth", serverCmd.PersistentFlags().Lookup("tab-width"))

	serverCmd.PersistentFlags().IntP("header", "H", 0, "number of header rows to fix")
	_ = viper.BindPFlag("general.Header", serverCmd.PersistentFlags().Lookup("header"))

	serverCmd.PersistentFlags().BoolP("disable-mouse", "", false, "disable mouse support")
	_ = viper.BindPFlag("general.DisableMouse", serverCmd.PersistentFlags().Lookup("disable-mouse"))

	serverCmd.PersistentFlags().BoolP("exit-write", "X", false, "output the current screen when exiting")
	_ = viper.BindPFlag("general.ExitWrite", serverCmd.PersistentFlags().Lookup("exit-write"))

	serverCmd.PersistentFlags().BoolP("quit-if-one-screen", "F", false, "quit if the output fits on one screen")
	_ = viper.BindPFlag("general.QuitSmall", serverCmd.PersistentFlags().Lookup("quit-if-one-screen"))

	serverCmd.PersistentFlags().BoolP("case-sensitive", "i", false, "case-sensitive in search")
	_ = viper.BindPFlag("general.CaseSensitive", serverCmd.PersistentFlags().Lookup("case-sensitive"))

	serverCmd.PersistentFlags().BoolP("alternate-rows", "C", false, "color to alternate rows")
	_ = viper.BindPFlag("general.AlternateRows", serverCmd.PersistentFlags().Lookup("alternate-rows"))

	serverCmd.PersistentFlags().BoolP("column-mode", "c", false, "column mode")
	_ = viper.BindPFlag("general.ColumnMode", serverCmd.PersistentFlags().Lookup("column-mode"))

	serverCmd.PersistentFlags().StringP("column-delimiter", "d", ",", "column delimiter")
	_ = viper.BindPFlag("general.ColumnDelimiter", serverCmd.PersistentFlags().Lookup("column-delimiter"))

	serverCmd.PersistentFlags().BoolP("line-number", "n", false, "line number")
	_ = viper.BindPFlag("general.LineNumMode", serverCmd.PersistentFlags().Lookup("line-number"))
}
