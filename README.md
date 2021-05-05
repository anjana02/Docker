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
    <td>docker attach</td>
    <td>Attach local standard input, output, and error streams to a running container</td>
  </tr>

  
  <tr>
    <td>docker build</td>
    <td>Build an image from a Dockerfile</td>
  </tr>

  
  <tr>
    <td> docker builder</td>
    <td>Manage builds</td>
  </tr>

  
  <tr>
    <td> docker checkpoint</td>
    <td>Manage checkpoints</td>
  </tr>

  
  <tr>
    <td> docker commit</td>
    <td>Create a new image from a container’s changes</td>
  </tr>

  
  <tr>
    <td> docker config</td>
    <td>Manage Docker configs</td>
  </tr>

  
  <tr>
    <td>docker container</td>
    <td>Manage containers</td>
  </tr>

  
  <tr>
    <td> docker context</td>
    <td>Manage contexts</td>
  </tr>

  
  <tr>
    <td> docker cp </td>
    <td>Copy files/folders between a container and the local filesystem</td>
  </tr>

  
  <tr>
    <td>docker create </td>
    <td>Create a new container</td>
  </tr>

  
  <tr>
    <td> docker diff</td>
    <td>Inspect changes to files or directories on a container’s filesystem</td>
  </tr>

  
  <tr>
    <td> docker events</td>
    <td>Get real time events from the server</td>
  </tr>

  
  <tr>
    <td>docker exec </td>
    <td>Run a command in a running container</td>
  </tr>

  
  <tr>
    <td> docker export </td>
    <td>Export a container’s filesystem as a tar archive</td>
  </tr>

  
  <tr>
    <td>docker history </td>
    <td>Show the history of an image</td>
  </tr>

  
  <tr>
    <td>ocker image</td>
    <td>Manage images</td>
  </tr>

  
  <tr>
    <td> docker images</td>
    <td>List images</td>
  </tr>

  
  <tr>
    <td> docker import</td>
    <td>Import the contents from a tarball to create a filesystem image</td>
  </tr>

  
  <tr>
    <td> docker info </td>
    <td>Display system-wide information</td>
  </tr>

  
  <tr>
    <td>docker inspect</td>
    <td>Return low-level information on Docker objects</td>
  </tr>

  
  <tr>
    <td>docker kill</td>
    <td>Kill one or more running containers</td>
  </tr>

  
  <tr>
    <td>docker load </td>
    <td>Load an image from a tar archive or STDIN</td>
  </tr>

  
  <tr>
    <td>docker login </td>
    <td>Log in to a Docker registry</td>
  </tr>

  
  <tr>
    <td>docker logout </td>
    <td>Log out from a Docker registry</td>
  </tr>

  
  <tr>
    <td>docker logs </td>
    <td>Fetch the logs of a container</td>
  </tr>

  
  <tr>
    <td>docker manifest </td>
    <td>Manage Docker image manifests and manifest lists</td>
  </tr>

  
  <tr>
    <td> docker network </td>
    <td>Manage networks</td>
  </tr>

  
  <tr>
    <td>docker node </td>
    <td>Manage Swarm nodes</td>
  </tr>

  
  <tr>
    <td> docker pause </td>
    <td>Pause all processes within one or more containers</td>
  </tr>

  
  <tr>
    <td>docker plugin </td>
    <td>Manage plugins</td>
  </tr>

  
  <tr>
    <td>docker port </td>
    <td>List port mappings or a specific mapping for the container</td>
  </tr>

  
  <tr>
    <td> docker ps</td>
    <td>List all running containers</td>
  </tr>
  
  <tr>
    <td>docker ps -a </td>
    <td>List all containers</td>
  </tr>
  
  <tr>
    <td>docker pull </td>
    <td>Pull an image or a repository from a registry</td>
  </tr>

  
  <tr>
    <td> docker push </td>
    <td>Push an image or a repository to a registry</td>
  </tr>

  
  <tr>
    <td> docker rename </td>
    <td>Rename a container</td>
  </tr>

  
  <tr>
    <td> docker restart </td>
    <td>Restart one or more containers</td>
  </tr>

  
  <tr>
    <td>docker rm </td>
    <td>Remove one or more containers</td>
  </tr>

  
  <tr>
    <td>docker rmi </td>
    <td>Remove one or more images</td>
  </tr>

  
  <tr>
    <td> docker run </td>
    <td>Run a command in a new container</td>
  </tr>

  
  <tr>
    <td> docker save </td>
    <td>Save one or more images to a tar archive (streamed to STDOUT by default)</td>
  </tr>

  
  <tr>
    <td>docker search</td>
    <td>Search the Docker Hub for images</td>
  </tr>

  
  <tr>
    <td>docker secret</td>
    <td>Manage Docker secrets</td>
  </tr>

  
  <tr>
    <td>docker service</td>
    <td>Manage services</td>
  </tr>

  
  <tr>
    <td>docker stack</td>
    <td>Manage Docker stacks</td>
  </tr>

  
  <tr>
    <td>docker start</td>
    <td>Start one or more stopped containers</td>
  </tr>

  
  <tr>
    <td>docker stats</td>
    <td>Display a live stream of container(s) resource usage statistics</td>
  </tr>

  
  <tr>
    <td>docker stop</td>
    <td>Stop one or more running containers</td>
  </tr>

  
  <tr>
    <td> docker swarm</td>
    <td>Manage Swarm</td>
  </tr>

  
  <tr>
    <td> docker system</td>
    <td>Manage Docker</td>
  </tr>

  
  <tr>
    <td> docker tag</td>
    <td>Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE</td>
  </tr>

  
  <tr>
    <td> docker top</td>
    <td>Display the running processes of a container</td>
  </tr>

  
  <tr>
    <td> docker trust</td>
    <td>Manage trust on Docker images</td>
  </tr>

  
  <tr>
    <td>  docker unpause</td>
    <td>Unpause all processes within one or more containers</td>
  </tr>

  
  <tr>
    <td> docker update</td>
    <td>Update configuration of one or more containers</td>
  </tr>

  
  <tr>
    <td> docker version</td>
    <td>Show the Docker version information</td>
  </tr>

  
  <tr>
    <td> docker volume</td>
    <td>Manage volumes</td>
  </tr>

  
  <tr>
    <td>docker wait</td>
    <td>Block until one or more containers stop, then print their exit codes</td>
  </tr>

</tbody>
</table>




 

 
