FROM --platform=linux/amd64 busybox:1.35.0-uclibc as busybox

FROM --platform=linux/amd64 gcr.io/distroless/base-debian11

COPY --from=busybox /bin/sh /bin/sh
COPY --from=busybox /bin/ls /bin/ls

COPY --from=busybox /bin/mkdir /bin/mkdir
COPY --from=busybox /bin/cat /bin/cat

ENV LOCO_PASTER_API_PORT=8000
COPY loco-paster /
COPY dist/ /dist/


ENTRYPOINT ["/loco-paster"]
