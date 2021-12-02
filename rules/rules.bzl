def _pkg_tar_impl(ctx):
    out = ctx.label.name + ".tar"

    ctx.actions.run(
        inputs = ctx.files.srcs,
        outputs = [],
        executable = ctx.executable._tar_bin,
    )

def _file_size_impl(ctx):
    # stat -f%z rules/1.txt
    # stat --format="%s" rules/1.txt
    # ls -ln rules/1.txt | awk '{print $5}'
    # wc -c < rules/1.txt | awk '{print $1}'

    ctx.actions.run_shell(
        outputs = [],
        command = "echo 1"
    )

pkg_tar = rule(
    doc = """Create tarball archive from provided input files""",
    implementation = _pkg_tar_impl,
    attrs = {
        "package_dir": attr.string(default = "/"),
        "srcs": attr.label_list(allow_files = True),
        # Implicit dependencies.
        "_tar_bin": attr.label(
            default = Label("//cmd/tar"),
            cfg = "exec",
            executable = True,
            allow_files = True,
        ),
        "out": attr.output(doc = "The generated file"),
    },
)

file_size = rule(
    doc = """Print out file size in bytes for provided file""",
    implementation = _file_size_impl,
    attrs = {
        "file": attr.label(
            doc = "Input file to get byte size for",
            mandatory = True,
            allow_single_file = True,
        ),
    },
    executable = True
)
