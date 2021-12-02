def _pkg_tar_impl(ctx):
    out = ctx.name + ".tar"
    pass

def _file_size_impl(ctx):
    # stat -f%z
    # stat --format="%s"
    # ls | awk ?

    pass

pkg_tar = rule(
    implementation = _pkg_tar_impl,
)

file_size = rule(
    implementation = _file_size_impl,
)
