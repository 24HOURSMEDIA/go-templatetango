#!/usr/bin/env bash
set -e

ME_DIR="$(realpath "$(dirname "$0")")"
PROJECT_DIR="${PROJECT_DIR:-$(realpath "$ME_DIR/..")}"

set -a
source "${PROJECT_DIR}/build_dockerhub.env"
set +a

check_env() {
    local var_name="$1"
    local var_value=$(eval echo "\$$var_name")
    if [ -z "$var_value" ]; then
        echo "Error: $var_name is undefined or empty"
        exit 1
    fi
}

check_env "VERSION"
check_env "IMAGE_NAME"
check_env "PLATFORMS"
check_env "BUILDER_NAME"

echo "Building and pushing image $IMAGE_NAME:$VERSION for platforms $PLATFORMS"
# ask to proceed or exit with code 0
read -p "Continue? [y/N] " -n 1 -r
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "Aborting"
    exit 0
fi

(
  cd ${PROJECT_DIR}

  # Set the desired image name and tag
  AUTO_IMAGE_TAG=${VERSION}
  IMAGE_TAG=${IMAGE_TAG:-"${AUTO_IMAGE_TAG}"}


  # Check if the builder already exists, and remove it if it does
  if docker buildx inspect $BUILDER_NAME > /dev/null 2>&1; then
      echo "Builder $BUILDER_NAME already exists, removing..."
      docker buildx rm $BUILDER_NAME
  fi

  # Create a new builder instance
  docker buildx create --name ${BUILDER_NAME} --use

  # Start up the builder instance
  docker buildx inspect ${BUILDER_NAME} --bootstrap

  # Build and push the image for multiple architectures
  docker buildx build --platform ${PLATFORMS} \
      -t $IMAGE_NAME:$IMAGE_TAG \
      --push \
      .

  # Clean up the builder instance
  docker buildx rm ${BUILDER_NAME}
)