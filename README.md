##Introduction

This is a simple heka output plugin writes message to beanstalkd queues.

##Dependencies

You need to install [kr's golang binding library for beanstalk](https://github.com/kr/beanstalk).

To include this library in building process, you need to add the line below in $HEKA_ROOT/cmake/externals.cmake:

>git_clone(https://github.com/kr/beanstalk.git master)

##How to install

You can refer to [heka's documentation](https://hekad.readthedocs.org/en/v0.9.2/installing.html#building-hekad-with-external-plugins).
