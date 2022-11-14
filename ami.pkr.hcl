source "amazon-ebs" "ubuntu" {
  ami_name      = "ubuntu"
  instance_type = "t2.micro"
  region        = "us-east-1"
  source_ami    = "${var.amis["ubuntu"]}"
  ssh_username  = "ubuntu"
}

source "amazon-ebs" "debian" {
  ami_name      = "debian"
  instance_type = "t2.micro"
  region        = "us-east-1"
  source_ami    = "${var.amis["debian"]}"
  ssh_username  = "admin"
}

source "amazon-ebs" "redhat" {
  ami_name      = "redhat"
  instance_type = "t2.micro"
  region        = "us-east-1"
  source_ami    = "${var.amis["redhat"]}"
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
