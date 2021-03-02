provider "google" {
  credentials = file("keys/devs-laboratory-2e33003dec33.json")
  project     = "devs-laboratory"
  region      = "asia-southeast1-a"
}

resource "google_storage_bucket" "backend_storage" {
  name          = "iac-backend-gcs-tfstate-laboratory"
  storage_class = "STANDARD"
  location      = "ASIA-SOUTHEAST1"
}

