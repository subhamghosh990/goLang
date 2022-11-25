module Revision

go 1.18

replace (
	abc v0.0.0 => ./grpc/proto
	client v0.0.0 => ./grpc/client
	//httpserver v0.0.0 => ./httpserver
	server v0.0.0 => ./grpc/server
)

require github.com/gorilla/mux v1.8.0
