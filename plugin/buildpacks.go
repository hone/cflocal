package plugin

import "github.com/sclevine/forge"

var SystemBuildpacks forge.Buildpacks = []forge.Buildpack{
	{
		Name:       "ruby_buildpack",
		URL:        "https://codon-buildpacks.s3.amazonaws.com/buildpacks/heroku/ruby.tgz",
		VersionURL: "https://raw.githubusercontent.com/sclevine/cflocal-data/master/versions/ruby-buildpack",
	},
	{
		Name:       "java_buildpack",
		URL:        "https://codon-buildpacks.s3.amazonaws.com/buildpacks/heroku/java.tgz",
		VersionURL: "https://raw.githubusercontent.com/sclevine/cflocal-data/master/versions/java-buildpack",
	},
	{
		Name:       "go_buildpack",
		URL:        "https://codon-buildpacks.s3.amazonaws.com/buildpacks/heroku/go.tgz",
		VersionURL: "https://raw.githubusercontent.com/sclevine/cflocal-data/master/versions/go-buildpack",
	},
	{
		Name:       "python_buildpack",
		URL:        "https://codon-buildpacks.s3.amazonaws.com/buildpacks/heroku/python.tgz",
		VersionURL: "https://raw.githubusercontent.com/sclevine/cflocal-data/master/versions/python-buildpack",
	},
	{
		Name:       "php_buildpack",
		URL:        "https://codon-buildpacks.s3.amazonaws.com/buildpacks/heroku/php.tgz",
		VersionURL: "https://raw.githubusercontent.com/sclevine/cflocal-data/master/versions/php-buildpack",
	},
	{
		Name:       "scala_buildpack",
		URL:        "https://codon-buildpacks.s3.amazonaws.com/buildpacks/heroku/scala.tgz",
		VersionURL: "https://raw.githubusercontent.com/sclevine/cflocal-data/master/versions/binary-buildpack",
	},
	{
		Name:       "gradle_buildpack",
		URL:        "https://codon-buildpacks.s3.amazonaws.com/buildpacks/heroku/gradle.tgz",
		VersionURL: "https://raw.githubusercontent.com/sclevine/cflocal-data/master/versions/binary-buildpack",
	},
	{
		Name:       "nodejs_buildpack",
		URL:        "https://codon-buildpacks.s3.amazonaws.com/buildpacks/heroku/nodejs.tgz",
		VersionURL: "https://raw.githubusercontent.com/sclevine/cflocal-data/master/versions/nodejs-buildpack",
	},
}
