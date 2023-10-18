description
-----------

WIP

This project aims to demonstrate how to use bazel as a monorepo build system
to build and test applications and their corresponding oci images in separate
bazel sub modules.


1. The [./components/go-hello](./components/go-hello) contains a minimalist golang hello world project,
yet with enough code to require the introduction of the gazelle tool.

1. The [./oci-images/go-hello](./oci-images/go-hello) contains a module which shall build and test
an OCI image, with a dependency to the go binary produced by the [./components/go-hello](./components/go-hello)
module.


Questions
---------

1. How to run gazelle [./components/go-hello](./components/go-hello)

Is this the correct way to use gazelle for a submodule ?

``` sh
bazel run @go-hello//:gazelle -- update components/go-hello 
```


1. How to build [./components/go-hello](./components/go-hello)

Is this the correct way to build a sub module ?

``` sh
bazel build @go-hello
```

1. How to build [./oci-images/go-hello](./oci-images/go-hello)

Is this the correct build and test oci images ?

``` sh
bazel test @go-hello-image
```

1. Cross modules dependencies

Why is it necessary to declare go-hello in both [./MODULE.bazel](./MODULE.bazel) and [./oci-images/go-hello/MODULE.bazel](./oci-images/go-hello/MODULE.bazel) to build the oci image ? 
When declaring only in  [./oci-images/go-hello/MODULE.bazel](./oci-images/go-hello/MODULE.bazel), running the following command fails.

``` sh
bazel test @go-hello-image
Loading: 
ERROR: Error computing the main repository mapping: module not found in registries: go-hello@_
```


1. How to build &  test all the modules ?

Is it possible to build all modules without having to run one command per module ?
