;(async function() {
  let connect = (await import("@dagger.io/dagger")).connect
  connect(async (client) => {
	cache = process.env.CACHE
	await buildJenkins(client.pipeline("no_cache"), cache).sync()
  }, {LogOutput: process.stdout})
})()

// Example pipeline to build Jenkins from a git ref
function buildJenkins(client, cache_level) {
  repo = "https://github.com/jenkinsci/jenkins"

  source = client.git(repo).branch("master").tree()
  maven = client.cacheVolume("maven")

  builder = client.container().pipeline("Build Jenkins")
  .from("eclipse-temurin:17-focal")

  if(cache_level != "all" && cache_level != "layers") {
    builder = builder.withEnvVariable("BUST", `${Date.now()}`)
  }
  builder = builder.withExec(["apt-get", "update"])
  .withExec(["apt-get", "install", "-y", "git"])
  .withWorkdir("/tmp")
  .withExec(["wget", "https://dlcdn.apache.org/maven/maven-3/3.9.4/binaries/apache-maven-3.9.4-bin.tar.gz"])
  .withExec(["tar", "xvf", "apache-maven-3.9.4-bin.tar.gz"])
  .withExec(["mv", "apache-maven-3.9.4", "/opt/"])
  .withEnvVariable("M2_HOME", "/opt/apache-maven-3.9.4")
  .withDirectory("/src", source)
  .withWorkdir("/src")
  if(cache_level != "none" && cache_level != "layers") {
    builder = builder.withMountedCache("/root/.m2", maven)
  }
  builder = builder.withExec(["sh", "-c", "/opt/apache-maven-3.9.4/bin/mvn -am -pl war,bom -Pquick-build clean install"])

  return builder
}
