job "fikryfahrezy-los-inmem" {
  datacenters = ["dc1"]

  group "webserver" {
    network {
      port "http" {
        to = 4000
      }
    }

    task "los" {
      driver = "docker"

      config {
        image = "yuuuka111/fikryfahrezy-los-inmem"

        ports = ["http"]
      }

      resources {
        cpu    = 256
        memory = 256
      }
    }
  }
}
