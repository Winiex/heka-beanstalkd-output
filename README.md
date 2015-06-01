##Introduction

This is a simple heka output plugin writes message to beanstalkd queues.

##Dependencies

You need to install [kr's golang binding library for beanstalk](https://github.com/kr/beanstalk).

To include this library in building process, you need to add the line below in $HEKA_ROOT/cmake/externals.cmake:

>git_clone(https://github.com/kr/beanstalk.git master)

##How to install

You can refer to [heka's documentation](https://hekad.readthedocs.org/en/v0.9.2/installing.html#building-hekad-with-external-plugins).

##Configurations

There exist three specific configurations for this plugin.

1. "beanstalkd_host": the host where the beanstalkd daemon is listening.
2. "beanstalkd_port": the port where the beanstalkd daemon is listening.
3. "beanstalkd_tube": the beanstalkd tube you want to put messages in.

An example may be like below:

>[BeanstalkdOutput]
>message_matcher = "TRUE"
>encoder = "PayloadEncoder"
>beanstalkd_host = "127.0.0.1"
>beanstalkd_port = "11000"
>beanstalkd_tube = "default"
