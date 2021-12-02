def _pkg_tar_impl(ctx):
    out = ctx.label.name + ".tar"

    ctx.actions.run(
        inputs = ctx.files.srcs,
        outputs = [],
        executable = ctx.executable._tar_bin,
    )

def _file_size_impl(ctx):
    executable = ctx.actions.declare_file(ctx.label.name)
    ctx.actions.expand_template(
        template = ctx.file._script_template,
        output = executable,
        substitutions = {
            "{INPUT}": ctx.file.file.path,
        },
        is_executable = True,
    )
    return [DefaultInfo(
        executable = executable,
        runfiles = ctx.runfiles(files = [ctx.file.file])
    )]

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
        "_script_template": attr.label(
            allow_single_file = True,
            default = "file_size.sh",
        ),
    },
    executable = True
)
