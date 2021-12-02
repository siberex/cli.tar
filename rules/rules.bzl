def _pkg_tar_impl(ctx):
    out = ctx.actions.declare_file(ctx.label.name + ".tar")

    args = ctx.actions.args()
    if ctx.attr.package_dir:
        args.add("--dir", ctx.attr.package_dir)
    args.add(out.path)
    args.add_all(ctx.files.srcs)
    args.use_param_file("@%s")

    ctx.actions.run(
        inputs = ctx.files.srcs,
        outputs = [out],
        executable = ctx.executable._tar_bin,
        arguments = [args],
        use_default_shell_env = True,
    )
    return [DefaultInfo(
        files = depset([out]),
        runfiles = ctx.runfiles(files = [out]),
    )]

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
        runfiles = ctx.runfiles(files = [ctx.file.file]),
    )]

pkg_tar = rule(
    doc = """Create tarball archive from provided input files""",
    implementation = _pkg_tar_impl,
    attrs = {
        "package_dir": attr.string(default = ""),
        "srcs": attr.label_list(
            allow_files = True,
            mandatory = True,
        ),
        # Implicit dependencies.
        "_tar_bin": attr.label(
            default = Label("//cmd/tar"),
            cfg = "exec",
            executable = True,
        ),
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
    executable = True,
)
