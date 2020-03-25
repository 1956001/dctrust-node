version=`build/kvant version`
echo $version

git add version
git add bin_name
git add release
git commit -m "$version"
git push