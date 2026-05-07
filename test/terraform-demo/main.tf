terraform {
  required_providers {
    random = {
      source = "hashicorp/random"
    }

    local = {
      source = "hashicorp/local"
    }
  }
}

resource "random_pet" "name" {
  length = 2
}

resource "random_string" "version" {
  length  = 8
  special = false
}

resource "local_file" "example" {
  filename = "${path.module}/hello.txt"

  content = <<TEXT
pet     = ${random_pet.name.id}
version = ${random_string.version.result}
time    = ${timestamp()}
TEXT
}

output "pet_name" {
  value = random_pet.name.id
}
