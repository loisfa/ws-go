package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request: ", r.URL)
	filePath := r.URL.Path[1:]
	http.ServeFile(w, r, filePath)
}

var (
	DEFAULT_PORT = 8080
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve the content on the server",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt64("port")
		http.HandleFunc("/", ServeFile)
		fmt.Println("Server running on port ", port)
		http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	},
}

func Execute() {
	if err := serveCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	serveCmd.Flags().Int64P("port", "p", 8080, "port on which the server runs")
}
