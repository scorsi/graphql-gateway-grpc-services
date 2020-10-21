workspace(name = "prometer")

################################################################################
# Setup Bazel project
################################################################################

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

################################################################################
# Go dependencies
################################################################################

### RULES GO

RULES_GO_VERSION = "0.23.3"

RULES_GO_SHA256 = "a8d6b1b354d371a646d2f7927319974e0f9e52f73a2452d2b3877118169eb6bb"

http_archive(
    name = "io_bazel_rules_go",
    sha256 = RULES_GO_SHA256,
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v{v}/rules_go-v{v}.tar.gz".format(v = RULES_GO_VERSION),
        "https://github.com/bazelbuild/rules_go/releases/download/v{v}/rules_go-v{v}.tar.gz".format(v = RULES_GO_VERSION),
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

### RULES GAZELLE

RULES_GAZELLE_VERSION = "0.21.1"

RULES_GAZELLE_SHA256 = "cdb02a887a7187ea4d5a27452311a75ed8637379a1287d8eeb952138ea485f7d"

http_archive(
    name = "bazel_gazelle",
    sha256 = RULES_GAZELLE_SHA256,
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v{v}/bazel-gazelle-v{v}.tar.gz".format(v = RULES_GAZELLE_VERSION),
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v{v}/bazel-gazelle-v{v}.tar.gz".format(v = RULES_GAZELLE_VERSION),
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

load("//:DEPENDENCIES.bzl", "go_repositories")

# gazelle:repository_macro DEPENDENCIES.bzl%go_repositories
go_repositories()

################################################################################
# Proto dependencies
################################################################################

### RULES PROTO

http_archive(
    name = "rules_proto",
    sha256 = "602e7161d9195e50246177e7c55b2f39950a9cf7366f74ed5f22fd45750cd208",
    strip_prefix = "rules_proto-97d8af4dc474595af3900dd85cb3a29ad28cc313",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_proto/archive/97d8af4dc474595af3900dd85cb3a29ad28cc313.tar.gz",
        "https://github.com/bazelbuild/rules_proto/archive/97d8af4dc474595af3900dd85cb3a29ad28cc313.tar.gz",
    ],
)

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

################################################################################
# Jsonnet dependencies
################################################################################

### RULES JSONNET

http_archive(
    name = "io_bazel_rules_jsonnet",
    sha256 = "7f51f859035cd98bcf4f70dedaeaca47fe9fbae6b199882c516d67df416505da",
    strip_prefix = "rules_jsonnet-0.3.0",
    urls = ["https://github.com/bazelbuild/rules_jsonnet/archive/0.3.0.tar.gz"],
)
load("@io_bazel_rules_jsonnet//jsonnet:jsonnet.bzl", "jsonnet_repositories")

jsonnet_repositories()

load("@jsonnet_go//bazel:repositories.bzl", "jsonnet_go_repositories")

jsonnet_go_repositories()

load("@jsonnet_go//bazel:deps.bzl", "jsonnet_go_dependencies")

jsonnet_go_dependencies()

################################################################################
# Js dependencies
################################################################################

### RULES NODEJS

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
http_archive(
    name = "build_bazel_rules_nodejs",
    sha256 = "64a71a64ac58b8969bb19b1c9258a973b6433913e958964da698943fb5521d98",
    urls = ["https://github.com/bazelbuild/rules_nodejs/releases/download/2.2.1/rules_nodejs-2.2.1.tar.gz"],
)

load("@build_bazel_rules_nodejs//:index.bzl", "node_repositories")

node_repositories(package_json = ["//:package.json"])

################################################################################
# Kubernetes & Docker dependencies
################################################################################

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "4521794f0fba2e20f3bf15846ab5e01d5332e587e9ce81629c7f96c793bb7036",
    strip_prefix = "rules_docker-0.14.4",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.14.4/rules_docker-v0.14.4.tar.gz"],
)

load("@io_bazel_rules_docker//repositories:repositories.bzl", container_repositories = "repositories")

container_repositories()

load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")

container_deps()

load("@io_bazel_rules_docker//repositories:pip_repositories.bzl", "pip_deps")

pip_deps()

load("@io_bazel_rules_docker//go:image.bzl", go_image_repos = "repositories")

go_image_repos()

http_archive(
    name = "io_bazel_rules_k8s",
    sha256 = "d91aeb17bbc619e649f8d32b65d9a8327e5404f451be196990e13f5b7e2d17bb",
    strip_prefix = "rules_k8s-0.4",
    urls = ["https://github.com/bazelbuild/rules_k8s/releases/download/v0.4/rules_k8s-v0.4.tar.gz"],
)

load("@io_bazel_rules_k8s//k8s:k8s.bzl", "k8s_repositories")

k8s_repositories()

load("@io_bazel_rules_k8s//k8s:k8s_go_deps.bzl", k8s_go_deps = "deps")

k8s_go_deps()
