package zenformation

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

const (
	ceType    = "com.zendesk.tag.create"
	ceSubject = "New Zendesk Comprehend"
)

// Transformation creates an interface for this package so it can be re-used to create others
type Transformation interface {
	Start(ctx context.Context) error
}

// Zenformation holds our input and output as it moves along
type Zenformation struct {
	in  *In
	out *Out

	CeClient cloudevents.Client
}

// Start Returns if stopCh is closed or Send() returns an error.
func (t *Zenformation) Start(stopCh <-chan struct{}) error {
	fmt.Println("Starting Zenformation")

	if err := t.CeClient.StartReceiver(context.Background(), t.receive); err != nil {
		return err
	}
	return nil
}

func (t *Zenformation) receive(ctx context.Context, event cloudevents.Event) error {

	if err := event.DataAs(t.in); err != nil {
		fmt.Println("error occured unmarshaling data")
		t.ErrorHandler(err)
		return err
	}

	t.askComprehend(t.in.Description)

	//conver to int64 to match zendesk target spec
	i, err := strconv.Atoi(t.in.ID)
	if err != nil {
		fmt.Println("error converting id ")
		t.ErrorHandler(err)
		return err
	}

	t.out.ID = int64(i)

	rD, err := setResponseData(t)
	if err != nil {
		fmt.Println("error setting response data for cloud event")
		t.ErrorHandler(err)
		return err
	}

	if result := t.CeClient.Send(context.Background(), *rD); !cloudevents.IsACK(result) {
		fmt.Println("error sending cloud event")
		fmt.Println("failed to send cloud event, %v", err)
		t.ErrorHandler(err)
		return err
	}

	fmt.Println("Sent")
	fmt.Println("Sent %q", rD.String())
	fmt.Println(rD.String()) //test

	return nil
}

func setResponseData(t *Zenformation) (*cloudevents.Event, error) {
	transformEvent := cloudevents.NewEvent(cloudevents.VersionV1)
	transformEvent.SetID("122311")
	transformEvent.SetType(ceType)
	transformEvent.SetSource("some.origin/intance")
	transformEvent.SetSubject(ceSubject)
	transformEvent.SetTime(time.Now())

	err := transformEvent.SetData("application/json", t.out)
	if err != nil {
		fmt.Println("error occured setting CE data, %v", err)
		return nil, err
	}

	return &transformEvent, nil

}

func (t *Zenformation) askComprehend(txt string) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	c := comprehend.New(sess)
	dSI := comprehend.DetectSentimentInput{}
	dSI.SetLanguageCode(os.Getenv("LANGUAGE"))
	dSI.SetText(txt)

	req, resp := c.DetectSentimentRequest(&dSI)
	err := req.Send()
	if err == nil { // resp is now filled
		fmt.Print("got response from Comprehiend: ")
		fmt.Println(*resp.Sentiment)
		t.out.Tag = *resp.Sentiment
		//fmt.Println(*resp)
	} else {
		fmt.Println("error occured requesting from AWS Comprehiend")
		t.ErrorHandler(err)
	}
}

// ErrorHandler is a horrible error handler that should be re written :)
func (t *Zenformation) ErrorHandler(err error) {
	//t.logger.Error("An error ocurred", zap.Error(err))
	fmt.Println("an error has occured:")
	fmt.Println(err)

}
