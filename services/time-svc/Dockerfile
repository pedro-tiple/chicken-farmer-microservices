# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds

#FROM golang as builder

# Must be outside gopath
#WORKDIR /user-svc

#COPY go.mod go.mod
#COPY main main
#RUN CGO_ENABLED=0 go build -o user-svc main/main.go

FROM alpine

# Avoid using root
RUN adduser -D unpuser
USER unpuser

#COPY --from=builder /user-svc/user-svc /user-svc
COPY time-svc /time-svc

CMD ["/time-svc"]