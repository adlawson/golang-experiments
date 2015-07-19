# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|

    go_version = "1.4.2"
    go_path = "/srv"

    config.vm.box = "ubuntu/trusty64"

    config.vm.synced_folder ".", "#{go_path}/src/github.com/adlawson/golang-experiments"

    config.vm.provision :shell, :inline => "touch .hushlogin"
    config.vm.provision :shell, :inline => "locale-gen en_GB.UTF.8"
    config.vm.provision :shell, :inline => "apt-get update --fix-missing"
    config.vm.provision :shell, :inline => "apt-get install -q -y g++ make git curl vim"

    config.vm.provision :shell, :inline => "mkdir -p #{go_path}/bin #{go_path}/pkg"
    config.vm.provision :shell, :inline => "echo 'fetching go#{go_version}.linux-amd64' && wget -q https://storage.googleapis.com/golang/go#{go_version}.linux-amd64.tar.gz"
    config.vm.provision :shell, :inline => "tar -C /usr/local -xzf go#{go_version}.linux-amd64.tar.gz"
    config.vm.provision :shell, :inline => "echo 'export PATH=$PATH:/usr/local/go/bin:#{go_path}/bin' > .bash_profile"
    config.vm.provision :shell, :inline => "echo 'export GOPATH=#{go_path}' >> .bash_profile"
    config.vm.provision :shell, :inline => "chown -R vagrant:vagrant #{go_path}"

end
