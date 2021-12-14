FROM centos:7

USER root
#变量
ENV version=go1.17.5 \
    GOROOT=/root/go \
    GOPATH=/root/Applications/Go \
    GOROOT_BOOTSTRAP=/root/go


#配置Go环境变量
RUN echo "export GOROOT=$HOME/go"  >> $HOME/.bash_profile 
RUN echo "export PATH=$PATH:$GOROOT/bin"  >> $HOME/.bash_profile  
RUN echo "export GOPATH=$HOME/Applications/Go"  >> $HOME/.bash_profile  && mkdir $GOPATH
RUN echo "export GOROOT_BOOTSTRAP=$HOME/${version}}"  >> $HOME/.bash_profile
RUN source $HOME/.bash_profile
#安装C工具
RUN yum install -y wget bison ed gawk gcc libc6-dev make

#编译包
RUN wget https://storage.googleapis.com/golang/${version}.src.tar.gz  && \
    tar -zxvf ${version}.src.tar.gz  && \
    mv go $GOROOT  && \
    cd $GOROOT/src && \
    ./all.bash

CMD [ "/bin/bash" ]
WORKDIR  $HOME