PWD := ${CURDIR}
DOCKER_IMAGE_NAME := "wwwpc"
CONTAINER_REMOVE_FLAG := "--rm"
BASE_URL := "https://tamadalab.github.io/kani"
HUGO_THEME := "cayman-hugo-theme"
JOJOMI_VERSION := 0.69.0

build:
	docker run ${CONTAINER_REMOVE_FLAG} --name ${DOCKER_IMAGE_NAME}_build -v "${PWD}":/src -v ${PWD}/public:/output -e HUGO_THEME=$(HUGO_THEME) -e HUGO_BASEURL=${BASE_URL} jojomi/hugo:${JOJOMI_VERSION}
	rm public/favicon* public/apple-touch-icon*

start:
	docker run ${CONTAINER_REMOVE_FLAG} -d --name ${DOCKER_IMAGE_NAME} -p 1313:1313 -v "${PWD}":/src -v "$(PWD)"/public:/output -e HUGO_THEME=$(HUGO_THEME) -e HUGO_WATCH="true" -e HUGO_BASEURL=${BASE_URL} jojomi/hugo:${JOJOMI_VERSION}

stop:
	docker stop ${DOCKER_IMAGE_NAME}
