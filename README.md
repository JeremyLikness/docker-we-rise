# Contain Your Excitement: Hands on Docker 

This is a presentation that was written for the [We Rise Women in Tech Conference](https://werise.tech/). It provides a series of hands-on examples for walking through your first Docker container. 

## Getting Started 

First, install [Docker](https://docker.com) then watch the [deck](./ContainYourExcitement.pptx) to get familiar with Docker. 

### First Step: Hello, World 

Type `docker run hello-world` and read the output. Congratulations! You just pulled your first Docker image and ran it.

### Second Step: Web Server 

Let's do something a little more advanced. In the `00-Hello` directory there is a single HTML file and a [Dockerfile](https://docs.docker.com/engine/reference/builder/) definition. The Dockerfile contains a set of instructions to build an image. In this case, it will use the [nginx](https://hub.docker.com/_/nginx/) webserver image, copy the HTML file, and create a self-contained webserver. 

Build the container: `docker build -t we-rise .` 

Then run it: `docker run -d -p 80:80 we-rise` 

Navigate to [http://localhost:80](http://localhost:80) in your browser to see it running. Now list the running processes:

`docker ps` 

You will see a name for the process that was assigned by the docker image, usually two words separated by an underscore. If, for example, the name is `commodore_breadbox` you can stop the running container:

`docker stop commodore_breadbox` 

Then remove it (this does not delete the image). 

`docker rm commodore_breadbox` 

Now look at the size of your image. 

Windows: `docker images | find "we-rise"`

Unix: `docker images | grep "we-rise"` 

Pretty large for a single file. Let's see if we can do better.

### Third Step: Tiny Web Server 

Navigate to `01-Hello-Small` and build the container: 

`docker build -t we-rise-small .` 

Run it and confirm it is working. This container uses the HTTP daemon that is part of [busybox](https://hub.docker.com/_/busybox/), an extremely small distribution. See for yourself: 

`docker images | grep "we-rise-small"` 

Now that's pretty handy. Let's take it another level. 

### Fourth Step: Go Webserver 

Change to `02-Hello-Small-Go\web` and rebuild the `we-rise-small` image with the same command from the previous step. This file now looks for a script at port 8080 on localhost. Of course, the script isn't there, so navigate up one level to `02-Hello-Small-Go` and build the service: 

`docker build -t we-rise-svc .` 

Even if you don't have [Go](https://golang.org/) installed, the image has everything needed to take the `webserver.go` file and build a `Go` app inside the container. Run it: 

`docker run -d -p 8080:8080 we-rise-svc` 

You can browse to the service directly at [http://localhost:8080](http://localhost:8080) and see it in action by refreshing the web page that is running from the previous step. The text `Greetings from service ???` should change to a number, indicating the web app has now successfully interacted with the service.

Be sure to stop and remove your containers to clean up! 

### Fifth Step: Go Small 

There are two ways to "go small." The first, if you are running version 17.05 or later, is to use the multi-stage build. Run:

`docker build -t we-rise-svc -f Dockerfile.multi .` 

If you have an older version of Docker, you can build locally. This step requires [installing Go](https://golang.org/dl/). There is a provided script in `03-Hello-Small-Go-Small` that will build the Go binary, then use the Dockerfile to create a small image by using the [scratch](https://hub.docker.com/_/scratch/) image which is empty, then adding the Go binary. You can see after running the script from a bash shell (if you are on Windows you can use the one that comes with [git](https://git-scm.com/downloads)) like this: `./gobuild.sh`

After either approach, you can check out the size of the new service:

`docker ps | grep "we-rise-svc"` 

Now that's small! We have a client/server web app that is a 1 megabyte image and a 6 megabyte image! 

### Last Step: Compose 

Stop and clean up your containers, then try something a little different. The entire app can be run with a simple command:

`docker-compose up` 

This uses the `docker-compose.yml` file to orchestrate several services. Once it is up and running you can navigate to your local host to confirm it works. Now open another command line, navigate to the same directory and run this command: 

`docker-compose scale gosvc=4` 

You'll see several new images starting. Give them a few seconds to start, then refresh your web page. You'll see the number change. This is because we included a proxy that automatically load balances across multiple instances of the service. Pretty powerful, no?

Type `docker-compose down` to tear down the containers and networks. 

Thanks! For more tutorials and information like this, follow me on Twitter: [@JeremyLikness](https://twitter.com/JeremyLikness). 

![Jeremy Likness](http://jeremylikness.com/signature.gif)