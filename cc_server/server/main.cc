#include <memory>
#include <string>

#include <grpcpp/ext/proto_server_reflection_plugin.h> // For server reflection
#include <grpcpp/grpcpp.h>
#include <grpcpp/health_check_service_interface.h>    // For health checking

// Include your service implementation
#include "cc_server/server/ping_service_impl.h"

void RunServer(const std::string& server_address) {
    // Instantiate your service implementation
    project_root::cc_server::PingServiceImpl service_impl;

    // Enable health checking and server reflection
    // Health checking is good practice.
    // Server reflection allows tools like grpcurl to inspect the service.
    grpc::EnableDefaultHealthCheckService(true);
    grpc::reflection::InitProtoReflectionServerBuilderPlugin();

    // Build the server
    grpc::ServerBuilder builder;

    // Listen on the given address without any authentication mechanism (Insecure)
    // For production, you'd use secure credentials.
    builder.AddListeningPort(server_address, grpc::InsecureServerCredentials());

    // Register your service implementation with the builder.
    builder.RegisterService(&service_impl);

    // Assemble the server.
    std::unique_ptr<grpc::Server> server(builder.BuildAndStart());
    if (!server) {
        std::cerr << "Server failed to start on " << server_address << std::endl;
        return;
    }
    std::cout << "Server listening on " << server_address << std::endl;

    // Wait for the server to shutdown.
    // This call will block until the server is shut down (e.g., by Ctrl+C).
    server->Wait();
}

int main(int argc, char** argv) {
    std::string server_address = "0.0.0.0:7070";
    if (argc > 1) {
        server_address = argv[1]; // Allow overriding address from command line
    }

    RunServer(server_address);

    return 0;
}