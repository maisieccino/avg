package main

import (
	"context"
	"fmt"
	"github.com/mbellgb/avg/pkg/server"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"os"
	"strconv"
	"time"
)

var meanCmd = &cobra.Command{
	Use:   "mean [VALUES...]",
	Short: "Calculate mean average of values",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Calculating average for ", args)
		url := cmd.Flags().Lookup("url").Value.String()
		conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithTimeout(time.Second*10))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error connecting: %v\n", err)
			os.Exit(1)
		}
		svc := server.NewGRPCClient(conn)
		values := []int32{}
		for _, arg := range args {
			v, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Could not convert '%s' to int32\n", arg)
			}
			values = append(values, int32(v))
		}
		mean, err := svc.Mean(context.Background(), values)
		if err != nil {

			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(mean)
	},
}

func init() {
	rootCmd.AddCommand(meanCmd)
}
