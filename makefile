GO := go

all: im-build user-build api-build file-build

im-build:
	cd im && $(GO) build

user-build:
	cd logic-user && $(GO) build

api-build:
	cd api-gateway && $(GO) build

file-build:
	cd file-server && $(GO) build
