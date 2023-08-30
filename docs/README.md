# modcacheclear

**modcacheclear** tool clears Go [module cache](https://go.dev/ref/mod#module-cache).

**modcacheclear** without any flags prints versions of all modules from the [module cache](https://go.dev/ref/mod#module-cache).

**modcacheclear** tool supports the following flags:

***-p*** regular expression pattern to match module path

`modcacheclear -p ^github.com/solsw` prints versions of all modules starting with 'github.com/solsw' from the [module cache](https://go.dev/ref/mod#module-cache).

***-a*** remove all module versions

`modcacheclear -p github.com/solsw/modhelper -a` removes all versions of 'github.com/solsw/modhelper' module from the [module cache](https://go.dev/ref/mod#module-cache); if the parent directory then becomes empty it is also deleted.

***-v*** number of module versions to keep

`modcacheclear -p github.com/solsw/modhelper -v 2` removes all but two versions with [maximum](https://semver.org/#spec-item-11) [SemVer](https://semver.org/#semantic-versioning-specification-semver) of 'github.com/solsw/modhelper' module from the [module cache](https://go.dev/ref/mod#module-cache).

`modcacheclear -v 1` removes all but one version with [maximum](https://semver.org/#spec-item-11) [SemVer](https://semver.org/#semantic-versioning-specification-semver) of all modules from the [module cache](https://go.dev/ref/mod#module-cache).

***-printskipped*** print skipped module paths
