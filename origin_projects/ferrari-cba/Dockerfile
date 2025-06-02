# final stage
FROM ubuntu:20.04

ENV TZ Asia/Seoul
ENV APP_HOME=/app/config/

RUN apt-get update;apt-get install -y cron tzdata ca-certificates;
RUN dpkg-reconfigure -f noninteractive tzdata

COPY ./dist/crontab.txt   /etc/cron.d/
COPY ./dist/ /app/

RUN chmod 777 /app/

RUN cd /app/
RUN mkdir /app/config
RUN chmod 0644 /etc/cron.d/crontab.txt
RUN chmod 0755 /app/script.sh
RUN crontab /etc/cron.d/crontab.txt

ENTRYPOINT cron start -f
