terraform {
  required_version = ">= 1.5.0"

  required_providers {
    random = {
      source  = "hashicorp/random"
      version = "~> 3.6"
    }

    local = {
      source  = "hashicorp/local"
      version = "~> 2.5"
    }
  }
}

provider "random" {}
provider "local" {}

resource "random_pet" "name" {
  length = 2
}

resource "local_file" "example" {
  filename = "${path.module}/generated.txt"

  content = <<EOT
tforge demo
pet = ${random_pet.name.id}
EOT
}

output "pet_name" {
  value = random_pet.name.id
}
