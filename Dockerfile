FROM debian:bookworm-slim

WORKDIR /
# ADD bin /bin/

RUN apt-get update && apt-get install -y ca-certificates

CMD ["/bin/sh"]
