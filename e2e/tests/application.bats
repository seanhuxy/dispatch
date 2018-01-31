#!/usr/bin/env bats

set -o pipefail

load helpers
load variables

@test "Application creation" {

    run dispatch get application
    assert_success

    # create applications "foo" and "bar"
    run dispatch create application foo-app
    assert_success

    run dispatch create app bar-app
    assert_success

    run_with_retry "dispatch get app foo-app --json | jq -r .status" "READY" 4 5
    run_with_retry "dispatch get app bar-app --json | jq -r .status" "READY" 4 5
}

@test "Base image creation" {

    run dispatch get base-image
    assert_success

    # Create base image "base-nodejs6"
    run dispatch create base-image base-nodejs6 $DOCKER_REGISTRY/$BASE_IMAGE_NODEJS6 --language nodejs6 --public
    assert_success

    # Ensure starting status is "INITIALIZED". Wait 20 seconds for status "READY"
    run_with_retry "dispatch get base-image base-nodejs6 --json | jq -r .status" "INITIALIZED" 1 0
    run_with_retry "dispatch get base-image base-nodejs6 --json | jq -r .status" "READY" 4 5
}

@test "Image creation" {

    run dispatch get image
    assert_success

    run dispatch create image foo-app-image base-nodejs6 --application foo-app
    assert_success
    run dispatch create image bar-app-image base-nodejs6 --application bar-app
    assert_success

    run_with_retry "dispatch get image foo-app-image --application foo-app --json | jq -r .status" "READY" 4 5
    run_with_retry "dispatch get image bar-app-image --application bar-app --json | jq -r .status" "READY" 4 5
}

@test "Secret creation" {
    run dispatch create secret foo-app-open-sesame -a foo-app ${DISPATCH_ROOT}/examples/nodejs6/secret.json
    echo_to_log
    assert_success
}

@test "Function creation" {
    run dispatch create function nodejs6 foo-app-func ${DISPATCH_ROOT}/examples/nodejs6/i-have-a-secret.js --secret open-sesame
    echo_to_log
    assert_success

    run_with_retry "dispatch get function foo-app-func --json | jq -r .status" "READY" 8 5
}


@test "API creation" {

}

@test "Event creation" {

}

