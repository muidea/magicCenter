FROM centos

LABEL Author="rangh"

RUN mkdir -p /var/app

WORKDIR /var/app

COPY resource.tar /tmp/
RUN cd /tmp \
 && tar -xvf resource.tar -C /var/app/ \
 && rm -f resource.tar

ADD magicCenter /var/app/magicCenter
ADD bootstrap.sh /var/app/bootstrap.sh
ADD wait-for-it.sh /var/app/wait-for-it.sh

RUN chmod +x /var/app/magicCenter
RUN chmod +x /var/app/bootstrap.sh
RUN chmod +x /var/app/wait-for-it.sh

ENTRYPOINT [ "/var/app/bootstrap.sh" ]

EXPOSE 8080
