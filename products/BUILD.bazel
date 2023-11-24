load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "products_lib",
    srcs = ["repo.go"],
    importpath = "github.com/testcontainers-community/testcontainers-bazel-demo/products",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_jackc_pgx_v5//:go_default_library",
    ],
)

go_test(
    name = "products_lib_test",
    srcs = ["repo_test.go"],
    embed = [":products_lib"],
    deps = [
        "@com_github_stretchr_testify//assert",
        "@com_github_testcontainers_testcontainers_go//:testcontainers-go",
        "@com_github_testcontainers_testcontainers_go//wait",
        "@com_github_testcontainers_testcontainers_go_modules_postgres//:go_default_library",
    ],
)
