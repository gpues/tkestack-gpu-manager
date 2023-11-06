
FROM $base_img

ARG version
ARG commit


RUN echo "/usr/local/nvidia/lib" > /etc/ld.so.conf.d/nvidia.conf && \
    echo "/usr/local/nvidia/lib64" >> /etc/ld.so.conf.d/nvidia.conf

ENV PATH=$PATH:/usr/local/nvidia/bin

# cgroup
VOLUME ["/sys/fs/cgroup"]

# display
EXPOSE 5678

COPY start.sh /
COPY copy-bin-lib.sh /

CMD ["/start.sh"]
