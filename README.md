# capdns

capdns is a network capture utility designed specifically for DNS traffic. This utility is based on tcpdump.

Some of its features include:

- Understands both IPv4 and IPv6
- Captures UDP, TCP, and IP fragments.

## problem background
In the dns test, packet capture is a common method, but dns requests are very frequent, which interferes a lot with the packet capture results. Sometimes it is necessary to only capture packages related to a specific domain name.

## Dependencies
To install the dependencies under CentOS

`yum -y install tcpdump`

## Instructions
`~./capdns -domain www.infvie.com (default "www.infvie.com")`


![Image text](https://mirrors.infvie.org//image/capdns/20220209125233.png)

![Image text](https://mirrors.infvie.org//image/capdns/20220209125149.png)

## Inspiration
https://github.com/DNS-OARC/dnscap