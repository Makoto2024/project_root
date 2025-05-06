#include "cc_server/server/ping_service_impl.h"

#include <string>

#include <grpcpp/support/status.h>

namespace project_root {
namespace cc_server {

::grpc::Status PingServiceImpl::Ping(
    ::grpc::ServerContext* context,
    const ::project_root::cc_server::PingRequest* request,
    ::project_root::cc_server::PingResponse* response
) {

    const std::string req_msg = request->msg();
    if (req_msg.empty()) {
        response->set_msg("ping;@NULL@");
    } else {
        response->set_msg("ping;" + req_msg);
    }
    return ::grpc::Status::OK;
}

} // namespace cc_server
} // namespace project_root