package build

import (
	"fastpm/src/downloader"
	"fastpm/src/extracter"
	"fastpm/src/logger"
	"fastpm/src/symlink"
)

func InstallPackages(pkgs []downloader.Package, target string, concurrency int) error {
    downloader.EnsureDir("./cache")
    symlink.EnsureDir(target)
    downloader.DownloadAll(pkgs, concurrency)
    for _, p := range pkgs {
        extractor.ExtractTarGz("./cache/"+p.Name, target)
        symlink.Create("./node_modules/"+p.Name, target+"/"+p.Name)
        logger.Info("Installed " + p.Name)
    }
    return nil
}
