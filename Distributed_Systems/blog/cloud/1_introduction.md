# Introduction to Cloud Computing

I've been wanting to share what as I explore the field of distributed systems and what better way to start with a blog post about clouds.
The idea of this post is to present the pre-requisite knowledge needed to come up with the working definition of a cloud.

## How did it all begin?
In the 1950s, even before the primitive ideas of a distributed system were implemented in the form of [ARPANET](https://en.wikipedia.org/wiki/ARPANET), users would connect to mainframes using terminals, not needing any processing power of their own. These mainframes were utilised to perform bulk data processing by large organizations for applications like nationwide census, consumer statistics and processing transactions. The downside to mainframes is that they are expensive and extremely difficult to set up. Therefore, most universities and research centres had just one mainframe. This can be considered the predecessor to the cloud. The mainframe also gave birth to the concept of  [Virtualization](https://en.wikipedia.org/wiki/Virtualization) where the resources of the mainframe, like the computing elements and memory were logically divided and distributed among users where jobs could be submitted along with creating a new instance of the operating system for each user. Before virtualization, the mainframe could perform only one job at a time. Virtualization allowed the concept of [Time-Sharing](https://en.wikipedia.org/wiki/Time-sharing) which basically allows a large number of users to concurrently interact with a single computer.
 
![an IBM Mainframe from 1964](https://upload.wikimedia.org/wikipedia/commons/7/7d/IBM_704_mainframe.gif)
An IBM Mainframe from 1964

Further, in the mid 1980s, SoftPC was created which is a software emulator of x86 hardware. It allowed DOS applications to be run on UNIX machines, making development a whole lot less expensive. VMware was founded around this time as well and sold the VMware workstation which was later evolved into the ESX server. The ESX server did not require a host operating system to run on, it could run on pure hardware which allowed it to host many instnaces of virtual machines with better perfromance. 

Today, a "cloud" has three important resources/aspects toward it:
1. Computing Resources
2. Storage
3. Network

Clouds today are data oriented and can be thought of as 

lots of storage + compute cycles nearby.

The predecessors of clouds that I've talked about above, worked exactly the other way around.

Today's clouds have four main
components:

1. Massive Scale
The implementations of the early ideas of distributed systems in the 20th century did not really scale up. Clouds on the other hand were designed to scale up massively. As of today, Microsoft's OneDrive caters to 55 million active users per month!
2. No Upfront Commitment
As I've mentioned above, if you wanted to use a mainframe you had to physically buy an entire one or pay an extremely high amount to be a lifetime user. Today's cloud solutions allow you to use resources on a per use basis and even allow you to specify the level of each resource. Clouds provide different resources *as a service*. For example, **HaaS**(Hardware as a Service) gives you only hardware resources and no form of security, **IaaS**(Infrastructure as a Service) provides you with the hardware resources along with operating system images which makes it easier to get started with using the cloud.
3. Data intensive
Unlike its predecessors, clouds are data oriented and are used either as stores for large volumes or data or are used for bulk data processing. The idea if to group data into [nodes](https://en.wikipedia.org/wiki/Computer_cluster) that are close to each other as much as possible to reduce latency during processing.
4. New Cloud Computing Paradigms
I haven't researched enough about this topic to talk about it but one paradigm that I have studied is called MapReduce. It's a technique that's used for processing a lot of data. It is broken up into two phases, the Map phase and the Reduce phase. The Map phase processes the data and generates key-value pairs similar to a Python dictionary with the value being the number of occurences of the keys. However, it does not group elements with same keys together, so from what I've understood, the value corresponding to all keys is one. What the Reduce phase does is it groups similar keys together and feeds each group to a compute node for processing. The logic for "grouping" keys is handled by a partitioning function. I will explain MapReduce in detail in the next blog post.


-> TODO Arrive at a clear definition of what a cloud is today. Need to add some intermediate information
