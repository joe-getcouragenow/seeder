# dev

Now all we need are two things

1. Mage

mage, instead of make
- each repo just needs a mage.go, that references seed repo. 
- the golang runtime then will automatically run the code in the seed repo.

github action for mage:
https://github.com/magefile/mage-action

ex: https://github.com/magefile/awesome-mage

The main file we will need:
https://github.com/mysteriumnetwork/go-ci/blob/master/mage.go

https://github.com/microsoft/cobalt/blob/master/test-harness/magefile.go

proto: https://github.com/kraman/mage-helpers/blob/master/protobuf/protobuf.go
- very nice an clean

git: https://github.com/denkhaus/magelib/tree/master/git

project structure: https://github.com/voyages-sncf-technologies/mageproj
- looks really good
- makes you independent of Github too.


https://github.com/openthread/ot-ns/blob/master/script/pack-web
- go-bindata smart stuff
- https://github.com/openthread/ot-ns/blob/master/web/web.go
	- properly loads the flutter web AND has the improbable go proxy that runs as its own binay BUT is controllable.
	- script to install is here: https://github.com/openthread/ot-ns/blob/master/script/install-deps
- Everything in this lib is perfect for us

https://nakedsecurity.sophos.com/2020/02/03/google-launches-open-source-security-key-project-opensk/
- Use this !!
- nRF52840-dongle.
	- software for it: https://github.com/google/OpenSK
- and DICE for MASTER KEY ! https://www.crowdsupply.com/dicekeys/dicekeys
- They keep making NEW FIDO standards and so hence the hardware you guys from Yubikey is NOT updatable. Hence why this open approahc is way smarter
	- See: https://github.com/google/OpenSK/issues/106
- Will need WebAuthn and FIDO golang code integrated into our stack, so the hardware works with it.


Protobufs
- https://github.com/izumin5210/grapi
- in CMD has standard scaffolding commands that we can use from mage.
- we can refactor the code with this. The code is a mess right now.
- and it will set us up for the Mod layer being WASM, so we can run on cloudflare workers or our own version of it.
	- the sys layer is the data layer and will serve to the wasm layer.
- start to use protoc-gen-validate !!!
	- used here well: https://github.com/Zenithar/go-monorepo/search?q=protoc-gen-validate
		- gens for all languages 
- the mistake Alex make is to assume golang in sys-shared !!
	- validation cant happen in mod now
	- models cant be used by mod now.

https://github.com/Zenithar/go-repo-template
- mage is in : https://github.com/Zenithar/go-repo-template/tree/master/.mage
- used on this https://github.com/Zenithar/go-spotigraph
	- which uses https://github.com/Zenithar/go-pkg
- The archi looks like ours !
https://github.com/Zenithar/solid
- really good authz based on oid

Mono repo tools

- Flutter based but same for golang really
- Tells a main what others to run generate and lint on.
- Same goes for our Git tools, because we want to checking many git repos with one commit.
https://github.com/google/mono_repo.dart/blob/master/mono_repo/build.yaml
- Example of mono repo being used here: https://github.com/google/dart-neats
	- shows how useful it is !

i18n
- We need to adopt their approach
- https://medium.com/flutter/announcing-flutter-1-22-44f146009e5f
- https://docs.google.com/document/d/10e0saTfAv32OZLRmONy866vnaw0I2jwL8zukykpgWBc/edit


Widgets
https://github.com/google/flutter.widgets
- thir list of widgets are exactly the things we need
- Ability to scroll to a position for examle in a list.

2. Seeder

Use golang to seed a repo from the seeder repo.

SO from any repo, you can seed it.

https://github.com/go-make/make/tree/master/cmd/go-make
- this does it with golang templates