# Benchmarks


***Note: The following benchmarks were done on a machine with AMD Ryzen 5 5600H @ 4.28 GHz (12 Core CPU)***

Potentially high traffic routes are stress tested using apache workbench to simulate 1000 concurrent users. 100,000 Requests are loaded up with a concurrency limit of 1000 (i.e 1000 requests are sent out at a single time) using the following command:

```bash
ab -n 100000 -c 1000 -H "Authorization: Bearer <AuthToken>" http://localhost:4000/api/<whatever-route>
```

## GET /api/items : Fetches all items on the menu

```
Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:
Server Hostname:        localhost
Server Port:            4000

Document Path:          /api/items
Document Length:        968 bytes

Concurrency Level:      1000
Time taken for tests:   7.651 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      110000000 bytes
HTML transferred:       96800000 bytes
Requests per second:    13070.98 [#/sec] (mean)
Time per request:       76.505 [ms] (mean)
Time per request:       0.077 [ms] (mean, across all concurrent requests)
Transfer rate:          14041.09 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   36  10.2     36      61
Processing:    20   40  10.4     40      63
Waiting:        0   17  13.3     13      58
Total:         41   76   2.6     76      85

Percentage of the requests served within a certain time (ms)
  50%     76
  66%     77
  75%     77
  80%     77
  90%     78
  95%     79
  98%     80
  99%     81
 100%     85 (longest request)
```


## GET /api/items/1 : Fetches a specific item ( itemID = 1 )


```
Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:
Server Hostname:        localhost
Server Port:            4000

Document Path:          /api/items/1
Document Length:        121 bytes

Concurrency Level:      1000
Time taken for tests:   7.626 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      25300000 bytes
HTML transferred:       12100000 bytes
Requests per second:    13112.74 [#/sec] (mean)
Time per request:       76.262 [ms] (mean)
Time per request:       0.076 [ms] (mean, across all concurrent requests)
Transfer rate:          3239.77 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   36  10.4     36      68
Processing:    19   39  10.5     40      71
Waiting:        0   16  13.1     13      65
Total:         42   76   3.0     76      96

Percentage of the requests served within a certain time (ms)
  50%     76
  66%     76
  75%     77
  80%     77
  90%     78
  95%     79
  98%     81
  99%     86
 100%     96 (longest request)
```

## POST /api/orders : 

After tweaking the MySQL MaxOpenConns and MaxIdleConns, this is the best benchmark I got for creating 100,000 orders with a concurrency limit of 1000 requests.

```
Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:
Server Hostname:        localhost
Server Port:            4000

Document Path:          /api/orders
Document Length:        Variable

Concurrency Level:      1000
Time taken for tests:   10.034 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      31277727 bytes
Total body sent:        35400000
HTML transferred:       18077727 bytes
Requests per second:    9966.24 [#/sec] (mean)
Time per request:       100.339 [ms] (mean)
Time per request:       0.100 [ms] (mean, across all concurrent requests)
Transfer rate:          3044.15 [Kbytes/sec] received
                        3445.36 kb/s sent
                        6489.52 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   35   7.0     36      53
Processing:    13   65  17.8     62     356
Waiting:        4   53  18.1     50     342
Total:         14  100  17.0     99     397

Percentage of the requests served within a certain time (ms)
  50%     99
  66%    101
  75%    103
  80%    104
  90%    109
  95%    119
  98%    142
  99%    169
 100%    397 (longest request)
 ```




### Optimizations Currently Implemented
- In-memory cache for critical routes
- Cached authentication middleware