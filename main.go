package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func main() {
	prefix := flag.String("prefix", "", "Parameter store prefix")
	flag.Parse()
	if *prefix == "" {
		exitWithError("Error: required flag 'prefix' missing.")
	}

	sess := session.Must(session.NewSession())
	client := ssm.New(sess, aws.NewConfig().WithRegion("eu-west-1"))

	input := &ssm.GetParametersByPathInput{
		Path:           prefix,
		Recursive:      aws.Bool(true),
		WithDecryption: aws.Bool(true),
	}

	output, err := client.GetParametersByPath(input)
	if err != nil {
		exitWithError(fmt.Sprintf("Error: unable to retrieve from parameter store - %s.", err.Error()))
	}

	fmt.Print(asKV(output.Parameters, *prefix))
}

func asKV(params []*ssm.Parameter, prefix string) string {
	builder := strings.Builder{}

	for _, param := range params {
		name := *param.Name
		value := *param.Value

		builder.WriteString(fmt.Sprintf("%s='%s'\n", cleanPrefix(name, prefix), cleanValue(value)))
	}

	return builder.String()
}

func cleanValue(s string) string {
	r := strings.NewReplacer("'", "\\'")
	return r.Replace(s)
}

func cleanPrefix(s, prefix string) string {
	r := strings.NewReplacer(prefix+"/", "", ".", "_", "/", "_", "-", "_")
	return r.Replace(s)
}

func exitWithError(msg string) {
	fmt.Fprintf(os.Stderr, msg)
	os.Exit(1)
}
