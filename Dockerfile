FROM busybox

WORKDIR /root/
COPY build/blog ./
ENTRYPOINT [ "blog" ]