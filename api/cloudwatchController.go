package cloudwatchcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/gorilla/mux"
)

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
func AddRoutes(r *mux.Router) {

	r = r.PathPrefix("/cloudwatch").Subrouter()

	r.HandleFunc("/loggroups", LogGroups).Methods("POST")
	r.HandleFunc("/logstreams", LogStreams).Methods("POST")
}

//w *myResponseWriter
func parseJSON(w http.ResponseWriter, r *http.Request, lgsr *LogGroupSearchRequest) bool {
	if err := json.NewDecoder(r.Body).Decode(lgsr); err != nil {
		json.NewEncoder(w).Encode("Invalid JSON Request")
		return false
	}
	return true
}

//LogGroups Accepts Json
func LogGroups(w http.ResponseWriter, r *http.Request) {

	var lgsr LogGroupSearchRequest

	if !parseJSON(w, r, &lgsr) {
		return
	}

	fmt.Print(lgsr)
	fmt.Print(&lgsr)

	var groupInput cloudwatchlogs.DescribeLogGroupsInput
	// groupInput.Limit = lgsr.count
	// groupInput.LogGroupNamePrefix = lgsr.prefix

	result, err := svc.DescribeLogGroups(&groupInput)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(result)
	}

	//return result
}

//LogStreams Accepts Json
func LogStreams(w http.ResponseWriter, r *http.Request) {

	var streamInput cloudwatchlogs.DescribeLogStreamsInput
	streamInput.LogGroupName = aws.String("testloggroup")

	result, err := svc.DescribeLogStreams(&streamInput)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(result)
	}

	//return result
}

//LogGroupSearchRequest ...
type LogGroupSearchRequest struct {
	//this probably isnt right
	Count  int64
	Prefix string
}
