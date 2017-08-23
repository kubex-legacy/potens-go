package potens

import (
	"time"
	"github.com/kubex/proto-go/eventpipe"
	"go.uber.org/zap"
	"github.com/kubex/potens-go/services"
	"github.com/golang/protobuf/ptypes"
)

func (app *Application) connectToEventPipe() error {
	if app.services.eventPipeClient == nil {
		eventPipeConn, err := app.GetServiceConnection(services.EventPipe().Key())
		if err != nil {
			return err
		}
		app.services.eventPipeClient = eventpipe.NewEventPipeClient(eventPipeConn)
	}
	return nil
}

func (app *Application) SendEvent(distributorID, projectID, eventID string, attributes ...*eventpipe.Event_Attribute) {
	go app.SendEventSync(distributorID, projectID, eventID, attributes...)
}

func (app *Application) SendEventSync(distributorID, projectID, eventID string, attributes ...*eventpipe.Event_Attribute) error {
	app.connectToEventPipe()
	eventCtx, cancel := app.GrpcTimeoutContext(time.Second * 10)
	defer cancel()
	resp, err := app.services.eventPipeClient.Write(eventCtx, &eventpipe.Event{
		DistributorId: distributorID,
		ProjectId:     projectID,
		VendorId:      app.VendorID(),
		AppId:         app.AppID(),
		EventId:       eventID,
		Occurred:      ptypes.TimestampNow(),
		Attributes:    attributes,
	})
	if err != nil {
		app.Log().Error("Event Write Failed", zap.Error(err))
	} else if resp != nil && !resp.Success {
		app.Log().Error("Failed to write event")
	}

	return err
}
