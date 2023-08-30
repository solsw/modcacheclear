# modcacheclear

**modcacheclear** tool clears Go [module cache](https://go.dev/ref/mod#module-cache).

**modcacheclear** operates on modules from the [module cache](https://go.dev/ref/mod#module-cache) which have [go.mod](https://go.dev/ref/mod#go-mod-file) file and valid [SemVer](https://semver.org/#semantic-versioning-specification-semver) in the corresponding directory name.

**modcacheclear** without any flags prints [versions](https://go.dev/ref/mod#versions) of all modules (identified with [module paths](https://go.dev/ref/mod#module-path)) from the [module cache](https://go.dev/ref/mod#module-cache).

**modcacheclear** tool supports the following flags:

***-p*** regular expression pattern to match [module path](https://go.dev/ref/mod#module-path)

`modcacheclear -p ^github.com/solsw` prints [versions](https://go.dev/ref/mod#versions) of all modules starting with 'github.com/solsw' from the [module cache](https://go.dev/ref/mod#module-cache).

***-v*** number of module [versions](https://go.dev/ref/mod#versions) to keep

`modcacheclear -p github.com/solsw/modhelper -v 2` removes all but two [versions](https://go.dev/ref/mod#versions) with [maximum](https://semver.org/#spec-item-11) [SemVer](https://semver.org/#semantic-versioning-specification-semver) of 'github.com/solsw/modhelper' module from the [module cache](https://go.dev/ref/mod#module-cache).

`modcacheclear -v 1` removes all but one [version](https://go.dev/ref/mod#versions) with [maximum](https://semver.org/#spec-item-11) [SemVer](https://semver.org/#semantic-versioning-specification-semver) of all modules from the [module cache](https://go.dev/ref/mod#module-cache).

***-a*** remove all module [versions](https://go.dev/ref/mod#versions)

`modcacheclear -p github.com/solsw/modhelper -a` removes all [versions](https://go.dev/ref/mod#versions) of 'github.com/solsw/modhelper' module from the [module cache](https://go.dev/ref/mod#module-cache); if then the parent directory becomes empty, it is also deleted.

***-printskipped*** print skipped [module paths](https://go.dev/ref/mod#module-path)

***-nooutput*** produce no output except errors
