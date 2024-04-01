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
}

service FileService {
  rpc CreateOAS (CreateOASRequest) returns (CreateOASResponse);
}
