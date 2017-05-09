FROM alpine
ADD k8guard-report /
ADD templates /templates
EXPOSE 3001
ENTRYPOINT ["/k8guard-report"]
