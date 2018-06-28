# Terms I've come across in Distributed Systems

**Availability** - the proportion of time a system is in a functioning condition. If a user cannot access the system, it is said to be unavailable.

**Fault Tolerance** - ability of a system to behave in a well-defined manner once faults occur

**latency** - The state of being latent; delay, a period between the initiation of something and the occurrence. (**latent** From Latin latens, latentis, present participle of lateo ("lie hidden"). Existing or present but concealed or inactive.)
latency is really the time between when something happened and the time it has an impact or becomes visible.
For example, imagine that you are infected with an airborne virus that turns people into zombies. The latent period is the time between when you became infected, and when you turn into a zombie. That's latency: the time during which something that has already happened is concealed from view.

**difference between error and anamoly** - an error is incorrect behavior, while an anomaly is unexpected behavior. If you were smarter, you'd expect the anomalies to occur.

**column oriented DBMS** -  A column-oriented DBMS"(or columnar database management system) is a database management system (DBMS) that stores data tables by column rather than by row.

**Network partition** - A network partition refers to network decomposition into relatively independent subnets for their separate optimization as well as network split due to the failure of network devices. In both cases the partition-tolerant behavior of subnets is expected. This means that even after network is partitioned into multiple sub-systems, it still works correctly.

For example, in a network with multiple subnets where nodes A and B are located in one subnet and nodes C and D are in another, a partition occurs if the switch between the two subnets fails. In that case nodes A and B can no longer communicate with nodes C and D, but all nodes A-D work the same as before.

**Difference between when partitioning and replication is used** - You partition data over multiple nodes because it makes it more conducive for parallel processing.
 Data is *replicated* over multiple nodes to reduce the distance (and potential latency) between that data and their compute nodes/ or to reduce the latency between the client and the server. This allows more servers to take part in the computation. It improves availability because now, more nodes need to fail in order for data to become unavailable. However, since there are now independent copies of the data that has to be kept in sync on multiple machines - this means ensuring that the replication follows a consistency model. Strength of a consistency model is inversely proportion to latency.



