## Project Description

SkyBind(TM) Console extends Portainer to create a simple, general-purpose console for managing micro-sites (small purpose-built websites).

## Project Scope

### 1) VM Management:
 
Sample features: Start VM, Restart VM, Stop VM.

NOTE: Simplicity is a core value so a feature-complete cross-cloud console is out of scope for this project. Cloud-specific consoles should continue to be relied on for features that are out-of-scope for the SkyBind(TM) Console. 

### 2) Tenant / Site Management:

 Sample features: Create | Edit | Delete Tenant Site.

## Operating Vision

We believe in contributing to the growth of great open source projects rather than diluting the open source pool by forking a totally new initiative. If the Portainer Management team like the added SkyBind(TM) Console functionality and see it as in-scope for the Portainer vision then it will be backported to Portainer and this project will evolve to operate as a new Portainer feature incubator. In either case, this project is intending to be a long-term incremental investment in an open source Micro-site Console.

## Repository Management

This repository will be managed as a two-way branch of the Portainer master branch to facilitate backporting of bug fixes and enhancements that the Portainer Management team would like to add to Portainer. To facilitate this bi-directionality, enhancements will be isolated to new files and new folders where feasible.

This repository will be synchronised with Portainer after every major release to the Portainer master branch.

NOTE1: This SkyBind(TM) fork of Portainer is synchronised with the Portainer master branch not the Portainer development branch so there may be new Portainer features visible to Portainer developers that are not visible to SkyBind(TM) Console developers.

NOTE2: The Portainer Build files & directories have not been retained in this fork. For simplicity our preference is to build without using Grunt and Yarn.

NOTE3: The Portainer API directory has not been retained in this fork at this time to simplify the initial build. The API will be added back in after release 1.0.

## Contributions

This project will not be open for contributions until the first major release which is expected to be 20 December 2019. Work-in-progress will be made visible on a caveat-emptor basis on the skybind-master branch until release 1.0 at which time a skybind-development branch will become the home for future updates.

## Licensing

SkyBind(TM) Console / Portainer is licensed under the zlib license. See [LICENSE](./LICENSE) for reference.

SkyBind(TM) Console / Portainer also contains the following code, which is licensed under the [MIT license](https://opensource.org/licenses/MIT):

- UI For Docker: Copyright (c) 2013-2016 Michael Crosby (crosbymichael.com), Kevan Ahlquist (kevanahlquist.com), Anthony Lapenna (portainer.io)

- rdash-angular: Copyright (c) [2014] [Elliot Hesp]
