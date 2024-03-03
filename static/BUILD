load("@rules_pkg//:pkg.bzl", "pkg_tar")

pkg_tar(
    name = "static_layer",
    srcs = [
        ":styles.css",
    ],
    package_dir = "static",
    visibility = ["//visibility:public"],
)
