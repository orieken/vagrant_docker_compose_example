# vagrant_docker_compose_example
example with vagrant and docker containers 




## Parts

## Install on OsX

* docker

``` 
brew install docker
```

* boot2docker

```
brew install boot2docker
```

* docker-compose

```
brew install docker-compose
```

* Vagrant

```
brew cask install vagrant
```

* vagrant-manager

``` 
brew cask install vagrant-manager
```

* vagrant-docker-compose plugin (I got an error the first time so i ran it twice and worked fine)

```
vagrant plugin install vagrant-docker-compose
```

## Running

* Start boot2docker

```
boot2docker up
```

* Set all of the output

```
export DOCKER_HOST=
export DOCKER_CERT_PATH=
export DOCKER_TLS_VERIFY=
```

* run

```
docker-compose build
docker-compose up
```

* get the ip of boot2docker

```
boot2docker ip
```

* check app
* boot2docker ip the port we opened on the container


## GO DEPS

* http://labix.org/mgo
* https://godoc.org/github.com/julienschmidt/httprouter