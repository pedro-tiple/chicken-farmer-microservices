# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds

#FROM golang as builder

# Must be outside gopath
#WORKDIR /barn-svc

#COPY go.mod go.mod
#COPY main main
#RUN CGO_ENABLED=0 go build -o barn-svc main/barn.go

FROM alpine

# Avoid using root
RUN adduser -D unpuser
USER unpuser

#COPY --from=builder /barn-svc/barn-svc /barn-svc
COPY barn-svc /barn-svc
ENV PORT 8080

CMD ["/barn-svc"]