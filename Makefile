GOPKG ?=	moul.io/drunken-bishop
DOCKER_IMAGE ?=	moul/drunken-bishop
GOBINS ?=	.

all: test install

-include rules.mk
