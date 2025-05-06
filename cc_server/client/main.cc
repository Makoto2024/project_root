#include <iostream>
#include <memory>
#include <string>

#include <grpcpp/grpcpp.h> // Core gRPC headers

#include "absl/log/log.h"
#include "absl/log/check.h"

// Include the generated gRPC headers. The path is based on your Bazel setup
// and the package name in your .proto file.
#include "cc_server/protos/ping_service.grpc.pb.h"

// Using directives for convenience, referencing the generated C++ namespaces/classes
using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;
using project_root::cc_server::PingRequest;  // From your .proto
using project_root::cc_server::PingResponse; // From your .proto
using project_root::cc_server::PingService;  // From your .proto

class PingServiceClient {
public:
    // Constructor: Initializes the stub using a gRPC channel.
    PingServiceClient(std::shared_ptr<Channel> channel)
        : stub_(PingService::NewStub(channel)) {}

    // Calls the Ping RPC method on the server.
    std::string SendPing(const std::string& user_message) {
        // 1. Prepare the request object.
        PingRequest request;
        request.set_msg(user_message);

        // 2. Prepare an empty response object to be filled by the server.
        PingResponse response;

        // 3. Create a ClientContext. This can be used to set deadlines, metadata, etc.
        ClientContext context;
        // Example: Set a deadline for the RPC call (e.g., 1 second)
        // std::chrono::system_clock::time_point deadline =
        //     std::chrono::system_clock::now() + std::chrono::seconds(1);
        // context.set_deadline(deadline);

        // 4. Make the actual RPC call.
        // This is a synchronous (blocking) call.
        Status status = stub_->Ping(&context, request, &response);

        // 5. Process the response and status.
        if (!status.ok()) {
            LOG(FATAL)
                << "RPC failed with error code: " << status.error_code()
                << ", message: " << status.error_message() << std::endl;
            exit(1);
        }
        return response.msg();
    }

private:
    // The stub provides the client-side API for the service methods.
    std::unique_ptr<PingService::Stub> stub_;
};

int main(int argc, char** argv) {
    // Determine the server address. Default to localhost:7070.
    const std::string server_address = "localhost:7070";

    // Create a gRPC channel to the server.
    // grpc::InsecureChannelCredentials() means the connection is not encrypted (no TLS).
    // For production, you would typically use secure credentials.
    std::shared_ptr<Channel> channel =
        grpc::CreateChannel(server_address, grpc::InsecureChannelCredentials());

    if (!channel) {
        LOG(ERROR) << "Failed to create channel to " << server_address << std::endl;
        exit(1);
    }
    LOG(INFO) << "Client attempting to connect to: " << server_address << std::endl;

    // Create the client object.
    PingServiceClient client(channel);

    // Determine the message to send.
    const std::string message_to_send = (argc > 1) ?
        std::string(argv[1]) :
        "Hello from C++ Client @ " + std::to_string(
            std::chrono::system_clock::now().time_since_epoch().count());
    LOG(INFO) << "Client sending message: \"" << message_to_send << "\"" << std::endl;

    // Call the SendPing method.
    const std::string reply = client.SendPing(message_to_send);

    // Print the server's reply.
    LOG(INFO) << "Client received reply: \"" << reply << "\"" << std::endl;
    return 0;
}
