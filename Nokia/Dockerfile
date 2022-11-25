FROM mykafka:9092
RUN mkdir /root/subham

COPY aaa.txt /root/subham/
COPY queue /root/subham/
RUN cd /root/subham/ && go build queue.go

#RUN pwd
#RUN export PATH=$PATH:/usr/local/go/bin
