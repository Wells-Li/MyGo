FROM centos:7

USER root
#变量
ENV version=go1.17.5 \
    GOROOT=/root/go \
    GOPATH=/root/Applications/Go \
    GOROOT_BOOTSTRAP=/root/go1.17


#配置Go环境变量
RUN echo "export GOROOT=$HOME/go"  >> $HOME/.bash_profile 
RUN echo "export PATH=$PATH:$GOROOT/bin"  >> $HOME/.bash_profile  
RUN echo "export GOPATH=$HOME/Applications/Go"  >> $HOME/.bash_profile  && mkdir -p $GOPATH
RUN echo "export GOROOT_BOOTSTRAP=$HOME/go1.17"  >> $HOME/.bash_profile
RUN echo "export DISABLE_NET_TESTS=1" $HOME/.bash_profile 
RUN source $HOME/.bash_profile
#安装C工具
RUN yum install -y wget bison ed gawk gcc libc6-dev make git


#编译包
RUN wget https://studygolang.com/dl/golang/go1.17.5.linux-amd64.tar.gz > /dev/null 2>&1 && \
    tar -zxvf go1.17.5.linux-amd64.tar.gz -C $HOME/ > /dev/null 2>&1 && \
    cp -r $HOME/go $GOROOT_BOOTSTRAP  && \
    #编译安装go
    cd $GOROOT/src &&   ./all.bash && \
    cp $HOME/go/bin/*  /usr/bin/ && \
    rm -rf GOROOT_BOOTSTRAP=/root/go1.17 && \
    rm -rf go1.17.5.linux-amd64.tar.gz
CMD [ "/bin/bash" ]