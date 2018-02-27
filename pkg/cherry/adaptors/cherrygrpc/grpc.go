package cherrygrpc

import (
	"encoding/json"
	"net/http"

	"git.containerum.net/ch/kube-client/pkg/cherry"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	JSONMarshal   = json.Marshal
	JSONUnmarshal = json.Unmarshal
)

var httpToGRPCCode = map[int]codes.Code{
	444:                                     codes.Canceled,
	http.StatusOK:                           codes.OK,
	http.StatusBadRequest:                   codes.InvalidArgument,
	http.StatusRequestTimeout:               codes.DeadlineExceeded,
	http.StatusNotFound:                     codes.NotFound,
	http.StatusConflict:                     codes.AlreadyExists,
	http.StatusForbidden:                    codes.PermissionDenied,
	http.StatusInsufficientStorage:          codes.ResourceExhausted,
	http.StatusPreconditionFailed:           codes.FailedPrecondition,
	http.StatusGatewayTimeout:               codes.Aborted,
	http.StatusRequestedRangeNotSatisfiable: codes.OutOfRange,
	http.StatusNotImplemented:               codes.Unimplemented,
	http.StatusInternalServerError:          codes.Internal,
	http.StatusServiceUnavailable:           codes.Unavailable,
	http.StatusUnauthorized:                 codes.Unauthenticated,
}

// ToGRPC -- convert cherry error to grpc error for passing between services using grpc
func ToGRPC(errToPass *cherry.Err) error {
	data, err := JSONMarshal(errToPass)
	if err != nil {
		data = append(data, []byte("; with error "+err.Error())...)
	}
	code, mapped := httpToGRPCCode[errToPass.StatusHTTP]
	if !mapped {
		code = codes.Unknown
	}
	return status.Error(code, string(data))
}

// FromGRPC -- convert grpc error to cherry error. Sets ok to true if conversion was successful, false otherwise
func FromGRPC(errToCheck error) (ret *cherry.Err, ok bool) {
	grpcErr, ok := status.FromError(errToCheck)
	if !ok {
		return
	}
	err := JSONUnmarshal([]byte(grpcErr.Message()), &ret)
	ok = err == nil
	return
}
