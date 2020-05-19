# Version Monitor

A system to keep track of and monitor software version in use by a team or organization. The basic idea is having a overview which software version are we running and what is the latest stable version.

## Getting started

Create a configuration file named `version-monitor.yaml`, use `example.yaml` as starting point. The system supports different adapters for information gathering. See [adapters](./pkg/adapters/README.md) for details.

Run the system like:

```shell
git clone https://github.com/fielmann-ag/version-monitor.git
cd version-monitor
go build
cp example.yaml version-monitor.yaml
./version-monitor
```

This will spawn a http-server on port 8080, check URL [http://localhost:8080/](http://localhost:8080/) on localhost.

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Authors

- Alex Klinkert [aklinkert](https://github.com/aklinkert)
- Heinrich Grotjohann [heinrichgrt](https://github.com/heinrichgrt)
- Roman Messer [rmmsr](https://github.com/rmmsr)

# License

```
Copyright 2020 Fielmann AG

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
