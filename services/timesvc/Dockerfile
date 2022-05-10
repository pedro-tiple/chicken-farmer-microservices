# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds

#FROM golang as builder

# Must be outside gopath
#WORKDIR /farmersvc

#COPY go.mod go.mod
#COPY main main
#RUN CGO_ENABLED=0 go build -o farmersvc main/main.go

FROM alpine

# Avoid using root
RUN adduser -D unpuser
USER unpuser

#COPY --from=builder /farmersvc/farmersvc /farmersvc
COPY timesvc /timesvc

CMD ["/timesvc"]