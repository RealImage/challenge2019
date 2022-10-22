package main

import (
	"flag"

	"github.com/c1pca/challenge2019/internal/app"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	pflag.String("input", "../../input/input.csv", "input path for flagname")
	pflag.String("partners", "../../input/partners.csv", "partners path for flagname")
	pflag.String("output", "../../output/", "output path for flagname")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	input := viper.GetString("input")
	partners := viper.GetString("partners")
	output := viper.GetString("output")

	app.FindPartnerWithMinCost(input, partners, output)
}
