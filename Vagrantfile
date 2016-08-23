Vagrant.configure(2) do |config|
  config.vm.box = "Ubuntu14"
  config.vm.network :forwarded_port, guest: 9000, host: 9000
  config.vm.network "private_network", ip: "192.168.33.10"
  config.vm.synced_folder "./", "/vagrant", :nfs => true
  config.vm.provider "virtualbox" do |vb|
    vb.memory = "1024"
  end
end
