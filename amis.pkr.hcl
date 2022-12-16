data "amazon-ami" "ubuntu_x86" {
  filters = {
    virtualization-type = "hvm"
    name                = "ubuntu/images/*ubuntu-kinetic-22.10-amd64-server-*"
    root-device-type    = "ebs"
    architecture        = "x86_64"
  }
  owners      = ["099720109477"]
  most_recent = true
}

data "amazon-ami" "ubuntu_arm" {
  filters = {
    virtualization-type = "hvm"
    name                = "ubuntu/images/*ubuntu-kinetic-22.10-amd64-server-*"
    root-device-type    = "ebs"
    architecture        = "arm64"
  }
  owners      = ["099720109477"]
  most_recent = true
}

data "amazon-ami" "debian_x86" {
  filters = {
    virtualization-type = "hvm"
    name                = "debian-11-amd64-*"
    root-device-type    = "ebs"
    architecture        = "x86_64"
  }
  owners      = ["136693071363"]
  most_recent = true
}

data "amazon-ami" "debian_arm" {
  filters = {
    virtualization-type = "hvm"
    name                = "debian-11-amd64-*"
    root-device-type    = "ebs"
    architecture        = "arm64"
  }
  owners      = ["136693071363"]
  most_recent = true
}

data "amazon-ami" "redhat_x86" {
  filters = {
    virtualization-type = "hvm"
    name                = "RHEL-9.0.0_HVM-20221027-x86_64-1-Hourly2-GP2"
    root-device-type    = "ebs"
    architecture        = "x86_64"
  }
  owners      = ["309956199498"]
  most_recent = true
}

data "amazon-ami" "redhat_arm" {
  filters = {
    virtualization-type = "hvm"
    name                = "RHEL-9.1.0_HVM-20221101-arm64-2-Hourly2-GP2"
    root-device-type    = "ebs"
    architecture        = "arm64"
  }
  owners      = ["309956199498"]
  most_recent = true
}


source "amazon-ebs" "ubuntu_x86" {
  ami_name      = "nixtune-ubuntu-x86"
  instance_type = "t2.micro"
  region        = "us-east-1"
  source_ami    = data.amazon-ami.ubuntu_x86.id
  ssh_username  = "ubuntu"
}

source "amazon-ebs" "ubuntu_arm" {
  ami_name      = "nixtune-ubuntu-arm64"
  instance_type = "t2.micro"
  region        = "us-east-1"
  source_ami    = data.amazon-ami.ubuntu_arm.id
  ssh_username  = "ubuntu"
}

source "amazon-ebs" "debian_x86" {
  ami_name      = "nixtune-debian-x86"
  instance_type = "t2.micro"
  region        = "us-east-1"
  source_ami    = data.amazon-ami.debian_x86.id
  ssh_username  = "admin"
}

source "amazon-ebs" "debian_arm" {
  ami_name      = "nixtune-debian-arm64"
  instance_type = "t2.micro"
  region        = "us-east-1"
  source_ami    = data.amazon-ami.debian_arm.id
  ssh_username  = "admin"
}

source "amazon-ebs" "redhat_x86" {
  ami_name      = "nixtune-redhat-x86"
  instance_type = "t2.micro"
  region        = "us-east-1"
  source_ami    = data.amazon-ami.redhat_x86.id
  ssh_username  = "ec2-user"
}

source "amazon-ebs" "redhat_arm" {
  ami_name      = "nixtune-redhat-arm64"
  instance_type = "t2.micro"
  region        = "us-east-1"
  source_ami    = data.amazon-ami.redhat_arm.id
  ssh_username  = "ec2-user"
}

build {
  sources = [
    "source.amazon-ebs.ubuntu_x86",
    "source.amazon-ebs.ubuntu_arm",
    "source.amazon-ebs.debian_x86",
    "source.amazon-ebs.debian_arm",
    "source.amazon-ebs.redhat_x86",
    "source.amazon-ebs.redhat_arm"
  ]

  provisioner "shell" {
    script = "./setup.sh"
  }
}
