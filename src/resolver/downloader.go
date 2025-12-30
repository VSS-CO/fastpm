package resolver

import "fastpm/src/downloader"

func Download(pkg downloader.Package) error {
    return downloader.DownloadWithProgress(pkg.URL, "./cache/"+pkg.Name)
}
