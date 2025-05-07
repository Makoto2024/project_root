#include "cc_server/server/ping_service_impl.h"

#include <grpcpp/channel.h>
#include <grpcpp/support/status_code_enum.h>
#include <grpcpp/support/status.h>
#include <grpcpp/grpcpp.h>


#include <string>

#include "absl/log/log.h"
#include "absl/log/check.h"

#include "golang_server/protos/pong_service.pb.h"
#include "golang_server/protos/pong_service.grpc.pb.h"

using grpc::Channel;
using grpc::Status;
using grpc::ClientContext;


namespace {

bool call_pong(const std::string &msg, std::string &reply) {
    constexpr char target[] = "golang-server:8080";
    std::shared_ptr<Channel> channel =
        grpc::CreateChannel(target, grpc::InsecureChannelCredentials());
    if (!channel) {
        LOG(ERROR) << "Failed to create channel to " << target << std::endl;
        return false;
    }
    LOG(INFO) << "Client attempting to connect to: " << target << std::endl;


    project_root::golang_server::PongRequest request;
    request.set_msg(msg);

    project_root::golang_server::PongResponse response;
    grpc::ClientContext context;
    const grpc::Status status =
        project_root::golang_server::PongService::NewStub(channel)->Pong(&context, request, &response);
    if (!status.ok()) {
        LOG(ERROR)
            << "RPC failed with error code: " << status.error_code()
            << ", message: " << status.error_message() << std::endl;
        return false;
    }
    reply = response.msg();
    return true;

}

}  // namespace


namespace project_root {
namespace cc_server {

::grpc::Status PingServiceImpl::Ping(
    ::grpc::ServerContext* context,
    const ::project_root::cc_server::PingRequest* request,
    ::project_root::cc_server::PingResponse* response
) {
    // If msg contains ;ping, then just return the passed in msg.
    const std::string req_msg = request->msg();
    if (std::string::npos != req_msg.find(";ping")) {
        response->set_msg(req_msg);
        return ::grpc::Status::OK;
    }

    // Otherwise, call PongService.Pong with {req_msg};ping
    // and then return its responded msg.
    std::string pong_resp_msg;
    if (!call_pong(req_msg + ";ping", pong_resp_msg)) {
        return ::grpc::Status::CANCELLED;
    }
    response->set_msg(pong_resp_msg);
    return ::grpc::Status::OK;
}

} // namespace cc_server
} // namespace project_root