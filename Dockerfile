FROM busybox

WORKDIR /root/
COPY build/blog ./
# COPY uploadfile/ ./uploadfile/
ENTRYPOINT [ "./blog" ]