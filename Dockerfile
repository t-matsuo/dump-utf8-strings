from centos:centos7.7.1908

RUN yum install -y epel-release
RUN yum install -y golang
RUN yum install -y glibc-static
RUN mkdir /build
WORKDIR /build
ENTRYPOINT ./compile.sh
