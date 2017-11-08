# CloudGo小项目简述：  
该项目简单地实现了一些web功能，在访问“/”时显示“Hello world！”，访问“/user/:id”时显示“Hello id！”  

在具体实现上，采用了Beego这个web框架，在对比revel，Negroni等选项后，选择了beego。  
Beego具有以下的优点:    
1.国人开发，丰富的中文文档和本土化的社区。  
2.提供的工具丰富，利用bee工具包可以方便的快速部署项目，实现自动化测试等等。
3.高度模块化，低耦合，可以在开发中独立使用beego的模块，可操作性强。  
4.强大方便的监控功能，实时监控当前server的运行状态。  
  
综上，我选择了Beego并且阅读了beego的技术文档。
实际上用beego构建web应用十分简单直观，很快就完成了cloudgo。  
  
使用curl来测试`curl -v http://loacalhost:8080/`(程序跑在8080端口）
```
admin@ubuntu:~$ curl -v http://localhost:8080/user/flow
* Hostname was NOT found in DNS cache
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /user/flow HTTP/1.1
> User-Agent: curl/7.35.0
> Host: localhost:8080
> Accept: */*
> 
< HTTP/1.1 200 OK
* Server beegoServer:1.9.0 is not blacklisted
< Server: beegoServer:1.9.0
< Date: Wed, 08 Nov 2017 06:43:34 GMT
< Content-Length: 12
< Content-Type: text/plain; charset=utf-8
< 
hello flow!
* Connection #0 to host localhost left intact
```
然后用ab来测试`ab -n 1000 -c 100 http://localhost:8080/user/flow`  
```
admin@ubuntu:~$ ab -n 1000 -c 100 http://localhost:8080/user/flow
This is ApacheBench, Version 2.3 <$Revision: 1528965 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        beegoServer:1.9.0
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/flow
Document Length:        12 bytes

Concurrency Level:      100
Time taken for tests:   0.097 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      156000 bytes
HTML transferred:       12000 bytes
Requests per second:    10265.78 [#/sec] (mean)
Time per request:       9.741 [ms] (mean)
Time per request:       0.097 [ms] (mean, across all concurrent requests)
Transfer rate:          1563.93 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.5      0       3
Processing:     0    9   8.0      8      38
Waiting:        0    9   8.0      8      38
Total:          0    9   8.0      8      39

Percentage of the requests served within a certain time (ms)
  50%      8
  66%     11
  75%     14
  80%     15
  90%     19
  95%     28
  98%     31
  99%     35
 100%     39 (longest request)
```
ab命令中-n指定了总的请求数，-c指定了并发数（即同一时间的请求）
需要关注的数据有`Requests per second`:吞吐量，每秒中可以处理的请求。
两个`Time per request`,一个是用户平均等待时间，一个是服务器平均处理请求用的时间。
