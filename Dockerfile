ARG BASEIMAGE=busybox
FROM $BASEIMAGE

USER 65534

ARG BINARY=configmap-reload
COPY ./$BINARY /configmap-reload

ENTRYPOINT ["/configmap-reload"]
