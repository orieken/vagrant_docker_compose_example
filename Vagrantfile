# -*- mode: ruby -*-
# vi: set ft=ruby :


Vagrant.configure(2) do |config|

  config.vm.box = "ubuntu/trusty64"

  config.vm.network(:forwarded_port, guest: 8080, host: 8080)

  config.vm.provision :docker
  config.vm.provision :docker_compose, yml: "/vagrant/docker-compose.yml", rebuild: true, project_name: "myproject", run: "always"

end
