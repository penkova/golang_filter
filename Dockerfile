FROM bitnami/minideb
ADD main /
EXPOSE 30
CMD ["/main"]