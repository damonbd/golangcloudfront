package cloudwatchcontroller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/gorilla/mux"
)

var tpl *template.Template
var svc *cloudwatchlogs.CloudWatchLogs

func init() {
	createSession()
}

func createSession() {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	svc = cloudwatchlogs.New(sess)
}

//AddRoutes ...
func AddRoutes(r *mux.Router, mainTpl *template.Template) {
	tpl = mainTpl

	r = r.PathPrefix("/cloudwatch").Subrouter()

	r.HandleFunc("/loggroups", LogGroups)
	r.HandleFunc("/logstreams", LogStreams)
}

//LogGroups Landing page for log groups
func LogGroups(w http.ResponseWriter, r *http.Request) {

	var groupInput cloudwatchlogs.DescribeLogGroupsInput

	result, err := svc.DescribeLogGroups(&groupInput)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(result)
	}

	//tpl.ExecuteTemplate(w, "testsubroute.html", nil)
}

//LogStreams Landing page for log streams after selecting a log group
func LogStreams(w http.ResponseWriter, r *http.Request) {

	var streamInput cloudwatchlogs.DescribeLogStreamsInput
	streamInput.LogGroupName = aws.String("testloggroup")

	result, err := svc.DescribeLogStreams(&streamInput)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(result)
	}

	//tpl.ExecuteTemplate(w, "testsubroute.html", nil)
}
