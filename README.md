# seeder

seeder is a developer tool that is used from a module.

The Modular nature of the architecture means that an seed.yaml can exist in any module
In the seed.yaml you declare other modules you are dependent on.

Why another packaging tool ?

- we have specialised needs for modularity.
- we can extend this for ops things like deployment and migrations by hookig into the sys-core Corba CLI, and adding extra meta data to the seed-yaml that tells us the IP address of the Server.
- Dev and CI needs OS specific dependencies in order to run
- go:generate for each module has many dependencies that are now .mk files, which is not sustainable for Multi OS builds.


## Constructs

Each modules needs a seed.yaml and a manifest.yaml

- seed.yaml defines all modules you need
- manifest.yaml defines all resources you use, so that the seeder can inspect and get them

## functions

deps

- manages dependencies
- this is typically various OS tools and any custom tools needed
- is OS aware.

git

- manages git operations for you. Only basic ones, and you can drop down to git if you need to egt fancy
- used by mods if in CI an extra git repo is needed

seed

- developer and ops tool for managing modules and their sub modules.
- inspects the seed.yaml to see all other Modules you use. 
- allows mono repos and single repos for Modules. 
- via git or local path, it retrieves the sub modules manifest.
    - respects git tags, so that versioned modules and apps can exist.
- from the manifest it retrieves any base config the app needs
    - works with uber config
    - routes
    - any other config
    - merges the config into a base.yml, ensurign you are up to date

manifest

- developer tool for generating everything you need to build modules.
- this wraps all the .mk code we are using
- in a module your go:generate can then just call "seeder gen X" to have seeder do the work for you.
    - you can pick and choose if you want to call seeder wrapper funcions for highly granular functions
- outputs the final manifest.yaml that seed tool will then use.


config

- merges many configs into one.
- this is used by the mod tool after is gets the manifest, grabs the config, and copies it up to your main.