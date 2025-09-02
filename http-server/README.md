# Http Server 101

A basic http server should be capable of the following tasks:-
* Process dynamic requests:- Process incoming requests from users who browse the website, log into their accounts or post images.
* Serve static assets:- Server js, css and images to browsers to create a dynamic experience for the user.
* Accept Connections:- The Http server must listen on a specific port to be able to accept connections from the internet.

# Networking 101

The network (or internet) operates using IP addresses and ports.
* IP addresses:- These are used to identify a specific device on the network(like your computer or a web browser)
* Port:- Identifies a specific application or service running on that device. Ports allow multiple programs on the same device to communicate over the network without interference.
    - Each network service(like an HTTP server, FTP server, etc) listens on a specific port for incoming traffic. The port number tells the operating system where to direct network packets.
    - Clients (like web browsers) send requests to the server by specifying both its IP address and port number. For example, when you type `http:example.com`, your browser connects to the server's IP on port 80 (the default HTTP)
    - Without listening on a specific port, the operating system wouldn't know where to forward the incoming network traffic. The port identifies which application on the machine should handle the incoming connection.

## Binding to specific IPs

* It refers to the process where an application (like a web server) is instructed to listen for incoming connections only on a specific network interface or IP address and port (of the machine it is running on).
* Every machine has one or more network interfaces, each associated with an IP address (or multiple IPs, in some cases). By binding to a particular IP, the server limits where it will accept connections from.

* Network Interfaces:- This is the 'software representation' of a network connection. Each interface is typically linked to an IP address. For example a machine might have:
    - localhost(loopback interface): This is typically represented by the IP `127.0.0.1`, which refers to local machine itself.
    - LAN IP (Local Area Network):- For communication with a local network (like your home or office network). For example: 192.168.1.10.
    - Public IP (Internet Facing IP):- This is the IP address visible to the outside world, assigned by an ISP(Internet Service Provider).

### Binding to all interfaces

* If you want your server to accept connections on any network interface, you can bind it to a special IP address called `0.0.0.0`. This IP doesn't represent a real address but tells the operating system to accept connections on all available interfaces (both local and public)

### Binding to multiple IPs

* It is also possible to bind to multiple specific IP addresses (though this requires more advanced networking configurations). This can be useful if you have multiple network interfaces and want to control which interfaces handle which requests.

* Following are the various scenarios when its useful:-
    - Multiple Network Interfaces
    - Virtual IPs
    - IPv4 and IPv6 addresses
    - Hosting Multiple Services or Websites
    - Public IP pool or subnet
    - IP Aliasing

#### Multiple Network Interfaces

* A machine can have more than 1 network interface card (NIC), and each NIC can be assigned a different public IP address. Each NIC would typically be connected to a different network or internet service provider (ISP). 

* Example 1:- A web server might have one NIC connected to one ISP and another NIC connected to a different ISP for redundancy(if one network goes down, the other can still handle traffic).

* Example 2:- A web hosting provider might assign multiple public IP addresses to a single server. This can allow the server to host multiple SSL certificates(before SNI, each SSL certificate needed its own IP address) or offer dedicated IPs for clients.

#### Virtual IPs (VIPs)

* A machine can have multiple IP addresses assigned to a single NIC using virtual IPs. This allows a single physical interface to respond to multiple public IPs.

* VIPs are commonly used in high-availability(HA) or load balancing scenarios. For example, a load balancer or reverse proxy might have several public IP addresses to manage traffic to multiple services.  

* Example:- A load balancer might use multiple public IPs to distribute traffic between different servers or services. This helps in distributing the traffic more effectively and increases fault tolerance.

#### IPv4 and IPv6

* A machine can simulataneously have both IPv4 and IPv6 addresses. This dual stack configuration is becoming more common as IPv6 adoption grows. Both of these IPs can be public and the machine can accept traffic on both.

* Example:- Security and Isolation. In such a case, different services on a server might be assigned different public IP addresses to isolated for security purposes. For instance, a database service might have one public IP with strict firewall rules, while a public-facing web service might have another IP with more open access.

#### Hostig multiple services or websites 

* When a machine is hosting multiple websites or services, it may be beneficial or necessary to to assign a different IP public address to each service.

* SNI(Server Name Indication) is used here over SSL, which is an extension of TLS protocol. It allows the client to specify the hostname of the server it is trying to connect during the initial SSL/TLS handshake.

#### Public IP Pool or Subnet

* Some organisations or service providers own a pool of public IP addresses. They can assign multiple public IP addresses to a single machine to handle different services (eg, a web server, email server or file server on different public IPs) 

* Example:- This is common in datacenters, where each machine might be assigned a set of public IP addresses from the available pool.

#### IP Aliasing

* it is a method by which a single network interface is assigned multiple IP addresses, often on the same subnet. This can be used for things like hosting different services on different public IPs, or for traffic management purposes.
