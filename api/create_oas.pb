syntax = "proto3";

package rosetta;
option go_package = "internal/api";

message File {
  string path = 1;
  bytes content = 2;
}

message CreateOASRequest {
  string apiKey = 1;
  string language = 2;
  bytes spec = 3;
  repeated File files = 4;
}

message CreateOASResponse {
  string spec = 1;
  string error = 2;
}

// *** GRPC Health Checking Protocol ***
// https://github.com/grpc/grpc/blob/master/doc/health-checking.md

message HealthCheckRequest {
  string service = 1;
}

message HealthCheckResponse {
  enum ServingStatus {
    UNKNOWN = 0;
    SERVING = 1;
    NOT_SERVING = 2;
    SERVICE_UNKNOWN = 3;  // Used only by the Watch method.
  }
  ServingStatus status = 1;
}

service FileService {
  rpc Check(HealthCheckRequest) returns (HealthCheckResponse);
  rpc CreateOAS (CreateOASRequest) returns (CreateOASResponse);
}
