group "default" {
	targets = [ "build" ]
}

target "deps" {
	context = "."
	dockerfile = ".github/Dockerfile"
	target = "deps"
}

target "build" {
	name = app
	context = "."
	dockerfile = ".github/Dockerfile"
	matrix = {
		app = [ "scalar" ]
	}
	tags = [ "ghcr.io/clementd64/x/${app}:latest" ]
	args = {
		APP = app
	}
	contexts = {
		deps = "target:deps"
	}
}
