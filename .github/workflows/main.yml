workflow "Build workflow" {
  on = "push"
  resolves = ["push"]
}

action "build" {
  uses = "actions/docker/cli@master"
  args = "build -t go-static-web-docker ."
}

action "login" {
  uses = "actions/docker/login@master"
  needs = ["build"]
  secrets = ["DOCKER_USERNAME", "DOCKER_PASSWORD"]
}

action "tag" {
  uses = "actions/docker/cli@master"
  needs = ["login"]
  args = "tag go-static-web-docker jackytck/go-static-web-docker:v0.0.1"
}

action "push" {
  uses = "actions/docker/cli@master"
  needs = ["tag"]
  args = "push jackytck/go-static-web-docker:v0.0.1"
}
