variable "amis" {
  type = map(string)
  default = {
    "debian" = "ami-09a41e26df464c548"
    "redhat" = "ami-05723c3b9cf4bf4ff"
  }
}