package potens

import (
	"github.com/kubex/potens-go/services"
	"github.com/kubex/proto-go/socket"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Socket helper for sockets
func (app *Application) Socket() *socketHandler {
	if app.services.socketHandler == nil {
		con, err := app.GetServiceConnection(services.Socket().Key())
		if err != nil {
			app.Log().Fatal("Unable to connect to Sockets Server", zap.String("error", err.Error()))
		}
		app.services.socketHandler = newHandler(con, app.GetGrpcContext(), app.VendorID(), app.AppID())
	}
	return app.services.socketHandler
}

type socketHandler struct {
	client     socket.SocketClient
	connection *grpc.ClientConn
	ctx        context.Context
	appID      string
	vendorID   string
}

func newHandler(cc *grpc.ClientConn, ctx context.Context, vendor string, appID string) *socketHandler {
	return &socketHandler{connection: cc, ctx: ctx, client: socket.NewSocketClient(cc), appID: appID, vendorID: vendor}
}

func (h *socketHandler) Subscribe(socketID string, channelName string) (*socket.PublishResponse, error) {
	return h.client.Subscribe(h.ctx, &socket.SubscribeMessage{
		SessionId: socketID,
		Channel:   channelName,
	})
}

func (h *socketHandler) SendMessage(channelName string, action string, payload string) (*socket.PublishResponse, error) {
	return h.client.Publish(h.ctx, &socket.SocketMessage{
		Channel: channelName,
		Action:  action,
		Payload: payload,
	})
}
