# goplus

Custom golang compiler with language improvements including C++ extensions

After several years of writing, maintaining, and deploying go code to various customers, clients, and employers, the language needs to be liberated with essential tool chain improvements.

The friendly and intelligent core go team at Google has limited bandwidth to implement features and those features require a long process of community discussion and review.  The intense scrutiny having to "convince" the committee verbally or in a community proposal would lead itself to never getting the features implemented.

GoPlus is a backward compatible golang compiler but with very focused improvements made in the runtime and compiler to allow for some much desired syntactical preferences.

## Supported Platforms

As of early June, 2024, the only "tested" platform is that of MacOS running a Linux Arm64 virtual machine.  Also published is an Linux-Amd64 build but it has actually not been tested on a live machine.  (Maybe I can do it with Rosetta later on the Mac)

Native MacOS does not support a reasonable environment for MacOS Virtual Machines so I will be delivering it as a native machine installation package when the next available time slot becomes available.

Native Windows AMD64 builds are still being finalized.  And the Windows ARM64 should become more available along with building the Delve Debugger for Windows ARM64 (super alpha right now)

You can always use the MacOS Arm64 built compiler to cross compile to any of the supported platforms, but the toolchain is really only built right now to run on a MacOS Arm64 device.

And the Visual Studio Code IDE is the only planned IDE that will be tested.  

## Installation

For developer simplicity, the compiler is shipped as a DevContainer.  So there should be very little effort required to get this custom compiler running in your environment.

On your MacOS M1,M2,M3,M4 computer, install the Microsoft Dev Containers Visual Studio code extension.  Then install Orbstack Docker Virtual Machine to process your docker containers.

Docker Desktop SHOULD work, but it has not been tested, only Orbstack community edition.

Useful brew command: 

`brew install orbstack`

## Testing

Once the necessary software has been installed, open the project in a container and give it some time to initialize on the first pass.

Once your IDE has settled out and all of the dependencies have auto-installed, you should be able to set a breakpoint in main.go and step through the syntax examples.

## Syntax improvements

The challenge is to imagine C++ like syntax improvements while maintaining backward compatibility with idiomatic golang.

This custom compiler adds class constructor semantics to allow for less developer toil and provide greater efficiency on complex structural composition.

You can emulate a class inheritance hierarchy while not getting tripped up on virtual methods or dispatch tables.  

Also added is Operator Overloading semantics.  You can define new operators and also define their behavior.  Included are both array [] and bracket {} syntax.  Look at the sample code for examples.

If you want to verify that the GoPlus compiler is active, enter "go plus" on the command line and it should confirm its identity.

## Bugs / Feature Requests

If you have a syntax construct that fails to compile, copy the code into a small snippet and create an issue: https://github.com/markdiener10/goplus/issues

I will enjoy reviewers coming up with syntactical use cases I did not think of.

And if for some reason, you would like to see a feature added to the compiler that you know will not be accepted by the larger community, feel free to open an issue and we can discuss it.

## Road Map

I have additional C++ related syntax improvements and I want to more fully automate my continuous integration and delivery pipeline so that new versions of GoPlus are merged with the latest release of the community golang toolchain.

And when I have some free cycles, native Windows and MacOS packages will be forthcoming.

Have a great day!