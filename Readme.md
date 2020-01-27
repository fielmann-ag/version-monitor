# Version Monitor
A system to keep track of and monitor software version in use by a team or organization. The basic idea is having a overview which software version are we running and what is the latest stable.  

## Getting started
Create a configuration file. The default is `configuration.yaml`, you can start from `example.yaml` as starting point. The system supports different adapters for information gathering. See [adapters](./pkg/adapters/README.md) for details. 


Run the system like:
```
git clone 
cd version-monitor
Config=example.yaml ./version-monitor
```
This will spawn a http-server on port 8080. 
Check [here](http://localhost:8080/) your system on localhost. 


## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.


## Authors
Alex Pinnecke  [aklinkert](https://github.com/aklinkert)
Heinrich Grotjohann [heinrichgrt](https://github.com/heinrichgrt) 
Roman Messer [rmmsr](https://github.com/rmmsr)