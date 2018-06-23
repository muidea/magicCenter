FROM centos

LABEL Author="rangh"

RUN mkdir -p /var/app

WORKDIR /var/app

COPY res.tar /tmp/
RUN cd /tmp \
 && tar -xvf res.tar -C /var/app/ \
 && rm -f res.tar

ADD magicCenter /var/app/magicCenter
ADD bootstrap.sh /var/app/bootstrap.sh

RUN chmod +x /var/app/magicCenter
RUN chmod +x /var/app/bootstrap.sh

ENTRYPOINT [ "/var/app/bootstrap.sh" ]

EXPOSE 8080
