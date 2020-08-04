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
		exit("Error: required flag 'prefix' missing.")
	}

	sess := session.Must(session.NewSession())
	client := ssm.New(sess, aws.NewConfig().WithRegion("eu-west-1"))

	input := &ssm.GetParametersByPathInput{
		Path:      prefix,
		Recursive: aws.Bool(true),
	}

	output, err := client.GetParametersByPath(input)
	if err != nil {
		exit(fmt.Sprintf("Error: unable to retrieve from parameter store - %s.", err.Error()))
	}

	fmt.Print(asKV(output.Parameters))
}

func asKV(params []*ssm.Parameter) string {
	builder := strings.Builder{}

	for _, param := range params {
		name := *param.Name
		value := *param.Value

		builder.WriteString(fmt.Sprintf("%s=%s\n", clean(name), value))
	}

	return builder.String()
}

func clean(s string) string {
	r := strings.NewReplacer(".", "_", "/", "_")
	return r.Replace(s[1:]) // strip leading '/'
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
