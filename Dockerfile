# Builder
FROM golang:1.18.3 AS builder

ARG GITHUB_PATH
ARG BRANCH

WORKDIR /go/src/
RUN git clone --branch $BRANCH $GITHUB_PATH
WORKDIR /go/src/quiz-ics-manager-api
RUN make build

# ics-manager-api

FROM golang:1.18.3 as server

COPY --from=builder /go/src/quiz-ics-manager-api/ics-manager-api /bin/
COPY --from=builder /go/src/quiz-ics-manager-api/config.toml /etc/

EXPOSE 8082

CMD ["/bin/ics-manager-api"]
