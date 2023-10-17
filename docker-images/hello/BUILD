load("@aspect_bazel_lib//lib:transitions.bzl", "platform_transition_filegroup")
load("@container_structure_test//:defs.bzl", "container_structure_test")
load("@rules_oci//oci:defs.bzl", "oci_image", "oci_tarball")
load("@rules_pkg//:pkg.bzl", "pkg_tar")

# Put app go_binary into a tar layer.
pkg_tar(
    name = "hello_layer",
    srcs = ["@go-hello:binary"],
    # If the binary depends on RUNFILES, uncomment the attribute below.
    # include_runfiles = True
)

oci_image(
    name = "image",
    base = "@distroless_base",
    entrypoint = ["/hello"],
    tars = [":hello_layer"],
)

# This is the target that should be released to the target platform
platform_transition_filegroup(
    name = "transitioned_image",
    srcs = [":image"],
    target_platform = select({
        "@platforms//cpu:arm64": "@rules_go//go/toolchain:linux_arm64",
        "@platforms//cpu:x86_64": "@rules_go//go/toolchain:linux_amd64",
    }),
)

oci_tarball(
    name = "tarball",
    # Use the image built for the exec platform rather than the target platform
    image = ":image",
    repo_tags = ["bzl-poc/go-hello:latest"],
)

container_structure_test(
    name = "test",
    configs = ["test.yaml"],
    # Use the image built for the exec platform rather than the target platform
    image = ":image",
)
