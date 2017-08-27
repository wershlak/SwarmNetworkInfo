# Swarm Network Info

I just hacked this together to give me a good understanding of how the swarm networks were working under the hood.

Here's an example from my swarm test environment

```
+------------+-------------------+---------+-----------------------------------------------------+-------------+---------------+
|    NODE    |      SERVICE      | STATUS  |                        IMAGE                        |   NETWORK   |    ADDRESS    |
+------------+-------------------+---------+-----------------------------------------------------+-------------+---------------+
| swarmtest1 | elk_logspout      | running | bekt/logspout-logstash:latest                       | elk_default | 10.0.0.7/24   |
| swarmtest3 | elk_kibana        | running | docker.elastic.co/kibana/kibana:5.3.2               | ingress     | 10.255.0.6/16 |
|            |                   |         |                                                     | elk_default | 10.0.0.11/24  |
| swarmtest1 | elk_elasticsearch | running | docker.elastic.co/elasticsearch/elasticsearch:5.3.2 | elk_default | 10.0.0.12/24  |
| swarmtest3 | elk_logspout      | running | bekt/logspout-logstash:latest                       | elk_default | 10.0.0.8/24   |
| swarmtest1 | elk_logstash      | running | docker.elastic.co/logstash/logstash:5.3.2           | elk_default | 10.0.0.9/24   |
| swarmtest2 | elk_logspout      | running | bekt/logspout-logstash:latest                       | elk_default | 10.0.0.5/24   |
+------------+-------------------+---------+-----------------------------------------------------+-------------+---------------+
```

It shows each swarm task (container) and all the network attachments and IP Addresses
