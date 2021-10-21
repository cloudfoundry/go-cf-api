---
sidebar_position: 2
title: Load Tests
---

# Load Tests

To show the possible advantage of this PoC implementation in comparison to the Ruby implementation ([cloud_controller_ng](https://github.com/cloudfoundry/cloud_controller_ng)) we ran a few load tests to find out the limits of both implementations.

The aim of this was to get an idea how much performance a new implementation could potentially offer.
So we compared selected endpoints that we already fully implemented in the PoC Implementation.
The load test was conducted with [vegeta](https://github.com/tsenart/vegeta) and [netdata](https://github.com/netdata/netdata) to monitor vm and database resources.
## Go vs Ruby implementation

We executed the same test on both implementations on compareable setups.
For this we implemented quite a number of middleware functionality also in the PoC e.g. Logging, Metrics(Prometheus), Rate Limiting(Altough disabled for the test), Authorization and Role checking.
The test then ran the same api call against specific endpoints with varous number of requests per second.
These endpoints where fully implemented in the PoC and compliant with [CF V3 API 3.109.0](https://v3-apidocs.cloudfoundry.org/version/3.109.0).

The tested endpoints where the ones we also have implemented in this PoC at the time of testing. Namely:
- [/v3/info](https://v3-apidocs.cloudfoundry.org/version/3.109.0/#info)
- [/v3/security_groups/:guid](https://v3-apidocs.cloudfoundry.org/version/3.109.0/#get-a-security-group)
- [/v3/buildpacks](https://v3-apidocs.cloudfoundry.org/version/3.109.0/#list-buildpacks)

### Test setup
Firstly we seeded the database with entries to have realistic database processing times as well as a realistic result sets returned by postgres. We ran these tests on a db with:
- 50 Spaces
- 1000 Orgs
- 20 Buildpacks
- 40000 Security Groups

Other key setup facts where:
- The test was conducted on AWS
- We used cloud_controller_ng version `3.108.0` and the PoC implementation version `0.0.19`
- Both implementations ran with a single instance
- We executed these tests from a gorouter vm that had an instance type of `c4.8xlarge`
- We skipped the cloudfoundry routing layer by accessing the tested instances directly over the private network the gorouter is also part of.
- The PoC(Go) Implementation of the CloudController was deployed on a VM of the instance type `t3a.xlarge`
- Postgres database used was intentionally oversized to a `db.m5.2xlarge` instance to prevent the db being a bottleneck.
- The maximum amount of concurrently usable db connections for both implementations was set to `500`
- The `threadpool_size` config parameter in cloud_controller_ng was set to `1000` to theoretically make the implementations more equal in the maximum possible parallelism(although we know ruby is single threaded).
- The test was conducted with an admin user and disabled rate limiting on both implementations.

### /v3/info

The load test of the `/v3/info` endpoint tells us the maximum number of requests the web server on both implementations can handle without expensive buissiness logic, authentication, accessing other apis or the database. It represents the raw theoretical trouhput the underlying technology stack can serve.
##### Ruby Implementation

While 64 requests per second on the `v3/info` endpoint run quite fine on the ruby implementation, the instance breaks down at 128 requests per second.
![](/img/loadtests/v1/vegeta/ng_info_128.png)
Interesting is, that from the first second on, we have api error responses and they also take quite long to return(13 seconds). Also interesting is that the time they need to return decreases linearly until it matches the response times of successfull calls.
Only 30% of the requests are served correctly, which is even lower then the previously accieved 64 requests per second.
As one can see the error codes are also `502 Bad Gatewaty` so the behaviour we already know of cloud_controller_ng showed up. Most likely the cloud_controller_ng web process did not answer its own health endpoint(in time) that is checked by nginx so the instance got removed from local nginx load balancing and a `502` was served. Interestingly enough we also see `500 Internal Server Error` responses which should not happen as this means unhandled exceptions in the code.
So somewehere around 100 requests per second is the estimated hard limit of requests per second for any endpoint as `/v3/info` is one of the cheapest ones.

In the [VM Resources Section](#cpu-load-1) one can see that at 128 requests per second, the api is not limited by system resources. However [Cloud Controller Load](#cloudcontroller-load) shows a usage of 25% which means a full core. This is the expected hard limit of ruby hiting here as ruby runs single threaded.
##### Go Implementation

We initially thought to never be able to handle 4k requests per second on a 4 core VM so the highest measurement we have are 4k requests per second and that seems to work fine on the PoC implementation. In the resources tab
![](/img/loadtests/v1/vegeta/go_info_4096.png)

In the [VM Resources Section](#cpu-load) one can see that at 4096 reqeusts per second, the api is fully using all cpu resources and this is likely limited by that.In high load situations, the HAProxy process seems to consume a lot of cpu. In this case HAProxy consumed 2,5 Cores while the PoC consumed rughly a single core.
### /v3/buildpacks

The `/v3/buildpacks` endpoint is a endpoint that requires no role checking but authorization and does quite simple sql queries on just the buildpacks table. It shows the maximum load of a simple listing endpoint for both implementations.

##### Ruby Implementation

For the ruby implementation we saw that at a rate of 32 requests per seconds works quite well but just with a troughput of 27 requests per second, while 64 requests per seconds break the instance with already described behaviour in the [/v3/info](TODO) test.

![](/img/loadtests/v1/vegeta/ng_buildpacks_32.png)
![](/img/loadtests/v1/vegeta/ng_buildpacks_64.png)

So the estimated maximum would be around 30 requests per second since in the [VM Resources Section](#cpu-load-1) we see no overall resource limitation rather the ruby thread using its full single core already at 32 requests per second. Additionally the trouhput of 27 tells us that requests are piling up already.
##### Go Implementation

On the PoC implementation we see some requests timing out on the client side(30s timeout) at 1k requests per second.
![](/img/loadtests/v1/vegeta/go_buildpacks_1024.png)
They maybe still would have been served but our default client timeout defines that every request longer than 30 seconds counts as failed.
One can thus see around 550 requests per second as upper limit for this endpoint in the PoC as the average trouput was around that area.
### /v3/security_groups/:guid

The `/v3/security_groups/:guid` endpoint involves checking of roles and also has links to spaces. It has more complex queries and also trough the high number of security_groups in the database, we test how the implementations can cope with processing time of the DB. Ideally we wanted to test the listing security_groups endpoint but that was not yet implemented at the time of load testing in the PoC.
##### Ruby Implementation

For the ruby implementation we saw that again a rate of 64 requests per second overloads the cloud_controller_ng via using the whole single thread ruby has available.

![](/img/loadtests/v1/vegeta/ng_security_groups_64.png)

However this time the trouput at a rate of 32 requests per second the rate and trouhput matches so it seems the `/v3/security_groups/:guid` is a bit cheaper on the cpu. Altough the single ruby thread is already exhausted again at that rate. So we estimate the sustainable trouput beeing around 40 requests per second.
##### Go Implementation

For the `/v3/security_groups/:guid` endpoint we also see a limit around 1k requests per second. Although a bit higher than `/v3/buildpacks` as at 1k requests, all get served within the 30s timeout limit.

![](/img/loadtests/v1/vegeta/go_security_groups_1024.png)

Based on the troughput numbers, a sustained load of 750 requests per second seems realistic.

A rather interesting behaviour opposed to cloud_controller_ng can be seen at a rate of 2k requests per second.
The response times rise for nearly 20 seconds until all requests run into a timeout. Yet no server error was visible just an enourmous increase of latencies.
![](/img/loadtests/v1/vegeta/go_security_groups_2048.png)

### VM Resources

Here are the recordings of the system resources while executing the test. The full data is available as netdata dump in the [raw results tarball archive](#raw-measurements).
#### Go Implementation

To correlate the metrics to specific test runs, here are the timestamps of the test runs:
```shell
Targeting: https://10.0.65.75/v3/info
2021-11-03 08:56:07.314080 - Using attack rate: 8 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/info"}
2021-11-03 08:56:57.638468 - Using attack rate: 16 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/info"}
2021-11-03 08:57:47.913230 - Using attack rate: 24 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/info"}
2021-11-03 08:58:38.179441 - Using attack rate: 32 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/info"}
2021-11-03 08:59:28.498043 - Using attack rate: 64 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/info"}
2021-11-03 09:00:18.947272 - Using attack rate: 128 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/info"}
2021-11-03 09:01:09.172216 - Using attack rate: 256 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/info"}
2021-11-03 09:01:59.411845 - Using attack rate: 512 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/info"}
2021-11-03 09:02:49.780627 - Using attack rate: 1024 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/info"}
2021-11-03 09:03:40.089794 - Using attack rate: 2048 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/info"}
2021-11-03 09:04:30.480274 - Using attack rate: 4096 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/info"}


Targeting: https://10.0.65.75/v3/buildpacks
2021-11-03 09:05:40.888477 - Using attack rate: 8 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/buildpacks"}
2021-11-03 09:06:31.237477 - Using attack rate: 16 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/buildpacks"}
2021-11-03 09:07:21.503889 - Using attack rate: 24 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/buildpacks"}
2021-11-03 09:08:11.840010 - Using attack rate: 32 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/buildpacks"}
2021-11-03 09:09:02.160462 - Using attack rate: 64 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/buildpacks"}
2021-11-03 09:09:52.466723 - Using attack rate: 128 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/buildpacks"}
2021-11-03 09:10:42.759395 - Using attack rate: 256 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/buildpacks"}
2021-11-03 09:11:33.273219 - Using attack rate: 512 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/buildpacks"}
2021-11-03 09:12:23.564614 - Using attack rate: 1024 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/buildpacks"}
2021-11-03 09:13:38.855490 - Using attack rate: 2048 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/buildpacks"}
2021-11-03 09:14:59.447755 - Using attack rate: 4096 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/buildpacks"}


Targeting: https://10.0.65.75/v3/security_groups
2021-11-03 09:16:20.424447 - Using attack rate: 8 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/security_groups/619fb6a8-038c-4934-9285-1c2793cb3b1c"}
2021-11-03 09:17:10.786576 - Using attack rate: 16 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/security_groups/619fb6a8-038c-4934-9285-1c2793cb3b1c"}
2021-11-03 09:18:01.054648 - Using attack rate: 24 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/security_groups/619fb6a8-038c-4934-9285-1c2793cb3b1c"}
2021-11-03 09:18:51.309282 - Using attack rate: 32 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/security_groups/619fb6a8-038c-4934-9285-1c2793cb3b1c"}
2021-11-03 09:19:41.730875 - Using attack rate: 64 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/security_groups/619fb6a8-038c-4934-9285-1c2793cb3b1c"}
2021-11-03 09:20:32.027494 - Using attack rate: 128 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/security_groups/619fb6a8-038c-4934-9285-1c2793cb3b1c"}
2021-11-03 09:21:22.264486 - Using attack rate: 256 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/security_groups/619fb6a8-038c-4934-9285-1c2793cb3b1c"}
2021-11-03 09:22:12.736188 - Using attack rate: 512 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/security_groups/619fb6a8-038c-4934-9285-1c2793cb3b1c"}
2021-11-03 09:23:03.124116 - Using attack rate: 1024 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/security_groups/619fb6a8-038c-4934-9285-1c2793cb3b1c"}
2021-11-03 09:24:04.777251 - Using attack rate: 2048 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/security_groups/619fb6a8-038c-4934-9285-1c2793cb3b1c"}
2021-11-03 09:25:25.160157 - Using attack rate: 4096 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/security_groups/619fb6a8-038c-4934-9285-1c2793cb3b1c"}
```

##### CPU Load
![](/img/loadtests/v1/netdata/go_overall_cpu.png)
CPU is fully used, altough a syslog process(blackbox) always broke and went out of control in the middle of the tests and consumed a full core consistently but with a high nice value so it is not expected to have influenced the test. Only a system reboot helped to resolve the broken blackbox process as it will restart once killed and get stuck again.
###### HAProxy CPU Load
![](/img/loadtests/v1/netdata/go_haproxy_usage.png)

At high number of requests or error rates, the load of the haproxy process is quite substential. There might be optimizations in haproxy config possible here to improove overall performance of the api since it is cpu limited. When a less cpu is consumed by haproxy in a high load scenario, that could increase the trouhput of the go process. 
###### PoC CPU Load
![](/img/loadtests/v1/netdata/go_go_usage.png)
At maximum only 1-1.5 Cores are used. At high load HAProxy consumes the rest so 1.5-3 Cores. 
##### RAM Usage
![](/img/loadtests/v1/netdata/go_ram_usage.png)

##### TCP Connections
![](/img/loadtests/v1/netdata/go_tcp_connections.png)
##### DB connections and transactions per second
![](/img/loadtests/v1/netdata/go_db_transactions.png)
On can observe the DB connection pool is scaled up and down on demand. At maximum we observed around 5k transactions/s on the database. A more realisic peak load that is reached more constantly is around 3k-3.5k requests/s. But as we designed the Implementation to be the bottleneck and not the Database, the transactions per second is not the maximum one could get out of a single instance.

For this we conducted a [second test](#go---test-single-db-connection) with a single db connection to get a estimation of how many api requests we could serve with the currently limiting factor of 5k available db connections on AWS RDS instances, given that it could handle all of the transactions.

#### Ruby Implementation
To correlate the metrics to specific test runs, here are the timestamps of the test runs:
```shell
Targeting: https://10.0.65.66:9024/v3/info
2021-11-02 14:49:37.770203 - Using attack rate: 8 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/info"}
2021-11-02 14:50:28.070121 - Using attack rate: 16 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/info"}
2021-11-02 14:51:18.421874 - Using attack rate: 24 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/info"}
2021-11-02 14:52:08.730871 - Using attack rate: 32 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/info"}
2021-11-02 14:52:58.972751 - Using attack rate: 64 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/info"}
2021-11-02 14:53:49.302106 - Using attack rate: 128 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/info"}
2021-11-02 14:54:45.049050 - Using attack rate: 256 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/info"}
2021-11-02 14:55:39.396009 - Using attack rate: 512 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/info"}
2021-11-02 14:56:35.547554 - Using attack rate: 1024 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/info"}
2021-11-02 14:57:25.914571 - Using attack rate: 2048 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/info"}
2021-11-02 14:58:19.378228 - Using attack rate: 4096 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/info"}


Targeting: https://10.0.65.66:9024/v3/buildpacks
2021-11-02 14:59:17.465476 - Using attack rate: 8 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/buildpacks"}
2021-11-02 15:00:07.879977 - Using attack rate: 16 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/buildpacks"}
2021-11-02 15:00:58.330487 - Using attack rate: 24 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/buildpacks"}
2021-11-02 15:01:48.664109 - Using attack rate: 32 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/buildpacks"}
2021-11-02 15:02:45.669107 - Using attack rate: 64 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/buildpacks"}
2021-11-02 15:03:42.210203 - Using attack rate: 128 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/buildpacks"}
2021-11-02 15:04:51.524479 - Using attack rate: 256 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/buildpacks"}
2021-11-02 15:05:42.108656 - Using attack rate: 512 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/buildpacks"}
2021-11-02 15:07:12.638379 - Using attack rate: 1024 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/buildpacks"}
2021-11-02 15:09:04.571407 - Using attack rate: 2048 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/buildpacks"}
2021-11-02 15:10:19.833107 - Using attack rate: 4096 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/buildpacks"}


Targeting: https://10.0.65.66:9024/v3/security_groups
2021-11-02 15:11:33.580214 - Using attack rate: 8 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/security_groups/619fb6a8-038c-4934-9285-1c2793cb3b1c"}
2021-11-02 15:12:24.311953 - Using attack rate: 16 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/security_groups/619fb6a8-038c-4934-9285-1c2793cb3b1c"}
2021-11-02 15:13:14.636080 - Using attack rate: 24 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/security_groups/619fb6a8-038c-4934-9285-1c2793cb3b1c"}
2021-11-02 15:14:05.095217 - Using attack rate: 32 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/security_groups/619fb6a8-038c-4934-9285-1c2793cb3b1c"}
2021-11-02 15:14:55.409231 - Using attack rate: 64 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/security_groups/619fb6a8-038c-4934-9285-1c2793cb3b1c"}
2021-11-02 15:15:55.436128 - Using attack rate: 128 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/security_groups/619fb6a8-038c-4934-9285-1c2793cb3b1c"}
2021-11-02 15:16:57.788453 - Using attack rate: 256 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/security_groups/619fb6a8-038c-4934-9285-1c2793cb3b1c"}
2021-11-02 15:17:54.845722 - Using attack rate: 512 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/security_groups/619fb6a8-038c-4934-9285-1c2793cb3b1c"}
2021-11-02 15:19:05.253594 - Using attack rate: 1024 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/security_groups/619fb6a8-038c-4934-9285-1c2793cb3b1c"}
2021-11-02 15:20:11.579831 - Using attack rate: 2048 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/security_groups/619fb6a8-038c-4934-9285-1c2793cb3b1c"}
2021-11-02 15:21:24.921151 - Using attack rate: 4096 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.66:9024/v3/security_groups/619fb6a8-038c-4934-9285-1c2793cb3b1c"}
```
##### CPU Load
![](/img/loadtests/v1/netdata/ng_overall_cpu.png)
The CPU load is alot more spiky and even when the api responds with failures, not all resources are used.
###### Nginx Load
![](/img/loadtests/v1/netdata/ng_nginx_cpu.png)
At high loads nginx consumes a full core which is still a lot less than haproxy did.
###### CloudController Load
![](/img/loadtests/v1/netdata/ng_ng_cpu.png)
In this graph one can clearly see the real resource limitation beeing rubys single threaded limitation. Using a full core is the limit of a ruby process and once hit the api responds with failures or fails to answer reqeuests.
##### RAM Usage
![](/img/loadtests/v1/netdata/ng_ram.png)

##### TCP Connections
![](/img/loadtests/v1/netdata/ng_tcp_connections.png)
##### DB connections and transactions per second
![](/img/loadtests/v1/netdata/ng_db_transactions.png)
On can observe the DB connection pool is scaled up never down again. Connections would get closed once they time out after an hour. At maximum we observed around 450 transactions/s on the database. While a more realisic peak load that is reached more constantly is around 300-350 requests/s.

### Summary 

Maximum requests per second on this test setup
<table>
  <tbody>
    <tr>
      <td align="center">Endpoint</td>
      <td align="center">Go Implementation</td>
      <td align="center">Ruby Implementation</td>
      <td align="center">Improvement</td>
    </tr>
    <tr>
      <td align="center">/v3/info</td>
      <td align="center">4000</td>
      <td align="center">100</td>
      <td align="center">+3900%(40 Times more)</td>
    </tr>
    <tr>
      <td align="center">/v3/buildpacks</td>
      <td align="center">550</td>
      <td align="center">30</td>
      <td align="center">+1730%(18.3 Times more)</td>
    </tr>
    <tr>
      <td align="center">/v3/security_groups</td>
      <td align="center">750</td>
      <td align="center">40</td>
      <td align="center">+1775%(18.75 Times more)</td>
    </tr>
  </tbody>
</table>

Further Findings
- The Go Implementation can issue 10x more transactions per second (operations on the Database) over the same amount of db connections. Potentially even more but likely we do less db transactions per api requests and are cpu bound before being limited by postgres connections.
- The Go Implementation does not get pulled out of load balancing in an overload situation.
- The Go Implementation may wastes a lot of resources in HAProxy that may need to be looked at and tweaked.
- The Go Implementation uses more RAM in general and it increases/decreases more on peak loads. But that could be expected since it handles up to 40 times more requests.


We also briefly tested the PoC on a large vm (64 cores) how it would scale horizontally. We saw 30k+ `/v3/info` calls per second beeing handled by a single instance. At that point the emitter of the requests was to weak and one would need to set up a distributed load test to see real results.
Since this would be time consuming and not compareable against a single threaded implementation we stopped further load tests in that scale.
### Raw Measurements

All measurement Runs can be found in below table. At the bottom one can download all raw measurements in a tar.xz file.
How to look at the raw values is explained in [#viewing-raw-results](#viewing-raw-results)

<table>
  <tbody>
    <tr>
      <td align="center">Go Implementation</td>
      <td align="center">Ruby Implementation</td>
    </tr>
    <tr>
      <td align="center">
        GET /v3/info
        <ul>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-info_8.html">v3/info - 8 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-info_16.html">v3/info - 16 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-info_24.html">v3/info - 24 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-info_32.html">v3/info - 32 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-info_64.html">v3/info - 64 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-info_128.html">v3/info - 128 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-info_256.html">v3/info - 256 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-info_512.html">v3/info - 512 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-info_1024.html">v3/info - 1024 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-info_2048.html">v3/info - 2048 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-info_4096.html">v3/info - 4096 req/s</a> </li>
        </ul>
        GET /v3/buildpacks
        <ul>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-buildpacks_8.html">v3/buildpacks - 8 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-buildpacks_16.html">v3/buildpacks - 16 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-buildpacks_24.html">v3/buildpacks - 24 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-buildpacks_32.html">v3/buildpacks - 32 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-buildpacks_64.html">v3/buildpacks - 64 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-buildpacks_128.html">v3/buildpacks - 128 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-buildpacks_256.html">v3/buildpacks - 256 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-buildpacks_512.html">v3/buildpacks - 512 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-buildpacks_1024.html">v3/buildpacks - 1024 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-buildpacks_2048.html">v3/buildpacks - 2048 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-buildpacks_4096.html">v3/buildpacks - 4096 req/s</a> </li>
        </ul>
        GET /v3/security_groups/:guid
        <ul>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-security_groups_8.html">v3/security_groups - 8 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-security_groups_16.html">v3/security_groups - 16 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-security_groups_24.html">v3/security_groups - 24 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-security_groups_32.html">v3/security_groups - 32 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-security_groups_64.html">v3/security_groups - 64 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-security_groups_128.html">v3/security_groups - 128 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-security_groups_256.html">v3/security_groups - 256 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-security_groups_512.html">v3/security_groups - 512 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-security_groups_1024.html">v3/security_groups - 1024 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-security_groups_2048.html">v3/security_groups - 2048 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test/html/v3-security_groups_4096.html">v3/security_groups - 4096 req/s</a> </li>
        </ul>
      </td>
      <td align="center">
        GET /v3/info
        <ul>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-info_8.html">v3/info - 8 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-info_16.html">v3/info - 16 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-info_24.html">v3/info - 24 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-info_32.html">v3/info - 32 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-info_64.html">v3/info - 64 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-info_128.html">v3/info - 128 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-info_256.html">v3/info - 256 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-info_512.html">v3/info - 512 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-info_1024.html">v3/info - 1024 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-info_2048.html">v3/info - 2048 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-info_4096.html">v3/info - 4096 req/s</a> </li>
        </ul>
        GET /v3/buildpacks
        <ul>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-buildpacks_8.html">v3/buildpacks - 8 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-buildpacks_16.html">v3/buildpacks - 16 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-buildpacks_24.html">v3/buildpacks - 24 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-buildpacks_32.html">v3/buildpacks - 32 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-buildpacks_64.html">v3/buildpacks - 64 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-buildpacks_128.html">v3/buildpacks - 128 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-buildpacks_256.html">v3/buildpacks - 256 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-buildpacks_512.html">v3/buildpacks - 512 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-buildpacks_1024.html">v3/buildpacks - 1024 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-buildpacks_2048.html">v3/buildpacks - 2048 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-buildpacks_4096.html">v3/buildpacks - 4096 req/s</a> </li>
        </ul>
        GET /v3/security_groups/:guid
        <ul>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-security_groups_8.html">v3/security_groups - 8 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-security_groups_16.html">v3/security_groups - 16 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-security_groups_24.html">v3/security_groups - 24 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-security_groups_32.html">v3/security_groups - 32 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-security_groups_64.html">v3/security_groups - 64 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-security_groups_128.html">v3/security_groups - 128 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-security_groups_256.html">v3/security_groups - 256 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-security_groups_512.html">v3/security_groups - 512 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-security_groups_1024.html">v3/security_groups - 1024 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-security_groups_2048.html">v3/security_groups - 2048 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_ng_load_test/html/v3-security_groups_4096.html">v3/security_groups - 4096 req/s</a> </li>
        </ul>
      </td>
    </tr>
    <tr>
      <td align="center"><a target="_parent" href="/loadtests/v1/cc_go_load_test/cc_go_load_test.tar.xz">Raw Archive</a></td>
      <td align="center"><a target="_parent" href="/loadtests/v1/cc_ng_load_test/cc_ng_load_test.tar.xz">Raw Archive</a></td>
    </tr>
  </tbody>
</table>

## Go - Test single db connection

This second test aims not to compare against the cloud_controller_ng ruby implementation.
Instead it is meant to show how many transactions on the DB can be issued over a single DB connection.
This is important as currently a hard limit is set for e.g. AWS RDS instances. The biggest RDS instance can handle just [5k open connections](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Limits.html#RDS_Limits.MaxConnections) `LEAST({DBInstanceClassMemory/9531392}, 5000)`.

So when doing e.g. 30 `/v3/buildpacks` requests per second over 160 cloud_controller_ng instances (maximum number of instances on 5000 db connections with default db connection pool) we can just squeeze out 4.8k requests per second under optimal conditions(the db responding as if it wouldnt have load). This is currently a limiting factor for scaling the cloud_controller_ng further.

With this test we will see how many requests per second and db transactions could theoretically be served by the go implementation with 5000 connections and if this number is still a hard limit for the PoC implementations scaling.
### Test setup

The Setup was the same as in the [first test run](#test-setup).
However the the maximum amount of concurrently usable db connections was set to `1`.
Additionally just the PoC implementatio was tested.
### /v3/info

With a single DB Connection, the performance of the `/v3/info` is expected to be unchanged from the first test.
Yet we see less requests/s where handled but we assume this is likely related to vm burst budget. 
We should have used non burstable instance types for the test, we may redo this then but for now as rugh comparison we leave the test results as the are.

![](/img/loadtests/v1/vegeta/gosingle_info_4096.png)
### /v3/buildpacks

The buildpacks endpoint can now just handle around 250 requests per seconds which produce rughly 1k transactions per second on the database.
That in return means indeed the single db connection is fully used and the bottleneck.

![](/img/loadtests/v1/vegeta/gosingle_buildpacks_256.png)

At 512 requests per second the go implementation is overloaded in a way that api calls take longer and longer until eventually the 30 timeout of vegeta cancels the request for a subset of really slow requests (50%) and marks it as error.

![](/img/loadtests/v1/vegeta/gosingle_buildpacks_512.png)
### /v3/security_groups/:guid

The security_groups endpoint can now just handle around 320 requests per seconds which produce rughly 1k transactions per second on the database as well.
This endpoint is thus a bit cheaper than `v3/buildpacks` since it issues less or cheaper database statements.


So 256 requests per second serve perfectly fine.
![](/img/loadtests/v1/vegeta/gosingle_security_groups_256.png)
While it struggles a bit with 512 requests per second doing just around 320.
![](/img/loadtests/v1/vegeta/gosingle_security_groups_512.png)
Since the 1k transactions per second seems to stay the same in comparison to the `/v3/buildpacks` endpoint, one can expect this beeing a realistic limit over a single db connection.

### VM Resources

Metrics where basically the same between this run and the one with unlimited db connections.
The exception beeing DB connections and transactions per second of course.
#### DB connections and transactions per second
![](/img/loadtests/v1/netdata/gosingle_db_transactions.png)
The db connections show around 30 connections because another cloud_controller_ng instance is connected at the same time but without load.
### Summary 

We estimate the PoC implementation can serve around 1k transactions on a single db connection when not beeing limited by cpu, rather this single db connection. As 1k db transactions correlate to 320 `/v3/security_groups/:guid` calls and 250 `/v3/buildpacks` calls on could theoretically make `320req/s * 5000 connections = 1600000` so 1.6 Million requests per second over 5k db connections for `/v3/security_groups/:guid`. This of course wont be possible since the db could never serve so many transaction(`1000 db transactions per second * 5k connections = 5 Million`). But this indicates that for the PoC implementation, the number of db connections should not be a limit anymore rather VM and DB performance.
### Raw Measurements

All results can be found in below table.

<table>
  <tbody>
    <tr>
      <td align="center">Go Implementation</td>
    </tr>
    <tr>
      <td align="center">
        GET /v3/info
        <ul>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-info_8.html">v3/info - 8 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-info_16.html">v3/info - 16 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-info_24.html">v3/info - 24 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-info_32.html">v3/info - 32 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-info_64.html">v3/info - 64 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-info_128.html">v3/info - 128 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-info_256.html">v3/info - 256 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-info_512.html">v3/info - 512 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-info_1024.html">v3/info - 1024 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-info_2048.html">v3/info - 2048 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-info_4096.html">v3/info - 4096 req/s</a> </li>
        </ul>
        GET /v3/buildpacks
        <ul>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-buildpacks_8.html">v3/buildpacks - 8 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-buildpacks_16.html">v3/buildpacks - 16 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-buildpacks_24.html">v3/buildpacks - 24 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-buildpacks_32.html">v3/buildpacks - 32 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-buildpacks_64.html">v3/buildpacks - 64 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-buildpacks_128.html">v3/buildpacks - 128 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-buildpacks_256.html">v3/buildpacks - 256 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-buildpacks_512.html">v3/buildpacks - 512 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-buildpacks_1024.html">v3/buildpacks - 1024 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-buildpacks_2048.html">v3/buildpacks - 2048 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-buildpacks_4096.html">v3/buildpacks - 4096 req/s</a> </li>
        </ul>
        GET /v3/security_groups/:guid
        <ul>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-security_groups_8.html">v3/security_groups - 8 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-security_groups_16.html">v3/security_groups - 16 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-security_groups_24.html">v3/security_groups - 24 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-security_groups_32.html">v3/security_groups - 32 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-security_groups_64.html">v3/security_groups - 64 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-security_groups_128.html">v3/security_groups - 128 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-security_groups_256.html">v3/security_groups - 256 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-security_groups_512.html">v3/security_groups - 512 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-security_groups_1024.html">v3/security_groups - 1024 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-security_groups_2048.html">v3/security_groups - 2048 req/s</a> </li>
          <li><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/html/v3-security_groups_4096.html">v3/security_groups - 4096 req/s</a> </li>
        </ul>
      </td>
    </tr>
    <tr>
      <td align="center"><a target="_parent" href="/loadtests/v1/cc_go_load_test_single_db_connection/cc_go_load_test_single_db_connection.tar.xz">Raw Archive</a></td>
    </tr>
  </tbody>
</table>

## Viewing Raw Results
The the raw measurements tarballs linked above include:
- `vegeta.log` Terminal output as text file ([#vegetalog](#vegetalog))
- `netdata.snapshot` Vm and Database metrics during run as [netdata snapshot](https://learn.netdata.cloud/docs/dashboard/import-export-print-snapshot) ([#netdata-snapshot](#netdata-snapshot))
- `v3-.*_[0-9]*.bin` Binary [vegeta](https://github.com/tsenart/vegeta) files ([#binary-vegeta-files](#binary-vegeta-files))
- `html folder` Pre rendered analysis from the binary files ([#html-folder](#html-folder))


### Vegeta.log
Its a text file with (utc)timestamps when the runs where executed. You can use these timestamps to correlate metrics from the netdata snapshot to individual runs.

Example:
```bash
...
Targeting: https://10.0.65.75/v3/info
2021-11-03 07:56:07.314080 - Using attack rate: 8 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/info"}
2021-11-03 07:56:57.638468 - Using attack rate: 16 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/info"}
2021-11-03 07:57:47.913230 - Using attack rate: 24 - {"body": "", "header": {"Authorization": "REDACTED"}, "method": "GET", "url": "https://10.0.65.75/v3/info"}
...
```

### Netdata Snapshot
Netdata is a system monitoring tool one can [start as container locally](https://learn.netdata.cloud/docs/agent/packaging/docker).
You can then view the tool at http://localhost:9090.
It will monitor your local machine, but you can [import a previously made recording](https://learn.netdata.cloud/docs/dashboard/import-export-print-snapshot).
![](https://user-images.githubusercontent.com/1153921/114218399-360fb600-991e-11eb-8dea-fabd2bffc5b3.gif)
The file to use is `netdata.snapshot` inside the raw results archive.

### Binary Vegeta Files
The results archive also includes raw binary dumps of [vegeta](https://github.com/tsenart/vegeta).
The files have the `.bin` file ending. The name represents the tested endpoint as well as the load(requests per second).

E.g. `v3-buildpacks_32.bin` is a test of the `/v3/buildpacks` endpoint with `32 requests per second`.

These files can be piped to [vegeta](https://github.com/tsenart/vegeta) for ploting and analysis. See: https://github.com/tsenart/vegeta#report-command 

### HTML Folder
The HTML folder includes the output of `vegeta report` as well as `vegeta plot` commands combined in one viewable html page that can be opened in your browser.