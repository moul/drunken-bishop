GOPKG ?=	moul.io/drunken-bishop
DOCKER_IMAGE ?=	moul/drunken-bishop
GOBINS ?=	.
NPM_PACKAGES ?=	.

all: test install

-include rules.mk
