#include "cc_server/protos/ping_service.grpc.pb.h" // Generated gRPC header

namespace grpc {
class ServerContext;
class Status;
} // namespace grpc

// Using the package name for C++ namespace as defined in the proto
namespace project_root {
namespace cc_server {

class PingServiceImpl final : public PingService::Service {
public:
    // Constructor and Destructor (if needed for resource management)
    PingServiceImpl() = default;
    ~PingServiceImpl() override = default;

    // The Ping RPC method implementation
    ::grpc::Status Ping(::grpc::ServerContext* context,
                        const ::project_root::cc_server::PingRequest* request,
                        ::project_root::cc_server::PingResponse* response) override;
};

} // namespace cc_server
} // namespace project_root