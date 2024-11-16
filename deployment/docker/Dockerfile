FROM junebao857/blog_base AS builder

WORKDIR /build
COPY ../.. .
RUN make

FROM ubuntu:20.04 AS run
WORKDIR /app
COPY --from=builder /build/bin/juneblog /app/
COPY --from=builder /build/scripts/run.sh /app/

EXPOSE 8080 2998
ENTRYPOINT ["/app/run.sh"]
