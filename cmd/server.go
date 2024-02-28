package cmd

import (
	"bytes"
	"fmt"
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
		if err := doc.ControlReader(conn, nil); err != nil {
			log.Fatal(err)
		}
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
		fmt.Println(err)
		os.Exit(1)
	}

	s := SockAddr
	if err := doc.ControlReader(bytes.NewBufferString(s), nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ov, err := oviewer.NewOviewer(doc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ov.SetConfig(config)

	l, err := net.Listen("unix", SockAddr)
	if err != nil {
		fmt.Printf("listen error: %s", err.Error())
		os.Exit(1)
		log.Fatal("listen error:", err)
	}
	defer l.Close()

	go receive(ov, l)

	if err := ov.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(serverCmd)

	config = oviewer.NewConfig()
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config `file` (default is $XDG_CONFIG_HOME/ov/config.yaml)")

	rootCmd.PersistentFlags().StringVarP(&SockAddr, "socket", "p", SockAddr, "socket path ")

	// Config.General
	serverCmd.PersistentFlags().IntP("tab-width", "x", 8, "tab stop width")
	_ = viper.BindPFlag("general.TabWidth", serverCmd.PersistentFlags().Lookup("tab-width"))

	serverCmd.PersistentFlags().IntP("header", "H", 0, "number of header rows to fix")
	_ = viper.BindPFlag("general.Header", serverCmd.PersistentFlags().Lookup("header"))

	serverCmd.PersistentFlags().BoolP("alternate-rows", "C", false, "color to alternate rows")
	_ = viper.BindPFlag("general.AlternateRows", serverCmd.PersistentFlags().Lookup("alternate-rows"))

	serverCmd.PersistentFlags().BoolP("column-mode", "c", false, "column mode")
	_ = viper.BindPFlag("general.ColumnMode", serverCmd.PersistentFlags().Lookup("column-mode"))

	serverCmd.PersistentFlags().BoolP("line-number", "n", false, "line number")
	_ = viper.BindPFlag("general.LineNumMode", serverCmd.PersistentFlags().Lookup("line-number"))

	serverCmd.PersistentFlags().BoolP("wrap", "w", true, "wrap mode")
	_ = viper.BindPFlag("general.WrapMode", serverCmd.PersistentFlags().Lookup("wrap"))

	serverCmd.PersistentFlags().StringP("column-delimiter", "d", ",", "column delimiter")
	_ = viper.BindPFlag("general.ColumnDelimiter", serverCmd.PersistentFlags().Lookup("column-delimiter"))

	rootCmd.PersistentFlags().BoolP("follow-mode", "f", false, "follow mode")
	_ = viper.BindPFlag("general.FollowMode", rootCmd.PersistentFlags().Lookup("follow-mode"))

	rootCmd.PersistentFlags().BoolP("follow-all", "A", false, "follow all")
	_ = viper.BindPFlag("general.FollowAll", rootCmd.PersistentFlags().Lookup("follow-all"))

	// Config
	serverCmd.PersistentFlags().BoolP("disable-mouse", "", false, "disable mouse support")
	_ = viper.BindPFlag("DisableMouse", serverCmd.PersistentFlags().Lookup("disable-mouse"))

	serverCmd.PersistentFlags().BoolP("exit-write", "X", false, "output the current screen when exiting")
	_ = viper.BindPFlag("AfterWrite", serverCmd.PersistentFlags().Lookup("exit-write"))

	serverCmd.PersistentFlags().BoolP("quit-if-one-screen", "F", false, "quit if the output fits on one screen")
	_ = viper.BindPFlag("QuitSmall", serverCmd.PersistentFlags().Lookup("quit-if-one-screen"))

	serverCmd.PersistentFlags().BoolP("case-sensitive", "i", false, "case-sensitive in search")
	_ = viper.BindPFlag("CaseSensitive", serverCmd.PersistentFlags().Lookup("case-sensitive"))

	serverCmd.PersistentFlags().BoolP("debug", "", false, "debug mode")
	_ = viper.BindPFlag("Debug", serverCmd.PersistentFlags().Lookup("debug"))
}
