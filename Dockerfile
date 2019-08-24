FROM scratch
COPY server /
COPY cmd/dashboard/ /
CMD ["/server"]
