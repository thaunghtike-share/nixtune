data "amazon-ami" "ubuntu" {
  filters = {
    virtualization-type = "hvm"
    name                = "ubuntu/images/*ubuntu-focal-20.04-amd64-server-*"
    root-device-type    = "ebs"
  }
  owners      = ["099720109477"]
  most_recent = true
}

data "amazon-ami" "debian" {
  filters = {
    virtualization-type = "hvm"
    name                = "debian-11-amd64-*"
    root-device-type    = "ebs"
  }
  owners      = ["136693071363"]
  most_recent = true
}

data "amazon-ami" "redhat" {
  filters = {
    virtualization-type = "hvm"
    name                = "RHEL-9.0.0_HVM-20221027-x86_64-1-Hourly2-GP2"
    root-device-type    = "ebs"
    architecture        = "x86_64"
  }
  owners      = ["309956199498"]
  most_recent = true
}

source "amazon-ebs" "ubuntu" {
  ami_name      = "nixtune-ubuntu"
  instance_type = "t2.micro"
  region        = "us-east-1"
  source_ami    = data.amazon-ami.ubuntu.id
  ssh_username  = "ubuntu"
}

source "amazon-ebs" "debian" {
  ami_name      = "nixtune-debian"
  instance_type = "t2.micro"
  region        = "us-east-1"
  source_ami    = data.amazon-ami.debian.id
  ssh_username  = "admin"
}

source "amazon-ebs" "redhat" {
  ami_name      = "nixtune-redhat"
  instance_type = "t2.micro"
  region        = "us-east-1"
  source_ami    = data.amazon-ami.redhat.id
  ssh_username  = "ec2-user"
}

build {
  sources = [
    "source.amazon-ebs.ubuntu",
    "source.amazon-ebs.debian",
    "source.amazon-ebs.redhat"
  ]

  provisioner "shell" {
    script = "./setup.sh"
  }
}
