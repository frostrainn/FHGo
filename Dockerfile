FROM scratch
ADD server/app /
ENV TZ=Asia/Shanghai
CMD ["/app"]
