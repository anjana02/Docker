 What's Docker?
=======================

Docker is an open platform for developing,shipping,and running applications.
-------------


 Docker is a tool that allows developers, sys-admins etc. to easily deploy their applications in a sandbox (called containers) to run on the host operating system i.e. Linux. The key benefit of Docker is that it allows users to package an application with all of its dependencies into a standardized unit for software development.


what are containers?
==================

Containers, by contrast, isolate applications execution environments from one another, but share the underlying OS kernel.

They are almost measured in mb.

Use fewer resource than VMs.

And also it starts immediately.


What are images?
================

A Docker image is a file, comprised of multiple layers, that is used to execute code in a Docker container. An image is essentially built from the instructions for a complete and executable version of an application, which relies on the host OS kernel. When the Docker user runs an image, it can become one or multiple instances of that container.

Install Docker
==============
Follow these according to your OS : <br>
<a href="https://docs.docker.com/docker-for-mac/install/ "  target="_blank">Mac</a> <br>
<a href="https://docs.docker.com/engine/install/ubuntu/" target="_blank">Linux</a> <br>
<a href="https://docs.docker.com/docker-for-windows/install/" target="_blank">Windows</a> <br>

Once you are done installing Docker, test your Docker installation by running the following:
--------------------------------


$ docker run hello-world
-----------------------

Hello from Docker.
This message shows that your installation appears to be working correctly.
...

Docker Command
==============
The base command for the Docker CLI.

<table>
<thead>
  <tr>
    <td>Command</td>
    <td>Description</td>
  </tr>
</thead>
<tbody>

  
  <tr>
    <td><a href="/engine/reference/commandline/attach/">docker attach</a></td>
    <td>Attach local standard input, output, and error streams to a running container</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/build/">docker build</a></td>
    <td>Build an image from a Dockerfile</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/builder/">docker builder</a></td>
    <td>Manage builds</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/checkpoint/">docker checkpoint</a></td>
    <td>Manage checkpoints</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/commit/">docker commit</a></td>
    <td>Create a new image from a container’s changes</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/config/">docker config</a></td>
    <td>Manage Docker configs</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/container/">docker container</a></td>
    <td>Manage containers</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/context/">docker context</a></td>
    <td>Manage contexts</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/cp/">docker cp</a></td>
    <td>Copy files/folders between a container and the local filesystem</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/create/">docker create</a></td>
    <td>Create a new container</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/diff/">docker diff</a></td>
    <td>Inspect changes to files or directories on a container’s filesystem</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/events/">docker events</a></td>
    <td>Get real time events from the server</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/exec/">docker exec</a></td>
    <td>Run a command in a running container</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/export/">docker export</a></td>
    <td>Export a container’s filesystem as a tar archive</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/history/">docker history</a></td>
    <td>Show the history of an image</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/image/">docker image</a></td>
    <td>Manage images</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/images/">docker images</a></td>
    <td>List images</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/import/">docker import</a></td>
    <td>Import the contents from a tarball to create a filesystem image</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/info/">docker info</a></td>
    <td>Display system-wide information</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/inspect/">docker inspect</a></td>
    <td>Return low-level information on Docker objects</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/kill/">docker kill</a></td>
    <td>Kill one or more running containers</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/load/">docker load</a></td>
    <td>Load an image from a tar archive or STDIN</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/login/">docker login</a></td>
    <td>Log in to a Docker registry</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/logout/">docker logout</a></td>
    <td>Log out from a Docker registry</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/logs/">docker logs</a></td>
    <td>Fetch the logs of a container</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/manifest/">docker manifest</a></td>
    <td>Manage Docker image manifests and manifest lists</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/network/">docker network</a></td>
    <td>Manage networks</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/node/">docker node</a></td>
    <td>Manage Swarm nodes</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/pause/">docker pause</a></td>
    <td>Pause all processes within one or more containers</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/plugin/">docker plugin</a></td>
    <td>Manage plugins</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/port/">docker port</a></td>
    <td>List port mappings or a specific mapping for the container</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/ps/">docker ps</a></td>
    <td>List all running containers</td>
  </tr>
  
  <tr>
    <td><a href="/engine/reference/commandline/ps/">docker ps -a </a></td>
    <td>List all containers</td>
  </tr>
  
  <tr>
    <td><a href="/engine/reference/commandline/pull/">docker pull</a></td>
    <td>Pull an image or a repository from a registry</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/push/">docker push</a></td>
    <td>Push an image or a repository to a registry</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/rename/">docker rename</a></td>
    <td>Rename a container</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/restart/">docker restart</a></td>
    <td>Restart one or more containers</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/rm/">docker rm</a></td>
    <td>Remove one or more containers</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/rmi/">docker rmi</a></td>
    <td>Remove one or more images</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/run/">docker run</a></td>
    <td>Run a command in a new container</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/save/">docker save</a></td>
    <td>Save one or more images to a tar archive (streamed to STDOUT by default)</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/search/">docker search</a></td>
    <td>Search the Docker Hub for images</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/secret/">docker secret</a></td>
    <td>Manage Docker secrets</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/service/">docker service</a></td>
    <td>Manage services</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/stack/">docker stack</a></td>
    <td>Manage Docker stacks</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/start/">docker start</a></td>
    <td>Start one or more stopped containers</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/stats/">docker stats</a></td>
    <td>Display a live stream of container(s) resource usage statistics</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/stop/">docker stop</a></td>
    <td>Stop one or more running containers</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/swarm/">docker swarm</a></td>
    <td>Manage Swarm</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/system/">docker system</a></td>
    <td>Manage Docker</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/tag/">docker tag</a></td>
    <td>Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/top/">docker top</a></td>
    <td>Display the running processes of a container</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/trust/">docker trust</a></td>
    <td>Manage trust on Docker images</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/unpause/">docker unpause</a></td>
    <td>Unpause all processes within one or more containers</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/update/">docker update</a></td>
    <td>Update configuration of one or more containers</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/version/">docker version</a></td>
    <td>Show the Docker version information</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/volume/">docker volume</a></td>
    <td>Manage volumes</td>
  </tr>

  
  <tr>
    <td><a href="/engine/reference/commandline/wait/">docker wait</a></td>
    <td>Block until one or more containers stop, then print their exit codes</td>
  </tr>

</tbody>
</table>




 

 
