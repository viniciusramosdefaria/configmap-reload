ARG BASEIMAGE=alpine
FROM $BASEIMAGE

ARG BINARY=configmap-reload
COPY ./$BINARY /configmap-reload

ENTRYPOINT ["/configmap-reload"]
