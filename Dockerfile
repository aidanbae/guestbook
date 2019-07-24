FROM golang:1.12.4
RUN apt-get remove -y git ssh wget
RUN rm -rf /usr/bin/scp
RUN rm -rf /usr/bin/svn
RUN rm -rf /usr/bin/ssh

RUN groupadd tossinvest
RUN useradd --no-log-init -r -g tossinvest tossinvest
USER tossinvest
COPY . /go/
WORKDIR /go