load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/testcontainers-community/testcontainers-bazel-demo/products
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.work",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)
